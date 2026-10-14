package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 int8 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 string; F1 int8; F2 int64 }
func genFn0(p0 bool, p1 int8, p2 string) bool {
    fmt.Println(p0)
    switch 9223372036854775807 {
    case 1:
        for fi4 := 0; fi4 < 14; fi4++ {
            v5 := GenS1{F0: 1, F1: 80, F2: p1}
            _ = fi4
            _ = v5
        }
    }
    p2 = p2
    fmt.Println(p1)
    v6 := GenS2{F0: 2, F1: (p2 + p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((p0 || true) && (false && ((!true) || ((!(4294967295 < 508)) || true))))
}

func main() {
    var v1 int32 = -354
    var v2 int = 128
    v3 := []uint32{2147483648, 1, 42, 2}
    v1 = v1
    v7 := GenS3{F0: "a", F1: 80, F2: 42}
    v2 = v2
    v8 := GenS1{F0: v1, F1: v2, F2: -88}
    switch (v2 & v2) {
    case 1:
        v9 := GenS3{F0: "\n\t", F1: 42, F2: 42}
        _ = v9
    case 4294967295:
        fmt.Println(v1, v1)
    case -9223372036854775808:
        if ((3.14159 * ((1e10 / 1e10) - (1.5 / -2.25))) != (1.5 / float64(0.1))) {
            var v10 string = "foo bar"
            var v11 int8 = 100
            var v12 uint16 = 42
            _ = v10
            _ = v11
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        var v14 float64 = ((1.5 * 0.0) - (((float64(0.0) + (3.14159 + (0.1 - 1e10))) + 3.14159) / 3.14159))
        _ = fi13
        _ = v14
    }
    var v15 string = "hello"
    v1 = 0
    var v16 uint64 = 90
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
