package lengthconv

func CToM(c Centimiter) Meter { return Meter(c / 100) }

func MToC(m Meter) Centimiter { return Centimiter(m * 100) }
