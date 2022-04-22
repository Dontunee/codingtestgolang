# Instructions

Format: a 45 min pair programming session followed by 15 minutes of discussion and feedback.
Goal: We'll be creating a new golang http service together with net/http and writing unit tests using go test.
Preparation: Candidate should have a golang development environment set up on his laptop and be prepared to screen share as we work together.

# Getting started
1. Download this git repository to candidate's machine
2. Verify environment: use "go run ." to verify your http server works and go to http://localhost:8080 (you can kill the process once verified)
3. Running tests: use "go test -v ./..."

Data model:
- A listing is a house for sale
- each listing has an id, street address, price, status
- price is whole dollars between $0 and 100 billion
- status can be for sale, sold or expired