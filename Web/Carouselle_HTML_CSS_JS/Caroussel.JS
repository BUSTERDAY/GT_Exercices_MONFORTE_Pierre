const carouselSlide = document.querySelector('.carousel-slide');
const prevButton = document.querySelector('.carousel-prev');
const nextButton = document.querySelector('.carousel-next');
const images = document.querySelectorAll('.carousel-slide img');

let counter = 0;
const slideWidth = images[0].clientWidth;
const slideInterval = 3000; // temps en millisecondes entre chaque image
let slideTimer;

// Fonction pour faire défiler les images automatiquement
function autoSlide() {
  slideTimer = setInterval(() => {
    if (counter >= images.length - 1) {
      counter = 0;
    } else {
      counter++;
    }
    carouselSlide.style.transform = `translateX(${-slideWidth * counter}px)`;
  }, slideInterval);
}

// Écouteurs d'événements pour les boutons "Précédent" et "Suivant"
nextButton.addEventListener('click', () => {
  if (counter >= images.length - 1) {
    counter = 0;
  } else {
    counter++;
  }
  carouselSlide.style.transform = `translateX(${-slideWidth * counter}px)`;
  clearInterval(slideTimer);
});

prevButton.addEventListener('click', () => {
  if (counter <= 0) {
    counter = images.length - 1;
  } else {
    counter--;
  }
  carouselSlide.style.transform = `translateX(${-slideWidth * counter}px)`;
  clearInterval(slideTimer);
});

// Démarrer le défilement automatique
autoSlide();
