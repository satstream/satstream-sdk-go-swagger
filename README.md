# Go API client for satstream_go_sdk

Satstream API

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.17
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://satstream.io](https://satstream.io)

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./satstream_go_sdk"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.satstream.io/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressesApi* | [**GetAddressBalance**](docs/AddressesApi.md#getaddressbalance) | **Get** /addresses/{address}/balance | Get address balance
*AddressesApi* | [**GetAddressNonInscriptionUtxos**](docs/AddressesApi.md#getaddressnoninscriptionutxos) | **Get** /addresses/{address}/utxos | Get address non-inscription UTXOs
*AddressesApi* | [**GetAddressRuneBalance**](docs/AddressesApi.md#getaddressrunebalance) | **Get** /addresses/{address}/runes/{runeid} | Get address rune balance
*AddressesApi* | [**GetAddressRunesBalanceList**](docs/AddressesApi.md#getaddressrunesbalancelist) | **Get** /addresses/{address}/runes | Get address runes balance list
*AddressesApi* | [**GetAddressTimeframeBalance**](docs/AddressesApi.md#getaddresstimeframebalance) | **Get** /addresses/{address}/balance/timeframe | Get address timeframe balance
*BlocksApi* | [**GetBlockByHash**](docs/BlocksApi.md#getblockbyhash) | **Get** /blocks/hash/{hash} | Get block by hash
*BlocksApi* | [**GetBlockInfo**](docs/BlocksApi.md#getblockinfo) | **Get** /blocks/{height} | Get block info
*BlocksApi* | [**GetBlockTransactions**](docs/BlocksApi.md#getblocktransactions) | **Get** /blocks/{height}/transactions | Get block transactions
*BlocksApi* | [**GetCurrentBlockHeight**](docs/BlocksApi.md#getcurrentblockheight) | **Get** /blocks/current-height | Get current block height
*FeesApi* | [**GetRecommendedFees**](docs/FeesApi.md#getrecommendedfees) | **Get** /fees | Get recommended fees
*MempoolApi* | [**GetAddressMempoolTransactions**](docs/MempoolApi.md#getaddressmempooltransactions) | **Get** /mempool/addresses/{address}/transactions | Get address mempool transactions
*MempoolApi* | [**GetMempoolTransactionInfo**](docs/MempoolApi.md#getmempooltransactioninfo) | **Get** /mempool/transactions/{txid} | Get mempool transaction info
*MempoolApi* | [**GetMempoolTransactions**](docs/MempoolApi.md#getmempooltransactions) | **Get** /mempool/transactions | Get mempool transactions
*RunesApi* | [**GetRunesHolders**](docs/RunesApi.md#getrunesholders) | **Get** /runes/{runeId}/holders | Get rune holders
*RunesApi* | [**GetRunesInfo**](docs/RunesApi.md#getrunesinfo) | **Get** /runes/{runeId} | Get rune info
*RunesApi* | [**GetRunesInfoList**](docs/RunesApi.md#getrunesinfolist) | **Get** /runes | Get runes info list
*TransactionsApi* | [**BroadcastTransaction**](docs/TransactionsApi.md#broadcasttransaction) | **Post** /transactions/broadcast | Broadcast transaction
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /indexer/tx/{hash} | Get transaction
*TransactionsApi* | [**GetTransactionInfo**](docs/TransactionsApi.md#gettransactioninfo) | **Get** /transactions/{txid} | Get transaction info

## Documentation For Models

 - [AllOfresponsesGetAddressTimeframeBalanceItemBlockRange](docs/AllOfresponsesGetAddressTimeframeBalanceItemBlockRange.md)
 - [BigInt](docs/BigInt.md)
 - [GithubComSatstreamSsApiServerApiAddressesResponsesError](docs/GithubComSatstreamSsApiServerApiAddressesResponsesError.md)
 - [GithubComSatstreamSsApiServerApiBlocksResponsesError](docs/GithubComSatstreamSsApiServerApiBlocksResponsesError.md)
 - [GithubComSatstreamSsApiServerApiRunesResponsesError](docs/GithubComSatstreamSsApiServerApiRunesResponsesError.md)
 - [GithubComSatstreamSsApiServerApiTransactionsResponsesError](docs/GithubComSatstreamSsApiServerApiTransactionsResponsesError.md)
 - [GithubComSatstreamSsUtilsOrdinalsTerms](docs/GithubComSatstreamSsUtilsOrdinalsTerms.md)
 - [GithubComSatstreamSsUtilsOrdinalsTermsRange](docs/GithubComSatstreamSsUtilsOrdinalsTermsRange.md)
 - [GithubComSatstreamSsUtilsRpcBlock](docs/GithubComSatstreamSsUtilsRpcBlock.md)
 - [GithubComSatstreamSsUtilsRpcBtcTx](docs/GithubComSatstreamSsUtilsRpcBtcTx.md)
 - [GithubComSatstreamSsUtilsRpcPrevOut](docs/GithubComSatstreamSsUtilsRpcPrevOut.md)
 - [GithubComSatstreamSsUtilsRpcScriptPubKey](docs/GithubComSatstreamSsUtilsRpcScriptPubKey.md)
 - [GithubComSatstreamSsUtilsRpcScriptSig](docs/GithubComSatstreamSsUtilsRpcScriptSig.md)
 - [GithubComSatstreamSsUtilsRpcUtxoRune](docs/GithubComSatstreamSsUtilsRpcUtxoRune.md)
 - [GithubComSatstreamSsUtilsRpcVin](docs/GithubComSatstreamSsUtilsRpcVin.md)
 - [GithubComSatstreamSsUtilsRpcVout](docs/GithubComSatstreamSsUtilsRpcVout.md)
 - [GithubComSatstreamSsUtilsStoreTransactionDocument](docs/GithubComSatstreamSsUtilsStoreTransactionDocument.md)
 - [ResponsesBlockRange](docs/ResponsesBlockRange.md)
 - [ResponsesGetAddressBalance](docs/ResponsesGetAddressBalance.md)
 - [ResponsesGetAddressBalanceData](docs/ResponsesGetAddressBalanceData.md)
 - [ResponsesGetAddressMempoolTxs](docs/ResponsesGetAddressMempoolTxs.md)
 - [ResponsesGetAddressNonInscriptionUtxo](docs/ResponsesGetAddressNonInscriptionUtxo.md)
 - [ResponsesGetAddressNonInscriptionUtxoData](docs/ResponsesGetAddressNonInscriptionUtxoData.md)
 - [ResponsesGetAddressRuneBalance](docs/ResponsesGetAddressRuneBalance.md)
 - [ResponsesGetAddressRuneBalanceData](docs/ResponsesGetAddressRuneBalanceData.md)
 - [ResponsesGetAddressRunesBalanceList](docs/ResponsesGetAddressRunesBalanceList.md)
 - [ResponsesGetAddressRunesBalanceListData](docs/ResponsesGetAddressRunesBalanceListData.md)
 - [ResponsesGetAddressRunesBalanceListItem](docs/ResponsesGetAddressRunesBalanceListItem.md)
 - [ResponsesGetAddressTimeframeBalance](docs/ResponsesGetAddressTimeframeBalance.md)
 - [ResponsesGetAddressTimeframeBalanceData](docs/ResponsesGetAddressTimeframeBalanceData.md)
 - [ResponsesGetAddressTimeframeBalanceItem](docs/ResponsesGetAddressTimeframeBalanceItem.md)
 - [ResponsesGetBlockByHash](docs/ResponsesGetBlockByHash.md)
 - [ResponsesGetBlockHeight](docs/ResponsesGetBlockHeight.md)
 - [ResponsesGetBlockHeightData](docs/ResponsesGetBlockHeightData.md)
 - [ResponsesGetBlockInfo](docs/ResponsesGetBlockInfo.md)
 - [ResponsesGetBlockTransactions](docs/ResponsesGetBlockTransactions.md)
 - [ResponsesGetFees](docs/ResponsesGetFees.md)
 - [ResponsesGetFeesData](docs/ResponsesGetFeesData.md)
 - [ResponsesGetMempoolTransactions](docs/ResponsesGetMempoolTransactions.md)
 - [ResponsesGetMempoolTxInfo](docs/ResponsesGetMempoolTxInfo.md)
 - [ResponsesGetRuneHolders](docs/ResponsesGetRuneHolders.md)
 - [ResponsesGetRuneHoldersData](docs/ResponsesGetRuneHoldersData.md)
 - [ResponsesGetRuneInfo](docs/ResponsesGetRuneInfo.md)
 - [ResponsesGetRunesInfoList](docs/ResponsesGetRunesInfoList.md)
 - [ResponsesGetRunesInfoListData](docs/ResponsesGetRunesInfoListData.md)
 - [ResponsesGetTransaction](docs/ResponsesGetTransaction.md)
 - [ResponsesGetTxInfo](docs/ResponsesGetTxInfo.md)
 - [ResponsesGetTxInfoData](docs/ResponsesGetTxInfoData.md)
 - [ResponsesNonInscriptionUtxo](docs/ResponsesNonInscriptionUtxo.md)
 - [ResponsesPaginationInfo](docs/ResponsesPaginationInfo.md)
 - [ResponsesRuneHolder](docs/ResponsesRuneHolder.md)
 - [ResponsesRuneInfo](docs/ResponsesRuneInfo.md)
 - [ResponsesSendRawTransaction](docs/ResponsesSendRawTransaction.md)
 - [ResponsesSendRawTransactionData](docs/ResponsesSendRawTransactionData.md)

## Documentation For Authorization

## ApiKeyAuth
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

team@satstream.io
