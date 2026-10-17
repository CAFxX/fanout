package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    var v3 float64 = ((3.14159 * 1e10) + (1.5 * (1e10 * -2.25)))
    var v4 int = 4294967295
    var v5 float64 = 1.5
    switch "foo bar" {
    case "a":
        if ((!true) && true) {
            switch "hello" {
            case "hello":
                var v6 int16 = -596
                _ = v6
            case "foo bar":
                if ((v4 - (v4 - (v4 - -260))) <= 980) {
                    if (true && (-1 >= 0)) {
                        v5 -= v3
                        fmt.Println(v5, v5, v3)
                        v7 := GenS1{F0: false, F1: 40}
                        _ = v7
                    } else {
                        _ = len(v1)
                        for fi8 := 0; fi8 < 3; fi8++ {
                            var v9 float32 = float32(v5)
                            _ = fi8
                            _ = v9
                        }
                    }
                    fmt.Println(v5, v3, v4)
                } else {
                    v5 = v3
                    var v10 float64 = v5
                    _ = v10
                }
            }
            var v11 float64 = (0.0 / v3)
            v11 -= (v5 / (v5 / ((v3 / -2.25) / v3)))
            _ = v11
        }
    case "hello":
        v4 += v4
    }
    var v12 uint32 = 2147483648
    switch "hello" {
    case "a":
        v3 = (((((((1.5 + (((float64(1e10) - v5) + (-2.25 * v3)) - v5)) + 1.5) * float64((0.1 / float32(0.1)))) / float64((-2.25 - float32((v5 * v3))))) * v5) * ((float64(0.1) / ((0.1 + 3.14159) - v3)) + v5)) + v5)
    default:
        v3 = v5
    }
    for fi13 := 0; fi13 < 11; fi13++ {
        v12 -= v12
        v14 := GenS1{F0: (!(true && false)), F1: 1}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        v5 = (float64((1e10 - 0.1)) / (v3 + (v3 * v3)))
        delete(v2, "")
        _ = fi15
    }
    var v16 int16 = 0
    for fi17 := 0; fi17 < 18; fi17++ {
        v3 -= float64((1.5 / (1e10 * float32(v3))))
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int64 = -357
            fmt.Println(v5, v3)
            var v20 int16 = (v16 - (-(v16)))
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v3 = (float64((1e10 / -2.25)) - (float64(1e10) / (v5 + v5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
