function copyToClipboard() {
    var copyText = document.getElementById("wallet");
    copyText.select();
    copyText.setSelectionRange(0, 99999);
    navigator.clipboard.writeText(copyText.value);
    alert("The wallet id is copied!");
}

function getUserInfo() {
    // alert(window.location.href)
    const urlParams = new URLSearchParams(window.location.search);
    let userId = urlParams.get('user_id');

    let balance = document.getElementById("balance")
    balance.value = "1000 $" + userId
}