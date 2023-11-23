'use strict';

class Webflow {
    constructor(selector) {
        this.el = document.querySelector(selector);
        this.golang = new Go();
    }

    init() {
        WebAssembly.instantiateStreaming(fetch("webflow.wasm"), this.golang.importObject).then((result) => {
            this.golang.run(result.instance);
        });
    }

    render() {

    }
}
