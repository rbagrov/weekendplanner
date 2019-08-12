/* global document */
jQuery(document).ready(function(){

	/***
	Adding Google Map.
	***/

	/* Calling goMap() function, initializing map and adding markers. */
	jQuery('#map').goMap({
		maptype: 'ROADMAP',
        latitude: 43.057006,
        longitude: 24.950218, 
        zoom: 8,
        scaleControl: false,
        scrollwheel: false,
		markers: [
		]
	});

	/* Hiding all the markers on the map. */
	for (var i in $.goMap.markers) {
		if (this[i] !== 0) {
			$.goMap.showHideMarker(jQuery.goMap.markers[i], false);
		}
	}

	/* Revealing markers from the first group - 'airport' */
	$.goMap.showHideMarkerByGroup('airport', true);

	/* Processing clicks on the tabs under the map. Revealing corresponding to each tab markers. */
	jQuery('#industries-tabs ul li a').click(function(event) {
		/* Preventing default link action */
		event.preventDefault();
		/* Getting current marker group name. Link ID's and marker group names must coincide. */
		var markerGroup = jQuery(this).attr('id');
		/* Changing current active tab. */
		jQuery('#industries-tabs ul li').removeClass('active');
		jQuery(this).parent().addClass('active');
		/* Hiding all the markers on the map. */
		for (var i in jQuery.goMap.markers) {
			if (this[i] !== 0) {
				jQuery.goMap.showHideMarker(jQuery.goMap.markers[i], false);
			}
		}
		/* Revealing markers from the corresponding group. */
		jQuery.goMap.showHideMarkerByGroup(markerGroup, true);
	});

});
