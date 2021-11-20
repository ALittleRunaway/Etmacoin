let userId

function registerUser() {
    let login = window.document.getElementById("first_login").value
    let pass = window.document.getElementById("first_pass").value
    let pass_confirm = window.document.getElementById("first_pass_confirm").value

    if (pass !== pass_confirm) {
        alert("The passwords don't match!")
    } else {
        alert("You've been registered successfully!")
        let xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", "http://localhost:6006/new_user?login=" + login + "&pass=" + pass, false);
        xmlHttp.send()
        userId = JSON.parse(xmlHttp.responseText)["Id"]
        event.preventDefault();
        window.location = "http://localhost:6006/homepage?user_id=" + userId;
    }
}