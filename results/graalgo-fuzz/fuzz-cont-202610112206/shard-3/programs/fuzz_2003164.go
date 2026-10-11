package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint16 }
type GenS2 struct { F0 int16; F1 int8 }
func genFn0(p0 uint8) uint8 {
    var v6 uint32 = 2147483648
    v6 = v6
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 uint16 = 32767
    var v2 int = 9223372036854775807
    var v3 bool = (!(0 < 42))
    v4 := []uint32{}
    var v5 int = (v2 ^ (v2 + v2))
    for fi7 := 0; fi7 < 2; fi7++ {
        _ = len(v4)
        _ = fi7
    }
    v8 := GenS1{F0: ("foo bar" + "hello"), F1: v5, F2: (^(v1))}
    fmt.Println(v1)
    v1 = (v1 ^ (v1 + v1))
    for fi9 := 0; fi9 < 7; fi9++ {
        v2 -= -2147483648
        _ = fi9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        if true {
            v1 -= v1
            if false {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        }
        v1 = (^(v1))
        v4 = append(v4, 336)
        _ = fi10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if false {
            for fi12 := 0; fi12 < 5; fi12++ {
                if (true && (float32(0.0) < (-2.25 * 3.14159))) {
                    switch 42 {
                    case 127:
                        if v3 {
                            var v13 uint32 = 2147483648
                            if ((42 <= -3) || v3) {
                                v1 = (v1 + 2)
                                var v14 string = ""
                                if len(v4) > 0 {
                                    v4[0] = v13
                                }
                                _ = v14
                            } else {
                                v4 = append(v4, 2)
                                v3 = (!v3)
                            }
                            _ = v13
                        }
                        fmt.Println(fi11)
                    case 9223372036854775807:
                        var v15 uint = 0
                        _ = v15
                    case 1:
                        v5 -= v2
                    }
                } else {
                    fmt.Println(fi12, v3, v2)
                    if (((3.14159 / 0.1) + float32((1.5 / 3.14159))) == 0.0) {
                        v2 -= (fi12 / ((v2) | 1))
                    } else {
                        v3 = v3
                    }
                }
                var v16 float64 = (-2.25 / float64(-2.25))
                var v17 int16 = 1000
                _ = fi12
                _ = v16
                _ = v17
            }
        } else {
            fmt.Println(v5)
        }
        _ = fi11
    }
    v18 := GenS2{F0: 1, F1: 42}
    for fi19 := 0; fi19 < 13; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS2{F0: 0, F1: 127}
            if v3 {
                v22 := GenS2{F0: 537, F1: -128}
                _ = v22
            } else {
                switch (fi19 | (v5 | fi19)) {
                case 32768:
                    v5 += -2147483648
                default:
                    if true {
                        fmt.Println(fi20, v5)
                        v23 := GenS2{F0: 0, F1: -128}
                        _ = v23
                    }
                }
                var v24 uint8 = 0
                var v25 uint16 = v1
                _ = v24
                _ = v25
            }
            var v26 int = (v5 | (v5 | (v2 | v5)))
            _ = fi20
            _ = v21
            _ = v26
        }
        _ = fi19
    }
    var v27 int16 = -1
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
