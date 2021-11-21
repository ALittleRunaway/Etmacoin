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
        xmlHttp.open("GET", "/new_user?login=" + login + "&pass=" + pass, false);
        xmlHttp.send()
        userId = JSON.parse(xmlHttp.responseText)["Id"]
        event.preventDefault();
        window.location = "/homepage?user_id=" + userId;
    }
}
function loginUser() {
    let login = window.document.getElementById("login").value
    let pass = window.document.getElementById("pass").value

    let xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", "/login_user?login=" + login + "&pass=" + pass, false);
    xmlHttp.send()
    userId = JSON.parse(xmlHttp.responseText)["Id"]
    if (userId !== 0) {
        event.preventDefault();
        window.location = "/homepage?user_id=" + userId;
    } else {
        alert("Your password is wrong either you are not signed up!")
    }
}