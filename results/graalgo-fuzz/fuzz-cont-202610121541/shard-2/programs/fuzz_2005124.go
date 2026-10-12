package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 string; F2 int64 }
type GenS4 struct { F0 int64 }
func genFn0() int32 {
    switch 864 {
    case -1:
        v5 := []int{40}
        _ = v5
    case -9223372036854775808:
        var v6 uint64 = 2
        _ = v6
    case 128:
        if true {
            var v7 bool = true
            v8 := []int{128}
            _ = v7
            _ = v8
        } else {
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 uint64 = 269
                _ = fi9
                _ = v10
            }
        }
    default:
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 828, F1: -1}
            var v13 uint64 = 0
            v13 -= (v13 | 2)
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        if (!true) {
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 float32 = 1e10
                v17 := GenS1{F0: (fi15 - fi14), F1: 1}
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 int16 = 84
            _ = v18
        } else {
            var v19 uint = 42
            fmt.Println(fi14)
            var v20 int = fi14
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float32 = float32(3.14159)
            var v23 string = ""
            _ = fi21
            _ = v22
            _ = v23
        }
        switch fi14 {
        case 4294967295:
            if true {
                switch -1 {
                case 1:
                    v24 := []int{fi14}
                    _ = v24
                case 750:
                    v25 := []int{0, fi14}
                    _ = v25
                case -9223372036854775808:
                    var v26 int = (fi14 / (((fi14 / ((fi14) | 1))) | 1))
                    _ = v26
                default:
                    var v27 uint8 = 63
                    _ = v27
                }
                fmt.Println(fi14)
            } else {
                var v28 uint64 = 42
                v29 := []int{fi14, 9223372036854775807, 487, fi14}
                if ((!false) || false) {
                    v28 += (v28 ^ v28)
                    if (false || ((!true) || true)) {
                        var v30 string = ("foo bar" + "foo bar")
                        v30 = "foo bar"
                        v31 := GenS2{F0: (-(fi14)), F1: "a"}
                        _ = v30
                        _ = v31
                    }
                } else {
                    if len(v29) > 0 {
                        v29[0] = fi14
                    }
                    var v32 int32 = 1
                    _ = v32
                }
                _ = v28
                _ = v29
            }
        case 943:
            v33 := []int{fi14, -9223372036854775808, (fi14 | fi14)}
            _ = v33
        }
        _ = fi14
    }
    return 0
}
func genFn1(p0 bool, p1 int16) int16 {
    var v34 float64 = float64(1.5)
    v34 -= (((v34 * v34) / 1.0) - v34)
    for fi35 := 0; fi35 < 19; fi35++ {
        v36 := GenS1{F0: fi35, F1: 0}
        _ = fi35
        _ = v36
    }
    _ = p0
    _ = p1
    _ = v34
    return p1
}

func main() {
    var v1 string = ""
    var v2 uint8 = 128
    var v3 bool = (514 != 32767)
    v4 := make(map[string]bool)
    if (!(42 > 1)) {
        v2 -= v2
        v37 := GenS1{F0: 402, F1: 0}
        _ = v37
    }
    fmt.Println(v2, v3)
    delete(v4, v1)
    v38 := []int{1, 0, 32767}
    v39 := GenS2{F0: 32767, F1: (("\n\t" + (v1 + v1)) + (("\n\t" + v1) + "hello"))}
    var v40 int64 = 9223372036854775807
    if v3 {
        v1 = ""
        for fi41 := 0; fi41 < 11; fi41++ {
            v3 = v3
            var v42 uint64 = 0
            v42 -= 205
            _ = fi41
            _ = v42
        }
    }
    var v43 int8 = -1
    for fi44 := 0; fi44 < 18; fi44++ {
        for fi45 := 0; fi45 < 19; fi45++ {
            fmt.Println(v43, v40, v1)
            _ = fi45
        }
        _ = fi44
    }
    for fi46 := 0; fi46 < 7; fi46++ {
        var v47 int = (fi46 | fi46)
        for fi48 := 0; fi48 < 4; fi48++ {
            if (v1 > ("\n\t" + v1)) {
                var v49 int64 = (v40 - -513)
                v50 := GenS3{F0: 42, F1: (v1 + ("a" + v1)), F2: -9223372036854775808}
                if ((!((((v43 / (((-(v43))) | 1)) > 1) && (361 <= 317)) || v3)) && (v40 < -911)) {
                    var v51 uint64 = 771
                    _ = v51
                }
                _ = v49
                _ = v50
            }
            _ = fi48
        }
        _ = fi46
        _ = v47
    }
    v52 := GenS3{F0: 42, F1: "foo bar", F2: v40}
    v53 := GenS4{F0: v40}
    v54 := GenS2{F0: 0, F1: v1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v52
    _ = v53
    _ = v54
}
