fetch("/api/").then(response => {
    return response.json();
}).then(data => {
    const el = document.querySelector("select");
    data.forEach(c => {
        let item = document.createElement("option");
        item.text = c;
        el.add(item);
    });
    el.onchange = () => {
        window.location = "/c/" + el.selectedOptions[0].text;
    };
    M.AutoInit(document.body);
});
