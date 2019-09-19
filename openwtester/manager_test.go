package openwtester

import (
	"github.com/blocktree/openwallet/log"
	"github.com/blocktree/openwallet/openw"
	"github.com/blocktree/openwallet/openwallet"
	"path/filepath"
	"testing"
)

var (
	testApp        = "bitcoin-adapter"
	configFilePath = filepath.Join("conf")
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"DASH",
	}
	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "HELLO DASH", IsTrust: true, Password: "12345678"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()
	walletID := "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh"
	//walletID := "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh"
	account := &openwallet.AssetsAccount{Alias: "mainnetBTC", WalletID: walletID, Required: 1, Symbol: "DASH", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "12345678", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	//walletID := "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh"
	//accountID := "FqQBQ8Bn26GogR7UAu6e2ZVhrYYmKUpmBS7CSM1KLTTZ"
	walletID := "WGe5KPhw5RoikvCxgEk76emZ8PNtoPPBgh"
	accountID := "33LEEmDcZG18TcY6qBhMMyrZZ5hLi2fnwmQ3TdjxfBrR"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 4)
	if err != nil {
		log.Error(err)
		return
	}

	for i, w := range address {
		log.Infof("address[%d] : %s", i, w.Address)
	}

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	//walletID := "WAmTnvPKMWpJBqKk6cncFG3mTXz3iPmtzV"
	//accountID := "21Vn4NEmXT6DRy2EfdPTAJCS2kYTACTuconBer8AQ1cz"
	walletID := "WAmTnvPKMWpJBqKk6cncFG3mTXz3iPmtzV"
	accountID := "EPxkNBu6iMospC6aHQppv36UGY4mb1WqUE7oNZ7Xp9Df"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Infof("address[%d] : %s", i, w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
