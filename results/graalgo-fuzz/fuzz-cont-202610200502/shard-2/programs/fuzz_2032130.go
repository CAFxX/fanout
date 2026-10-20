package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint }
type GenS3 struct { F0 bool; F1 bool; F2 bool }

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 uint32 = 78
    var v4 uint32 = (v3 * v3)
    v1 -= (v1 >> ((1) & 63))
    if (42 >= 1) {
        var v5 uint = 2
        _ = v5
    }
    v3 = v3
    v4 = (v4 << ((2) & 63))
    switch "hello" {
    case "hello":
        v4 -= v3
    case "\n\t":
        var v6 string = "hello"
        _ = v6
    case "a":
        if true {
            if (258 == -110) {
                for fi7 := 0; fi7 < 11; fi7++ {
                    if false {
                        v1 = int32(fi7)
                        var v8 uint32 = 2
                        v9 := GenS1{F0: fi7, F1: 543, F2: -1}
                        _ = v8
                        _ = v9
                    }
                    v4 += v3
                    _ = fi7
                }
                switch 0 {
                case 1:
                    fmt.Println(v1)
                default:
                    v10 := GenS1{F0: 623, F1: 42, F2: -32768}
                    _ = v10
                }
            }
        }
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := GenS2{F0: 128, F1: 135, F2: 1}
        if false {
            var v13 bool = ((((3.14159 - 1e10) * float64(float32(float64(1e10)))) != ((float64(3.14159) * (1.5 + (3.14159 * (0.1 * -2.25)))) - (float64(0.0) - (0.1 * -2.25)))) || true)
            fmt.Println(v1)
            _ = v13
        } else {
            v4 -= (v3 ^ 0)
            v14 := []int{fi11, (fi11 | (fi11 * (-(fi11))))}
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 int8 = 42
    v16 := []int{-2147483648, 127, -505, -1}
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 6; fi18++ {
            if true {
                if (true && (18446744073709551615 >= 42)) {
                    v19 := GenS3{F0: true, F1: (false || true), F2: true}
                    v3 = 170
                    _ = v19
                }
                var v20 int16 = 1
                _ = v20
            }
            _ = fi18
        }
        fmt.Println(v15, v15)
        var v21 int16 = -818
        _ = fi17
        _ = v21
    }
    for fi22 := 0; fi22 < 17; fi22++ {
        v23 := GenS2{F0: 2, F1: 255, F2: 0}
        _ = fi22
        _ = v23
    }
    v4 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
