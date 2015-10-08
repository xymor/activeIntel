var node = document.createElement("img");
node.src = "http://tracking-wolf.mybluemix.net/track.gif?"+JSON.stringify(dataLayer); 
node.id = 'trkngPx21212';
document.body.appendChild(node);
setTimeout(function() { document.getElementById('trkngPx21212').style.display='none'; }, 1);