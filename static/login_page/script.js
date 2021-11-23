let userId

function validateRegistrationData(login, pass, pass_confirm) {
    if (login === "") {
        return "Login can not be empty!"
    } else if (pass === "") {
        return "Password can not be empty!"
    } else if (pass !== pass_confirm) {
        return "The passwords do not mach!"
    }
    return ""
}

function registerUser() {
    let login = window.document.getElementById("first_login").value
    let pass = window.document.getElementById("first_pass").value
    let pass_confirm = window.document.getElementById("first_pass_confirm").value

    let error = validateRegistrationData(login, pass, pass_confirm)

    if (error === "") {
        alert("You've been registered successfully!")
        let xmlHttp = new XMLHttpRequest();
        xmlHttp.open("GET", "/new_user?login=" + login + "&pass=" + pass, false);
        xmlHttp.send()
        userId = JSON.parse(xmlHttp.responseText)["Id"]
        event.preventDefault();
        window.location = "/homepage?user_id=" + userId;
    } else {
        alert(error)
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

setInputFilter(document.getElementById("first_login"), function(value) {
    return /^([A-Z]|[a-z]|[0-9]|_|-| )*$/.test(value);
});