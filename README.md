# sky-fiber-jsonrpc
json rpc 2.0 interface for sky-fiber

# install

go get osamingo/jsonrpc

go get github.com/osamingo/jsonrpc

# args

flag.StringVar(&NodeRpcAddress, "backend", "http://127.0.0.1:16420", "backend server web interface addr")

flag.StringVar(&ListenAddr, "port", "127.0.0.1:8081", "listen port")

# version

```
curl -X POST -H "Content-Type:application/json" -d '{"id":"12", "method":"version", "jsonrpc":"2.0"}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"version":"0.23.0","commit":"c347c4df349007e1f4a4db747d95d4d66e7c80a7","branch":""}}

```

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

## outputs by addresses

```
curl -X POST -H "Content-Type:application/json" -d '{"id":"12", "method":"outputs", "params":{"addrs":["EX8omhDyjKtc8zHGp1KZwn7usCndaoJxSe"]}, "jsonrpc":"2.0"}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"head_outputs":[{"hash":"e4df6feadd5e47f92815da8689fc200189ade33149992cd2bcd19ef1bcc553da","time":1527992207,"block_seq":1267,"src_tx":"295276e44d3b3cdc29fb52a8507b59d2aaa50d2658f524e9e8d85c9c72ee9968","address":"EX8omhDyjKtc8zHGp1KZwn7usCndaoJxSe","coins":"94.000000","hours":13516,"calculated_hours":19124}],"outgoing_outputs":[],"incoming_outputs":[]}}
```

## outputs by hashes 

```
curl -X POST -H "Content-Type:application/json" -d '{"id":"12", "method":"outputs", "params":{"hashes":["e4df6feadd5e47f92815da8689fc200189ade33149992cd2bcd19ef1bcc553da"]}, "jsonrpc":"2.0"}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"head_outputs":[{"hash":"e4df6feadd5e47f92815da8689fc200189ade33149992cd2bcd19ef1bcc553da","time":1527992207,"block_seq":1267,"src_tx":"295276e44d3b3cdc29fb52a8507b59d2aaa50d2658f524e9e8d85c9c72ee9968","address":"EX8omhDyjKtc8zHGp1KZwn7usCndaoJxSe","coins":"94.000000","hours":13516,"calculated_hours":19124}],"outgoing_outputs":[],"incoming_outputs":[]}}
```

## wallet balance

```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"walletBalance", "params":{"id":"samos_cli.wlt"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"confirmed":{"coins":94000000,"hours":19124},"predicted":{"coins":94000000,"hours":19124}}}
```

## balance

```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"balance", "params":{"addrs":["EX8omhDyjKtc8zHGp1KZwn7usCndaoJxSe"]}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"confirmed":{"coins":94000000,"hours":19124},"predicted":{"coins":94000000,"hours":19124}}}
```

## transaction

```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"transaction", "params":{"txid":"6532f4cad6ac0fad68967cb55c00084a3c202538387a0a7c60e11cf080870cf8"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"status":{"confirmed":true,"unconfirmed":false,"height":303,"block_seq":1000,"unknown":false},"time":0,"txn":{"length":220,"type":0,"txid":"6532f4cad6ac0fad68967cb55c00084a3c202538387a0a7c60e11cf080870cf8","inner_hash":"1c4ca2da02ae0243f745c54b8e277f9d6a9f97b6a82b48fe7665aca4948402cb","timestamp":1526806119,"sigs":["34faf1733e532c02c6bf36ba6d45b40e0bea9fecdbcd4b4121319e5a8cfed83657347a7c854cec50cca8174bc06c1190dfb1336208698927f7c1eba540c9c6b901"],"inputs":["8b5b4375f1a87ecca5d7a578fcc81f31c89748d999518c551431c0fe4d461f8c"],"outputs":[{"uxid":"52edc78d1b98e2c02ba019ea33904bb6f1fb724b87fe10e2bfc847d0944f24f4","dst":"yRqTGJRSeKwVrMyPBhzmVVLMdmPKcY5RuP","coins":"410.000000","hours":410},{"uxid":"65760f651bbaeb4276cd7d90e850731d20ba01789fd33717d6b1e0a9d9af3477","dst":"2gauuARNtVu3Rqqmh7qMp7FL9GGWLxkDVyX","coins":"915445.000000","hours":84089}]}}}
```

## address new
```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"addressNew", "params":{"id":"t1.wlt", "num":"1"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"addresses":["gjgjBjxNKAZsyo8w9sPpyxLz7ikCCiw771"]}}

```

## wallet create
```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"walletCreate", "params":{ "seed":"aaf dfd esxc 2ddf", "label":"t2wlt", "scan":"10"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"meta":{"coin":"skycoin","filename":"2018_06_01_82e1.wlt","label":"t2wlt","type":"deterministic","version":"0.1","crypto_type":"","timestamp":0,"encrypted":false},"entries":[{"address":"NUTy4CnirPi61BoSKubc2dQnXBkgAGVeN7","public_key":"02aa9f66f460a8e05ee431b96ea970defdcbc0bb3af4a1b579b8750f1b1d5224fe"}]}}

```

## wallet spend

```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"walletSpend", "params":{ "id":"t1.wlt", "dst":"NUTy4CnirPi61BoSKubc2dQnXBkgAGVeN7", "coins":"10000000"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"balance":{"confirmed":{"coins":20030000000,"hours":4721316},"predicted":{"coins":20020000000,"hours":3209871}},"txn":{"length":220,"type":0,"txid":"47cd3540d516470ad33f23549caaa10f25a1d0d0ae68914430c62511ec917489","inner_hash":"2343b16b65faa0207d5ac6aaa73a606694248c6a768a845f884421f6212d0baa","sigs":["186443c3316240a2969f7c33e3985fb7ee9d22555557184e293dd0a0d34b4004122c3cc172f9287f1f760b2e81117478f314b6f60b52600b4dc76a4a4265627e01"],"inputs":["a0adff987e0eb28979c67744d93188e19f60ee94e6e495b50796e7e758506053"],"outputs":[{"uxid":"330956610431962f586d3224c9c5948a46da4e9cc11f64ebef8cd1b71cfc1541","dst":"2fxav8p7QFkKk8TBwmE6wvu8S8VVEyvpX8C","coins":"9980.000000","hours":503815},{"uxid":"c26f6011b0b1c31a7309e12e03288ecb37b695d259caa7e78a340cc190148371","dst":"NUTy4CnirPi61BoSKubc2dQnXBkgAGVeN7","coins":"10.000000","hours":503815}]}}}
```

## create transaction

```
curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"transactionCreate", "params":{ "hours_selection": { "type": "auto", "mode": "share", "share_factor": "0.8" }, "wallet": { "id": "fortest1.wlt" }, "change_address": "24BPafbuQH9w73LMPVDYiaBfq8yzyuS2pPq", "to": [{ "address": "upqWLUuHgZJwEc661Ss7wcycdY6ruoGsKn", "coins": "4" }, { "address": "25owdr2fcHSmjyJnWUNieCJB5nZQQYxNGpP", "coins": "2" }] }}'  http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":{"transaction":{"length":257,"type":0,"txid":"37e28633faf3dae8ef82571cbaf4a7e9225bad206158f04b73720b69c87e3793","inner_hash":"59932086bef30e65ea7becbcbc09ff5f92891627bb1d79491251bcdcc2acf81b","fee":"58","sigs":["7b4aab41f35c0941d4eccb06ac6079fd025abb0099d11aeb12315f58e982fbac65a9b4a04c8ba8fec48d5983393e99c324f2a7cb51a9f6dd3cd71af05ab3ce1100"],"inputs":[{"uxid":"b6e46ed9a65b1fd7bf66b486f3d509a3ec0f118072bc17ac9061fe6cebadbea7","address":"24BPafbuQH9w73LMPVDYiaBfq8yzyuS2pPq","coins":"13.010000","hours":"0","calculated_hours":"116","timestamp":1528174897,"block":1286,"txid":"7f62b04ce7b3f5097938bfa181dbc8e5aaae09428dde90e8ebc77ac67ccf0e32"}],"outputs":[{"uxid":"921d3f76d52b848fddb215d82ac04964991cf8062b46c1fa8d3db22124b06f35","address":"upqWLUuHgZJwEc661Ss7wcycdY6ruoGsKn","coins":"4.000000","hours":"31"},{"uxid":"e9405981ae1afbf3491bff3c2735f8afc33010d0340c9e5a076099a747dc5b2f","address":"25owdr2fcHSmjyJnWUNieCJB5nZQQYxNGpP","coins":"2.000000","hours":"15"},{"uxid":"00a03eb2135f9a2a109b642d560049249c3c15d9b1294e6c5d59590e3a167049","address":"24BPafbuQH9w73LMPVDYiaBfq8yzyuS2pPq","coins":"7.010000","hours":"12"}]},"encoded_transaction":"010100000059932086bef30e65ea7becbcbc09ff5f92891627bb1d79491251bcdcc2acf81b010000007b4aab41f35c0941d4eccb06ac6079fd025abb0099d11aeb12315f58e982fbac65a9b4a04c8ba8fec48d5983393e99c324f2a7cb51a9f6dd3cd71af05ab3ce110001000000b6e46ed9a65b1fd7bf66b486f3d509a3ec0f118072bc17ac9061fe6cebadbea703000000008347f5baf0c640551e8251c516cf973ae5f888a000093d00000000001f00000000000000009c184f52c91c7eee4d8f0101698d2e84e9a0053180841e00000000000f0000000000000000980b2388daaab0f742f5b76392aa33680366113dd0f66a00000000000c00000000000000"}}

```

## inject transaction

```

curl -X POST -H "Content-Type:application/json" -d '{"jsonrpc":"2.0", "id":"12", "method":"transactionInject", "params":{"rawtx":"010100000059932086bef30e65ea7becbcbc09ff5f92891627bb1d79491251bcdcc2acf81b010000007b4aab41f35c0941d4eccb06ac6079fd025abb0099d11aeb12315f58e982fbac65a9b4a04c8ba8fec48d5983393e99c324f2a7cb51a9f6dd3cd71af05ab3ce110001000000b6e46ed9a65b1fd7bf66b486f3d509a3ec0f118072bc17ac9061fe6cebadbea703000000008347f5baf0c640551e8251c516cf973ae5f888a000093d00000000001f00000000000000009c184f52c91c7eee4d8f0101698d2e84e9a0053180841e00000000000f0000000000000000980b2388daaab0f742f5b76392aa33680366113dd0f66a00000000000c00000000000000"}}' http://127.0.0.1:8081/jrpc
{"id":"12","jsonrpc":"2.0","result":"37e28633faf3dae8ef82571cbaf4a7e9225bad206158f04b73720b69c87e3793"}

```
