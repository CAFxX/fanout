package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 bool; F1 int8 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint64, p1 int16, p2 uint32) int8 {
    v6 := []int{-247, 0, 2147483647}
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 5; fi8++ {
        switch (fi8 * (fi8 / (((fi8 / ((9223372036854775807) | 1))) | 1))) {
        case 32767:
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := GenS1{F0: fi9, F1: p1}
                var v11 uint16 = 71
                v12 := GenS2{F0: -1, F1: false, F2: -128}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        default:
            p2 = (^(p2))
        }
        var v13 uint = 459
        if true {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v14 := GenS3{F0: (p2 / (((p2 ^ uint32(p0))) | 1))}
            _ = v14
        }
        _ = fi8
        _ = v13
    }
    v15 := GenS3{F0: (p2 * (p2 + (p2 * p2)))}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := []uint32{2}
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 65535
    v4 := make(map[uint32]string)
    var v5 uint16 = (^(v3))
    var v16 int64 = -9223372036854775808
    fmt.Println(v5)
    var v17 int32 = 0
    var v18 uint32 = 2147483648
    switch 9223372036854775807 {
    case 779:
        for fi19 := 0; fi19 < 5; fi19++ {
            v17 = v17
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi19
        }
    case 42:
        switch 4294967295 {
        case -838:
            if false {
                for fi20 := 0; fi20 < 13; fi20++ {
                    if (("foo bar" == "a") || false) {
                        if (true && (1 != 392)) {
                            for fi21 := 0; fi21 < 11; fi21++ {
                                if true {
                                    v2 += (v2 | v2)
                                    var v22 uint16 = (v3 / (((v3 % ((42) | 1))) | 1))
                                    v1 = append(v1, 218)
                                    _ = v22
                                }
                                switch (^(fi21)) {
                                case -666:
                                    switch (fi20 << ((2) & 63)) {
                                    case 42:
                                        v1 = append(v1, v18)
                                    case -584:
                                        v5 = 255
                                    case 9223372036854775807:
                                        if true {
                                            v23 := GenS1{F0: fi20, F1: -210}
                                            var v24 uint16 = (v3 - 440)
                                            _ = v23
                                            _ = v24
                                        } else {
                                            switch "foo bar" {
                                            case "\n\t":
                                                var v25 uint64 = 4294967295
                                                _ = v25
                                            default:
                                                if (!(true || false)) {
                                                    v3 = (v5 * v5)
                                                    fmt.Println(fi21)
                                                } else {
                                                    v5 -= uint16(v18)
                                                }
                                            }
                                            var v26 uint32 = 42
                                            var v27 uint16 = (v3 - (^(v3)))
                                            _ = v26
                                            _ = v27
                                        }
                                    default:
                                        v28 := GenS4{F0: 1}
                                        var v29 int = (-(fi21))
                                        switch (((("foo bar" + (("hello" + "foo bar") + "\n\t")) + "") + "foo bar") + "a") {
                                        case "\n\t":
                                            v29 = fi20
                                        case "":
                                            var v30 int = -131
                                            _ = v30
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = 2147483648
                                            }
                                        }
                                        _ = v28
                                        _ = v29
                                    }
                                case 32767:
                                    v17 = (v17 + v17)
                                }
                                v18 += (^(v18))
                                _ = fi21
                            }
                            v31 := GenS2{F0: (v17 - (v17 >> ((4294967295) & 63))), F1: true, F2: 0}
                            switch fi20 {
                            case -2147483648:
                                v32 := GenS5{F0: 1, F1: -300}
                                _ = v32
                            case 373:
                                switch -715 {
                                case 4294967295:
                                    fmt.Println(v3, v18)
                                case 155:
                                    var v33 int32 = (v17 >> ((484) & 63))
                                    _ = v33
                                case -1:
                                    if (v5 < v5) {
                                        if (v5 == 57) {
                                            v34 := GenS4{F0: 42}
                                            v4[v18] = ("foo bar" + "hello")
                                            _ = v34
                                        }
                                        var v35 uint64 = 194
                                        _ = v35
                                    }
                                default:
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case 128:
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    if (45 <= int8(v18)) {
                                        v3 -= v3
                                    }
                                    var v37 int32 = (v17 >> ((4294967295) & 63))
                                    _ = fi36
                                    _ = v37
                                }
                            }
                            _ = v31
                        } else {
                            v2 = (^(v16))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v18 -= (^(v18))
                        }
                        switch fi20 {
                        case 127:
                            fmt.Println(v17, v3, v3)
                        case 0:
                            var v38 bool = false
                            _ = v38
                        default:
                            switch (fi20 + 2147483647) {
                            case 4294967295:
                                v5 = v5
                            case 2147483647:
                                if (!true) {
                                    v39 := GenS3{F0: v18}
                                    _ = v39
                                }
                            case 104:
                                v40 := GenS6{F0: false, F1: 127}
                                _ = v40
                            }
                        }
                    } else {
                        var v41 int = (fi20 - (fi20 << ((1) & 63)))
                        fmt.Println(v3, v17)
                        v3 = (^(v3))
                        _ = v41
                    }
                    _ = fi20
                }
                v42 := GenS7{F0: 42}
                switch ("foo bar" + ("" + "")) {
                case "foo bar":
                    fmt.Println(v18, v3, v16)
                case "hello":
                    v43 := GenS3{F0: (^(v18))}
                    _ = v43
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v42
            } else {
                var v44 float64 = float64(-2.25)
                fmt.Println(v2, v17, v3)
                if false {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v17 = v17
                        v46 := GenS3{F0: v18}
                        for fi47 := 0; fi47 < 11; fi47++ {
                            var v48 uint16 = (v5 + v5)
                            _ = fi47
                            _ = v48
                        }
                        _ = fi45
                        _ = v46
                    }
                    if (true && (false || true)) {
                        v49 := GenS4{F0: 18446744073709551615}
                        v50 := GenS8{F0: 128}
                        _ = v49
                        _ = v50
                    }
                    var v51 int64 = v16
                    _ = v51
                } else {
                    v18 += (^(v18))
                }
                _ = v44
            }
        case 1:
            for fi52 := 0; fi52 < 9; fi52++ {
                var v53 uint32 = 2
                for fi54 := 0; fi54 < 7; fi54++ {
                    v16 = (v16 + v16)
                    var v55 int8 = -37
                    v18 -= v18
                    if (!(v17 >= (v17 % ((-1) | 1)))) {
                        v56 := GenS9{F0: 42}
                        v57 := GenS5{F0: 0, F1: 42}
                        _ = v56
                        _ = v57
                    }
                    _ = fi54
                    _ = v55
                }
                _ = fi52
                _ = v53
            }
        case 128:
            for fi58 := 0; fi58 < 3; fi58++ {
                _ = len(v4)
                v3 -= (^(v5))
                _ = fi58
            }
        }
    }
    var v59 uint8 = 98
    var v60 uint16 = (v5 >> ((42) & 63))
    var v61 int8 = 100
    switch 1 {
    case -2147483648:
        v60 = v3
    case -234:
        var v62 uint64 = 42
        _ = v62
    case -1:
        var v63 uint16 = v5
        _ = v63
    default:
        v16 -= v16
    }
    v64 := GenS2{F0: (^(v17)), F1: true, F2: v61}
    fmt.Println(v5)
    v65 := GenS10{F0: v2}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v60
    _ = v61
    _ = v64
    _ = v65
}
