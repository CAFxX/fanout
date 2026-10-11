package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint16) uint {
    for fi5 := 0; fi5 < 17; fi5++ {
        var v6 int16 = -194
        _ = fi5
        _ = v6
    }
    v7 := []int{-397, 2147483647}
    _ = p0
    _ = v7
    return 237
}

func main() {
    var v1 uint64 = 0
    var v2 int32 = -1
    v3 := make(map[uint32]string)
    var v4 int = 1
    var v8 uint16 = 0
    for fi9 := 0; fi9 < 6; fi9++ {
        switch "foo bar" {
        case "a":
            switch (((("a" + "") + ("" + "\n\t")) + "\n\t") + "") {
            case "foo bar":
                if ("a" > "hello") {
                    var v10 uint32 = 2147483647
                    _ = v10
                }
            default:
                v11 := v3[1]
                _ = v11
            }
        default:
            switch ("" + "") {
            case "hello":
                v12 := GenS1{F0: v8, F1: fi9, F2: fi9}
                _ = v12
            case "":
                fmt.Println(v8, v8, fi9)
            default:
                fmt.Println(fi9)
            }
        }
        var v13 uint = genFn0(v8)
        v4 = 128
        _ = fi9
        _ = v13
    }
    v14 := GenS2{F0: 42}
    fmt.Println(v1, v1, v2)
    v1 = 18446744073709551615
    switch "foo bar" {
    case "\n\t":
        switch (^(v4)) {
        case 4294967295:
            v4 += (v4 % (((v4 - (-(v4)))) | 1))
        case 128:
            var v15 uint32 = 2
            _ = v15
        case 2147483647:
            delete(v3, 2147483648)
        default:
            var v16 uint = 293
            _ = v16
        }
    default:
        switch v4 {
        case -9223372036854775808:
            var v17 int16 = 0
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        var v19 uint16 = (v8 >> ((42) & 63))
        var v20 bool = (!(uint16(v1) > (v8 >> ((42) & 63))))
        _ = fi18
        _ = v19
        _ = v20
    }
    var v21 float64 = ((3.14159 - -2.25) * (1.5 * 1e10))
    v22 := GenS2{F0: 0}
    var v23 int = (v4 >> ((33) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
}
