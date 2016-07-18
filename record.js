var tabURL = ""

chrome.browserAction.onClicked.addListener(function(tab) {
    tabURL = tab.url
    chrome.tabs.executeScript(tab.id, {file: "jquery.js"}, function(){});
    chrome.tabs.sendMessage(tab.id, {text: 'report_back'}, recordTime);
});


// A function to use as callback
function recordTime(timeContent) {
    console.log('I received the following time: ' + timeContent);
    var postData = {
        "url": tabURL,
        "time": timeContent
    }
    var json_data = JSON.stringify(postData)
    $.ajax({
        type: "POST",
        url: "http://localhost:8080/record",
        data: json_data,
        success: function() {
            console.log("send time and url!");
        },
        error: function() {
            console.log("something went wrong with the send");
        }
    })

}