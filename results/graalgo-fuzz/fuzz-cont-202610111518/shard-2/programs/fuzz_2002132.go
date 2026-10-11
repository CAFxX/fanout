package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
func genFn0() string {
    v7 := GenS1{F0: 0, F1: 0}
    v8 := GenS1{F0: -2147483648, F1: 1}
    var v9 uint8 = 2
    if (32767 != 0) {
        var v10 uint16 = 2
        v10 = v10
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return (("a" + "") + "\n\t")
}
func genFn1(p0 int8) uint64 {
    p0 -= (p0 >> ((497) & 63))
    p0 = 1
    var v11 int16 = 42
    var v12 string = "a"
    _ = p0
    _ = v11
    _ = v12
    return 0
}

func main() {
    var v1 int8 = 0
    v2 := []uint32{0}
    var v3 int = 845
    var v4 float32 = (3.14159 + 1e10)
    v5 := make(map[string]int)
    var v6 int32 = 0
    switch "" {
    case "hello":
        if true {
            fmt.Println(v4, v6, v4)
            var v13 string = "hello"
            _ = len(v2)
            _ = v13
        } else {
            var v14 int = v3
            v4 = (v4 * v4)
            v3 += -2147483648
            _ = v14
        }
    }
    if ((v6 ^ (v6 << ((42) & 63))) <= 1) {
        var v15 uint = 2
        var v16 uint64 = 1
        _ = v15
        _ = v16
    }
    var v17 float32 = 3.14159
    if false {
        var v18 int32 = 2147483647
        _ = v18
    } else {
        for fi19 := 0; fi19 < 7; fi19++ {
            v6 = (v6 ^ (-(v6)))
            var v20 uint32 = 2
            _ = fi19
            _ = v20
        }
        switch v3 {
        case 1:
            fmt.Println(v1)
        case 9223372036854775807:
            var v21 uint = 304
            _ = v21
        }
    }
    if true {
        var v22 uint64 = 0
        _ = v22
    }
    v6 = (^(v6))
    if (v1 == (v1 & -80)) {
        var v23 int8 = v1
        var v24 uint8 = 127
        _ = v23
        _ = v24
    }
    _ = len(v5)
    var v25 float64 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
