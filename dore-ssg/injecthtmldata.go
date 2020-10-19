package main

var Prefix = []byte(`
<!DOCTYPE html>
<!--[if IE 8]><html class="no-js lt-ie9" lang="en" > <![endif]-->
<!--[if gt IE 8]><!-->
<html class="no-js" lang="en" >
<!--<![endif]-->
<head>
<meta charset="utf-8" />
<meta name="viewport" content="width=device-width, initial-scale=1.0" />
<meta property="og:description" content="Document Title" />
<meta property="og:image" content="" />
<meta property="og:site_name" content="" />
<meta property="og:type" content="" />
<meta property="og:title" content="Document Title" />
<meta property="og:type" content="website" />
<meta property="og:updated_time" content="" />
<meta property="og:url" content="" />
<title>Documentation</title>
<link rel="shortcut icon" href="static/img/favicon.ico"/>
<link rel="stylesheet" href="static/css/theme.css" type="text/css" />
<link rel="stylesheet" href="static/css/jquery-ui.min.css" type="text/css" />
</head>
<body class="wy-body-for-nav">
`)

var PrefixNav = []byte(`
<div class="wy-grid-for-nav">
`)

var PreContent = []byte(`
<section data-toggle="wy-nav-shift" class="wy-nav-content-wrap">
<div class="wy-nav-content">
<div class="rst-content">
<a href="index.html">Home</a>
`)

var PreContentEnd = []byte(`
<div role="main" class="document" itemscope="itemscope" itemtype="http://schema.org/Article">
<div itemprop="articleBody" class="toBeIndexed">
`)

var PostfixNav = []byte(`
</div><!-- /.toBeIndexed -->
</div>
<footer>
<div class="rst-footer-nav" role="navigation" aria-label="footer navigation">
`)

var PostfixNavEnd = []byte(`
</div></footer>
</div>
</div>
</section>
</div>
`)

var Postfix = []byte(`
<script type="text/javascript">
var DOCUMENTATION_OPTIONS = {
URL_ROOT:'./',
VERSION:'',
COLLAPSE_INDEX:false,
FILE_SUFFIX:'.html',
HAS_SOURCE:  true
};
</script>
<script type="text/javascript" src="static/js/jquery.min.js"></script>
<script type="text/javascript" src="static/js/jquery-ui.min.js"></script>
<script type="text/javascript" src="static/js/underscore.js"></script>
<script type="text/javascript" src="static/js/doctools.js"></script>
<script type="text/javascript" src="static/js/theme.js"></script>
<script type="text/javascript">
jQuery(function () {
SphinxRtdTheme.StickyNav.enable();
});
</script>
<script type="text/javascript" id="idPiwikScriptPlaceholder"></script>
<header>
<div class="header__container">
<a href="index.html"><img class="raw_img" height="64" src="img/bannerLogo.png" alt="Documentation" /></a>
`)

var PostfixTerminal = []byte(`
</div>
</header>
</body>
`)

var landingPageSearch = []byte(`
<div class="landing-search" role="search">
<form class="wy-form" action="search" method="get">
<input type="text" name="q" placeholder="Search this document" autofocus />
</form>
</div>
`)
