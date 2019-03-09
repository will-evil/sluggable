package sluggable

// Sluggabler is interface who use for create slug from object.
// IsSlugUnique uses for check slug unique
type Sluggabler interface {
	IsSlugUnique(slug string) (bool, error)
	SlugLang() (string, error)
}
