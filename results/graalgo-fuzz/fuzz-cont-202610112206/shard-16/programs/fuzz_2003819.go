package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }

func main() {
    var v1 int8 = 41
    var v2 int32 = 94
    var v3 uint8 = 2
    var v4 uint = 0
    var v5 uint32 = 2
    if true {
        var v6 float32 = -2.25
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := []int{fi7, (fi7 + fi7)}
            v9 := GenS1{F0: (fi7 | (fi7 ^ (fi7 / ((fi7) | 1)))), F1: v5}
            var v10 string = "\n\t"
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        _ = v6
    }
    switch 42 {
    case 551:
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: fi11, F1: (v5 % ((v5) | 1))}
            var v13 bool = (!((0.0 >= (((((-2.25 / (((1.5 - (float32(0.1) / -2.25)) * 1.5) - ((0.1 + 1.5) * (float32((float64(1e10) / 1e10)) / (float32(-2.25) / -2.25))))) - 0.0) * (3.14159 + float32(1e10))) * 1e10) + float32(1e10))) || false))
            switch (fi11 | -9223372036854775808) {
            case 2147483647:
                v14 := GenS1{F0: (fi11 >> ((v4) & 63)), F1: v5}
                _ = v14
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    case 0:
        v15 := []int{-664, 32767}
        _ = v15
    default:
        var v16 int = 9223372036854775807
        _ = v16
    }
    var v17 float32 = float32((((((float64(1e10) + 1e10) * 0.0) - ((0.0 + 0.1) / 1.5)) * float64(0.0)) / 0.1))
    v17 -= v17
    var v18 float32 = ((v17 - v17) + float32(1e10))
    var v19 uint64 = 18446744073709551615
    switch 32768 {
    case 9223372036854775807:
        var v20 int16 = -32768
        _ = v20
    default:
        var v21 float64 = 0.0
        _ = v21
    }
    v22 := []int{4294967295, 128, 128, 42}
    for fi23 := 0; fi23 < 20; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            switch "" {
            case "":
                switch "" {
                case "foo bar":
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                case "hello":
                    v25 := GenS1{F0: fi23, F1: v5}
                    _ = v25
                case "a":
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                }
            case "foo bar":
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v1)
                case "a":
                    fmt.Println(fi23)
                }
            case "\n\t":
                _ = len(v22)
            }
            v4 = uint(v5)
            _ = fi24
        }
        _ = fi23
    }
    v5 -= v5
    _ = len(v22)
    v26 := GenS2{F0: 255, F1: 153, F2: (v4 ^ (v4 >> ((v4) & 63)))}
    v17 += ((float32((0.0 - -2.25)) - (v18 + v17)) + v18)
    for fi27 := 0; fi27 < 16; fi27++ {
        v17 += -2.25
        fmt.Println(v19, v17, v2)
        v18 = 1.5
        _ = fi27
    }
    var v28 int = -2147483648
    var v29 uint = (v4 % (((^(v4))) | 1))
    var v30 int16 = -402
    v5 = (^(v5))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
