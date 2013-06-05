{{ define "container" }}
{{with .PageCube.Msg}}
<div class="container" >
  <div class="alert alert-success" >
  <h5>
    {{.}}
  </h5>
  </div>
</div>
{{ end }}
{{ end }}