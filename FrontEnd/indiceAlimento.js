const url = "http://localhost:8080/alimentos";

document.addEventListener("DOMContentLoaded", function (eventDOM) {
  if (!isUserLogged()) {
    window.location = 'login.html?reason=login_required';
  }
  const elementosTable = document.querySelector("#elementosTable tbody");

  obtenerAlimentos();

  document.getElementById("btnFiltrar").addEventListener("click", function (eventClick) {
    eventClick.preventDefault();
    obtenerAlimentos();
  })
})

async function obtenerAlimentos() {
  try {
    const nombre = document.getElementById("FiltroNombre").value;
    function capitalizarPrimeraLetra(texto) {
      if (!texto) return texto; 
      return texto
        .split(' ')
        .map(palabra => palabra.charAt(0).toUpperCase() + palabra.slice(1).toLowerCase())
        .join(' ');
    }
    const nombreCapitalizado = capitalizarPrimeraLetra(nombre);
    if (nombreCapitalizado == ""){
      await makeRequest(`${url}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
    }else{
    await makeRequest(`${url}?tipo=${nombreCapitalizado}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);}
  } catch (error) {
    console.error("Error [obtenerAlimentos]: ", error?.message || error)
  }
}

function successFn(data) {
  const elementosTable = document.querySelector("#elementosTable tbody");
  while (elementosTable.firstChild) {
    elementosTable.removeChild(elementosTable.firstChild);
  }
  data.forEach(elemento => {
    const row = document.createElement("tr"); 
    

    row.innerHTML = `
          <td>${elemento.NombreAlimento}</td>
          <td>${elemento.TipoAlimento}</td>
          <td>${elemento.PrecioUnitario}</td>
          <td>${elemento.CantidadActual}</td>
          <td>${elemento.CantidadMinima}</td>
          <td>${elemento.MomentoAlimento.join(', ')}</td>
          <td class="acciones">
            <a style="color: #362100" href="editarAlimento.html?id=${elemento.IDAlimento}&tipo=EDITAR">Editar</a> | 
            <a style="color: #362100" href="#" onclick="eliminarAlimento('${elemento.IDAlimento}')">Eliminar</a>
          </td>
      `; 
    elementosTable.appendChild(row);
  });
}
async function eliminarAlimento(id) {
  console.log("Intentando eliminar el alimento con ID:", id); 
  if (confirm("¿Estás seguro de que deseas eliminar este alimento?")) {
 try{ await makeRequest(`http://localhost:8080/alimentos/${id}`, Method.DELETE, null, ContentType.JSON, CallType.PRIVATE, successFn2, errorFn);
} catch (error) {
  console.error("Error [eliminarAlimento]: ", error?.message || error)
  
}}}
function successFn2(){
  alert("Eliminado con exito");
  obtenerAlimentos();}

function errorFn(status, response) {
  console.log("Falla:", response);
}