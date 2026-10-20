package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int64; F1 int8; F2 uint32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 uint; F2 bool }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 uint64 }

func main() {
    var v1 uint8 = 208
    var v2 uint64 = 2
    var v3 bool = true
    switch "foo bar" {
    case "\n\t":
        v4 := []int{54, 2147483647, 2147483647, -9223372036854775808}
        _ = v4
    case "":
        if true {
            v5 := GenS1{F0: 409, F1: -863, F2: 479}
            _ = v5
        }
    default:
        if ((v3 && ((v1 + v1) >= 0)) && (1 == 42)) {
            switch int(v2) {
            case 921:
                var v6 uint = 4294967295
                _ = v6
            case 573:
                v7 := GenS2{F0: 9223372036854775807}
                _ = v7
            case 1:
                if v3 {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        var v9 string = "\n\t"
                        _ = fi8
                        _ = v9
                    }
                    switch "hello" {
                    case "a":
                        switch ("hello" + "hello") {
                        case "\n\t":
                            var v10 float64 = float64((3.14159 - 1e10))
                            _ = v10
                        }
                    case "\n\t":
                        var v11 uint8 = 0
                        _ = v11
                    case "foo bar":
                        var v12 uint8 = v1
                        _ = v12
                    default:
                        if (-130 < 4294967295) {
                            v13 := GenS3{F0: 0, F1: 196}
                            if ((((-2.25 - 1.5) - 1.5) * float32(float64(3.14159))) > 0.0) {
                                for fi14 := 0; fi14 < 11; fi14++ {
                                    v15 := GenS4{F0: 250}
                                    fmt.Println(v2, v2, v1)
                                    v16 := []int{(fi14 % ((128) | 1)), -634, (fi14 % (((fi14 + (^(fi14)))) | 1)), (fi14 | (fi14 * fi14))}
                                    _ = fi14
                                    _ = v15
                                    _ = v16
                                }
                                v1 -= (v1 >> ((65535) & 63))
                            } else {
                                var v17 uint = 42
                                _ = v17
                            }
                            for fi18 := 0; fi18 < 2; fi18++ {
                                v19 := GenS3{F0: (v2 - uint64(v1)), F1: 2}
                                _ = fi18
                                _ = v19
                            }
                            _ = v13
                        }
                    }
                    v1 = (v1 - (v1 - v1))
                } else {
                    var v20 string = (("foo bar" + ("hello" + ("a" + ("" + (("" + "foo bar") + ("\n\t" + "\n\t")))))) + (("\n\t" + "\n\t") + ""))
                    var v21 uint16 = 255
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v23 := []int{42}
                        v24 := GenS5{F0: uint8(v21), F1: -133}
                        fmt.Println(fi22)
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
            }
        } else {
            switch ((("" + "foo bar") + "\n\t") + ("hello" + "a")) {
            case "hello":
                switch 2147483647 {
                case 1:
                    switch ("foo bar" + "hello") {
                    case "a":
                        v2 = 42
                    case "":
                        v25 := GenS5{F0: 62, F1: 9223372036854775807}
                        _ = v25
                    case "foo bar":
                        for fi26 := 0; fi26 < 19; fi26++ {
                            var v27 int64 = -301
                            for fi28 := 0; fi28 < 16; fi28++ {
                                var v29 uint32 = 65535
                                _ = fi28
                                _ = v29
                            }
                            _ = fi26
                            _ = v27
                        }
                    default:
                        v1 -= v1
                    }
                case 2147483647:
                    var v30 uint32 = 42
                    _ = v30
                case 590:
                    v3 = false
                default:
                    for fi31 := 0; fi31 < 4; fi31++ {
                        if v3 {
                            v32 := GenS4{F0: 921}
                            v2 = v2
                            _ = v32
                        }
                        _ = fi31
                    }
                }
            case "":
                if (42 >= 0) {
                    if v3 {
                        v33 := []int{}
                        if len(v33) > 0 {
                            v33[0] = -1
                        }
                        _ = v33
                    } else {
                        v3 = (((0.1 + -2.25) / -2.25) <= 0.1)
                        if (0 < -128) {
                            if v3 {
                                for fi34 := 0; fi34 < 17; fi34++ {
                                    for fi35 := 0; fi35 < 3; fi35++ {
                                        switch (("foo bar" + "hello") + "hello") {
                                        case "\n\t":
                                            var v36 float32 = 0.0
                                            _ = v36
                                        case "":
                                            v37 := []int{}
                                            _ = v37
                                        case "foo bar":
                                            switch ("a" + "") {
                                            case "hello":
                                                v1 = v1
                                            case "a":
                                                v38 := []int{}
                                                _ = v38
                                            case "foo bar":
                                                if ((!(1 != 1)) && false) {
                                                    var v39 uint = 1
                                                    var v40 uint64 = (v2 | v2)
                                                    switch "a" {
                                                    case "\n\t":
                                                        if true {
                                                            v1 = (v1 << ((v39) & 63))
                                                            fmt.Println(fi34, v40)
                                                        }
                                                    case "a":
                                                        v41 := GenS3{F0: (v2 << ((v39) & 63)), F1: 2}
                                                        _ = v41
                                                    case "hello":
                                                        v3 = v3
                                                    }
                                                    _ = v39
                                                    _ = v40
                                                }
                                            default:
                                                v42 := []int{fi35}
                                                switch (fi35 >> ((65535) & 63)) {
                                                case -430:
                                                    switch "hello" {
                                                    case "\n\t":
                                                        if (!(!(!v3))) {
                                                            v42 = append(v42, fi34)
                                                        }
                                                    case "a":
                                                        v43 := GenS6{F0: 1, F1: 0, F2: 2147483648}
                                                        _ = v43
                                                    case "":
                                                        var v44 int32 = -1
                                                        _ = v44
                                                    default:
                                                        if len(v42) > 0 {
                                                            v42[0] = (fi34 % ((fi34) | 1))
                                                        }
                                                    }
                                                }
                                                v45 := GenS1{F0: 0, F1: 317, F2: 0}
                                                _ = v42
                                                _ = v45
                                            }
                                        }
                                        if ((fi35 & (fi34 - fi35)) <= fi34) {
                                            v3 = ((3.14159 < 0.0) && v3)
                                            v2 += v2
                                            v3 = v3
                                            if (v2 <= v2) {
                                                if (true || false) {
                                                    v46 := GenS3{F0: (^(v2)), F1: 4294967295}
                                                    _ = v46
                                                }
                                            }
                                        }
                                        var v47 uint16 = 32767
                                        _ = fi35
                                        _ = v47
                                    }
                                    var v48 int = (fi34 | fi34)
                                    var v49 uint32 = 2147483648
                                    _ = fi34
                                    _ = v48
                                    _ = v49
                                }
                                var v50 float64 = float64(0.1)
                                var v51 int64 = -339
                                _ = v50
                                _ = v51
                            } else {
                                v52 := GenS3{F0: v2, F1: 902}
                                for fi53 := 0; fi53 < 20; fi53++ {
                                    v54 := GenS2{F0: 9223372036854775807}
                                    _ = fi53
                                    _ = v54
                                }
                                var v55 float64 = 1.5
                                _ = v52
                                _ = v55
                            }
                        }
                        var v56 bool = (2147483647 > 1)
                        _ = v56
                    }
                }
            case "\n\t":
                if v3 {
                    v1 = 128
                    var v57 bool = true
                    _ = v57
                }
            }
            v58 := []int{}
            if len(v58) > 0 {
                v58[0] = 9223372036854775807
            }
            _ = v58
        }
    }
    v59 := []int{}
    switch "hello" {
    case "\n\t":
        if (18446744073709551615 != 664) {
            var v60 uint64 = v2
            for fi61 := 0; fi61 < 3; fi61++ {
                v60 = v60
                var v62 int8 = -64
                v62 = int8(v2)
                _ = fi61
                _ = v62
            }
            _ = v60
        } else {
            if v3 {
                var v63 bool = (32767 >= uint16(v1))
                _ = v63
            }
            fmt.Println(v1)
            v64 := GenS7{F0: (0 != 713)}
            _ = v64
        }
    }
    _ = len(v59)
    var v65 int8 = 42
    var v66 float32 = (3.14159 + ((0.0 * -2.25) - (0.1 - (1e10 + (((0.1 * float32(((3.14159 / (-2.25 + (3.14159 - 1.5))) / -2.25))) - 3.14159) - -2.25)))))
    fmt.Println(v66, v2)
    if (!true) {
        v67 := GenS2{F0: 0}
        _ = v67
    } else {
        var v68 bool = (65535 > 4294967295)
        _ = v68
    }
    var v69 string = "\n\t"
    if (!(!v3)) {
        fmt.Println(v2, v2, v65)
        if v3 {
            if v3 {
                if false {
                    v70 := GenS6{F0: -62, F1: v65, F2: 2147483647}
                    v71 := GenS7{F0: ((0.0 >= 1e10) || true)}
                    v72 := GenS8{F0: v65, F1: 216, F2: v3}
                    _ = v70
                    _ = v71
                    _ = v72
                }
                fmt.Println(v66)
                v73 := GenS9{F0: true}
                _ = v73
            } else {
                fmt.Println(v66, v3, v65)
            }
            fmt.Println(v65, v65)
        } else {
            v69 = ""
            v74 := GenS4{F0: 2147483647}
            if len(v59) >= 2 {
                v59 = v59[1:len(v59)]
            }
            _ = v74
        }
    }
    v75 := GenS10{F0: v2}
    var v76 uint8 = v1
    v66 -= (((3.14159 + (float32(1e10) - v66)) + ((v66 / 1.0) * (v66 * -2.25))) - v66)
    switch v69 {
    case "foo bar":
        var v77 float32 = v66
        _ = v77
    default:
        var v78 float32 = v66
        _ = v78
    }
    var v79 bool = (1 >= 494)
    v80 := GenS4{F0: -39}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v59:", len(v59), v59)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v69:", v69)
    fmt.Println("v76:", v76)
    fmt.Println("v79:", v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v59
    _ = v65
    _ = v66
    _ = v69
    _ = v75
    _ = v76
    _ = v79
    _ = v80
}
