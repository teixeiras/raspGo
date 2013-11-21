
function AlertSystem() {
}

AlertSystem.prototype.alert = function(message) {
		var warning = $("<h4></h4>");
		warning.addClass("alert_warning");
		warning.html(message);
		warning.css({"position" : "fixed", "top" : "0px", "left" : "0px"});
	  	$("#main").append(warning);
	  	warning.delay(1000).fadeOut( 500 ,function () {
	  		warning.remove();
	  	});
	}

AlertSystem.prototype.error = function(message) {
		var warning = $("<h4></h4>");
		warning.addClass("alert_error");
		warning.html(message);
		warning.css({"position" : "absolute", "top" : "0px", "left" : "0px"});
	  	$("#main").append(warning);
	  	warning.delay(1000).fadeOut( 500 ,function () {
	  		warning.remove();
	  	});
}

AlertSystem.prototype.success = function (message) {
		var warning = $("<h4></h4>");
		warning.addClass("alert_success");
		warning.html(message);
		warning.css({"position" : "absolute", "top" : "0px", "left" : "0px"});
	  	$("#main").append(warning);
	  	warning.delay(1000).fadeOut( 500 ,function () {
	  		warning.remove();
	  	});
}

AlertSystem.prototype.informative = function(message) {
		var warning = $("<h4></h4>");
		warning.addClass("alert_info");
		warning.html(message);
		warning.css({"position" : "absolute", "top" : "0px", "left" : "0px"});
	  	$("#main").append(warning);
	  	warning.delay(1000).fadeOut( 500 ,function () {
	  		warning.remove();
	  	});
}

function Messages() {

};

Messages.prototype.messageHandler = function (message) {
	console.log(message);
	var object = JSON.parse(message);
	var alertSystem = new AlertSystem();
	alertSystem.informative(object);
}

var messagesObject = new Messages();


$(document).ready(function()  { 

		
  	    if (window["WebSocket"]) {
		    conn = new WebSocket("ws://"+window.location.host+"/ws");
		    conn.onclose = function(evt) {
		        
		    }
		    conn.onmessage = function(evt) {
		    	messagesObject.messageHandler(evt.data);
		    }

		    conn.onopen = function (event) {
		    	conn.send('{"module":"file.manager", "action": "listFiles", "args":"{}"}');
			}
		    	
		    

		} else {
		    appendLog($("<div><b>Your browser does not support WebSockets.</b></div>"))
		}
});