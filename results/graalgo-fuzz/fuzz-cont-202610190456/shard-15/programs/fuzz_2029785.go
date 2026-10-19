package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 int64; F1 string; F2 int }

func main() {
    v1 := make(map[string]string)
    var v2 bool = (-1 > 0)
    var v3 int32 = 2147483647
    var v4 int8 = -19
    var v5 uint16 = 65535
    var v6 int64 = 1
    var v7 int8 = (-(v4))
    v6 = (v6 & (v6 % ((v6) | 1)))
    v8 := GenS1{F0: (v2 || (("a" < "") && v2)), F1: 1}
    fmt.Println(v7, v2, v2)
    if (1000 == 42) {
        fmt.Println(v4, v7, v5)
    } else {
        v1[""] = ("hello" + "")
        v9 := v1["a"]
        _ = v9
    }
    fmt.Println(v3, v3)
    var v10 int8 = 100
    v11 := GenS1{F0: (v3 != v3), F1: 410}
    v12 := GenS2{F0: (v5 - v5)}
    if (!true) {
        v13 := []int{42}
        v10 = (^(v7))
        var v14 int32 = (v3 - (-(v3)))
        _ = v13
        _ = v14
    } else {
        if v2 {
            v3 = v3
            v6 = (-(v6))
        } else {
            v15 := []int{-929, 9223372036854775807, 439, 0}
            _ = v15
        }
        var v16 uint8 = 255
        v17 := GenS2{F0: 42}
        _ = v16
        _ = v17
    }
    if v2 {
        v2 = v2
        fmt.Println(v5, v5)
    } else {
        v18 := []int{1, 128, -20}
        v4 = (v10 - 42)
        _ = v18
    }
    if ((v3 * 0) < (v3 | 253)) {
        var v19 int16 = 810
        if true {
            var v20 uint64 = 4294967295
            if true {
                v2 = v2
            } else {
                v21 := GenS3{F0: 127, F1: (v4 / ((v10) | 1)), F2: 128}
                v22 := GenS2{F0: (^(v5))}
                v23 := v1[("a" + "\n\t")]
                _ = v21
                _ = v22
                _ = v23
            }
            v3 -= v3
            _ = v20
        } else {
            var v24 int = -526
            v25 := GenS4{F0: v6, F1: "foo bar", F2: (v24 >> ((42) & 63))}
            _ = v24
            _ = v25
        }
        v26 := []int{2147483647}
        _ = v19
        _ = v26
    }
    v27 := GenS3{F0: 0, F1: 0, F2: 146}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
