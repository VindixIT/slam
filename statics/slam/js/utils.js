function filtraTabela(input, tabelaNome, offset, colnum){
  var filter, table, tr, td, i, txtValue;
  filter = input.value.toUpperCase();
  table = document.getElementById(tabelaNome);
  tr = table.getElementsByTagName("tr");
  for (i = offset; i < tr.length; i++) {
    td = tr[i].getElementsByTagName("td")[colnum];
	console.log(td.innerText);
    if (td) {
      txtValue = td.textContent || td.innerText;
      if (txtValue.toUpperCase().indexOf(filter) > -1) {
        tr[i].style.display = "table-row";
      } else {
        tr[i].style.display = "none";
      }
    }       
  }
}
