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

function validatePassword(pass) {
    if (!(/[a-z]/.test(pass))) {
        return "There are no lowercase letters in your password!"
    } else if (!(/[A-Z]/.test(pass))) {
        return "There are no uppercase letters in your password!";
    } else if (!(/[0-9]/.test(pass))) {
        return "The are no digits in the password!";
    } else if (!(/[!@#$%^&*-]/.test(pass))) {
        return "The are no special characters in the password!";
    } else if (pass.length < 8) {
        return "The password has to be at least 8 characters long!"
    }
    return ""
}

function registerUser() {
    let login = window.document.getElementById("first_login").value.trim().replace(/\s\s+/g, ' ');
    let pass = window.document.getElementById("first_pass").value
    let pass_confirm = window.document.getElementById("first_pass_confirm").value

    let error = validateRegistrationData(login, pass, pass_confirm)
    let error_pass = validatePassword(pass)

    if (error === "") {
        if (error_pass === "") {
            alert("You've been registered successfully!")
            let xmlHttp = new XMLHttpRequest();
            xmlHttp.open("GET", "/new_user?login=" + login + "&pass=" + pass, false);
            xmlHttp.send()
            userId = JSON.parse(xmlHttp.responseText)["Id"]
            event.preventDefault();
            window.location = "/homepage?user_id=" + userId;
        } else {
            alert(error_pass)
        }
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

// Restricts input for the given textbox to the given inputFilter.
function setInputFilter(textbox, inputFilter) {
        textbox.addEventListener("input", function() {
            if (inputFilter(this.value)) {
                this.oldValue = this.value;
                this.oldSelectionStart = this.selectionStart;
                this.oldSelectionEnd = this.selectionEnd;
            } else if (this.hasOwnProperty("oldValue")) {
                this.value = this.oldValue;
                this.setSelectionRange(this.oldSelectionStart, this.oldSelectionEnd);
            } else {
                this.value = "";
            }});
}

console.log(document.getElementById("first_login"))

setInputFilter(document.getElementById("first_login"), function(value) {
    return /^[ A-Za-z-_\d]*$/.test(value);
});
