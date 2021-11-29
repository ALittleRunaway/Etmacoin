let userId = 0
let userLogin = ""
let userWallet = ""

function fillLatestTransactions(latestTransactions) {
    let tbl = document.getElementById("table_all_transactions")
    let tblBody = document.createElement("tbody");

    for (let i = 0; i < latestTransactions.length; i++) {
        let row = document.createElement("tr");
        let obj = latestTransactions[i];

        for (let key in obj){
            let value = ""
            if (key === "Amount") {
                value = obj[key] + " Ɇ";
            } else if (key === "Time") {
                value = obj[key].replaceAll("T", " ");
            } else {
                value = obj[key];
            }
            let cell = document.createElement("td");
            let cellText = document.createTextNode(value);
            cell.appendChild(cellText);
            row.appendChild(cell);
        }
        tblBody.appendChild(row);
    }
    tbl.appendChild(tblBody);
}

function fillUserTransactions(userTransactions) {
    let tbl = document.getElementById("table_user_transactions")
    let tblBody = document.createElement("tbody");

    for (let i = 0; i < userTransactions.length; i++) {
        let row = document.createElement("tr");
        let obj = userTransactions[i];

        for (let key in obj){
            let value = ""
            if (key === "Amount") {
                value = obj[key] + " Ɇ";
            } else if (key === "Time") {
                value = obj[key].replaceAll("T", " ");
            } else {
                value = obj[key];
            }
            let cell = document.createElement("td");
            let cellText = document.createTextNode(value);
            cell.appendChild(cellText);
            row.appendChild(cell);
        }
        tblBody.appendChild(row);
    }
    tbl.appendChild(tblBody);
}

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
    let userTransactions = JSON.parse(xmlHttp.responseText)["Transactions"]
    fillUserTransactions(userTransactions)

    xmlHttp.open("GET", "/latest_transactions", false);
    xmlHttp.send()
    let latestTransactions = JSON.parse(xmlHttp.responseText)["Transactions"]
    fillLatestTransactions(latestTransactions)
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

function SignOutRedirect() {
    let result = confirm("Are you sure you want to sign out?");
    if (result) {
        window.location = "/";
    } else {}
}

function ChangeTransactionsColorDown() {
    let TransactionsLink = document.getElementById("links_transactions")
    TransactionsLink.style.color = "#bfbfbf"
    setTimeout(() => {  TransactionsLink.style.color = "#555555" }, 50);
}

function ChangeAPIDocksColorDown() {
    let apiDocsLink = document.getElementById("links_api")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#3bb31e" }, 50);
}

function ChangeHomepageColorDown() {
    let apiDocsLink = document.getElementById("links_homepage")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#3bb31e" }, 50);
}

function ChangeSignOutColorDown() {
    let apiDocsLink = document.getElementById("links_sign_out")
    apiDocsLink.style.color = "#bfbfbf"
    setTimeout(() => {  apiDocsLink.style.color = "#3bb31e" }, 50);
}
