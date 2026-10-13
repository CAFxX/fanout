package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
func genFn0(p0 uint64, p1 uint16, p2 uint8) int {
    switch -85 {
    case 281:
        p2 = p2
    case 127:
        v5 := GenS1{F0: p0, F1: p2, F2: "hello"}
        _ = v5
    default:
        var v6 int8 = 42
        _ = v6
    }
    v7 := []int{1}
    _ = len(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 32768
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 208
    var v3 int32 = -1
    var v4 uint = 18446744073709551615
    v8 := []int{32768}
    if false {
        var v9 bool = (!true)
        var v10 float32 = 1e10
        if (-222 != -989) {
            v2 = (v2 % (((v2 % ((v2) | 1))) | 1))
            fmt.Println(v3, v2, v3)
        }
        _ = v10
        _ = v9
    } else {
        v3 -= v3
    }
    v2 = v2
    v11 := GenS1{F0: 4294967295, F1: 127, F2: "a"}
    v4 = (^(v4))
    v12 := v1[""]
    if (140 >= 15) {
        var v13 uint16 = 32767
        _ = v13
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if (2 < 1) {
            var v15 bool = true
            _ = v15
        } else {
            v16 := GenS1{F0: 42, F1: 128, F2: (((v12 + (v12 + ("a" + "foo bar"))) + "hello") + "foo bar")}
            if len(v8) > 0 {
                v8[0] = (-(fi14))
            }
            _ = v16
        }
        v17 := GenS1{F0: 0, F1: 220, F2: (v12 + v12)}
        if true {
            v3 += -2147483648
        } else {
            var v18 int64 = 1
            _ = v18
        }
        _ = fi14
        _ = v17
    }
    fmt.Println(v3)
    if (42 == (v2 | v2)) {
        var v19 uint16 = 2
        _ = v19
    } else {
        fmt.Println(v3, v12)
        v3 -= v3
        v4 = 18446744073709551615
    }
    fmt.Println(v12)
    v20 := GenS1{F0: 1, F1: 1, F2: v12}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
