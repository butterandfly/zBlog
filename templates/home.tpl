{{ define "container" }}
  <!-- container -->
  <div class="container">
  <div class="row">
    <div class="{{if .AppCube.NoRight}}span12{{else}}span8{{end}}">
      {{ with .PageCube.Articles}}
      {{range .}}
      <article class="well whitebg">
        <h2><a href="article?id={{.ID}}">{{.Title}}</a></h2>
        <p class="muted">
        <small><strong>Posted: </strong><em>{{niceTimeStr .CreateTime}}</em></small>
        </p>        
        <p>
          {{.Content}}
        </p>
      </article>
      {{end}}
      {{end}}
    </div>
    {{if not .AppCube.BlogCube.NoRight}}
    <div id="floatright" class="span4">
      {{range .PageCube.AFloats}}
      <div class="well well-small">
        {{.Content}}
      </div>
      {{end}}
      
    </div><!-- id="floatright" -->
    {{end}}
  </div><!-- end of .row -->
  </div>
  <!-- end of container -->
{{end}}
