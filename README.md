# sky-fiber-jsonrpc
json rpc 2.0 interface for sky-fiber

# install

go get osamingo/jsonrpc

go get github.com/osamingo/jsonrpc

# args

flag.StringVar(&NodeRpcAddress, "backend", "http://127.0.0.1:16420", "backend server web interface addr")

flag.StringVar(&ListenAddr, "port", "127.0.0.1:8081", "listen port")

# block

```sh
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"block", "params":{"seq":24}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"header":{"seq":24,"block_hash":"79124ce0a59cdf54457a6aa84c2123842f043174ae1e687dd6708515464d07c1","previous_block_hash":"e2da61df39cd248a20193887328369196e4b98ae1e749c2c5f08ac52ef464181","timestamp":1524273699,"fee":38245020,"version":0,"tx_body_hash":"62c8cbea2e671457784d6a93b3984c98727f53b9d2ac4ae5b765e75a7d18c6a5"},"body":{"txns":[{"length":294,"type":0,"txid":"62c8cbea2e671457784d6a93b3984c98727f53b9d2ac4ae5b765e75a7d18c6a5","inner_hash":"b7a60d4bfaf1c9924a262d5475920db82d5bf5c22c0c76ebb019e029614d66a9","sigs":["b06779364ddc38b6ad1a65044bc1b67dbe57117647d8b4e34b5d559965cb1e362895e08aa312c76246fc7e79c720f26b77eec5155df3888e5bf4924fde69d46501"],"inputs":["cedd3f18048df4274e3ad8bb2da0c5e8a1bdd4006ce6fe4fcb8503b0592ea1de"],"outputs":[{"uxid":"393a4464eabb2ecf2ed5d826e6f77788d932362639ad02051e5f47809b85049e","dst":"p88NvTpFFZUxcgffrcirrJgJikRtqZP7Cd","coins":"36.000000","hours":36},{"uxid":"91f5be2f0e04afbcd1042a7032fcc3055a593b75e6a7e6d65c5265d1500052ad","dst":"q2SZxPJoJLbnnN5NkrSJ3WvJ7rRbHXtzDu","coins":"36.000000","hours":36},{"uxid":"cb8cc1d929a9638b7f2c19d9485e4c210e6c6aa26425bf264a8670e8053bd283","dst":"rPw4CJ4WCoYs6FguGVLfeu77CcG2fXidrF","coins":"36.000000","hours":36},{"uxid":"5f91a3ab4d891b04e3fa6dd6997b769cc63eb5f44f4c9941f3ea1236fc27bee0","dst":"b9JKtor5PyDJTjSogY7rqeDhMqtwAwjXuq","coins":"999675.000000","hours":38244911}]}]},"size":294}}

```

# last blocks

```sh
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"blockLastN", "params":{"num":1}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"blocks":[{"header":{"seq":1302,"block_hash":"563c599dfe754dd43abf8dac6cfbf07587014bdae55d166f40838fdfc2ec22cf","previous_block_hash":"aea90d18c0e019a1c63249cd38106f756dbf24f215453f7a6b8d33352dc33478","timestamp":1528207017,"fee":963,"version":0,"tx_body_hash":"f9300c7eec727fe443e95d01f53ab4df6dbeffadd85985287e5b6f86d5ff35ec"},"body":{"txns":[{"length":220,"type":0,"txid":"f9300c7eec727fe443e95d01f53ab4df6dbeffadd85985287e5b6f86d5ff35ec","inner_hash":"08549c223e915405821de1b8c32aa44acf2fba5abbd628cee50234b4510eb743","sigs":["221f6d69f602ca00b8db7d303dc98191dc37d4cb5fc77d8fa13da38b5dc0a285696ae74a633534135e44385b08307f1ce5b74b7f1a1359e6a4254b96c938da2100"],"inputs":["057ab7d60a5275ed883ee6d3340d2f6b9adeb0d51d5311aff32fb9089b1b9f73"],"outputs":[{"uxid":"e5da661011d7dcbc32e5c25e233f53ac82ee458cd223d9ffe51dbdeaa30a469c","dst":"Mr3XDK9M2KfNjAcgAgw5sQ8ppRurCqXLyD","coins":"8.000000","hours":482},{"uxid":"26fb0e657f32703ed02df2f8d9c303a5e18c675071d74aa07b06ba23ed4d2cbb","dst":"qjjndKgpJMkWND4rmy1KH47sg4FvJf6juQ","coins":"2.000000","hours":481}]}]},"size":220}]}}

```

## block range

```sh
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"blockRange", "params":{"start":1000, "end":1000}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"blocks":[{"header":{"seq":1000,"block_hash":"2d19b70c4cdf8f62fd784cccfa42ec7c5b65cde27e048e43e451552bc2841248","previous_block_hash":"0a87dbb7296624dce79d6dfbd5159a7a9e6e258155d79639c4e69a22ee5b878d","timestamp":1526806119,"fee":84500,"version":0,"tx_body_hash":"6532f4cad6ac0fad68967cb55c00084a3c202538387a0a7c60e11cf080870cf8"},"body":{"txns":[{"length":220,"type":0,"txid":"6532f4cad6ac0fad68967cb55c00084a3c202538387a0a7c60e11cf080870cf8","inner_hash":"1c4ca2da02ae0243f745c54b8e277f9d6a9f97b6a82b48fe7665aca4948402cb","sigs":["34faf1733e532c02c6bf36ba6d45b40e0bea9fecdbcd4b4121319e5a8cfed83657347a7c854cec50cca8174bc06c1190dfb1336208698927f7c1eba540c9c6b901"],"inputs":["8b5b4375f1a87ecca5d7a578fcc81f31c89748d999518c551431c0fe4d461f8c"],"outputs":[{"uxid":"52edc78d1b98e2c02ba019ea33904bb6f1fb724b87fe10e2bfc847d0944f24f4","dst":"yRqTGJRSeKwVrMyPBhzmVVLMdmPKcY5RuP","coins":"410.000000","hours":410},{"uxid":"65760f651bbaeb4276cd7d90e850731d20ba01789fd33717d6b1e0a9d9af3477","dst":"2gauuARNtVu3Rqqmh7qMp7FL9GGWLxkDVyX","coins":"915445.000000","hours":84089}]}]},"size":220}]}}

```
