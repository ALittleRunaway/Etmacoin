let userId = 0
let userLogin = ""
let userWallet = ""

function getUserTransactionsInfo() {
    const urlParams = new URLSearchParams(window.location.search);
    userId = urlParams.get('user_id');

    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/get_user_info?user_id=" + userId, false);
    xmlHttp.send()
    userLogin = JSON.parse(xmlHttp.responseText)["Login"]
    userWallet = JSON.parse(xmlHttp.responseText)["Wallet"]
    let userBalance = JSON.parse(xmlHttp.responseText)["Balance"]
    event.preventDefault();

    xmlHttp.open("GET", "/user_transactions?user_id=" + userId, false);
    xmlHttp.send()

    xmlHttp.open("GET", "/latest_transactions", false);
    xmlHttp.send()

    // alert(userId + " " + userLogin + " " + userWallet)
}
