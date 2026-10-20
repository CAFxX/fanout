package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint; F1 int8 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 uint8 = 1
        v8 += v8
        v9 := GenS1{F0: 42, F1: v8, F2: 658}
        _ = fi7
        _ = v8
        _ = v9
    }
    if true {
        if false {
            var v10 float64 = (0.0 * 0.1)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 7; fi11++ {
                var v12 int32 = 0
                switch (("\n\t" + "a") + "\n\t") {
                case "foo bar":
                    switch ("foo bar" + (("hello" + "a") + "hello")) {
                    case "a":
                        var v13 int8 = 127
                        _ = v13
                    }
                case "\n\t":
                    var v14 int16 = 258
                    _ = v14
                default:
                    fmt.Println(fi11, fi11)
                }
                var v15 bool = (true || (1.5 != (3.14159 + ((0.1 + (0.1 + (0.1 + 1.5))) / (0.1 - (((-2.25 * -2.25) / 1.0) + 1.5))))))
                _ = fi11
                _ = v12
                _ = v15
            }
        }
    }
    if (566 <= 2) {
        var v16 int32 = -700
        _ = v16
    } else {
        var v17 float32 = ((1.5 / (((1.5 + 3.14159) / (1e10 + 0.1)) + (1.5 + (1.5 / (float32(float64(3.14159)) / 3.14159))))) + float32((1e10 / ((((0.1 - (3.14159 - 1e10)) / 1.0) + (float64((-2.25 - 0.0)) / (float64((0.1 * -2.25)) + float64(0.1)))) + 1e10))))
        _ = v17
    }
    return 2
}
func genFn1() int16 {
    v18 := GenS1{F0: -98, F1: 96, F2: 914}
    var v19 bool = (true && false)
    v20 := []int{1, 896, -1}
    _ = v18
    _ = v19
    _ = v20
    return -495
}

func main() {
    var v1 float64 = 1e10
    var v2 uint = 255
    var v3 int8 = 76
    var v4 uint8 = 42
    var v5 float64 = ((3.14159 * 0.0) * ((v1 - (v1 - ((float64((1.5 / 0.1)) + 0.0) - (v1 + v1)))) + float64(0.1)))
    var v6 int16 = 32767
    if false {
        v21 := GenS1{F0: v3, F1: v4, F2: 32767}
        for fi22 := 0; fi22 < 18; fi22++ {
            if true {
                var v23 uint32 = 2147483647
                _ = v23
            }
            for fi24 := 0; fi24 < 6; fi24++ {
                v25 := []int{fi22}
                _ = fi24
                _ = v25
            }
            v26 := []int{fi22}
            _ = fi22
            _ = v26
        }
        v27 := GenS2{F0: 18446744073709551615}
        _ = v21
        _ = v27
    }
    if true {
        v1 += 1.5
    } else {
        v28 := GenS1{F0: v3, F1: v4, F2: 0}
        v29 := []int{}
        fmt.Println(v2)
        _ = v28
        _ = v29
    }
    if (true && false) {
        v30 := []int{32768}
        v4 += genFn0()
        _ = v30
    }
    v6 = v6
    switch "\n\t" {
    case "\n\t":
        for fi31 := 0; fi31 < 11; fi31++ {
            v3 += (v3 * 1)
            for fi32 := 0; fi32 < 14; fi32++ {
                if false {
                    switch (fi31 << ((660) & 63)) {
                    case 42:
                        var v33 bool = false
                        _ = v33
                    case -2147483648:
                        var v34 int32 = 1
                        _ = v34
                    case -217:
                        v3 = 127
                    default:
                        v3 = 67
                    }
                }
                v35 := GenS2{F0: 543}
                var v36 float64 = 1e10
                v37 := GenS3{F0: false, F1: 0}
                v38 := []int{(fi31 + 131), (-(fi31)), -2147483648, int(v4)}
                var v39 uint8 = v4
                v40 := GenS4{F0: uint64(v39), F1: (v4 + v4)}
                _ = fi32
                _ = v35
                _ = v36
                _ = v37
                _ = v38
                _ = v39
                _ = v40
            }
            _ = fi31
        }
    case "":
        v41 := []int{-128}
        _ = v41
    default:
        v1 += ((0.0 + float64(((1e10 + 0.1) - float32(v1)))) / ((0.0 - (float64(0.1) / (((1.5 / v5) + ((0.1 / (v5 + 1.5)) - 0.0)) - 1e10))) / 1.0))
    }
    v4 -= 42
    v42 := GenS4{F0: 0, F1: (v4 << ((431) & 63))}
    if false {
        fmt.Println(v5, v3, v5)
        v4 = (^(v4))
    } else {
        var v43 int8 = (^(v3))
        for fi44 := 0; fi44 < 2; fi44++ {
            switch (("a" + "\n\t") + (("a" + "foo bar") + (("hello" + ("hello" + "hello")) + ("\n\t" + ("" + "a"))))) {
            case "foo bar":
                v45 := []int{(fi44 & fi44), fi44, fi44, fi44}
                _ = v45
            case "a":
                fmt.Println(fi44, v43)
            }
            _ = fi44
        }
        _ = v43
    }
    var v46 uint32 = 260
    v47 := GenS4{F0: 4294967295, F1: 2}
    var v48 uint16 = 0
    if ((float32(0.1) + 1.5) < 0.1) {
        if ((!(true || false)) || (false || true)) {
            var v49 int32 = 1
            v3 += 127
            var v50 uint64 = 2
            _ = v49
            _ = v50
        }
        var v51 uint8 = v4
        v52 := GenS3{F0: ((v3 + (-(v3))) != 127), F1: 0}
        _ = v51
        _ = v52
    } else {
        v53 := []int{9223372036854775807, 127}
        v54 := GenS1{F0: v3, F1: (^(v4)), F2: v48}
        _ = len(v53)
        _ = v53
        _ = v54
    }
    v2 = (v2 | v2)
    if true {
        v46 -= v46
        switch -107 {
        case 128:
            v55 := []int{9223372036854775807, -508, 128}
            _ = v55
        case 2147483647:
            fmt.Println(v1, v5)
        }
        v56 := []int{-445, int(v48), int(v4)}
        _ = v56
    }
    switch 1 {
    case 4294967295:
        v57 := GenS4{F0: 1, F1: (^(v4))}
        _ = v57
    case 32767:
        var v58 uint16 = (^(v48))
        _ = v58
    case 0:
        v59 := GenS5{F0: v4}
        _ = v59
    default:
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "foo bar":
            for fi60 := 0; fi60 < 17; fi60++ {
                for fi61 := 0; fi61 < 15; fi61++ {
                    fmt.Println(v48, v4)
                    var v62 int64 = -1
                    if false {
                        var v63 uint16 = (v48 + (^(v48)))
                        v4 -= 132
                        _ = v63
                    }
                    _ = fi61
                    _ = v62
                }
                if false {
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v3, v6, v2)
                    case "\n\t":
                        var v64 int8 = -34
                        _ = v64
                    default:
                        v65 := GenS6{F0: (v2 >> (((v2 & v2)) & 63)), F1: 100}
                        _ = v65
                    }
                    for fi66 := 0; fi66 < 9; fi66++ {
                        fmt.Println(v3, v2)
                        v67 := []int{(-(fi60)), (fi60 & (fi66 - fi66)), 1}
                        _ = fi66
                        _ = v67
                    }
                    v68 := GenS2{F0: 0}
                    _ = v68
                }
                _ = fi60
            }
        default:
            var v69 int8 = -48
            _ = v69
        }
    }
    var v70 int16 = (v6 >> ((v2) & 63))
    switch 1 {
    case 4294967295:
        v71 := []int{}
        _ = v71
    case 970:
        var v72 uint32 = 4294967295
        _ = v72
    case 42:
        v70 = (v70 - genFn1())
    default:
        v1 += v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
    _ = v70
}
