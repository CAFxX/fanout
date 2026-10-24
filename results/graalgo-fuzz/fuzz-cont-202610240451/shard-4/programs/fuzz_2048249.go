package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint16 = 65535
    var v2 string = "\n\t"
    v3 := make(map[string]uint32)
    var v4 uint = 69
    v5 := []int{9223372036854775807, 920}
    for fi6 := 0; fi6 < 18; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v1 += (v1 * (v1 + (^(v1))))
            v8 := GenS1{F0: 100, F1: 42}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v1 = v1
    v1 = 255
    v1 = v1
    fmt.Println(v1)
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            v11 := GenS2{F0: (fi9 & (-(fi9))), F1: 0}
            if (1 < int16(fi10)) {
                var v12 uint64 = 0
                var v13 uint64 = v12
                var v14 int8 = -19
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch -9223372036854775808 {
    case 0:
        v2 = (v2 + v2)
    case 32768:
        v1 += (^(v1))
    case -635:
        v15 := v3["a"]
        _ = v15
    default:
        v2 = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
