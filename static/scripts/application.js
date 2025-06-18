// Global variables for image gallery
let currentImageIndex = 0;
let galleryImages = [];

/**
 * Shows the modal image popup
 * @param {string} title title of the image
 * @param {string} excerpt text description of the image
 * @param {string} src image source URL
 */
function showImageModal(title, excerpt, src) {
    const modal = document.getElementById("modal");
    if (!modal) return;

    // Update gallery images array if needed
    if (galleryImages.length === 0) {
        // Get all images with their metadata
        const imageElements = document.querySelectorAll('img[data-modal-image]');
        galleryImages = Array.from(imageElements).map(img => ({
            src: img.dataset.modalImage || img.src,
            title: img.dataset.modalTitle || '',
            excerpt: img.dataset.modalExcerpt || ''
        }));
    }

    // Find current image index
    currentImageIndex = galleryImages.findIndex(img => img.src === src);
    if (currentImageIndex === -1) {
        // If image not found in gallery, add it
        currentImageIndex = galleryImages.length;
        galleryImages.push({
            src: src,
            title: title || '',
            excerpt: excerpt || ''
        });
    }
    updateModalContent();
    modal.showModal();
}

/**
 * Updates the modal content with current image data
 */
function updateModalContent() {
    const modalTitle = document.getElementById("modal-title");
    const modalExcerpt = document.getElementById("modal-excerpt");
    const modalImage = document.getElementById("modal-image");
    const imageCounter = document.getElementById("image-counter");
    const navigationDots = document.getElementById("navigation-dots");
    
    if (!modalTitle || !modalExcerpt || !modalImage || !imageCounter || !navigationDots) return;

    const currentImage = galleryImages[currentImageIndex];
    if (!currentImage) {
        console.error('No image found at index:', currentImageIndex);
        return;
    }

    // Update navigation dots
    navigationDots.innerHTML = galleryImages.map((_, index) => `
        <button 
            class="w-3 h-3 rounded-full transition-all duration-300 ${index === currentImageIndex ? 'bg-blue-500 scale-110' : 'bg-gray-300 hover:bg-gray-400 dark:bg-gray-600 dark:hover:bg-gray-500'}"
            onclick="navigateToImage(${index})"
            aria-label="Go to image ${index + 1}"
        ></button>
    `).join('');
    
    modalTitle.textContent = currentImage.title;
    modalExcerpt.textContent = currentImage.excerpt;
    
    // Load new image
    const img = new Image();
    img.onload = function() {
        modalImage.src = currentImage.src;
        modalImage.style.opacity = '1';
    };
    modalImage.style.opacity = '0.5';
    img.src = currentImage.src;
    
    // Update counter
    imageCounter.textContent = `Image ${currentImageIndex + 1} of ${galleryImages.length}`;
}

/**
 * Navigate to a specific image in the gallery
 * @param {number} index The index of the image to navigate to
 */
function navigateToImage(index) {
    if (index >= 0 && index < galleryImages.length) {
        currentImageIndex = index;
        updateModalContent();
    }
}

// MARK: Images Modal
document.addEventListener('DOMContentLoaded', function () {
    const modal = document.getElementById("modal");
    if (!modal) return;

    // Setup navigation buttons
    const prevButton = document.getElementById('prevImage');
    const nextButton = document.getElementById('nextImage');

    if (prevButton) {
        prevButton.addEventListener('click', (e) => {
            e.stopPropagation();
            currentImageIndex = (currentImageIndex - 1 + galleryImages.length) % galleryImages.length;
            updateModalContent();
        });
    }

    if (nextButton) {
        nextButton.addEventListener('click', (e) => {
            e.stopPropagation();
            currentImageIndex = (currentImageIndex + 1) % galleryImages.length;
            updateModalContent();
        });
    }

    // Keyboard navigation
    document.addEventListener('keydown', (e) => {
        if (!modal.open) return;
        
        switch(e.key) {
            case 'ArrowLeft':
                currentImageIndex = (currentImageIndex - 1 + galleryImages.length) % galleryImages.length;
                updateModalContent();
                break;
            case 'ArrowRight':
                currentImageIndex = (currentImageIndex + 1) % galleryImages.length;
                updateModalContent();
                break;
            case 'Escape':
                modal.close();
                break;
        }
    });

    // Close modal when clicking outside
    modal.addEventListener('click', function (event) {
        if (event.target.closest('#navigation-dots')) return; // Don't close when clicking navigation dots
        const rect = modal.getBoundingClientRect();
        const isInDialog = (rect.top <= event.clientY && event.clientY <= rect.top + rect.height &&
            rect.left <= event.clientX && event.clientX <= rect.left + rect.width);
        if (!isInDialog) {
            modal.close();
        }
    });

    // Reset gallery when modal is closed
    modal.addEventListener('close', () => {
        currentImageIndex = 0;
        galleryImages = [];
    });

    const menuToggle = document.getElementById('menu-toggle');
    if (menuToggle) {
        menuToggle.addEventListener('click', function () {
            const menu = document.getElementById('mobile-menu');
            if (menu) menu.classList.toggle('hidden');
        });
    }

    const themeToggle = document.getElementById('theme-toggle');
    const lightIcon = document.getElementById('light-icon');
    const darkIcon = document.getElementById('dark-icon');

    const html = document.documentElement;

    if (localStorage.getItem('color-theme') === 'dark') {
        html.classList.add('dark');
    } else {
        html.classList.remove('dark');
    }

    themeToggle.addEventListener('click', function () {
        // Toggle dark class on HTML element
        html.classList.toggle('dark');

        if (html.classList.contains('dark')) {
            localStorage.setItem('color-theme', 'dark');
        } else {
            localStorage.setItem('color-theme', 'light');
        }
    });

    const demoForm = document.getElementById('demo-form');
    if (demoForm) {
        demoForm.addEventListener('submit', function () {
            const event = new Event('verified');
            demoForm.dispatchEvent(event);
        });
    }
});
