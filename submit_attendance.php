<?php
// Check if the form was submitted
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // Get the student's name, batch, and food preference from the form
    $studentName = $_POST["student_name"];
    $batch = $_POST["batch"];
    $foodPreference = $_POST["food_preference"];

    // Create a string to store the attendance information
    $attendanceInfo = "Name: $studentName\nBatch: $batch\nFood Preference: $foodPreference\n";

    // Define the file where attendance data will be saved
    $attendanceFile = "attendance.txt";

    // Open the file for writing (create if it doesn't exist)
    $fileHandle = fopen($attendanceFile, "a");

    // Check if the file was opened successfully
    if ($fileHandle === false) {
        echo "Error: Unable to open the attendance file.";
    } else {
        // Write the attendance information to the file
        fwrite($fileHandle, $attendanceInfo . "\n");

        // Close the file
        fclose($fileHandle);

        // Display a success message to the user
        echo "Attendance submitted successfully!";
    }
} else {
    // If the form was not submitted via POST, display an error
    echo "Error: Form not submitted.";
}
?>
