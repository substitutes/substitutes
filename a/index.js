(() => {
    const klassen = () => {
        const _x = {
            a: [9, 8, 7, 6, 5],
            b: ["a", "b", "c", "d"],
            c: ["12", "11", "10"]
        };
        $.each(_x.a, (a, b) => {
            $.each(_x.b, (x, y) => {
                _x.c.push(b + y);
            });
        });
        return _x.c;
    };
    const e = $('select');
    e.material_select();
    $.getJSON("/api", data => {
        $.each(data, (a, b) => {
            $("#heute").append("<option value='" + b + "'>" + b + "</option>");
        });
        $.each(klassen(), (a, b) => {
            $("#alle").append("<option value='" + b + "'>" + b + "</option>");
        });
        e.material_select();
        e.change(function () {
            window.location.pathname = "/k/" + $(this).val();
        });
    }).catch(e => console.log(e));
})();