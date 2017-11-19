
var API = {
    base: "https://appengine.emoji-compress.com/",
    // base: "http://127.0.0.1:8080/",
    call: function (payload, handler, callback) {
        var targetUri = new URL(this.base + handler);
        //https://developer.mozilla.org/en-US/docs/Web/API/Request/Request
        var webRequest = new Request(targetUri,
            { method: "POST", mode: "cors", body: JSON.stringify(payload) });

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
}

// API.bytesmap.encode("a", (result, error) => { console.info(result); console.error(error) })

document.addEventListener('DOMContentLoaded', function () {

    var watchSimple = function (handler, method) {
        var idprefix = handler + "_" + method + "_"
        var button = document.getElementById(idprefix + "btn")
        var textIn = document.getElementById(idprefix + "in")
        var textOut = document.getElementById(idprefix + "out")
        var url = handler + "/" + method

        var dictIn = document.getElementById(idprefix + "dictionary_in")
        var dictOut = document.getElementById(idprefix + "dictionary_out")

        button.addEventListener("click", function () {

            if (dictOut) {
                //we are expecting a dictionary out so ...
                textOut.value = "Loading ❂ ❂ ❂ "
                dictOut.value = "Loading ❂ ❂ ❂ "
            } else {
                textOut.value = "Loading ❂ ❂ ❂ "
            }

            API.call({
                text: textIn.value,
                dict: (dictIn ? JSON.parse(dictIn.value) : null)
            }, url, function (apiResponse, error) {
                if (error) {
                    textOut.value = "Error: " + error
                    return
                }
                // console.info(apiResponse)

                if (dictOut) {
                    //we are expecting a dictionary out so ...
                    textOut.value = apiResponse.archive
                    dictOut.value = JSON.stringify(apiResponse.dict)
                } else {
                    textOut.value = apiResponse
                }
            })
        })
    }

    watchSimple("bytesmap", "encode")
    watchSimple("bytesmap", "decode")

    watchSimple("lz78", "encode")
    watchSimple("lz78", "decode")

    watchSimple("dictionary", "encode")
    watchSimple("dictionary", "decode")

}, false);

