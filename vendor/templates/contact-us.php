<?php
//If the form is submitted
if(isset($_POST['submit'])) {
	
	$comments = $_POST['message'];

	//Check to make sure that the name field is not empty
	if(trim($_POST['name']) == '') {
		$hasError = true;
	} else {
		$name = trim($_POST['name']);
	}


	//Check to make sure sure that a valid email address is submitted
	if(trim($_POST['email']) == '')  {
		$hasError = true;
	} else if (!eregi("^[A-Z0-9._%-]+@[A-Z0-9._%-]+\.[A-Z]{2,4}$", trim($_POST['email']))) {
		$hasError = true;
	} else {
		$email = trim($_POST['email']);
	}

	$website = trim($_POST['website']);

	//If there is no error, send the email
	if(!isset($hasError)) {
		$emailTo = 'raad.ajlouni@raconta.com'; //Put your own email address here
		$body = "Name: $name \n\nEmail: $email \n\nWebsite: $website \n\nComments:\n $comments";
		$headers = 'From: My Site: '. $emailTo . "\r\n" . 'Reply-To: ' . $email;

		mail($emailTo, 'Glocal', $body, $headers);
		$emailSent = true;
	}
}
?>
<!doctype html>
<html class="" lang="en">

	<head>

		<meta charset="utf-8">
		<title>Glocal</title>

		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no" />

		<link rel="stylesheet" href="css/jquery-ui-1.10.3.custom.min.css" />
		<link rel="stylesheet" href="css/jquery-selectbox.css" />
		<link rel="stylesheet" href="css/responsive-grid.css" />
		<link rel="stylesheet" href="css/styles.css" />
		<link id="theme-color" rel="stylesheet" href="css/green.css" />
		<link rel="stylesheet" href="css/header-1.css" />
		<link rel="stylesheet" href="css/responsive.css" />

		<!--[if lt IE 9]>
			<link rel="stylesheet" href="css/styles-ie8-and-down.css" />
		<![endif]-->

	</head>

	<body>

		<div id="theme-switcher" class="theme-switcher">
			<div id="switcher-toggle-button" class="switcher-toggle-button"></div>
			<div class="label">Colors:</div>
			<ul id="color-switcher" class="color-switcher">
				<li class="green"></li>
				<li class="red"></li>
				<li class="blue"></li>
				<li class="grey"></li>
				<li class="violet"></li>
				<li class="orange"></li>
			</ul>
			<div class="label">Layout:</div>
			<select id="layout-switcher" class="layout-switcher" name="layout-switcher">
				<option selected="" value="fullscreen">Fullscreen</option>
				<option value="boxed">Boxed</option>
			</select>
			<div class="label">Background:</div>
			<ul id="background-switcher" class="background-switcher">
				<li class="background-1"></li>
				<li class="background-2"></li>
				<li class="background-3"></li>
			</ul>
		</div>

		<header class="section">

			<div class="header-top-wrapper">
				<div class="zone-header-top zone clearfix">

					<div class="header-top-left container-8">

						<div class="user-links">
							<div class="login">
								<a href="#" id="login-link" class="login-link">Login</a>
								<form id="login-form" class="login-form">
									<input class="text-input-grey" type="text" placeholder="Login">
									<input class="text-input-grey" type="text" placeholder="Password">
									<a href="#" class="password-restore">Forgot Password?</a>
									<input class="button-2-colorful" type="submit" value="Login">
								</form>
							</div>
							<div class="register">
								<a href="#" id="register-link" class="register-link">Register</a>
								<form id="register-form" class="register-form">
									<input class="text-input-grey" type="text" placeholder="Login">
									<input class="text-input-grey" type="text" placeholder="Company Name">
									<input class="text-input-grey" type="text" placeholder="Contact Person Name">
									<input class="text-input-grey" type="text" placeholder="E-mail">
									<input class="text-input-grey" type="text" placeholder="Password">
									<input class="text-input-grey" type="text" placeholder="Repeat Password">
									<div href="#" class="fields-required">All Fields Are Required</div>
									<input class="button-2-colorful" type="submit" value="Register">
								</form>
							</div>
						</div>

					</div>

					<div class="header-top-right container-16">

						<div class="social-links block">
							<a href="http://www.facebook.com">
								<img src="images/facebook-icon.png" alt=""/>
							</a>
							<a href="http://www.google.com">
								<img src="images/google-icon.png" alt="" />
							</a>
							<a href="http://www.twitter.com">
								<img src="images/twitter-icon.png" alt="" />
							</a>
							<a href="http://www.linkedin.com">
								<img src="images/linkedin-icon.png" alt="" />
							</a>
							<a href="http://www.pinterest.com">
								<img src="images/pinterest-icon.png" alt="" />
							</a>
							<a href="http://www.dribbble.com">
								<img src="images/dribbble-icon.png" alt="" />
							</a>
						</div>

						<div class="languages block">
							<a href="#" class="en current">EN</a>
							<a href="#" class="fr">FR</a>
							<a href="#" class="it">IT</a>
							<a href="#" class="de">DE</a>
						</div>

					</div>

				</div><!-- end of .zone-header-top -->
			</div><!-- end of .header-top-wrapper -->

			<div class="header-wrapper">
				<div class="zone-header zone clearfix">

					<div class="header-left container-6">

						<div class="logo block">
							<a href="index.html">
								<img src="images/logo.png" alt="" />
							</a>
						</div>

					</div>

					<div class="header-right container-18">

						<div class="adverisement block">
							<a href="http://themeforest.net/">
								<img src="images/content/banner.jpg" alt="" />
							</a>
						</div>

					</div>

				</div><!-- end of .zone-header -->
			</div><!-- end of .header-wrapper -->

			<div class="main-menu-wrapper">
				<div class="zone-main-menu zone clearfix">

					<div class="main-menu-container container-24">

						<div class="main-menu block">
							<ul id="sf-menu">
								<li class="empty">
									<div></div>
								</li>
								<li class="first">
									<a href="index.html">HOME</a>
									<ul>
										<li class="first">
											<a href="index-street-view.html">Home (Street View)</a>
										</li>
										<li class="">
											<a href="index-slideshow.html">Home (Slideshow)</a>
										</li>
										<li class="">
											<a href="index-header2.html">Home (Header 2)</a>
										</li>
										<li class="">
											<a href="index-header3.html">Home (Header 3)</a>
										</li>
										<li class="">
											<a href="index-header4.html">Home (Header 4)</a>
										</li>
										<li class="">
											<a href="index-header5.html">Home (Header 5)</a>
										</li>
										<li class="">
											<a href="index-header6.html">Home (Header 6)</a>
										</li>
										<li class="last">
											<a href="index-header7.html">Home (Header 7)</a>
										</li>
									</ul>
								</li>
								<li class="">
									<a href="companies-listing.html">FEATURES</a>
									<ul>
										<li class="first">
											<a href="companies-listing.html">Companies Listing</a>
										</li>
										<li class="">
											<a href="shortcodes.html">Shortcodes</a>
										</li>
										<li class="">
											<a href="company-page.html">Company Page</a>
										</li>
										<li class="">
											<a href="company-page-2.html">Company Page 2</a>
										</li>
										<li class="">
											<a href="company-page-3.html">Company Page (Complex Ratings)</a>
										</li>
										<li class="">
											<a href="company-tabs.html">Company Page (With Tabs)</a>
										</li>
										<li class="">
											<a href="single-project.html">Single Project</a>
										</li>
										<li class="">
											<a href="portfolio-1.html">Portfolio (1 column)</a>
										</li>
										<li class="">
											<a href="portfolio-2.html">Portfolio (2 columns)</a>
										</li>
										<li class="">
											<a href="portfolio-3.html">Portfolio (3 columns)</a>
										</li>
										<li class="last">
											<a href="portfolio-4.html">Portfolio (4 columns)</a>
										</li>
									</ul>
								</li>
								<li class="">
									<a href="about-us.html">ABOUT US</a>
								</li>
								<li class="">
									<a href="price-register.html">SUBMIT LISTING</a>
								</li>
								<li class="neighbour-left">
									<a href="blog.html">BLOG</a>
									<ul>
										<li class="first last">
											<a href="blog-post.html">Blog post</a>
										</li>
									</ul>
								</li>
								<li class="active">
									<a href="contact-us.php">CONTACT US</a>
								</li>
								<li class="last neighbour-right">
									<a href="index.html">PURCHASE THIS THEME</a>
								</li>
								<li class="empty">
									<div></div>
								</li>
							</ul>
						</div>

					</div>

				</div><!-- end of .zone-main-menu -->
			</div><!-- end of .main-menu-wrapper -->

		</header>

		<section class="section content">

			<div class="content-wrapper">
				<div class="zone-content equalize zone clearfix">

					<div class="content-container container-16">

						<div class="contact-us block">
							<div class="block-title">
								<h1>Contact Us</h1>
							</div>
							<div class="comment-message">
								<div class="comment-message-title">
									Send Us a <span class="text-colorful">Message</span>
								</div>
								<form method="post" action="<?php echo $_SERVER['PHP_SELF']; ?>" id="contact-us-form" class="comment-message-form">
									<input type="text" name="name" class="text-input-grey name" placeholder="Name *" />
									<input type="text" name="email" class="text-input-grey email" placeholder="Email *" />
									<input type="text" name="website" class="text-input-grey website" placeholder="Website" />
									<textarea name="message" class="text-input-grey comment-message-main" placeholder="Your Comments Here"></textarea>
									<input type="submit" name="submit" value="Send Message" class="button-2-colorful" />
									<?php if(isset($hasError)) { //If errors are found ?>
										<div class="form-message text-colorful">Please check if you've filled all the fields with valid information. Thank you.</div>
									<?php } ?>
									<?php if(isset($emailSent) && $emailSent == true) { //If email is sent ?>
										<div class="form-message text-colorful">Thank you! Email successfully sent.</div>
									<?php } ?>
								</form>
							</div>
						</div>

						<div class="separator"></div>

						<div class="company-details block">
							<div class="company-address">
								<div class="details-title">Address Details:</div>
								<div class="detail address">1234 Street<br />Mountain View, CA 94043</div>
								<div class="detail phone">Phone: +1 123-456-7890<br />Fax: +1 123-456-7890</div>
								<div class="detail email">
									E-mail: <a href="mailto:email@example.com" class="text-colorful">email@example.com</a><br />Website: <a href="http://themeforest.net/" class="text-colorful">www.example.com</a>
								</div>
							</div>
							<div class="company-hours">
								<div class="details-title">Opening Hours:</div>
								<div class="detail">
									<span class="detail-label">Monday-Friday:</span>9am - 5pm
								</div>
								<div class="detail">
									<span class="detail-label">Saturday:</span>10am - 3pm
								</div>
								<div class="detail">
									<span class="detail-label">Sunday:</span>Closed
								</div>
							</div>
							<div class="clearfix"></div>
							<div class="company-map">
								<iframe src="https://maps.google.com/maps?q=33.874976,-117.566814&amp;num=1&amp;ie=UTF8&amp;t=m&amp;ll=33.878112,-117.566414&amp;spn=0.054727,0.20977&amp;z=12&amp;output=embed"></iframe>
							</div>
						</div>

					</div><!-- end of .content-container -->

					<div class="sidebar-container container-8">

						<div class="recently-added block">
							<div class="block-title">
								<h3>Recently Added</h3>
							</div>
							<ul class="entries-list">
								<li class="clearfix">
									<a href="#" class="thumbnail">
										<img src="images/content/sky.png" alt="" />
									</a>
									<a href="#" class="entry-title">Company Name</a>
									<div class="entry-excerpt">Lorem ipsum dolor sit amet, consectetur adipisicing elit.</div>
								</li>
								<li class="clearfix">
									<a href="#" class="thumbnail">
										<img src="images/content/text.png" alt="" />
									</a>
									<a href="#" class="entry-title">Another Company</a>
									<div class="entry-excerpt">Lorem ipsum dolor sit amet, consectetur adipisicing elit.</div>
								</li>
							</ul>
							<div class="two-images-banner clearfix">
								<a href="#">
									<img src="images/content/crayons.png" alt="" />
								</a>
								<a href="#">
									<img src="images/content/coins.png" alt="" />
								</a>
							</div>
						</div>

						<div class="latest-news block">
							<div class="block-title">
								<h3>Latest News</h3>
							</div>
							<ul class="entries-list">
								<li class="clearfix">
									<a href="#" class="thumbnail">
										<img src="images/content/coins.png" alt="" />
									</a>
									<a href="#" class="entry-title">Lorem Ipsum</a>
									<div class="entry-excerpt">Lorem ipsum dolor sit amet, consectetur adipisicing elit.</div>
								</li>
								<li class="clearfix">
									<a href="#" class="thumbnail">
										<img src="images/content/crayons.png" alt="" />
									</a>
									<a href="#" class="entry-title">Dolor Sit Amet</a>
									<div class="entry-excerpt">Lorem ipsum dolor sit amet, consectetur adipisicing elit.</div>
								</li>
							</ul>
							<div class="one-image-banner">
								<a href="#">
									<img src="images/content/handshake.png" alt="" />
								</a>
							</div>
						</div>

					</div><!-- end of .sidebar-container -->

				</div><!-- end of .zone-content -->
			</div><!-- end of .content-wrapper -->

		</section>

		<footer class="section">

			<div class="footer-wrapper">
				<div class="zone-footer zone clearfix">

					<div class="footer-container container-24">

						<div class="website-short-description block">
							<img src="images/logo-small.png" class="logo" alt="" />
							<div class="description-text">
								Donec venenatis, turpis vel hendrerit interdum, dui ligula ultricies purus, sed posuere libero dui id orci. Nam congue, pede vitae dapibus aliquet, elit magna vulputate arcu, vel tempus metus leo non est. Etiam sit amet lectus quis est congue mollis.
							</div>
						</div>

						<div class="twitter-feed block">
							<h3 class="title">Recent Tweets</h3>
							<div id="twitter-feed"></div>
						</div>

						<div class="recent-posts block">
							<h3 class="title">Recent Posts</h3>
							<ul>
								<li class="first">
									<a href="#" class="text-colorful">Lorem ipsum dolor sit amet</a>
								</li>
								<li>
									<a href="#" class="text-colorful">Proin nibh augue suscipit</a>
								</li>
								<li>
									<a href="#" class="text-colorful">Cras vel lorem</a>
								</li>
								<li class="last">
									<a href="#" class="text-colorful">Quisque semper justo at risus</a>
								</li>
							</ul>
						</div>

						<div class="flickr-feed block">
							<h3 class="title">Flickr Feed</h3>
							<div id="flickr-feed"></div>
						</div>

					</div>

				</div><!-- end of .zone-footer -->
			</div><!-- end of .footer-wrapper -->

			<div class="copyright-wrapper">
				<div class="zone-copyright zone clearfix">

					<div class="copyright-left-container container-12">

						<div class="copyright block">&copy; All Rights Reserved. Design by www.uou.ch.</div>

					</div>

					<div class="copyright-right-container container-12">

						<div class="social-links block">
							<a href="http://www.facebook.com">
								<img src="images/facebook-icon.png" alt="" />
							</a>
							<a href="http://www.google.com">
								<img src="images/google-icon.png" alt="" />
							</a>
							<a href="http://www.twitter.com">
								<img src="images/twitter-icon.png" alt="" />
							</a>
							<a href="http://www.linkedin.com">
								<img src="images/linkedin-icon.png" alt="" />
							</a>
							<a href="http://www.pinterest.com">
								<img src="images/pinterest-icon.png" alt="" />
							</a>
							<a href="http://www.dribbble.com">
								<img src="images/dribbble-icon.png" alt="" />
							</a>
						</div>

					</div>

				</div><!-- end of .zone-copyright -->
			</div><!-- end of .copyright-wrapper -->

		</footer>

		<div id="boxed-switch" class="boxed-switch">SWITCH TO BOXED VERSION</div>

		<script type="text/javascript" src="scripts/jquery-1.10.2.min.js"></script>
		<script type="text/javascript" src="scripts/jquery-ui-1.10.3.custom.min.js"></script>
		<script type="text/javascript" src="scripts/jquery.colorbox-min.js"></script>
		<script type="text/javascript" src="scripts/jquery.selectbox-0.6.1.js"></script>
		<script type="text/javascript" src="scripts/jquery.tweet.js"></script>
		<script type="text/javascript" src="scripts/jflickrfeed.min.js"></script>
		<script type="text/javascript" src="scripts/superfish.js"></script>
		<script type="text/javascript" src="scripts/jquery.mobilemenu.min.js"></script>
		<script type="text/javascript" src="scripts/jquery.placeholder.min.js"></script>
		<script type="text/javascript" src="scripts/scripts.js"></script>

	</body>

</html>