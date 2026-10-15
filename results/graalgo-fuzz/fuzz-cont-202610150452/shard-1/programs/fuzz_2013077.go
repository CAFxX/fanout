package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint) int8 {
    p0 = (^(p0))
    if (1 <= 2) {
        var v5 int8 = 25
        _ = v5
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            var v6 int64 = -853
            _ = v6
        }
    case "foo bar":
        p0 += (^(p0))
    }
    p0 += (p0 >> ((p0) & 63))
    _ = p0
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = ((("a" + "\n\t") + "foo bar") + "foo bar")
    var v3 uint8 = 42
    v4 := make(map[int]bool)
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int64 = 9223372036854775807
        _ = fi7
        _ = v8
    }
    var v9 float32 = 0.1
    v3 = (v3 << ((779) & 63))
    for fi10 := 0; fi10 < 12; fi10++ {
        v11 := []int{(^(fi10))}
        switch fi10 {
        case -2147483648:
            switch "a" {
            case "foo bar":
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: 497, F1: 65535}
                    _ = fi12
                    _ = v13
                }
            case "hello":
                v2 = "hello"
            case "a":
                v2 = "\n\t"
            default:
                var v14 uint = 255
                _ = v14
            }
        case -586:
            fmt.Println(v9, v2)
        default:
            v9 -= float32(1.5)
        }
        _ = fi10
        _ = v11
    }
    v9 -= 0.0
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            v16 := []int{128, (fi15 - (-(fi15))), (fi15 / ((fi15) | 1))}
            for fi17 := 0; fi17 < 15; fi17++ {
                switch v2 {
                case "hello":
                    switch v2 {
                    case "":
                        var v18 float32 = float32(1.5)
                        _ = v18
                    case "a":
                        var v19 float64 = 3.14159
                        _ = v19
                    case "foo bar":
                        _ = len(v4)
                    }
                case "foo bar":
                    v9 -= v9
                }
                v9 -= v9
                v2 = v2
                _ = fi17
            }
            _ = v16
        }
        v1[(fi15 % (((fi15 * fi15)) | 1))] = 2
        _ = fi15
    }
    var v20 uint8 = (v3 + (^(v3)))
    v3 = v3
    if ((-2.25 * 1e10) > 0.0) {
        if false {
            v21 := v1[32768]
            _ = v21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
