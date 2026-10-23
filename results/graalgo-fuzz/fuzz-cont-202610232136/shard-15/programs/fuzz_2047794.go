package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 string) int {
    if false {
        v4 := GenS1{F0: 42, F1: 786}
        _ = v4
    }
    var v5 float32 = 1e10
    _ = p0
    _ = v5
    return 896
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    v3 := make(map[string]bool)
    if true {
        v1 += (v1 - 1)
        var v6 int16 = -1
        _ = v6
    } else {
        var v7 uint = 18446744073709551615
        v1 -= v1
        switch ((("foo bar" + "\n\t") + "hello") + "hello") {
        case "":
            switch -2147483648 {
            case -1:
                v7 -= (v7 ^ (^(v7)))
            case -9223372036854775808:
                _ = len(v2)
            default:
                v1 += (v1 - v1)
            }
        case "\n\t":
            var v8 int16 = -409
            _ = v8
        default:
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS2{F0: 42}
                _ = fi9
                _ = v10
            }
        }
        _ = v7
    }
    var v11 int16 = -732
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 uint16 = 255
        var v14 uint64 = 42
        v1 += (v1 >> ((177) & 63))
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = (-(v1))
    v1 = (v1 >> ((0) & 63))
    v1 -= -2147483648
    v1 -= (-(v1))
    switch 1 {
    case -265:
        var v15 uint32 = 65535
        _ = v15
    case 0:
        v11 = (v11 * v11)
    default:
        var v16 float32 = 1.5
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
