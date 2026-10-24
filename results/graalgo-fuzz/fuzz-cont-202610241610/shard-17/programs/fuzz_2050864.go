package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 uint32; F1 int; F2 uint16 }
type GenS4 struct { F0 string; F1 int16 }
type GenS5 struct { F0 string; F1 uint64; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint; F2 uint8 }
type GenS7 struct { F0 int32; F1 int32; F2 string }
type GenS8 struct { F0 int64; F1 int32; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint8; F2 int8 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 uint64 }
func genFn0() uint16 {
    v6 := GenS1{F0: 2, F1: 2147483647}
    var v7 int = 1
    v7 -= v7
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16, p1 bool, p2 bool) uint32 {
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint64 = 948
        _ = fi8
        _ = v9
    }
    var v10 int32 = 1
    v11 := GenS1{F0: 32767, F1: 580}
    var v12 float32 = 3.14159
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 903
}

func main() {
    var v1 uint8 = 109
    v2 := make(map[int]string)
    var v3 uint8 = v1
    var v4 int = 4294967295
    var v5 int = v4
    v13 := GenS2{F0: (genFn1(-32768, false, (-2.25 < ((0.0 - -2.25) * 3.14159))) >= 732), F1: (true || false)}
    switch (-(v4)) {
    case 32767:
        if true {
            v14 := GenS3{F0: 870, F1: 127, F2: 0}
            _ = v14
        }
    case 127:
        v15 := []int{v4}
        _ = v15
    case 4294967295:
        fmt.Println(v5)
    default:
        fmt.Println(v4, v3)
    }
    v16 := v2[42]
    for fi17 := 0; fi17 < 7; fi17++ {
        if false {
            for fi18 := 0; fi18 < 9; fi18++ {
                v16 = "foo bar"
                switch fi17 {
                case 1:
                    v4 += fi18
                case 128:
                    v19 := []int{(v4 / (((v5 % ((fi18) | 1))) | 1)), (fi17 / ((v5) | 1))}
                    _ = v19
                default:
                    if (genFn0() <= 857) {
                        if true {
                            v20 := GenS4{F0: "\n\t", F1: 32767}
                            _ = v20
                        } else {
                            v1 -= v3
                            if false {
                                v21 := GenS5{F0: "a", F1: 18446744073709551615, F2: 2}
                                v22 := GenS2{F0: true, F1: (-527 > -1)}
                                _ = v21
                                _ = v22
                            } else {
                                v4 = (-(fi18))
                            }
                            if (v16 < v16) {
                                var v23 int8 = 0
                                var v24 uint8 = (v3 << ((42) & 63))
                                _ = v23
                                _ = v24
                            } else {
                                var v25 int = 32767
                                fmt.Println(v25)
                                v4 = int(v1)
                                if false {
                                    v26 := GenS6{F0: -1, F1: 65535, F2: (^(v3))}
                                    v1 += (v3 - v1)
                                    _ = v26
                                } else {
                                    switch v16 {
                                    case "a":
                                        if (-1 <= 42) {
                                            var v27 uint8 = (v3 & (^(v1)))
                                            _ = v27
                                        } else {
                                            delete(v2, -9223372036854775808)
                                        }
                                    case "hello":
                                        switch -9223372036854775808 {
                                        case 9223372036854775807:
                                            v4 = (^(v25))
                                        case -9223372036854775808:
                                            v28 := GenS7{F0: 1, F1: 2147483647, F2: ""}
                                            _ = v28
                                        case 4294967295:
                                            var v29 int64 = 0
                                            _ = v29
                                        }
                                        var v30 float64 = 3.14159
                                        if ((702 == 1) || false) {
                                            var v31 uint = 475
                                            switch v4 {
                                            case -368:
                                                v16 = v16
                                            case 127:
                                                if (1.5 != float32(v30)) {
                                                    fmt.Println(v4, v16, v31)
                                                    fmt.Println(v30, fi17, v1)
                                                }
                                                var v32 uint16 = genFn0()
                                                v32 = uint16(v31)
                                                v33 := GenS8{F0: 9223372036854775807, F1: -2147483648, F2: 1}
                                                _ = v32
                                                _ = v33
                                            case 4294967295:
                                                v5 = v5
                                            default:
                                                if false {
                                                    v1 = v3
                                                } else {
                                                    fmt.Println(fi17)
                                                    v34 := GenS8{F0: 1, F1: 2147483647, F2: -1}
                                                    fmt.Println(v4)
                                                    v4 -= fi17
                                                    var v35 bool = (!(570 != 4294967295))
                                                    _ = v34
                                                    _ = v35
                                                }
                                            }
                                            v36 := []int{fi17, (fi18 & (fi17 >> (((v31 | v31)) & 63)))}
                                            var v37 uint = 255
                                            _ = v31
                                            _ = v36
                                            _ = v37
                                        }
                                        _ = v30
                                    case "foo bar":
                                        v38 := GenS9{F0: 255, F1: (v3 * 120), F2: 45}
                                        _ = v38
                                    default:
                                        v39 := []int{(-(v4))}
                                        _ = v39
                                    }
                                }
                                v4 = v5
                                _ = v25
                            }
                            v40 := []int{(v5 & (v4 ^ (^(fi18))))}
                            _ = v40
                        }
                    } else {
                        v5 += int(v3)
                    }
                }
                _ = fi18
            }
            if true {
                var v41 uint64 = 2
                _ = v41
            } else {
                v4 -= (v4 | (v5 ^ (v5 | v5)))
                switch (v4 & 127) {
                case 1:
                    v42 := GenS10{F0: genFn1(329, false, ((32767 != -694) || (!true)))}
                    _ = v42
                case 9223372036854775807:
                    if (true && false) {
                        v43 := GenS3{F0: 65535, F1: int(v3), F2: 1}
                        if (true || (true && (!false))) {
                            for fi44 := 0; fi44 < 7; fi44++ {
                                v45 := v2[42]
                                var v46 uint16 = 65535
                                v3 = v3
                                _ = fi44
                                _ = v45
                                _ = v46
                            }
                            var v47 int64 = 42
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 12; fi48++ {
                            if false {
                                v49 := []int{}
                                var v50 int64 = 1
                                _ = v49
                                _ = v50
                            } else {
                                if false {
                                    if false {
                                        var v51 string = v16
                                        _ = v51
                                    }
                                } else {
                                    if true {
                                        var v52 int = (fi48 / ((fi48) | 1))
                                        var v53 int = (-(v52))
                                        var v54 uint16 = 61
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                    v5 = v4
                                }
                                if true {
                                    v4 -= fi17
                                    var v55 uint64 = 0
                                    v3 += (v3 / ((v1) | 1))
                                    _ = v55
                                } else {
                                    var v56 uint16 = 255
                                    var v57 uint64 = 42
                                    v56 -= 65535
                                    if (2147483647 > 1) {
                                        var v58 bool = (9223372036854775807 == 9223372036854775807)
                                        v59 := v2[-657]
                                        v60 := GenS8{F0: 160, F1: 2147483647, F2: 100}
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    } else {
                                        v16 = v16
                                    }
                                    _ = v56
                                    _ = v57
                                }
                            }
                            var v61 int = 127
                            _ = fi48
                            _ = v61
                        }
                        _ = v43
                    }
                case -9223372036854775808:
                    var v62 uint32 = 590
                    _ = v62
                }
                switch (("a" + (v16 + ("\n\t" + (v16 + v16)))) + v16) {
                case "foo bar":
                    switch (("\n\t" + ((v16 + "") + v16)) + (((v16 + v16) + ((("hello" + v16) + v16) + "")) + "foo bar")) {
                    case "":
                        fmt.Println(v16)
                    case "a":
                        if true {
                            switch v16 {
                            case "hello":
                                v16 = v16
                            default:
                                v63 := []int{v5, (-(fi17))}
                                _ = v63
                            }
                        }
                    default:
                        if true {
                            _ = len(v2)
                            v5 += (v5 ^ 128)
                        } else {
                            v4 = fi17
                            v64 := []int{}
                            var v65 uint8 = v3
                            _ = v64
                            _ = v65
                        }
                    }
                case "hello":
                    var v66 uint32 = 2147483647
                    _ = v66
                case "":
                    var v67 uint16 = 945
                    _ = v67
                }
            }
            v68 := GenS8{F0: 0, F1: 2147483647, F2: 1}
            _ = v68
        }
        for fi69 := 0; fi69 < 7; fi69++ {
            v70 := GenS11{F0: 0}
            var v71 uint8 = v1
            _ = fi69
            _ = v70
            _ = v71
        }
        fmt.Println(v1)
        _ = fi17
    }
    v3 = 1
    if (true && (0 <= 1)) {
        for fi72 := 0; fi72 < 6; fi72++ {
            fmt.Println(v3)
            _ = fi72
        }
        switch "foo bar" {
        case "":
            v3 -= v1
        case "a":
            for fi73 := 0; fi73 < 4; fi73++ {
                v4 = (fi73 >> ((2) & 63))
                v5 -= v5
                _ = fi73
            }
        }
    }
    v74 := GenS5{F0: ("" + (((("\n\t" + v16) + "hello") + "") + v16)), F1: 18446744073709551615, F2: 0}
    for fi75 := 0; fi75 < 4; fi75++ {
        for fi76 := 0; fi76 < 11; fi76++ {
            var v77 int32 = -2147483648
            _ = fi76
            _ = v77
        }
        v78 := GenS4{F0: v16, F1: 1}
        v4 += v5
        _ = fi75
        _ = v78
    }
    switch (v5 % (((v4 | 42)) | 1)) {
    case -816:
        var v79 float64 = 1.5
        _ = v79
    case -2147483648:
        var v80 string = "foo bar"
        _ = v80
    case 9223372036854775807:
        v81 := v2[32768]
        _ = v81
    default:
        if (42 <= 2147483647) {
            var v82 uint64 = 2
            _ = v82
        } else {
            var v83 uint8 = (^(v3))
            v84 := GenS4{F0: v16, F1: 1000}
            _ = v83
            _ = v84
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v74
}
