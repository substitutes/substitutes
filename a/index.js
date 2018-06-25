if (wrk.g("class")) {
    window.location = "/c/" + wrk.g("class");
}

fetch("/api/").then(response => {
    return response.json();
}).then(data => {
    const el = wrk.$("select");
    data.forEach(c => {
        let item = document.createElement("option");
        item.text = c;
        el.add(item);
    });
    el.onchange = () => {
        wrk.s("class", el.selectedOptions[0].text);
        window.location = "/c/" + el.selectedOptions[0].text;
    };
});