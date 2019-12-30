package lengthconv

import "fmt"

type Centimiter float64
type Meter float64

const (
	LichangHigh Meter=1.77
	SeaAltitude Meter=0
)

func (c Centimiter) String() string {return fmt.Sprintf("%g cm",c)}

func (m Meter) String() string {return fmt.Sprintf("%g m", m)}