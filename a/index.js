$(() => {
    $('select').material_select();
    $.getJSON("/api", data => {
        $.each(data, (a, b) => {
            $("select").append("<option value='" + b + "'>" + b + "</option>");
        });
        $('select').material_select();
        $("select").change(() => window.location.pathname = "/k/" + $("select").val());
    }).catch(e => console.log(e));
});