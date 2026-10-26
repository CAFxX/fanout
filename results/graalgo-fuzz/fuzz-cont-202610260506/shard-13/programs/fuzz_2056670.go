package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int32 }
func genFn0(p0 int, p1 uint32) uint64 {
    fmt.Println(p1, p0)
    var v6 uint8 = 0
    v7 := []int{}
    if false {
        p1 = uint32(v6)
    }
    if false {
        switch "\n\t" {
        case "a":
            fmt.Println(p1, v6)
        case "":
            v8 := GenS1{F0: -2147483648}
            _ = v8
        default:
            var v9 int64 = 42
            _ = v9
        }
        p1 = p1
    } else {
        v10 := GenS1{F0: -1}
        var v11 uint8 = (v6 / (((v6 >> ((18446744073709551615) & 63))) | 1))
        var v12 uint32 = (^(p1))
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 uint, p2 uint16) uint8 {
    if (127 != 32) {
        p1 += p1
    } else {
        v13 := []int{int(p0)}
        _ = v13
    }
    var v14 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    return uint8(p1)
}

func main() {
    var v1 bool = false
    v2 := []int{1, -592, -853, -2147483648}
    var v3 uint = 65535
    v4 := make(map[uint32]int)
    v5 := []uint32{2147483647}
    v15 := GenS1{F0: 1}
    if v1 {
        v3 = (v3 << (((v3 | v3)) & 63))
    } else {
        if len(v5) > 0 {
            v5[0] = 433
        }
    }
    v3 += v3
    for fi16 := 0; fi16 < 1; fi16++ {
        v3 -= v3
        _ = fi16
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    switch (("\n\t" + "hello") + "") {
    case "\n\t":
        fmt.Println(v1, v3)
    case "a":
        for fi17 := 0; fi17 < 16; fi17++ {
            if (((-2.25 - 0.0) + (float32(0.1) + 3.14159)) >= 3.14159) {
                delete(v4, 813)
            }
            _ = fi17
        }
    default:
        v18 := GenS2{F0: "\n\t"}
        _ = v18
    }
    var v19 int64 = 1
    v20 := GenS2{F0: ""}
    _ = len(v4)
    for fi21 := 0; fi21 < 15; fi21++ {
        if (int32(v19) > 42) {
            for fi22 := 0; fi22 < 15; fi22++ {
                if (false || ("hello" != ("" + "\n\t"))) {
                    v3 += 65535
                }
                v23 := GenS1{F0: 2147483647}
                v24 := GenS3{F0: v1, F1: 1}
                var v25 float32 = (1.5 * -2.25)
                var v26 float64 = (3.14159 + 1e10)
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 string = "a"
            _ = v27
        }
        fmt.Println(v1, v1, v3)
        if false {
            fmt.Println(v3, fi21, v19)
            v5 = append(v5, 958)
        }
        _ = fi21
    }
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 int = (-(fi28))
        v30 := GenS3{F0: v1, F1: 1}
        _ = fi28
        _ = v29
        _ = v30
    }
    v1 = v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
