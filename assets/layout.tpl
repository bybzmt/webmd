<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<link href="/assets/base.css" type="text/css" rel="stylesheet" />

<script type="text/javascript" src="/assets/jquery-3.1.0.min.js"></script>

<link href="/assets/google-code-prettify/prettify.css" type="text/css" rel="stylesheet" />
<script type="text/javascript" src="/assets/google-code-prettify/prettify.js"></script>
</head>
<body>

<div id="left">
</div>

<div id="main">
{{.md}}
</div>

</body>
<script>
$("code").each(function(){
    if (this.className && this.className.indexOf("language") == 0) {
        $(this).addClass("prettyprint linenums");
    }
});

prettyPrint()

var idx = location.href.lastIndexOf('/')

if (idx > 0) {
    $("#left").load(location.href.substr(0, idx));
} else {
    $("#left").load('/');
}
</script>
</html>
