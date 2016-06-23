window.onload = function () {
    button = document.getElementById("boton");
    var base_url = window.location.origin;

    function showCrc(e) {
        e.preventDefault();
        
        var respuesta = document.getElementById("respuesta");
        if (respuesta !== null) {
            respuesta.parentNode.removeChild(respuesta);
        }
        
        var msg = document.getElementById('msg').value;
        var poly = document.getElementById('poly').value;

        var xhr = new XMLHttpRequest();
        xhr.open('POST', encodeURI(base_url + '/crc'));
        xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
        xhr.onload = function () {
            if (xhr.status !== 200) {
                console.log('request failed.');
            } else {
                document.getElementById('resultados').insertAdjacentHTML('afterbegin', '<p id="respuesta">' + xhr.responseText + '</p>');
            }
        }
        xhr.send(encodeURI('msg=' + msg + '&poly=' + poly));
    }

    button.addEventListener('click', showCrc, 'false');
}