
function fetchCourts(){
    fetch('http://localhost:3000/courts')
    .then(res => res.json())
    .then((out) => {
        console.log(out);
        const element = document.getElementById("courts")
        const br = document.createElement("br");
        for(let i = 0; i < out.length; i++){
            let para = document.createElement("p");
            let node = document.createTextNode("Name: " + JSON.stringify(out[i].name));
            para.appendChild(node);
            para.appendChild(br);
            node = document.createTextNode("Location: " + JSON.stringify(out[i].location));
            para.appendChild(node);
            para.appendChild(br);
            node = document.createTextNode("Reservations: " + JSON.stringify(out[i].reserved));
            para.appendChild(node);
            para.appendChild(br);
            element.appendChild(para);
        }
    })
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
