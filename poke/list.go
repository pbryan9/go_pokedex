package poke

func cmdList(c *Config, args ...string) error {
	c.Pokedex.List()
	return nil
}
