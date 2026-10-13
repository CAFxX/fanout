package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }

func main() {
    v1 := []int{32768, 127}
    var v2 uint32 = 2
    var v3 float64 = 1e10
    switch "a" {
    case "hello":
        _ = len(v1)
    }
    v4 := GenS1{F0: 255, F1: 2}
    if (false && true) {
        switch 32768 {
        case -259:
            v3 = ((float64(1.5) * v3) + (v3 * (v3 - float64(-2.25))))
        case 42:
            switch 9223372036854775807 {
            case 433:
                var v5 float64 = ((1.5 * 1e10) + 0.1)
                _ = v5
            case 927:
                v6 := GenS1{F0: 255, F1: 810}
                _ = v6
            case 32768:
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := GenS1{F0: 207, F1: 32767}
                    v2 = 2147483648
                    _ = fi7
                    _ = v8
                }
            default:
                fmt.Println(v2, v2)
            }
        }
        fmt.Println(v2, v2)
        v3 = v3
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            v1 = append(v1, 49)
            switch "\n\t" {
            case "foo bar":
                v2 -= v2
            case "a":
                var v11 uint8 = 128
                _ = v11
            case "":
                var v12 float32 = (0.1 - (3.14159 / float32(v3)))
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    v1 = append(v1, 461)
    _ = len(v1)
    _ = len(v1)
    for fi13 := 0; fi13 < 16; fi13++ {
        v3 -= float64(3.14159)
        _ = len(v1)
        _ = fi13
    }
    var v14 uint64 = 179
    if ((((1 != v2) || true) && (!(4294967295 == 42))) || false) {
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                v16 := GenS1{F0: 330, F1: 1}
                _ = v16
            }
            _ = fi15
        }
        var v17 int16 = -32768
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
