document.addEventListener('DOMContentLoaded', function() {
    // Mobile menu elements
    const mobileMenuBtn = document.querySelector('.mobile-menu-btn');
    const navLinks = document.querySelector('.nav-links');
    const menuIcon = mobileMenuBtn.querySelector('i');

    // Toggle mobile menu
    function toggleMenu() {
        navLinks.classList.toggle('active');
        if (navLinks.classList.contains('active')) {
            menuIcon.classList.remove('fa-bars');
            menuIcon.classList.add('fa-times');
            document.body.style.overflow = 'hidden'; // Блокируем скролл страницы
        } else {
            menuIcon.classList.remove('fa-times');
            menuIcon.classList.add('fa-bars');
            document.body.style.overflow = ''; // Восстанавливаем скролл
        }
    }

    // Event listeners
    mobileMenuBtn.addEventListener('click', toggleMenu);

    // Close menu when clicking on links
    document.querySelectorAll('.nav-links a').forEach(link => {
        link.addEventListener('click', function() {
            if (window.innerWidth <= 992) { // Только для мобильного вида
                navLinks.classList.remove('active');
                menuIcon.classList.remove('fa-times');
                menuIcon.classList.add('fa-bars');
                document.body.style.overflow = ''; // Восстанавливаем скролл
            }
        });
    });

    // Smooth scrolling for anchor links
    document.querySelectorAll('a[href^="#"]').forEach(anchor => {
        anchor.addEventListener('click', function(e) {
            const targetId = this.getAttribute('href');
            if (targetId === '#') return;

            e.preventDefault();
            const targetElement = document.querySelector(targetId);
            if (targetElement) {
                window.scrollTo({
                    top: targetElement.offsetTop - 80,
                    behavior: 'smooth'
                });
            }
        });
    });

    // Responsive check on resize
    window.addEventListener('resize', function() {
        if (window.innerWidth > 992) {
            // На больших экранах убедимся, что меню закрыто
            navLinks.classList.remove('active');
            menuIcon.classList.remove('fa-times');
            menuIcon.classList.add('fa-bars');
            document.body.style.overflow = '';
        }
    });

    // Animation on scroll
    function animateOnScroll() {
        const elements = document.querySelectorAll('.animate-card, .animate-fade-in, .animate-slide-up');

        elements.forEach(element => {
            const elementPosition = element.getBoundingClientRect().top;
            const windowHeight = window.innerHeight;

            if (elementPosition < windowHeight - 100) {
                element.style.animationPlayState = 'running';
            }
        });
    }

    // Initial animations
    animateOnScroll();
    window.addEventListener('scroll', animateOnScroll);
});