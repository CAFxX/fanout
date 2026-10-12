package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 string; F2 uint16 }
type GenS7 struct { F0 string; F1 int16 }
func genFn0(p0 string, p1 string, p2 int64) uint32 {
    v4 := []int{-5, -9223372036854775808, 2147483647, 0}
    v5 := GenS1{F0: 42, F1: 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 67
}
func genFn1(p0 uint8, p1 int64) bool {
    if false {
        v6 := []int{}
        _ = v6
    }
    var v7 uint8 = 128
    var v8 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (4294967295 >= genFn0("a", "\n\t", p1))
}

func main() {
    var v1 uint8 = 44
    var v2 float32 = 0.0
    var v3 float64 = -2.25
    var v9 uint32 = 0
    v10 := GenS2{F0: 18446744073709551615}
    var v11 int64 = 0
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 15; fi13++ {
            if (((-2.25 >= v2) && false) || (((610 >= 42) || false) || false)) {
                var v14 uint = 4294967295
                var v15 int = fi13
                _ = v14
                _ = v15
            }
            fmt.Println(v9, v3)
            v1 = v1
            v16 := []int{fi12, 0, fi12}
            v17 := GenS3{F0: ("foo bar" + ("" + "")), F1: (v9 / (((v9 * v9)) | 1))}
            if (18446744073709551615 < 1) {
                fmt.Println(v9, v2, v11)
                if (116 < (v1 + v1)) {
                    var v18 uint = 2
                    v18 = (v18 * v18)
                    _ = v18
                } else {
                    if (v3 < v3) {
                        var v19 float32 = v2
                        v20 := GenS4{F0: 880}
                        _ = v19
                        _ = v20
                    }
                }
                var v21 int32 = -170
                _ = v21
            }
            _ = fi13
            _ = v16
            _ = v17
        }
        v3 = 0.0
        v22 := []int{2147483647, -304, fi12}
        _ = fi12
        _ = v22
    }
    if (!((!true) || true)) {
        var v23 int8 = -128
        var v24 uint32 = (^(v9))
        v25 := []int{42, 4294967295}
        _ = v23
        _ = v24
        _ = v25
    } else {
        v26 := GenS5{F0: v1}
        v3 = v3
        _ = v26
    }
    var v27 bool = (0 <= 42)
    switch "foo bar" {
    case "a":
        v28 := GenS4{F0: 2147483647}
        _ = v28
    default:
        if v27 {
            if v27 {
                switch (("foo bar" + ("" + "hello")) + "foo bar") {
                case "hello":
                    fmt.Println(v1, v1, v3)
                case "a":
                    var v29 int8 = 0
                    _ = v29
                }
            }
        } else {
            v30 := GenS4{F0: 1}
            _ = v30
        }
    }
    var v31 uint32 = v9
    v32 := GenS6{F0: 1, F1: "foo bar", F2: 1}
    if true {
        var v33 float32 = 1e10
        var v34 uint = 2
        var v35 float64 = (float64(((v33 - (v2 / v2)) - float32(float64((1.5 * v33))))) * (v3 / 1.5))
        _ = v33
        _ = v34
        _ = v35
    } else {
        if (215 <= -2147483648) {
            v36 := []int{-2147483648, 127}
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            _ = v36
        }
        var v37 string = "foo bar"
        _ = v37
    }
    var v38 int32 = 2147483647
    var v39 float64 = v3
    switch 1 {
    case 32768:
        v40 := GenS7{F0: "a", F1: -32768}
        _ = v40
    case 32767:
        var v41 uint16 = 42
        _ = v41
    case 0:
        v2 = float32(v39)
    }
    var v42 uint64 = 159
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v42:", v42)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v39
    _ = v42
    _ = v9
}
