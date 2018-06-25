wrk.$("#clear").onclick = ev => {
    wrk.s("class", "");
    window.location = '/';
};

fetch("/api/c/" + window.location.pathname.split('/')[window.location.pathname.split('/').length-1]).then(res => {
    return res.json();
}).then(data => {
    document.querySelector("#title").innerText = data.meta.date.replace("Vertretungen", "Substitutes").split("/")[0];
    document.querySelector("#subtitle").innerHTML = data.meta.class + ", " + data.meta.date.split("/")[1];
    if (!data.meta.extended)
        Array.from(document.getElementsByClassName("hide-extended")).forEach(a => a.remove());
    data.substitutes.forEach(substitute => {
        // TODO: Smart fill for this
        if (data.meta.extended) {
            document.querySelector("tbody").innerHTML += "<tr class='text-lighten-2'><td>" +
                substitute.hour + "</td><td>" + substitute.time + "</td><td>" + substitute.teacher.replace("?", " => ") +
                "</td><td>" + substitute.subject.replace("?", " => ") + "</td><td>" + substitute.room + "</td><td>" + substitute.type.replace("Vertretung", "Substitute") +
                "</td><td>" + substitute.notes + "</td><td>" + substitute.reason + "</td></tr>";
        } else {
            document.querySelector("tbody").innerHTML += "<tr class='text-lighten-2'><td>" +
                substitute.hour + "</td><td>" + substitute.teacher.replace("?", " => ") +
                "</td><td>" + substitute.subject.replace("?", " => ") + "</td><td>" + substitute.room + "</td><td>" + substitute.type.replace("Vertretung", "Substitute") +
                "</td><td>" + substitute.notes + "</td></tr>";
        }
    });
});