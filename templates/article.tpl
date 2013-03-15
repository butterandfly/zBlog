{{ define "container" }}
  <!-- container -->
  <div class="container">
  <div class="row">
    <div class="span9">
      <article class="well whitebg">
        <h2><a href="article/?id={{.Id}}">{{.Title}}</a></h2>
        <p class="muted">
        <small><strong>Posted: </strong><em>{{.CreateTime}}</em></small>
        </p>        
        <p>
          {{.Content}}
        </p>
      </article>
    </div>
  </div><!-- end of .row -->
  </div>
  <!-- end of container -->
{{ end }}