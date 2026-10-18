package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 uint8; F1 bool }
type GenS7 struct { F0 int; F1 int16 }
type GenS8 struct { F0 uint8; F1 uint8; F2 uint64 }
func genFn0() string {
    var v4 uint8 = 0
    if true {
        v4 -= v4
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 5; fi6++ {
                if ("a" == ("a" + "\n\t")) {
                    v4 -= 2
                } else {
                    if true {
                        v7 := []int{}
                        var v8 bool = ("foo bar" > "\n\t")
                        var v9 int16 = 653
                        _ = v7
                        _ = v8
                        _ = v9
                    }
                    switch -2147483648 {
                    case -75:
                        v4 -= (v4 ^ 127)
                    case 0:
                        v10 := []int{}
                        _ = v10
                    }
                    var v11 int = fi6
                    var v12 int64 = -192
                    var v13 uint16 = uint16(v4)
                    _ = v11
                    _ = v12
                    _ = v13
                }
                switch "a" {
                case "a":
                    var v14 uint32 = 2
                    v14 += (v14 & v14)
                    v15 := []int{0, (fi6 >> ((2) & 63)), 1, 32768}
                    _ = v14
                    _ = v15
                case "hello":
                    if true {
                        if false {
                            v16 := []int{(-(fi6)), fi5, fi5}
                            _ = v16
                        } else {
                            v4 += (v4 % ((v4) | 1))
                            fmt.Println(fi5, v4)
                            if ((-32768 != -32768) && false) {
                                var v17 uint32 = 2147483647
                                _ = v17
                            }
                        }
                        switch (-(fi5)) {
                        case -2147483648:
                            var v18 string = "hello"
                            _ = v18
                        case -9223372036854775808:
                            var v19 int16 = -335
                            _ = v19
                        case 2147483647:
                            if true {
                                switch fi5 {
                                case -9223372036854775808:
                                    var v20 int32 = 2147483647
                                    _ = v20
                                case -2147483648:
                                    v4 -= (^(v4))
                                }
                                switch "a" {
                                case "":
                                    v4 -= v4
                                    var v21 uint32 = 1
                                    v22 := []int{}
                                    _ = v21
                                    _ = v22
                                case "\n\t":
                                    var v23 bool = ((-9223372036854775808 != int64(v4)) && true)
                                    _ = v23
                                default:
                                    if true {
                                        if (1 < 42) {
                                            v24 := GenS1{F0: 65535, F1: -1}
                                            _ = v24
                                        }
                                        fmt.Println(fi5, v4)
                                        v25 := []int{}
                                        _ = v25
                                    } else {
                                        v26 := GenS2{F0: -2147483648}
                                        v27 := []int{(-(fi6))}
                                        if true {
                                            if len(v27) > 0 {
                                                v27[0] = fi6
                                            }
                                        }
                                        _ = v26
                                        _ = v27
                                    }
                                }
                                switch fi5 {
                                case 2147483647:
                                    var v28 int8 = -16
                                    _ = v28
                                case -2147483648:
                                    v29 := GenS3{F0: -9223372036854775808, F1: uint64(v4)}
                                    _ = v29
                                }
                            } else {
                                v4 -= v4
                                var v30 float64 = ((-2.25 / -2.25) * -2.25)
                                v31 := []int{-717, fi5}
                                _ = v30
                                _ = v31
                            }
                        default:
                            v4 = 128
                        }
                        v32 := GenS4{F0: false}
                        _ = v32
                    }
                }
                var v33 int16 = 1000
                _ = fi6
                _ = v33
            }
            _ = fi5
        }
        for fi34 := 0; fi34 < 20; fi34++ {
            v4 = v4
            var v35 int = 359
            _ = fi34
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        for fi37 := 0; fi37 < 14; fi37++ {
            v4 = 128
            if true {
                v38 := GenS2{F0: -2147483648}
                _ = v38
            } else {
                v4 = (v4 << ((18446744073709551615) & 63))
            }
            _ = fi37
        }
        v4 -= (v4 % ((v4) | 1))
        v39 := GenS3{F0: -250, F1: 4294967295}
        _ = fi36
        _ = v39
    }
    if false {
        v40 := GenS5{F0: 42, F1: "\n\t"}
        switch 32767 {
        case 128:
            if true {
                v4 = 0
            } else {
                if (980 < 255) {
                    v41 := []int{32767}
                    _ = v41
                } else {
                    var v42 int64 = 0
                    fmt.Println(v4, v42)
                    _ = v42
                }
                if false {
                    v43 := GenS6{F0: v4, F1: true}
                    v44 := []int{740}
                    _ = v43
                    _ = v44
                }
                fmt.Println(v4)
            }
        case 1:
            fmt.Println(v4)
        default:
            if (false && (32767 <= 420)) {
                var v45 int32 = -896
                v46 := []int{0, 2147483647}
                v45 += v45
                _ = v45
                _ = v46
            }
        }
        _ = v40
    }
    _ = v4
    return "hello"
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 2147483648
    var v3 int16 = 32767
    for fi47 := 0; fi47 < 7; fi47++ {
        v48 := GenS7{F0: (fi47 - 42), F1: (v3 | 941)}
        v49 := []int{}
        _ = fi47
        _ = v48
        _ = v49
    }
    var v50 uint64 = uint64(v2)
    if false {
        v2 = 555
        var v51 int64 = 1
        v52 := GenS2{F0: 1}
        _ = v51
        _ = v52
    }
    if (20 > 127) {
        v1 = (v1 / (((v1 << ((1) & 63))) | 1))
        if false {
            if (128 <= 2) {
                var v53 uint32 = v2
                var v54 int8 = -1
                if (-316 >= 32768) {
                    v55 := GenS8{F0: 0, F1: 255, F2: uint64(v2)}
                    _ = v55
                }
                _ = v53
                _ = v54
            }
            v56 := []int{-691, 128, 9223372036854775807, 898}
            v1 -= v1
            _ = v56
        }
    } else {
        var v57 int64 = 9223372036854775807
        var v58 string = "foo bar"
        _ = v57
        _ = v58
    }
    var v59 uint = 2
    fmt.Println(v3, v59)
    if (-316 != 387) {
        v60 := []int{}
        _ = v60
    } else {
        switch "hello" {
        case "a":
            v61 := []int{674, -9223372036854775808, 1, 0}
            _ = v61
        }
        v2 -= 648
    }
    for fi62 := 0; fi62 < 15; fi62++ {
        fmt.Println(v1)
        fmt.Println(v1, v2)
        _ = fi62
    }
    if (true || false) {
        fmt.Println(v2)
        if true {
            if (-2.25 == 3.14159) {
                v50 = v50
                for fi63 := 0; fi63 < 15; fi63++ {
                    v64 := []int{(fi63 * fi63), -9223372036854775808, fi63}
                    var v65 bool = (255 < 2)
                    v50 = v50
                    _ = fi63
                    _ = v64
                    _ = v65
                }
            }
        }
        fmt.Println(v3, v59)
    }
    for fi66 := 0; fi66 < 12; fi66++ {
        fmt.Println(v3, v50, v50)
        _ = fi66
    }
    v67 := GenS7{F0: 0, F1: 596}
    v59 = v59
    v59 = (v59 | v59)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v50:", v50)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v50
    _ = v59
    _ = v67
}
