package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int; F1 uint16 }
func genFn0(p0 bool) int {
    var v7 int32 = 0
    if p0 {
        v8 := []int{752, 32767}
        if (0 != -879) {
            p0 = (2 == 729)
        } else {
            v9 := GenS1{F0: 1000, F1: 34, F2: 100}
            var v10 uint64 = 754
            _ = v10
            _ = v9
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            p0 = true
            if p0 {
                v12 := GenS2{F0: 42, F1: fi11}
                _ = v12
            }
            var v13 uint64 = 4294967295
            _ = fi11
            _ = v13
        }
        _ = v8
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        var v15 int8 = -112
        var v16 uint = 18446744073709551615
        var v17 int32 = v7
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v7
    return -508
}

func main() {
    v1 := []uint32{53, 0, 2}
    var v2 uint16 = 2
    var v3 uint = 4294967295
    var v4 int = -9223372036854775808
    var v5 int = 2147483647
    var v6 int8 = -113
    v18 := GenS3{F0: 0, F1: 65535}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 = int8(v2)
    switch (("hello" + "a") + "foo bar") {
    case "foo bar":
        var v19 int32 = 70
        _ = v19
    case "a":
        if ((int64(v5) == -1) && true) {
            if len(v1) > 0 {
                v1[0] = 2
            }
        } else {
            switch 0 {
            case -1:
                v20 := GenS4{F0: -9223372036854775808, F1: 172}
                _ = v20
            }
        }
    default:
        var v21 uint8 = 25
        _ = v21
    }
    var v22 int = v5
    var v23 int32 = 2147483647
    _ = len(v1)
    fmt.Println(v6)
    v24 := GenS2{F0: 646, F1: -1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
