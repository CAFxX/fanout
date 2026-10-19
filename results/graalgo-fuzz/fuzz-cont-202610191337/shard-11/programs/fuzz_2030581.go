package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 bool; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int8 }
func genFn0(p0 int64, p1 uint32) int8 {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int64 = (p0 + (^(p0)))
        var v6 int = (fi4 * 9223372036854775807)
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int16 = int16(p0)
    if (!(-2.25 != (0.1 * 0.0))) {
        for fi8 := 0; fi8 < 6; fi8++ {
            if (1e10 > 1e10) {
                fmt.Println(p1, v7)
                v9 := GenS1{F0: 65535}
                _ = v9
            } else {
                var v10 float32 = 1.5
                var v11 uint16 = 0
                fmt.Println(v11, v11, v7)
                _ = v10
                _ = v11
            }
            var v12 uint8 = 42
            if (true || (v7 != (v7 % ((v7) | 1)))) {
                var v13 uint16 = 666
                _ = v13
            } else {
                v14 := GenS1{F0: 32767}
                p0 = 389
                _ = v14
            }
            _ = fi8
            _ = v12
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            if (0.1 > 0.1) {
                var v16 int64 = -9223372036854775808
                if true {
                    v17 := GenS1{F0: 42}
                    var v18 uint16 = 986
                    for fi19 := 0; fi19 < 18; fi19++ {
                        v7 -= (-(v7))
                        var v20 int8 = 127
                        _ = fi19
                        _ = v20
                    }
                    _ = v17
                    _ = v18
                }
                fmt.Println(p0)
                _ = v16
            }
            if true {
                for fi21 := 0; fi21 < 16; fi21++ {
                    if (fi15 <= (fi15 + fi21)) {
                        v22 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615, F2: "foo bar"}
                        v23 := GenS1{F0: 675}
                        switch (fi15 ^ 0) {
                        case 268:
                            if ((!false) || true) {
                                var v24 uint = 51
                                var v25 int = fi21
                                _ = v24
                                _ = v25
                            } else {
                                v26 := []int{(fi21 % ((168) | 1))}
                                _ = v26
                            }
                        }
                        _ = v22
                        _ = v23
                    }
                    v27 := GenS3{F0: (65535 < 32767), F1: uint64(p1), F2: 988}
                    _ = fi21
                    _ = v27
                }
                for fi28 := 0; fi28 < 20; fi28++ {
                    if (!false) {
                        fmt.Println(p1)
                        if false {
                            fmt.Println(p1, p1)
                        } else {
                            v7 = v7
                        }
                        v7 = v7
                    } else {
                        v29 := []int{fi15}
                        _ = v29
                    }
                    var v30 float64 = 0.1
                    _ = fi28
                    _ = v30
                }
                switch "hello" {
                case "hello":
                    for fi31 := 0; fi31 < 14; fi31++ {
                        if true {
                            p1 = (p1 + 2147483647)
                            p1 -= p1
                        }
                        var v32 uint = 0
                        switch (fi31 & -121) {
                        case -21:
                            if true {
                                p0 = (p0 * (p0 % (((p0 + (-(p0)))) | 1)))
                                v33 := []int{fi15, fi31}
                                _ = v33
                            } else {
                                v34 := GenS3{F0: (39 == uint64(v32)), F1: 42, F2: uint64(v32)}
                                _ = v34
                            }
                        default:
                            p1 += p1
                        }
                        _ = fi31
                        _ = v32
                    }
                case "foo bar":
                    if (-1 == 100) {
                        v35 := []int{(fi15 | -492), 128, fi15, (fi15 ^ fi15)}
                        p1 = p1
                        p0 = p0
                        _ = v35
                    }
                case "":
                    if (true || false) {
                        var v36 int8 = -124
                        if true {
                            v7 -= 1
                        }
                        v37 := GenS1{F0: 578}
                        _ = v36
                        _ = v37
                    } else {
                        for fi38 := 0; fi38 < 5; fi38++ {
                            if true {
                                p1 += 1
                                p1 = p1
                                switch fi15 {
                                case 32767:
                                    var v39 uint32 = 1
                                    _ = v39
                                case 128:
                                    p1 -= 0
                                case 42:
                                    var v40 float64 = 0.1
                                    _ = v40
                                default:
                                    v41 := GenS2{F0: 1, F1: 253, F2: ("\n\t" + "foo bar")}
                                    _ = v41
                                }
                            }
                            if (127 >= 128) {
                                fmt.Println(v7)
                                var v42 int32 = -1
                                _ = v42
                            }
                            p1 = (p1 + p1)
                            v7 = (v7 << ((4294967295) & 63))
                            var v43 int16 = v7
                            fmt.Println(fi15, v7, p0)
                            fmt.Println(p0, p1)
                            _ = fi38
                            _ = v43
                        }
                        for fi44 := 0; fi44 < 3; fi44++ {
                            v45 := GenS1{F0: 2}
                            _ = fi44
                            _ = v45
                        }
                    }
                }
            }
            fmt.Println(v7, fi15)
            _ = fi15
        }
    }
    _ = p0
    _ = p1
    _ = v7
    return -115
}

func main() {
    var v1 bool = (706 > 0)
    v2 := make(map[uint32]string)
    v3 := make(map[string]int)
    v46 := GenS4{F0: 199}
    v47 := GenS5{F0: 259}
    v1 = (!v1)
    v3[(("" + "") + ("" + "a"))] = 2147483647
    v1 = v1
    v48 := GenS2{F0: 1, F1: 355, F2: ("hello" + (("a" + ("\n\t" + "foo bar")) + "a"))}
    v49 := []int{-606}
    v50 := GenS1{F0: 1}
    if ((0.1 > 0.1) && v1) {
        if true {
            for fi51 := 0; fi51 < 19; fi51++ {
                for fi52 := 0; fi52 < 16; fi52++ {
                    v53 := GenS6{F0: 32767}
                    var v54 string = ("hello" + "")
                    if len(v49) >= 2 {
                        v49 = v49[1:len(v49)]
                    }
                    _ = fi52
                    _ = v53
                    _ = v54
                }
                v1 = (v1 || ((fi51 + fi51) == (fi51 << ((18446744073709551615) & 63))))
                fmt.Println(v1)
                _ = fi51
            }
            v55 := v2[2147483648]
            _ = v55
        } else {
            var v56 string = "foo bar"
            if (v56 <= v56) {
                v1 = true
                v57 := GenS5{F0: 4294967295}
                v49 = append(v49, 4294967295)
                _ = v57
            } else {
                if (32767 != 747) {
                    if v1 {
                        v56 = v56
                    } else {
                        fmt.Println(v1)
                        var v58 uint = 1
                        _ = v58
                    }
                    for fi59 := 0; fi59 < 18; fi59++ {
                        var v60 int32 = 2147483647
                        _ = fi59
                        _ = v60
                    }
                    fmt.Println(v1)
                }
            }
            if (1 >= -1) {
                switch -587 {
                case 128:
                    switch 1 {
                    case 1:
                        v1 = false
                    case 42:
                        v56 = "\n\t"
                    default:
                        v61 := GenS1{F0: 0}
                        _ = v61
                    }
                case 32768:
                    if (true && v1) {
                        for fi62 := 0; fi62 < 11; fi62++ {
                            var v63 int32 = 42
                            _ = fi62
                            _ = v63
                        }
                    }
                default:
                    var v64 int8 = -1
                    _ = v64
                }
                var v65 int64 = 9223372036854775807
                _ = v65
            }
            _ = v56
        }
        var v66 uint = 979
        for fi67 := 0; fi67 < 14; fi67++ {
            v66 = (v66 * (v66 | (v66 * (v66 | (v66 ^ v66)))))
            var v68 string = ("" + (("hello" + "\n\t") + "foo bar"))
            v1 = ((128 == 2) || v1)
            _ = fi67
            _ = v68
        }
        _ = v66
    } else {
        fmt.Println(v1)
    }
    for fi69 := 0; fi69 < 3; fi69++ {
        var v70 int = fi69
        if len(v49) >= 2 {
            v49 = v49[1:len(v49)]
        }
        v71 := GenS1{F0: 32767}
        _ = fi69
        _ = v70
        _ = v71
    }
    v1 = true
    for fi72 := 0; fi72 < 8; fi72++ {
        var v73 uint16 = 255
        v74 := GenS7{F0: 42, F1: 42, F2: 100}
        v1 = ((true || ((v1 && true) || true)) || true)
        _ = fi72
        _ = v73
        _ = v74
    }
    for fi75 := 0; fi75 < 7; fi75++ {
        fmt.Println(fi75, v1)
        v76 := GenS3{F0: (false || false), F1: 991, F2: 4294967295}
        var v77 bool = v1
        _ = fi75
        _ = v76
        _ = v77
    }
    if (255 != 2) {
        for fi78 := 0; fi78 < 2; fi78++ {
            fmt.Println(v1)
            if v1 {
                if v1 {
                    for fi79 := 0; fi79 < 9; fi79++ {
                        v80 := v2[2147483648]
                        _ = fi79
                        _ = v80
                    }
                    var v81 uint32 = 2
                    if false {
                        var v82 uint8 = 2
                        _ = v82
                    } else {
                        v81 = 2147483648
                        fmt.Println(fi78, fi78)
                    }
                    _ = v81
                }
                if (!v1) {
                    v83 := v3["\n\t"]
                    for fi84 := 0; fi84 < 1; fi84++ {
                        if (964 == 814) {
                            v85 := GenS3{F0: true, F1: 545, F2: 490}
                            _ = v85
                        } else {
                            var v86 uint64 = 18446744073709551615
                            if (65535 == uint16(v86)) {
                                var v87 string = "hello"
                                var v88 int64 = -827
                                _ = v87
                                _ = v88
                            }
                            _ = v86
                        }
                        _ = fi84
                    }
                    var v89 uint = 65535
                    _ = v83
                    _ = v89
                }
                v90 := GenS7{F0: 127, F1: 903, F2: -128}
                _ = v90
            } else {
                fmt.Println(fi78)
                if (255 == 1) {
                    for fi91 := 0; fi91 < 9; fi91++ {
                        var v92 float64 = 1.5
                        if false {
                            switch (^(fi91)) {
                            case -1:
                                v1 = (2 < 4294967295)
                            case 837:
                                if ((^(fi78)) == fi78) {
                                    if len(v49) >= 2 {
                                        v49 = v49[1:len(v49)]
                                    }
                                    v1 = (v92 >= (v92 + (0.1 * 0.0)))
                                    var v93 uint64 = 475
                                    _ = v93
                                }
                            case -890:
                                v1 = v1
                            }
                            v94 := GenS1{F0: 65535}
                            v1 = true
                            _ = v94
                        }
                        v92 = (1e10 - 3.14159)
                        _ = fi91
                        _ = v92
                    }
                    v95 := GenS1{F0: 65535}
                    v96 := v3[("\n\t" + ("" + (("" + "a") + "hello")))]
                    _ = v95
                    _ = v96
                } else {
                    var v97 uint16 = 42
                    _ = v97
                }
            }
            _ = fi78
        }
        if (-1 >= 1) {
            var v98 float32 = float32((1e10 + (float64(-2.25) * 1e10)))
            _ = v98
        }
        var v99 uint64 = 4294967295
        _ = v99
    } else {
        switch ("foo bar" + "") {
        case "":
            _ = len(v2)
        case "\n\t":
            var v100 int16 = -1
            _ = v100
        case "a":
            var v101 uint32 = 1
            _ = v101
        }
    }
    var v102 float32 = float32(1.5)
    var v103 int = -408
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v49:", len(v49), v49)
    _ = v1
    _ = v102
    _ = v103
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
