function editUser(e) {
	var editForm = document.getElementById('edit-form');
	editForm.style.display = 'block';
	var userId = e.parentNode.parentNode.childNodes[3].innerText;
	var userName = e.parentNode.parentNode.childNodes[5].innerText;
	var userUsername = e.parentNode.parentNode.childNodes[7].innerText;
	var userEmail = e.parentNode.parentNode.childNodes[9].innerText;
	var userMobile = e.parentNode.parentNode.childNodes[11].innerText;
	var userRole = e.parentNode.parentNode.childNodes[15].childNodes[1].value;
	document.getElementById('userIdToUpdate').value = userId;
	document.getElementById('userName').value = userName;
	document.getElementById('userUsername').value = userUsername;
	document.getElementById('userEmail').value = userEmail;
	document.getElementById('userMobile').value = userMobile;
	document.getElementById('RoleForUpdate').value = userRole;
	document.getElementById('userName').focus();
}

function deleteUser(e) {
	var deleteForm = document.getElementById('delete-form');
	deleteForm.style.display = 'block';
	var userId = e.parentNode.parentNode.childNodes[3].innerText;
	document.getElementById('userIdToDelete').value = userId;
}

function openCreateUser(){
	document.getElementById('create-form').style.display='block';
	document.getElementById('NameUserForInsert').focus();
}

function naoVazio(...fields){
  fields.forEach(field => {
		console.log(field);
		if(field.length==0) {
			return false;
		}
	});
	return true;
}

function createUser(){
	let name = document.getElementById('NameUserForInsert').value;
	let username = document.getElementById('UserName').value;
	let password = document.getElementById('Password').value;
	let confirmation = document.getElementById('Confirmation').value;
	let email = document.getElementById('Email').value;
	let mobile = document.getElementById('Mobile').value;
	let qtdAtendimentos = document.getElementById('QtdAtendimentos').value;
	let tipoEspecialidade = document.getElementById('TipEspecialidade').value;
	let outraEspecialidade = document.getElementById('OutraEspecialidade').value;
	if(password == confirmation && 
		naoVazio(name, username, password, confirmation, email)){
		let xmlhttp;
		xmlhttp = new XMLHttpRequest();
		xmlhttp.onreadystatechange=function()
		{
				console.log(xmlhttp.response);
				if (xmlhttp.readyState==4 && xmlhttp.status==200)
				{
					document.getElementById('create-form').style.display='none';
					document.getElementById("messageText").innerText = "Sucesso";
					document.getElementById("message").style.display="block";
				} else {
					document.getElementById("Errors").innerText = xmlhttp.response;
					document.getElementById("error-message").style.display="block";
					return;		
				}
		}
		xmlhttp.open("POST","/createUser",true);
		let params = "name="+name+
		"&username="+username+
		"&password="+password+
		"&email="+email+
		"&mobile="+mobile+
		"&qtdAtendimentos="+qtdAtendimentos+
		"&tipoEspecialidade="+tipoEspecialidade+
		"&outraEspecialidade="+outraEspecialidade;
		xmlhttp.send(params);
	} else {
		let errorMsg = "Falta preencher campos do formulário.";
		document.getElementById("Errors").innerText = errorMsg;
		document.getElementById("error-message").style.display="block";
		return;		
	}
}