package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0(p0 int32, p1 uint16, p2 bool) uint {
    var v4 uint32 = 730
    v5 := []int{}
    if (0 == 42) {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 int64 = int64(v4)
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            p2 = p2
            p1 += (p1 >> ((65535) & 63))
            switch -241 {
            case 9223372036854775807:
                v4 = 987
            }
            _ = fi8
        }
    }
    switch "a" {
    case "":
        var v9 uint32 = (v4 - 2)
        _ = v9
    case "a":
        switch "a" {
        case "hello":
            var v10 int8 = 100
            _ = v10
        case "foo bar":
            var v11 float32 = -2.25
            _ = v11
        default:
            v12 := GenS1{F0: 582, F1: 9223372036854775807, F2: (1 <= 9223372036854775807)}
            _ = v12
        }
    case "hello":
        var v13 int32 = (^(p0))
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 string, p1 int8) uint16 {
    v14 := GenS1{F0: 255, F1: -1, F2: true}
    v15 := GenS2{F0: 255, F1: p0, F2: 42}
    var v16 uint32 = 2
    var v17 int = 4294967295
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 464
    var v3 uint16 = 65535
    v18 := GenS2{F0: (v3 | (v3 * (v3 % ((v3) | 1)))), F1: "\n\t", F2: 274}
    if (1e10 != float64((3.14159 * 3.14159))) {
        v19 := []int{9223372036854775807, -19, -365, 32768}
        var v20 uint8 = 1
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: (v3 << ((76) & 63)), F1: "a", F2: 4294967295}
    if (9223372036854775807 < int64(v2)) {
        v1[(v2 << ((18446744073709551615) & 63))] = ("hello" + "a")
    }
    if false {
        v3 += 2
        v22 := GenS2{F0: v3, F1: (("\n\t" + ("foo bar" + ((("hello" + "") + ("foo bar" + "a")) + "foo bar"))) + "hello"), F2: v2}
        _ = v22
    } else {
        v3 += 217
    }
    v2 += (v2 + v2)
    v23 := []int{}
    var v24 uint16 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
}
