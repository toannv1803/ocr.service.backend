<!DOCTYPE html>
<html lang="en">
  <head>
    <title>File Upload</title>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="http://localhost:2020/upload"
      method="post"
    >
      <input type="file" name="file" />
      <input type="submit" value="upload" />
    </form>
  </body>
</html>