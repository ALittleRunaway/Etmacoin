let userId = 0
let userLogin = ""
let userWallet = ""

async function sha256(message) {
    let msgBuffer = new TextEncoder().encode(message);
    let hashBuffer = await crypto.subtle.digest('SHA-256', msgBuffer);
    const hashArray = Array.from(new Uint8Array(hashBuffer));
    const hashHex = hashArray.map(b => b.toString(16).padStart(2, '0')).join('');
    return hashHex;
}

function copyToClipboard() {
    var copyText = document.getElementById("wallet");
    copyText.select();
    copyText.setSelectionRange(0, 99999);
    navigator.clipboard.writeText(copyText.value);
    alert("The wallet id is copied!");
}

function getGreetingPhrase() {
    let time = new Date().getHours();
    if ((time >= 5) && (time < 12)) {
        return "Good morning, "
    } else if ((time >= 12) && (time < 17)) {
        return "Good afternoon, "
    } else if ((time >= 17) && (time < 23)) {
        return "Good evening, "
    } else {
        return "Good night, "
    }
}

function getUserInfo() {
    const urlParams = new URLSearchParams(window.location.search);
    userId = urlParams.get('user_id');

    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/get_user_info?user_id=" + userId, false);
    xmlHttp.send()
    userLogin = JSON.parse(xmlHttp.responseText)["Login"]
    userWallet = JSON.parse(xmlHttp.responseText)["Wallet"]
    let userBalance = JSON.parse(xmlHttp.responseText)["Balance"]
    event.preventDefault();


    let login = document.getElementById("user_login")
    login.innerHTML = getGreetingPhrase() + userLogin + "!"
    let balance = document.getElementById("balance")
    balance.value = userBalance + " Ɇ"
    let wallet = document.getElementById("wallet")
    wallet.value = userWallet
}

function validateTransactionData(recipient, amount) {
    if (recipient === "") {
        return "The recipient wallet has to be specified!"
    }
    if (recipient === userWallet) {
        userWallet.value = ""
        return "You can't send the transaction to yourself!"
    }
    if (amount === "") {
        return "The amount of the transaction has to be specified!"
    }
    let balance = document.getElementById("balance").value.split(' ')[0]
    if (parseInt(amount) > parseInt(balance)) {
        return "There is not enough ɆtmaCoin on your balance."
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

        xmlHttp.open("GET", "/get_user_info?user_id=" + userId, false);
        xmlHttp.send()
        let userBalance = JSON.parse(xmlHttp.responseText)["Balance"]
        event.preventDefault();
        let balance = document.getElementById("balance")
        balance.value = userBalance + " Ɇ"
    } else {
        alert(error)
    }
}

function chooseRandomRecipient() {
    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/random_wallet?user_id=" + userId, false);
    xmlHttp.send()
    let userWallet = JSON.parse(xmlHttp.responseText)["Wallet"]
    event.preventDefault();
    let wallet = document.getElementById("recipient")
    wallet.value = userWallet
}

// https://localhost:6006/new_transaction?user_id=2&recipient_wallet=c387354a-54a5-4890-9474-c7352d8bac38&amount=10

// Restricts input for the given textbox to the given inputFilter function.
function setInputFilter(textbox, inputFilter) {
    ["input", "keydown", "keyup", "mousedown", "mouseup", "select", "contextmenu", "drop"].forEach(function(event) {
        textbox.addEventListener(event, function() {
            if (inputFilter(this.value)) {
                this.oldValue = this.value;
                this.oldSelectionStart = this.selectionStart;
                this.oldSelectionEnd = this.selectionEnd;
            } else if (this.hasOwnProperty("oldValue")) {
                this.value = this.oldValue;
                this.setSelectionRange(this.oldSelectionStart, this.oldSelectionEnd);
            } else {
                this.value = "";
            }
        });
    });
}

setInputFilter(document.getElementById("amount"), function(value) {
    return /^\d*$/.test(value);
});

function TransactionsRedirect() {
    window.location = "/transactions?user_id=" + userId;
}

function APIRedirect() {
    window.location = "/api_docs";
}

function ChangeTransactionsColorDown() {
    let TransactionsLink = document.getElementById("links_transactions")
    TransactionsLink.style.color = "#bfbfbf"
    setTimeout(() => {  TransactionsLink.style.color = "#3bb31e" }, 50);
}

function ChangeAPIDocksColorDown() {
    let apiDocsLink = document.getElementById("links_api")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#3bb31e" }, 50);
}
