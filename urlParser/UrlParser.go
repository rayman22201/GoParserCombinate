// A urlParser the uses a Parser Combinator to do the dirty work
// @TODO: add more detailed error messages into the parser.
package urlParser

import (
	"github.com/Apartments24-7/goSpider/parserCombinator"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type urlList struct {
	UniqueURLs   []URLStruct
	possibleURLs []string
	uniqueMap    map[string]bool
	lock         sync.RWMutex
	wg           sync.WaitGroup
}

type URLParseError struct {
	Msg string
}

func (e *URLParseError) Error() string {
	return e.Msg
}

type URLStruct struct {
	Protocol string
	Port     string
	Base     []string
	Path     string
	Filename string
	FileExt  string
	Query    string
	Anchor   string
}

func (url URLStruct) DebugString() string {
	return strings.Join([]string{
		"Protocol: " + url.Protocol,
		"Base: " + strings.Join(url.Base, ", "),
		"Port: " + url.Port,
		"Path: " + url.Path,
		"Filename: " + url.Filename,
		"File Extension: " + url.FileExt,
		"Anchor Tag: " + url.Anchor,
		"Query String: " + url.Query,
	}, "\n")
}

func (url URLStruct) String() string {
	var protocol, port, fileExt, anchor, query string
	switch url.Protocol {
	case "HTTP":
		protocol = "http://"
	case "HTTPS":
		protocol = "https://"
	case "FTP":
		protocol = "ftp://"
	}
	if url.Port != "" {
		port = ":" + url.Port
	}
	if url.FileExt != "" {
		fileExt = "." + url.FileExt
	}
	if url.Anchor != "" {
		anchor = "#" + url.Anchor
	}
	if url.Query != "" {
		query = "?" + url.Query
	}
	return protocol + strings.Join(url.Base, ".") + port + url.Path + url.Filename + fileExt + anchor + query
}

func RelativeURLToAbsolute(relative URLStruct, domain URLStruct) URLStruct {
	// Basically, merge the two structs, with the protocol, domain, and port comming form the
	// "domain" struct and everythinbg else from the "relative" struct.

	var path string
	path = relative.Path
	if strings.Index(relative.Path, ".") == 0 {
		path = domain.Path + relative.Path
	}

	return URLStruct{
		Protocol : domain.Protocol,
		Port : domain.Port,
		Base : domain.Base,
		Path : path,
		Filename : relative.Filename,
		FileExt : relative.FileExt,
		Query : relative.Query,
		Anchor : relative.Anchor,
	}
}

var containsURL = regexp.MustCompile(`(\'|\"|\()([\w\-?:\/&*\.=;,\%\[\]#]( )?)+(\'|\"|\))`)

func ParseBody(body string, whiteList []string) []URLStruct {
	parURLList := urlList{
		uniqueMap: make(map[string]bool),
	}

	parURLList.possibleURLs = containsURL.FindAllString(body, -1)

	parURLList.wg.Add(len(parURLList.possibleURLs))
	for _, possibleURL := range parURLList.possibleURLs {
		go func(testString string) {
			defer parURLList.wg.Done()
			// split on spaces to allow for lists of urls
			// Hangles cases like the HTMLL5 srcset tag
			possibleURLSplit := strings.Split(testString, " ")

			for _, possibleSubURL := range possibleURLSplit {
				url, err := ParseValidURL(possibleSubURL, whiteList)
				if err == nil {
					parURLList.lock.RLock()
					_, ok := parURLList.uniqueMap[url.String()]
					parURLList.lock.RUnlock()
					if !ok {
						parURLList.lock.Lock()
						// CAS semantics. We have to double check the value
						// because it could have changed since we unlocked the Reader lock.
						_, ok := parURLList.uniqueMap[url.String()]
						if !ok {
							parURLList.UniqueURLs = append(parURLList.UniqueURLs, *url)
							parURLList.uniqueMap[url.String()] = true
						}
						parURLList.lock.Unlock()
					}
				}
			}
		}(possibleURL)
	}
	parURLList.wg.Wait()

	return parURLList.UniqueURLs
}

func ParseValidURL(possibleURL string, whiteList []string) (*URLStruct, error) {
	possibleURL = strings.Trim(possibleURL, "'")
	possibleURL = strings.Trim(possibleURL, "\"")
	possibleURL = strings.Trim(possibleURL, "(")
	possibleURL = strings.Trim(possibleURL, ")")
	result, _ := urlParser(possibleURL)

	if result == nil {
		return nil, &URLParseError{"Unable to Parse a Valid URL from string: " + possibleURL}
	}
	url := result.(URLStruct)

	if !whiteListApproved(url.Base, whiteList) {
		return nil, &URLParseError{"Domain not found on the white list."}
	}

	return &url, nil

}

func whiteListApproved(urlSegments []string, whiteList []string) bool {
	if len(whiteList) > 0 && len(urlSegments) > 0 {
		// Using a linear search here. Could use a map for faster performance,
		// but I expect the white list to be very small,
		// and arrays provide a nicer interface for the "white list" concept.
		found := true
		for _, whiteURL := range whiteList {
			whiteURLSegments := strings.Split(whiteURL, ".")
			if len(urlSegments) == len(whiteURLSegments) {
				for i, urlSegment := range urlSegments {
					found = whiteURLSegments[i] == "*" || urlSegment == whiteURLSegments[i]
					if !found {
						break
					}
				}
				if found {
					break
				}
			} else {
				found = false
			}
		}
		return found
	}
	return true
}

//----------------------------Use Parser Combinator to do the Parsing-----------------------------//
func urlParser(input string) (parserCombinator.ParseNode, string) {
	populateURLStruct := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		url := URLStruct{}
		url.Protocol = nodes[0].(string)

		switch port := nodes[2].(type) {
		case int:
			url.Port = strconv.Itoa(port)
		}

		switch base := nodes[1].(type) {
		case []string:
			url.Base = base
		case string:
			if len(base) > 0 {
				url.Base = []string{base}
			}
		}
		url.Query = nodes[5].(string)
		url.Anchor = nodes[4].(string)

		switch pathSegments := nodes[3].(type) {
		case []string:
			url.Path = pathSegments[0]
			if len(pathSegments) > 1 {
				url.Filename = pathSegments[1]
			}
			if len(pathSegments) > 2 {
				url.FileExt = pathSegments[2]
			}
		}

		if len(url.Base) == 0 && url.Filename == "" && (url.Path == "" || url.Path == "/") {
			return nil
		}

		return url
	}

	//url := [ protocol ] [ ipV6 | ipV4 | domain ] [ port ] [ path ] [ queryString | anchorTag ]
	absURL := parserCombinator.Or(passFirstNode, ipv6, ipv4, domain)
	isUrl := parserCombinator.And(
		populateURLStruct,
		maybe(protocol),
		maybe(absURL),
		maybe(port),
		maybe(path),
		maybe(anchorString),
		maybe(queryString))

	return isUrl(input)
}

func path(input string) (parserCombinator.ParseNode, string) {
	processPath := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		var pathStr string

		relative := nodes[0].(string)
		if relative != "" {
			switch relative {
			case "PARENTDIR":
				pathStr += ".."
			case "CURDIR":
				pathStr += "."
			}
		}

		pathList := nodes[2].([]parserCombinator.ParseNode)
		for i, pathSegment := range pathList {
			if i < len(pathList)-1 {
				pathStr += "/" + pathSegment.(string)
			}
		}
		// Attempt to tease a file name and extension out of the last segment of the path.
		fileName := pathList[len(pathList)-1].(string)
		fileNameParts := strings.SplitN(fileName, ".", 2)
		if len(fileNameParts) > 1 {
			fileExt := fileNameParts[1]
			fileName := strings.Replace(fileName, "."+fileExt, "", 1)
			return []string{pathStr + "/", fileName, fileExt}
		} else {
			return []string{pathStr + "/", fileName}
		}
	}

	return parserCombinator.And(
		processPath,
		maybe(superRelativeBase),
		forwardSlash,
		parserCombinator.ListOf(parserCombinator.PassAllNodes, pathSegment, forwardSlash))(input)
}

func domain(input string) (parserCombinator.ParseNode, string) {
	processDomainSegmentList := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		tld := nodes[len(nodes)-1].(string)
		if len(nodes) < 2 {
			if len(tld) < 2 {
				return nil
			}
		} else {
			domain := nodes[len(nodes)-2].(string)
			if len(tld) < 2 || len(domain) < 2 {
				return nil
			}
		}
		return nodes
	}

	processDomain := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		if nodes[2] == nil {
			return nil
		}
		var domainParts []string

		domainParts = append(domainParts, nodes[0].(string))
		for _, segment := range nodes[2].([]parserCombinator.ParseNode) {
			domainParts = append(domainParts, segment.(string))
		}

		return domainParts
	}

	domainSegmentList := parserCombinator.ListOf(processDomainSegmentList, domainSegment, dot)

	return parserCombinator.And(processDomain, domainSegment, dot, domainSegmentList)(input)
}

func subnet(input string) (parserCombinator.ParseNode, string) {
	result, output := parseInt(input)
	if result == nil || result.(int) > 255 {
		return nil, input
	}
	return result, output
}

func ipv4(input string) (parserCombinator.ParseNode, string) {
	processIPV4 := func(nodes ...parserCombinator.ParseNode) parserCombinator.ParseNode {
		var ipAddress string
		for _, node := range nodes {
			switch t := node.(type) {
			case int:
				ipAddress += strconv.Itoa(t)
			case string:
				ipAddress += "."
			}
		}
		return ipAddress
	}
	return parserCombinator.And(processIPV4, subnet, dot, subnet, dot, subnet, dot, subnet)(input)
}

func port(input string) (parserCombinator.ParseNode, string) {
	return parserCombinator.And(passSecondNode, colon, parseInt)(input)
}

func queryString(input string) (parserCombinator.ParseNode, string) {
	return parserCombinator.And(passSecondNode, questionMark, queryOrAnchorStringSegment)(input)
}

func anchorString(input string) (parserCombinator.ParseNode, string) {
	return parserCombinator.And(passSecondNode, hashTag, queryOrAnchorStringSegment)(input)
}

//---------------------Utility Functions------------------------------------------------------------
var passFirstNode = parserCombinator.PassNthNode(1)
var passSecondNode = parserCombinator.PassNthNode(2)
var maybe = func(p parserCombinator.Parser) parserCombinator.Parser {
	return parserCombinator.Maybe(passFirstNode, p)
}

//----------------------Basic Terminals-------------------------------------------------------------
var colon = parserCombinator.CharParser(':', "COLON")

var dot = parserCombinator.CharParser('.', "DOT")

var questionMark = parserCombinator.CharParser('?', "QUESTIONMARK")

var hashTag = parserCombinator.CharParser('#', "HASHTAG")

var forwardSlash = parserCombinator.CharParser('/', "FORWARDSLASH")

func protocol(input string) (parserCombinator.ParseNode, string) {
	switch {
	case strings.HasPrefix(input, "https://"):
		return "HTTPS", strings.Replace(input, "https://", "", 1)
	case strings.HasPrefix(input, "http://"):
		return "HTTP", strings.Replace(input, "http://", "", 1)
	case strings.HasPrefix(input, "ftp://"):
		return "FTP", strings.Replace(input, "ftp://", "", 1)
	case strings.HasPrefix(input, "//"):
		return "HTTP", strings.Replace(input, "//", "", 1)
	default:
		return nil, input
	}
}

func superRelativeBase(input string) (parserCombinator.ParseNode, string) {
	switch {
	case strings.HasPrefix(input, ".."):
		return "PARENTDIR", strings.Replace(input, "..", "", 1)
	case strings.HasPrefix(input, "."):
		return "CURDIR", strings.Replace(input, ".", "", 1)
	default:
		return nil, input
	}
}

//---------------------Terminals that need Regexps--------------------------------------------------
var domainSegment = parserCombinator.RegexpParser(regexp.MustCompile(`^[\w-]+`))

var queryOrAnchorStringSegment = parserCombinator.RegexpParser(regexp.MustCompile(`^[\w-,\=\[\]\*\~\.\&\!\$\@\:\;\+\'\/]+`))

var pathSegment = parserCombinator.RegexpParser(regexp.MustCompile(`^[\w-,\.\:\%]+`))

// Note: The ipv6 regex has worked well in my tests but may miss some edge cases,
// as ipv6 can take several forms. It is much more complicated than ipv4.
// @See: https://www.ietf.org/rfc/rfc2732.txt
// @See: https://www.ietf.org/rfc/rfc2460.txt
var ipv6 = parserCombinator.RegexpParser(regexp.MustCompile(`\[([0-9A-Fa-f]{1,4}(:|\.)|:){5,7}[0-9A-Fa-f]{1,4}\]`))

var isInt = regexp.MustCompile(`^[0-9]+`)

func parseInt(input string) (parserCombinator.ParseNode, string) {
	numberStr := isInt.FindString(input)
	if numberStr == "" {
		return nil, input
	}
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return nil, input
	}

	return number, strings.Replace(input, numberStr, "", 1)
}
