package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint }
type GenS8 struct { F0 int64; F1 int32 }
type GenS9 struct { F0 string; F1 int16; F2 int16 }
type GenS10 struct { F0 uint64; F1 int }
type GenS11 struct { F0 int }
func genFn0(p0 bool) int {
    var v7 float64 = (1e10 / (0.1 * 3.14159))
    for fi8 := 0; fi8 < 9; fi8++ {
        if true {
            for fi9 := 0; fi9 < 14; fi9++ {
                var v10 uint64 = 675
                _ = fi9
                _ = v10
            }
            v11 := GenS1{F0: ("a" + ("hello" + ("a" + "a")))}
            _ = v11
        }
        var v12 string = "hello"
        _ = fi8
        _ = v12
    }
    if (!p0) {
        for fi13 := 0; fi13 < 9; fi13++ {
            for fi14 := 0; fi14 < 20; fi14++ {
                v15 := []int{fi14, 128}
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
        fmt.Println(p0)
        var v16 float32 = 3.14159
        _ = v16
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        if p0 {
            if p0 {
                switch (fi17 * fi17) {
                case 703:
                    v7 -= v7
                case 128:
                    var v18 uint32 = 796
                    _ = v18
                case -748:
                    v19 := GenS2{F0: 42, F1: -21, F2: 0}
                    _ = v19
                }
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                p0 = (!(42 == 2147483647))
                v21 := GenS3{F0: 386, F1: 858}
                _ = fi20
                _ = v21
            }
            var v22 uint8 = 42
            _ = v22
        }
        _ = fi17
    }
    _ = p0
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 int8, p1 int16) bool {
    v23 := GenS4{F0: (-(p0)), F1: 65535, F2: 645}
    var v24 uint64 = 534
    v25 := []int{int(v24), -107, -856}
    for fi26 := 0; fi26 < 11; fi26++ {
        v27 := GenS5{F0: p1}
        fmt.Println(p0, fi26)
        fmt.Println(p0)
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    _ = v25
    return false
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = false
    v3 := make(map[uint32]int)
    var v4 int = 1
    var v5 int64 = 42
    var v6 bool = v2
    v2 = v2
    v2 = true
    fmt.Println(v5, v2)
    var v28 uint = 42
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := GenS6{F0: 0, F1: 676, F2: 0}
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        v28 -= (v28 & v28)
        _ = fi31
    }
    var v32 uint16 = 1
    fmt.Println(v32)
    if v2 {
        if (!(0.1 == -2.25)) {
            if v2 {
                v4 -= (v4 >> ((255) & 63))
                fmt.Println(v28, v28, v32)
                if v2 {
                    fmt.Println(v28, v32)
                } else {
                    var v33 int16 = int16(v5)
                    switch 1 {
                    case 32767:
                        v34 := GenS7{F0: 1, F1: (633 < 0), F2: 55}
                        _ = v34
                    case -1:
                        v35 := GenS7{F0: 0, F1: false, F2: v28}
                        _ = v35
                    case 9223372036854775807:
                        switch "" {
                        case "\n\t":
                            switch (v4 * v4) {
                            case 4294967295:
                                fmt.Println(v28, v32)
                            }
                        case "hello":
                            if v2 {
                                v36 := []int{v4, int(v28), v4}
                                v37 := GenS7{F0: uint8(v32), F1: (!false), F2: (v28 % ((2) | 1))}
                                _ = v36
                                _ = v37
                            } else {
                                if (!((0.1 + (0.0 * (0.1 - float64((1e10 * (0.0 * 0.0)))))) < -2.25)) {
                                    v38 := GenS8{F0: v5, F1: 0}
                                    var v39 uint8 = 66
                                    fmt.Println(v39, v39, v32)
                                    _ = v38
                                    _ = v39
                                } else {
                                    if (42 >= 3) {
                                        switch "a" {
                                        case "foo bar":
                                            var v40 int = (v4 + genFn0(v2))
                                            _ = v40
                                        }
                                        if (!v2) {
                                            var v41 int = v4
                                            _ = v41
                                        } else {
                                            fmt.Println(v33, v28)
                                            if v2 {
                                                v42 := []int{v4, v4}
                                                v6 = v2
                                                _ = v42
                                            } else {
                                                var v43 uint64 = 1
                                                var v44 uint8 = 127
                                                _ = len(v3)
                                                _ = v43
                                                _ = v44
                                            }
                                        }
                                    } else {
                                        var v45 int = 127
                                        switch "hello" {
                                        case "hello":
                                            if v6 {
                                                if ((1e10 + float32(1.5)) == (3.14159 / float32(3.14159))) {
                                                    v6 = true
                                                    var v46 int16 = v33
                                                    switch "a" {
                                                    case "hello":
                                                        v3[uint32(v32)] = (v45 << ((4294967295) & 63))
                                                    case "a":
                                                        v6 = (0 >= 18446744073709551615)
                                                    case "foo bar":
                                                        for fi47 := 0; fi47 < 11; fi47++ {
                                                            var v48 int16 = v46
                                                            for fi49 := 0; fi49 < 13; fi49++ {
                                                                if (!(("" > (("foo bar" + "a") + "\n\t")) && (v2 && ((v28 >> (((v28 % ((v28) | 1))) & 63)) > uint(v32))))) {
                                                                    var v50 int = (-(v45))
                                                                    fmt.Println(v5, v48, v28)
                                                                    v32 += (v32 ^ (v32 >> (((v28 / (((v28 - v28)) | 1))) & 63)))
                                                                    _ = v50
                                                                } else {
                                                                    switch v4 {
                                                                    case 0:
                                                                        var v51 string = ("hello" + "hello")
                                                                        _ = v51
                                                                    case 2147483647:
                                                                        var v52 int32 = -2147483648
                                                                        _ = v52
                                                                    }
                                                                }
                                                                v53 := []int{(^(fi47)), -1}
                                                                v2 = v6
                                                                switch (fi49 << ((v28) & 63)) {
                                                                case 9223372036854775807:
                                                                    var v54 bool = ((v32 >> ((v28) & 63)) >= (v32 + (v32 / ((v32) | 1))))
                                                                    _ = v54
                                                                case 2147483647:
                                                                    var v55 uint64 = 0
                                                                    _ = v55
                                                                default:
                                                                    if v2 {
                                                                        var v56 string = ((("" + "hello") + ("foo bar" + ("a" + "hello"))) + "foo bar")
                                                                        v57 := GenS9{F0: (v56 + v56), F1: v46, F2: (-(v33))}
                                                                        v58 := GenS3{F0: 2, F1: -1}
                                                                        v4 = v45
                                                                        _ = v56
                                                                        _ = v57
                                                                        _ = v58
                                                                    } else {
                                                                        v6 = v6
                                                                        var v59 int64 = v5
                                                                        _ = v59
                                                                    }
                                                                }
                                                                _ = fi49
                                                                _ = v53
                                                            }
                                                            _ = fi47
                                                            _ = v48
                                                        }
                                                    }
                                                    _ = v46
                                                } else {
                                                    v3[980] = v4
                                                    switch "a" {
                                                    case "a":
                                                        switch "\n\t" {
                                                        case "":
                                                            v2 = false
                                                        case "hello":
                                                            if ((9223372036854775807 <= -1) && v2) {
                                                                v5 -= (v5 >> (((v28 & 129)) & 63))
                                                            }
                                                        case "a":
                                                            v45 += (v4 & genFn0((!v6)))
                                                        }
                                                    case "":
                                                        v32 -= v32
                                                    case "foo bar":
                                                        v28 += (v28 + 4294967295)
                                                    }
                                                }
                                            }
                                        default:
                                            fmt.Println(v45)
                                        }
                                        _ = v45
                                    }
                                    for fi60 := 0; fi60 < 5; fi60++ {
                                        if v2 {
                                            v61 := GenS10{F0: 4294967295, F1: fi60}
                                            _ = v61
                                        } else {
                                            switch (-(fi60)) {
                                            case -1:
                                                v5 += -624
                                            case -9223372036854775808:
                                                v3[404] = (fi60 << ((v28) & 63))
                                            case 188:
                                                v33 = (v33 | 1000)
                                            default:
                                                if (uint8(v28) == 0) {
                                                    v4 -= fi60
                                                }
                                            }
                                            var v62 uint = (v28 - 18446744073709551615)
                                            v32 += v32
                                            _ = v62
                                        }
                                        _ = fi60
                                    }
                                }
                                var v63 uint = 0
                                fmt.Println(v2, v2, v5)
                                _ = v63
                            }
                        case "a":
                            if ((v5 + (v5 * 1)) < (v5 ^ 1)) {
                                var v64 int = (v4 - v4)
                                v28 = uint(v32)
                                _ = v64
                            } else {
                                var v65 int32 = 2147483647
                                switch (v4 - (v4 << ((762) & 63))) {
                                case 32768:
                                    var v66 uint64 = 42
                                    _ = v66
                                case 42:
                                    v67 := []int{(v4 * (v4 * v4))}
                                    _ = v67
                                default:
                                    fmt.Println(v2, v5, v4)
                                }
                                _ = v65
                            }
                        }
                    default:
                        switch v4 {
                        case 0:
                            var v68 int8 = -124
                            _ = v68
                        case -109:
                            v69 := GenS11{F0: -37}
                            _ = v69
                        case -717:
                            v4 -= (v4 / ((v4) | 1))
                        default:
                            switch v4 {
                            case 873:
                                v28 += v28
                            case 32767:
                                _ = len(v3)
                            }
                        }
                    }
                    v33 = (v33 ^ int16(v28))
                    _ = v33
                }
            } else {
                v2 = (v4 < -2147483648)
            }
            var v70 uint32 = 65535
            v3[458] = genFn0(v2)
            _ = v70
        } else {
            v2 = (42 >= v5)
            v28 = v28
            var v71 uint8 = 0
            _ = v71
        }
        var v72 float64 = 3.14159
        _ = v72
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
