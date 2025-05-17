/**
 * Shows the modal image popup
 * @param {string} title title of the image
 * @param {string} excerpt text description of the image
 * @param {string} src image source URL
 */
function showImageModal(title, excerpt, src) {
    const modal = document.getElementById("modal");
    const modalTitle = document.getElementById("modal-title");
    const modalExcerpt = document.getElementById("modal-excerpt");
    const modalImage = document.getElementById("modal-image");
    
    if (modal && modalTitle && modalExcerpt && modalImage) {
        modalTitle.textContent = title;
        modalExcerpt.textContent = excerpt;
        modalImage.src = src;
        modal.showModal();
    }
}

// MARK: Images Modal
document.addEventListener('DOMContentLoaded', function () {
  const menuToggle = document.getElementById('menu-toggle');
  if (menuToggle) {
    menuToggle.addEventListener('click', function () {
      const menu = document.getElementById('mobile-menu');
      if (menu) menu.classList.toggle('hidden');
    });
  }
});

  // Initialize modal functionality
  const modal = document.getElementById("modal");
  if (modal) {
    // Close modal when clicking outside
    modal.addEventListener('click', function (event) {
      const rect = modal.getBoundingClientRect();
      const isInDialog = (rect.top <= event.clientY && event.clientY <= rect.top + rect.height &&
        rect.left <= event.clientX && event.clientX <= rect.left + rect.width);
      if (!isInDialog) {
        modal.close();
      }
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

  document.getElementById('demo-form').addEventListener('submit', function () {
    const event = new Event('verified');
    const elem = document.querySelector("#demo-form");
    elem.dispatchEvent(event);
  })
