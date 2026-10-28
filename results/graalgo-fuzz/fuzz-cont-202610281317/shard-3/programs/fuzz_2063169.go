package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int64; F2 uint64 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint8 }
func genFn0(p0 uint, p1 uint32) uint64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        p1 = (^(p1))
        _ = fi4
    }
    var v5 uint = p0
    _ = p0
    _ = p1
    _ = v5
    return 480
}
func genFn1(p0 string, p1 bool, p2 uint8) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 uint16 = 841
        switch fi6 {
        case 127:
            p0 = p0
        case 42:
            v7 -= (^(v7))
        default:
            for fi8 := 0; fi8 < 16; fi8++ {
                if (951 <= 0) {
                    if p1 {
                        p2 += p2
                        var v9 float64 = (1.5 * 0.0)
                        v7 += (v7 - 255)
                        _ = v9
                    }
                } else {
                    var v10 uint = 255
                    var v11 int64 = 42
                    var v12 int16 = -454
                    _ = v10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 1}
                v7 = (v7 - uint16(p2))
                if true {
                    if p1 {
                        fmt.Println(v7, p2)
                        v14 := GenS1{F0: (v7 ^ 848)}
                        v15 := []int{fi6}
                        _ = v14
                        _ = v15
                    }
                    p1 = (1 <= 728)
                    p2 -= p2
                    var v16 bool = false
                    switch p0 {
                    case "hello":
                        v17 := []int{fi6, fi6}
                        _ = v17
                    }
                    var v18 uint = 18446744073709551615
                    var v19 float64 = 1e10
                    v20 := []int{}
                    fmt.Println(v19, p0, v18)
                    v20 = append(v20, (fi8 + fi8))
                    if (v19 <= v19) {
                        var v21 uint = v18
                        var v22 int8 = 1
                        var v23 bool = true
                        v7 = (v7 + (v7 + uint16(v21)))
                        _ = v21
                        _ = v22
                        _ = v23
                    } else {
                        p2 = (p2 >> ((v18) & 63))
                        var v24 uint16 = (^(v7))
                        if true {
                            p2 += (^(p2))
                            fmt.Println(p0)
                        } else {
                            v25 := GenS2{F0: p2, F1: (^(v24)), F2: 0}
                            switch ("" + "\n\t") {
                            case "hello":
                                var v26 int64 = -1
                                _ = v26
                            case "foo bar":
                                v27 := GenS3{F0: 0}
                                _ = v27
                            default:
                                fmt.Println(p0, p1, v7)
                            }
                            _ = v25
                        }
                        _ = v24
                    }
                    _ = len(v20)
                    switch "" {
                    case "foo bar":
                        v20 = append(v20, fi8)
                    case "hello":
                        var v28 uint64 = genFn0(v18, 223)
                        var v29 bool = (127 >= 100)
                        _ = v28
                        _ = v29
                    }
                    _ = v16
                    _ = v18
                    _ = v19
                    _ = v20
                }
                fmt.Println(p2)
                switch (fi6 >> ((65535) & 63)) {
                case 2147483647:
                    v7 = v7
                case 0:
                    var v30 int8 = -1
                    _ = v30
                case 127:
                    if p1 {
                        switch "hello" {
                        case "a":
                            fmt.Println(p0, p2, fi8)
                        case "foo bar":
                            if false {
                                var v31 uint8 = 2
                                _ = v31
                            } else {
                                v32 := []int{(fi8 % ((fi6) | 1))}
                                v33 := GenS4{F0: 575, F1: 642, F2: 4294967295}
                                p1 = (42 >= 0)
                                if p1 {
                                    var v34 int = 127
                                    _ = v34
                                }
                                var v35 int32 = -832
                                _ = v32
                                _ = v33
                                _ = v35
                            }
                        }
                        p1 = (!(p1 || true))
                        v36 := GenS1{F0: 671}
                        _ = v36
                    }
                    var v37 int32 = -2147483648
                    var v38 int8 = 1
                    _ = v37
                    _ = v38
                default:
                    p1 = (!(!false))
                }
                v39 := []int{}
                _ = fi8
                _ = v13
                _ = v39
            }
        }
        _ = fi6
        _ = v7
    }
    p2 -= 42
    if p1 {
        var v40 int16 = 268
        _ = v40
    }
    var v41 int32 = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v41
    return 42
}

func main() {
    var v1 int8 = 28
    var v2 float64 = 1.5
    var v3 int8 = -128
    for fi42 := 0; fi42 < 3; fi42++ {
        v3 = (v3 >> ((255) & 63))
        if (!false) {
            v2 = float64((-2.25 + (1e10 - 1e10)))
        } else {
            fmt.Println(v2, v1)
        }
        v43 := []int{(fi42 << ((18446744073709551615) & 63)), fi42, (fi42 - (^(fi42)))}
        _ = fi42
        _ = v43
    }
    if (18446744073709551615 < 2) {
        v2 = (v2 - v2)
    }
    v44 := []int{}
    if len(v44) >= 2 {
        v44 = v44[1:len(v44)]
    }
    v45 := GenS2{F0: 42, F1: 65535, F2: 0}
    if (("a" + ("hello" + "\n\t")) < ("\n\t" + "")) {
        v46 := GenS4{F0: -1, F1: -9223372036854775808, F2: genFn0(18446744073709551615, 2)}
        if (!(!true)) {
            var v47 uint = 0
            v44 = append(v44, 4294967295)
            _ = v47
        } else {
            var v48 bool = ((true || (127 <= 127)) || ((3.14159 - (1e10 - ((0.0 - 0.1) / 1e10))) != -2.25))
            _ = v48
        }
        _ = v46
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v49 string = ""
        _ = v49
    case 42:
        if (!true) {
            v50 := GenS2{F0: 42, F1: 65535, F2: 42}
            switch ("" + "foo bar") {
            case "a":
                if true {
                    var v51 uint8 = 67
                    v51 = v51
                    _ = v51
                }
            case "":
                var v52 bool = (!(244 == -213))
                _ = v52
            default:
                v53 := GenS2{F0: 127, F1: 157, F2: -4}
                _ = v53
            }
            v54 := GenS4{F0: -1, F1: 0, F2: 42}
            _ = v50
            _ = v54
        } else {
            v2 = v2
            for fi55 := 0; fi55 < 3; fi55++ {
                v3 = v3
                _ = fi55
            }
            for fi56 := 0; fi56 < 1; fi56++ {
                v57 := GenS5{F0: 0, F1: 255}
                v1 -= v3
                _ = fi56
                _ = v57
            }
        }
    }
    fmt.Println(v3, v1, v1)
    for fi58 := 0; fi58 < 6; fi58++ {
        var v59 int8 = 100
        _ = fi58
        _ = v59
    }
    v1 += (v3 & 0)
    var v60 uint16 = 0
    for fi61 := 0; fi61 < 12; fi61++ {
        v44 = append(v44, -3)
        v3 = v1
        switch "" {
        case "\n\t":
            v2 += 0.0
        case "a":
            _ = len(v44)
        case "foo bar":
            for fi62 := 0; fi62 < 12; fi62++ {
                var v63 uint8 = 0
                if (1e10 <= ((v2 / (v2 + 3.14159)) * (3.14159 * float64(1.5)))) {
                    v64 := GenS5{F0: 1000, F1: (v60 & 320)}
                    _ = v64
                } else {
                    v65 := GenS4{F0: 2147483647, F1: 399, F2: 1}
                    var v66 uint = 967
                    v67 := GenS3{F0: -1}
                    _ = v65
                    _ = v66
                    _ = v67
                }
                fmt.Println(v3)
                _ = fi62
                _ = v63
            }
        default:
            v68 := GenS6{F0: 860, F1: 2}
            _ = v68
        }
        _ = fi61
    }
    fmt.Println(v1, v3)
    v69 := GenS3{F0: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v44
    _ = v45
    _ = v60
    _ = v69
}
