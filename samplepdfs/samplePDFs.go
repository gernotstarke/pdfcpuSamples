// Package samplepdfs contain constant declarations for several sample PDF files,
// plus their (relative) path information, so that other test cases can be DRY.
package samplepdfs

// private constant

// public constants

// DirPrefix is the relative path to the directory containing samples
const DirPrefix = "../samplepdfs/"

// OnePageFile has exactly one page
const OnePageFile = DirPrefix + "sample-A4-portrait-1pg.pdf"

// ThreePageFile has exactly three pages
const ThreePageFile = DirPrefix + "sample-A4-portrait-3pgs.pdf"

// DisguisedMarkdownFile is a file with PDF extension, but markdown content
const DisguisedMarkdownFile = DirPrefix + "md-disguised-as.pdf"
