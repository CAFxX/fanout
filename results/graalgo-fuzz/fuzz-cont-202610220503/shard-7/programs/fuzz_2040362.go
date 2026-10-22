package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint }
type GenS2 struct { F0 int16; F1 int; F2 uint8 }
type GenS3 struct { F0 uint32; F1 int32; F2 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 int64 }
type GenS6 struct { F0 uint64; F1 uint8; F2 string }
type GenS7 struct { F0 int; F1 int8; F2 int32 }
type GenS8 struct { F0 bool; F1 int16; F2 string }
type GenS9 struct { F0 uint16; F1 uint32 }
type GenS10 struct { F0 uint32; F1 string }
type GenS11 struct { F0 int16; F1 int; F2 uint16 }
func genFn0(p0 int16, p1 int16) int32 {
    var v7 string = "a"
    v8 := GenS1{F0: -1, F1: 42, F2: 255}
    p1 = (p1 | (p0 + p1))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 int8, p1 int64) bool {
    var v9 uint8 = 42
    p0 -= 42
    v10 := []int{int(p0)}
    v11 := GenS1{F0: (p1 + 0), F1: p1, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (!false)
}

func main() {
    var v1 uint = 1
    var v2 uint = 0
    v3 := []string{"", "\n\t", "a"}
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v6 int32 = -1
    v2 = (^(v2))
    if true {
        for fi12 := 0; fi12 < 13; fi12++ {
            for fi13 := 0; fi13 < 12; fi13++ {
                if (241 < 20) {
                    v2 = v2
                    v2 = v1
                    var v14 int = (-(fi12))
                    _ = v14
                } else {
                    var v15 bool = (!(v6 == v6))
                    var v16 bool = (4294967295 == 42)
                    _ = v15
                    _ = v16
                }
                var v17 bool = true
                v18 := GenS1{F0: -952, F1: 9223372036854775807, F2: (^(v1))}
                fmt.Println(v1, v1)
                _ = fi13
                _ = v17
                _ = v18
            }
            var v19 int32 = (v6 >> ((1) & 63))
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    switch (fi20 + (^(fi12))) {
                    case 879:
                        fmt.Println(v1)
                    case 0:
                        v21 := GenS1{F0: -9223372036854775808, F1: -1, F2: v2}
                        _ = v21
                    default:
                        if (("foo bar" + "a") == "hello") {
                            var v22 uint64 = 18446744073709551615
                            _ = v22
                        } else {
                            switch "\n\t" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = "hello"
                                }
                            case "foo bar":
                                v2 = (v2 / (((v2 & (v2 + 4294967295))) | 1))
                            }
                            v23 := GenS2{F0: 42, F1: fi20, F2: 42}
                            v24 := v4[42]
                            v1 -= (^(v2))
                            v25 := GenS2{F0: -281, F1: v24, F2: 51}
                            v1 = v1
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    }
                    _ = fi20
                }
            } else {
                v2 = v2
                v26 := GenS3{F0: 2, F1: v19, F2: -1}
                v6 += (v19 >> (((^(v2))) & 63))
                _ = v26
            }
            _ = fi12
            _ = v19
        }
        _ = len(v4)
    } else {
        v27 := GenS4{F0: 128}
        var v28 uint64 = 728
        v29 := GenS5{F0: -112, F1: 9223372036854775807}
        _ = v27
        _ = v28
        _ = v29
    }
    v1 = (v2 + 337)
    if (1.5 <= (1.5 / float64(1.5))) {
        if ((false || (v6 == 42)) && (!false)) {
            switch "foo bar" {
            case "":
                v2 = v1
            case "hello":
                v30 := GenS6{F0: 42, F1: 0, F2: "\n\t"}
                _ = v30
            default:
                v31 := GenS6{F0: 1, F1: 11, F2: ""}
                _ = v31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                if (!false) {
                    var v33 float32 = 1.5
                    if (-23 == 9223372036854775807) {
                        v33 += -2.25
                        for fi34 := 0; fi34 < 10; fi34++ {
                            v35 := GenS2{F0: -32768, F1: fi32, F2: 169}
                            _ = fi34
                            _ = v35
                        }
                        for fi36 := 0; fi36 < 5; fi36++ {
                            var v37 int32 = (v6 / ((v6) | 1))
                            v38 := v5[fi36]
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    } else {
                        v39 := GenS7{F0: (fi32 - fi32), F1: 0, F2: (v6 / (((v6 >> (((v1 ^ (v1 & 47))) & 63))) | 1))}
                        v6 = v6
                        _ = v39
                    }
                    fmt.Println(v2, v6)
                    _ = v33
                }
                if true {
                    if (true && (1 <= 42)) {
                        v2 += 322
                    }
                    if false {
                        if true {
                            v40 := GenS8{F0: true, F1: 1, F2: (("\n\t" + "\n\t") + (("hello" + (("foo bar" + (("a" + "a") + "\n\t")) + "hello")) + (("hello" + "") + ("" + "\n\t"))))}
                            var v41 bool = ((false || ((0 == 18446744073709551615) && (!false))) && (!false))
                            _ = v40
                            _ = v41
                        } else {
                            for fi42 := 0; fi42 < 11; fi42++ {
                                v43 := GenS4{F0: (fi42 & fi42)}
                                var v44 uint16 = 42
                                var v45 uint = v1
                                if true {
                                    fmt.Println(v2, fi42, fi42)
                                }
                                if (((true || (-1 >= 1)) && true) || false) {
                                    var v46 int32 = v6
                                    _ = v46
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                _ = fi42
                                _ = v43
                                _ = v44
                                _ = v45
                            }
                            for fi47 := 0; fi47 < 15; fi47++ {
                                if true {
                                    fmt.Println(fi32, v6)
                                    v2 = v1
                                } else {
                                    var v48 uint8 = 242
                                    if false {
                                        var v49 uint64 = 0
                                        v48 = (^(v48))
                                        _ = v49
                                    } else {
                                        var v50 int = fi47
                                        _ = v50
                                    }
                                    v2 -= (v2 << (((^(v1))) & 63))
                                    _ = v48
                                }
                                v1 -= (v1 * v2)
                                _ = fi47
                            }
                        }
                    }
                    if ((v6 - genFn0(-910, 42)) == (-(v6))) {
                        for fi51 := 0; fi51 < 10; fi51++ {
                            v52 := GenS8{F0: (!false), F1: 42, F2: (("hello" + "hello") + (("foo bar" + "foo bar") + "\n\t"))}
                            v1 = 137
                            var v53 float64 = -2.25
                            _ = fi51
                            _ = v52
                            _ = v53
                        }
                        fmt.Println(fi32, v6, v1)
                        v1 += (v1 / (((^(v2))) | 1))
                    } else {
                        var v54 int = 128
                        var v55 int64 = 42
                        v55 = (v55 << ((v1) & 63))
                        _ = v54
                        _ = v55
                    }
                } else {
                    v3 = append(v3, "")
                    v1 = (v1 - 65535)
                }
                _ = fi32
            }
        } else {
            for fi56 := 0; fi56 < 14; fi56++ {
                fmt.Println(fi56, v1)
                if (!false) {
                    switch (("hello" + "\n\t") + (("hello" + "foo bar") + "\n\t")) {
                    case "a":
                        v2 -= (v1 | (v2 | v1))
                    case "hello":
                        var v57 uint16 = 518
                        _ = v57
                    default:
                        v58 := GenS2{F0: 0, F1: 759, F2: 0}
                        _ = v58
                    }
                } else {
                    v6 = v6
                    switch (("hello" + "hello") + ("hello" + "a")) {
                    case "a":
                        if false {
                            v1 -= (v2 / ((v1) | 1))
                        }
                    case "":
                        fmt.Println(v2, v2)
                    case "hello":
                        for fi59 := 0; fi59 < 10; fi59++ {
                            var v60 int = (fi59 | 42)
                            _ = fi59
                            _ = v60
                        }
                    default:
                        v6 = v6
                    }
                }
                switch "\n\t" {
                case "foo bar":
                    var v61 uint = 4294967295
                    _ = v61
                case "":
                    v2 = v1
                default:
                    v1 -= 0
                }
                _ = fi56
            }
            v2 += v2
            _ = len(v5)
        }
    } else {
        if true {
            var v62 bool = true
            _ = v62
        }
    }
    switch 2147483647 {
    case -1:
        v63 := GenS6{F0: 4294967295, F1: 127, F2: "foo bar"}
        _ = v63
    case -692:
        v64 := v4[1]
        _ = v64
    }
    v5[2147483647] = ("a" + "a")
    var v65 uint16 = 649
    switch 4294967295 {
    case 2147483647:
        fmt.Println(v6)
    }
    var v66 uint64 = 2
    switch (("\n\t" + ("\n\t" + "")) + "foo bar") {
    case "":
        switch "foo bar" {
        case "hello":
            if (false || true) {
                v66 = v66
                var v67 uint32 = 2
                _ = v67
            } else {
                _ = len(v3)
                for fi68 := 0; fi68 < 6; fi68++ {
                    var v69 int = (fi68 ^ fi68)
                    for fi70 := 0; fi70 < 16; fi70++ {
                        var v71 uint8 = 42
                        _ = fi70
                        _ = v71
                    }
                    v72 := GenS9{F0: v65, F1: 4294967295}
                    _ = fi68
                    _ = v69
                    _ = v72
                }
            }
        case "":
            var v73 float64 = 0.1
            _ = v73
        }
    }
    switch 431 {
    case 1:
        v74 := GenS10{F0: 522, F1: "hello"}
        _ = v74
    }
    for fi75 := 0; fi75 < 15; fi75++ {
        if true {
            fmt.Println(v65, v2, v66)
        } else {
            var v76 int = fi75
            var v77 uint64 = 400
            for fi78 := 0; fi78 < 4; fi78++ {
                var v79 int16 = 1000
                v80 := v4[519]
                var v81 int = v76
                fmt.Println(v66, v6, v77)
                _ = fi78
                _ = v79
                _ = v80
                _ = v81
            }
            _ = v76
            _ = v77
        }
        var v82 int32 = -59
        _ = fi75
        _ = v82
    }
    if true {
        v83 := GenS11{F0: 32767, F1: -1, F2: uint16(v66)}
        _ = v83
    }
    v84 := v4[2147483647]
    for fi85 := 0; fi85 < 6; fi85++ {
        var v86 uint16 = v65
        _ = fi85
        _ = v86
    }
    fmt.Println(v84, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v84
}
