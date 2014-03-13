<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <title>Golang技术博客</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">

    <!-- Le styles -->
    <link href="/static/css/bootstrap.css" rel="stylesheet">
    <style type="text/css">
      body {
        padding-top: 60px;
        padding-bottom: 40px;
      }
      .sidebar-nav {
        padding: 9px 0;
      }
    </style>
    <link href="/static/css/bootstrap-responsive.css" rel="stylesheet">

    <!-- HTML5 shim, for IE6-8 support of HTML5 elements -->
    <!--[if lt IE 9]>
      <script src="http://html5shim.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->

    <!-- Fav and touch icons -->
    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="/static/ico/apple-touch-icon-144-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="114x114" href="/static/ico/apple-touch-icon-114-precomposed.png">
    <link rel="apple-touch-icon-precomposed" sizes="72x72" href="/static/ico/apple-touch-icon-72-precomposed.png">
    <link rel="apple-touch-icon-precomposed" href="/static/ico/apple-touch-icon-57-precomposed.png">
    <link rel="shortcut icon" href="/static/ico/favicon.png">
  </head>

  <body>

    <div class="navbar navbar-inverse navbar-fixed-top">
      <div class="navbar-inner">
        <div class="container-fluid">
          <a class="btn btn-navbar" data-toggle="collapse" data-target=".nav-collapse">
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
          </a>
          <a class="brand" href="#">Golang 技术博客</a>
          <div class="nav-collapse collapse">            
            <ul class="nav">
               <li class="active"><a href="/">首页</a></li>
               <li><a href="/about">关于我们</a></li>
               <li><a href="/contact">联系我们</a></li>
				{{if .Username}}
					<li><a style="color:#ffffff;margin-left:600px;">欢迎：{{.Username}}</a></li>
					<li><a href="/logout">退出</a></li>
				{{else}}
					<li>
					<form action="/login" method="post" style="height:20px;">
						<input type="text" name="username" placeholder="Username" value="" style="margin:5px;"/>
						<input type="text" name="password" placeholder="Password" value="" style="margin:5px;"/>
						<button type="submit" class="btn btn-primary" style="margin:5px;">Login</button>
						<a href="/reg" style="color:#ffffff;">Register</a>
					</form>
					</li>
				{{end}}
            </ul>
          </div><!--/.nav-collapse -->
        </div>
      </div>

		
    </div>

    <div class="container-fluid">
      <div class="row-fluid">
        <div class="span3">
			
			
			
			<form action="/" method="get">
		    <fieldset>
				
					<div class="input-append input-prepend">
						<span class="add-on"><i class="icon-search"></i></span>
						<input type="text" name="s" id="search" placeholder="Search" value="" />
						<button type="submit" class="btn btn-primary">Search</button>
					</div>
		      
		    </fieldset>
			</form>
      <form enctype="multipart/form-data" action="/upload" method="post">
          <input type="file" name="uploadfile" />
          <input type="hidden" name="token" value="{{.}}" />
          <input type="submit" value="upload" />
      </form>
		

          <div class="well sidebar-nav">
            <ul class="nav nav-list">
              <li class="nav-header">博客文章</li>
			{{range .BlogList}}
				{{if .Cru}}
				   <li class="active"><a href="/{{.Blog.Id}}">{{.Blog.Title}}</a></li>
				{{else}}
				   <li><a href="/{{.Blog.Id}}">{{.Blog.Title}}</a></li>
				{{end}}
			{{end}}
            </ul>
          </div><!--/.well -->
        </div><!--/span-->
        <div class="span9">
          <div class="hero-unit">
				{{.Content}}
          </div>         
        </div><!--/span-->
      </div><!--/row-->

      <hr>

      <footer>
        <p>&copy; Company 2012</p>
      </footer>

    </div><!--/.fluid-container-->

    <!-- Le javascript
    ================================================== -->
    <!-- Placed at the end of the document so the pages load faster -->
    <script src="/static/js/jquery.js"></script>
    <script src="/static/js/bootstrap-transition.js"></script>
    <script src="/static/js/bootstrap-alert.js"></script>
    <script src="/static/js/bootstrap-modal.js"></script>
    <script src="/static/js/bootstrap-dropdown.js"></script>
    <script src="/static/js/bootstrap-scrollspy.js"></script>
    <script src="/static/js/bootstrap-tab.js"></script>
    <script src="/static/js/bootstrap-tooltip.js"></script>
    <script src="/static/js/bootstrap-popover.js"></script>
    <script src="/static/js/bootstrap-button.js"></script>
    <script src="/static/js/bootstrap-collapse.js"></script>
    <script src="/static/js/bootstrap-carousel.js"></script>
    <script src="/static/js/bootstrap-typeahead.js"></script>

  </body>
</html>
