package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int; F1 uint16; F2 string }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int16; F1 string }
type GenS10 struct { F0 uint; F1 int8; F2 uint }
type GenS11 struct { F0 uint8; F1 int16 }
type GenS12 struct { F0 int64 }
func genFn0(p0 int16, p1 int, p2 uint16) int8 {
    p2 = (p2 & (p2 % ((42) | 1)))
    v5 := []int{(p1 + p1), (p1 / ((-2147483648) | 1)), (p1 << ((42) & 63))}
    if (uint(p2) != 4) {
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 int32 = -148
            v8 := GenS1{F0: 42, F1: 1, F2: 976}
            if (1.5 <= 3.14159) {
                v9 := GenS2{F0: p2, F1: 42}
                _ = v9
            } else {
                p2 = p2
                if false {
                    var v10 int32 = 0
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 uint8 = 145
                        var v13 float32 = 0.0
                        if (1e10 >= 0.0) {
                            p2 -= (p2 * 32767)
                        } else {
                            v14 := GenS3{F0: (p2 | (p2 ^ (^(p2)))), F1: 42}
                            p2 = (p2 | (p2 ^ p2))
                            if len(v5) > 0 {
                                v5[0] = (p1 + fi6)
                            }
                            _ = v14
                        }
                        v15 := GenS4{F0: 1}
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v15
                    }
                    p1 += (-(p1))
                    _ = v10
                } else {
                    if false {
                        fmt.Println(p1, v7, p2)
                        var v16 bool = (!true)
                        if v16 {
                            v17 := GenS4{F0: 0}
                            v5 = append(v5, p1)
                            _ = v17
                        } else {
                            for fi18 := 0; fi18 < 3; fi18++ {
                                var v19 int = (p1 + (-(fi6)))
                                if (18446744073709551615 >= 4294967295) {
                                    if ((42 < 2) || false) {
                                        p2 = (^(p2))
                                        _ = len(v5)
                                        var v20 uint32 = 2
                                        var v21 float64 = (1.5 - 0.1)
                                        fmt.Println(v19, v20, fi18)
                                        fmt.Println(v16, v21)
                                        _ = v20
                                        _ = v21
                                    }
                                    p2 += p2
                                    if (10 != 0) {
                                        v19 -= (fi18 | (-(fi18)))
                                    } else {
                                        var v22 uint = 1
                                        switch "hello" {
                                        case "\n\t":
                                            if (2147483648 < 42) {
                                                switch "hello" {
                                                case "foo bar":
                                                    if v16 {
                                                        var v23 int8 = -66
                                                        _ = v23
                                                    } else {
                                                        if len(v5) > 0 {
                                                            v5[0] = 42
                                                        }
                                                    }
                                                case "a":
                                                    fmt.Println(p0, v16)
                                                case "":
                                                    v24 := GenS4{F0: 0}
                                                    _ = v24
                                                }
                                            }
                                            var v25 int64 = -9223372036854775808
                                            _ = v25
                                        case "foo bar":
                                            fmt.Println(v7, v22, v16)
                                        case "hello":
                                            if v16 {
                                                v19 -= fi6
                                                v26 := GenS3{F0: p2, F1: -999}
                                                _ = v26
                                            } else {
                                                switch -1 {
                                                case 32767:
                                                    var v27 int8 = -2
                                                    _ = v27
                                                default:
                                                    var v28 int8 = -118
                                                    _ = v28
                                                }
                                                v7 = v7
                                            }
                                        }
                                        _ = v22
                                    }
                                } else {
                                    v29 := GenS5{F0: "\n\t"}
                                    v30 := GenS3{F0: 255, F1: -203}
                                    _ = v29
                                    _ = v30
                                }
                                p0 -= 32767
                                var v31 uint = 255
                                fmt.Println(v16)
                                v7 -= (v7 % (((v7 % (((^(v7))) | 1))) | 1))
                                p2 = (p2 & p2)
                                _ = fi18
                                _ = v19
                                _ = v31
                            }
                            var v32 int32 = (v7 ^ v7)
                            _ = v32
                        }
                        _ = v16
                    }
                }
            }
            _ = fi6
            _ = v7
            _ = v8
        }
        v5 = append(v5, -2147483648)
        var v33 uint8 = 1
        _ = v33
    }
    v34 := GenS2{F0: (^(p2)), F1: 844}
    _ = p0
    _ = p1
    _ = p2
    _ = v34
    _ = v5
    return 100
}
func genFn1(p0 uint16, p1 int, p2 uint16) uint16 {
    v35 := GenS6{F0: 936, F1: p1}
    p2 += (p2 | p0)
    fmt.Println(p2, p2)
    var v36 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v35
    _ = v36
    return p0
}

func main() {
    var v1 int32 = 1
    var v2 uint = 65535
    var v3 uint32 = 2147483647
    v4 := make(map[uint32]string)
    fmt.Println(v1, v2)
    fmt.Println(v2, v2)
    for fi37 := 0; fi37 < 13; fi37++ {
        v38 := GenS3{F0: uint16(v3), F1: -9223372036854775808}
        switch "" {
        case "\n\t":
            for fi39 := 0; fi39 < 5; fi39++ {
                if true {
                    if (0 < 592) {
                        var v40 int64 = 9223372036854775807
                        _ = v40
                    } else {
                        if false {
                            fmt.Println(v3, v2, v2)
                        }
                        var v41 uint16 = 779
                        _ = v41
                    }
                    v42 := []int{fi39, fi37, (fi39 * 648), (fi39 + 4294967295)}
                    _ = v42
                }
                _ = fi39
            }
        case "foo bar":
            var v43 uint64 = 18446744073709551615
            _ = v43
        case "hello":
            fmt.Println(v3)
        }
        _ = fi37
        _ = v38
    }
    v3 -= 729
    var v44 int32 = (v1 << (((v2 / (((v2 * v2)) | 1))) & 63))
    v3 = v3
    v1 += (v1 >> ((v2) & 63))
    v45 := GenS5{F0: ""}
    if (false && (0.0 < 1e10)) {
        fmt.Println(v3, v3, v1)
        if false {
            if false {
                _ = len(v4)
                var v46 uint8 = 2
                _ = v46
            } else {
                for fi47 := 0; fi47 < 13; fi47++ {
                    if false {
                        var v48 uint32 = v3
                        v44 -= (v1 << (((v2 << (((v2 + 0)) & 63))) & 63))
                        _ = v48
                    }
                    _ = fi47
                }
                var v49 string = ""
                v3 = 2147483647
                _ = v49
            }
            if (1e10 != float32(1e10)) {
                v50 := GenS7{F0: 127, F1: 65535, F2: "\n\t"}
                v51 := []int{}
                v52 := GenS8{F0: (true && ("a" > "hello"))}
                _ = v50
                _ = v51
                _ = v52
            } else {
                v2 = (^(v2))
                fmt.Println(v1, v44, v2)
            }
        } else {
            switch 32768 {
            case -669:
                v53 := GenS2{F0: 378, F1: 0}
                _ = v53
            case -1:
                v1 = v44
            }
            var v54 uint32 = 4294967295
            v55 := GenS9{F0: 138, F1: ""}
            _ = v54
            _ = v55
        }
        if false {
            v4[(v3 | v3)] = "foo bar"
            switch 42 {
            case 0:
                if false {
                    v56 := GenS9{F0: 213, F1: ("foo bar" + "a")}
                    for fi57 := 0; fi57 < 3; fi57++ {
                        if (-126 != 100) {
                            fmt.Println(v2)
                            v44 = v44
                        } else {
                            for fi58 := 0; fi58 < 9; fi58++ {
                                if (2147483647 != v3) {
                                    switch (fi57 & fi58) {
                                    case 9223372036854775807:
                                        var v59 uint8 = 90
                                        _ = v59
                                    case 0:
                                        var v60 int8 = -128
                                        _ = v60
                                    case 712:
                                        v3 = v3
                                    }
                                } else {
                                    v61 := GenS6{F0: v3, F1: 140}
                                    if (int8(fi58) <= 42) {
                                        v62 := GenS7{F0: (fi58 + 781), F1: 454, F2: ((("foo bar" + "foo bar") + "a") + "hello")}
                                        fmt.Println(v3, v1)
                                        var v63 uint64 = 622
                                        v64 := []int{(fi57 ^ (^(fi58))), fi58}
                                        _ = v62
                                        _ = v63
                                        _ = v64
                                    } else {
                                        v65 := []int{1, (fi58 + (fi57 << (((v2 & (v2 | 2))) & 63))), -517}
                                        if (v3 > v3) {
                                            v66 := GenS1{F0: 1, F1: v3, F2: 18446744073709551615}
                                            _ = v66
                                        } else {
                                            var v67 int64 = 0
                                            v68 := GenS10{F0: 255, F1: int8(fi57), F2: (v2 ^ 42)}
                                            _ = v67
                                            _ = v68
                                        }
                                        _ = v65
                                    }
                                    _ = v61
                                }
                                _ = fi58
                            }
                        }
                        v69 := GenS6{F0: v3, F1: 678}
                        _ = fi57
                        _ = v69
                    }
                    switch "foo bar" {
                    case "hello":
                        if true {
                            for fi70 := 0; fi70 < 4; fi70++ {
                                v1 = (v44 | v44)
                                v71 := GenS6{F0: v3, F1: -217}
                                _ = fi70
                                _ = v71
                            }
                        } else {
                            var v72 uint8 = 1
                            fmt.Println(v44)
                            _ = v72
                        }
                    case "\n\t":
                        if false {
                            var v73 int32 = -2147483648
                            _ = v73
                        } else {
                            if true {
                                v2 = (^(v2))
                                _ = len(v4)
                                v74 := GenS7{F0: 42, F1: 136, F2: ""}
                                _ = v74
                            } else {
                                for fi75 := 0; fi75 < 17; fi75++ {
                                    v76 := []int{}
                                    for fi77 := 0; fi77 < 7; fi77++ {
                                        var v78 int16 = -1
                                        _ = fi77
                                        _ = v78
                                    }
                                    var v79 bool = (18446744073709551615 <= v2)
                                    _ = fi75
                                    _ = v76
                                    _ = v79
                                }
                                v2 += v2
                                for fi80 := 0; fi80 < 11; fi80++ {
                                    v81 := GenS7{F0: fi80, F1: 32767, F2: ("\n\t" + "foo bar")}
                                    _ = fi80
                                    _ = v81
                                }
                            }
                        }
                    case "a":
                        fmt.Println(v44, v3, v2)
                    default:
                        for fi82 := 0; fi82 < 1; fi82++ {
                            var v83 uint = (^(v2))
                            var v84 float32 = float32(-2.25)
                            _ = fi82
                            _ = v83
                            _ = v84
                        }
                    }
                    _ = v56
                }
            case 32767:
                v85 := GenS6{F0: 984, F1: 43}
                _ = v85
            case -2147483648:
                v86 := GenS11{F0: 255, F1: -551}
                _ = v86
            default:
                if (968 <= (^(v3))) {
                    v2 = v2
                    v2 += (v2 ^ (v2 + (^(v2))))
                } else {
                    v1 = (v1 % ((v44) | 1))
                }
            }
        }
    }
    v44 = v44
    for fi87 := 0; fi87 < 2; fi87++ {
        v44 = v44
        _ = fi87
    }
    if (false && (0 != -1)) {
        _ = len(v4)
    } else {
        for fi88 := 0; fi88 < 3; fi88++ {
            v89 := GenS12{F0: 42}
            _ = fi88
            _ = v89
        }
        v1 += v44
    }
    for fi90 := 0; fi90 < 12; fi90++ {
        var v91 uint32 = v3
        _ = fi90
        _ = v91
    }
    var v92 int8 = 1
    delete(v4, v3)
    v2 += (v2 >> (((^(v2))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v92
}
