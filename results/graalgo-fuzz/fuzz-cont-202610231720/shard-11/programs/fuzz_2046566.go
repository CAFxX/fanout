package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 int32; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 int64; F1 uint }
type GenS7 struct { F0 bool; F1 int32; F2 uint16 }
type GenS8 struct { F0 int64; F1 uint8; F2 int16 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 uint; F1 uint8; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint32 }
type GenS13 struct { F0 int }
type GenS14 struct { F0 int; F1 int16; F2 uint16 }
type GenS15 struct { F0 uint8; F1 uint16; F2 uint }
type GenS16 struct { F0 bool }
type GenS17 struct { F0 int64 }
type GenS18 struct { F0 int64 }
func genFn0() int8 {
    switch ("a" + "a") {
    case "a":
        var v5 int = 1
        _ = v5
    }
    var v6 uint8 = 1
    v7 := []int{42}
    var v8 bool = false
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 180
    var v3 int16 = -32768
    var v4 uint = 716
    v9 := GenS1{F0: (!true), F1: "", F2: 226}
    v10 := []int{}
    switch 42 {
    case 127:
        if (false || true) {
            if false {
                var v11 uint16 = 644
                switch -1 {
                case 127:
                    v10 = append(v10, 32768)
                case -436:
                    v4 = (v4 & (v4 + (v4 | v4)))
                case 128:
                    v12 := GenS1{F0: false, F1: ("hello" + "foo bar"), F2: 4294967295}
                    _ = v12
                }
                if (false && (1.5 != 0.0)) {
                    v13 := GenS1{F0: (100 > 42), F1: ("\n\t" + ("a" + ("hello" + ("foo bar" + "foo bar")))), F2: 1}
                    for fi14 := 0; fi14 < 12; fi14++ {
                        fmt.Println(v2, v2, fi14)
                        fmt.Println(v11)
                        if len(v10) > 0 {
                            v10[0] = (fi14 | fi14)
                        }
                        _ = fi14
                    }
                    _ = v13
                } else {
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v3 = v3
                        for fi16 := 0; fi16 < 17; fi16++ {
                            fmt.Println(v2, v4, v3)
                            v2 += (v2 - (v2 << ((v4) & 63)))
                            _ = fi16
                        }
                        var v17 string = "\n\t"
                        _ = fi15
                        _ = v17
                    }
                    v18 := GenS1{F0: false, F1: "hello", F2: 810}
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if false {
                            v20 := GenS2{F0: ""}
                            if (2 <= 778) {
                                if (!true) {
                                    v21 := GenS3{F0: 42, F1: -1, F2: genFn0()}
                                    _ = v21
                                } else {
                                    v4 = (v4 * (^(v4)))
                                    var v22 float64 = 3.14159
                                    _ = v22
                                }
                                switch fi19 {
                                case 83:
                                    if (-534 <= (-(v3))) {
                                        for fi23 := 0; fi23 < 20; fi23++ {
                                            if false {
                                                v11 = v11
                                                fmt.Println(v4, v4, v2)
                                                var v24 bool = (0.0 > 1e10)
                                                _ = v24
                                            }
                                            _ = fi23
                                        }
                                        v3 = v3
                                        v11 -= v11
                                    } else {
                                        var v25 float64 = -2.25
                                        _ = v25
                                    }
                                case -1:
                                    v26 := GenS1{F0: true, F1: "a", F2: 42}
                                    _ = v26
                                }
                                fmt.Println(v3, v3)
                            } else {
                                v2 += 255
                                var v27 bool = (true && false)
                                if (100 != 48) {
                                    v11 += v11
                                }
                                _ = v27
                            }
                            if (true && (!(!(!true)))) {
                                for fi28 := 0; fi28 < 17; fi28++ {
                                    switch (-(fi28)) {
                                    case -9223372036854775808:
                                        v29 := GenS2{F0: "hello"}
                                        _ = v29
                                    case -2147483648:
                                        v30 := GenS4{F0: 4294967295}
                                        if false {
                                            var v31 uint = (^(v4))
                                            v11 = (^(v11))
                                            _ = v31
                                        }
                                        var v32 float64 = -2.25
                                        var v33 string = "foo bar"
                                        v32 += (v32 + (v32 / (v32 + (1.5 * v32))))
                                        switch (v33 + "hello") {
                                        case "\n\t":
                                            v3 -= (^(v3))
                                        case "hello":
                                            var v34 uint8 = 128
                                            _ = v34
                                        case "":
                                            var v35 int64 = 42
                                            _ = v35
                                        default:
                                            v11 += v11
                                        }
                                        _ = v30
                                        _ = v32
                                        _ = v33
                                    case 586:
                                        var v36 float64 = (1.5 / (0.1 - float64(1e10)))
                                        _ = v36
                                    default:
                                        if true {
                                            var v37 bool = (!true)
                                            fmt.Println(v3, fi19)
                                            switch "a" {
                                            case "a":
                                                if (42 == 0) {
                                                    delete(v1, 2)
                                                    if (v37 || (3.14159 < 0.0)) {
                                                        if (499 >= 232) {
                                                            switch "foo bar" {
                                                            case "":
                                                                if v37 {
                                                                    var v38 string = ("\n\t" + "\n\t")
                                                                    _ = v38
                                                                }
                                                            case "hello":
                                                                var v39 int64 = 9223372036854775807
                                                                _ = v39
                                                            default:
                                                                fmt.Println(fi28, v2, v3)
                                                            }
                                                            if len(v10) >= 2 {
                                                                v10 = v10[1:len(v10)]
                                                            }
                                                        } else {
                                                            switch "foo bar" {
                                                            case "":
                                                                var v40 uint64 = 1
                                                                _ = v40
                                                            default:
                                                                fmt.Println(v37)
                                                            }
                                                            v41 := GenS2{F0: ("foo bar" + "\n\t")}
                                                            v2 += v2
                                                            _ = v41
                                                        }
                                                    }
                                                    switch "foo bar" {
                                                    case "":
                                                        var v42 uint32 = 42
                                                        _ = v42
                                                    case "foo bar":
                                                        v43 := GenS5{F0: (^(v4)), F1: uint64(v4)}
                                                        _ = v43
                                                    default:
                                                        var v44 float64 = 3.14159
                                                        _ = v44
                                                    }
                                                } else {
                                                    var v45 float64 = 0.1
                                                    fmt.Println(v4, fi28, v45)
                                                    v3 = (v3 * v3)
                                                    fmt.Println(v45)
                                                    var v46 int = fi28
                                                    _ = v45
                                                    _ = v46
                                                }
                                            }
                                            _ = v37
                                        } else {
                                            v2 -= v2
                                            var v47 uint64 = 2
                                            _ = v47
                                        }
                                        var v48 uint64 = 2
                                        if false {
                                            var v49 uint16 = (v11 << (((v4 >> ((v4) & 63))) & 63))
                                            switch "hello" {
                                            case "a":
                                                delete(v1, 224)
                                            }
                                            if true {
                                                _ = len(v10)
                                                v50 := GenS5{F0: v4, F1: v48}
                                                switch ("foo bar" + ("" + ("hello" + "\n\t"))) {
                                                case "hello":
                                                    switch 127 {
                                                    case 2147483647:
                                                        if true {
                                                            v11 = (v49 | uint16(v4))
                                                            v49 += (v49 + (v49 << (((v4 | v4)) & 63)))
                                                            v51 := GenS6{F0: 9223372036854775807, F1: v4}
                                                            _ = v51
                                                        }
                                                    }
                                                case "\n\t":
                                                    v2 = (v2 + (v2 % ((178) | 1)))
                                                }
                                                _ = v50
                                            }
                                            v1[4294967295] = false
                                            v52 := GenS7{F0: false, F1: -990, F2: 605}
                                            _ = v49
                                            _ = v52
                                        }
                                        if false {
                                            switch (fi28 / ((fi28) | 1)) {
                                            case -1:
                                                if (0 >= 0) {
                                                    fmt.Println(v48, v3)
                                                }
                                            default:
                                                v53 := GenS4{F0: 2}
                                                if true {
                                                    if (genFn0() >= 0) {
                                                        v2 = v2
                                                        fmt.Println(v11, fi28, fi28)
                                                        v54 := GenS5{F0: v4, F1: v48}
                                                        _ = v54
                                                    } else {
                                                        v55 := v1[541]
                                                        _ = v55
                                                    }
                                                    _ = len(v10)
                                                } else {
                                                    switch ("foo bar" + "foo bar") {
                                                    case "hello":
                                                        if false {
                                                            if (!true) {
                                                                fmt.Println(fi28, v3)
                                                            } else {
                                                                if (!false) {
                                                                    v56 := GenS8{F0: -190, F1: (v2 % ((v2) | 1)), F2: (v3 ^ (v3 - v3))}
                                                                    if false {
                                                                        switch -917 {
                                                                        case -1:
                                                                            v57 := GenS9{F0: 42, F1: 2147483647}
                                                                            _ = v57
                                                                        case 42:
                                                                            v58 := v1[0]
                                                                            _ = v58
                                                                        }
                                                                        fmt.Println(fi19)
                                                                    } else {
                                                                        v2 += (v2 << (((v4 / (((v4 + 286)) | 1))) & 63))
                                                                    }
                                                                    var v59 uint64 = v48
                                                                    _ = v56
                                                                    _ = v59
                                                                }
                                                                v60 := GenS10{F0: v4, F1: (v2 / (((v2 << ((v4) & 63))) | 1)), F2: -8}
                                                                switch "a" {
                                                                case "":
                                                                    v61 := GenS11{F0: (fi19 << ((uint(v11)) & 63))}
                                                                    _ = v61
                                                                }
                                                                _ = v60
                                                            }
                                                            delete(v1, 499)
                                                            switch ("foo bar" + "\n\t") {
                                                            case "hello":
                                                                v11 += 32767
                                                            case "":
                                                                v62 := GenS6{F0: -9223372036854775808, F1: 65535}
                                                                _ = v62
                                                            case "\n\t":
                                                                v63 := v1[0]
                                                                _ = v63
                                                            default:
                                                                fmt.Println(v11, fi28)
                                                            }
                                                        }
                                                    case "a":
                                                        v11 -= 32767
                                                    default:
                                                        fmt.Println(v48)
                                                    }
                                                }
                                                v11 += (v11 / ((uint16(v48)) | 1))
                                                _ = v53
                                            }
                                            v3 = int16(fi19)
                                        }
                                        switch fi28 {
                                        case 1:
                                            fmt.Println(fi19, v48)
                                        case 622:
                                            delete(v1, 2)
                                        case 128:
                                            switch (fi19 % (((fi28 << ((v4) & 63))) | 1)) {
                                            case -9223372036854775808:
                                                _ = len(v1)
                                            case 9223372036854775807:
                                                if true {
                                                    fmt.Println(v48, v2, v48)
                                                    switch (fi28 | fi19) {
                                                    case 718:
                                                        v64 := GenS8{F0: -895, F1: v2, F2: -621}
                                                        _ = v64
                                                    default:
                                                        v65 := GenS1{F0: true, F1: "", F2: v48}
                                                        _ = v65
                                                    }
                                                    if ((v2 <= v2) && false) {
                                                        var v66 float32 = (3.14159 + 3.14159)
                                                        _ = v66
                                                    } else {
                                                        delete(v1, 2147483647)
                                                        var v67 uint8 = uint8(v11)
                                                        var v68 int8 = -1
                                                        fmt.Println(fi28)
                                                        v69 := v1[438]
                                                        var v70 int = 42
                                                        _ = v67
                                                        _ = v68
                                                        _ = v69
                                                        _ = v70
                                                    }
                                                } else {
                                                    switch (^(fi19)) {
                                                    case -2147483648:
                                                        var v71 uint8 = v2
                                                        _ = v71
                                                    case 32767:
                                                        _ = len(v1)
                                                    }
                                                    var v72 int = (fi19 / ((-2147483648) | 1))
                                                    var v73 uint32 = 4294967295
                                                    _ = v72
                                                    _ = v73
                                                }
                                                v74 := GenS12{F0: 0}
                                                _ = v74
                                            case 0:
                                                delete(v1, 2147483648)
                                            default:
                                                switch (fi19 ^ fi19) {
                                                case 9223372036854775807:
                                                    switch ("foo bar" + ("hello" + ("a" + "\n\t"))) {
                                                    case "foo bar":
                                                        var v75 int64 = -9223372036854775808
                                                        _ = v75
                                                    default:
                                                        switch 682 {
                                                        case 468:
                                                            v76 := GenS4{F0: uint32(v48)}
                                                            _ = v76
                                                        case 584:
                                                            v11 = (v11 >> (((v4 % ((401) | 1))) & 63))
                                                        case 32767:
                                                            v48 += v48
                                                        }
                                                    }
                                                case 32767:
                                                    if ((!true) || true) {
                                                        v4 += v4
                                                        var v77 int32 = 0
                                                        _ = v77
                                                    } else {
                                                        var v78 int64 = 1
                                                        _ = v78
                                                    }
                                                    switch "a" {
                                                    case "":
                                                        v79 := v1[65535]
                                                        _ = v79
                                                    case "foo bar":
                                                        _ = len(v1)
                                                    case "hello":
                                                        v3 = v3
                                                    default:
                                                        if false {
                                                            var v80 int64 = 42
                                                            _ = v80
                                                        } else {
                                                            v48 += 2
                                                            v2 = 2
                                                            if (1 <= v11) {
                                                                var v81 int = 1
                                                                v10 = append(v10, -504)
                                                                _ = v81
                                                            } else {
                                                                fmt.Println(v4)
                                                            }
                                                        }
                                                    }
                                                    v3 = (-(v3))
                                                case 975:
                                                    if false {
                                                        v82 := GenS13{F0: fi19}
                                                        var v83 int = 969
                                                        fmt.Println(v83, v83)
                                                        _ = v82
                                                        _ = v83
                                                    }
                                                }
                                            }
                                        }
                                        _ = len(v1)
                                        v11 += (v11 | v11)
                                        _ = v48
                                    }
                                    var v84 int64 = -9223372036854775808
                                    var v85 int = (fi28 - fi19)
                                    v11 -= (v11 ^ v11)
                                    if ((1e10 <= 0.1) && (float64((((1.5 * 1e10) / 1.0) - ((3.14159 - 0.1) - float32(0.1)))) > ((3.14159 * 3.14159) - float64(float32((0.0 / -2.25)))))) {
                                        v2 -= v2
                                        v84 += 42
                                    } else {
                                        var v86 uint8 = (v2 - v2)
                                        v3 += -649
                                        v4 += uint(v11)
                                        _ = v86
                                    }
                                    _ = fi28
                                    _ = v84
                                    _ = v85
                                }
                                v11 += v11
                            }
                            _ = v20
                        } else {
                            v2 = (v2 + (v2 | 1))
                            var v87 float64 = 0.0
                            _ = v87
                        }
                        v3 -= 0
                        v11 -= v11
                        _ = fi19
                    }
                    _ = v18
                }
                _ = v11
            }
            if false {
                v88 := GenS14{F0: 407, F1: (v3 >> ((v4) & 63)), F2: 419}
                var v89 float32 = ((float32((0.0 - 1e10)) + (1e10 - ((-2.25 + 1e10) - -2.25))) * 0.1)
                switch "a" {
                case "foo bar":
                    v10 = append(v10, -612)
                case "":
                    for fi90 := 0; fi90 < 12; fi90++ {
                        for fi91 := 0; fi91 < 19; fi91++ {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            _ = fi91
                        }
                        _ = fi90
                    }
                case "a":
                    v10 = append(v10, 32767)
                }
                _ = v88
                _ = v89
            } else {
                var v92 int = 2147483647
                _ = len(v1)
                _ = v92
            }
            v2 = v2
        }
    }
    for fi93 := 0; fi93 < 10; fi93++ {
        var v94 uint64 = 42
        v95 := GenS11{F0: (^(fi93))}
        v96 := GenS15{F0: v2, F1: 42, F2: (v4 << ((v4) & 63))}
        _ = fi93
        _ = v94
        _ = v95
        _ = v96
    }
    v97 := GenS16{F0: false}
    v2 += (v2 & v2)
    for fi98 := 0; fi98 < 3; fi98++ {
        v2 = 42
        v2 = (^(v2))
        _ = fi98
    }
    v99 := GenS17{F0: 42}
    fmt.Println(v2, v2)
    for fi100 := 0; fi100 < 19; fi100++ {
        _ = len(v1)
        _ = fi100
    }
    switch 128 {
    case 127:
        for fi101 := 0; fi101 < 19; fi101++ {
            if (!false) {
                v102 := GenS18{F0: 42}
                if (uint16(v2) == 1) {
                    for fi103 := 0; fi103 < 20; fi103++ {
                        var v104 bool = (127 <= (v2 / (((v2 + v2)) | 1)))
                        v104 = v104
                        v104 = ((v2 ^ (^(v2))) != v2)
                        _ = fi103
                        _ = v104
                    }
                    if (!((!(!(float64(-2.25) > (0.1 + 1e10)))) || ((true && (549 >= 1)) || (!((uint16(v4) < uint16(v4)) && false))))) {
                        v2 = (v2 * (^(v2)))
                    }
                } else {
                    var v105 bool = (!((0.0 != 1e10) && (false || (!(false && false)))))
                    _ = v105
                }
                _ = v102
            }
            v2 -= v2
            _ = fi101
        }
    case -9223372036854775808:
        v106 := GenS18{F0: 1}
        _ = v106
    }
    var v107 uint32 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v107:", v107)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v107
    _ = v2
    _ = v3
    _ = v4
    _ = v9
    _ = v97
    _ = v99
}
