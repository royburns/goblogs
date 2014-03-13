<div class="row-fluid">
                    
<div class="btn-toolbar">
    <a class="btn btn-primary" href="/admin/addblog"><i class="icon-plus"></i>添加博客</a>
  <div class="btn-group">
  </div>
</div>
<div class="well">
    <table class="table">
      <thead>
        <tr>
          <th>博客id</th>
          <th>博客标题</th>
		  <th>博客发布时间</th>
          <th style="width: 26px;"></th>
        </tr>
      </thead>
      <tbody>
        {{range .Blogs}}
          <tr>
	 
          <td>{{.id}}</td>
          <td>{{.title}}</td>
		  <td>{{.created}}</td>
          <td>
              <a href="/admin/editblog/{{.id}}"><i class="icon-pencil"></i></a>
              <a href="/admin/delblog/{{.id}}" role="button" data-toggle="modal"><i class="icon-remove"></i></a>
          </td>
	        
        </tr>
        {{end}}
      </tbody>
    </table>
</div>


<div class="pagination">

		{{.PageStr}}
    
</div>

