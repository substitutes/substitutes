let currentTeacher = document.location.href.substr(document.location.href.lastIndexOf("/") + 1);


fetch("/api/t/" + currentTeacher).catch(e => M.toast({html: e})).then(res => {
    return res.json();
}).then(data => {
    document.querySelector("#spinner").remove();
    document.querySelector("#heading").innerHTML = data.teacher + " (" + data.date.replace("Substitutes ", "") + ")";

    if (!data.substitutes) {
        M.toast({html: "The teacher does not have any substitutes"});
        return;
    }
    data.substitutes.forEach(substitute => {
        // TODO: Smart fill for this
        document.querySelector("tbody").innerHTML += "<tr class='text-lighten-2'><td>" +
            substitute.hour + "</td><td>" + substitute.classes + "</td><td>" + substitute.subject +
            "</td><td>" + substitute.room + "</td><td>" + substitute.type.replace("Vertretung", "Substitute") + "</td><td>" + substitute.teacher + "</td></tr>";
    });
});
