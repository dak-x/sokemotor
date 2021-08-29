// Send the entire html page to the Index Html endpoint of sokemotor.

var htmlString = document.getElementsByTagName("html")[0].innerText
var url = document.URL
var today = new Date()

jsonData = {
  "url": url,
  "dom": htmlString,
  "lastaccessed": today.toISOString()
}

console.log(jsonData)

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


postData('http://localhost:8080/indexHTML', jsonData)
  .then(data => {
    console.log(data); // JSON data parsed by `data.json()` call
  });
