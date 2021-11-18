

function registerUser() {
    let login = window.document.getElementById("first_login").value
    let pass = window.document.getElementById("first_pass").value
    let pass_confirm = window.document.getElementById("first_pass_confirm").value

    if (pass !== pass_confirm) {
        alert("The passwords don't match!")
    } else {
        alert("You've been registered successfully!")
        event.preventDefault();
        window.location = "http://localhost:6006/homepage";
    }
}