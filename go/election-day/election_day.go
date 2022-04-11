package electionday

import "strconv"

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	var votes = initialVotes
	return &votes
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if nil == counter {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	if nil == counter {
		return
	}
	*counter += increment
}

// NewElectionResult creates a new election result
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var elect ElectionResult
	elect.Name = candidateName
	elect.Votes = votes
	return &elect

}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	elected := *result
	display := elected.Name + " (" + strconv.Itoa(elected.Votes) + ")"
	return display

}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	_, exists := results[candidate]
	if exists {
		results[candidate]--
	}
}
