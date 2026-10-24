package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 string }
type GenS3 struct { F0 bool; F1 string; F2 int }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int32; F1 uint32; F2 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int64; F1 uint16 }
type GenS9 struct { F0 bool }
func genFn0(p0 uint16, p1 uint) uint8 {
    if (-699 >= 1) {
        p0 = (p0 | uint16(p1))
        v7 := []int{658, 126, -9223372036854775808, -272}
        var v8 int16 = -32768
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: 1}
        var v10 uint8 = 127
        fmt.Println(p1, v10, v10)
        _ = v10
        _ = v9
    }
    p0 = (p0 ^ 2)
    _ = p0
    _ = p1
    return 131
}
func genFn1(p0 int32, p1 string, p2 bool) uint8 {
    p0 += p0
    v11 := GenS2{F0: (p0 * (^(p0))), F1: 42, F2: ""}
    var v12 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 uint = 255
    var v3 int32 = 59
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v6 int64 = 1
    if (1 != 2) {
        v1 += v1
        if false {
            switch "\n\t" {
            case "a":
                for fi13 := 0; fi13 < 6; fi13++ {
                    v6 = v6
                    switch -808 {
                    case 32768:
                        v14 := []int{}
                        _ = v14
                    case 128:
                        v15 := v4[v1]
                        _ = v15
                    }
                    if false {
                        v3 += v3
                        var v16 int32 = (v3 % (((v3 ^ v3)) | 1))
                        var v17 int8 = -30
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                }
            case "foo bar":
                if (!(v2 > v2)) {
                    var v18 float32 = 0.0
                    v19 := GenS2{F0: v3, F1: 1, F2: ("a" + "\n\t")}
                    _ = v18
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 10; fi20++ {
                        v2 = (v2 + v2)
                        _ = len(v5)
                        var v21 uint64 = 1
                        _ = fi20
                        _ = v21
                    }
                }
            }
            v22 := GenS3{F0: false, F1: "", F2: 529}
            v2 -= v2
            _ = v22
        } else {
            var v23 uint = (v2 & (v2 | (v2 / (((v2 + (v2 >> ((v2) & 63)))) | 1))))
            if false {
                if (!(uint64(v1) <= 42)) {
                    for fi24 := 0; fi24 < 6; fi24++ {
                        switch "foo bar" {
                        case "\n\t":
                            v25 := GenS1{F0: -128}
                            _ = v25
                        case "":
                            v23 = v23
                        default:
                            v26 := GenS3{F0: (!true), F1: ((("\n\t" + (("foo bar" + "") + "foo bar")) + "foo bar") + (("foo bar" + ("foo bar" + "hello")) + "")), F2: fi24}
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v2 = 87
                            v28 := GenS2{F0: v3, F1: 42, F2: ("foo bar" + ("a" + (("foo bar" + "\n\t") + "")))}
                            _ = fi27
                            _ = v28
                        }
                        _ = fi24
                    }
                    delete(v4, uint32(v2))
                } else {
                    switch "hello" {
                    case "hello":
                        for fi29 := 0; fi29 < 14; fi29++ {
                            if (!false) {
                                v30 := GenS2{F0: v3, F1: 1, F2: "hello"}
                                _ = v30
                            } else {
                                switch 0 {
                                case 127:
                                    if true {
                                        if (!(!(!true))) {
                                            _ = len(v4)
                                        }
                                    } else {
                                        fmt.Println(v3, v23, v2)
                                    }
                                case 32768:
                                    for fi31 := 0; fi31 < 1; fi31++ {
                                        if ((^(v23)) <= (^(v23))) {
                                            v23 -= v23
                                            v6 -= v6
                                        }
                                        v32 := GenS4{F0: ((fi29 >> (((v2 << (((v2 & v2)) & 63))) & 63)) != (fi29 << ((v23) & 63))), F1: 604}
                                        if (true && (fi31 >= (fi31 << ((v2) & 63)))) {
                                            v33 := GenS1{F0: 42}
                                            var v34 int = (fi31 * (^(fi31)))
                                            var v35 int8 = -128
                                            _ = v33
                                            _ = v34
                                            _ = v35
                                        }
                                        _ = fi31
                                        _ = v32
                                    }
                                case 42:
                                    switch "\n\t" {
                                    case "hello":
                                        fmt.Println(v1, v3)
                                    case "a":
                                        switch fi29 {
                                        case 1:
                                            fmt.Println(v6, v3)
                                        }
                                    case "":
                                        for fi36 := 0; fi36 < 14; fi36++ {
                                            v2 = 4294967295
                                            _ = fi36
                                        }
                                    default:
                                        for fi37 := 0; fi37 < 19; fi37++ {
                                            v2 = v23
                                            v2 -= 65535
                                            _ = fi37
                                        }
                                    }
                                default:
                                    if (3.14159 > ((3.14159 - 0.1) + float64(float32(0.1)))) {
                                        v1 -= (v1 >> (((v23 - (v23 & v2))) & 63))
                                        v23 -= (v23 & v23)
                                    } else {
                                        v3 += (v3 * 2147483647)
                                    }
                                }
                                v38 := GenS2{F0: v3, F1: 18446744073709551615, F2: "\n\t"}
                                if true {
                                    v1 = 4294967295
                                    v39 := GenS5{F0: 2, F1: false}
                                    _ = v39
                                }
                                _ = v38
                            }
                            v3 = -1
                            if false {
                                switch ("\n\t" + "") {
                                case "\n\t":
                                    v40 := GenS6{F0: 793, F1: (v1 % (((v1 & (v1 ^ v1))) | 1)), F2: -32768}
                                    _ = v40
                                }
                                v41 := GenS6{F0: v3, F1: (v1 << ((v23) & 63)), F2: 42}
                                var v42 int32 = (v3 & v3)
                                _ = v41
                                _ = v42
                            }
                            _ = fi29
                        }
                    case "foo bar":
                        for fi43 := 0; fi43 < 9; fi43++ {
                            if false {
                                for fi44 := 0; fi44 < 9; fi44++ {
                                    v45 := GenS7{F0: 0}
                                    v1 += (v1 | v1)
                                    if false {
                                        switch (("a" + ("a" + "foo bar")) + ("hello" + ("a" + "foo bar"))) {
                                        case "\n\t":
                                            v46 := GenS8{F0: (v6 % ((679) | 1)), F1: 32767}
                                            _ = v46
                                        case "hello":
                                            v1 = (v1 - v1)
                                        case "":
                                            if (fi44 > (fi43 * fi44)) {
                                                fmt.Println(v23, v1)
                                                if (3.14159 <= (3.14159 * -2.25)) {
                                                    var v47 uint = 879
                                                    _ = v47
                                                }
                                                v48 := GenS9{F0: (!false)}
                                                v49 := GenS9{F0: false}
                                                _ = len(v5)
                                                _ = v48
                                                _ = v49
                                            } else {
                                                fmt.Println(v1, v2, v3)
                                                v50 := []int{(fi43 ^ (fi43 - fi43)), fi44}
                                                var v51 uint8 = 1
                                                _ = v50
                                                _ = v51
                                            }
                                        }
                                        fmt.Println(v23)
                                    }
                                    _ = fi44
                                    _ = v45
                                }
                                for fi52 := 0; fi52 < 8; fi52++ {
                                    v53 := GenS9{F0: ((v1 ^ v1) != v1)}
                                    _ = fi52
                                    _ = v53
                                }
                                fmt.Println(v23, v2)
                            } else {
                                for fi54 := 0; fi54 < 10; fi54++ {
                                    v6 = v6
                                    if (255 > genFn0(255, 111)) {
                                        var v55 uint32 = v1
                                        _ = v55
                                    }
                                    _ = fi54
                                }
                            }
                            var v56 int16 = 1000
                            v57 := GenS9{F0: (2 >= 42)}
                            _ = fi43
                            _ = v56
                            _ = v57
                        }
                    case "":
                        v58 := []int{389, -502, 9223372036854775807, -1}
                        _ = v58
                    }
                }
            }
            fmt.Println(v1, v2)
            _ = v23
        }
    }
    switch ("" + "a") {
    case "\n\t":
        for fi59 := 0; fi59 < 10; fi59++ {
            switch (fi59 / ((fi59) | 1)) {
            case 775:
                for fi60 := 0; fi60 < 18; fi60++ {
                    _ = len(v4)
                    v3 -= v3
                    var v61 uint = (v2 & (v2 + v2))
                    var v62 uint32 = 2147483648
                    _ = fi60
                    _ = v61
                    _ = v62
                }
            case 9223372036854775807:
                v1 += 448
            }
            _ = fi59
        }
    case "hello":
        _ = len(v4)
    default:
        for fi63 := 0; fi63 < 16; fi63++ {
            v3 = v3
            fmt.Println(v1, v2)
            var v64 int64 = -9223372036854775808
            _ = fi63
            _ = v64
        }
    }
    var v65 uint32 = 4294967295
    var v66 uint16 = 65535
    for fi67 := 0; fi67 < 2; fi67++ {
        fmt.Println(v6, v1, v6)
        _ = fi67
    }
    if false {
        for fi68 := 0; fi68 < 6; fi68++ {
            var v69 uint8 = uint8(v1)
            _ = fi68
            _ = v69
        }
        v1 -= (v1 ^ uint32(v2))
    }
    for fi70 := 0; fi70 < 15; fi70++ {
        var v71 uint = v2
        _ = fi70
        _ = v71
    }
    if true {
        for fi72 := 0; fi72 < 16; fi72++ {
            v73 := GenS6{F0: v3, F1: (v65 ^ (^(v65))), F2: 42}
            _ = fi72
            _ = v73
        }
        v3 = (v3 ^ (^(v3)))
    } else {
        var v74 bool = false
        var v75 int = 9223372036854775807
        _ = v74
        _ = v75
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
}
