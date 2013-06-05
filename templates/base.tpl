{{define "page" }}
<!doctype html>
<html lang="en">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
{{with .AppCube.BlogCube}}
<title>{{.BlogName}}</title>
{{end}}
<link href="/statics/bootstrap/css/bootstrap.css" rel="stylesheet" />
<style type="text/css">
body {
  padding-top: 60px;
  padding-bottom: 40px;
}
.bigFont {
  font-size: 140%;
}
.width0 {
  width: 100%;
}
.whitebg {
  background-color: #ffffff;
}
.colorbg {
  background-color: #eeefff;
}
.colorborder-pink {
  border-color: #FF99CC;
}
.colorborder-green {
  border-color: #66FF99;
}
.colorborder-blue {
  border-color: #00CCFF;
}
.colorborder-purple {
  border-color: #CC99FF;
}
.colorborder-oringe {
  border-color: #FFA347;
}
.colorborder-lgreen {
  border-color: #99FF99;
}

</style>
<link href="/statics/bootstrap/css/bootstrap-responsive.css" rel="stylesheet">
<script type="text/javascript" src="/statics/js/jquery.js"></script>
<script type="text/javascript" src="/statics/js/xheditor/xheditor.js"></script>
</head>
<body>
  <!-- nav -->

    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container">
          <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </a>
              {{with .AppCube.BlogCube}}
          <a class="brand" href="/">{{.BlogName}}</a>
              {{end}}
          <div class="nav-collapse collapse">
            <ul class="nav">
              {{with .PageCube.ActiveNav}}
              <li class="{{if strEqual . "home"}}active{{end}}"><a href="/">Home</a></li>
              {{end}}
            </ul>
             <form class="navbar-form pull-right">
              <a href="manageblog" class="btn btn-info">Management</a>
            </form> 
          </div><!--/.nav-collapse -->
        </div>
      </div>
    </div>
  <!-- end of nav -->

  {{ template "container" .}}

  <!-- javascript lib -->
<!-- my javasciprt -->
<!-- Le javascript
    ================================================== -->
<!-- Placed at the end of the document so the pages load faster -->
</body>
</html>
{{end}}

