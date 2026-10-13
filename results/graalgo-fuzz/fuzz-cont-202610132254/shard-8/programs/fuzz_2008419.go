package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int32; F1 int64; F2 uint16 }
type GenS8 struct { F0 int8; F1 int32 }
type GenS9 struct { F0 int8; F1 string }
func genFn0() int16 {
    var v6 int16 = -32768
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 bool = (259 == (fi7 | (fi7 | -296)))
        _ = fi7
        _ = v8
    }
    v6 += (v6 | v6)
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            v11 := []int{fi10, (fi10 * (fi10 | (-(fi9))))}
            if (42 <= 0) {
                v6 = (v6 * (^(v6)))
                _ = len(v11)
                fmt.Println(v6)
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    _ = v6
    return v6
}

func main() {
    var v1 string = ""
    var v2 int8 = 112
    var v3 uint8 = 12
    v4 := make(map[uint32]string)
    var v5 uint32 = 289
    if true {
        _ = len(v4)
        fmt.Println(v2)
    } else {
        switch (v1 + "hello") {
        case "\n\t":
            if false {
                var v12 uint = 255
                switch "" {
                case "hello":
                    switch (v1 + v1) {
                    case "\n\t":
                        var v13 int32 = 2147483647
                        _ = v13
                    }
                case "":
                    var v14 int16 = genFn0()
                    _ = v14
                }
                v3 = (v3 ^ v3)
                _ = v12
            }
        }
        v15 := v4[65535]
        switch 9223372036854775807 {
        case 371:
            switch -1 {
            case -907:
                var v16 float64 = (3.14159 + 0.0)
                _ = v16
            case -9223372036854775808:
                v17 := v4[v5]
                _ = v17
            case 42:
                if ((v5 ^ 2) > (v5 ^ v5)) {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            v20 := GenS1{F0: 56, F1: 0, F2: (v5 >> ((65535) & 63))}
                            v21 := GenS1{F0: 65535, F1: 282, F2: 2147483648}
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        _ = fi18
                    }
                    if (v2 < (-(v2))) {
                        if false {
                            fmt.Println(v15, v1)
                            var v22 string = "\n\t"
                            fmt.Println(v3)
                            _ = v22
                        } else {
                            v23 := GenS2{F0: 24, F1: genFn0(), F2: -889}
                            switch -479 {
                            case 883:
                                v24 := []int{-571, 925, -1}
                                _ = v24
                            }
                            _ = v23
                        }
                    } else {
                        v15 = (((v15 + v15) + "\n\t") + "hello")
                    }
                }
            }
        case -1:
            var v25 float64 = (0.1 / -2.25)
            _ = v25
        }
        _ = v15
    }
    fmt.Println(v3)
    if (1 != 9223372036854775807) {
        v26 := []int{416, 128}
        switch -1 {
        case 32767:
            var v27 string = ("foo bar" + "hello")
            _ = v27
        case 9223372036854775807:
            if (1e10 < -2.25) {
                if false {
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 uint64 = 0
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        v4[(v5 << ((383) & 63))] = "hello"
                        _ = fi28
                        _ = v29
                    }
                    if (false || true) {
                        v26 = append(v26, -9223372036854775808)
                        v1 = "\n\t"
                    } else {
                        for fi30 := 0; fi30 < 8; fi30++ {
                            if true {
                                v3 += v3
                                fmt.Println(v2)
                                if (!(32767 == 0)) {
                                    var v31 int = (fi30 << ((0) & 63))
                                    _ = v31
                                } else {
                                    var v32 uint = 255
                                    fmt.Println(v1, v1)
                                    if (!(v5 == (v5 | (^(v5))))) {
                                        fmt.Println(v1, v2, v1)
                                        v33 := GenS3{F0: v1, F1: 32767}
                                        fmt.Println(v5)
                                        _ = v33
                                    } else {
                                        switch v1 {
                                        case "\n\t":
                                            v5 = uint32(v32)
                                        case "hello":
                                            v5 += 531
                                        }
                                    }
                                    _ = v32
                                }
                            }
                            v4[(v5 % ((v5) | 1))] = v1
                            _ = fi30
                        }
                    }
                    var v34 int16 = 0
                    _ = v34
                } else {
                    _ = len(v4)
                }
                v1 = (("a" + "foo bar") + ("a" + ("\n\t" + "a")))
                for fi35 := 0; fi35 < 3; fi35++ {
                    if true {
                        switch (-(fi35)) {
                        case -813:
                            var v36 int = 0
                            _ = v36
                        case 42:
                            v37 := GenS1{F0: 32767, F1: 4294967295, F2: 0}
                            _ = v37
                        case 32768:
                            delete(v4, (^(v5)))
                        }
                        if false {
                            v1 = ""
                            v38 := GenS4{F0: (!false)}
                            _ = v38
                        }
                        switch v1 {
                        case "\n\t":
                            if ((v2 % ((127) | 1)) <= v2) {
                                var v39 int = (-(fi35))
                                var v40 uint = 646
                                if (v5 != (v5 << (((v40 >> ((255) & 63))) & 63))) {
                                    for fi41 := 0; fi41 < 8; fi41++ {
                                        v42 := GenS5{F0: uint32(v40)}
                                        var v43 uint = (v40 & (v40 | (v40 + (v40 & (v40 << ((v40) & 63))))))
                                        if ((v43 | v43) >= (v43 * v40)) {
                                            v1 = ((((v1 + "") + (v1 + "hello")) + v1) + v1)
                                            delete(v4, 885)
                                        }
                                        v40 = (v40 * v43)
                                        _ = fi41
                                        _ = v42
                                        _ = v43
                                    }
                                }
                                _ = v39
                                _ = v40
                            }
                        }
                    }
                    switch v1 {
                    case "hello":
                        if len(v26) > 0 {
                            v26[0] = -2147483648
                        }
                    }
                    if len(v26) > 0 {
                        v26[0] = fi35
                    }
                    _ = fi35
                }
            } else {
                if true {
                    switch 2147483647 {
                    case 127:
                        if (true && (!false)) {
                            var v44 uint64 = uint64(v3)
                            _ = len(v4)
                            _ = v44
                        } else {
                            v45 := GenS3{F0: v1, F1: 0}
                            for fi46 := 0; fi46 < 9; fi46++ {
                                v47 := GenS1{F0: 32767, F1: 1, F2: (v5 * v5)}
                                _ = fi46
                                _ = v47
                            }
                            var v48 float64 = -2.25
                            _ = v45
                            _ = v48
                        }
                    case -285:
                        for fi49 := 0; fi49 < 5; fi49++ {
                            if len(v26) > 0 {
                                v26[0] = (fi49 % ((42) | 1))
                            }
                            delete(v4, (v5 | v5))
                            _ = fi49
                        }
                    }
                }
                for fi50 := 0; fi50 < 17; fi50++ {
                    var v51 uint64 = 28
                    fmt.Println(v2)
                    delete(v4, (v5 % ((2) | 1)))
                    _ = fi50
                    _ = v51
                }
                var v52 string = v1
                _ = v52
            }
        case 1:
            var v53 uint64 = 4294967295
            _ = v53
        default:
            for fi54 := 0; fi54 < 1; fi54++ {
                fmt.Println(v3, v1)
                _ = fi54
            }
        }
        v2 -= v2
        _ = v26
    } else {
        fmt.Println(v1)
    }
    var v55 uint64 = uint64(v3)
    var v56 uint8 = v3
    v56 += 178
    v57 := GenS5{F0: v5}
    fmt.Println(v55, v3)
    for fi58 := 0; fi58 < 8; fi58++ {
        var v59 uint16 = 767
        if (255 < 4294967295) {
            fmt.Println(v59, v55)
            var v60 uint64 = v55
            _ = v60
        }
        _ = fi58
        _ = v59
    }
    if false {
        if true {
            v61 := []int{0, 43, -2147483648}
            _ = v61
        } else {
            for fi62 := 0; fi62 < 9; fi62++ {
                fmt.Println(v55, v2, fi62)
                var v63 int16 = genFn0()
                for fi64 := 0; fi64 < 20; fi64++ {
                    var v65 int8 = 42
                    var v66 bool = false
                    v2 = (v2 * (v65 ^ -1))
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                _ = fi62
                _ = v63
            }
            v3 -= v56
        }
        switch "foo bar" {
        case "a":
            v55 = v55
        case "":
            for fi67 := 0; fi67 < 9; fi67++ {
                fmt.Println(v2, v56)
                _ = fi67
            }
        case "foo bar":
            for fi68 := 0; fi68 < 11; fi68++ {
                v56 = 128
                if true {
                    switch (fi68 ^ 32768) {
                    case 128:
                        v69 := GenS6{F0: (^(v5))}
                        _ = v69
                    default:
                        v70 := GenS5{F0: v5}
                        _ = v70
                    }
                    for fi71 := 0; fi71 < 5; fi71++ {
                        var v72 float64 = (3.14159 * 1e10)
                        _ = fi71
                        _ = v72
                    }
                    for fi73 := 0; fi73 < 5; fi73++ {
                        if true {
                            var v74 float64 = 1.5
                            _ = v74
                        }
                        var v75 int32 = 0
                        v75 += (v75 % ((v75) | 1))
                        _ = fi73
                        _ = v75
                    }
                } else {
                    var v76 uint16 = 65535
                    _ = v76
                }
                v77 := GenS7{F0: 1, F1: 42, F2: 42}
                _ = fi68
                _ = v77
            }
        default:
            v78 := []int{133, 952}
            _ = v78
        }
    }
    switch 127 {
    case 128:
        if (int32(v55) == -1) {
            v79 := GenS8{F0: v2, F1: 0}
            for fi80 := 0; fi80 < 8; fi80++ {
                fmt.Println(v3)
                for fi81 := 0; fi81 < 17; fi81++ {
                    v82 := GenS5{F0: (v5 ^ (v5 + (v5 % ((v5) | 1))))}
                    v83 := v4[v5]
                    v83 = (v1 + "a")
                    _ = fi81
                    _ = v82
                    _ = v83
                }
                _ = fi80
            }
            _ = v79
        }
    }
    for fi84 := 0; fi84 < 4; fi84++ {
        switch v1 {
        case "\n\t":
            v5 = (v5 & 2147483647)
        case "foo bar":
            if ((!(!(319 > -363))) && (-56 >= 42)) {
                v56 += (^(v3))
                for fi85 := 0; fi85 < 14; fi85++ {
                    v86 := GenS9{F0: (v2 * v2), F1: ((v1 + v1) + v1)}
                    var v87 uint64 = v55
                    var v88 bool = false
                    _ = fi85
                    _ = v86
                    _ = v87
                    _ = v88
                }
                v89 := []int{fi84, -288, fi84}
                _ = v89
            }
        default:
            var v90 uint16 = 2
            _ = v90
        }
        if (false && (fi84 >= fi84)) {
            fmt.Println(v1)
        } else {
            var v91 int64 = 0
            v91 = (v91 & (v91 % ((v91) | 1)))
            _ = v91
        }
        var v92 int16 = genFn0()
        _ = fi84
        _ = v92
    }
    for fi93 := 0; fi93 < 8; fi93++ {
        var v94 uint = 255
        _ = fi93
        _ = v94
    }
    var v95 uint64 = v55
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
    _ = v56
    _ = v57
    _ = v95
}
