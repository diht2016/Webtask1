var root = document.getElementById("root");
var ul = document.createElement("ul");
root.appendChild(ul);

function put(text) {
	var li = document.createElement("li");
	var span = document.createElement("span");
	span.textContent = text;
	var del = document.createElement("button");
	del.textContent = "Удалить";
	del.onclick = function() {
		ul.removeChild(li);
	}
	li.appendChild(span);
	li.appendChild(del);
	ul.appendChild(li);
}

put("Сделать задание #3 по web-программированию");

var inp = document.createElement("input");
inp.id = "add_task_input";
var add = document.createElement("button");
add.id = "add_task";
add.textContent = "Добавить";
add.onclick = function() {
	put(inp.value);
}
root.appendChild(inp);
root.appendChild(add);
