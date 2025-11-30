document.addEventListener("DOMContentLoaded", function () {
    if (!isUserLogged()) {
      window.location = 'login.html?reason=login_required';
    }
  
    document.getElementById("formAlimento").addEventListener("submit", async function (event) {
      event.preventDefault(); 
      
      const momentos = obtenerMomentosSeleccionados();
      
      if (momentos.length === 0) {
        alert("Debes seleccionar al menos un momento.");
        return; 
      }
  
      await agregarReceta(momentos); 
      this.reset(); 
    });
  });
  

  function obtenerMomentosSeleccionados() {
    const momentosSeleccionados = [];
    const momentos = ["Almuerzo", "Merienda", "Cena", "Desayuno"];
  
    momentos.forEach(momento => {
      const checkbox = document.getElementById(momento);
      if (checkbox.checked) {
        momentosSeleccionados.push(checkbox.value);
      }
    });
  
    return momentosSeleccionados;
  }
  
  async function agregarReceta(momentos) {
    try {
      const nombre = document.getElementById("NombreAlimento").value;
      const tipo = document.getElementById("TipoAlimento").value;
      const precio = parseFloat(document.getElementById("PrecioUnitario").value);
      const cantidadActual = parseInt(document.getElementById("CantidadActual").value);
      const cantidadMinima = parseInt(document.getElementById("CantidadMinima").value);
  
      const data = {
        NombreAlimento: nombre,
        TipoAlimento: tipo,
        PrecioUnitario: precio,
        CantidadActual: cantidadActual,
        CantidadMinima: cantidadMinima,
        MomentoAlimento: momentos,
      };
  
      await makeRequest(
        "http://localhost:8080/alimentos",
        Method.POST,
        data,
        ContentType.JSON,
        CallType.PRIVATE,
        successFn(),
        errorFn
      );
    } catch (error) {
      console.error("Error : ", error?.message || error);
      alert("Carga fallida. Intente nuevamente.");
    }
  }
  function errorFn(){
    console.error("Error al guardar los cambios del alimento");
    alert("Hubo un problema al guardar los cambios.");
  }
  function successFn() {
    alert("Alimento agregado exitosamente");
    window.location = 'indiceAlimento.html';
  }
  
