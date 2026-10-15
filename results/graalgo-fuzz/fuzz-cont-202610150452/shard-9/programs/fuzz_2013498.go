package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint32 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 bool }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int }
func genFn0(p0 bool) bool {
    for fi4 := 0; fi4 < 19; fi4++ {
        switch 32768 {
        case -2147483648:
            if p0 {
                var v5 uint64 = 2
                _ = v5
            } else {
                p0 = (42 != 157)
                v6 := []int{fi4}
                if len(v6) > 0 {
                    v6[0] = (-(fi4))
                }
                _ = v6
            }
        case -227:
            v7 := GenS1{F0: 587, F1: -85, F2: 821}
            _ = v7
        case -424:
            var v8 bool = p0
            _ = v8
        }
        fmt.Println(fi4, p0)
        _ = fi4
    }
    v9 := []int{9223372036854775807, 617}
    _ = p0
    _ = v9
    return p0
}
func genFn1(p0 int8, p1 uint, p2 uint) uint32 {
    p0 -= (p0 << (((p1 * p2)) & 63))
    for fi10 := 0; fi10 < 15; fi10++ {
        fmt.Println(p2, p0)
        _ = fi10
    }
    var v11 uint16 = 42
    var v12 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (v12 ^ 65535)
}

func main() {
    v1 := []uint32{1, 2147483648, 42}
    var v2 uint8 = 0
    v3 := []int{9223372036854775807, 653}
    v2 = (v2 * v2)
    v2 = v2
    for fi13 := 0; fi13 < 5; fi13++ {
        _ = len(v1)
        _ = fi13
    }
    var v14 float32 = 0.0
    switch "hello" {
    case "a":
        if len(v1) > 0 {
            v1[0] = 42
        }
    case "":
        var v15 float64 = 0.1
        _ = v15
    default:
        var v16 string = "hello"
        _ = v16
    }
    switch -96 {
    case 32768:
        switch (("a" + "") + "hello") {
        case "foo bar":
            v14 = 3.14159
        case "a":
            var v17 uint32 = 65535
            _ = v17
        default:
            if (-37 > -32768) {
                v18 := GenS1{F0: 65535, F1: 0, F2: 2147483648}
                v19 := GenS1{F0: 0, F1: -239, F2: genFn1(-128, 801, 4294967295)}
                var v20 string = (("" + ("foo bar" + "hello")) + "")
                _ = v18
                _ = v19
                _ = v20
            }
        }
    case 127:
        for fi21 := 0; fi21 < 17; fi21++ {
            v22 := GenS1{F0: 0, F1: -580, F2: 121}
            _ = fi21
            _ = v22
        }
    }
    var v23 uint32 = 4294967295
    v24 := GenS2{F0: ("\n\t" + "a"), F1: 1}
    if (4294967295 == 2147483647) {
        var v25 int64 = 9223372036854775807
        _ = v25
    } else {
        var v26 uint = 929
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v27 := GenS2{F0: ((((("foo bar" + "") + (("foo bar" + "") + "\n\t")) + ("\n\t" + ("hello" + ("" + ("a" + ("\n\t" + "\n\t")))))) + "hello") + ""), F1: (v2 ^ 128)}
        _ = v26
        _ = v27
    }
    v2 = 255
    v28 := GenS3{F0: "", F1: ((!(428 != 0)) && ((false && false) || true)), F2: (1 < 1)}
    var v29 bool = true
    if v29 {
        fmt.Println(v14, v23)
        v30 := GenS4{F0: 42, F1: 32767}
        _ = v30
    } else {
        var v31 uint64 = 0
        _ = v31
    }
    var v32 uint = 0
    var v33 float64 = (float64(v14) - (-2.25 + ((((1e10 - 1e10) / (float64(float32(0.1)) / (3.14159 + float64((((3.14159 - (v14 + 1.5)) * (v14 * (v14 * (v14 + (-2.25 / v14))))) - v14))))) / ((1e10 + (float64(((float32(-2.25) + v14) - v14)) + 0.0)) + ((-2.25 - 0.1) + -2.25))) * 3.14159)))
    switch "" {
    case "hello":
        fmt.Println(v32)
    case "a":
        var v34 int64 = 9223372036854775807
        _ = v34
    }
    if (v32 < 18446744073709551615) {
        switch 9223372036854775807 {
        case -2147483648:
            var v35 bool = (((1e10 - v33) - ((float64(0.1) - (float64(v14) + 0.0)) / float64(v14))) <= (float64(3.14159) / 1e10))
            _ = v35
        default:
            if (v23 <= (v23 - v23)) {
                switch ("foo bar" + (("hello" + "a") + "a")) {
                case "hello":
                    var v36 uint16 = 32767
                    _ = v36
                }
                v37 := GenS5{F0: 128}
                var v38 int = 128
                _ = v37
                _ = v38
            }
        }
    } else {
        var v39 uint = 782
        _ = v39
    }
    v14 += v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v33
}
