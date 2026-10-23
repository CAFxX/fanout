package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 uint8, p1 string) int16 {
    var v4 int = 42
    var v5 int64 = -925
    var v6 float64 = 1.5
    v6 -= v6
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 int8) int16 {
    if (183 < 2) {
        fmt.Println(p1, p1)
        for fi7 := 0; fi7 < 7; fi7++ {
            p1 = p1
            _ = fi7
        }
        if ((42 >= 1) && ((-2.25 != -2.25) || true)) {
            v8 := GenS1{F0: "hello", F1: (p0 | p0), F2: "foo bar"}
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 float64 = 0.0
                v10 = v10
                var v11 uint32 = 582
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS2{F0: 210}
            _ = v12
            _ = v8
        }
    } else {
        if (27 == 132) {
            v13 := GenS3{F0: -511, F1: 2147483648, F2: 255}
            p1 = -128
            v14 := []int{-9223372036854775808}
            _ = v13
            _ = v14
        } else {
            switch ((("\n\t" + ("" + "\n\t")) + "a") + "") {
            case "hello":
                var v15 uint64 = 464
                _ = v15
            case "":
                var v16 uint16 = 42
                _ = v16
            default:
                for fi17 := 0; fi17 < 18; fi17++ {
                    switch "foo bar" {
                    case "a":
                        var v18 int64 = 0
                        _ = v18
                    case "":
                        p1 = (p1 - p1)
                    case "foo bar":
                        var v19 uint = (p0 << ((2) & 63))
                        _ = v19
                    default:
                        var v20 int32 = 42
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 int = (fi17 | (fi21 * fi21))
                        fmt.Println(fi17, v22)
                        v22 = (-(fi17))
                        _ = fi21
                        _ = v22
                    }
                    _ = fi17
                }
            }
            v23 := GenS4{F0: ("foo bar" + "hello")}
            _ = v23
        }
    }
    fmt.Println(p0, p1)
    p0 += (p0 >> (((p0 / ((250) | 1))) & 63))
    _ = p0
    _ = p1
    return 1
}

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 4294967295
    v3 := make(map[int]bool)
    fmt.Println(v2)
    var v24 uint = 1
    v2 -= v2
    v24 = (v24 & 889)
    v2 = (^(v2))
    var v25 uint32 = (v2 & 42)
    var v26 int = -2147483648
    v27 := []int{}
    v25 -= (v2 ^ (v25 * (v2 + 382)))
    var v28 uint = (v24 / ((18) | 1))
    v24 = v24
    if ((0.0 * (0.0 * (((1e10 - float64((0.0 - float32(1e10)))) - 0.0) / 1e10))) < -2.25) {
        if true {
            var v29 int = (v26 >> ((v28) & 63))
            v25 += (v2 & 936)
            var v30 int = (v29 & (v26 ^ -2147483648))
            _ = v29
            _ = v30
        }
        switch v26 {
        case 9223372036854775807:
            var v31 bool = true
            _ = v31
        case -210:
            v32 := GenS2{F0: 2}
            _ = v32
        }
        if (true || false) {
            var v33 string = ("foo bar" + "hello")
            v34 := GenS2{F0: 1}
            _ = v33
            _ = v34
        } else {
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 uint = v24
                _ = fi35
                _ = v36
            }
            switch (("\n\t" + "hello") + "\n\t") {
            case "\n\t":
                v28 = v28
            case "hello":
                v37 := v3[127]
                _ = v37
            }
            v28 = (v24 | 855)
        }
    } else {
        var v38 int = 0
        _ = v38
    }
    if (!false) {
        v2 += v2
        if (!(false || ("a" <= "a"))) {
            v3[127] = (127 < 1)
            v39 := GenS2{F0: 0}
            _ = v39
        }
    } else {
        delete(v3, (v26 + -978))
        switch ("foo bar" + "a") {
        case "a":
            var v40 float64 = float64(float32(0.1))
            _ = v40
        case "foo bar":
            v41 := GenS1{F0: ("foo bar" + "a"), F1: 85, F2: "foo bar"}
            _ = v41
        default:
            for fi42 := 0; fi42 < 10; fi42++ {
                v25 += (v25 % ((2147483647) | 1))
                _ = fi42
            }
        }
        v2 = (v25 + v25)
    }
    if true {
        var v43 bool = (!(((-(v26)) > -151) && (100 >= 170)))
        _ = v43
    }
    fmt.Println(v2)
    for fi44 := 0; fi44 < 14; fi44++ {
        var v45 int16 = 42
        v46 := GenS5{F0: 18446744073709551615}
        var v47 uint64 = 1
        _ = fi44
        _ = v45
        _ = v46
        _ = v47
    }
    v2 += v25
    v2 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
