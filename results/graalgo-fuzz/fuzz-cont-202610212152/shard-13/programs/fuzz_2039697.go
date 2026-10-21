package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 string }
type GenS2 struct { F0 int32; F1 int32; F2 int64 }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string, p1 int64) int16 {
    p0 = ((p0 + p0) + ("\n\t" + p0))
    v5 := []int{-32, 609, 874, -75}
    var v6 uint8 = 128
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 586
}
func genFn1(p0 int8, p1 uint32) int16 {
    var v7 uint8 = 150
    var v8 int64 = -725
    var v9 int = 32767
    v10 := GenS1{F0: 1, F1: (("" + "") + "\n\t"), F2: ""}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return genFn0("\n\t", v8)
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    var v3 uint64 = 663
    var v4 float32 = 0.0
    var v11 int32 = 42
    v3 = v1
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: v1, F1: "", F2: "foo bar"}
        for fi14 := 0; fi14 < 9; fi14++ {
            switch (("foo bar" + "") + "foo bar") {
            case "":
                v15 := v2[(fi14 | fi14)]
                _ = v15
            }
            fmt.Println(v1, v1, v11)
            if false {
                v16 := GenS1{F0: v3, F1: "foo bar", F2: ("a" + "a")}
                if (!(!(false || (42 >= 2147483648)))) {
                    v3 = (^(v3))
                }
                _ = v16
            }
            fmt.Println(fi14)
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    v17 := []int{127, 0}
    if (2 != 255) {
        v18 := GenS2{F0: (v11 << ((509) & 63)), F1: -629, F2: 0}
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        var v20 uint8 = 127
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: (v11 >> ((18446744073709551615) & 63)), F1: -777, F2: 1}
    var v22 bool = (((!(-1 >= 42)) || (!false)) && false)
    var v23 bool = v22
    if (1 >= 23) {
        fmt.Println(v4, v11, v4)
        switch ("" + "a") {
        case "hello":
            var v24 int8 = -47
            _ = v24
        case "\n\t":
            v1 += v1
        case "foo bar":
            v23 = v23
        default:
            v25 := GenS3{F0: 815, F1: 0}
            _ = v25
        }
        v26 := v2[9223372036854775807]
        _ = v26
    } else {
        if (v11 == (v11 - (v11 >> ((258) & 63)))) {
            fmt.Println(v22, v23)
            switch -722 {
            case 32768:
                if v22 {
                    fmt.Println(v22, v22)
                } else {
                    var v27 uint32 = 592
                    _ = v27
                }
            case -2147483648:
                fmt.Println(v4, v23)
            default:
                var v28 int64 = 609
                _ = v28
            }
        }
        var v29 uint32 = 139
        fmt.Println(v29, v3, v3)
        _ = v29
    }
    if len(v17) > 0 {
        v17[0] = 32767
    }
    var v30 uint8 = 127
    if false {
        v22 = ((v30 + v30) < 2)
    }
    var v31 bool = v23
    switch 32768 {
    case 32767:
        v3 -= (v3 << ((42) & 63))
    case 0:
        switch "foo bar" {
        case "hello":
            _ = len(v2)
        case "foo bar":
            v22 = (((-2.25 + 0.0) - 1e10) < 1.5)
        case "\n\t":
            var v32 uint8 = v30
            _ = v32
        default:
            fmt.Println(v4)
        }
    default:
        v30 += v30
    }
    var v33 uint16 = 2
    switch ("a" + "foo bar") {
    case "\n\t":
        fmt.Println(v4, v1, v23)
    case "":
        var v34 int8 = -37
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v4
}
