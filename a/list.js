(() => {
    $.getJSON("/api/c/{{ .class }}", data => {
        $.each(data.substitutes, (a, b) => {
            (b.type === "Cancelled") ? $("tbody").append(
                "<tr class='teal-text text-lighten-2'><td>" +
                b.hour + "</td><td>" + b.time + "</td><td>" + b.teacher +
                "</td><td>" + b.subject + "</td><td>" + b.room + "</td><td>" + b.type.replace("Vertretung", "Substitute") +
                "</td><td>" + b.reason + "</td></tr>") : $("tbody").append(
                "<tr><td>" + b.hour + "</td><td>" + b.time + "</td><td>" +
                b.teacher +
                "</td><td>" + b.subject + "</td><td>" + b.room + "</td><td>" + b.type.replace("Vertretung", "Substitute") +
                "</td><td>" + b.reason + "</td></tr>");
        });
        $("h4").html(data.meta.date.replace("Vertretungen", "Substitutes").split("/")[0]);
        $("#title").html(data.meta.class);
    }).catch(m => Materialize.toast(m.status + ": " + m.responseJSON.message));
});
