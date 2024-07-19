package poke

func cmdPokedex(c *Config, args ...string) error {
	c.Pokedex.Pokedex()
	return nil
}
