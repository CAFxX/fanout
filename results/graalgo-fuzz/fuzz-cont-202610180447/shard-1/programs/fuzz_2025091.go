package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint16; F1 string; F2 uint8 }
func genFn0(p0 uint16, p1 uint32) int32 {
    var v4 string = "a"
    v5 := []int{-122, 2147483647, 0}
    fmt.Println(p1, v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -1
}
func genFn1(p0 int, p1 int64, p2 uint) int64 {
    if true {
        v6 := GenS1{F0: (^(p2)), F1: p2, F2: 415}
        _ = v6
    }
    if (255 == 255) {
        p0 += p0
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{p0, (p0 >> ((1) & 63)), (fi8 * fi8)}
                p0 = (fi7 + fi8)
                _ = len(v9)
                _ = fi8
                _ = v9
            }
            v10 := GenS2{F0: 778, F1: 42}
            _ = fi7
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 13; fi11++ {
            v12 := []int{p0}
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
            var v13 int64 = (p1 / ((-9223372036854775808) | 1))
            _ = fi11
            _ = v12
            _ = v13
        }
        fmt.Println(p0)
        for fi14 := 0; fi14 < 18; fi14++ {
            v15 := GenS3{F0: uint32(p2)}
            _ = fi14
            _ = v15
        }
    }
    v16 := []int{p0, (p0 >> (((p2 % ((p2) | 1))) & 63)), 9223372036854775807, p0}
    fmt.Println(p1, p0)
    p2 = (p2 | 255)
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return (p1 ^ p1)
}

func main() {
    v1 := make(map[int]bool)
    var v2 float64 = (((0.0 + 0.0) / (1e10 + 0.0)) * 1.5)
    v3 := make(map[int]bool)
    for fi17 := 0; fi17 < 11; fi17++ {
        v2 += -2.25
        if false {
            v18 := GenS2{F0: 4294967295, F1: 255}
            var v19 int32 = 1
            _ = v18
            _ = v19
        }
        _ = fi17
    }
    if (42 == 2) {
        for fi20 := 0; fi20 < 18; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 int8 = -128
                v23 := GenS4{F0: 1}
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = fi20
        }
        if true {
            switch -2147483648 {
            case 1:
                var v24 uint16 = 1
                _ = v24
            case 32767:
                fmt.Println(v2)
            case 127:
                for fi25 := 0; fi25 < 12; fi25++ {
                    v26 := GenS3{F0: 42}
                    _ = fi25
                    _ = v26
                }
            default:
                switch "\n\t" {
                case "":
                    v27 := GenS2{F0: 1, F1: 4294967295}
                    _ = v27
                default:
                    if false {
                        var v28 bool = (4294967295 == 42)
                        _ = v28
                    } else {
                        v2 = 1e10
                        v2 = v2
                        var v29 int16 = 0
                        _ = v29
                    }
                }
            }
        }
    } else {
        v2 = v2
    }
    v30 := []int{-431, -132, 32767}
    v2 = v2
    fmt.Println(v2)
    var v31 uint = 1
    if true {
        if len(v30) > 0 {
            v30[0] = -2147483648
        }
        v31 -= v31
        var v32 int64 = -196
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v34 := v1[57]
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 int64 = 9223372036854775807
            if (v34 && (51 != -128)) {
                var v37 int64 = 666
                v34 = v34
                _ = v37
            } else {
                v38 := GenS1{F0: (^(v31)), F1: 1, F2: 65535}
                switch (fi35 << ((v31) & 63)) {
                case 42:
                    switch ("a" + "foo bar") {
                    case "hello":
                        v2 = (0.1 * v2)
                    case "foo bar":
                        v36 = (v36 * 9223372036854775807)
                    default:
                        var v39 uint8 = 0
                        _ = v39
                    }
                case 863:
                    v36 = (v36 | 294)
                default:
                    v31 += v31
                }
                var v40 string = "hello"
                _ = v38
                _ = v40
            }
            _ = fi35
            _ = v36
        }
        v41 := GenS2{F0: 749, F1: v31}
        _ = fi33
        _ = v34
        _ = v41
    }
    _ = len(v30)
    for fi42 := 0; fi42 < 9; fi42++ {
        switch fi42 {
        case -1:
            if (42 >= 2147483647) {
                v43 := GenS5{F0: v31, F1: -867}
                if ((0.1 + 0.0) <= 3.14159) {
                    var v44 int16 = 277
                    var v45 uint16 = 255
                    _ = v44
                    _ = v45
                }
                _ = v43
            }
        case 0:
            v1[-1] = true
        }
        if (!(84 <= 207)) {
            v46 := GenS6{F0: genFn0(134, 0), F1: 2}
            var v47 float32 = 0.0
            _ = v46
            _ = v47
        } else {
            var v48 uint8 = 1
            switch (("hello" + (("" + ("hello" + ("" + "hello"))) + "\n\t")) + "foo bar") {
            case "a":
                switch "a" {
                case "a":
                    var v49 int64 = -1
                    _ = v49
                case "":
                    for fi50 := 0; fi50 < 1; fi50++ {
                        var v51 int32 = 1
                        if (true && (false && true)) {
                            var v52 float64 = (v2 - (1.5 - (((0.1 * 3.14159) + (-2.25 * (float64(3.14159) * (((v2 + v2) * (0.0 * float64(((-2.25 - 1e10) + (-2.25 / (1e10 * 1e10)))))) - v2)))) * float64(0.1))))
                            delete(v3, fi42)
                            _ = v52
                        }
                        _ = fi50
                        _ = v51
                    }
                }
            }
            _ = v48
        }
        var v53 int8 = -121
        _ = fi42
        _ = v53
    }
    if true {
        for fi54 := 0; fi54 < 19; fi54++ {
            switch (fi54 % ((0) | 1)) {
            case 42:
                if len(v30) >= 2 {
                    v30 = v30[1:len(v30)]
                }
            }
            var v55 int32 = -2147483648
            _ = fi54
            _ = v55
        }
    } else {
        if len(v30) > 0 {
            v30[0] = 876
        }
        var v56 float64 = v2
        v57 := GenS4{F0: genFn0(1, 65535)}
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 7; fi58++ {
        var v59 float32 = (3.14159 + (-2.25 / float32(v2)))
        for fi60 := 0; fi60 < 5; fi60++ {
            switch fi60 {
            case -316:
                _ = len(v3)
            }
            var v61 int32 = 822
            _ = fi60
            _ = v61
        }
        v62 := GenS6{F0: 2147483647, F1: 1}
        _ = fi58
        _ = v59
        _ = v62
    }
    var v63 float32 = 1.5
    v64 := GenS7{F0: 49, F1: "\n\t", F2: 42}
    v63 = 0.0
    v31 -= v31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v63
    _ = v64
}
