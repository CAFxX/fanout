package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64, p1 int8) int64 {
    v5 := GenS1{F0: 262}
    v6 := GenS2{F0: ("a" > ""), F1: 42, F2: 292}
    fmt.Println(p1, p0)
    p0 += 0
    var v7 float32 = 1.5
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 int64) uint {
    var v8 float64 = 1.5
    v8 = (v8 / 0.1)
    if (4294967295 == 42) {
        var v9 bool = (-2147483648 == 1)
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := []int{}
            _ = fi10
            _ = v11
        }
        _ = v9
    } else {
        p0 = p0
    }
    _ = p0
    _ = v8
    return 255
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 int64 = 1
    fmt.Println(v4, v4)
    if ((0.0 == 3.14159) || false) {
        v4 = genFn0(v4, v1)
        var v12 uint32 = 2
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := []int{}
        v15 := GenS1{F0: 2147483648}
        var v16 int32 = 126
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v4, v4)
        _ = fi17
    }
    fmt.Println(v1)
    for fi18 := 0; fi18 < 1; fi18++ {
        if (!false) {
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := GenS3{F0: 18446744073709551615, F1: ("hello" + "")}
                switch 32768 {
                case 679:
                    v1 -= (v1 + 0)
                case 869:
                    v21 := v2[fi18]
                    _ = v21
                default:
                    var v22 bool = (v1 != v1)
                    _ = v22
                }
                _ = fi19
                _ = v20
            }
            fmt.Println(v1, v4, fi18)
            delete(v3, fi18)
        }
        v4 += (v4 - (v4 | 619))
        var v23 int16 = 1000
        _ = fi18
        _ = v23
    }
    v2[830] = 2
    if false {
        if true {
            switch -1 {
            case 0:
                if (((float32(0.1) + (1e10 * 0.1)) + (((0.1 + -2.25) / -2.25) * (((1e10 - 3.14159) * 1.5) / (-2.25 + ((3.14159 + 0.0) / 1e10))))) != ((-2.25 + 0.0) + 1.5)) {
                    fmt.Println(v1, v1)
                    if ((2 == 42) && ((!(!false)) && true)) {
                        v24 := GenS3{F0: 403, F1: "foo bar"}
                        var v25 bool = (!(!false))
                        _ = v24
                        _ = v25
                    }
                    v26 := GenS2{F0: ((!((((((-2.25 + 0.0) - 1e10) / 1.0) <= -2.25) && false) || ((-(v4)) <= genFn0(42, v1)))) || false), F1: 699, F2: 2}
                    _ = v26
                }
            case -1:
                switch 32767 {
                case 9223372036854775807:
                    fmt.Println(v4)
                case 128:
                    var v27 int32 = 42
                    _ = v27
                case -845:
                    var v28 uint32 = 816
                    _ = v28
                }
            default:
                var v29 uint8 = 128
                _ = v29
            }
            var v30 float64 = 0.0
            _ = v30
        }
        var v31 float64 = (-2.25 * (1e10 + 3.14159))
        for fi32 := 0; fi32 < 2; fi32++ {
            var v33 bool = (((("\n\t" + ("" + ("foo bar" + "foo bar"))) < (("foo bar" + (((((("hello" + ("hello" + ((("" + ("a" + "")) + "") + "a"))) + "\n\t") + "hello") + "") + ("\n\t" + "")) + (("" + "\n\t") + "\n\t"))) + "")) && true) || true)
            _ = fi32
            _ = v33
        }
        _ = v31
    }
    if true {
        v34 := GenS4{F0: 65535, F1: 255}
        _ = v34
    } else {
        for fi35 := 0; fi35 < 10; fi35++ {
            var v36 int = (fi35 + (fi35 - (fi35 - 127)))
            _ = fi35
            _ = v36
        }
        v37 := GenS3{F0: 18446744073709551615, F1: "hello"}
        _ = v37
    }
    switch "" {
    case "foo bar":
        var v38 int64 = (v4 & (v4 >> ((2) & 63)))
        _ = v38
    case "a":
        v39 := v2[299]
        _ = v39
    default:
        v4 += (v4 + 709)
    }
    switch "hello" {
    case "hello":
        var v40 int64 = v4
        _ = v40
    case "":
        var v41 float32 = (((1e10 * float32((0.1 - 0.1))) + 1e10) * (float32(0.0) / -2.25))
        _ = v41
    }
    for fi42 := 0; fi42 < 3; fi42++ {
        switch (^(fi42)) {
        case 2147483647:
            var v43 float64 = 3.14159
            _ = v43
        default:
            v44 := []int{(fi42 + (fi42 | 475)), fi42}
            _ = v44
        }
        _ = len(v3)
        v45 := GenS5{F0: 65, F1: -128}
        _ = fi42
        _ = v45
    }
    fmt.Println(v4, v4)
    var v46 bool = (!false)
    v1 -= v1
    fmt.Println(v1)
    var v47 int64 = 9223372036854775807
    fmt.Println(v1, v4, v47)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
}
