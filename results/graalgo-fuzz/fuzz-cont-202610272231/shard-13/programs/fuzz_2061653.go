package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int16 }

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{1}
    var v3 uint32 = 948
    var v4 string = "foo bar"
    var v5 int = -1
    var v6 bool = true
    v7 := GenS1{F0: 42}
    var v8 uint8 = 160
    v1 = (v1 / v1)
    fmt.Println(v8, v3)
    v9 := GenS2{F0: 0, F1: 167}
    var v10 uint8 = v8
    fmt.Println(v6, v5, v3)
    v6 = true
    v10 = uint8(v3)
    v8 -= (v10 * v8)
    if ((v4 >= ((v4 + ((v4 + "") + ((v4 + v4) + (v4 + v4)))) + ((v4 + v4) + v4))) || false) {
        for fi11 := 0; fi11 < 9; fi11++ {
            v12 := GenS1{F0: 32767}
            var v13 uint16 = 1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    if v6 {
        if false {
            v14 := GenS3{F0: v4, F1: 434}
            switch (-(v5)) {
            case 535:
                _ = len(v2)
            case -798:
                switch v5 {
                case 127:
                    v4 = (v4 + v4)
                case 522:
                    fmt.Println(v4)
                case 1:
                    switch "hello" {
                    case "\n\t":
                        v8 = uint8(v3)
                    default:
                        v15 := GenS4{F0: (^(v3)), F1: -32768}
                        _ = v15
                    }
                default:
                    var v16 bool = v6
                    _ = v16
                }
            default:
                v17 := GenS2{F0: v10, F1: 491}
                _ = v17
            }
            v18 := GenS1{F0: 0}
            _ = v14
            _ = v18
        }
        var v19 float32 = (float32(v1) * -2.25)
        _ = v19
    }
    var v20 int16 = 0
    v8 = v10
    switch v4 {
    case "foo bar":
        fmt.Println(v4, v1)
    default:
        fmt.Println(v8)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
