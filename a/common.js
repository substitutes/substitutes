// substitutes - common lib - @fronbasal

const wrk = {
    $: a => document.querySelector(a),
    e$: localStorage != null,
    m$: () => wrk.i("[dev] localStorage not available."),
    i: console.info,
    s: (a, b) => wrk.e$ ? localStorage.setItem(a, b) : wrk.m$(),
    g: a => wrk.e$ ? localStorage.getItem(a) : wrk.m$()
};