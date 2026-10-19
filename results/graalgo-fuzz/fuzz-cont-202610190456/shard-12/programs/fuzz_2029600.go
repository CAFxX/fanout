package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint8; F2 string }
func genFn0() uint32 {
    var v4 uint32 = 350
    v4 -= (v4 << ((255) & 63))
    _ = v4
    return v4
}
func genFn1(p0 int8, p1 int64, p2 uint8) int8 {
    switch -1 {
    case 0:
        if (!(1000 == 42)) {
            var v5 uint = 42
            _ = v5
        }
    case 9223372036854775807:
        var v6 int32 = -975
        _ = v6
    case 128:
        v7 := GenS1{F0: -242, F1: p2, F2: true}
        _ = v7
    }
    fmt.Println(p2, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    v1 := []int{0}
    var v2 int = 128
    var v3 uint = 132
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = v2
        _ = fi8
        _ = v9
    }
    fmt.Println(v3)
    if (42 < 704) {
        v1 = append(v1, 1)
    }
    v1 = append(v1, 2147483647)
    for fi10 := 0; fi10 < 16; fi10++ {
        v3 -= (v3 & v3)
        switch "" {
        case "":
            v11 := GenS1{F0: fi10, F1: 0, F2: ((!false) || false)}
            _ = v11
        case "a":
            switch ("a" + "foo bar") {
            case "a":
                var v12 uint8 = 2
                _ = v12
            default:
                v2 += v2
            }
        default:
            v13 := GenS2{F0: 96, F1: 168}
            _ = v13
        }
        v14 := GenS3{F0: 0}
        _ = fi10
        _ = v14
    }
    fmt.Println(v2)
    var v15 int8 = 1
    v3 += v3
    v16 := GenS3{F0: -34}
    var v17 uint64 = 4294967295
    var v18 uint = v3
    v19 := GenS4{F0: (("" + ("" + "foo bar")) + "hello"), F1: 255, F2: "a"}
    v17 = v17
    var v20 int16 = 197
    v20 -= v20
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
