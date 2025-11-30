const url = "http://localhost:8080/recetas";

document.addEventListener("DOMContentLoaded", function (eventDOM) {
  if (!isUserLogged()) {
    window.location = 'login.html?reason=login_required';
  }

  obtenerRecetas();

  document.getElementById("btnFiltrar").addEventListener("click", function (eventClick) {
    eventClick.preventDefault();
    obtenerRecetas();
  })
})

async function obtenerRecetas() {
  try {
    const momentoAlimento = document.getElementById('FiltroMomentoAlimento').value;
    const nombreAlimento = document.getElementById('FiltroNombreAlimento').value;
    const tipoAlimento = document.getElementById('FiltroTipoAlimento').value;

    if (nombreAlimento == "" && momentoAlimento == "" && tipoAlimento == ""){
      await makeRequest(`${url}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
    }else{
      await makeRequest(`${url}?momento=${momentoAlimento}&nombre=${nombreAlimento}&tipo=${tipoAlimento}`, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);
     }} catch (error) {
    console.error("Error [obtenerRecetas]: ", error?.message || error)
  }
}

function successFn(data) {
  const elementosTable = document.querySelector("#elementosTable tbody");

  while (elementosTable.firstChild) {
    elementosTable.removeChild(elementosTable.firstChild);
  }

  data.forEach(receta => {
    const row = document.createElement("tr"); 
    const alimentos = receta.AlimentosNecesarios.map(alimento => `${alimento.NombreAlimento} (${alimento.CantidadNecesaria})`).join(', ');
    row.innerHTML = `
          <td>${receta.NombreReceta}</td>
          <td>${receta.MomentoReceta}</td>
          <td>${alimentos}</td>
          <td class="acciones">
            <a style="color: #362100" href="editarReceta.html?id=${receta.ID}&tipo=EDITAR">Editar</a> | 
            <a style="color: #362100" href="#" onclick="eliminarReceta('${receta.ID}')">Eliminar</a>
          </td>
      `; 
    elementosTable.appendChild(row);
  }); 
}


async function eliminarReceta(id) {
  console.log("Intentando eliminar la receta con ID:", id); 
  if (confirm("¿Estás seguro de que deseas eliminar esta receta?")) {
 try{ await makeRequest(`http://localhost:8080/recetas/${id}`, Method.DELETE, null, ContentType.JSON, CallType.PRIVATE, successFn2, errorFn);
} catch (error) {
  console.error("Error [eliminarReceta]: ", error?.message || error)
  
}}}
function successFn2() {
  alert("Eliminada con éxito");
  location.reload(); 
}

function errorFn(status, response) {
  console.log("Falla:", response);
  alert("Error al eliminarla", error)
}