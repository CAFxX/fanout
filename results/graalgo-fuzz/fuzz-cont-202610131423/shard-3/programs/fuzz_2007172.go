package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint64) int {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint = 4294967295
        var v9 int64 = -189
        v9 = v9
        _ = fi7
        _ = v8
        _ = v9
    }
    if (1 == 42) {
        var v10 int32 = -2147483648
        var v11 uint64 = 2
        v12 := GenS1{F0: 255, F1: false, F2: true}
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    return 96
}
func genFn1(p0 string) string {
    p0 = "hello"
    p0 = ("" + "")
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint64 = 4294967295
        v15 := []int{(fi13 ^ fi13), (fi13 >> ((36) & 63))}
        v14 = v14
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    return p0
}

func main() {
    var v1 uint = 4294967295
    v2 := make(map[string]string)
    var v3 int16 = -32768
    var v4 uint8 = 42
    var v5 uint16 = 775
    var v6 int16 = v3
    v16 := GenS2{F0: -304, F1: v4}
    v6 -= v3
    v5 -= 255
    v1 -= (^(v1))
    fmt.Println(v5, v3)
    v6 = 32767
    var v17 int32 = 1
    fmt.Println(v5)
    v1 = 763
    delete(v2, "foo bar")
    v5 = v5
    if (65535 != 187) {
        if false {
            v4 = v4
            v18 := GenS1{F0: v4, F1: ("" >= ""), F2: true}
            v19 := GenS1{F0: (v4 + (^(v4))), F1: (true && false), F2: true}
            _ = v18
            _ = v19
        }
    } else {
        var v20 int = 998
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
