package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 int; F2 uint8 }
type GenS5 struct { F0 int32; F1 int8; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint16; F2 uint }
type GenS8 struct { F0 string; F1 uint64; F2 uint64 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int; F1 int64 }
type GenS11 struct { F0 int; F1 int32; F2 int64 }
func genFn0(p0 int64, p1 int8, p2 uint16) string {
    var v7 string = "\n\t"
    v7 = ("foo bar" + (("" + "\n\t") + v7))
    v7 = v7
    v8 := GenS1{F0: 127}
    v7 = v7
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "hello"
}
func genFn1(p0 int8) int {
    if false {
        var v9 uint64 = 18446744073709551615
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint = 4294967295
            if (true || true) {
                v12 := []int{}
                v9 = (v9 & (v9 - v9))
                _ = v12
            } else {
                var v13 uint = uint(v9)
                fmt.Println(v9)
                v14 := []int{(fi10 + -2147483648), (fi10 + (-(fi10)))}
                _ = v13
                _ = v14
            }
            v15 := GenS2{F0: ((v9 > (v9 % ((4294967295) | 1))) || false), F1: (-1 >= 42)}
            _ = fi10
            _ = v11
            _ = v15
        }
        _ = v9
    } else {
        var v16 int = 872
        for fi17 := 0; fi17 < 10; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                v19 := GenS1{F0: (fi17 * fi18)}
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            switch "foo bar" {
            case "hello":
                for fi21 := 0; fi21 < 20; fi21++ {
                    p0 += p0
                    if (2 > 5) {
                        v22 := GenS3{F0: (0 >= 734)}
                        _ = v22
                    } else {
                        if (0 < 661) {
                            v16 = v16
                            if false {
                                v23 := []int{(-(fi20))}
                                _ = v23
                            } else {
                                var v24 int8 = p0
                                v25 := []int{}
                                _ = v24
                                _ = v25
                            }
                        } else {
                            v26 := []int{(-(fi21)), fi21, (v16 - fi21), -9223372036854775808}
                            _ = v26
                        }
                        v27 := GenS1{F0: 540}
                        fmt.Println(fi21, fi21)
                        _ = v27
                    }
                    var v28 string = ("foo bar" + "\n\t")
                    _ = fi21
                    _ = v28
                }
            default:
                var v29 uint32 = 1
                _ = v29
            }
            if (!true) {
                var v30 float64 = 0.1
                v31 := GenS1{F0: (fi20 >> ((42) & 63))}
                var v32 int32 = 2147483647
                _ = v30
                _ = v31
                _ = v32
            }
            v33 := []int{}
            _ = fi20
            _ = v33
        }
        _ = v16
    }
    switch 32767 {
    case 32767:
        var v34 uint16 = 284
        _ = v34
    case 751:
        if false {
            var v35 float32 = 0.0
            fmt.Println(p0, v35)
            v36 := GenS3{F0: true}
            _ = v35
            _ = v36
        }
    case 1:
        var v37 uint64 = 18446744073709551615
        _ = v37
    default:
        v38 := []int{32767, 32767, 597}
        _ = v38
    }
    v39 := GenS1{F0: 1}
    v40 := GenS4{F0: 127, F1: 127, F2: 127}
    fmt.Println(p0)
    _ = p0
    _ = v39
    _ = v40
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    var v3 int = -1
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 int16 = 32767
    switch "" {
    case "a":
        switch "\n\t" {
        case "foo bar":
            if true {
                var v41 uint32 = 42
                _ = v41
            } else {
                fmt.Println(v6, v3)
                v42 := GenS1{F0: v3}
                var v43 int16 = (v6 ^ v6)
                _ = v42
                _ = v43
            }
        case "":
            v1[65535] = 4294967295
        case "a":
            if (543 < 0) {
                if (!(!false)) {
                    var v44 int = (v3 & (v3 % (((v3 - v3)) | 1)))
                    v45 := []int{}
                    v6 = v6
                    _ = v44
                    _ = v45
                }
            }
        }
    }
    v46 := GenS5{F0: -286, F1: 1, F2: 42}
    v3 -= (-(v3))
    v3 = v3
    for fi47 := 0; fi47 < 18; fi47++ {
        var v48 int = -9223372036854775808
        switch "" {
        case "\n\t":
            v1[65535] = v3
        }
        v49 := GenS6{F0: 127}
        _ = fi47
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 4; fi50++ {
        if true {
            var v51 string = "foo bar"
            _ = v51
        }
        if false {
            v6 = (v6 - int16(v3))
        }
        var v52 string = ("foo bar" + "foo bar")
        _ = fi50
        _ = v52
    }
    if true {
        fmt.Println(v3, v6)
        delete(v5, 4294967295)
        for fi53 := 0; fi53 < 2; fi53++ {
            var v54 int = genFn1(106)
            for fi55 := 0; fi55 < 6; fi55++ {
                if true {
                    _ = len(v2)
                    delete(v1, 2147483647)
                } else {
                    v56 := GenS3{F0: false}
                    _ = v56
                }
                v57 := GenS3{F0: true}
                fmt.Println(fi53, fi53)
                _ = fi55
                _ = v57
            }
            v58 := GenS7{F0: 163, F1: 673, F2: 42}
            _ = fi53
            _ = v54
            _ = v58
        }
    } else {
        if true {
            fmt.Println(v6)
            if true {
                v59 := GenS8{F0: (("a" + ("foo bar" + "hello")) + ("foo bar" + "")), F1: 330, F2: 2}
                _ = v59
            }
        } else {
            switch v3 {
            case 2147483647:
                switch 2147483647 {
                case 32767:
                    fmt.Println(v3)
                }
            case 455:
                for fi60 := 0; fi60 < 11; fi60++ {
                    v3 = genFn1(68)
                    var v61 int16 = v6
                    fmt.Println(v61)
                    _ = fi60
                    _ = v61
                }
            case -749:
                v62 := GenS4{F0: 127, F1: 9223372036854775807, F2: 127}
                _ = v62
            default:
                for fi63 := 0; fi63 < 20; fi63++ {
                    switch -539 {
                    case -2147483648:
                        v6 += (^(v6))
                    case -217:
                        v64 := GenS9{F0: true}
                        _ = v64
                    }
                    _ = fi63
                }
            }
            var v65 int8 = 0
            var v66 uint16 = 0
            _ = v65
            _ = v66
        }
        if (0 < v6) {
            v3 = v3
            v67 := []int{}
            _ = v67
        } else {
            var v68 int = v3
            _ = v68
        }
    }
    if false {
        if true {
            v3 -= (^(v3))
            for fi69 := 0; fi69 < 10; fi69++ {
                var v70 float32 = -2.25
                var v71 int32 = 42
                var v72 int64 = 0
                _ = fi69
                _ = v70
                _ = v71
                _ = v72
            }
        } else {
            var v73 int16 = (-(v6))
            switch genFn1(1) {
            case 473:
                fmt.Println(v73, v6)
            case 42:
                v74 := GenS8{F0: "", F1: 4294967295, F2: 2}
                _ = v74
            case 32768:
                if (2147483647 == -50) {
                    switch v3 {
                    case 4294967295:
                        var v75 string = (("hello" + "") + "\n\t")
                        _ = v75
                    case -1:
                        var v76 bool = true
                        _ = v76
                    }
                } else {
                    for fi77 := 0; fi77 < 2; fi77++ {
                        var v78 uint16 = 2
                        for fi79 := 0; fi79 < 13; fi79++ {
                            switch genFn1(-1) {
                            case 2147483647:
                                switch "foo bar" {
                                case "":
                                    fmt.Println(v6, v6, v6)
                                }
                                v6 = v73
                            case 4294967295:
                                fmt.Println(fi79)
                            }
                            _ = fi79
                        }
                        _ = fi77
                        _ = v78
                    }
                    v73 += v6
                }
            default:
                v80 := []int{v3, (v3 % ((0) | 1))}
                _ = v80
            }
            _ = v73
        }
        var v81 string = ("\n\t" + "hello")
        if ((-1 >= 1) || (!true)) {
            if false {
                v82 := GenS4{F0: 205, F1: v3, F2: 225}
                if true {
                    if (false || true) {
                        var v83 uint8 = 110
                        _ = v83
                    }
                }
                v3 = 0
                _ = v82
            } else {
                v84 := GenS8{F0: v81, F1: 0, F2: 1}
                delete(v4, v3)
                fmt.Println(v3, v3)
                _ = v84
            }
        }
        _ = v81
    } else {
        v6 = (v6 & v6)
    }
    switch (v3 >> ((255) & 63)) {
    case -9223372036854775808:
        if (!false) {
            if true {
                var v85 uint8 = 127
                v85 -= v85
                if (255 >= 2) {
                    if (((942 != 35) && (993 > v3)) || false) {
                        v6 -= 32767
                    } else {
                        fmt.Println(v85, v3)
                        var v86 int32 = 2147483647
                        v6 = (v6 / (((-(v6))) | 1))
                        _ = v86
                    }
                    v85 -= (v85 / (((v85 >> ((65535) & 63))) | 1))
                }
                _ = v85
            } else {
                for fi87 := 0; fi87 < 3; fi87++ {
                    v6 = v6
                    _ = fi87
                }
                var v88 string = ("\n\t" + (("foo bar" + ("foo bar" + "foo bar")) + "a"))
                for fi89 := 0; fi89 < 13; fi89++ {
                    switch (v88 + v88) {
                    case "foo bar":
                        for fi90 := 0; fi90 < 10; fi90++ {
                            v91 := []int{v3, (fi89 & genFn1(127)), (^(fi89))}
                            _ = fi90
                            _ = v91
                        }
                    case "":
                        if (2 >= 42) {
                            switch -1 {
                            case -9223372036854775808:
                                fmt.Println(v6)
                            case 9223372036854775807:
                                v6 = v6
                            default:
                                fmt.Println(v6, v6, v88)
                            }
                        } else {
                            if (true && (0 == 263)) {
                                v6 = 32767
                            }
                            for fi92 := 0; fi92 < 19; fi92++ {
                                if true {
                                    v93 := GenS4{F0: 35, F1: 1, F2: 2}
                                    _ = v93
                                } else {
                                    if (1 != -534) {
                                        if true {
                                            v94 := GenS9{F0: false}
                                            switch (v3 - fi89) {
                                            case 197:
                                                v3 = 834
                                            }
                                            _ = v94
                                        } else {
                                            fmt.Println(v3, fi92, v3)
                                            fmt.Println(fi92)
                                        }
                                        v6 = v6
                                        if (186 != 127) {
                                            var v95 int = fi89
                                            v96 := GenS10{F0: fi89, F1: 0}
                                            _ = v95
                                            _ = v96
                                        } else {
                                            v3 += (^(v3))
                                            v88 = ("hello" + "")
                                            var v97 float64 = ((0.0 + 3.14159) / (1.5 - 0.0))
                                            var v98 uint64 = 152
                                            v99 := []int{}
                                            _ = v97
                                            _ = v98
                                            _ = v99
                                        }
                                    }
                                    v100 := GenS5{F0: 753, F1: 127, F2: 1}
                                    var v101 float64 = 1e10
                                    _ = v100
                                    _ = v101
                                }
                                if (true || false) {
                                    v88 = "foo bar"
                                    v3 = (fi89 ^ (fi92 / ((fi89) | 1)))
                                    var v102 int32 = -2147483648
                                    _ = v102
                                } else {
                                    v103 := GenS11{F0: (-(fi92)), F1: 2147483647, F2: 0}
                                    _ = v103
                                }
                                switch (("foo bar" + v88) + v88) {
                                case "\n\t":
                                    fmt.Println(v88, v6, v6)
                                default:
                                    var v104 uint8 = 44
                                    _ = v104
                                }
                                _ = fi92
                            }
                        }
                    }
                    v3 -= fi89
                    _ = fi89
                }
                _ = v88
            }
            delete(v2, genFn1(79))
            fmt.Println(v3, v6)
        } else {
            fmt.Println(v6, v3)
            var v105 int16 = v6
            _ = v105
        }
    case 127:
        switch (v3 * (v3 & (^(v3)))) {
        case 128:
            v6 -= (v6 | v6)
        case 1:
            if (1e10 <= 0.0) {
                var v106 int64 = -9223372036854775808
                var v107 int64 = v106
                _ = v106
                _ = v107
            }
        }
    case -1:
        v4[(-(v3))] = (false || (v6 > -1))
    }
    switch v3 {
    case 1:
        var v108 uint = 65535
        _ = v108
    case 127:
        v3 -= (^(v3))
    case 4294967295:
        v6 = (v6 - v6)
    }
    fmt.Println(v3, v3)
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 << ((2) & 63))
    case "foo bar":
        var v109 uint32 = 2147483648
        _ = v109
    default:
        for fi110 := 0; fi110 < 4; fi110++ {
            v111 := GenS9{F0: ((255 >= 0) || ((v6 ^ v6) != (v6 | v6)))}
            _ = fi110
            _ = v111
        }
    }
    fmt.Println(v6, v3)
    v112 := []int{v3, v3}
    v113 := GenS4{F0: 127, F1: 42, F2: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v112:", len(v112), v112)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v112
    _ = v113
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
