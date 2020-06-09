function creatNewElement() {
    var li = document.createElement('li');
    var InputValue = document.getElementById("Input").value;
    var textNode = document.createTextNode(InputValue);
    li.appendChild(textNode);

    if (InputValue === '') {
        alert("Hey, il n'y a rien à faire")
    } else {
        document.getElementById("completed-tasks").appendChild(li);
    }

    document.getElementById("Input").value = "";




}