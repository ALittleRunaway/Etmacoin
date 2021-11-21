let userId = 0

function copyToClipboard() {
    var copyText = document.getElementById("wallet");
    copyText.select();
    copyText.setSelectionRange(0, 99999);
    navigator.clipboard.writeText(copyText.value);
    alert("The wallet id is copied!");
}

function getUserInfo() {
    const urlParams = new URLSearchParams(window.location.search);
    userId = urlParams.get('user_id');

    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/get_user_info?user_id=" + userId, false);
    xmlHttp.send()
    let userLogin = JSON.parse(xmlHttp.responseText)["Login"]
    let userWallet = JSON.parse(xmlHttp.responseText)["Wallet"]
    let userBalance = JSON.parse(xmlHttp.responseText)["Balance"]
    event.preventDefault();

    let balance = document.getElementById("balance")
    balance.value = userBalance + " $⩢⪑⫚"
    let wallet = document.getElementById("wallet")
    wallet.value = userWallet
}

function validateTransactionData(recipient, amount) {
    if (recipient === "") {
        return "The recipient wallet has to be specified!"
    }
    if (amount === "") {
        return "The amount of the transaction has to be specified!"
    }
    let balance = document.getElementById("balance").value.split(' ')[0]
    if (parseInt(amount) > parseInt(balance)) {
        return "There is not enough EtmaCoin on your balance."
    }
    return ""
}

function sendNewTransaction() {
    let recipient_wallet = document.getElementById("recipient")
    let amount = document.getElementById("amount")
    let message = document.getElementById("message")

    let error = validateTransactionData(recipient_wallet.value, amount.value)

    if (error === "") {
        let url = "/new_transaction?user_id=" + userId + "&recipient_wallet=" + recipient_wallet.value +
            "&amount=" + amount.value + "&message=" + message.value

        let xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", url, false);
        xmlHttp.send()

        let transactionSendSummary = JSON.parse(xmlHttp.responseText)["Response"]

        alert(transactionSendSummary)
        recipient_wallet.value = ""
        amount.value = ""
        message.value = ""
    } else {
        alert(error)
    }


}