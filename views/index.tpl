<a class="btn btn-default" href="#" role="button">Login</a>
<!-- <a class="btn btn-default" href="/users/new" role="button">Sign Up</a> -->
<a class="btn btn-primary btn-lg" href="/users/new" role="button" data-toggle="modal" data-target="#myModalNorm">
	Sign Up
</a>


<html>
<head>
	<title>Sample Application</title>
	<!-- Bootstrap -->
	<link rel="stylesheet" href="/static/css/bootstrap.min.css">
	<link rel="stylesheet" href="/static/css/bootstrap-theme.min.css">
	<!-- <script src="/static/js/bootstrap.min.js"></script> -->
	<script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
</head>
<body>
	<!-- Modal -->
	<div class="modal fade" id="myModalNorm" tabindex="-1" role="dialog" 
	aria-labelledby="myModalLabel" aria-hidden="true">
	<div class="modal-dialog">
		<div class="modal-content">
			<!-- Modal Header -->
			<div class="modal-header">
				<button type="button" class="close" 
				data-dismiss="modal">
				<span aria-hidden="true">&times;</span>
				<span class="sr-only">Close</span>
			</button>
			<h4 class="modal-title" id="myModalLabel">
				Sign Up
			</h4>
		</div>

		<!-- Modal Body -->
		<div class="modal-body">

			<form accept-charset="utf-8" role="form" class="form-horizontal" method="POST" action='{{urlfor "UsersController.Create"}}'>
				<div class="form-group">
					<label for="exampleInputUsername">Username</label>
					<input type="text" class="form-control" id="username" placeholder="Enter Username" name="username">
				</div>
				<div class="form-group">
					<label for="exampleInputPassword1">Password</label>
					<input type="password" class="form-control" id="exampleInputPassword1" placeholder="Enter Password" name="password">
				</div>
				<div class="form-group">
					<label for="exampleInputPasswordConfirmation1">Password Confirmation</label>
					<input type="password" class="form-control" id="exampleInputPasswordConfirmation1" placeholder="Password" name="password_confirmation">
				</div>
				<button type="submit" class="btn btn-default">Submit</button>
			</form>
		</body>
	</div>
</div>
</div>
</div>
</html>
