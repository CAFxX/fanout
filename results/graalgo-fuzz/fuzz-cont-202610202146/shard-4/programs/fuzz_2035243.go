package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint, p2 bool) int8 {
    fmt.Println(p0, p0)
    p0 = (p0 + (p0 - p0))
    v7 := GenS1{F0: 560, F1: 127, F2: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 100
}
func genFn1(p0 uint, p1 uint64, p2 int32) uint {
    var v8 uint32 = 445
    p2 = p2
    v9 := GenS2{F0: 975}
    fmt.Println(p1)
    var v10 int8 = -35
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (p0 % (((p0 + 504)) | 1))
}

func main() {
    var v1 uint16 = 65535
    var v2 uint32 = 655
    var v3 float64 = 0.0
    var v4 bool = true
    v5 := []uint32{(v2 * 455), v2}
    v6 := make(map[uint32]string)
    if v4 {
        fmt.Println(v2, v1, v2)
        v3 = -2.25
        v6[1] = "\n\t"
    } else {
        v11 := GenS2{F0: 0}
        v1 = v1
        v12 := GenS1{F0: -1, F1: 42, F2: -9223372036854775808}
        _ = v11
        _ = v12
    }
    v2 -= v2
    var v13 int64 = 42
    v5 = append(v5, v2)
    fmt.Println(v2, v2, v1)
    var v14 int8 = -128
    switch 2147483647 {
    case 127:
        v13 -= v13
    case 91:
        if (128 != 4294967295) {
            if ((-2147483648 <= 42) && true) {
                v15 := GenS2{F0: genFn1(42, 962, -2147483648)}
                _ = v15
            }
            v4 = (-2.25 != (1e10 - ((1.5 + 3.14159) + 1e10)))
        }
    default:
        fmt.Println(v13, v14, v14)
    }
    var v16 uint = 255
    if (false && (v4 || (474 <= 2147483647))) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        var v17 float64 = v3
        _ = v17
    }
    if (42 > 2147483647) {
        for fi18 := 0; fi18 < 2; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                switch "\n\t" {
                case "a":
                    if len(v5) > 0 {
                        v5[0] = (v2 / ((v2) | 1))
                    }
                case "foo bar":
                    v16 = (v16 ^ 889)
                case "":
                    fmt.Println(fi19, fi18)
                }
                _ = fi19
            }
            v20 := GenS2{F0: v16}
            _ = fi18
            _ = v20
        }
        if false {
            v4 = v4
            _ = len(v6)
            var v21 float64 = (v3 / v3)
            _ = v21
        }
        fmt.Println(v14, v2, v1)
    }
    fmt.Println(v2, v16, v13)
    v6[(v2 * (v2 + (^(v2))))] = "a"
    for fi22 := 0; fi22 < 12; fi22++ {
        for fi23 := 0; fi23 < 17; fi23++ {
            _ = len(v5)
            _ = fi23
        }
        var v24 uint32 = 0
        v25 := GenS3{F0: "", F1: -2147483648, F2: 1}
        _ = fi22
        _ = v24
        _ = v25
    }
    var v26 float64 = (float64(((-2.25 + 3.14159) * -2.25)) * (v3 + v3))
    v6[v2] = ""
    switch "" {
    case "":
        v26 = 0.1
    case "foo bar":
        switch "a" {
        case "a":
            v27 := GenS1{F0: 1, F1: (v14 + v14), F2: 32767}
            _ = v27
        case "\n\t":
            v28 := GenS2{F0: (^(v16))}
            _ = v28
        default:
            var v29 int32 = -1
            _ = v29
        }
    case "\n\t":
        if (!v4) {
            v16 -= 981
            v6[(v2 << (((^(v16))) & 63))] = ("foo bar" + "\n\t")
            var v30 int64 = (v13 % ((v13) | 1))
            _ = v30
        } else {
            fmt.Println(v16, v14, v3)
            v1 = 575
        }
    default:
        var v31 uint32 = 419
        _ = v31
    }
    if (!v4) {
        switch 2147483647 {
        case 32767:
            v32 := GenS3{F0: ((("foo bar" + ("\n\t" + ((("\n\t" + "a") + "") + ((("hello" + (("\n\t" + "foo bar") + (("\n\t" + ((("\n\t" + "a") + ("\n\t" + "a")) + (("hello" + "hello") + (("hello" + ("foo bar" + "")) + "foo bar")))) + "foo bar"))) + "a") + "foo bar")))) + "\n\t") + ""), F1: -1, F2: 18446744073709551615}
            _ = v32
        default:
            _ = len(v6)
        }
        for fi33 := 0; fi33 < 4; fi33++ {
            var v34 uint32 = (v2 - 2147483647)
            if len(v5) > 0 {
                v5[0] = (v2 - v2)
            }
            var v35 int64 = (v13 & (v13 ^ v13))
            _ = fi33
            _ = v34
            _ = v35
        }
        var v36 int32 = 518
        _ = v36
    } else {
        v37 := GenS4{F0: uint64(v16)}
        var v38 int64 = (v13 + v13)
        v6[(v2 ^ (v2 ^ (v2 * v2)))] = "foo bar"
        _ = v37
        _ = v38
    }
    var v39 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
