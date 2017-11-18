
var API = {
    base: "http://appengine.emoji-compress.com/",
    // base: "http://127.0.0.1:8080/",
    _run: function (payload, handler, callback) {
        var targetUri = new URL(this.base + handler);
        var webRequest = new Request(targetUri,
            { method: "POST", mode: "cors", body: JSON.stringify(payload) });

        //https://developer.mozilla.org/en-US/docs/Web/API/Request/Request
        fetch(webRequest)
            .then(function (response) {
                if (response.status == 200) return response.json();
                else throw new Error('Something went wrong on api server!');
            })
            .then(function (responseJson) {
                if (responseJson.ok == false) {
                    throw new Error(responseJson.error)
                }
                callback(responseJson.response, null)
            })
            .catch(function (reason) {
                console.error(reason)
                callback(null, reason)
            })
    },
    bytesmap: {
        encode: function (original, callback) {
            var payload = { text: original }
            var handler = "bytesmap/encode"
            API._run(payload, handler, callback)
        },
        decode: function (original, callback) {
            var payload = { text: original }
            var handler = "bytesmap/decode"
            API._run(payload, handler, callback)
        }
    }
}

// API.bytesmap.encode("a", (result, error) => { console.info(result); console.error(error) })

document.addEventListener('DOMContentLoaded', function () {

    var watch = function (handler, func) {
        var idprefix = handler + "_" + func + "_"
        var button = document.getElementById(idprefix + "btn")
        var textIn = document.getElementById(idprefix + "in")
        var textOut = document.getElementById(idprefix + "out")

        button.addEventListener("click", function () {
            API[handler][func](textIn.value, function (result, error) {
                if (error) {
                    textOut.value = "Error: " + error
                    return
                }
                textOut.value = result
            })
        })
    }

    watch("bytesmap", "encode")
    watch("bytesmap", "decode")

}, false);

