// i dont know why reservee is captial lol ill find it later B)
function fetchCourts(){
    fetch('http://localhost:3000/courts')
    .then(res => res.json())
    .then((out) => {
        console.log(out);
        const element = document.getElementById("courts")
        const br = document.createElement("br");
        for(let i = 0; i < out.length; i++){
                let para = document.createElement("p");
                let br = document.createElement("br");
                //Create button information
                let btn = document.createElement("button");
                btn.innerHTML = "Reserve " + JSON.stringify(out[i].name);
                btn.onclick = function () {
                    var reservee = document.getElementById("fname").value;
                    var date = document.querySelector('input[type="date"]').value;
                    var time = document.querySelector('input[type="time"]').value;
                    var court = out[i].name;
                    postTime(court, reservee, date, time)
                }
                //Create court information
                let node = document.createTextNode(" Name: " + JSON.stringify(out[i].name));
                para.appendChild(btn);
                para.appendChild(node);
                element.appendChild(br);
                node = document.createTextNode(" Location: " + JSON.stringify(out[i].location));
                para.appendChild(node);
                element.appendChild(br);
                reservationsLabel = document.createTextNode(" Reservations: ")
                para.appendChild(reservationsLabel);
                for( let j = 0; j < out[i].reserved.length; j++){
                    if(out[i].reserved[j].Reservee === "None"){
                        ;
                    }else{
                        node = document.createTextNode(" " + JSON.stringify(out[i].reserved[j].Reservee) + " on: " + JSON.stringify(out[i].reserved[j].date) + " at: " + JSON.stringify(out[i].reserved[j].time) + ", ");
                        para.appendChild(node);
                        element.appendChild(br);
                    }
                }
                //-------------------------------
                element.appendChild(para);       
        }
    })
    .catch(err => console.log(err));
}

function postTime(courtName, reserveeInput, dateInput, timeInput){
    let data = {court: courtName,
                reservee: reserveeInput,
                date: dateInput,
                time: timeInput};

    console.log(data);

    if(reserveeInput == "" || dateInput == "" || timeInput ==""){
        window.alert("ERROR: Please make sure to input a name, email, time and date!");
    }else{
        postReservation(data);
    }
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
