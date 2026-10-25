package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }

func main() {
    var v1 int64 = 0
    v2 := []uint32{1, 2147483648}
    var v3 uint8 = 89
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 0
    switch -9223372036854775808 {
    case -1:
        v3 = 154
    default:
        var v6 bool = true
        _ = v6
    }
    v7 := GenS1{F0: 2, F1: 0}
    if (2 == 0) {
        v3 += (v3 & (v5 % (((^(v5))) | 1)))
        switch "hello" {
        case "hello":
            v5 = (v3 / ((v3) | 1))
        case "a":
            v8 := GenS1{F0: 3, F1: 1}
            _ = v8
        case "\n\t":
            var v9 uint = 447
            _ = v9
        default:
            for fi10 := 0; fi10 < 9; fi10++ {
                if (-90 == 42) {
                    delete(v4, 2147483648)
                    for fi11 := 0; fi11 < 1; fi11++ {
                        _ = len(v2)
                        var v12 uint16 = 2
                        _ = fi11
                        _ = v12
                    }
                }
                _ = fi10
            }
        }
        var v13 uint16 = 255
        _ = v13
    } else {
        var v14 int16 = -291
        v15 := GenS1{F0: 186, F1: 763}
        _ = v14
        _ = v15
    }
    v5 = (v3 + (^(v5)))
    fmt.Println(v1)
    var v16 int16 = int16(v3)
    v16 = v16
    v3 = (v5 * v3)
    v1 -= v1
    if (!(1 < 54)) {
        switch 2147483647 {
        case 42:
            var v17 string = ("\n\t" + "foo bar")
            _ = v17
        case 9223372036854775807:
            var v18 int8 = -128
            _ = v18
        default:
            var v19 int8 = 42
            _ = v19
        }
        var v20 string = "hello"
        _ = v20
    }
    var v21 int = 32768
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
