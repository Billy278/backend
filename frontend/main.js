const menuToggle = document.querySelector(".toggle");
const navItem = document.querySelector(".list ul");
menuToggle.addEventListener("click", () => {
  navItem.classList.toggle("slide");
});
