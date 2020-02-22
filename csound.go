package csound

// File provides a DSL for generating csound CSD (aka "unified file format") files.
type File struct {
	Synthesizer Synthesizer `xml:"CsoundSynthesizer"`
}

// Synthesizer represents a csound synthesizer.
type Synthesizer struct {
	Options Options `xml:"CsOptions"`
}

// Options represents csound options.
type Options struct {
}
