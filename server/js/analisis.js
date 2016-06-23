window.onload = function () {
    button = document.getElementById("boton");
    var base_url = window.location.origin;

    function showAnalisis(e) {
        e.preventDefault();
        
        var lista = document.getElementById("lista");
        if (lista !== null) {
            lista.parentNode.removeChild(lista);
        }
        
        var poly = document.getElementById('poly').value;

        var xhr = new XMLHttpRequest();
        xhr.open('POST', encodeURI(base_url + '/hamming'));
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        xhr.onload = function () {
            if (xhr.status !== 200) {
                console.log('request failed.');
            } else {
                h = '<ul id="lista">'
                texts = xhr.responseText.split('\n')
                for (var index = 0; index < texts.length-1; index++) {
                    h += '<li>' + texts[index] + '</li>'
                }
                h += '</ul>'
                document.getElementById('resultados').insertAdjacentHTML('afterbegin', h);
            }
        }
        xhr.send(encodeURI('poly=' + poly));
    }
    
    button.addEventListener('click', showAnalisis, 'false');
}