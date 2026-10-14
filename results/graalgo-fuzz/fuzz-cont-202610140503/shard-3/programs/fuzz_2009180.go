package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 int8 }

func main() {
    var v1 bool = (2 < 65535)
    var v2 uint = 457
    var v3 int16 = 1
    v4 := []int{9223372036854775807, 32767}
    var v5 float64 = 0.0
    for fi6 := 0; fi6 < 6; fi6++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
    }
    _ = len(v4)
    v3 = (v3 ^ -1)
    var v7 uint16 = 32767
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint64 = 784
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    if v1 {
        var v10 int = -535
        switch int(v7) {
        case -2147483648:
            var v11 string = "foo bar"
            _ = v11
        case 0:
            for fi12 := 0; fi12 < 14; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    var v14 float32 = float32((0.0 * 0.1))
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        case -1:
            v15 := GenS1{F0: (v3 / ((1000) | 1)), F1: 18446744073709551615, F2: 72}
            _ = v15
        default:
            fmt.Println(v3)
        }
        fmt.Println(v10)
        _ = v10
    }
    var v16 uint8 = 30
    for fi17 := 0; fi17 < 15; fi17++ {
        var v18 uint64 = 4294967295
        for fi19 := 0; fi19 < 2; fi19++ {
            v20 := GenS1{F0: (v3 * (v3 % ((471) | 1))), F1: (v18 | 340), F2: 6}
            v5 = v5
            _ = len(v4)
            v18 = v18
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
