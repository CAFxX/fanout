package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 int16; F1 int64; F2 uint8 }
func genFn0(p0 int32, p1 int8, p2 bool) int32 {
    v6 := []int{273, 275}
    switch ("foo bar" + (((("\n\t" + "\n\t") + "") + "a") + "\n\t")) {
    case "":
        p0 = (-(p0))
    case "\n\t":
        var v7 int = -87
        _ = v7
    }
    if true {
        switch ("foo bar" + "a") {
        case "foo bar":
            v8 := GenS1{F0: -912, F1: 1, F2: 65535}
            _ = v8
        default:
            var v9 uint16 = 736
            _ = v9
        }
        var v10 int = 141
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 161
}
func genFn1(p0 uint64) string {
    p0 -= 309
    v11 := GenS1{F0: -1, F1: 42, F2: 65535}
    p0 = p0
    p0 += (p0 - p0)
    var v12 bool = (!(true && false))
    _ = p0
    _ = v11
    _ = v12
    return "a"
}

func main() {
    v1 := []int{42, 4294967295, 2147483647}
    v2 := make(map[int]int)
    var v3 int32 = -2147483648
    var v4 int32 = (v3 + (v3 - v3))
    var v5 int8 = 127
    var v13 uint32 = 2147483647
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float32 = 0.1
        var v16 uint64 = 42
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v3, v13, v3)
    if true {
        v17 := GenS1{F0: -1, F1: 978, F2: 504}
        _ = v17
    } else {
        v5 += (v5 & (v5 + (v5 >> ((963) & 63))))
        switch ("foo bar" + (("\n\t" + "\n\t") + "foo bar")) {
        case "foo bar":
            v3 += (v3 << ((366) & 63))
        }
    }
    v18 := GenS1{F0: 1, F1: -1, F2: 396}
    v19 := GenS1{F0: -1, F1: 127, F2: 42}
    var v20 uint8 = 255
    _ = len(v1)
    var v21 int32 = v3
    v22 := GenS2{F0: 0, F1: 623, F2: v20}
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    v21 = genFn0(v4, v5, true)
    v13 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
