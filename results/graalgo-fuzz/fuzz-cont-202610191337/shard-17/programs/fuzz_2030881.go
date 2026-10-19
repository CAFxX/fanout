package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint64 }
func genFn0(p0 int8, p1 uint32) int {
    v5 := GenS1{F0: 0}
    if (-2.25 >= ((0.0 * ((float32(-2.25) + 0.0) - 0.0)) + 0.0)) {
        var v6 uint64 = 2
        p1 = (p1 & 42)
        _ = v6
    } else {
        p1 -= 284
        var v7 bool = (p1 >= 42)
        if ((p1 * 1) >= 65535) {
            v8 := []int{-1, 9223372036854775807, 9223372036854775807}
            fmt.Println(v7, p1, p0)
            _ = v8
        } else {
            v9 := []int{239, 4294967295}
            if (0 < uint8(p1)) {
                var v10 bool = v7
                _ = v10
            }
            _ = v9
        }
        _ = v7
    }
    var v11 uint8 = 134
    v11 -= v11
    if true {
        v12 := GenS1{F0: 186}
        for fi13 := 0; fi13 < 8; fi13++ {
            fmt.Println(fi13, fi13, fi13)
            _ = fi13
        }
        _ = v12
    } else {
        var v14 float32 = 0.0
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v5
    return -799
}

func main() {
    var v1 uint64 = 42
    var v2 int32 = 42
    v3 := []string{"\n\t", ""}
    var v4 uint64 = (v1 >> ((1) & 63))
    if (v1 == 18446744073709551615) {
        if true {
            if true {
                var v15 uint = 4294967295
                _ = v15
            }
        }
        v16 := GenS1{F0: 4294967295}
        v3 = append(v3, "a")
        _ = v16
    }
    var v17 uint16 = 511
    v18 := GenS1{F0: 42}
    fmt.Println(v17)
    var v19 int32 = (v2 * v2)
    v20 := GenS2{F0: "a", F1: 42}
    var v21 uint16 = (v17 * 160)
    v22 := GenS3{F0: 2, F1: 478, F2: (v1 & v4)}
    v19 = (v19 ^ v2)
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS3{F0: 255, F1: -32768, F2: (v1 / ((695) | 1))}
        v25 := GenS1{F0: 4294967295}
        v3 = append(v3, "\n\t")
        _ = fi23
        _ = v24
        _ = v25
    }
    v19 -= (v2 / (((v2 - 764)) | 1))
    switch ("hello" + "\n\t") {
    case "foo bar":
        fmt.Println(v2, v4, v4)
    case "a":
        var v26 int32 = -2147483648
        _ = v26
    }
    if false {
        v19 -= (^(v19))
        for fi27 := 0; fi27 < 17; fi27++ {
            v3 = append(v3, "")
            v19 += (v2 | (v19 << ((65535) & 63)))
            v28 := GenS3{F0: 128, F1: 32767, F2: v1}
            _ = fi27
            _ = v28
        }
    } else {
        v3 = append(v3, "")
        v19 = (v19 ^ v2)
        var v29 uint32 = 2147483647
        _ = v29
    }
    v2 += v19
    var v30 int = -852
    var v31 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
