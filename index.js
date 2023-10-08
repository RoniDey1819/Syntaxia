document.addEventListener('DOMContentLoaded', function () {
    const nonVegCheckbox = document.getElementById('checkbox-nonveg');
    const vegCheckbox = document.getElementById('checkbox-veg');
    const form = document.getElementById('details-form');

    // Function to check if at least one checkbox is selected
    function validateCheckboxSelection() {
        return nonVegCheckbox.checked || vegCheckbox.checked;
    }

    // Add a submit event listener to the form
    form.addEventListener('submit', function (event) {
        if (!validateCheckboxSelection()) {
            alert('Please select at least one food preference (Non-Veg or Veg).');
            event.preventDefault(); // Prevent form submission
        }
    });

    nonVegCheckbox.addEventListener('change', function () {
        if (nonVegCheckbox.checked) {
            vegCheckbox.checked = false;
        }
    });

    vegCheckbox.addEventListener('change', function () {
        if (vegCheckbox.checked) {
            nonVegCheckbox.checked = false;
        }
    });
});