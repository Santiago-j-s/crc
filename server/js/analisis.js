var analisis = (function () {
    var base_url = window.location.origin;
    var button;

    function inicializar() {
        button = document.getElementById('boton')
    }

    function post(url) {
        var xhr = new XMLHttpRequest();
        xhr.open('POST', url);
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        return xhr
    }

    function remover(elem) {
        if (elem !== null) {
            elem.parentNode.removeChild(elem)
        }
    }

    function showAnalisis(e) {
        e.preventDefault();

        var lista = document.getElementById('lista');
        remover(lista)

        var error = document.getElementsByName('error').item(0);
        remover(error);

        var url = encodeURI(base_url + '/hamming')
        var xhr = post(url)
        xhr.onload = function () {
            if (xhr.status !== 200) {
                p = "<p name='error'>" + xhr.response + "</p>";
                document.getElementById('errores').insertAdjacentHTML('afterbegin', p);
            } else {
                h = '<ul id="lista">'
                texts = xhr.responseText.split('\n')
                for (var index = 0; index < texts.length - 1; index++) {
                    h += '<li>' + texts[index] + '</li>'
                }
                h += '</ul>'
                document.getElementById('resultados').insertAdjacentHTML('afterbegin', h);
            }
        }

        var poly = document.getElementById('poly').value;
        var data = encodeURI('poly=' + poly)
        xhr.send(data);
    }

    function binds() {
        button.addEventListener('click', showAnalisis, 'false');
    }

    function init() {
        inicializar();
        binds();
    }

    return { init: init }
})();

window.onload = function () {
    analisis.init();
}