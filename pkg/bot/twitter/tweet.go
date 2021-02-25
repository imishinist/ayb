package twitter

func (c *Client) Tweet(text string) error {
	_, _, err := c.inner.Statuses.Update(text, nil)
	if err != nil {
		return err
	}
	return nil
}
