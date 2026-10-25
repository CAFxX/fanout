package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 uint }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16 }
func genFn0() int32 {
    var v4 int = 0
    for fi5 := 0; fi5 < 11; fi5++ {
        switch 214 {
        case -1:
            v6 := GenS1{F0: 42, F1: 988, F2: 65535}
            _ = v6
        default:
            var v7 int32 = 42
            _ = v7
        }
        _ = fi5
    }
    var v8 uint8 = 162
    if true {
        v9 := GenS1{F0: 1, F1: 905, F2: 4294967295}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 10; fi10++ {
            v11 := []int{}
            if (416 != -973) {
                var v12 uint8 = (v8 * v8)
                var v13 bool = ((2147483648 < 919) || false)
                if v13 {
                    if (!((!(!v13)) && v13)) {
                        var v14 uint = 119
                        if (18446744073709551615 >= (v14 ^ v14)) {
                            v15 := GenS1{F0: 0, F1: 32767, F2: v14}
                            if len(v11) > 0 {
                                v11[0] = 891
                            }
                            _ = v15
                        }
                        v16 := GenS2{F0: -1, F1: ("a" + "a")}
                        _ = v14
                        _ = v16
                    } else {
                        var v17 int16 = -1
                        if len(v11) > 0 {
                            v11[0] = (v4 * (fi10 >> ((65535) & 63)))
                        }
                        _ = v17
                    }
                    if v13 {
                        _ = len(v11)
                        v11 = append(v11, 42)
                    } else {
                        _ = len(v11)
                    }
                } else {
                    fmt.Println(fi10)
                    v18 := GenS3{F0: -763, F1: 239, F2: (v12 + v8)}
                    if (115 > v12) {
                        fmt.Println(v8, v12, v8)
                        _ = len(v11)
                    }
                    _ = v18
                }
                _ = v12
                _ = v13
            } else {
                if true {
                    _ = len(v11)
                    v8 -= 42
                    var v19 float64 = 0.1
                    _ = v19
                } else {
                    fmt.Println(v4, fi10)
                    for fi20 := 0; fi20 < 20; fi20++ {
                        var v21 int64 = 0
                        fmt.Println(v21, v8)
                        _ = fi20
                        _ = v21
                    }
                }
                var v22 uint8 = (^(v8))
                _ = v22
            }
            v23 := GenS3{F0: 843, F1: 384, F2: v8}
            _ = fi10
            _ = v11
            _ = v23
        }
        var v24 uint = 65535
        for fi25 := 0; fi25 < 15; fi25++ {
            v4 += (-(fi25))
            _ = fi25
        }
        _ = v24
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        var v27 bool = true
        if (v27 || v27) {
            if false {
                for fi28 := 0; fi28 < 20; fi28++ {
                    switch "foo bar" {
                    case "\n\t":
                        v29 := []int{(fi28 | -9223372036854775808)}
                        _ = v29
                    case "hello":
                        var v30 float64 = 0.1
                        _ = v30
                    case "":
                        switch fi26 {
                        case 2147483647:
                            switch "a" {
                            case "a":
                                if v27 {
                                    v31 := []int{-2147483648}
                                    v27 = (!v27)
                                    _ = v31
                                }
                            }
                            v27 = true
                        default:
                            switch fi26 {
                            case 2147483647:
                                var v32 int16 = 1000
                                _ = v32
                            case 452:
                                var v33 int16 = 42
                                _ = v33
                            default:
                                if (5 == 255) {
                                    if (((v27 || (18446744073709551615 <= 2)) || (!v27)) && (1000 <= -441)) {
                                        switch fi26 {
                                        case 9223372036854775807:
                                            v34 := []int{fi28, (v4 ^ (fi28 / (((fi26 / ((fi26) | 1))) | 1)))}
                                            _ = v34
                                        case 127:
                                            var v35 float32 = float32(-2.25)
                                            _ = v35
                                        }
                                        v8 += v8
                                        var v36 int16 = -1
                                        _ = v36
                                    } else {
                                        v37 := GenS4{F0: -337}
                                        v38 := GenS3{F0: 1, F1: 0, F2: v8}
                                        v39 := []int{fi28, (fi26 | (fi28 >> ((65535) & 63))), fi26}
                                        _ = v37
                                        _ = v38
                                        _ = v39
                                    }
                                }
                            }
                        }
                    default:
                        var v40 uint16 = uint16(v8)
                        _ = v40
                    }
                    fmt.Println(fi26, v8)
                    switch fi26 {
                    case -9223372036854775808:
                        var v41 uint64 = 2
                        _ = v41
                    case 0:
                        var v42 uint32 = 724
                        _ = v42
                    case 1:
                        v27 = v27
                    default:
                        switch ("a" + "hello") {
                        case "a":
                            v43 := GenS5{F0: 133}
                            _ = v43
                        }
                    }
                    _ = fi28
                }
                fmt.Println(fi26, v8)
            }
            var v44 int8 = 0
            _ = v44
        }
        _ = fi26
        _ = v27
    }
    _ = v4
    _ = v8
    return 0
}
func genFn1(p0 bool) uint64 {
    p0 = (true || (!(p0 && p0)))
    var v45 uint32 = 2147483648
    p0 = p0
    v46 := []int{32768}
    _ = p0
    _ = v45
    _ = v46
    return 4294967295
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 234
    var v3 bool = false
    var v47 int32 = -2147483648
    var v48 uint8 = 94
    fmt.Println(v48, v2)
    var v49 bool = true
    switch "a" {
    case "hello":
        fmt.Println(v49, v47)
    }
    v50 := []int{1, 127}
    for fi51 := 0; fi51 < 3; fi51++ {
        var v52 int = (-(fi51))
        fmt.Println(fi51, v2)
        _ = fi51
        _ = v52
    }
    _ = len(v50)
    v53 := GenS2{F0: 9223372036854775807, F1: ("hello" + "hello")}
    if v49 {
        fmt.Println(v3)
    }
    var v54 float64 = 1e10
    var v55 float64 = v54
    v3 = true
    var v56 int64 = 42
    v57 := GenS5{F0: 576}
    if (v3 && v49) {
        v56 += 0
    }
    fmt.Println(v55)
    var v58 float64 = (float64(1e10) * 0.0)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("slicelen v50:", len(v50), v50)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v47
    _ = v48
    _ = v49
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
}
