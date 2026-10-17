package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool, p1 int) bool {
    p0 = (0 < 100)
    switch "a" {
    case "hello":
        var v7 uint32 = 948
        _ = v7
    default:
        var v8 bool = p0
        _ = v8
    }
    var v9 uint8 = 255
    var v10 int16 = 135
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return true
}
func genFn1(p0 bool, p1 int64) int {
    v11 := []int{}
    var v12 int = 32768
    var v13 int16 = 699
    switch v12 {
    case 0:
        var v14 uint32 = 67
        _ = v14
    case 2147483647:
        if p0 {
            v12 = (v12 + 42)
        }
    }
    var v15 bool = (!(false && (p0 || (int(p1) == (v12 & (^(v12)))))))
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v15
    return v12
}

func main() {
    v1 := []uint32{0}
    var v2 bool = (0 > 636)
    v3 := make(map[int]string)
    var v4 int32 = -1
    v5 := []string{("" + ""), "hello"}
    var v6 float32 = -2.25
    if ((v4 - -339) < v4) {
        switch "" {
        case "a":
            var v16 int32 = (v4 & v4)
            _ = v16
        }
        switch "hello" {
        case "":
            var v17 bool = false
            _ = v17
        case "hello":
            delete(v3, -702)
        }
        switch 127 {
        case 32767:
            delete(v3, 1)
        case 42:
            var v18 int16 = 1000
            _ = v18
        case 4294967295:
            v4 -= (v4 & 1)
        default:
            fmt.Println(v4, v2)
        }
    }
    var v19 uint8 = 42
    for fi20 := 0; fi20 < 1; fi20++ {
        switch (fi20 << ((65535) & 63)) {
        case -2147483648:
            var v21 uint16 = 32767
            _ = v21
        }
        v22 := GenS1{F0: (fi20 + (fi20 / (((fi20 * (fi20 % (((fi20 - (fi20 / (((fi20 >> ((42) & 63))) | 1)))) | 1)))) | 1)))}
        v23 := GenS2{F0: "", F1: v4, F2: 127}
        _ = fi20
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 bool = v2
        if (9223372036854775807 < (fi24 - genFn1((306 != 1), 1))) {
            v26 := GenS1{F0: int(v19)}
            if len(v1) > 0 {
                v1[0] = 2
            }
            v3[2147483647] = "a"
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    switch -9223372036854775808 {
    case 4294967295:
        v6 = float32(0.1)
    }
    fmt.Println(v4, v6, v19)
    v27 := GenS3{F0: 42}
    for fi28 := 0; fi28 < 6; fi28++ {
        v29 := GenS4{F0: 255}
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
