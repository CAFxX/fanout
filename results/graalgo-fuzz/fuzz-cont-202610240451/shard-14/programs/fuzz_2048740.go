package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 int16; F1 int64 }
type GenS6 struct { F0 uint32; F1 uint64 }
func genFn0(p0 uint8, p1 string) int {
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 uint16 = 42
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 381
    var v10 float64 = float64(1e10)
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 915
}
func genFn1() int64 {
    if ((127 < 127) || ((!false) && false)) {
        var v11 uint8 = 0
        fmt.Println(v11)
        v12 := GenS1{F0: false}
        _ = v11
        _ = v12
    }
    v13 := GenS2{F0: 127}
    _ = v13
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]int)
    var v3 uint = 42
    var v4 int16 = 680
    var v5 float32 = 1.5
    var v6 uint64 = 18446744073709551615
    v14 := []int{9223372036854775807, 32767, genFn0(uint8(v6), "foo bar"), 716}
    if false {
        v6 -= (v6 ^ 4294967295)
        if false {
            if ((1 > 758) || ("\n\t" == ("" + ("\n\t" + "a")))) {
                var v15 int16 = 153
                _ = v15
            } else {
                var v16 uint64 = v6
                _ = v16
            }
        } else {
            v17 := v1[0]
            _ = v17
        }
        if false {
            switch -9223372036854775808 {
            case -375:
                v5 += (v5 + float32((0.1 + 0.1)))
            case 9223372036854775807:
                var v18 int8 = 57
                _ = v18
            }
        } else {
            var v19 uint64 = v6
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        var v21 bool = false
        _ = fi20
        _ = v21
    }
    delete(v1, 589)
    v22 := GenS3{F0: ((v6 & 1) < v6), F1: 858, F2: ""}
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := GenS4{F0: 0, F1: 2}
        _ = fi23
        _ = v24
    }
    if (2 != 2) {
        var v25 int16 = (v4 / (((v4 / ((v4) | 1))) | 1))
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v3)
            if (false || (!true)) {
                if (false || ((v6 << ((v3) & 63)) != v6)) {
                    fmt.Println(v3, v3, v3)
                    v4 = v25
                    for fi27 := 0; fi27 < 14; fi27++ {
                        v25 += (v25 ^ v4)
                        fmt.Println(v25)
                        _ = fi27
                    }
                } else {
                    var v28 uint8 = 2
                    _ = v28
                }
                v14 = append(v14, fi26)
                v5 = (v5 * v5)
            }
            v29 := GenS1{F0: (false || false)}
            _ = fi26
            _ = v29
        }
        var v30 uint64 = (^(v6))
        _ = v25
        _ = v30
    } else {
        v3 -= 42
        _ = len(v14)
        if false {
            if (!true) {
                if (((32767 != 255) || (-53 == 0)) && (32767 == 4294967295)) {
                    v3 = v3
                }
                for fi31 := 0; fi31 < 5; fi31++ {
                    switch (^(fi31)) {
                    case 54:
                        v6 -= 18446744073709551615
                    case -2147483648:
                        var v32 float64 = ((1.5 / 1.0) * (((0.1 / -2.25) * float64(0.0)) * 3.14159))
                        _ = v32
                    }
                    delete(v1, 331)
                    switch (("" + ("" + "\n\t")) + "\n\t") {
                    case "foo bar":
                        switch fi31 {
                        case -1:
                            if false {
                                _ = len(v1)
                                switch -609 {
                                case 0:
                                    var v33 int32 = 2147483647
                                    _ = v33
                                case 42:
                                    var v34 uint = v3
                                    _ = v34
                                case -162:
                                    fmt.Println(v4, v3, v5)
                                }
                            } else {
                                v35 := GenS1{F0: ((float64((float32(0.1) - 0.1)) < (((1e10 + -2.25) + 1e10) + 1e10)) && true)}
                                for fi36 := 0; fi36 < 5; fi36++ {
                                    var v37 int8 = 1
                                    _ = fi36
                                    _ = v37
                                }
                                _ = v35
                            }
                        }
                    case "a":
                        fmt.Println(fi31, v4)
                    default:
                        var v38 float32 = float32(0.0)
                        _ = v38
                    }
                    _ = fi31
                }
            } else {
                if true {
                    v6 = 18446744073709551615
                    v14 = append(v14, 1)
                } else {
                    var v39 float32 = v5
                    _ = v39
                }
                fmt.Println(v6)
            }
        } else {
            fmt.Println(v3, v4)
            v40 := GenS5{F0: -746, F1: 322}
            fmt.Println(v4, v6, v4)
            _ = v40
        }
    }
    v41 := GenS5{F0: -235, F1: genFn1()}
    for fi42 := 0; fi42 < 1; fi42++ {
        fmt.Println(v6)
        _ = fi42
    }
    if (127 >= 58) {
        v43 := GenS4{F0: 0, F1: 42}
        fmt.Println(v4, v3)
        _ = v43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        _ = len(v1)
        if (1 < 92) {
            v45 := v1[0]
            var v46 int8 = 1
            _ = v45
            _ = v46
        } else {
            v47 := GenS3{F0: (!(false && true)), F1: 255, F2: ("\n\t" + "hello")}
            _ = v47
        }
        _ = fi44
    }
    if (false || false) {
        if false {
            var v48 string = ""
            var v49 uint32 = 2147483647
            fmt.Println(v6)
            _ = v48
            _ = v49
        }
    } else {
        var v50 int32 = -1
        for fi51 := 0; fi51 < 15; fi51++ {
            v4 = (v4 << ((1) & 63))
            var v52 int = fi51
            for fi53 := 0; fi53 < 14; fi53++ {
                fmt.Println(fi53)
                var v54 string = "hello"
                _ = fi53
                _ = v54
            }
            _ = fi51
            _ = v52
        }
        _ = v50
    }
    if (-20 >= 127) {
        v55 := GenS1{F0: (-2.25 >= ((0.1 / 1e10) * 0.0))}
        switch 642 {
        case -1:
            v56 := GenS5{F0: (v4 - -259), F1: int64(v6)}
            _ = v56
        case 820:
            for fi57 := 0; fi57 < 8; fi57++ {
                for fi58 := 0; fi58 < 15; fi58++ {
                    fmt.Println(fi58, v6)
                    delete(v1, 2147483648)
                    _ = fi58
                }
                if (true && false) {
                    v6 = 4294967295
                } else {
                    v5 = v5
                    v5 = -2.25
                    v59 := GenS5{F0: v4, F1: 1}
                    _ = v59
                }
                var v60 int16 = v4
                _ = fi57
                _ = v60
            }
        case -515:
            if false {
                _ = len(v14)
            } else {
                if ((2 != 122) || (574 > 0)) {
                    v5 = 1e10
                    v5 = 1e10
                } else {
                    for fi61 := 0; fi61 < 12; fi61++ {
                        v6 = v6
                        v62 := GenS4{F0: -1, F1: 963}
                        if (65535 != 65535) {
                            var v63 int = 689
                            v64 := GenS6{F0: 541, F1: 15}
                            _ = v63
                            _ = v64
                        }
                        _ = fi61
                        _ = v62
                    }
                    var v65 uint16 = 65535
                    _ = v65
                }
            }
        }
        v5 -= v5
        _ = v55
    }
    v5 = ((float32((1.5 / (-2.25 - 0.1))) + v5) / v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
