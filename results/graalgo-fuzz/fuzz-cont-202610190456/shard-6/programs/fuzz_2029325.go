package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint16; F1 int32 }
type GenS8 struct { F0 bool; F1 uint64 }
func genFn0() int32 {
    var v5 bool = (false && false)
    var v6 uint32 = 981
    fmt.Println(v6, v6)
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 uint32 = 2147483647
    v4 := make(map[string]int)
    for fi7 := 0; fi7 < 12; fi7++ {
        v1 = (v1 << ((198) & 63))
        if true {
            v8 := GenS1{F0: "\n\t", F1: fi7}
            _ = v8
        } else {
            var v9 uint = 447
            if (!(!true)) {
                v1 -= (v1 << ((v9) & 63))
                v10 := GenS2{F0: (^(v1))}
                v11 := GenS3{F0: false}
                _ = v10
                _ = v11
            }
            _ = v9
        }
        v12 := GenS3{F0: true}
        _ = fi7
        _ = v12
    }
    fmt.Println(v3, v3)
    var v13 uint8 = 255
    for fi14 := 0; fi14 < 9; fi14++ {
        switch int(v1) {
        case 568:
            for fi15 := 0; fi15 < 8; fi15++ {
                if true {
                    fmt.Println(v1)
                    v16 := v4[(("\n\t" + ("" + "")) + "hello")]
                    var v17 uint32 = (v3 - 0)
                    _ = v16
                    _ = v17
                } else {
                    fmt.Println(fi15)
                    var v18 int = (^(fi15))
                    _ = v18
                }
                _ = fi15
            }
        case -121:
            var v19 int = fi14
            _ = v19
        default:
            if (v1 == genFn0()) {
                var v20 uint16 = 42
                _ = v20
            }
        }
        _ = fi14
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v3 += (v3 - (v3 >> ((18446744073709551615) & 63)))
        _ = fi21
    }
    var v22 uint8 = 42
    fmt.Println(v22, v22)
    fmt.Println(v13)
    for fi23 := 0; fi23 < 1; fi23++ {
        v1 = (v1 % (((v1 * int32(fi23))) | 1))
        if (false || true) {
            v24 := []int{fi23}
            _ = len(v2)
            if false {
                for fi25 := 0; fi25 < 13; fi25++ {
                    if (!((-(fi25)) > (fi23 + fi25))) {
                        var v26 uint16 = 42
                        v1 = 80
                        _ = len(v4)
                        v27 := GenS1{F0: ("foo bar" + "hello"), F1: fi25}
                        _ = v26
                        _ = v27
                    } else {
                        var v28 int = fi25
                        _ = v28
                    }
                    _ = fi25
                }
                var v29 string = ("hello" + ("hello" + ("foo bar" + ("\n\t" + "a"))))
                v30 := GenS4{F0: 647, F1: 196, F2: 504}
                _ = v29
                _ = v30
            } else {
                fmt.Println(v1, v3)
                for fi31 := 0; fi31 < 15; fi31++ {
                    fmt.Println(v3, v1)
                    if (true || false) {
                        switch fi23 {
                        case 996:
                            var v32 uint8 = v22
                            _ = v32
                        default:
                            var v33 uint64 = 1
                            _ = v33
                        }
                        if (636 >= 18446744073709551615) {
                            v13 = (v22 ^ 83)
                            fmt.Println(fi31, v13, fi31)
                            var v34 int = 32767
                            _ = v34
                        } else {
                            var v35 float32 = 0.1
                            _ = v35
                        }
                        var v36 int32 = int32(v13)
                        switch (fi31 + fi23) {
                        case -352:
                            v36 -= (-(v1))
                        case 42:
                            var v37 int16 = 32767
                            _ = v37
                        case 9223372036854775807:
                            v3 += (^(v3))
                        default:
                            var v38 int16 = 42
                            _ = v38
                        }
                        _ = v36
                    }
                    if (!false) {
                        var v39 float64 = -2.25
                        v3 = v3
                        switch (^(fi31)) {
                        case -1:
                            v22 += v13
                        case -9223372036854775808:
                            if (2147483648 != v3) {
                                if true {
                                    var v40 uint16 = 32767
                                    fmt.Println(v3, v40, v39)
                                    _ = v40
                                }
                                v41 := GenS1{F0: "", F1: int(v13)}
                                _ = v41
                            }
                        case 2147483647:
                            if true {
                                var v42 uint16 = 522
                                if false {
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                }
                                v13 -= (v22 >> ((0) & 63))
                                v42 -= v42
                                v39 = v39
                                v13 += v22
                                v1 = int32(fi31)
                                var v43 int8 = -1
                                switch ("hello" + "\n\t") {
                                case "":
                                    var v44 int8 = v43
                                    _ = v44
                                }
                                v1 = (v1 + (^(v1)))
                                _ = v42
                                _ = v43
                            } else {
                                v4[("hello" + "a")] = -1
                                v3 += (v3 | (v3 + (v3 * v3)))
                            }
                        }
                        _ = v39
                    }
                    _ = fi31
                }
                fmt.Println(v22, v1, v1)
            }
            _ = v24
        } else {
            v45 := GenS5{F0: (v3 / (((v3 >> ((1) & 63))) | 1)), F1: 498}
            for fi46 := 0; fi46 < 17; fi46++ {
                v13 = v22
                var v47 float32 = (1.5 + 1e10)
                var v48 string = ("\n\t" + "hello")
                switch fi23 {
                case -1:
                    v2[9223372036854775807] = (941 > uint64(v13))
                case 1:
                    if true {
                        if true {
                            v49 := GenS1{F0: ((("\n\t" + ("\n\t" + (v48 + v48))) + v48) + v48), F1: (fi23 / ((fi46) | 1))}
                            v50 := GenS5{F0: 309, F1: 2}
                            var v51 int16 = 4
                            _ = v49
                            _ = v50
                            _ = v51
                        }
                        var v52 int64 = 86
                        if true {
                            if false {
                                v53 := GenS6{F0: (false || (!(835 > 42))), F1: v52}
                                _ = v53
                            }
                            fmt.Println(fi23, fi46)
                        }
                        v22 = v22
                        if ((v48 + (v48 + v48)) != "") {
                            v54 := GenS5{F0: 2147483647, F1: 18446744073709551615}
                            v48 = v48
                            var v55 float64 = 3.14159
                            _ = v54
                            _ = v55
                        }
                        _ = v52
                    }
                case 32767:
                    v1 = genFn0()
                default:
                    var v56 uint16 = 157
                    _ = v56
                }
                var v57 int64 = 0
                v58 := GenS1{F0: "", F1: 32768}
                _ = fi46
                _ = v47
                _ = v48
                _ = v57
                _ = v58
            }
            v3 += v3
            _ = v45
        }
        _ = fi23
    }
    for fi59 := 0; fi59 < 3; fi59++ {
        for fi60 := 0; fi60 < 13; fi60++ {
            v13 -= v13
            v61 := GenS7{F0: 65535, F1: v1}
            v62 := GenS5{F0: (v3 % ((v3) | 1)), F1: 995}
            _ = fi60
            _ = v61
            _ = v62
        }
        for fi63 := 0; fi63 < 14; fi63++ {
            v64 := GenS4{F0: -282, F1: v22, F2: 1}
            v2[32767] = true
            v65 := GenS8{F0: (32767 <= -32768), F1: 328}
            _ = fi63
            _ = v64
            _ = v65
        }
        for fi66 := 0; fi66 < 13; fi66++ {
            v13 = 156
            v1 -= v1
            v1 -= 2147483647
            _ = fi66
        }
        _ = fi59
    }
    for fi67 := 0; fi67 < 15; fi67++ {
        var v68 uint = 0
        _ = fi67
        _ = v68
    }
    for fi69 := 0; fi69 < 17; fi69++ {
        switch "foo bar" {
        case "hello":
            delete(v4, "")
        }
        var v70 int8 = 0
        _ = fi69
        _ = v70
    }
    _ = len(v2)
    v71 := GenS8{F0: true, F1: 42}
    delete(v4, "hello")
    if false {
        v72 := []int{9223372036854775807}
        _ = v72
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v71
}
