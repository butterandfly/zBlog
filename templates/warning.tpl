{{ define "container" }}
{{with .PageCube.Msg}}
<div class="container" >
  <div class="alert alert-error" >
  <h5>
    {{.}}
  </h5>
  </div>
</div>
{{ end }}
{{ end }}