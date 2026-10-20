package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint64; F1 string; F2 uint }
type GenS3 struct { F0 int32; F1 int8; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 uint8 }
func genFn0() uint16 {
    v4 := GenS1{F0: -819, F1: 650}
    var v5 int8 = 100
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 int, p1 int16) int16 {
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(fi6)
        p1 = 1000
        fmt.Println(fi6, fi6, fi6)
        _ = fi6
    }
    if false {
        if false {
            if (9223372036854775807 < -9223372036854775808) {
                for fi7 := 0; fi7 < 1; fi7++ {
                    v8 := GenS1{F0: -32768, F1: 270}
                    _ = fi7
                    _ = v8
                }
                p0 = p0
            }
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS2{F0: 431, F1: "a", F2: 255}
                _ = fi9
                _ = v10
            }
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    p0 -= (p0 ^ p0)
                    _ = fi12
                }
                var v13 int8 = 1
                _ = fi11
                _ = v13
            }
            if true {
                switch ("hello" + "\n\t") {
                case "a":
                    fmt.Println(p0, p1)
                case "foo bar":
                    v14 := []int{}
                    _ = v14
                case "\n\t":
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(p0, p0)
                    case "":
                        v15 := []int{}
                        _ = v15
                    }
                }
            }
        }
        p1 = p1
    } else {
        var v16 int8 = 100
        _ = v16
    }
    var v17 int16 = 568
    switch ("hello" + (((("a" + "foo bar") + "") + "hello") + "\n\t")) {
    case "":
        p1 += (p1 + v17)
    case "hello":
        var v18 float32 = 1e10
        _ = v18
    }
    _ = p0
    _ = p1
    _ = v17
    return v17
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 uint16 = 32767
    v19 := GenS2{F0: 0, F1: "", F2: 1}
    if (-1 <= v1) {
        v3 = genFn0()
        if false {
            fmt.Println(v1)
        }
        v1 += v1
    } else {
        for fi20 := 0; fi20 < 9; fi20++ {
            var v21 int8 = v1
            switch fi20 {
            case -2147483648:
                v1 += v21
            case 930:
                for fi22 := 0; fi22 < 15; fi22++ {
                    var v23 uint64 = 0
                    _ = fi22
                    _ = v23
                }
            }
            v1 += v21
            _ = fi20
            _ = v21
        }
    }
    v3 = (v3 | 42)
    v24 := GenS3{F0: 2147483647, F1: (v1 % (((v1 >> ((368) & 63))) | 1)), F2: -1}
    v25 := GenS2{F0: 604, F1: "", F2: 596}
    for fi26 := 0; fi26 < 19; fi26++ {
        v27 := []int{}
        v3 = (v3 * v3)
        for fi28 := 0; fi28 < 9; fi28++ {
            v1 = (v1 ^ v1)
            v1 = (v1 - v1)
            _ = fi28
        }
        _ = fi26
        _ = v27
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        var v30 uint64 = 1
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        fmt.Println(fi31, fi31, v1)
        v3 = (^(v3))
        _ = fi31
    }
    if false {
        var v32 int64 = 0
        v33 := GenS2{F0: 409, F1: ("foo bar" + "a"), F2: 1}
        var v34 int64 = v32
        _ = v32
        _ = v33
        _ = v34
    }
    if ((65535 <= 2147483648) || true) {
        switch ("foo bar" + "") {
        case "foo bar":
            if ("foo bar" < ("foo bar" + "foo bar")) {
                for fi35 := 0; fi35 < 20; fi35++ {
                    if false {
                        var v36 uint = 18446744073709551615
                        v37 := []int{}
                        _ = v36
                        _ = v37
                    } else {
                        fmt.Println(v1, fi35)
                        var v38 int16 = -1
                        switch "foo bar" {
                        case "":
                            for fi39 := 0; fi39 < 17; fi39++ {
                                v3 -= (^(v3))
                                _ = fi39
                            }
                        case "hello":
                            v38 = 1000
                        case "a":
                            v2[""] = ("hello" + "\n\t")
                        default:
                            v3 = 65535
                        }
                        _ = v38
                    }
                    var v40 uint8 = 42
                    var v41 uint32 = 1
                    _ = fi35
                    _ = v40
                    _ = v41
                }
            } else {
                var v42 int8 = v1
                v2[((("hello" + ("foo bar" + "")) + ("a" + ("" + "hello"))) + "")] = ("\n\t" + ("\n\t" + ("" + (("hello" + ("" + "\n\t")) + ""))))
                fmt.Println(v1, v42)
                _ = v42
            }
        }
        for fi43 := 0; fi43 < 3; fi43++ {
            if true {
                v44 := GenS4{F0: 4294967295}
                _ = len(v2)
                _ = v44
            }
            if (384 <= 9) {
                var v45 int64 = 42
                var v46 int16 = genFn1(fi43, 32767)
                _ = v45
                _ = v46
            } else {
                switch 32767 {
                case 127:
                    for fi47 := 0; fi47 < 12; fi47++ {
                        if (!false) {
                            v48 := GenS1{F0: -66, F1: 394}
                            _ = v48
                        }
                        v1 -= v1
                        _ = fi47
                    }
                case 42:
                    var v49 uint64 = 36
                    _ = v49
                default:
                    _ = len(v2)
                }
            }
            _ = fi43
        }
        var v50 int64 = -60
        _ = v50
    } else {
        var v51 int64 = 9223372036854775807
        _ = v51
    }
    var v52 uint8 = 1
    for fi53 := 0; fi53 < 20; fi53++ {
        var v54 float64 = ((3.14159 + (0.0 + 0.0)) * (0.1 / 1.5))
        if (76 > 600) {
            v55 := GenS5{F0: v52, F1: genFn1(127, -264), F2: v52}
            _ = v55
        }
        _ = fi53
        _ = v54
    }
    var v56 uint16 = genFn0()
    var v57 uint64 = 4294967295
    v1 -= 0
    switch -1 {
    case 32768:
        var v58 uint8 = 19
        _ = v58
    default:
        v57 -= 246
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v52
    _ = v56
    _ = v57
}
