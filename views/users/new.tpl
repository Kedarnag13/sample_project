<html>
<head>
<title>Sample Application</title>
</head>
<body>
<form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='{{urlfor "UsersController.Post"}}'>
  <div class="form-group">
    <label for="exampleInputUsername">Username</label>
    <input type="text" class="form-control" id="username" placeholder="Email" name="username">
  </div>
  <div class="form-group">
    <label for="exampleInputPassword1">Password</label>
    <input type="password" class="form-control" id="exampleInputPassword1" placeholder="Password" name="password">
  </div>
  <div class="form-group">
    <label for="exampleInputPasswordConfirmation1">Password Confirmation</label>
    <input type="password" class="form-control" id="exampleInputPasswordConfirmation1" placeholder="Password" name="password_confirmation">
  </div>
  <button type="submit" class="btn btn-default">Submit</button>
</form>
</body>
</html>