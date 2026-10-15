package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint64; F1 uint32; F2 bool }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int8; F1 int16; F2 uint }
type GenS8 struct { F0 int; F1 bool; F2 int64 }

func main() {
    var v1 float32 = 1e10
    var v2 uint16 = 42
    var v3 int32 = 42
    var v4 uint32 = 65535
    var v5 int64 = -833
    v2 = v2
    v6 := GenS1{F0: 146, F1: true}
    var v7 uint = 2
    if false {
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: ("hello" + "\n\t")}
            _ = fi8
            _ = v9
        }
        v2 -= (v2 ^ (v2 + (v2 >> ((v7) & 63))))
        switch "a" {
        case "":
            var v10 int16 = 0
            _ = v10
        case "foo bar":
            v11 := GenS2{F0: "foo bar"}
            _ = v11
        }
    }
    if (1 <= 32767) {
        if false {
            switch 1 {
            case -1:
                v12 := GenS2{F0: ("hello" + "\n\t")}
                _ = v12
            default:
                v7 = uint(v4)
            }
            if false {
                v1 = (((1e10 * v1) * ((v1 / (v1 - -2.25)) * v1)) * 3.14159)
            }
        } else {
            if ((v3 | (v3 ^ 1)) != v3) {
                if (2 >= 101) {
                    fmt.Println(v5)
                    if (!true) {
                        var v13 int64 = v5
                        v5 = v5
                        _ = v13
                    } else {
                        var v14 int64 = 9223372036854775807
                        _ = v14
                    }
                    var v15 int16 = int16(v7)
                    _ = v15
                } else {
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 uint64 = 42
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        if (v1 < (1.5 - 1.5)) {
                            v3 = v3
                            var v19 int = (fi18 >> (((^(v7))) & 63))
                            _ = v19
                        } else {
                            if true {
                                v3 = 0
                                switch fi18 {
                                case 964:
                                    var v20 uint64 = 2
                                    _ = v20
                                case -889:
                                    var v21 uint64 = 0
                                    _ = v21
                                default:
                                    switch 128 {
                                    case 32767:
                                        v22 := GenS2{F0: ""}
                                        _ = v22
                                    }
                                }
                                if true {
                                    for fi23 := 0; fi23 < 5; fi23++ {
                                        var v24 bool = ((false || (false && false)) || (9223372036854775807 <= 0))
                                        if ("foo bar" != "\n\t") {
                                            if false {
                                                var v25 uint64 = 42
                                                if v24 {
                                                    v26 := GenS3{F0: 192, F1: (v25 * (v25 | (v25 >> ((18446744073709551615) & 63))))}
                                                    if false {
                                                        v27 := []int{(fi18 % (((fi23 * (-(fi18)))) | 1)), 32768}
                                                        v7 = (v7 % ((uint(v4)) | 1))
                                                        _ = v27
                                                    }
                                                    _ = v26
                                                }
                                                _ = v25
                                            }
                                        } else {
                                            v28 := GenS2{F0: ("" + "hello")}
                                            _ = v28
                                        }
                                        _ = fi23
                                        _ = v24
                                    }
                                    for fi29 := 0; fi29 < 3; fi29++ {
                                        v1 += v1
                                        v2 = v2
                                        v2 += v2
                                        _ = fi29
                                    }
                                }
                            } else {
                                for fi30 := 0; fi30 < 20; fi30++ {
                                    v31 := GenS2{F0: "foo bar"}
                                    var v32 bool = (!(v2 != 102))
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                                if true {
                                    if false {
                                        var v33 int16 = 591
                                        _ = v33
                                    }
                                    switch 368 {
                                    case -9223372036854775808:
                                        var v34 float64 = (((1e10 / 1e10) - 1e10) + -2.25)
                                        _ = v34
                                    default:
                                        v3 -= 1
                                    }
                                }
                                if true {
                                    v35 := []int{(fi18 - 4294967295), 1}
                                    fmt.Println(fi18, v5, v4)
                                    switch 32767 {
                                    case 0:
                                        for fi36 := 0; fi36 < 16; fi36++ {
                                            v7 = 210
                                            _ = fi36
                                        }
                                    default:
                                        for fi37 := 0; fi37 < 7; fi37++ {
                                            switch 4294967295 {
                                            case 497:
                                                v38 := GenS1{F0: 42, F1: (1.5 >= (((0.1 * -2.25) * (1e10 / 1.5)) - 0.0))}
                                                _ = v38
                                            case 83:
                                                v39 := GenS4{F0: 0, F1: 0}
                                                _ = v39
                                            default:
                                                var v40 float32 = v1
                                                _ = v40
                                            }
                                            var v41 uint8 = uint8(v4)
                                            _ = fi37
                                            _ = v41
                                        }
                                    }
                                    _ = v35
                                } else {
                                    var v42 float32 = v1
                                    _ = v42
                                }
                            }
                        }
                        fmt.Println(fi18)
                        _ = fi18
                    }
                }
                v43 := []int{947}
                for fi44 := 0; fi44 < 5; fi44++ {
                    v43 = append(v43, (fi44 + fi44))
                    v1 -= v1
                    _ = fi44
                }
                _ = v43
            } else {
                if true {
                    v7 = 169
                }
                var v45 float64 = (3.14159 / 1.5)
                _ = v45
            }
        }
    }
    v46 := []int{305, 128, 997}
    fmt.Println(v5)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v4)
    case "":
        for fi47 := 0; fi47 < 2; fi47++ {
            for fi48 := 0; fi48 < 15; fi48++ {
                if (1 <= (^(v2))) {
                    v7 -= v7
                }
                v2 = (v2 << (((v7 << ((737) & 63))) & 63))
                fmt.Println(fi48, v2, v5)
                if true {
                    var v49 uint8 = 1
                    if true {
                        switch int(v3) {
                        case 2147483647:
                            v50 := GenS5{F0: uint64(v7), F1: 126, F2: (true && true)}
                            _ = v50
                        default:
                            if (1 == 32767) {
                                v3 = int32(v49)
                                var v51 bool = true
                                _ = v51
                            } else {
                                fmt.Println(fi48, v2)
                                v7 -= v7
                                fmt.Println(fi48)
                            }
                        }
                    }
                    _ = v49
                } else {
                    var v52 uint16 = v2
                    var v53 uint8 = 128
                    v7 -= 482
                    _ = v52
                    _ = v53
                }
                fmt.Println(v4)
                v54 := GenS6{F0: uint64(v7)}
                v55 := GenS7{F0: 42, F1: int16(fi47), F2: v7}
                _ = fi48
                _ = v54
                _ = v55
            }
            _ = fi47
        }
    case "a":
        var v56 bool = (true || false)
        _ = v56
    }
    fmt.Println(v4, v3)
    for fi57 := 0; fi57 < 5; fi57++ {
        fmt.Println(v3, v1, v7)
        _ = fi57
    }
    for fi58 := 0; fi58 < 18; fi58++ {
        v59 := GenS8{F0: fi58, F1: false, F2: (-(v5))}
        _ = fi58
        _ = v59
    }
    switch -438 {
    case 32767:
        v7 -= v7
    case -9223372036854775808:
        var v60 uint = v7
        _ = v60
    case 42:
        if len(v46) > 0 {
            v46[0] = 9223372036854775807
        }
    default:
        v7 = v7
    }
    var v61 int64 = v5
    v7 += (^(v7))
    v46 = append(v46, -9223372036854775808)
    for fi62 := 0; fi62 < 6; fi62++ {
        var v63 uint64 = 962
        v61 = (v61 >> (((v7 << (((v7 << (((v7 - 310)) & 63))) & 63))) & 63))
        _ = fi62
        _ = v63
    }
    v2 -= (v2 & 32767)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
    _ = v61
    _ = v7
}
