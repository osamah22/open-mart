// var navLinks = document.querySelectorAll("nav a");
// for (var i = 0; i < navLinks.length; i++) {
// 	var link = navLinks[i]
// 	if (link.getAttribute('href') == window.location.pathname) {
// 		link.classList.add("live");
// 		break;
// 	}
// }
(function () {
    const dropdown = document.getElementById("userDropdown");
    const button = document.getElementById("userDropdownBtn");
    const arrow = document.getElementById("menu-arrow");

    if (!dropdown || !button || !arrow) return;

    button.addEventListener("click", function (e) {
        e.stopPropagation();

        const isOpen = dropdown.classList.toggle("open");

        button.setAttribute("aria-expanded", isOpen);

        // Toggle arrow direction
        arrow.classList.toggle("fa-chevron-down", !isOpen);
        arrow.classList.toggle("fa-chevron-up", isOpen);
    });

    document.addEventListener("click", function () {
        dropdown.classList.remove("open");
        button.setAttribute("aria-expanded", "false");

        arrow.classList.add("fa-chevron-down");
        arrow.classList.remove("fa-chevron-up");
    });
})();
