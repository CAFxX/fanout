package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 int16, p1 bool) uint8 {
    fmt.Println(p1)
    var v4 uint8 = 128
    _ = p0
    _ = p1
    _ = v4
    return v4
}

func main() {
    var v1 int8 = 43
    var v2 int = 4294967295
    v3 := make(map[uint32]string)
    for fi5 := 0; fi5 < 14; fi5++ {
        switch ("foo bar" + ("a" + (("" + ("hello" + ("foo bar" + "hello"))) + "\n\t"))) {
        case "\n\t":
            for fi6 := 0; fi6 < 20; fi6++ {
                switch "a" {
                case "hello":
                    var v7 uint64 = 18446744073709551615
                    _ = v7
                case "":
                    v1 = (v1 << ((18446744073709551615) & 63))
                }
                var v8 uint64 = 2
                _ = fi6
                _ = v8
            }
        case "":
            var v9 uint32 = 65535
            _ = v9
        case "hello":
            v10 := []int{508}
            _ = v10
        default:
            var v11 int32 = -445
            _ = v11
        }
        v1 -= (v1 - (v1 - (v1 >> ((1) & 63))))
        v1 = (^(v1))
        _ = fi5
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        v1 = (v1 >> ((1) & 63))
        _ = fi12
    }
    v13 := GenS1{F0: (v1 / ((-1) | 1)), F1: -73}
    switch 127 {
    case -281:
        var v14 uint8 = 2
        _ = v14
    }
    var v15 int64 = 1
    fmt.Println(v2, v2, v1)
    v16 := []int{}
    v17 := GenS2{F0: v1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
