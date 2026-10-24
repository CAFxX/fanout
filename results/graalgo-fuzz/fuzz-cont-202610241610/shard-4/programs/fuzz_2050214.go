package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 float64 = 1e10
    var v2 int = 9223372036854775807
    v3 := []int{v2, (v2 ^ (-(v2)))}
    var v4 uint32 = 42
    v3 = append(v3, (v2 * v2))
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            v2 = (^(fi5))
            _ = fi5
        }
    } else {
        fmt.Println(v4)
        for fi6 := 0; fi6 < 8; fi6++ {
            v2 -= v2
            v4 += (v4 + v4)
            _ = len(v3)
            _ = fi6
        }
        var v7 int32 = 2147483647
        _ = v7
    }
    var v8 int16 = -32768
    v9 := GenS1{F0: (-(v8)), F1: 1, F2: 2}
    _ = len(v3)
    for fi10 := 0; fi10 < 6; fi10++ {
        v1 = v1
        _ = fi10
    }
    switch v2 {
    case -541:
        for fi11 := 0; fi11 < 9; fi11++ {
            if (1 != 510) {
                var v12 uint64 = 42
                v12 -= v12
                _ = v12
            }
            var v13 string = ""
            if true {
                switch 2147483647 {
                case -694:
                    v1 -= v1
                case 42:
                    var v14 uint64 = 437
                    _ = v14
                case 127:
                    var v15 bool = true
                    _ = v15
                }
            }
            _ = fi11
            _ = v13
        }
    default:
        v8 -= v8
    }
    if false {
        v16 := GenS2{F0: 128}
        var v17 uint64 = 46
        var v18 float32 = (0.0 * -2.25)
        _ = v16
        _ = v17
        _ = v18
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 float64 = v1
            _ = fi19
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
