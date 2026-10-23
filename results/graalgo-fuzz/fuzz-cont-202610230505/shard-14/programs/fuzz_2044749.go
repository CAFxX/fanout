package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 string; F2 int }
type GenS5 struct { F0 int; F1 uint64 }
func genFn0(p0 bool, p1 int, p2 string) int {
    v7 := []int{p1, p1, (p1 >> ((4294967295) & 63)), 9223372036854775807}
    switch p1 {
    case 42:
        switch ("hello" + "") {
        case "foo bar":
            v8 := GenS1{F0: 1, F1: (127 <= 0), F2: 0}
            _ = v8
        case "hello":
            v9 := GenS1{F0: 0, F1: p0, F2: 17}
            _ = v9
        }
    case -895:
        if (false || (127 > 0)) {
            for fi10 := 0; fi10 < 14; fi10++ {
                p1 = (p1 & fi10)
                _ = fi10
            }
            var v11 float64 = -2.25
            _ = v11
        } else {
            switch (p1 + 0) {
            case -481:
                p0 = p0
            case 2147483647:
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int32 = 42
                    var v14 uint64 = 820
                    fmt.Println(p1, fi12, fi12)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            v15 := GenS1{F0: 18446744073709551615, F1: (-643 == 1), F2: 100}
            if p0 {
                if len(v7) > 0 {
                    v7[0] = 9223372036854775807
                }
                v16 := GenS2{F0: 699, F1: -128}
                if len(v7) > 0 {
                    v7[0] = (p1 + p1)
                }
                _ = v16
            }
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int32, p1 uint) int {
    if (42 != 1) {
        p0 = (-(p0))
        v17 := []int{}
        _ = v17
    } else {
        v18 := GenS1{F0: 4294967295, F1: ((65535 != 750) || (!true)), F2: 0}
        v19 := []int{-1, -9223372036854775808, 127, 4294967295}
        if (true || (-739 > p0)) {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            p0 = 1
            switch "foo bar" {
            case "":
                if (2 <= 642) {
                    if (0 != -128) {
                        var v20 int32 = 621
                        if ((!((1.5 == (((float64(1e10) - 1.5) + 3.14159) + 1.5)) && true)) || (22 != 100)) {
                            p0 += (p0 << ((p1) & 63))
                        } else {
                            v21 := GenS3{F0: (p0 | p0)}
                            _ = v21
                        }
                        var v22 int32 = (p0 ^ (p0 + (p0 * (v20 % ((p0) | 1)))))
                        _ = v20
                        _ = v22
                    }
                    var v23 uint16 = 681
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(v23, p0)
                    default:
                        fmt.Println(v23)
                    }
                    _ = v23
                } else {
                    switch ("\n\t" + ("a" + "\n\t")) {
                    case "hello":
                        var v24 bool = false
                        _ = v24
                    case "":
                        var v25 bool = true
                        _ = v25
                    default:
                        for fi26 := 0; fi26 < 20; fi26++ {
                            p1 -= p1
                            _ = fi26
                        }
                    }
                    var v27 uint64 = 0
                    _ = v27
                }
            }
        }
        _ = v18
        _ = v19
    }
    p0 -= (p0 >> ((p1) & 63))
    for fi28 := 0; fi28 < 19; fi28++ {
        v29 := []int{}
        _ = fi28
        _ = v29
    }
    if true {
        if (255 < 65535) {
            switch "\n\t" {
            case "\n\t":
                switch "foo bar" {
                case "":
                    v30 := []int{}
                    _ = v30
                case "foo bar":
                    fmt.Println(p0, p0)
                case "a":
                    p0 += p0
                default:
                    p1 = (p1 & (^(p1)))
                }
            case "a":
                var v31 int = 9223372036854775807
                _ = v31
            case "foo bar":
                var v32 uint32 = 422
                _ = v32
            default:
                var v33 uint16 = 32767
                _ = v33
            }
            p0 = p0
        }
        var v34 bool = false
        _ = v34
    } else {
        v35 := GenS3{F0: p0}
        for fi36 := 0; fi36 < 9; fi36++ {
            if true {
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 int8 = 3
                    var v39 uint8 = uint8(p1)
                    p1 = p1
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            } else {
                p1 = p1
            }
            fmt.Println(p1, fi36, fi36)
            _ = fi36
        }
        _ = v35
    }
    var v40 float64 = 1e10
    _ = p0
    _ = p1
    _ = v40
    return 127
}

func main() {
    var v1 int32 = -1
    var v2 int = 127
    var v3 int = (v2 - v2)
    v4 := make(map[int]uint32)
    v5 := []uint32{974, 2147483648, 2147483647}
    var v6 float32 = (1.5 / 3.14159)
    v3 += v2
    v3 = -845
    var v41 float32 = float32(-2.25)
    var v42 string = ("\n\t" + ("" + "hello"))
    v6 = float32(3.14159)
    fmt.Println(v3)
    var v43 int = v3
    delete(v4, (v43 ^ v3))
    v2 -= (v43 / ((-9223372036854775808) | 1))
    fmt.Println(v2, v1)
    if true {
        switch "hello" {
        case "hello":
            for fi44 := 0; fi44 < 7; fi44++ {
                v45 := GenS4{F0: (-(v3)), F1: v42, F2: genFn1((v1 + v1), 0)}
                _ = fi44
                _ = v45
            }
        case "a":
            var v46 float32 = (0.0 - 1.5)
            _ = v46
        case "\n\t":
            var v47 uint16 = 0
            _ = v47
        default:
            if false {
                fmt.Println(v42, v6, v6)
                var v48 float64 = ((1e10 * 0.0) * (1e10 + 0.1))
                _ = v48
            }
        }
        for fi49 := 0; fi49 < 15; fi49++ {
            for fi50 := 0; fi50 < 9; fi50++ {
                var v51 int32 = v1
                fmt.Println(v42, v43)
                if (!((0.0 == 0.0) || (true && (!false)))) {
                    var v52 int16 = -1
                    if true {
                        var v53 float64 = -2.25
                        v6 += v6
                        var v54 uint = 993
                        _ = v53
                        _ = v54
                    }
                    var v55 uint8 = 128
                    v56 := GenS5{F0: (fi50 - (fi50 + (fi50 << ((1) & 63)))), F1: 2}
                    var v57 int16 = -32768
                    _ = v52
                    _ = v55
                    _ = v56
                    _ = v57
                } else {
                    v3 = (fi50 | v3)
                    fmt.Println(v2, v1, fi50)
                    v51 -= (-(v51))
                    switch v3 {
                    case -541:
                        v42 = v42
                    case 32768:
                        var v58 uint16 = 0
                        _ = v58
                    case 1:
                        v3 = v43
                    default:
                        v1 = 1
                    }
                    switch genFn1(0, 1) {
                    case -344:
                        var v59 uint32 = 65535
                        _ = v59
                    default:
                        if true {
                            var v60 uint64 = 908
                            _ = v60
                        }
                    }
                    delete(v4, 9223372036854775807)
                }
                _ = fi50
                _ = v51
            }
            v61 := v4[(fi49 & genFn0(true, 42, "\n\t"))]
            fmt.Println(v2, v42, v43)
            _ = fi49
            _ = v61
        }
    }
    v1 += (-(v1))
    if (v1 < (^(v1))) {
        var v62 int8 = -128
        for fi63 := 0; fi63 < 4; fi63++ {
            var v64 uint32 = 876
            switch (v42 + ("\n\t" + v42)) {
            case "a":
                _ = len(v4)
            case "foo bar":
                var v65 int32 = v1
                _ = v65
            default:
                switch "\n\t" {
                case "a":
                    v64 = (^(v64))
                case "\n\t":
                    for fi66 := 0; fi66 < 3; fi66++ {
                        v67 := GenS3{F0: (-(v1))}
                        var v68 string = ("foo bar" + v42)
                        _ = fi66
                        _ = v67
                        _ = v68
                    }
                default:
                    _ = len(v4)
                }
            }
            var v69 uint64 = 42
            _ = fi63
            _ = v64
            _ = v69
        }
        _ = v62
    } else {
        v2 -= v3
    }
    fmt.Println(v43, v1)
    v42 = v42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
