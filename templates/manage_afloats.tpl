{{ define "container" }}
<div class="container">
  <div class="row">
  <div class="span8">
    <section class="well">
    <h4>公告列表：</h4>
    <form action="manageaf" method="post">
    {{with .PageCube.AFloats}}
    <ul>
      {{ range . }}
      <li>
        <a href="article?id={{.ID}}" class="bigFont">
          <strong>{{.Title}}</strong>
        </a> 
        <small>Posted: <em>{{niceTimeStr .CreateTime}}</em></small>
        <button class="btn btn-mini btn-danger" type="submit" name="delete" value="{{.ID}}">删除</button>
        <a href="editaf?id={{.ID}}" class="btn btn-mini btn-info">编辑</a>
      </li>
      {{ end }}
    </ul>
    {{ end }}
    </form>
    </section>
  </div> <!-- end of .span8 -->
  <div class="span4">
    <section class="well">
    <a href="/addaf" class="btn btn-large btn-primary btn-block">添加公告</a>
    </section>
  </div>
  </div>
</div>
{{ end }}