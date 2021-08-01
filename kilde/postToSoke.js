// Send the entire html page to the Index Html endpoint of sokemotor.

var htmlString = document.getElementsByTagName("html")[0].innerHTML
var url = document.URL
console.log(url)

// POST method implementation:
async function postData(url = '', data = {}) {
    const response = await fetch(url, {
      method: 'POST', 
      mode: 'cors', 
      cache: 'no-cache', 
      credentials: 'same-origin', 
      headers: {
        'Content-Type': 'application/json'
      },
      redirect: 'follow', 
      referrerPolicy: 'no-referrer', 
      body: JSON.stringify(data) 
    });
    return response.json();
  }
  
postData('http://localhost:8080/indexHtml', { "url" : url, "dom": htmlString})
    .then(data => {
      console.log(data); // JSON data parsed by `data.json()` call
    });
  