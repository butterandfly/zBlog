{{ define "container"}}
{{with .PageCube.AFloat}}
<div class="container">
  <div class="span8">
    <article class="well">
      <form action="editaf" method="post">
        <label for="title">标题</label>
        <input type="text" name="title" id="title" value="{{.Title}}" />
        <label for="content">HTML</label>
        <div>
          <textarea name="content" id="content" rows="20" style="width:640px;">{{.Content}}</textarea>
        </div>
        <input type="hidden" name="id" value="{{.ID}}" />
        <input type="submit" value="提交" />
      </form>
    </article>
  </div>
</div>
{{ end }}
{{ end }}