package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 int; F2 int64 }
type GenS5 struct { F0 int32; F1 int8; F2 int8 }
func genFn0() int8 {
    if (false && (265 < 2147483647)) {
        for fi7 := 0; fi7 < 4; fi7++ {
            fmt.Println(fi7)
            var v8 uint8 = 128
            _ = fi7
            _ = v8
        }
    } else {
        var v9 float32 = -2.25
        var v10 uint8 = 1
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 524, F1: "a"}
    v12 := GenS1{F0: 9223372036854775807, F1: "foo bar"}
    _ = v11
    _ = v12
    return 100
}
func genFn1() bool {
    var v13 float32 = -2.25
    v14 := GenS1{F0: -9223372036854775808, F1: "foo bar"}
    if false {
        v13 += (((v13 / float32((((1.5 * (0.0 / (float64(3.14159) - -2.25))) * 3.14159) - 0.1))) + float32(0.1)) + v13)
        if (42 >= 333) {
            var v15 string = (("hello" + ("\n\t" + ("hello" + "a"))) + "")
            v16 := []int{-1, 128, -438}
            switch "foo bar" {
            case "hello":
                var v17 int = -1
                _ = v17
            }
            _ = v15
            _ = v16
        }
    } else {
        var v18 uint = 4294967295
        var v19 uint32 = 241
        _ = v18
        _ = v19
    }
    _ = v13
    _ = v14
    return (18446744073709551615 != 4294967295)
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 128
    v3 := []uint32{}
    var v4 uint8 = v2
    var v5 uint16 = 265
    v6 := make(map[uint32]bool)
    for fi20 := 0; fi20 < 11; fi20++ {
        v2 += 1
        if true {
            var v21 int64 = 9223372036854775807
            v2 = 128
            fmt.Println(fi20, fi20, v21)
            _ = v21
        } else {
            var v22 bool = false
            v3 = append(v3, 2147483647)
            _ = v22
        }
        _ = fi20
    }
    switch "a" {
    case "foo bar":
        v3 = append(v3, 2147483648)
    case "\n\t":
        var v23 uint64 = uint64(v2)
        _ = v23
    default:
        v24 := GenS1{F0: -1, F1: ("" + "a")}
        _ = v24
    }
    v6[65535] = ((!false) && false)
    if true {
        if false {
            var v25 int64 = -1
            _ = v25
        }
        var v26 uint32 = 209
        _ = v26
    } else {
        v5 = 2
        var v27 float64 = 1e10
        v5 = v5
        _ = v27
    }
    v5 += 909
    v4 += 255
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v28 := GenS2{F0: (v5 - v5), F1: 1}
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 3; fi30++ {
            v31 := GenS1{F0: -1, F1: ("a" + ("a" + "\n\t"))}
            _ = fi30
            _ = v31
        }
        fmt.Println(v5)
        v32 := v6[68]
        _ = fi29
        _ = v32
    }
    _ = len(v6)
    if true {
        if (18446744073709551615 <= 255) {
            delete(v6, 2)
            var v33 uint64 = 458
            _ = v33
        }
        v34 := GenS3{F0: 9223372036854775807}
        var v35 float32 = (1e10 / 1.0)
        _ = v34
        _ = v35
    } else {
        v5 -= v5
        var v36 bool = false
        _ = v36
    }
    v4 = 185
    if (42 != 128) {
        fmt.Println(v2, v4)
        if (true || ((18446744073709551615 <= 255) || true)) {
            if len(v3) > 0 {
                v3[0] = 704
            }
            if false {
                var v37 uint8 = v2
                if (-2147483648 != 276) {
                    fmt.Println(v5, v37, v4)
                    v5 = v5
                    v4 = (v37 & 42)
                }
                _ = v37
            } else {
                v38 := GenS4{F0: 42, F1: 707, F2: -915}
                _ = v38
            }
            var v39 int16 = 42
            _ = v39
        } else {
            v40 := GenS4{F0: 704, F1: 1, F2: -131}
            _ = v40
        }
    } else {
        v2 = v2
        v41 := GenS5{F0: 2147483647, F1: 42, F2: 100}
        v42 := v1[0]
        _ = v41
        _ = v42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
