package comment

import (
	"context"

	"github.com/google/go-github/github"

	client "github.com/mattmoor/bindings/pkg/github"
)

func CleanupOlder(ctx context.Context, name, owner, repo string, number int) error {
	ghc, err := client.New(ctx)
	if err != nil {
		return err
	}

	var ids []int64

	lopt := &github.IssueListCommentsOptions{}
	for {
		comments, resp, err := ghc.Issues.ListComments(ctx, owner, repo, number, lopt)
		if err != nil {
			return err
		}
		for _, comment := range comments {
			if HasSignature(name, comment.GetBody()) {
				ids = append(ids, comment.GetID())
			}
		}
		if resp.NextPage == 0 {
			break
		}
		lopt.Page = resp.NextPage
	}

	for _, id := range ids {
		_, err := ghc.Issues.DeleteComment(ctx, owner, repo, id)
		if err != nil {
			return err
		}
	}

	return nil
}
