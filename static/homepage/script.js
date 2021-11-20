function copyToClipboard() {
    var copyText = document.getElementById("wallet");
    copyText.select();
    copyText.setSelectionRange(0, 99999);
    navigator.clipboard.writeText(copyText.value);
    alert("The wallet id is copied!");
}

function getUserInfo() {
    const urlParams = new URLSearchParams(window.location.search);
    let userId = urlParams.get('user_id');

    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "http://localhost:6006/get_user_info?user_id=" + userId, false);
    xmlHttp.send()
    userLogin = JSON.parse(xmlHttp.responseText)["Login"]
    userWallet = JSON.parse(xmlHttp.responseText)["Wallet"]
    userBalance = JSON.parse(xmlHttp.responseText)["Balance"]
    event.preventDefault();

    let balance = document.getElementById("balance")
    balance.value = userBalance + " $"
    let wallet = document.getElementById("wallet")
    wallet.value = userWallet
}