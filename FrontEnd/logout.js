document.addEventListener("DOMContentLoaded", function(eventDOM) {
    // Configura el botón de logout
    document.getElementById("btnLogout").addEventListener("click", function(eventClick) {
        eventClick.preventDefault();

        // Lógica para el logout
        logout();
    });
});

function logout() {
    localStorage.removeItem("authToken"); 
    window.location = 'login.html';
}