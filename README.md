# 🔑 Passphrase Generator

This is a simple Go-based passphrase generator that creates strong, memorable passphrases using a mix of syllables and digits. It's inspired by [Ricky Mondello's blog post](https://rmondello.com/2024/10/07/apple-passwords-generated-strong-password-format/) about Apple's strong password format.

## How It Works

- The passphrase is generated using random syllables in a consonant-vowel-consonant-(consonant)-vowel-consonant pattern.
- One syllable is capitalized, and a random digit is inserted before or after one of the syllables.
- Passphrases are output in the format: `syllable-syllable-syllable-syllable`.

## Requirements

- Go
