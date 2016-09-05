function handleResponse(data){
	$(".working").slideUp();

	if(data.ok){
		$("#dl").attr("href", data.url);
		var link = data.url;
		if(navigator.platform.toUpperCase().indexOf('LINUX')>=0 || navigator.platform.toUpperCase().indexOf('WIN')>=0){
			link = "https://www.google.com/calendar/render?cid=" + link;
		} else {
			link = link.replace("http://", "webcal://")
		}
		$("#cal").attr('href', link);
		$(".success").slideDown();
	} else {
		$(".warning").show();
		$("form").slideDown();
	}
}

$(document).ready(function(){
	$("form").on("submit", function(e){
		e.preventDefault();
		$(this).slideUp();
		$(".working").slideDown();

		$.ajax({
			type: "POST",
			url: $(this).attr("action"),
			data: $(this).serialize(),
			success: handleResponse,
			dataType: "json"
		})
	})
});