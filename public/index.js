/*
 *  :)
 *  lmk what cool ways you guys hack my little project
 */

var getJSON = function(url, callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'json';
    xhr.onload = function() {
      var status = xhr.status;
      if (status === 200) {
        callback(null, xhr.response);
      } else {
        callback(status, xhr.response);
      }
    };
    xhr.send();
};

const leaderBoard = document.getElementById("leaders");

getJSON("data", (s, data) => {
  if(s !== null){
    console.log(`Error: ${s}`)
  } else {
    data.forEach((address,ind) => {
      const listItem = document.createElement('li');
      listItem.appendChild(document.createTextNode(`${address[1]} - ${address[0]}`));

      if(ind >= 10){
        listItem.style.opacity = 0.5;
        listItem.hidden = true;
      }

      leaderBoard.appendChild(listItem);
    })
    
    document.getElementById("loading").remove();
  }
})

const registerName = document.getElementById("name");

function register() {
  window.location.assign("play?name="+registerName.value)
}

function unhide() {
  document.querySelectorAll("li").forEach(item => {
    item.hidden = false;
  });

  document.getElementById("unhide").remove();
}
