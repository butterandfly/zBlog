{{ define "container" }}
  <!-- container -->
  <div class="container">
  <div class="row">
    <div class="span4">
      <section class="well whitebg">
        {{with .AppCube.BlogCube}}
        <form action="manageblog" method="post">
          <label for="blogname">博客名：</label>
          <input type="text" name="blogname" id="blogname" value="{{.BlogName}}" />
          <br />
          <label for="noRight">显示右侧漂浮栏：</label>
          <input type="radio" name="noRight" value="0" /> 显示 
          <input type="radio" name="noRight" value="1" /> 隐藏<br/>
          <br/>
          <input type="submit" value="保存" />
        </form>
        {{end}}
      </section>
    </div>
    <div class="span4">
      <section class="well whitebg">
        <a href="/managearticle" class="btn btn-large btn-primary btn-block">管理文章</a>
      </section>
    </div>
    <div class="span4">
      <section class="well whitebg">
        <a href="/manageaf" class="btn btn-large btn-info btn-block">管理公告</a>
      </section>
    </div>
  </div><!-- end of .row -->
  </div>
  <!-- end of container -->
{{ end }}