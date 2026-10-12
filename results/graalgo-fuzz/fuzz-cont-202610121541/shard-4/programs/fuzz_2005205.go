package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 int32; F1 bool; F2 uint8 }
type GenS6 struct { F0 uint32; F1 int32; F2 bool }

func main() {
    var v1 int = 42
    var v2 uint64 = 42
    var v3 float32 = 0.1
    v4 := []int{v1}
    if (float64(v3) != (3.14159 + 3.14159)) {
        v3 = ((1e10 - ((v3 - v3) * v3)) - v3)
        for fi5 := 0; fi5 < 2; fi5++ {
            var v6 float64 = (3.14159 - 0.1)
            v1 += -574
            _ = fi5
            _ = v6
        }
    } else {
        var v7 int64 = -9223372036854775808
        v2 = v2
        v2 += 4294967295
        _ = v7
    }
    if false {
        var v8 int16 = 1000
        v9 := GenS1{F0: 42}
        _ = v8
        _ = v9
    } else {
        if len(v4) > 0 {
            v4[0] = (v1 % ((-126) | 1))
        }
        var v10 string = "a"
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := GenS2{F0: 2, F1: 1, F2: 0}
            fmt.Println(v3, v3, fi11)
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    v13 := GenS3{F0: 198}
    if (41 < 127) {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 float64 = 1.5
            for fi16 := 0; fi16 < 17; fi16++ {
                v1 = (v1 / ((fi16) | 1))
                v15 = (-2.25 * (0.1 - ((v15 + 1.5) / v15)))
                _ = fi16
            }
            switch (fi14 % ((fi14) | 1)) {
            case 0:
                for fi17 := 0; fi17 < 7; fi17++ {
                    v2 = 477
                    v3 = ((((1e10 * v3) - (v3 + v3)) + ((((v3 - 0.1) / v3) / float32((v15 + v15))) / (v3 + v3))) + (v3 - (v3 * v3)))
                    var v18 uint32 = 4294967295
                    _ = fi17
                    _ = v18
                }
            default:
                v15 += 0.1
            }
            _ = fi14
            _ = v15
        }
        if true {
            if (false && true) {
                v19 := GenS3{F0: -947}
                for fi20 := 0; fi20 < 7; fi20++ {
                    v1 = -1
                    _ = len(v4)
                    v21 := GenS4{F0: 127, F1: -128, F2: 1}
                    _ = fi20
                    _ = v21
                }
                _ = v19
            } else {
                v1 -= v1
                fmt.Println(v2)
                var v22 uint16 = 2
                _ = v22
            }
            v2 = (v2 / ((0) | 1))
        }
        if ((v1 < v1) || false) {
            v1 -= 319
            switch "" {
            case "a":
                v1 = (v1 & v1)
            }
            v2 -= 35
        } else {
            v2 -= (v2 | 0)
            v23 := GenS1{F0: 255}
            _ = v23
        }
    } else {
        var v24 uint = 42
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (float64(1.5) <= 0.1) {
            var v25 float32 = v3
            if (!(!false)) {
                v24 = (v24 + (v24 * 786))
            } else {
                v2 -= v2
            }
            _ = v25
        }
        _ = v24
    }
    switch (v1 >> ((1) & 63)) {
    case 457:
        var v26 uint8 = 127
        _ = v26
    default:
        v27 := GenS2{F0: 65535, F1: 847, F2: 2}
        _ = v27
    }
    var v28 float64 = 1e10
    var v29 int8 = 42
    v3 += ((1e10 / (((v3 + (v3 / 3.14159)) - 1.5) - (1e10 / (v3 + ((v3 / ((1.5 + -2.25) / 1e10)) - v3))))) / v3)
    switch (v1 << ((152) & 63)) {
    case 32:
        var v30 uint32 = 949
        _ = v30
    case -1:
        v29 += v29
    default:
        for fi31 := 0; fi31 < 17; fi31++ {
            if true {
                v32 := GenS5{F0: 0, F1: false, F2: 1}
                var v33 uint32 = 410
                v34 := GenS6{F0: (v33 & 4294967295), F1: 42, F2: (false && (1 != 494))}
                _ = v32
                _ = v33
                _ = v34
            } else {
                var v35 int8 = (v29 - (^(v29)))
                _ = v35
            }
            _ = fi31
        }
    }
    fmt.Println(v2, v2)
    var v36 bool = true
    v36 = (("hello" < "foo bar") || v36)
    v37 := GenS6{F0: 4294967295, F1: 0, F2: (("a" + "") == ("\n\t" + ""))}
    switch (v1 | 32768) {
    case 1:
        var v38 int32 = 2147483647
        _ = v38
    }
    var v39 string = ("a" + ("foo bar" + "\n\t"))
    v40 := GenS4{F0: 255, F1: v29, F2: -865}
    v28 = ((-2.25 / -2.25) + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v40
}
