
function fetchCourts(){
    fetch('http://localhost:3000/courts')
    .then(res => res.json())
    .then((out) => { document.getElementById("courts").innerText = JSON.stringify(out); })
    .catch(err => console.log(err));
}

function postTime(courtName){
    let data = {court: courtName,
                reservee: 'Deez',
                date: '9-12-10',
                time: '8:43 PM'};
    
    postReservation(data);
}

function postReservation(data){

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
