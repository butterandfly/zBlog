{{ define "container" }}
<div class="container">
  <div class="well">
<form action="addadmin" method="post">
  <label for="acount">请输入作为管理员的google账户：</label>
  <input type="text" name="acount" />@gmail.com
  <br />
  <input type="submit" value="确定" />
</form>  
</div>
</div>
{{ end }}