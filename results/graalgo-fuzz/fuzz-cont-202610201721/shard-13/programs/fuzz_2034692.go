package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS8 struct { F0 int64; F1 uint; F2 int8 }
func genFn0() int {
    if false {
        v6 := []int{9223372036854775807, 9223372036854775807, 42}
        v7 := GenS1{F0: 4294967295, F1: -34, F2: "\n\t"}
        var v8 int8 = -1
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 float32 = -2.25
    _ = v9
    return -2147483648
}
func genFn1(p0 string) uint {
    if (false || (false || (!false))) {
        var v10 int16 = 1
        var v11 string = p0
        v12 := []int{-1, 128}
        _ = v10
        _ = v11
        _ = v12
    }
    if (-2147483648 > 691) {
        var v13 float32 = 1e10
        _ = v13
    } else {
        for fi14 := 0; fi14 < 4; fi14++ {
            var v15 int64 = -343
            for fi16 := 0; fi16 < 6; fi16++ {
                v17 := []int{genFn0(), (fi16 / ((genFn0()) | 1))}
                _ = fi16
                _ = v17
            }
            _ = fi14
            _ = v15
        }
        p0 = (p0 + ((p0 + (p0 + (p0 + p0))) + p0))
        if false {
            v18 := GenS1{F0: 1, F1: 42, F2: p0}
            _ = v18
        }
    }
    p0 = p0
    var v19 uint = 945
    var v20 uint = 985
    _ = p0
    _ = v19
    _ = v20
    return v20
}

func main() {
    v1 := []uint32{2, 2}
    var v2 uint64 = 0
    v3 := []uint32{2147483647, 0, 4294967295, 394}
    var v4 int = -263
    v5 := make(map[string]uint32)
    v2 += (^(v2))
    v21 := GenS2{F0: 915}
    var v22 float64 = 1e10
    _ = len(v5)
    v2 = (v2 + v2)
    v4 -= (-(v4))
    var v23 uint = 18446744073709551615
    for fi24 := 0; fi24 < 5; fi24++ {
        switch "a" {
        case "hello":
            delete(v5, "a")
        }
        _ = len(v5)
        var v25 uint32 = 444
        _ = fi24
        _ = v25
    }
    var v26 int32 = 1
    v5["foo bar"] = 4294967295
    switch "\n\t" {
    case "":
        switch genFn0() {
        case 0:
            switch 42 {
            case 9223372036854775807:
                v23 -= (v23 & v23)
            case 127:
                fmt.Println(v4, v23, v26)
            case 70:
                for fi27 := 0; fi27 < 4; fi27++ {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    var v28 uint = v23
                    _ = fi27
                    _ = v28
                }
            }
        default:
            if (v22 > v22) {
                fmt.Println(v22, v23, v26)
                v29 := GenS3{F0: v26}
                _ = v29
            } else {
                v30 := GenS4{F0: -9223372036854775808, F1: "a"}
                _ = v30
            }
        }
    case "foo bar":
        var v31 int8 = int8(v4)
        _ = v31
    default:
        if false {
            v2 = 0
        } else {
            for fi32 := 0; fi32 < 8; fi32++ {
                v33 := GenS5{F0: v4}
                switch (v4 * (fi32 >> (((v23 + 708)) & 63))) {
                case 9223372036854775807:
                    var v34 uint16 = 0
                    _ = v34
                case 0:
                    v26 = (v26 + (-(v26)))
                case 32768:
                    if false {
                        v35 := GenS6{F0: 255, F1: ("a" + "")}
                        fmt.Println(v2, v26)
                        _ = v35
                    } else {
                        v23 = (v23 + v23)
                        v23 -= v23
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi32
                _ = v33
            }
            v36 := GenS7{F0: 448, F1: 42, F2: 127}
            _ = v36
        }
    }
    switch v4 {
    case 32767:
        v37 := GenS6{F0: 163, F1: "foo bar"}
        _ = v37
    case 446:
        for fi38 := 0; fi38 < 10; fi38++ {
            switch "foo bar" {
            case "":
                v4 += (v4 | fi38)
            default:
                v39 := GenS8{F0: 9223372036854775807, F1: genFn1(("" + "hello")), F2: 42}
                _ = v39
            }
            fmt.Println(v22)
            v22 = v22
            _ = fi38
        }
    }
    switch v4 {
    case 128:
        if (-128 >= -50) {
            var v40 int32 = (-(v26))
            v22 = (v22 + (v22 / (v22 / (-2.25 - v22))))
            v26 = (v26 - v26)
            _ = v40
        } else {
            v26 = v26
            _ = len(v3)
            for fi41 := 0; fi41 < 10; fi41++ {
                v42 := v5[""]
                fmt.Println(v4)
                var v43 uint8 = 255
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    case 0:
        v4 = v4
    default:
        if ("foo bar" >= ("" + ("hello" + "a"))) {
            v26 -= (v26 ^ 42)
        } else {
            v5[("hello" + "\n\t")] = 4294967295
            v4 = (v4 & (v4 << ((2) & 63)))
            switch "a" {
            case "\n\t":
                v44 := GenS7{F0: v2, F1: int8(v4), F2: 145}
                _ = v44
            }
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
