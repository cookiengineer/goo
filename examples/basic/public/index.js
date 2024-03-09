
document.addEventListener("DOMContentLoaded", () => {

	const buttons = Array.from(document.querySelectorAll("main section button[data-view]"));

	if (buttons.length > 0) {

		buttons.forEach((button) => {

			button.addEventListener("click", () => {

				let view = document.querySelector('main section[data-view="' + button.getAttribute('data-view') + '"]');
				if (view !== null) {

					let others = Array.from(document.querySelectorAll('main section[data-view]')).filter((o) => o !== view);
					if (others.length > 0) {
						others.forEach((element) => {
							element.classList.remove("active");
							element.classList.add("inactive");
						});
					}

					view.classList.remove("inactive");
					view.classList.add("active");

				}

			});

		});

	}

});
