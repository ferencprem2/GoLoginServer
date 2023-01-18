var button = document.getElementById("getInfosButton")

button.addEventListener("click", () => {
    let email = document.querySelector("#emailInput")
    let password = document.querySelector("#passwordInput")
    let result = document.querySelector(".result")

    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8000/login"

    xhr.open("POST", url, true)
    xhr.setRequestHeader("Content-Type", "application/json");

    xhr.onreadystatechange = function () {
        if  (xhr.readyState === 4 && xhr.status === 200){
            result.innerHTML = this.responseText
        }
    };

    var data = JSON.stringify({"email": email.value, "password": password.value})

    xhr.send(data);
})
