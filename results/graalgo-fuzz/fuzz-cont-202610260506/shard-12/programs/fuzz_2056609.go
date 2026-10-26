package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint = 0
    var v3 uint32 = 2147483648
    var v4 uint16 = 872
    var v5 bool = (-58 < 13)
    fmt.Println(v4, v2, v5)
    fmt.Println(v4)
    var v6 int64 = 0
    switch 9223372036854775807 {
    case 533:
        v3 -= v3
    case 127:
        fmt.Println(v2, v3, v3)
    default:
        var v7 bool = v5
        _ = v7
    }
    switch 42 {
    case -881:
        if v5 {
            var v8 uint64 = 1
            if v5 {
                v9 := GenS1{F0: 1, F1: (v8 % ((v8) | 1)), F2: 42}
                var v10 uint32 = (^(v3))
                _ = v10
                _ = v9
            }
            _ = v8
        }
    case 128:
        if v5 {
            v11 := GenS1{F0: (v3 ^ (v3 * (v3 - 65535))), F1: 0, F2: 1000}
            _ = v11
        }
    case -276:
        switch ("a" + "") {
        case "hello":
            v3 += 2
        case "a":
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := v1[v3]
                _ = len(v1)
                _ = fi12
                _ = v13
            }
        }
    }
    v1[(v3 % ((2147483648) | 1))] = uint32(v4)
    v14 := []int{}
    var v15 string = ("a" + "")
    v4 = (v4 / (((^(v4))) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
