{{ define "container" }}
  <!-- container -->
  <div class="container">
  <div class="row">
{{with .PageCube.Article}}
    <div class="span9">
      <article class="well whitebg">
        <h2><a href="article/?id={{.ID}}">{{.Title}}</a></h2>
        <p class="muted">
        <small><strong>Posted: </strong><em>{{niceTimeStr .CreateTime}}</em></small>
        </p>        
        <p>
          {{.Content}}
        </p>
      </article>
    </div>
  {{end}}
  </div><!-- end of .row -->
  </div>
  <!-- end of container -->
{{ end }}