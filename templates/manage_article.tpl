{{ define "container" }}
<div class="container">
  <div class="row">
{{with .PageCube.Articles}}
  <div class="span8">
    <section class="well">
    <h4>文章列表：</h4>
    <form action="managearticle" method="post">
    <ul>
      {{ range . }}
      <li>
        <a href="article?id={{.ID}}" class="bigFont">
          <strong>{{.Title}}</strong>
        </a> 
        <small>Posted: <em>{{niceTimeStr .CreateTime}}</em></small>
        <button class="btn btn-mini btn-danger" type="submit" name="delete" value="{{.ID}}">删除</button>
        <a href="editarticle?id={{.ID}}" class="btn btn-mini btn-info">编辑</a>
      </li>
      {{ end }}
    </ul>
    </form>
    </section>
  </div> <!-- end of .span8 -->
{{ end }}
  <div class="span4">
    <section class="well">
    <a href="/addarticle" class="btn btn-large btn-primary btn-block">添加文章</a>
    </section>
  </div>
  </div>
</div>
{{ end }}