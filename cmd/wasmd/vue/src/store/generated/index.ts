// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import CosmWasmWasmdCosmwasmWasmV1 from './CosmWasm/wasmd/cosmwasm.wasm.v1'


export default { 
  CosmWasmWasmdCosmwasmWasmV1: load(CosmWasmWasmdCosmwasmWasmV1, 'cosmwasm.wasm.v1'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}
