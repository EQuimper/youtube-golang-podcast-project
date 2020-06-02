package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"ytpodcast/feeds"
	"ytpodcast/graph/generated"
	"ytpodcast/graph/model"
	"ytpodcast/itunes"
	"ytpodcast/utils"
)

func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search(term)
	if err != nil {
		return nil, err
	}

	var podcasts []*model.Podcast

	for _, res := range res.Results {
		podcast := &model.Podcast{
			Artist:        res.ArtistName,
			PodcastName:   res.TrackName,
			FeedURL:       res.FeedURL,
			Thumbnail:     res.ArtworkURL100,
			EpisodesCount: res.TrackCount,
			Genres:        res.Genres,
		}

		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feedItems []*model.FeedItem

	for _, item := range res.Channel.Item {
		feedItem := &model.FeedItem{
			PubDate:     item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       utils.CheckNullString(item.Image.Href),
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
