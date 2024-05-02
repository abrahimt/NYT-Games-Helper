// Read Dictionary: Load a dictionary of words into memory. This dictionary will be used to check if potential solutions are valid English words.
// Filter Words by Length: Filter the dictionary to include only words of the same length as the target word in the Wordle game.
// Guessing Loop:
// Start a loop to make guesses until the correct word is found or the maximum number of guesses is reached.
// Make an initial guess based on common patterns or randomly selected from the filtered dictionary.
// After each guess, receive feedback on the correctness of the guess:
// The number of correct letters in the correct positions (green marks).
// The number of correct letters in incorrect positions (yellow marks).
// Use the feedback to refine subsequent guesses:
// Eliminate words from the dictionary that do not match the feedback.
// Generate new guesses based on the remaining potential solutions.
// Display Results:
// Once the correct word is found, display it along with the number of guesses made.
// If the maximum number of guesses is reached without finding the correct word, indicate that the solver failed to find the solution.