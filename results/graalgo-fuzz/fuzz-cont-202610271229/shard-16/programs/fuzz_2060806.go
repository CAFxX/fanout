package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint; F1 string }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 string; F2 uint8 }
type GenS7 struct { F0 uint16; F1 uint16; F2 uint32 }
type GenS8 struct { F0 bool; F1 uint }
func genFn0(p0 int8) int64 {
    if false {
        for fi5 := 0; fi5 < 13; fi5++ {
            v6 := GenS1{F0: 0, F1: ("hello" + "a")}
            _ = fi5
            _ = v6
        }
    }
    switch 1 {
    case -1:
        switch 32768 {
        case 334:
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 string = "\n\t"
                var v9 uint = 1
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            v10 := GenS2{F0: "\n\t", F1: 0, F2: -1}
            _ = v10
        }
    case 42:
        v11 := GenS3{F0: 255, F1: 71}
        _ = v11
    case 127:
        v12 := []int{}
        _ = v12
    }
    if true {
        if false {
            var v13 int32 = 2147483647
            _ = v13
        } else {
            v14 := []int{-178}
            var v15 int32 = 38
            _ = v14
            _ = v15
        }
        var v16 bool = (49 >= 128)
        _ = v16
    } else {
        if false {
            v17 := GenS3{F0: 255, F1: 2}
            _ = v17
        }
        v18 := GenS3{F0: 652, F1: 2147483647}
        v19 := []int{}
        _ = v18
        _ = v19
    }
    if false {
        v20 := []int{-548}
        switch "a" {
        case "foo bar":
            fmt.Println(p0)
        case "a":
            var v21 uint = 1
            _ = v21
        case "":
            p0 -= 31
        default:
            v22 := GenS4{F0: 1, F1: (("\n\t" + (("a" + "a") + "hello")) + "")}
            _ = v22
        }
        _ = v20
    }
    _ = p0
    return 843
}

func main() {
    var v1 string = ("\n\t" + "a")
    var v2 uint32 = 65535
    v3 := make(map[uint32]int)
    var v4 uint32 = (v2 | 926)
    if (int32(v4) == 2147483647) {
        switch v1 {
        case "hello":
            v1 = v1
        case "\n\t":
            var v23 bool = true
            _ = v23
        }
        if (!false) {
            v24 := GenS3{F0: 559, F1: (^(v4))}
            _ = v24
        } else {
            if false {
                for fi25 := 0; fi25 < 13; fi25++ {
                    var v26 uint64 = 640
                    _ = fi25
                    _ = v26
                }
            } else {
                switch v1 {
                case "hello":
                    var v27 uint32 = 4294967295
                    _ = v27
                case "foo bar":
                    fmt.Println(v1, v2, v2)
                default:
                    if ((-928 < -32768) && false) {
                        v4 = (v4 / ((v2) | 1))
                    }
                }
            }
        }
    } else {
        var v28 uint = 2
        var v29 int32 = -1
        if ("a" < "foo bar") {
            v4 = v2
        } else {
            v28 = v28
            for fi30 := 0; fi30 < 13; fi30++ {
                for fi31 := 0; fi31 < 19; fi31++ {
                    if true {
                        v32 := v3[(v4 >> ((uint(v2)) & 63))]
                        _ = v32
                    }
                    delete(v3, (v2 >> ((4294967295) & 63)))
                    if false {
                        v33 := GenS5{F0: 127}
                        _ = v33
                    } else {
                        var v34 int64 = 42
                        var v35 bool = (v28 >= v28)
                        _ = v34
                        _ = v35
                    }
                    _ = fi31
                }
                for fi36 := 0; fi36 < 4; fi36++ {
                    v37 := GenS6{F0: 128, F1: v1, F2: 2}
                    _ = fi36
                    _ = v37
                }
                delete(v3, (^(v4)))
                _ = fi30
            }
        }
        _ = v28
        _ = v29
    }
    _ = len(v3)
    v1 = (v1 + v1)
    switch "hello" {
    case "hello":
        if false {
            if (!true) {
                var v38 uint = 0
                _ = v38
            }
            if (((("hello" <= v1) && false) || false) && (1 == 42)) {
                var v39 bool = false
                var v40 int = -2147483648
                _ = v39
                _ = v40
            } else {
                if true {
                    switch "\n\t" {
                    case "hello":
                        switch "\n\t" {
                        case "\n\t":
                            v3[2147483647] = 128
                        case "hello":
                            if (!false) {
                                if (1.5 < 0.0) {
                                    var v41 uint8 = 144
                                    v42 := []int{-1, 4294967295}
                                    v41 = 42
                                    _ = v41
                                    _ = v42
                                } else {
                                    v4 += (v4 << ((65535) & 63))
                                    if true {
                                        fmt.Println(v4, v2, v4)
                                        fmt.Println(v2, v1)
                                        var v43 uint16 = 42
                                        _ = v43
                                    } else {
                                        for fi44 := 0; fi44 < 3; fi44++ {
                                            delete(v3, (v4 ^ v2))
                                            var v45 float32 = (1.5 - 0.0)
                                            _ = fi44
                                            _ = v45
                                        }
                                    }
                                    var v46 uint16 = uint16(v2)
                                    _ = v46
                                }
                            } else {
                                fmt.Println(v4)
                                fmt.Println(v1, v2)
                            }
                        case "foo bar":
                            v4 = (v4 % (((^(v4))) | 1))
                        default:
                            for fi47 := 0; fi47 < 8; fi47++ {
                                if false {
                                    if (1 > 2) {
                                        v2 = v4
                                        v48 := []int{4294967295, fi47, fi47}
                                        _ = v48
                                    }
                                } else {
                                    var v49 uint8 = 12
                                    var v50 bool = (!false)
                                    _ = v49
                                    _ = v50
                                }
                                _ = fi47
                            }
                        }
                    case "foo bar":
                        delete(v3, (v4 - v2))
                    case "\n\t":
                        v51 := GenS5{F0: 104}
                        _ = v51
                    default:
                        switch "" {
                        case "hello":
                            v52 := GenS2{F0: v1, F1: 2, F2: 32767}
                            _ = v52
                        case "foo bar":
                            var v53 int32 = -1
                            _ = v53
                        case "\n\t":
                            var v54 float64 = 1.5
                            _ = v54
                        }
                    }
                    var v55 int32 = -2147483648
                    fmt.Println(v4, v4, v2)
                    _ = v55
                }
            }
            var v56 float64 = (float64(0.0) * 1e10)
            _ = v56
        }
    case "":
        v2 += v2
    case "foo bar":
        v57 := GenS4{F0: 101, F1: "foo bar"}
        _ = v57
    default:
        var v58 int16 = -32768
        _ = v58
    }
    switch -9223372036854775808 {
    case 32767:
        v4 = (v4 % ((v4) | 1))
    case -14:
        v59 := []int{127, 401}
        _ = v59
    default:
        v60 := GenS5{F0: uint8(v4)}
        _ = v60
    }
    var v61 int8 = 104
    if (uint16(v4) == 32767) {
        for fi62 := 0; fi62 < 19; fi62++ {
            var v63 bool = false
            _ = fi62
            _ = v63
        }
    }
    v64 := GenS4{F0: 196, F1: v1}
    var v65 uint64 = 18446744073709551615
    var v66 uint16 = 906
    for fi67 := 0; fi67 < 1; fi67++ {
        v61 = (v61 & 100)
        _ = fi67
    }
    if true {
        v68 := GenS7{F0: 1, F1: v66, F2: uint32(v65)}
        if false {
            v66 -= (v66 - 65535)
            if false {
                v1 = v1
                v69 := v3[65535]
                for fi70 := 0; fi70 < 6; fi70++ {
                    v71 := GenS1{F0: (v66 << ((18446744073709551615) & 63)), F1: v1}
                    var v72 int16 = 42
                    var v73 float64 = 1e10
                    _ = fi70
                    _ = v71
                    _ = v72
                    _ = v73
                }
                _ = v69
            }
        } else {
            var v74 int32 = -900
            _ = v74
        }
        _ = v68
    }
    var v75 float64 = 0.0
    v76 := GenS4{F0: 65535, F1: v1}
    if true {
        if true {
            fmt.Println(v65, v65, v66)
            for fi77 := 0; fi77 < 19; fi77++ {
                v78 := []int{fi77, -158, fi77}
                fmt.Println(v2, v4, v4)
                for fi79 := 0; fi79 < 19; fi79++ {
                    if len(v78) >= 2 {
                        v78 = v78[1:len(v78)]
                    }
                    v75 += v75
                    _ = fi79
                }
                _ = fi77
                _ = v78
            }
        }
    }
    v1 = v1
    v4 -= v4
    switch "hello" {
    case "foo bar":
        v80 := GenS8{F0: (32767 != -32768), F1: 18446744073709551615}
        _ = v80
    default:
        fmt.Println(v2, v66)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v61:", v61)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v61
    _ = v64
    _ = v65
    _ = v66
    _ = v75
    _ = v76
}
