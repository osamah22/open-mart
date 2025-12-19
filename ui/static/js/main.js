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

    if (!dropdown || !button) return;

    button.addEventListener("click", function (e) {
        e.stopPropagation();
        dropdown.classList.toggle("open");
        button.setAttribute(
            "aria-expanded",
            dropdown.classList.contains("open")
        );
    });

    document.addEventListener("click", function () {
        dropdown.classList.remove("open");
        button.setAttribute("aria-expanded", "false");
    });
})();
