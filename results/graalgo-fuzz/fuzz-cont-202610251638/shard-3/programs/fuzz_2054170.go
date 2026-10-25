package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 uint16, p1 uint32, p2 uint64) int8 {
    if (189 < 0) {
        switch "\n\t" {
        case "":
            var v5 uint64 = (p2 ^ (p2 - (p2 << ((598) & 63))))
            _ = v5
        case "foo bar":
            for fi6 := 0; fi6 < 19; fi6++ {
                var v7 uint16 = (p0 + 1)
                var v8 uint = 2
                _ = fi6
                _ = v7
                _ = v8
            }
        default:
            p2 -= (^(p2))
        }
    }
    p1 += 1
    _ = p0
    _ = p1
    _ = p2
    return 59
}
func genFn1(p0 int64, p1 int64) int32 {
    var v9 float64 = (3.14159 + (0.1 / (1.5 + 3.14159)))
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -2147483648
}

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]int)
    v3 := make(map[uint32]int)
    v4 := make(map[int]string)
    v11 := GenS1{F0: 24, F1: 890}
    v12 := GenS1{F0: -128, F1: 1}
    fmt.Println(v1)
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := v4[42]
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int = fi13
            if false {
                var v17 int8 = -110
                if (false && false) {
                    v18 := []int{v16, (fi13 - int(v1)), fi15}
                    fmt.Println(fi13)
                    v16 = 0
                    fmt.Println(fi13)
                    _ = v18
                }
                _ = v17
            } else {
                v14 = "hello"
                var v19 uint = 42
                v20 := GenS1{F0: 1, F1: -156}
                _ = v19
                _ = v20
            }
            v21 := v2[4294967295]
            _ = fi15
            _ = v16
            _ = v21
        }
        _ = fi13
        _ = v14
    }
    v22 := GenS1{F0: -12, F1: int64(v1)}
    if (0 <= -1) {
        if ("a" != "") {
            var v23 uint8 = 255
            v2[-998] = 127
            var v24 bool = true
            _ = v23
            _ = v24
        } else {
            for fi25 := 0; fi25 < 1; fi25++ {
                var v26 int = (-(fi25))
                _ = fi25
                _ = v26
            }
            fmt.Println(v1)
        }
        var v27 uint64 = 4294967295
        for fi28 := 0; fi28 < 18; fi28++ {
            v29 := GenS1{F0: -1, F1: -180}
            _ = fi28
            _ = v29
        }
        _ = v27
    } else {
        v30 := GenS1{F0: -47, F1: 194}
        var v31 bool = false
        switch -2147483648 {
        case 32767:
            if (!v31) {
                v32 := []int{}
                if (128 < 255) {
                    fmt.Println(v31, v1)
                    var v33 float32 = float32(1.5)
                    _ = v33
                }
                _ = v32
            } else {
                v34 := []int{-365, 42}
                _ = v34
            }
        }
        _ = v30
        _ = v31
    }
    v35 := []int{}
    for fi36 := 0; fi36 < 20; fi36++ {
        var v37 int32 = 106
        _ = fi36
        _ = v37
    }
    v1 += 75
    switch 127 {
    case 4294967295:
        v1 = (v1 | (^(v1)))
    case -2147483648:
        var v38 bool = true
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v35
    _ = v4
}
