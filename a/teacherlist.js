fetch("/api/t/").catch(e => M.toast({html: e})).then(response => {
    return response.json();
}).then(data => {
    document.querySelector("#spinner").remove();
    document.querySelector("#heading").innerHTML = data.date.replace("Substitutes", "Teachers");
    const el = document.querySelector(".collection");
    data.teachers.forEach(c => {
        el.innerHTML += "<a class='collection-item indigo-text' href='/t/" + c + "'>" + c + "</a>";
    });
    M.AutoInit(document.body);
});
