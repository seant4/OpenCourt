
function fetchCourts(){
    fetch('http://localhost:3000/courts')
    .then(res => res.json())
    .then((out) => { document.getElementById("courts").innerText = JSON.stringify(out); })
    .catch(err => console.log(err));
}

function postReservation(){
    let data = {court: 'One',
                reservee: 'Jon',
                date: '9-10-21',
                time: '9:21 AM'};

    fetch('http://localhost:3000/courts', {
        method: "POST",
        mode: 'cors',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    }).then(res => {
        console.log("request complete, response: ", res);
    });
}

fetchCourts();
postReservation();
