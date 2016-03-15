package urlParser

import (
	"github.com/Apartments24-7/goSpider/parserCombinator"
	"testing"
)

func TestURLParserCombinator(t *testing.T) {
	var result parserCombinator.ParseNode
	var leftovers string

	result, leftovers = urlParser("https://127.0.0.1:8080/some/path")
	if result == nil || leftovers != "" {
		t.Log(result, leftovers)
		t.Fail()
	}
	result, leftovers = urlParser("../some/path?query=string")
	if result == nil || leftovers != "" {
		t.Log(result, leftovers)
		t.Fail()
	}
	result, leftovers = urlParser("some/path?query=string")
	if result != nil || leftovers == "" {
		t.Log(result, leftovers)
		t.Fail()
	}
}

func TestParseList(t *testing.T) {
	uniqueURLs := ParseBody(testList, []string{})
	for _, url := range uniqueURLs {
		t.Log(url.DebugString())
	}

	if len(uniqueURLs) != 16 {
		t.Log("UniqueURLs length: ", len(uniqueURLs))
		t.Fail()
	}
}

func TestParseWithWhiteList(t *testing.T) {
	whiteList := []string{"apt24-7.com", "test.apt24-7.com"}
	uniqueURLs := ParseBody(testList, whiteList)

	for _, url := range uniqueURLs {
		t.Log(url)
	}
	t.Log("UniqueURLs length: ", len(uniqueURLs))

	if len(uniqueURLs) != 6 {
		t.Fail()
	}
}

func TestParseBody(t *testing.T) {
	uniqueURLs := ParseBody(testPage, []string{})

	t.Log("UniqueURLs length: ", len(uniqueURLs))
	if len(uniqueURLs) != 146 {
		t.Fail()
	}
}

func TestParseCSS(t *testing.T) {
	uniqueURLs := ParseBody(testCSS, []string{})

	t.Log("UniqueURLs length: ", len(uniqueURLs))
	for _, url := range uniqueURLs {
		t.Log(url)
	}
	if len(uniqueURLs) != 7 {
		t.Fail()
	}
}

func BenchmarkParseBody(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseBody(testPage, []string{})
	}
}

func BenchmarkParseCSS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseBody(testCSS, []string{})
	}
}

//---------------Test Data--------------------------------------------------------------------------
var testList = `
<img src="test_file_only.jpg" />

<img src="/test_relative_file.jpg" />
<img src="./relative/test.jpg" />
<img src="../deep/relative/test.jpg" />
<img src=".././deepest_relative_test.jpg" />

<img src="small.jpg" srcset="medium.jpg 1000w, large.jpg 2000w" alt="srcset test">


<img src="http://[::FFFF:129.144.52.38]:80/index.gif" alt="ipv6" />

<a href="http://[1080::8:800:200C:417A]/foo" >ipv6</a>
<a href="192.168.1.2/foo" >ipv4</a>
<a href="192.168.1.2:23/foo" >ipv4 w port</a>
<script src="https://test.apt24-7.com/some/resource.js?num=1234" />
<script src="test.apt24-7.com/no/protocol.js#anchortag" />
<script src="test.apt24-7.com/no/protocol.js#anchortag" />
<script src="test-anchor-root.apt24-7.com/#anchortag" />
<script src="test-anchor-no-path.apt24-7.com#anchortag" />
`

var testPage = `
<!DOCTYPE html>
<html lang="en" dir="ltr" class="client-nojs">
<head>
<meta charset="UTF-8" />
<title>Recursive descent parser - Wikipedia, the free encyclopedia</title>
<script>document.documentElement.className = document.documentElement.className.replace( /(^|\s)client-nojs(\s|$)/, "$1client-js$2" );</script>
<script>window.RLQ = window.RLQ || []; window.RLQ.push( function () {
mw.config.set({"wgCanonicalNamespace":"","wgCanonicalSpecialPageName":!1,"wgNamespaceNumber":0,"wgPageName":"Recursive_descent_parser","wgTitle":"Recursive descent parser","wgCurRevisionId":671418181,"wgRevisionId":671418181,"wgArticleId":70089,"wgIsArticle":!0,"wgIsRedirect":!1,"wgAction":"view","wgUserName":null,"wgUserGroups":["*"],"wgCategories":["Articles lacking in-text citations from February 2009","All articles lacking in-text citations","Parsing algorithms","Articles with example C code"],"wgBreakFrames":!1,"wgPageContentLanguage":"en","wgPageContentModel":"wikitext","wgSeparatorTransformTable":["",""],"wgDigitTransformTable":["",""],"wgDefaultDateFormat":"dmy","wgMonthNames":["","January","February","March","April","May","June","July","August","September","October","November","December"],"wgMonthNamesShort":["","Jan","Feb","Mar","Apr","May","Jun","Jul","Aug","Sep","Oct","Nov","Dec"],"wgRelevantPageName":"Recursive_descent_parser","wgRelevantArticleId":70089,
"wgIsProbablyEditable":!0,"wgRestrictionEdit":[],"wgRestrictionMove":[],"wgMediaViewerOnClick":!0,"wgMediaViewerEnabledByDefault":!0,"wikilove-recipient":"","wikilove-anon":0,"wgWikiEditorEnabledModules":{"toolbar":!0,"dialogs":!0,"preview":!1,"publish":!1},"wgBetaFeaturesFeatures":[],"wgVisualEditor":{"pageLanguageCode":"en","pageLanguageDir":"ltr","usePageImages":!0,"usePageDescriptions":!0},"wgGatherShouldShowTutorial":!0,"wgULSAcceptLanguageList":["en-us","en"],"wgULSCurrentAutonym":"English","wgFlaggedRevsParams":{"tags":{"status":{"levels":1,"quality":2,"pristine":3}}},"wgStableRevisionId":null,"wgCategoryTreePageCategoryOptions":"{\"mode\":0,\"hideprefix\":20,\"showcount\":true,\"namespaces\":false}","wgNoticeProject":"wikipedia","wgWikibaseItemId":"Q1323264","wgVisualEditorToolbarScrollOffset":0});mw.loader.implement("user.options",function($,jQuery){mw.user.options.set({"variant":"en"});});mw.loader.implement("user.tokens",function($,jQuery){mw.user.tokens.set({"editToken":"+\\","patrolToken":"+\\","watchToken":"+\\"});});mw.loader.load(["mediawiki.page.startup","mediawiki.legacy.wikibits","ext.centralauth.centralautologin","mmv.head","ext.gadget.WatchlistBase","ext.gadget.WatchlistGreenIndicators","ext.visualEditor.desktopArticleTarget.init","ext.uls.init","ext.uls.interface","ext.centralNotice.bannerController","skins.vector.js"]);
} );</script>
<link rel="stylesheet" href="https://en.wikipedia.org/w/load.php?debug=false&amp;lang=en&amp;modules=ext.gadget.WatchlistBase%2CWatchlistGreenIndicators%7Cext.uls.nojs%7Cext.visualEditor.desktopArticleTarget.noscript%7Cmediawiki.legacy.shared%7Cmediawiki.sectionAnchor%7Cmediawiki.skinning.interface%7Cmediawiki.ui.button%7Cskins.vector.styles%7Cwikibase.client.init&amp;only=styles&amp;skin=vector&amp;*" />
<link rel="stylesheet" href="https://en.wikipedia.org/w/load.php?debug=false&amp;lang=en&amp;modules=mediawiki.legacy.commonPrint&amp;only=styles&amp;skin=vector&amp;*" media="print" />
<meta name="ResourceLoaderDynamicStyles" content="" />
<link rel="stylesheet" href="https://en.wikipedia.org/w/load.php?debug=false&amp;lang=en&amp;modules=site&amp;only=styles&amp;skin=vector&amp;*" />
<style>a:lang(ar),a:lang(kk-arab),a:lang(mzn),a:lang(ps),a:lang(ur){text-decoration:none}</style>
<script async="" src="https://en.wikipedia.org/w/load.php?debug=false&amp;lang=en&amp;modules=startup&amp;only=scripts&amp;skin=vector&amp;*"></script>
<meta name="generator" content="MediaWiki 1.26wmf19" />
<link rel="alternate" href="android-app://org.wikipedia/http/en.m.wikipedia.org/wiki/Recursive_descent_parser" />
<link rel="alternate" type="application/x-wiki" title="Edit this page" href="/w/index.php?title=Recursive_descent_parser&amp;action=edit" />
<link rel="edit" title="Edit this page" href="/w/index.php?title=Recursive_descent_parser&amp;action=edit" />
<link rel="apple-touch-icon" href="/static/apple-touch/wikipedia.png" />
<link rel="shortcut icon" href="/static/favicon/wikipedia.ico" />
<link rel="search" type="application/opensearchdescription+xml" href="/w/opensearch_desc.php" title="Wikipedia (en)" />
<link rel="EditURI" type="application/rsd+xml" href="//en.wikipedia.org/w/api.php?action=rsd" />
<link rel="copyright" href="//creativecommons.org/licenses/by-sa/3.0/" />
<link rel="alternate" type="application/atom+xml" title="Wikipedia Atom feed" href="/w/index.php?title=Special:RecentChanges&amp;feed=atom" />
<link rel="canonical" href="https://en.wikipedia.org/wiki/Recursive_descent_parser" />
<link rel="dns-prefetch" href="//meta.wikimedia.org" />
<!--[if lt IE 7]><style type="text/css">body{behavior:url("/w/static/1.26wmf19/skins/Vector/csshover.min.htc")}</style><![endif]-->
</head>
<body class="mediawiki ltr sitedir-ltr ns-0 ns-subject page-Recursive_descent_parser skin-vector action-view">
		<div id="mw-page-base" class="noprint"></div>
		<div id="mw-head-base" class="noprint"></div>
		<div id="content" class="mw-body" role="main">
			<a id="top"></a>

							<div id="siteNotice"><!-- CentralNotice --></div>
						<div class="mw-indicators">
</div>
			<h1 id="firstHeading" class="firstHeading" lang="en">Recursive descent parser</h1>
									<div id="bodyContent" class="mw-body-content">
									<div id="siteSub">From Wikipedia, the free encyclopedia</div>
								<div id="contentSub"></div>
												<div id="jump-to-nav" class="mw-jump">
					Jump to:					<a href="#mw-head">navigation</a>, 					<a href="#p-search">search</a>
				</div>
				<div id="mw-content-text" lang="en" dir="ltr" class="mw-content-ltr"><table class="metadata plainlinks ambox ambox-style ambox-More_footnotes" role="presentation">
<tr>
<td class="mbox-image">
<div style="width:52px"><img alt="" src="//upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Text_document_with_red_question_mark.svg/40px-Text_document_with_red_question_mark.svg.png" width="40" height="40" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Text_document_with_red_question_mark.svg/60px-Text_document_with_red_question_mark.svg.png 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/a/a4/Text_document_with_red_question_mark.svg/80px-Text_document_with_red_question_mark.svg.png 2x" data-file-width="48" data-file-height="48" /></div>
</td>
<td class="mbox-text"><span class="mbox-text-span">This article includes a <a href="/wiki/Wikipedia:Citing_sources" title="Wikipedia:Citing sources">list of references</a>, but <b>its sources remain unclear</b> because it has <b>insufficient <a href="/wiki/Wikipedia:Citing_sources#Inline_citations" title="Wikipedia:Citing sources">inline citations</a></b>. <span class="hide-when-compact">Please help to <a href="/wiki/Wikipedia:WikiProject_Fact_and_Reference_Check" title="Wikipedia:WikiProject Fact and Reference Check">improve</a> this article by <a href="/wiki/Wikipedia:When_to_cite" title="Wikipedia:When to cite">introducing</a> more precise citations.</span> <small><i>(February 2009)</i></small></span></td>
</tr>
</table>
<p>In <a href="/wiki/Computer_science" title="Computer science">computer science</a>, a <b>recursive descent parser</b> is a kind of <a href="/wiki/Top-down_parsing" title="Top-down parsing">top-down</a> <a href="/wiki/Parsing" title="Parsing">parser</a> built from a set of <a href="/wiki/Mutual_recursion" title="Mutual recursion">mutually recursive</a> procedures (or a non-recursive equivalent) where each such <a href="/wiki/Procedure_(computer_science)" title="Procedure (computer science)" class="mw-redirect">procedure</a> usually implements one of the <a href="/wiki/Production_(computer_science)" title="Production (computer science)">productions</a> of the <a href="/wiki/Formal_grammar" title="Formal grammar">grammar</a>. Thus the structure of the resulting program closely mirrors that of the grammar it recognizes.<sup id="cite_ref-1" class="reference"><a href="#cite_note-1"><span>[</span>1<span>]</span></a></sup></p>
<p>A <i>predictive parser</i> is a recursive descent parser that does not require <a href="/wiki/Backtracking" title="Backtracking">backtracking</a>. Predictive parsing is possible only for the class of <a href="/wiki/LL_parser" title="LL parser">LL(<i>k</i>)</a> grammars, which are the <a href="/wiki/Context-free_grammar" title="Context-free grammar">context-free grammars</a> for which there exists some positive integer <i>k</i> that allows a recursive descent parser to decide which production to use by examining only the next <i>k</i> tokens of input. The LL(<i>k</i>) grammars therefore exclude all ambiguous grammars, as well as all grammars that contain <a href="/wiki/Left_recursion" title="Left recursion">left recursion</a>. Any context-free grammar can be transformed into an equivalent grammar that has no left recursion, but removal of left recursion does not always yield an LL(<i>k</i>) grammar. A predictive parser runs in <a href="/wiki/Linear_time" title="Linear time" class="mw-redirect">linear time</a>.</p>
<p>Recursive descent with backtracking is a technique that determines which production to use by trying each production in turn. Recursive descent with backtracking is not limited to LL(k) grammars, but is not guaranteed to terminate unless the grammar is LL(k). Even when they terminate, parsers that use recursive descent with backtracking may require <a href="/wiki/Exponential_time" title="Exponential time" class="mw-redirect">exponential time</a>.</p>
<p>Although predictive parsers are widely used, and are frequently chosen if writing a parser by hand, programmers often prefer to use a table-based parser produced by a <a href="/wiki/Parser_generator" title="Parser generator" class="mw-redirect">parser generator</a>, either for an LL(<i>k</i>) language or using an alternative parser, such as <a href="/wiki/LALR_parser" title="LALR parser">LALR</a> or <a href="/wiki/LR_parser" title="LR parser">LR</a>. This is particularly the case if a grammar is not in <a href="/wiki/LL_parser" title="LL parser">LL(<i>k</i>)</a> form, as transforming the grammar to LL to make it suitable for predictive parsing is involved. Predictive parsers can also be automatically generated, using tools like <a href="/wiki/ANTLR" title="ANTLR">ANTLR</a>.</p>
<p>Predictive parsers can be depicted using transition diagrams for each non-terminal symbol where the edges between the initial and the final states are labelled by the symbols (terminals and non-terminals) of the right side of the production rule.<sup id="cite_ref-2" class="reference"><a href="#cite_note-2"><span>[</span>2<span>]</span></a></sup></p>
<p></p>
<div id="toc" class="toc">
<div id="toctitle">
<h2>Contents</h2>
</div>
<ul>
<li class="toclevel-1 tocsection-1"><a href="#Example_parser"><span class="tocnumber">1</span> <span class="toctext">Example parser</span></a>
<ul>
<li class="toclevel-2 tocsection-2"><a href="#C_implementation"><span class="tocnumber">1.1</span> <span class="toctext">C implementation</span></a></li>
</ul>
</li>
<li class="toclevel-1 tocsection-3"><a href="#See_also"><span class="tocnumber">2</span> <span class="toctext">See also</span></a></li>
<li class="toclevel-1 tocsection-4"><a href="#References"><span class="tocnumber">3</span> <span class="toctext">References</span></a></li>
<li class="toclevel-1 tocsection-5"><a href="#External_links"><span class="tocnumber">4</span> <span class="toctext">External links</span></a></li>
</ul>
</div>
<p></p>
<h2><span class="mw-headline" id="Example_parser">Example parser</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit&amp;section=1" title="Edit section: Example parser">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<p>The following <a href="/wiki/EBNF" title="EBNF" class="mw-redirect">EBNF</a>-like <a href="/wiki/Formal_grammar" title="Formal grammar">grammar</a> (for <a href="/wiki/Niklaus_Wirth" title="Niklaus Wirth">Niklaus Wirth</a>'s <a href="/wiki/PL/0" title="PL/0">PL/0</a> programming language, from <i><a href="/wiki/Algorithms_%2B_Data_Structures_%3D_Programs" title="Algorithms + Data Structures = Programs">Algorithms + Data Structures = Programs</a></i>) is in <a href="/wiki/LL_parser" title="LL parser">LL(1)</a> form:</p>
<div class="mw-highlight mw-content-ltr" dir="ltr">
<pre>
 <span class="k">program </span><span class="o">=</span> <span class="k">block </span><span class="s2">"."</span> <span class="p">.</span>
 
 <span class="k">block </span><span class="o">=</span>
     <span class="p">[</span><span class="s2">"const"</span> <span class="k">ident </span><span class="s2">"="</span> <span class="k">number </span><span class="p">{</span><span class="s2">","</span> <span class="k">ident </span><span class="s2">"="</span> <span class="k">number</span><span class="p">}</span> <span class="s2">";"</span><span class="p">]</span>
     <span class="p">[</span><span class="s2">"var"</span> <span class="k">ident </span><span class="p">{</span><span class="s2">","</span> <span class="k">ident</span><span class="p">}</span> <span class="s2">";"</span><span class="p">]</span>
     <span class="p">{</span><span class="s2">"procedure"</span> <span class="k">ident </span><span class="s2">";"</span> <span class="k">block </span><span class="s2">";"</span><span class="p">}</span> <span class="k">statement </span><span class="p">.</span>
 
 <span class="k">statement </span><span class="o">=</span>
     <span class="k">ident </span><span class="s2">":="</span> <span class="k">expression</span>
     <span class="p">|</span> <span class="s2">"call"</span> <span class="k">ident</span>
     <span class="p">|</span> <span class="s2">"begin"</span> <span class="k">statement </span><span class="p">{</span><span class="s2">";"</span> <span class="k">statement </span><span class="p">}</span> <span class="s2">"end"</span>
     <span class="p">|</span> <span class="s2">"if"</span> <span class="k">condition </span><span class="s2">"then"</span> <span class="k">statement</span>
     <span class="p">|</span> <span class="s2">"while"</span> <span class="k">condition </span><span class="s2">"do"</span> <span class="k">statement </span><span class="p">.</span>
 
 <span class="k">condition </span><span class="o">=</span>
     <span class="s2">"odd"</span> <span class="k">expression</span>
     <span class="p">|</span> <span class="k">expression </span><span class="p">(</span><span class="s2">"="</span><span class="p">|</span><span class="s2">"#"</span><span class="p">|</span><span class="s2">"&lt;"</span><span class="p">|</span><span class="s2">"&lt;="</span><span class="p">|</span><span class="s2">"&gt;"</span><span class="p">|</span><span class="s2">"&gt;="</span><span class="p">)</span> <span class="k">expression </span><span class="p">.</span>
 
 <span class="k">expression </span><span class="o">=</span> <span class="p">[</span><span class="s2">"+"</span><span class="p">|</span><span class="s2">"-"</span><span class="p">]</span> <span class="k">term </span><span class="p">{(</span><span class="s2">"+"</span><span class="p">|</span><span class="s2">"-"</span><span class="p">)</span> <span class="k">term</span><span class="p">}</span> <span class="p">.</span>
 
 <span class="k">term </span><span class="o">=</span> <span class="k">factor </span><span class="p">{(</span><span class="s2">"*"</span><span class="p">|</span><span class="s2">"/"</span><span class="p">)</span> <span class="k">factor</span><span class="p">}</span> <span class="p">.</span>
 
 <span class="k">factor </span><span class="o">=</span>
     <span class="k">ident</span>
     <span class="p">|</span> <span class="k">number</span>
     <span class="p">|</span> <span class="s2">"("</span> <span class="k">expression </span><span class="s2">")"</span> <span class="p">.</span>
</pre></div>
<p><a href="/wiki/Terminal_symbol" title="Terminal symbol" class="mw-redirect">Terminals</a> are expressed in quotes. Each <a href="/wiki/Nonterminal_symbol" title="Nonterminal symbol" class="mw-redirect">nonterminal</a> is defined by a rule in the grammar, except for <i>ident</i> and <i>number</i>, which are assumed to be implicitly defined.</p>
<h3><span class="mw-headline" id="C_implementation">C implementation</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit&amp;section=2" title="Edit section: C implementation">edit</a><span class="mw-editsection-bracket">]</span></span></h3>
<p>What follows is an implementation of a recursive descent parser for the above language in <a href="/wiki/C_(programming_language)" title="C (programming language)">C</a>. The parser reads in source code, and exits with an error message if the code fails to parse, exiting silently if the code parses correctly.</p>
<p>Notice how closely the predictive parser below mirrors the grammar above. There is a procedure for each nonterminal in the grammar. Parsing descends in a top-down manner, until the final nonterminal has been processed. The program fragment depends on a global variable, <i>sym</i>, which contains the current symbol from the input, and the function <i>nextsym</i>, which updates <i>sym</i> when called.</p>
<p>The implementations of the functions <i>nextsym</i> and <i>error</i> are omitted for simplicity.</p>
<div class="mw-highlight mw-content-ltr" dir="ltr">
<pre>
<span class="k">typedef</span> <span class="k">enum</span> <span class="p">{</span><span class="n">ident</span><span class="p">,</span> <span class="n">number</span><span class="p">,</span> <span class="n">lparen</span><span class="p">,</span> <span class="n">rparen</span><span class="p">,</span> <span class="n">times</span><span class="p">,</span> <span class="n">slash</span><span class="p">,</span> <span class="n">plus</span><span class="p">,</span>
    <span class="n">minus</span><span class="p">,</span> <span class="n">eql</span><span class="p">,</span> <span class="n">neq</span><span class="p">,</span> <span class="n">lss</span><span class="p">,</span> <span class="n">leq</span><span class="p">,</span> <span class="n">gtr</span><span class="p">,</span> <span class="n">geq</span><span class="p">,</span> <span class="n">callsym</span><span class="p">,</span> <span class="n">beginsym</span><span class="p">,</span> <span class="n">semicolon</span><span class="p">,</span>
    <span class="n">endsym</span><span class="p">,</span> <span class="n">ifsym</span><span class="p">,</span> <span class="n">whilesym</span><span class="p">,</span> <span class="n">becomes</span><span class="p">,</span> <span class="n">thensym</span><span class="p">,</span> <span class="n">dosym</span><span class="p">,</span> <span class="n">constsym</span><span class="p">,</span> <span class="n">comma</span><span class="p">,</span>
    <span class="n">varsym</span><span class="p">,</span> <span class="n">procsym</span><span class="p">,</span> <span class="n">period</span><span class="p">,</span> <span class="n">oddsym</span><span class="p">}</span> <span class="n">Symbol</span><span class="p">;</span>

<span class="n">Symbol</span> <span class="n">sym</span><span class="p">;</span>
<span class="kt">void</span> <span class="nf">nextsym</span><span class="p">(</span><span class="kt">void</span><span class="p">);</span>
<span class="kt">void</span> <span class="nf">error</span><span class="p">(</span><span class="k">const</span> <span class="kt">char</span> <span class="n">msg</span><span class="p">[]);</span>

<span class="kt">int</span> <span class="nf">accept</span><span class="p">(</span><span class="n">Symbol</span> <span class="n">s</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">sym</span> <span class="o">==</span> <span class="n">s</span><span class="p">)</span> <span class="p">{</span>
        <span class="n">nextsym</span><span class="p">();</span>
        <span class="k">return</span> <span class="mi">1</span><span class="p">;</span>
    <span class="p">}</span>
    <span class="k">return</span> <span class="mi">0</span><span class="p">;</span>
<span class="p">}</span>

<span class="kt">int</span> <span class="nf">expect</span><span class="p">(</span><span class="n">Symbol</span> <span class="n">s</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">s</span><span class="p">))</span>
        <span class="k">return</span> <span class="mi">1</span><span class="p">;</span>
    <span class="n">error</span><span class="p">(</span><span class="s">"expect: unexpected symbol"</span><span class="p">);</span>
    <span class="k">return</span> <span class="mi">0</span><span class="p">;</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">factor</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">ident</span><span class="p">))</span> <span class="p">{</span>
        <span class="p">;</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">number</span><span class="p">))</span> <span class="p">{</span>
        <span class="p">;</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">lparen</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">expression</span><span class="p">();</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">rparen</span><span class="p">);</span>
    <span class="p">}</span> <span class="k">else</span> <span class="p">{</span>
        <span class="n">error</span><span class="p">(</span><span class="s">"factor: syntax error"</span><span class="p">);</span>
        <span class="n">nextsym</span><span class="p">();</span>
    <span class="p">}</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">term</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="n">factor</span><span class="p">();</span>
    <span class="k">while</span> <span class="p">(</span><span class="n">sym</span> <span class="o">==</span> <span class="n">times</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">slash</span><span class="p">)</span> <span class="p">{</span>
        <span class="n">nextsym</span><span class="p">();</span>
        <span class="n">factor</span><span class="p">();</span>
    <span class="p">}</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">expression</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">sym</span> <span class="o">==</span> <span class="n">plus</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">minus</span><span class="p">)</span>
        <span class="n">nextsym</span><span class="p">();</span>
    <span class="n">term</span><span class="p">();</span>
    <span class="k">while</span> <span class="p">(</span><span class="n">sym</span> <span class="o">==</span> <span class="n">plus</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">minus</span><span class="p">)</span> <span class="p">{</span>
        <span class="n">nextsym</span><span class="p">();</span>
        <span class="n">term</span><span class="p">();</span>
    <span class="p">}</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">condition</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">oddsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">expression</span><span class="p">();</span>
    <span class="p">}</span> <span class="k">else</span> <span class="p">{</span>
        <span class="n">expression</span><span class="p">();</span>
        <span class="k">if</span> <span class="p">(</span><span class="n">sym</span> <span class="o">==</span> <span class="n">eql</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">neq</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">lss</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">leq</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">gtr</span> <span class="o">||</span> <span class="n">sym</span> <span class="o">==</span> <span class="n">geq</span><span class="p">)</span> <span class="p">{</span>
            <span class="n">nextsym</span><span class="p">();</span>
            <span class="n">expression</span><span class="p">();</span>
        <span class="p">}</span> <span class="k">else</span> <span class="p">{</span>
            <span class="n">error</span><span class="p">(</span><span class="s">"condition: invalid operator"</span><span class="p">);</span>
            <span class="n">nextsym</span><span class="p">();</span>
        <span class="p">}</span>
    <span class="p">}</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">statement</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">ident</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">becomes</span><span class="p">);</span>
        <span class="n">expression</span><span class="p">();</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">callsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">ident</span><span class="p">);</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">beginsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="k">do</span> <span class="p">{</span>
            <span class="n">statement</span><span class="p">();</span>
        <span class="p">}</span> <span class="k">while</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">semicolon</span><span class="p">));</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">endsym</span><span class="p">);</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">ifsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">condition</span><span class="p">();</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">thensym</span><span class="p">);</span>
        <span class="n">statement</span><span class="p">();</span>
    <span class="p">}</span> <span class="k">else</span> <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">whilesym</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">condition</span><span class="p">();</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">dosym</span><span class="p">);</span>
        <span class="n">statement</span><span class="p">();</span>
    <span class="p">}</span> <span class="k">else</span> <span class="p">{</span>
        <span class="n">error</span><span class="p">(</span><span class="s">"statement: syntax error"</span><span class="p">);</span>
        <span class="n">nextsym</span><span class="p">();</span>
    <span class="p">}</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">block</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">constsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="k">do</span> <span class="p">{</span>
            <span class="n">expect</span><span class="p">(</span><span class="n">ident</span><span class="p">);</span>
            <span class="n">expect</span><span class="p">(</span><span class="n">eql</span><span class="p">);</span>
            <span class="n">expect</span><span class="p">(</span><span class="n">number</span><span class="p">);</span>
        <span class="p">}</span> <span class="k">while</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">comma</span><span class="p">));</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">semicolon</span><span class="p">);</span>
    <span class="p">}</span>
    <span class="k">if</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">varsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="k">do</span> <span class="p">{</span>
            <span class="n">expect</span><span class="p">(</span><span class="n">ident</span><span class="p">);</span>
        <span class="p">}</span> <span class="k">while</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">comma</span><span class="p">));</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">semicolon</span><span class="p">);</span>
    <span class="p">}</span>
    <span class="k">while</span> <span class="p">(</span><span class="n">accept</span><span class="p">(</span><span class="n">procsym</span><span class="p">))</span> <span class="p">{</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">ident</span><span class="p">);</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">semicolon</span><span class="p">);</span>
        <span class="n">block</span><span class="p">();</span>
        <span class="n">expect</span><span class="p">(</span><span class="n">semicolon</span><span class="p">);</span>
    <span class="p">}</span>
    <span class="n">statement</span><span class="p">();</span>
<span class="p">}</span>

<span class="kt">void</span> <span class="nf">program</span><span class="p">(</span><span class="kt">void</span><span class="p">)</span> <span class="p">{</span>
    <span class="n">nextsym</span><span class="p">();</span>
    <span class="n">block</span><span class="p">();</span>
    <span class="n">expect</span><span class="p">(</span><span class="n">period</span><span class="p">);</span>
<span class="p">}</span>
</pre></div>
<h2><span class="mw-headline" id="See_also">See also</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit&amp;section=3" title="Edit section: See also">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<div class="noprint portal tright" style="border:solid #aaa 1px;margin:0.5em 0 0.5em 1em">
<table style="background:#f9f9f9;font-size:85%;line-height:110%;max-width:175px">
<tr style="vertical-align:middle">
<td style="text-align:center"><a href="/wiki/File:Internet_map_1024.jpg" class="image"><img alt="Portal icon" src="//upload.wikimedia.org/wikipedia/commons/thumb/d/d2/Internet_map_1024.jpg/28px-Internet_map_1024.jpg" width="28" height="28" class="noviewer" srcset="//upload.wikimedia.org/wikipedia/commons/thumb/d/d2/Internet_map_1024.jpg/42px-Internet_map_1024.jpg 1.5x, //upload.wikimedia.org/wikipedia/commons/thumb/d/d2/Internet_map_1024.jpg/56px-Internet_map_1024.jpg 2x" data-file-width="1280" data-file-height="1280" /></a></td>
<td style="padding:0 0.2em;vertical-align:middle;font-style:italic;font-weight:bold"><a href="/wiki/Portal:Computer_science" title="Portal:Computer science">Computer science portal</a></td>
</tr>
</table>
</div>
<ul>
<li><a href="/wiki/JavaCC" title="JavaCC">JavaCC</a> – a recursive descent parser generator</li>
<li><a href="/wiki/Coco/R" title="Coco/R">Coco/R</a> – a recursive descent parser generator</li>
<li><a href="/wiki/ANTLR" title="ANTLR">ANTLR</a> – a recursive descent parser generator</li>
<li><a href="/wiki/Parsing_expression_grammar" title="Parsing expression grammar">Parsing expression grammar</a> – another form representing recursive descent grammar</li>
<li><a href="/wiki/Spirit_Parser_Framework" title="Spirit Parser Framework">Spirit Parser Framework</a> – a C++ recursive descent parser generator framework requiring no pre-compile step</li>
<li><a href="/wiki/Tail_recursive_parser" title="Tail recursive parser">Tail recursive parser</a> – a variant of the recursive descent parser</li>
<li><a href="/wiki/Parboiled_(Java)" title="Parboiled (Java)">parboiled (Java)</a> – a recursive descent PEG parsing library for <a href="/wiki/Java_(programming_language)" title="Java (programming language)">Java</a></li>
<li><a href="/wiki/Recursive_ascent_parser" title="Recursive ascent parser">Recursive ascent parser</a></li>
<li><a rel="nofollow" class="external text" href="http://sourceforge.net/projects/bnf2xml/">bnf2xml</a> Markup input with XML tags using advanced BNF matching. (a top town LL recursive parser, front to back text, no compiling of lexor is needed or used)</li>
<li><a rel="nofollow" class="external text" href="https://metacpan.org/module/Parse::RecDescent">Parse::RecDescent</a>: A versatile recursive descent <a href="/wiki/Perl" title="Perl">Perl</a> module.</li>
<li><a rel="nofollow" class="external text" href="http://pyparsing.sourceforge.net/">pyparsing</a>: A versatile <a href="/wiki/Python_(programming_language)" title="Python (programming language)">Python</a> recursive parsing module that is not recursive descent (<a rel="nofollow" class="external text" href="http://mail.python.org/pipermail/python-list/2007-November/421649.html">python-list post</a>).</li>
<li><a rel="nofollow" class="external text" href="http://jparsec.codehaus.org/">Jparsec</a> a Java port of Haskell's Parsec module.</li>
</ul>
<h2><span class="mw-headline" id="References">References</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit&amp;section=4" title="Edit section: References">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<div class="reflist" style="list-style-type: decimal;">
<ol class="references">
<li id="cite_note-1"><span class="mw-cite-backlink"><b><a href="#cite_ref-1">^</a></b></span> <span class="reference-text"><span class="citation book">Burge, W.H. (1975). <i>Recursive Programming Techniques</i>. <a href="/wiki/International_Standard_Book_Number" title="International Standard Book Number">ISBN</a>&#160;<a href="/wiki/Special:BookSources/0-201-14450-6" title="Special:BookSources/0-201-14450-6">0-201-14450-6</a>.</span><span title="ctx_ver=Z39.88-2004&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ARecursive+descent+parser&amp;rft.au=Burge%2C+W.H.&amp;rft.aulast=Burge%2C+W.H.&amp;rft.btitle=Recursive+Programming+Techniques&amp;rft.date=1975&amp;rft.genre=book&amp;rft.isbn=0-201-14450-6&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook" class="Z3988"><span style="display:none;">&#160;</span></span></span></li>
<li id="cite_note-2"><span class="mw-cite-backlink"><b><a href="#cite_ref-2">^</a></b></span> <span class="reference-text"><span class="citation book"><a href="/wiki/Alfred_V._Aho" title="Alfred V. Aho" class="mw-redirect">Aho, Alfred V.</a>; Sethi, Ravi; <a href="/wiki/Jeffrey_Ullman" title="Jeffrey Ullman">Ullman, Jeffrey</a> (1986). <i>Compilers: Principles, Techniques and Tools</i> (first ed.). Addison Wesley. p.&#160;183.</span><span title="ctx_ver=Z39.88-2004&amp;rfr_id=info%3Asid%2Fen.wikipedia.org%3ARecursive+descent+parser&amp;rft.au=Aho%2C+Alfred+V.&amp;rft.aufirst=Alfred+V.&amp;rft.aulast=Aho&amp;rft.au=Sethi%2C+Ravi&amp;rft.au=Ullman%2C+Jeffrey&amp;rft.btitle=Compilers%3A+Principles%2C+Techniques+and+Tools&amp;rft.date=1986&amp;rft.edition=first&amp;rft.genre=book&amp;rft.pages=183&amp;rft.pub=Addison+Wesley&amp;rft_val_fmt=info%3Aofi%2Ffmt%3Akev%3Amtx%3Abook" class="Z3988"><span style="display:none;">&#160;</span></span></span></li>
</ol>
</div>
<p><span class="citation foldoc">This article is based on material taken from the <a href="/wiki/Free_On-line_Dictionary_of_Computing" title="Free On-line Dictionary of Computing">Free On-line Dictionary of Computing</a> prior to 1 November 2008 and incorporated under the "relicensing" terms of the <a href="/wiki/GNU_Free_Documentation_License" title="GNU Free Documentation License">GFDL</a>, version 1.3 or later.</span></p>
<ul>
<li><i><a href="/wiki/Compilers:_Principles,_Techniques,_and_Tools" title="Compilers: Principles, Techniques, and Tools">Compilers: Principles, Techniques, and Tools</a></i>, first edition, Alfred V Aho, Ravi Sethi, and Jeffrey D Ullman, in particular Section 4.4.</li>
<li><i>Modern Compiler Implementation in Java, Second Edition</i>, Andrew Appel, 2002, <a href="/wiki/Special:BookSources/052182060X" class="internal mw-magiclink-isbn">ISBN 0-521-82060-X</a>.</li>
<li><i>Recursive Programming Techniques</i>, W.H. Burge, 1975, <a href="/wiki/Special:BookSources/0201144506" class="internal mw-magiclink-isbn">ISBN 0-201-14450-6</a></li>
<li><i>Crafting a Compiler with C</i>, Charles N Fischer and Richard J LeBlanc, Jr, 1991, <a href="/wiki/Special:BookSources/0805321667" class="internal mw-magiclink-isbn">ISBN 0-8053-2166-7</a>.</li>
<li><i>Compiling with C# and Java</i>, Pat Terry, 2005, <a href="/wiki/Special:BookSources/032126360X" class="internal mw-magiclink-isbn">ISBN 0-321-26360-X</a>, 624</li>
<li><i>Algorithms + Data Structures = Programs</i>, Niklaus Wirth, 1975, <a href="/wiki/Special:BookSources/0130224189" class="internal mw-magiclink-isbn">ISBN 0-13-022418-9</a></li>
<li><i>Compiler Construction</i>, Niklaus Wirth, 1996, <a href="/wiki/Special:BookSources/0201403536" class="internal mw-magiclink-isbn">ISBN 0-201-40353-6</a></li>
</ul>
<h2><span class="mw-headline" id="External_links">External links</span><span class="mw-editsection"><span class="mw-editsection-bracket">[</span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit&amp;section=5" title="Edit section: External links">edit</a><span class="mw-editsection-bracket">]</span></span></h2>
<ul>
<li><a rel="nofollow" class="external text" href="http://www.mollypages.org/page/grammar/index.mp">Introduction to Parsing</a> - an easy to read introduction to parsing, with a comprehensive section on recursive descent parsing</li>
<li><a rel="nofollow" class="external text" href="http://teaching.idallen.com/cst8152/98w/recursive_decent_parsing.html">How to turn a Grammar into C code</a> - a brief tutorial on implementing recursive descent parser</li>
<li><a rel="nofollow" class="external text" href="http://lukaszwrobel.pl/blog/math-parser-part-3-implementation">Simple mathematical expressions parser</a> in <a href="/wiki/Ruby_(programming_language)" title="Ruby (programming language)">Ruby</a></li>
<li><a rel="nofollow" class="external text" href="http://effbot.org/zone/simple-top-down-parsing.htm">Simple Top Down Parsing in Python</a></li>
<li><a rel="nofollow" class="external text" href="http://compilers.iecc.com/crenshaw/">Jack W. Crenshaw: <i>Let's Build A Compiler</i> (1988-1995)</a>, in <a href="/wiki/Pascal_(programming_language)" title="Pascal (programming language)">Pascal</a>, with <a href="/wiki/Assembly_language" title="Assembly language">assembly language</a> output, using a "keep it simple" approach</li>
<li><a rel="nofollow" class="external text" href="http://www.cs.nott.ac.uk/~gmh/pearl.pdf">Functional Pearls: Monadic Parsing in Haskell</a></li>
</ul>


<!-- 
NewPP limit report
Parsed by mw1168
CPU time usage: 0.148 seconds
Real time usage: 0.194 seconds
Preprocessor visited node count: 316/1000000
Preprocessor generated node count: 0/1500000
Post‐expand include size: 7561/2097152 bytes
Template argument size: 13/2097152 bytes
Highest expansion depth: 6/40
Expensive parser function count: 1/500
Lua time usage: 0.051/10.000 seconds
Lua memory usage: 1.61 MB/50 MB
-->

<!-- 
Transclusion expansion time report (%,ms,calls,template)
100.00%  134.795      1 - -total
 51.70%   69.689      1 - Template:More_footnotes
 32.89%   44.338      1 - Template:Ambox
 31.83%   42.909      1 - Template:Reflist
 26.08%   35.157      2 - Template:Cite_book
 10.29%   13.871      1 - Template:Portal
  2.00%    2.697      1 - Template:FOLDOC
-->

<!-- Saved in parser cache with key enwiki:pcache:idhash:70089-0!*!0!!en!4!* and timestamp 20150725103852 and revision id 671418181
 -->
<noscript><img src="//en.wikipedia.org/wiki/Special:CentralAutoLogin/start?type=1x1" alt="" title="" width="1" height="1" style="border: none; position: absolute;" /></noscript></div>					<div class="printfooter">
						Retrieved from "<a dir="ltr" href="https://en.wikipedia.org/w/index.php?title=Recursive_descent_parser&amp;oldid=671418181">https://en.wikipedia.org/w/index.php?title=Recursive_descent_parser&amp;oldid=671418181</a>"					</div>
				<div id='catlinks' class='catlinks'><div id="mw-normal-catlinks" class="mw-normal-catlinks"><a href="/wiki/Help:Category" title="Help:Category">Categories</a>: <ul><li><a href="/wiki/Category:Parsing_algorithms" title="Category:Parsing algorithms">Parsing algorithms</a></li></ul></div><div id="mw-hidden-catlinks" class="mw-hidden-catlinks mw-hidden-cats-hidden">Hidden categories: <ul><li><a href="/wiki/Category:Articles_lacking_in-text_citations_from_February_2009" title="Category:Articles lacking in-text citations from February 2009">Articles lacking in-text citations from February 2009</a></li><li><a href="/wiki/Category:All_articles_lacking_in-text_citations" title="Category:All articles lacking in-text citations">All articles lacking in-text citations</a></li><li><a href="/wiki/Category:Articles_with_example_C_code" title="Category:Articles with example C code">Articles with example C code</a></li></ul></div></div>				<div class="visualClear"></div>
							</div>
		</div>
		<div id="mw-navigation">
			<h2>Navigation menu</h2>

			<div id="mw-head">
									<div id="p-personal" role="navigation" class="" aria-labelledby="p-personal-label">
						<h3 id="p-personal-label">Personal tools</h3>
						<ul>
							<li id="pt-createaccount"><a href="/w/index.php?title=Special:UserLogin&amp;returnto=Recursive+descent+parser&amp;type=signup" title="You are encouraged to create an account and log in; however, it is not mandatory">Create account</a></li><li id="pt-login"><a href="/w/index.php?title=Special:UserLogin&amp;returnto=Recursive+descent+parser" title="You're encouraged to log in; however, it's not mandatory. [o]" accesskey="o">Log in</a></li>						</ul>
					</div>
									<div id="left-navigation">
										<div id="p-namespaces" role="navigation" class="vectorTabs" aria-labelledby="p-namespaces-label">
						<h3 id="p-namespaces-label">Namespaces</h3>
						<ul>
															<li  id="ca-nstab-main" class="selected"><span><a href="/wiki/Recursive_descent_parser"  title="View the content page [c]" accesskey="c">Article</a></span></li>
															<li  id="ca-talk"><span><a href="/wiki/Talk:Recursive_descent_parser"  title="Discussion about the content page [t]" accesskey="t" rel="discussion">Talk</a></span></li>
													</ul>
					</div>
										<div id="p-variants" role="navigation" class="vectorMenu emptyPortlet" aria-labelledby="p-variants-label">
												<h3 id="p-variants-label">
							<span>Variants</span><a href="#"></a>
						</h3>

						<div class="menu">
							<ul>
															</ul>
						</div>
					</div>
									</div>
				<div id="right-navigation">
										<div id="p-views" role="navigation" class="vectorTabs" aria-labelledby="p-views-label">
						<h3 id="p-views-label">Views</h3>
						<ul>
															<li id="ca-view" class="selected"><span><a href="/wiki/Recursive_descent_parser" >Read</a></span></li>
															<li id="ca-edit"><span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=edit"  title="Edit this page [e]" accesskey="e">Edit</a></span></li>
															<li id="ca-history" class="collapsible"><span><a href="/w/index.php?title=Recursive_descent_parser&amp;action=history"  title="Past versions of this page [h]" accesskey="h">View history</a></span></li>
													</ul>
					</div>
										<div id="p-cactions" role="navigation" class="vectorMenu emptyPortlet" aria-labelledby="p-cactions-label">
						<h3 id="p-cactions-label"><span>More</span><a href="#"></a></h3>

						<div class="menu">
							<ul>
															</ul>
						</div>
					</div>
										<div id="p-search" role="search">
						<h3>
							<label for="searchInput">Search</label>
						</h3>

						<form action="/w/index.php" id="searchform">
							<div id="simpleSearch">
							<input type="search" name="search" placeholder="Search" title="Search Wikipedia [f]" accesskey="f" id="searchInput" /><input type="hidden" value="Special:Search" name="title" /><input type="submit" name="fulltext" value="Search" title="Search Wikipedia for this text" id="mw-searchButton" class="searchButton mw-fallbackSearchButton" /><input type="submit" name="go" value="Go" title="Go to a page with this exact name if it exists" id="searchButton" class="searchButton" />							</div>
						</form>
					</div>
									</div>
			</div>
			<div id="mw-panel">
				<div id="p-logo" role="banner"><a class="mw-wiki-logo" href="/wiki/Main_Page"  title="Visit the main page"></a></div>
						<div class="portal" role="navigation" id='p-navigation' aria-labelledby='p-navigation-label'>
			<h3 id='p-navigation-label'>Navigation</h3>

			<div class="body">
									<ul>
						<li id="n-mainpage-description"><a href="/wiki/Main_Page" title="Visit the main page [z]" accesskey="z">Main page</a></li><li id="n-contents"><a href="/wiki/Portal:Contents" title="Guides to browsing Wikipedia">Contents</a></li><li id="n-featuredcontent"><a href="/wiki/Portal:Featured_content" title="Featured content – the best of Wikipedia">Featured content</a></li><li id="n-currentevents"><a href="/wiki/Portal:Current_events" title="Find background information on current events">Current events</a></li><li id="n-randompage"><a href="/wiki/Special:Random" title="Load a random article [x]" accesskey="x">Random article</a></li><li id="n-sitesupport"><a href="https://donate.wikimedia.org/wiki/Special:FundraiserRedirector?utm_source=donate&amp;utm_medium=sidebar&amp;utm_campaign=C13_en.wikipedia.org&amp;uselang=en" title="Support us">Donate to Wikipedia</a></li><li id="n-shoplink"><a href="//shop.wikimedia.org" title="Visit the Wikimedia store">Wikipedia store</a></li>					</ul>
							</div>
		</div>
			<div class="portal" role="navigation" id='p-interaction' aria-labelledby='p-interaction-label'>
			<h3 id='p-interaction-label'>Interaction</h3>

			<div class="body">
									<ul>
						<li id="n-help"><a href="/wiki/Help:Contents" title="Guidance on how to use and edit Wikipedia">Help</a></li><li id="n-aboutsite"><a href="/wiki/Wikipedia:About" title="Find out about Wikipedia">About Wikipedia</a></li><li id="n-portal"><a href="/wiki/Wikipedia:Community_portal" title="About the project, what you can do, where to find things">Community portal</a></li><li id="n-recentchanges"><a href="/wiki/Special:RecentChanges" title="A list of recent changes in the wiki [r]" accesskey="r">Recent changes</a></li><li id="n-contactpage"><a href="//en.wikipedia.org/wiki/Wikipedia:Contact_us">Contact page</a></li>					</ul>
							</div>
		</div>
			<div class="portal" role="navigation" id='p-tb' aria-labelledby='p-tb-label'>
			<h3 id='p-tb-label'>Tools</h3>

			<div class="body">
									<ul>
						<li id="t-whatlinkshere"><a href="/wiki/Special:WhatLinksHere/Recursive_descent_parser" title="List of all English Wikipedia pages containing links to this page [j]" accesskey="j">What links here</a></li><li id="t-recentchangeslinked"><a href="/wiki/Special:RecentChangesLinked/Recursive_descent_parser" title="Recent changes in pages linked from this page [k]" accesskey="k">Related changes</a></li><li id="t-upload"><a href="/wiki/Wikipedia:File_Upload_Wizard" title="Upload files [u]" accesskey="u">Upload file</a></li><li id="t-specialpages"><a href="/wiki/Special:SpecialPages" title="A list of all special pages [q]" accesskey="q">Special pages</a></li><li id="t-permalink"><a href="/w/index.php?title=Recursive_descent_parser&amp;oldid=671418181" title="Permanent link to this revision of the page">Permanent link</a></li><li id="t-info"><a href="/w/index.php?title=Recursive_descent_parser&amp;action=info" title="More information about this page">Page information</a></li><li id="t-wikibase"><a href="//www.wikidata.org/wiki/Q1323264" title="Link to connected data repository item [g]" accesskey="g">Wikidata item</a></li><li id="t-cite"><a href="/w/index.php?title=Special:CiteThisPage&amp;page=Recursive_descent_parser&amp;id=671418181" title="Information on how to cite this page">Cite this page</a></li>					</ul>
							</div>
		</div>
			<div class="portal" role="navigation" id='p-coll-print_export' aria-labelledby='p-coll-print_export-label'>
			<h3 id='p-coll-print_export-label'>Print/export</h3>

			<div class="body">
									<ul>
						<li id="coll-create_a_book"><a href="/w/index.php?title=Special:Book&amp;bookcmd=book_creator&amp;referer=Recursive+descent+parser">Create a book</a></li><li id="coll-download-as-rdf2latex"><a href="/w/index.php?title=Special:Book&amp;bookcmd=render_article&amp;arttitle=Recursive+descent+parser&amp;oldid=671418181&amp;writer=rdf2latex">Download as PDF</a></li><li id="t-print"><a href="/w/index.php?title=Recursive_descent_parser&amp;printable=yes" title="Printable version of this page [p]" accesskey="p">Printable version</a></li>					</ul>
							</div>
		</div>
			<div class="portal" role="navigation" id='p-lang' aria-labelledby='p-lang-label'>
			<h3 id='p-lang-label'>Languages</h3>

			<div class="body">
									<ul>
						<li class="interlanguage-link interwiki-ar"><a href="//ar.wikipedia.org/wiki/%D8%A7%D9%84%D8%AA%D8%B1%D9%85%D9%8A%D8%B2_%D8%A7%D9%84%D8%AA%D9%83%D8%B1%D8%A7%D8%B1%D9%8A_%D8%A7%D9%84%D9%86%D9%85%D9%88%D8%B0%D8%AC%D9%8A" title="الترميز التكراري النموذجي – Arabic" lang="ar" hreflang="ar">العربية</a></li><li class="interlanguage-link interwiki-cs"><a href="//cs.wikipedia.org/wiki/Anal%C3%BDza_rekurzivn%C3%ADm_sestupem" title="Analýza rekurzivním sestupem – Czech" lang="cs" hreflang="cs">Čeština</a></li><li class="interlanguage-link interwiki-de"><a href="//de.wikipedia.org/wiki/Rekursiver_Abstieg" title="Rekursiver Abstieg – German" lang="de" hreflang="de">Deutsch</a></li><li class="interlanguage-link interwiki-ko"><a href="//ko.wikipedia.org/wiki/%EB%90%98%EB%B6%80%EB%A6%84_%ED%95%98%ED%96%A5_%EA%B5%AC%EB%AC%B8_%EB%B6%84%EC%84%9D" title="되부름 하향 구문 분석 – Korean" lang="ko" hreflang="ko">한국어</a></li><li class="interlanguage-link interwiki-ja"><a href="//ja.wikipedia.org/wiki/%E5%86%8D%E5%B8%B0%E4%B8%8B%E9%99%8D%E6%A7%8B%E6%96%87%E8%A7%A3%E6%9E%90" title="再帰下降構文解析 – Japanese" lang="ja" hreflang="ja">日本語</a></li><li class="interlanguage-link interwiki-pt"><a href="//pt.wikipedia.org/wiki/Analisador_sint%C3%A1tico_descendente_recursivo" title="Analisador sintático descendente recursivo – Portuguese" lang="pt" hreflang="pt">Português</a></li><li class="interlanguage-link interwiki-ru"><a href="//ru.wikipedia.org/wiki/%D0%9C%D0%B5%D1%82%D0%BE%D0%B4_%D1%80%D0%B5%D0%BA%D1%83%D1%80%D1%81%D0%B8%D0%B2%D0%BD%D0%BE%D0%B3%D0%BE_%D1%81%D0%BF%D1%83%D1%81%D0%BA%D0%B0" title="Метод рекурсивного спуска – Russian" lang="ru" hreflang="ru">Русский</a></li><li class="interlanguage-link interwiki-sr"><a href="//sr.wikipedia.org/wiki/Analizator_rekurzivnim_spustom" title="Analizator rekurzivnim spustom – Serbian" lang="sr" hreflang="sr">Српски / srpski</a></li><li class="interlanguage-link interwiki-uk"><a href="//uk.wikipedia.org/wiki/%D0%A0%D0%B5%D0%BA%D1%83%D1%80%D1%81%D0%B8%D0%B2%D0%BD%D0%B8%D0%B9_%D1%81%D0%BF%D1%83%D1%81%D0%BA" title="Рекурсивний спуск – Ukrainian" lang="uk" hreflang="uk">Українська</a></li><li class="uls-p-lang-dummy"><a href="#"></a></li>					</ul>
				<div class='after-portlet after-portlet-lang'><span class="wb-langlinks-edit wb-langlinks-link"><a href="//www.wikidata.org/wiki/Q1323264#sitelinks-wikipedia" title="Edit interlanguage links" class="wbc-editpage">Edit links</a></span></div>			</div>
		</div>
				</div>
		</div>
		<div id="footer" role="contentinfo">
							<ul id="footer-info">
											<li id="footer-info-lastmod"> This page was last modified on 14 July 2015, at 15:45.</li>
											<li id="footer-info-copyright">Text is available under the <a rel="license" href="//en.wikipedia.org/wiki/Wikipedia:Text_of_Creative_Commons_Attribution-ShareAlike_3.0_Unported_License">Creative Commons Attribution-ShareAlike License</a><a rel="license" href="//creativecommons.org/licenses/by-sa/3.0/" style="display:none;"></a>;
additional terms may apply.  By using this site, you agree to the <a href="//wikimediafoundation.org/wiki/Terms_of_Use">Terms of Use</a> and <a href="//wikimediafoundation.org/wiki/Privacy_policy">Privacy Policy</a>. Wikipedia® is a registered trademark of the <a href="//www.wikimediafoundation.org/">Wikimedia Foundation, Inc.</a>, a non-profit organization.</li>
									</ul>
							<ul id="footer-places">
											<li id="footer-places-privacy"><a href="//wikimediafoundation.org/wiki/Privacy_policy" title="wikimedia:Privacy policy">Privacy policy</a></li>
											<li id="footer-places-about"><a href="/wiki/Wikipedia:About" title="Wikipedia:About">About Wikipedia</a></li>
											<li id="footer-places-disclaimer"><a href="/wiki/Wikipedia:General_disclaimer" title="Wikipedia:General disclaimer">Disclaimers</a></li>
											<li id="footer-places-contact"><a href="//en.wikipedia.org/wiki/Wikipedia:Contact_us">Contact Wikipedia</a></li>
											<li id="footer-places-developers"><a href="https://www.mediawiki.org/wiki/Special:MyLanguage/How_to_contribute">Developers</a></li>
											<li id="footer-places-mobileview"><a href="//en.m.wikipedia.org/w/index.php?title=Recursive_descent_parser&amp;mobileaction=toggle_view_mobile" class="noprint stopMobileRedirectToggle">Mobile view</a></li>
									</ul>
										<ul id="footer-icons" class="noprint">
											<li id="footer-copyrightico">
							<a href="//wikimediafoundation.org/"><img src="/static/images/wikimedia-button.png" srcset="/static/images/wikimedia-button-1.5x.png 1.5x, /static/images/wikimedia-button-2x.png 2x" width="88" height="31" alt="Wikimedia Foundation"/></a>						</li>
											<li id="footer-poweredbyico">
							<a href="//www.mediawiki.org/"><img src="https://en.wikipedia.org/static/1.26wmf19/resources/assets/poweredby_mediawiki_88x31.png" alt="Powered by MediaWiki" srcset="https://en.wikipedia.org/static/1.26wmf19/resources/assets/poweredby_mediawiki_132x47.png 1.5x, https://en.wikipedia.org/static/1.26wmf19/resources/assets/poweredby_mediawiki_176x62.png 2x" width="88" height="31" /></a>						</li>
									</ul>
						<div style="clear:both"></div>
		</div>
		<script>window.RLQ = window.RLQ || []; window.RLQ.push( function () {
mw.loader.state({"ext.globalCssJs.site":"ready","ext.globalCssJs.user":"ready","user":"ready","user.groups":"ready"});
} );</script>
<link rel="stylesheet" href="https://en.wikipedia.org/w/load.php?debug=false&amp;lang=en&amp;modules=ext.cite.styles%7Cext.gadget.DRN-wizard%2CReferenceTooltips%2Ccharinsert%2Cfeatured-articles-links%2CrefToolbar%2Cswitcher%2Cteahouse%7Cext.pygments%2CwikimediaBadges&amp;only=styles&amp;skin=vector&amp;*" />
<script>window.RLQ = window.RLQ || []; window.RLQ.push( function () {
mw.loader.load(["ext.cite.a11y","mediawiki.toc","mediawiki.action.view.postEdit","site","mediawiki.user","mediawiki.hidpi","mediawiki.page.ready","mediawiki.searchSuggest","ext.cirrusSearch.loggingSchema","mmv.bootstrap.autostart","ext.eventLogging.subscriber","ext.wikimediaEvents","ext.wikimediaEvents.statsd","ext.wikimediaEvents.geoFeatures","ext.navigationTiming","ext.gadget.teahouse","ext.gadget.ReferenceTooltips","ext.gadget.DRN-wizard","ext.gadget.charinsert","ext.gadget.refToolbar","ext.gadget.switcher","ext.gadget.featured-articles-links","ext.visualEditor.targetLoader","schema.UniversalLanguageSelector","ext.uls.eventlogger","ext.uls.interlanguage"]);
} );</script><script>window.RLQ = window.RLQ || []; window.RLQ.push( function () {
mw.config.set({"wgBackendResponseTime":92,"wgHostname":"mw1074"});
} );</script>
	</body>
</html>
`

var testCSS = `.mw-cite-backlink,.cite-accessibility-label{-moz-user-select:none;-webkit-user-select:none;-ms-user-select:none;user-select:none}sup.reference{unicode-bidi:-moz-isolate;unicode-bidi:-webkit-isolate;unicode-bidi:isolate}ol.references li:target,sup.reference:target{background-color:#def;background-color:rgba(0,127,255,0.133)}@media print{.mw-cite-backlink{display:none}}.referencetooltip{position:absolute;list-style:none;list-style-image:none;opacity:0;font-size:10px;margin:0;z-index:5;padding:0}.referencetooltip li{border:#080086 2px solid;max-width:260px;padding:10px 8px 13px 8px;margin:0px;background-color:#F7F7F7;box-shadow:2px 4px 2px rgba(0,0,0,0.3);-moz-box-shadow:2px 4px 2px rgba(0,0,0,0.3);-webkit-box-shadow:2px 4px 2px rgba(0,0,0,0.3)}.referencetooltip li+li{margin-left:7px;margin-top:-2px;border:0;padding:0;height:3px;width:0px;background-color:transparent;box-shadow:none;-moz-box-shadow:none;-webkit-box-shadow:none;border-top:12px #080086 solid;border-right:7px transparent solid;border-left:7px transparent solid}.referencetooltip>li+li::after{content:'';border-top:8px #F7F7F7 solid;border-right:5px transparent solid;border-left:5px transparent solid;margin-top:-12px;margin-left:-5px;z-index:1;height:0px;width:0px;display:block}.client-js body .referencetooltip li li{border:none;box-shadow:none;-moz-box-shadow:none;-webkit-box-shadow:none;height:auto;width:auto;margin:auto;padding:0;position:static}.RTflipped{padding-top:13px}.referencetooltip.RTflipped li+li{position:absolute;top:2px;border-top:0;border-bottom:12px #080086 solid}.referencetooltip.RTflipped li+li::after{border-top:0;border-bottom:8px #F7F7F7 solid;position:absolute;margin-top:7px}.RTsettings{float:right;height:24px;width:24px;cursor:pointer;background-image:url(//upload.wikimedia.org/wikipedia/commons/thumb/7/77/Gear_icon.svg/24px-Gear_icon.svg.png);background-image:linear-gradient(transparent,transparent),url(//upload.wikimedia.org/wikipedia/commons/7/77/Gear_icon.svg);margin-top:-9px;margin-right:-7px;-webkit-transition:opacity 0.15s;-moz-transition:opacity 0.15s;-o-transition:opacity 0.15s;-ms-transition:opacity 0.15s;transition:opacity 0.15s;opacity:0.6;filter:alpha(opacity=60)}.RTsettings:hover{opacity:1;filter:alpha(opacity=100)}.RTTarget{border:#080086 2px solid}.skin-vector li.GA,.skin-monobook li.GA,.skin-modern li.GA{list-style-image:url(//upload.wikimedia.org/wikipedia/commons/4/42/Monobook-bullet-ga.png)}.skin-vector li.FA,.skin-monobook li.FA{list-style-image:url(//upload.wikimedia.org/wikipedia/commons/d/d4/Monobook-bullet-star.png)}.skin-modern li.FA{list-style-image:url(//upload.wikimedia.org/wikipedia/commons/thumb/2/2c/Modern-bullet-star.svg/9px-Modern-bullet-star.svg.png)}.wp-teahouse-question-form{position:absolute;margin-left:auto;margin-right:auto;background-color:#f4f3f0;border:1px solid #a7d7f9;padding:1em}#wp-th-question-ask{float:right}.wp-teahouse-ask a.external{background-image:none !important}.wp-teahouse-respond-form{position:absolute;margin-left:auto;margin-right:auto;background-color:#f4f3f0;border:1px solid #a7d7f9;padding:1em}.wp-th-respond{float:right}.wp-teahouse-respond a.external{background-image:none !important}.mw-highlight .hll{background-color:#ffffcc }.mw-highlight{background:#f8f8f8}.mw-highlight .c{color:#408080;font-style:italic }.mw-highlight .err{border:1px solid #FF0000 }.mw-highlight .k{color:#008000;font-weight:bold }.mw-highlight .o{color:#666666 }.mw-highlight .cm{color:#408080;font-style:italic }.mw-highlight .cp{color:#BC7A00 }.mw-highlight .c1{color:#408080;font-style:italic }.mw-highlight .cs{color:#408080;font-style:italic }.mw-highlight .gd{color:#A00000 }.mw-highlight .ge{font-style:italic }.mw-highlight .gr{color:#FF0000 }.mw-highlight .gh{color:#000080;font-weight:bold }.mw-highlight .gi{color:#00A000 }.mw-highlight .go{color:#888888 }.mw-highlight .gp{color:#000080;font-weight:bold }.mw-highlight .gs{font-weight:bold }.mw-highlight .gu{color:#800080;font-weight:bold }.mw-highlight .gt{color:#0044DD }.mw-highlight .kc{color:#008000;font-weight:bold }.mw-highlight .kd{color:#008000;font-weight:bold }.mw-highlight .kn{color:#008000;font-weight:bold }.mw-highlight .kp{color:#008000 }.mw-highlight .kr{color:#008000;font-weight:bold }.mw-highlight .kt{color:#B00040 }.mw-highlight .m{color:#666666 }.mw-highlight .s{color:#BA2121 }.mw-highlight .na{color:#7D9029 }.mw-highlight .nb{color:#008000 }.mw-highlight .nc{color:#0000FF;font-weight:bold }.mw-highlight .no{color:#880000 }.mw-highlight .nd{color:#AA22FF }.mw-highlight .ni{color:#999999;font-weight:bold }.mw-highlight .ne{color:#D2413A;font-weight:bold }.mw-highlight .nf{color:#0000FF }.mw-highlight .nl{color:#A0A000 }.mw-highlight .nn{color:#0000FF;font-weight:bold }.mw-highlight .nt{color:#008000;font-weight:bold }.mw-highlight .nv{color:#19177C }.mw-highlight .ow{color:#AA22FF;font-weight:bold }.mw-highlight .w{color:#bbbbbb }.mw-highlight .mb{color:#666666 }.mw-highlight .mf{color:#666666 }.mw-highlight .mh{color:#666666 }.mw-highlight .mi{color:#666666 }.mw-highlight .mo{color:#666666 }.mw-highlight .sb{color:#BA2121 }.mw-highlight .sc{color:#BA2121 }.mw-highlight .sd{color:#BA2121;font-style:italic }.mw-highlight .s2{color:#BA2121 }.mw-highlight .se{color:#BB6622;font-weight:bold }.mw-highlight .sh{color:#BA2121 }.mw-highlight .si{color:#BB6688;font-weight:bold }.mw-highlight .sx{color:#008000 }.mw-highlight .sr{color:#BB6688 }.mw-highlight .s1{color:#BA2121 }.mw-highlight .ss{color:#19177C }.mw-highlight .bp{color:#008000 }.mw-highlight .vc{color:#19177C }.mw-highlight .vg{color:#19177C }.mw-highlight .vi{color:#19177C }.mw-highlight .il{color:#666666 }.mw-highlight{direction:ltr;unicode-bidi:embed}code code.mw-highlight{background-color:transparent;border:0;padding:0}.mw-highlight .err{border:0}.mw-highlight .hll{display:block}.mw-highlight.mw-content-ltr .lineno{float:left}.mw-highlight.mw-content-rtl .lineno{float:right}.badge-goodarticle,.badge-recommendedarticle{list-style-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAkAAAANBAMAAACJLlk1AAAAAXNSR0IArs4c6QAAACFQTFRFAAAAbGxsbm5ucXFxc3Nzd3d3enp6e3t7fX19f39/iYmJftgPsQAAAAF0Uk5TAEDm2GYAAAAwSURBVAhbY2DACSaASQEQwSoKJFoEBVUaGJgFBYWBHEVBAyBpCCaFkoCKuAoYqhgAamoEibSxbiMAAAAASUVORK5CYII=);list-style-image:url(https://en.wikipedia.org/static/1.26wmf19/extensions/Wikidata/extensions/WikimediaBadges/resources/images/badge-silver-star.png?5da14)!ie}.badge-featuredarticle,.badge-featuredportal,.badge-featuredlist{list-style-image:url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAkAAAANBAMAAACJLlk1AAAAJFBMVEX8/vz8wjT8viT83pT8ylT88tz80mz8xkT82oT84pz8xjz81nwKqN0kAAAAAXRSTlMAQObYZgAAADFJREFUCNdjYMAJJoBJBRDBpgYkTJSURAwY2JWUtIAcIaUFQLIITCpuBCpibWBoZQAAfGEGEiyyYyMAAAAASUVORK5CYII=);list-style-image:url(https://en.wikipedia.org/static/1.26wmf19/extensions/Wikidata/extensions/WikimediaBadges/resources/images/badge-golden-star.png?c1f32)!ie}`
