function copyToClipboard() {
    var copyText = document.getElementById("wallet");
    copyText.select();
    copyText.setSelectionRange(0, 99999);
    navigator.clipboard.writeText(copyText.value);
    alert("The wallet id is copied!");
}