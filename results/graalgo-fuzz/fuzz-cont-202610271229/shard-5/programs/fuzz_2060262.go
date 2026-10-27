package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint32, p1 bool) int32 {
    if (100 <= 100) {
        v7 := []int{9223372036854775807, -9223372036854775808, -2147483648, 1}
        p0 = (p0 * p0)
        if len(v7) > 0 {
            v7[0] = 827
        }
        _ = v7
    } else {
        v8 := GenS1{F0: 172, F1: 719}
        _ = v8
    }
    switch ("\n\t" + "a") {
    case "hello":
        switch "a" {
        case "foo bar":
            var v9 uint8 = 128
            _ = v9
        case "\n\t":
            v10 := []int{}
            _ = v10
        case "hello":
            if (798 < p0) {
                fmt.Println(p1, p1)
                v11 := []int{-2147483648}
                _ = v11
            } else {
                var v12 float64 = -2.25
                var v13 float32 = (0.0 + 1e10)
                _ = v12
                _ = v13
            }
        default:
            v14 := GenS1{F0: 783, F1: 255}
            _ = v14
        }
    case "foo bar":
        fmt.Println(p0)
    }
    p0 -= p0
    var v15 uint32 = p0
    _ = p0
    _ = p1
    _ = v15
    return -933
}
func genFn1(p0 int8, p1 int64) uint64 {
    p1 = (p1 / (((p1 - p1)) | 1))
    if false {
        v16 := []int{9223372036854775807, 897, 521}
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        var v17 int32 = -1
        _ = v16
        _ = v17
    }
    var v18 float64 = (3.14159 / 1e10)
    _ = p0
    _ = p1
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-407, 9223372036854775807, -351, -9223372036854775808}
    var v3 int16 = 719
    v4 := make(map[int]bool)
    var v5 uint16 = 65535
    var v6 int32 = 0
    v1[-2147483648] = 128
    v3 -= (^(v3))
    if false {
        var v19 uint32 = 42
        _ = v19
    } else {
        if (42 == -672) {
            v1[-1] = 595
            var v20 uint32 = 1
            _ = v20
        } else {
            if false {
                var v21 float32 = -2.25
                _ = v21
            }
            switch 0 {
            case 1:
                v22 := GenS1{F0: 683, F1: 255}
                _ = v22
            case 32767:
                fmt.Println(v3, v3, v3)
            }
        }
    }
    switch "" {
    case "hello":
        v23 := GenS1{F0: 656, F1: v5}
        _ = v23
    case "":
        v6 -= v6
    }
    v3 = v3
    if true {
        v6 -= (v6 ^ v6)
        var v24 float32 = (-2.25 - 3.14159)
        if false {
            v2 = append(v2, -984)
            switch (("foo bar" + "foo bar") + ("\n\t" + "")) {
            case "hello":
                fmt.Println(v5, v3)
            default:
                if true {
                    v5 = v5
                    for fi25 := 0; fi25 < 10; fi25++ {
                        v5 += v5
                        var v26 uint16 = 279
                        fmt.Println(v24)
                        _ = fi25
                        _ = v26
                    }
                }
            }
            fmt.Println(v24)
        } else {
            var v27 int = 1
            if (!(((v5 - v5) != 255) || true)) {
                var v28 uint32 = 2147483647
                _ = v28
            } else {
                v29 := GenS2{F0: (v6 % ((int32(v3)) | 1)), F1: (^(v5))}
                if (0 != -1) {
                    v1[(v27 * v27)] = (v27 % (((^(v27))) | 1))
                } else {
                    v2 = append(v2, (v27 % ((1) | 1)))
                    v30 := GenS2{F0: (v6 & genFn0(0, false)), F1: (v5 ^ v5)}
                    if len(v2) > 0 {
                        v2[0] = (v27 + -9223372036854775808)
                    }
                    _ = v30
                }
                _ = v29
            }
            _ = v27
        }
        _ = v24
    } else {
        var v31 int64 = -954
        _ = v31
    }
    var v32 int8 = -128
    var v33 uint32 = 42
    v3 -= v3
    v32 -= (v32 >> ((696) & 63))
    for fi34 := 0; fi34 < 9; fi34++ {
        for fi35 := 0; fi35 < 2; fi35++ {
            delete(v1, (fi35 | (-(fi34))))
            _ = fi35
        }
        var v36 bool = (1.5 != -2.25)
        _ = fi34
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
