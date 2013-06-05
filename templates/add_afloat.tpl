{{ define "container"}}
<div class="container">
  <div class="span8">
    <article class="well">
      <form action="addaf" method="post">
        <label for="title">标题</label>
        <input type="text" name="title" id="title" />
        <label for="content">Html</label>
        <div>
          <textarea name="content" id="content" rows="20" style="width:640px;"></textarea>
        </div>
        <input type="submit" value="提交" />
      </form>
    </article>
  </div>
</div>
{{ end }}