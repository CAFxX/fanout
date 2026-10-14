package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint8; F1 string; F2 int }
type GenS7 struct { F0 int64; F1 bool }
func genFn0(p0 uint64) int8 {
    v5 := []int{}
    for fi6 := 0; fi6 < 14; fi6++ {
        v5 = append(v5, 418)
        for fi7 := 0; fi7 < 2; fi7++ {
            p0 += 763
            _ = fi7
        }
        _ = fi6
    }
    var v8 int64 = 1
    _ = p0
    _ = v5
    _ = v8
    return 95
}
func genFn1() uint16 {
    var v9 string = ""
    v10 := GenS1{F0: 65535}
    for fi11 := 0; fi11 < 13; fi11++ {
        if (!true) {
            var v12 int = fi11
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 int32 = 736
                _ = fi13
                _ = v14
            }
            _ = v12
        }
        _ = fi11
    }
    var v15 uint32 = 65535
    v16 := []int{}
    _ = v10
    _ = v15
    _ = v16
    _ = v9
    return 242
}

func main() {
    var v1 uint32 = 4294967295
    var v2 bool = ("" >= ("a" + ""))
    v3 := []uint32{(v1 ^ v1), v1, (v1 * v1), v1}
    var v4 int8 = 127
    if ((true && (787 != genFn1())) && (!(65535 >= 255))) {
        var v17 uint64 = 4294967295
        var v18 int8 = v4
        var v19 int64 = 1
        _ = v17
        _ = v18
        _ = v19
    }
    v4 -= (v4 * v4)
    if len(v3) > 0 {
        v3[0] = v1
    }
    v2 = ((!(1 == 255)) || ((v2 && v2) || v2))
    v1 -= (^(v1))
    v20 := GenS1{F0: 32767}
    switch -2147483648 {
    case 128:
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 bool = (true && v2)
            _ = fi21
            _ = v22
        }
    case 127:
        if (-254 < 42) {
            for fi23 := 0; fi23 < 17; fi23++ {
                for fi24 := 0; fi24 < 18; fi24++ {
                    if (-1 >= 1) {
                        v2 = (42 >= 65535)
                    } else {
                        v25 := GenS2{F0: -32768, F1: fi23}
                        _ = v25
                    }
                    v26 := GenS3{F0: -29}
                    _ = fi24
                    _ = v26
                }
                if (false && (("\n\t" > "hello") || v2)) {
                    if ((2 > 65535) || ((!true) || (42 >= int64(fi23)))) {
                        fmt.Println(fi23)
                        v3 = append(v3, (^(v1)))
                    }
                } else {
                    v27 := GenS1{F0: genFn1()}
                    if len(v3) > 0 {
                        v3[0] = v1
                    }
                    _ = v27
                }
                if false {
                    v1 = v1
                    v2 = ((0.1 >= 1.5) || v2)
                }
                _ = fi23
            }
            var v28 int8 = 40
            _ = v28
        }
    case 130:
        var v29 int64 = 1
        _ = v29
    default:
        switch 128 {
        case 806:
            v1 = (v1 >> ((uint(v1)) & 63))
        case 2147483647:
            v1 += v1
        default:
            for fi30 := 0; fi30 < 17; fi30++ {
                v31 := GenS1{F0: 576}
                _ = fi30
                _ = v31
            }
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        if false {
            v33 := GenS1{F0: 255}
            fmt.Println(fi32, v2, v4)
            fmt.Println(v4)
            _ = v33
        }
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 uint32 = (^(v1))
            var v36 uint64 = 18446744073709551615
            fmt.Println(v2, v4)
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi32
    }
    var v37 int = -990
    v4 = 2
    v4 = 42
    if v2 {
        v1 -= (v1 * v1)
        for fi38 := 0; fi38 < 13; fi38++ {
            v39 := GenS1{F0: 65535}
            _ = fi38
            _ = v39
        }
        if v2 {
            fmt.Println(v2, v2)
            var v40 bool = v2
            _ = v40
        } else {
            fmt.Println(v1)
            if (1.5 > 1e10) {
                v37 = (v37 - (v37 & v37))
                fmt.Println(v37)
            } else {
                var v41 int16 = -32768
                for fi42 := 0; fi42 < 20; fi42++ {
                    if v2 {
                        var v43 int64 = -1
                        for fi44 := 0; fi44 < 15; fi44++ {
                            fmt.Println(fi42, fi42)
                            var v45 float64 = 1e10
                            _ = fi44
                            _ = v45
                        }
                        _ = v43
                    }
                    _ = fi42
                }
                for fi46 := 0; fi46 < 17; fi46++ {
                    switch ("a" + "a") {
                    case "hello":
                        v37 = v37
                    default:
                        fmt.Println(fi46, v2)
                    }
                    v2 = v2
                    _ = fi46
                }
                _ = v41
            }
        }
    }
    var v47 bool = (!v2)
    switch "hello" {
    case "foo bar":
        if (695 == 160) {
            v48 := GenS1{F0: 2}
            _ = v48
        }
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        for fi49 := 0; fi49 < 11; fi49++ {
            _ = len(v3)
            switch fi49 {
            case -522:
                switch ("hello" + ("" + "\n\t")) {
                case "hello":
                    if ((!v47) || (127 != 255)) {
                        if len(v3) > 0 {
                            v3[0] = (^(v1))
                        }
                        if v47 {
                            _ = len(v3)
                            var v50 float32 = (float32(1e10) - (1e10 / 1.5))
                            fmt.Println(v50, v2, v37)
                            _ = v50
                        } else {
                            v51 := GenS4{F0: 2147483647, F1: 2}
                            v4 += (v4 + v4)
                            fmt.Println(v4)
                            _ = v51
                        }
                    }
                }
            default:
                var v52 int8 = 0
                _ = v52
            }
            _ = fi49
        }
    default:
        switch v37 {
        case -644:
            for fi53 := 0; fi53 < 9; fi53++ {
                switch fi53 {
                case -38:
                    switch 1 {
                    case 2147483647:
                        if v47 {
                            if true {
                                _ = len(v3)
                                var v54 uint64 = 703
                                _ = v54
                            } else {
                                v1 = 2
                            }
                        } else {
                            if true {
                                if true {
                                    if v47 {
                                        v55 := GenS5{F0: 0}
                                        var v56 bool = v47
                                        if ((163 < 127) || v56) {
                                            if (!(767 == 1)) {
                                                var v57 uint64 = 18446744073709551615
                                                fmt.Println(v56, v4)
                                                _ = v57
                                            }
                                        }
                                        _ = v55
                                        _ = v56
                                    } else {
                                        for fi58 := 0; fi58 < 10; fi58++ {
                                            v1 -= v1
                                            _ = fi58
                                        }
                                        v59 := GenS6{F0: 127, F1: (("hello" + "hello") + "\n\t"), F2: v37}
                                        var v60 int16 = -32768
                                        _ = v59
                                        _ = v60
                                    }
                                    v1 -= v1
                                }
                            } else {
                                v37 = (fi53 + (^(fi53)))
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v61 := GenS4{F0: 1, F1: 255}
                                _ = v61
                            }
                        }
                    case 32768:
                        fmt.Println(v1, fi53, fi53)
                    default:
                        for fi62 := 0; fi62 < 13; fi62++ {
                            fmt.Println(v47, v47)
                            v63 := GenS7{F0: -790, F1: (!v47)}
                            switch 128 {
                            case -2147483648:
                                if len(v3) > 0 {
                                    v3[0] = v1
                                }
                            case 9223372036854775807:
                                var v64 float32 = 3.14159
                                _ = v64
                            case 0:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                v65 := GenS6{F0: 0, F1: ("a" + ""), F2: fi53}
                                v66 := GenS2{F0: 471, F1: fi62}
                                _ = v65
                                _ = v66
                            default:
                                v4 -= (v4 | (v4 * v4))
                            }
                            _ = fi62
                            _ = v63
                        }
                    }
                }
                switch int(v4) {
                case 0:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -1:
                    v4 = v4
                case 9223372036854775807:
                    fmt.Println(v1, v2)
                }
                switch v37 {
                case 0:
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                case -2147483648:
                    var v67 float32 = (1.5 + 1e10)
                    _ = v67
                case 649:
                    v2 = false
                default:
                    if v47 {
                        for fi68 := 0; fi68 < 8; fi68++ {
                            v47 = v47
                            v47 = v2
                            var v69 float32 = 0.1
                            _ = fi68
                            _ = v69
                        }
                        for fi70 := 0; fi70 < 14; fi70++ {
                            var v71 int8 = v4
                            var v72 float64 = ((-2.25 + 1.5) / 0.1)
                            _ = fi70
                            _ = v71
                            _ = v72
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v73 uint16 = 0
                        v74 := GenS3{F0: 1000}
                        _ = v73
                        _ = v74
                    }
                }
                _ = fi53
            }
        }
    }
    var v75 string = "a"
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v47
    _ = v75
}
