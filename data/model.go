package data

import "io"

type Decodable interface {
	FromJSON(reader io.Reader) error
}

type Encodable interface {
	ToJSON(writer io.Writer) error
}

type Codable interface {
	FromJSON(reader io.Reader) error
	ToJSON(writer io.Writer) error
}
