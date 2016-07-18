
chrome.runtime.onMessage.addListener(
	function (msg, sender, sendResponse) {
	    if (msg.text === 'report_back') {
			if (document.getElementsByClassName("ytp-time-current").length > 0) {
		        sendResponse(document.getElementsByClassName("ytp-time-current")[0].innerHTML);
			};
			
	    }
	    return true;
});



