<html>
<head>
<title></title>
</head>
<body>
<form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='{{urlfor "UsersController.Post"}}'>
    Username: <input type="text" name="username"><br />
    Password: <input type="password" name="password"><br />
    Password Confirmation: <input type="password" name="password_confirmation">
    <input type="submit" value="Sign Up">
</form>
</body>
</html>