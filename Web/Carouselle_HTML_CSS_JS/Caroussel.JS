const carouselSlide = document.querySelector('.carousel-slide');
const prevButton = document.querySelector('.carousel-prev');
const nextButton = document.querySelector('.carousel-next');
const images = document.querySelectorAll('.carousel-slide img');

let counter = 0;
const slideWidth = images[0].clientWidth;

nextButton.addEventListener('click', () => {
  if (counter >= images.length - 1) return;
  counter++;
  carouselSlide.style.transform = `translateX(${-slideWidth * counter}px)`;
});

prevButton.addEventListener('click', () => {
  if (counter <= 0) return;
  counter--;
  carouselSlide.style.transform = `translateX(${-slideWidth * counter}px)`;
});
