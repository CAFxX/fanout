package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int32; F2 int16 }
type GenS5 struct { F0 int16; F1 int8 }
type GenS6 struct { F0 int8; F1 string; F2 int16 }
type GenS7 struct { F0 uint8; F1 uint }
type GenS8 struct { F0 uint; F1 uint32; F2 int8 }
type GenS9 struct { F0 int64; F1 bool; F2 bool }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 bool; F1 uint8; F2 int }
func genFn0(p0 int, p1 int32, p2 int64) string {
    for fi4 := 0; fi4 < 11; fi4++ {
        var v5 uint8 = 69
        p1 += (-(p1))
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 uint8 = 122
        v8 := GenS1{F0: 4294967295}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 18446744073709551615
    v10 := []int{p0}
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return ("foo bar" + (("" + "a") + ""))
}
func genFn1(p0 uint16, p1 bool, p2 uint8) bool {
    if (!(-329 != 128)) {
        for fi11 := 0; fi11 < 14; fi11++ {
            for fi12 := 0; fi12 < 17; fi12++ {
                fmt.Println(p0)
                _ = fi12
            }
            switch ("" + "") {
            case "\n\t":
                if p1 {
                    var v13 bool = true
                    _ = v13
                }
            case "foo bar":
                var v14 bool = (p1 && p1)
                _ = v14
            }
            _ = fi11
        }
        v15 := GenS2{F0: 4294967295, F1: 65535}
        var v16 int16 = 0
        _ = v15
        _ = v16
    }
    if (((0.0 + 0.1) != (0.0 + 0.1)) || ((2147483647 == -457) || (!(!p1)))) {
        v17 := GenS3{F0: 711}
        _ = v17
    } else {
        p0 = 255
    }
    if true {
        for fi18 := 0; fi18 < 18; fi18++ {
            switch "hello" {
            case "":
                p2 -= p2
            case "\n\t":
                var v19 bool = p1
                _ = v19
            }
            var v20 uint64 = 2
            _ = fi18
            _ = v20
        }
        p2 = p2
        fmt.Println(p2, p2, p0)
    }
    p1 = (true && (!(true && p1)))
    p2 += (p2 | (^(p2)))
    _ = p0
    _ = p1
    _ = p2
    return (p2 == (^(p2)))
}

func main() {
    var v1 float32 = (1.5 + 0.0)
    var v2 uint64 = 612
    var v3 string = "a"
    for fi21 := 0; fi21 < 13; fi21++ {
        v3 = "hello"
        for fi22 := 0; fi22 < 8; fi22++ {
            v23 := GenS4{F0: -364, F1: 42, F2: 1000}
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    switch 128 {
    case 0:
        v2 -= (v2 * 1)
    }
    v24 := GenS5{F0: 42, F1: 127}
    if (((true || false) || false) && (9223372036854775807 < 128)) {
        for fi25 := 0; fi25 < 1; fi25++ {
            var v26 uint16 = 0
            var v27 uint8 = 100
            fmt.Println(v3)
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    switch "a" {
    case "hello":
        v28 := GenS6{F0: 127, F1: (v3 + v3), F2: int16(v2)}
        _ = v28
    case "foo bar":
        if (!true) {
            var v29 uint64 = (v2 + (^(v2)))
            if (255 == 63) {
                v30 := []int{911, 370}
                var v31 int16 = -341
                _ = v30
                _ = v31
            }
            _ = v29
        } else {
            for fi32 := 0; fi32 < 3; fi32++ {
                switch v3 {
                case "\n\t":
                    v33 := []int{fi32, (fi32 - 9223372036854775807)}
                    _ = v33
                default:
                    fmt.Println(v1, fi32)
                }
                _ = fi32
            }
        }
    default:
        for fi34 := 0; fi34 < 15; fi34++ {
            for fi35 := 0; fi35 < 3; fi35++ {
                if ((1.5 - 1e10) < float64((0.1 + v1))) {
                    v36 := []int{(fi34 % ((4294967295) | 1))}
                    v37 := GenS7{F0: 2, F1: 2}
                    v3 = v3
                    v38 := GenS4{F0: fi34, F1: -1, F2: 940}
                    var v39 int64 = 1
                    v39 = (v39 / ((v39) | 1))
                    _ = v36
                    _ = v37
                    _ = v38
                    _ = v39
                } else {
                    var v40 string = (v3 + ((v3 + v3) + "hello"))
                    v1 -= v1
                    if (true || ((v2 / ((18446744073709551615) | 1)) != 18446744073709551615)) {
                        v1 -= (float32(((1.5 * (0.1 * float64(-2.25))) + 0.0)) / ((v1 + v1) - (0.1 * (1.5 - v1))))
                        switch "foo bar" {
                        case "hello":
                            var v41 int16 = 0
                            _ = v41
                        case "":
                            v42 := []int{2147483647, (fi35 % (((fi35 / ((fi35) | 1))) | 1))}
                            _ = v42
                        case "\n\t":
                            v43 := []int{(fi34 / ((32768) | 1)), 0, (fi34 + -583), 539}
                            _ = v43
                        default:
                            v40 = v40
                        }
                    } else {
                        v44 := []int{fi35, (fi35 + 0)}
                        v3 = "foo bar"
                        _ = v44
                    }
                    _ = v40
                }
                v1 = v1
                v45 := GenS8{F0: 4294967295, F1: 481, F2: 1}
                _ = fi35
                _ = v45
            }
            var v46 uint64 = 2
            _ = fi34
            _ = v46
        }
    }
    switch -2147483648 {
    case 9223372036854775807:
        switch "hello" {
        case "hello":
            v1 += v1
        case "":
            var v47 float32 = float32(3.14159)
            _ = v47
        case "a":
            v3 = v3
        default:
            for fi48 := 0; fi48 < 18; fi48++ {
                fmt.Println(v2, fi48, v3)
                var v49 uint16 = 0
                _ = fi48
                _ = v49
            }
        }
    }
    fmt.Println(v2, v3)
    v50 := []int{32768, -34, 800, 32768}
    v51 := GenS9{F0: 0, F1: (255 == 455), F2: false}
    v3 = ("\n\t" + ((v3 + v3) + v3))
    v3 = (v3 + v3)
    switch -112 {
    case -1:
        for fi52 := 0; fi52 < 20; fi52++ {
            if (201 > 238) {
                v53 := GenS4{F0: fi52, F1: -2147483648, F2: -32768}
                _ = v53
            }
            _ = fi52
        }
    case 32767:
        switch ("foo bar" + "hello") {
        case "":
            fmt.Println(v1)
        case "hello":
            v54 := GenS10{F0: 0, F1: uint8(v2)}
            _ = v54
        case "a":
            if false {
                _ = len(v50)
                v55 := GenS11{F0: false, F1: 129, F2: 32767}
                for fi56 := 0; fi56 < 11; fi56++ {
                    var v57 int16 = -921
                    _ = fi56
                    _ = v57
                }
                _ = v55
            }
        }
    case -69:
        if len(v50) > 0 {
            v50[0] = 9223372036854775807
        }
    }
    fmt.Println(v3)
    var v58 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v50
    _ = v51
    _ = v58
}
