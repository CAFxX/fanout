package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 bool }
func genFn0(p0 uint8) int64 {
    var v5 int64 = 9223372036854775807
    if (v5 <= v5) {
        if (true || true) {
            fmt.Println(v5, p0)
            v5 = (v5 ^ v5)
        }
        p0 = p0
    }
    _ = p0
    _ = v5
    return v5
}

func main() {
    var v1 float64 = 3.14159
    var v2 float64 = 3.14159
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2147483648, 421, 453, 2147483648}
    fmt.Println(v2, v2)
    v6 := GenS1{F0: 92, F1: ("foo bar" + ""), F2: (4294967295 < 4294967295)}
    switch 942 {
    case 32768:
        v7 := GenS1{F0: 217, F1: "", F2: (false || false)}
        _ = v7
    case -2147483648:
        v8 := GenS1{F0: 1, F1: ("\n\t" + (("foo bar" + "hello") + (("" + "hello") + ""))), F2: (357 > 357)}
        _ = v8
    default:
        for fi9 := 0; fi9 < 20; fi9++ {
            v2 += v2
            _ = fi9
        }
    }
    var v10 int32 = 2147483647
    fmt.Println(v10, v10)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint8 = 127
        _ = fi11
        _ = v12
    }
    v10 = v10
    switch ("\n\t" + "a") {
    case "hello":
        v13 := GenS2{F0: true}
        _ = v13
    case "a":
        var v14 int16 = -1
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
