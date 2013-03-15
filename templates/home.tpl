{{ define "container" }}
  <!-- container -->
  <div class="container">
  <div class="row">
    <div class="span8">
      {{range .}}
      <article class="well whitebg">
        <h2><a href="article?id={{.Id}}">{{.Title}}</a></h2>
        <p class="muted">
        <small><strong>Posted: </strong><em>{{.CreateTime}}</em></small>
        </p>        
        <p>
          {{.Content}}
        </p>
      </article>
      {{end}}
    </div>
    <div id="floatright" class="span4">
      <div class="well well-small">
        <dl>
          <dt>关于本站</dt>
          <dd>本站使用了Google App Engine，由go语言编写；同时使用了Bootstrap框架。</dd>
          <p></p>
          <dt>关于本人</dt>
          <dd>本人是一名资质尚浅的ios开发人员，当然这说法可能很快要改变。</dd>
        </dl>
      </div>
      <div class="well well-small">
        <dl>
          <dt>新浪微博</dt>
          <dd><a href="http://weibo.com/zenol">weibo.com/zenol</a></dd>
            <dt>Facebook</dt>
            <dd><a href="http://www.facebook.com/zero.hero.lin">www.facebook.com/zero.hero.lin</a></dd>
            <dt>QQ</dt>
            <dd>108653682</dd>
            </dl>
      </div>
    </div><!-- id="floatright" -->
  </div><!-- end of .row -->
  </div>
  <!-- end of container -->
{{ end }}