fetch("/api/").catch(e => M.toast({html: e})).then(response => {
    return response.json();
}).then(data => {
    const el = document.querySelector(".collection");
    data.forEach(c => {
        el.innerHTML += "<a class='collection-item indigo-text' href='/c/" + c + "'>" + c + "</a>";
    });
    el.onchange = () => {
        window.location = "/c/" + el.selectedOptions[0].text;
    };
    M.AutoInit(document.body);
});
