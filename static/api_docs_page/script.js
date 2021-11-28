let userId = 0
let userLogin = ""
let userWallet = ""

function getUserAPIDocsInfo() {
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


function TransactionsRedirect() {
    window.location = "/transactions?user_id=" + userId;
}

function APIRedirect() {
    window.location = "/api_docs?user_id=" + userId;
}

function HomepageRedirect() {
    window.location = "/homepage?user_id=" + userId;
}

function ChangeTransactionsColorDown() {
    let TransactionsLink = document.getElementById("links_transactions")
    TransactionsLink.style.color = "#bfbfbf"
    setTimeout(() => {  TransactionsLink.style.color = "#3bb31e" }, 50);
}

function ChangeAPIDocksColorDown() {
    let apiDocsLink = document.getElementById("links_api")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#555555" }, 50);
}

function ChangeHomepageColorDown() {
    let apiDocsLink = document.getElementById("links_homepage")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#3bb31e" }, 50);
}
