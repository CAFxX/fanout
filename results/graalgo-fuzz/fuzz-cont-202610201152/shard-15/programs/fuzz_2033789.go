package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 int8 }
func genFn0() string {
    v5 := []int{127, -46, 42, 32767}
    v5 = append(v5, 32767)
    for fi6 := 0; fi6 < 2; fi6++ {
        if false {
            for fi7 := 0; fi7 < 8; fi7++ {
                var v8 int8 = 42
                if len(v5) > 0 {
                    v5[0] = fi6
                }
                _ = fi7
                _ = v8
            }
            switch "a" {
            case "":
                var v9 uint32 = 2
                _ = v9
            case "hello":
                var v10 bool = false
                _ = v10
            }
        } else {
            if (true || false) {
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 uint16 = 266
                    _ = fi11
                    _ = v12
                }
                var v13 uint64 = 498
                if (!(!(!(1 <= 255)))) {
                    if false {
                        switch "" {
                        case "hello":
                            var v14 string = "\n\t"
                            _ = v14
                        case "foo bar":
                            v13 = (v13 & (v13 | v13))
                        }
                    } else {
                        fmt.Println(fi6, v13)
                        v15 := GenS1{F0: 121}
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 17; fi16++ {
                        _ = len(v5)
                        v13 += 4294967295
                        if false {
                            var v17 bool = (0 < 42)
                            v5 = append(v5, (fi6 | fi6))
                            var v18 int = fi6
                            var v19 bool = v17
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        _ = fi16
                    }
                    _ = len(v5)
                }
                _ = v13
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                v21 := GenS2{F0: 42, F1: 42}
                if (true && (-1 == 42)) {
                    fmt.Println(fi20)
                    fmt.Println(fi20)
                    fmt.Println(fi6, fi20)
                }
                var v22 bool = (255 > 620)
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        var v23 uint32 = 657
        _ = fi6
        _ = v23
    }
    for fi24 := 0; fi24 < 8; fi24++ {
        for fi25 := 0; fi25 < 12; fi25++ {
            if (true && false) {
                if (-2.25 < float32((1e10 + (((0.1 * 1.5) + ((0.1 / 1.5) / float64(-2.25))) * float64(0.1))))) {
                    var v26 int16 = 32767
                    _ = v26
                } else {
                    var v27 uint64 = 42
                    var v28 float64 = (3.14159 + 3.14159)
                    fmt.Println(fi25, fi24)
                    _ = v27
                    _ = v28
                }
                fmt.Println(fi24)
            }
            var v29 float32 = -2.25
            _ = fi25
            _ = v29
        }
        switch ("\n\t" + ("foo bar" + "")) {
        case "hello":
            if (true && (42 == 32767)) {
                if (65535 != 489) {
                    var v30 string = ((("\n\t" + "a") + "a") + "foo bar")
                    _ = v30
                } else {
                    v31 := GenS1{F0: 128}
                    var v32 uint64 = 992
                    _ = v31
                    _ = v32
                }
            }
        }
        v5 = append(v5, (-(fi24)))
        _ = fi24
    }
    _ = v5
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 1
    var v3 float32 = 0.0
    v4 := make(map[int]int)
    v33 := GenS1{F0: 2}
    for fi34 := 0; fi34 < 10; fi34++ {
        if (false && false) {
            var v35 float64 = (1.5 / (float64(v3) + -2.25))
            _ = v35
        } else {
            var v36 int = fi34
            v2 = (v2 & (v2 % ((2) | 1)))
            if true {
                delete(v1, "hello")
            }
            _ = v36
        }
        if true {
            switch ("a" + "\n\t") {
            case "\n\t":
                v37 := GenS3{F0: 42, F1: 517}
                _ = v37
            case "hello":
                switch (fi34 | fi34) {
                case 0:
                    if false {
                        fmt.Println(v2, fi34, v2)
                        var v38 string = "hello"
                        _ = v38
                    } else {
                        v39 := GenS4{F0: 886, F1: (32767 <= 65535)}
                        v2 -= (v2 ^ v2)
                        v2 = (v2 >> ((255) & 63))
                        _ = v39
                    }
                }
            default:
                if false {
                    v40 := GenS5{F0: fi34}
                    switch "\n\t" {
                    case "":
                        for fi41 := 0; fi41 < 11; fi41++ {
                            if (1 <= 1) {
                                var v42 uint64 = v2
                                delete(v1, "")
                                _ = v42
                            }
                            _ = fi41
                        }
                    default:
                        var v43 int32 = -1
                        _ = v43
                    }
                    var v44 int8 = 34
                    _ = v40
                    _ = v44
                } else {
                    fmt.Println(fi34, v2, v2)
                    var v45 int = (fi34 << ((18446744073709551615) & 63))
                    for fi46 := 0; fi46 < 11; fi46++ {
                        delete(v4, 1)
                        _ = fi46
                    }
                    _ = v45
                }
            }
            switch -842 {
            case -310:
                v2 = (v2 ^ (v2 + (^(v2))))
            case 32768:
                for fi47 := 0; fi47 < 11; fi47++ {
                    var v48 int64 = 1
                    v48 = (v48 | v48)
                    _ = fi47
                    _ = v48
                }
            }
            if true {
                for fi49 := 0; fi49 < 18; fi49++ {
                    fmt.Println(v2, v2, fi34)
                    v3 = 3.14159
                    switch -9223372036854775808 {
                    case 127:
                        v50 := GenS2{F0: 65535, F1: 18446744073709551615}
                        _ = v50
                    case 1:
                        fmt.Println(fi34, fi49, fi49)
                    }
                    _ = fi49
                }
                v51 := GenS5{F0: (fi34 * -313)}
                fmt.Println(v3)
                _ = v51
            } else {
                var v52 int16 = -305
                _ = v52
            }
        }
        if false {
            var v53 int64 = -1
            var v54 uint = 65535
            _ = v53
            _ = v54
        } else {
            var v55 uint32 = 2147483648
            v56 := GenS6{F0: uint64(v55), F1: (v55 * 115), F2: 42}
            v57 := GenS6{F0: v2, F1: (^(v55)), F2: 574}
            _ = v55
            _ = v56
            _ = v57
        }
        _ = fi34
    }
    v2 = v2
    var v58 int64 = 962
    v2 -= (v2 & v2)
    v59 := GenS2{F0: 245, F1: (v2 - (v2 ^ v2))}
    if true {
        for fi60 := 0; fi60 < 18; fi60++ {
            fmt.Println(v58, fi60)
            _ = fi60
        }
        var v61 string = "hello"
        _ = v61
    } else {
        if (817 < 26) {
            for fi62 := 0; fi62 < 4; fi62++ {
                v63 := []int{(fi62 & -2147483648), (fi62 % (((-(fi62))) | 1)), -1, fi62}
                for fi64 := 0; fi64 < 1; fi64++ {
                    fmt.Println(v2, fi62, v2)
                    var v65 bool = false
                    v63 = append(v63, fi62)
                    v66 := GenS5{F0: fi62}
                    _ = fi64
                    _ = v65
                    _ = v66
                }
                if len(v63) > 0 {
                    v63[0] = fi62
                }
                _ = fi62
                _ = v63
            }
            fmt.Println(v3, v3, v2)
            v67 := GenS6{F0: v2, F1: 2147483647, F2: 42}
            _ = v67
        }
    }
    switch -595 {
    case 127:
        v68 := v4[32768]
        _ = v68
    case 32768:
        if false {
            delete(v4, 127)
        } else {
            var v69 uint64 = v2
            var v70 int = 2147483647
            if false {
                for fi71 := 0; fi71 < 8; fi71++ {
                    fmt.Println(v2, v58)
                    if false {
                        v72 := GenS4{F0: 42, F1: ((v58 != (v58 - v58)) && ((("\n\t" + ("a" + "a")) + "hello") <= "foo bar"))}
                        _ = v72
                    } else {
                        var v73 bool = true
                        if (32767 <= uint16(v2)) {
                            v73 = (255 == 748)
                        } else {
                            fmt.Println(v69, v2)
                        }
                        _ = v73
                    }
                    if (((fi71 << ((1) & 63)) == (fi71 | (-(fi71)))) || (v69 >= (v2 - 525))) {
                        var v74 uint32 = 42
                        for fi75 := 0; fi75 < 5; fi75++ {
                            var v76 string = ("a" + "hello")
                            if false {
                                if (int16(v74) < 755) {
                                    if (42 == 0) {
                                        v77 := []int{fi71, (fi71 << ((2) & 63)), v70, (v70 * fi71)}
                                        v78 := GenS7{F0: 42}
                                        v69 = 0
                                        _ = v77
                                        _ = v78
                                    }
                                } else {
                                    var v79 uint = 18446744073709551615
                                    v80 := GenS2{F0: (^(v79)), F1: 633}
                                    _ = v79
                                    _ = v80
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v81 bool = (42 != 888)
                                    _ = v81
                                case "a":
                                    v70 -= (fi71 - fi75)
                                case "hello":
                                    var v82 bool = (!false)
                                    fmt.Println(v70)
                                    _ = v82
                                default:
                                    var v83 uint64 = v2
                                    _ = v83
                                }
                            }
                            var v84 int32 = -199
                            _ = fi75
                            _ = v76
                            _ = v84
                        }
                        v74 += (v74 ^ 650)
                        _ = v74
                    }
                    _ = fi71
                }
                fmt.Println(v70, v70, v3)
                v2 += v2
            } else {
                fmt.Println(v58)
                if (("\n\t" > ("hello" + "")) || ((v69 ^ 42) >= v69)) {
                    v2 -= v2
                    if (((v3 / 0.1) - -2.25) == v3) {
                        for fi85 := 0; fi85 < 4; fi85++ {
                            v86 := []int{v70}
                            _ = fi85
                            _ = v86
                        }
                        v2 = 226
                        var v87 uint64 = (v2 & (v69 & v2))
                        _ = v87
                    } else {
                        delete(v1, "hello")
                        var v88 int32 = -2147483648
                        _ = v88
                    }
                }
            }
            _ = v69
            _ = v70
        }
    case 4294967295:
        var v89 int8 = -78
        _ = v89
    default:
        v3 -= (3.14159 - v3)
    }
    for fi90 := 0; fi90 < 10; fi90++ {
        var v91 uint = 4294967295
        v3 -= ((3.14159 / 0.1) / ((v3 / v3) / (v3 / 3.14159)))
        v92 := v4[fi90]
        _ = fi90
        _ = v91
        _ = v92
    }
    v3 += v3
    for fi93 := 0; fi93 < 18; fi93++ {
        v58 -= v58
        v94 := []int{}
        _ = fi93
        _ = v94
    }
    v95 := GenS8{F0: int8(v2)}
    switch 4294967295 {
    case -9223372036854775808:
        var v96 string = ""
        _ = v96
    case 32767:
        v97 := GenS6{F0: (v2 >> ((236) & 63)), F1: 42, F2: 308}
        _ = v97
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v58:", v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v58
    _ = v59
    _ = v95
}
