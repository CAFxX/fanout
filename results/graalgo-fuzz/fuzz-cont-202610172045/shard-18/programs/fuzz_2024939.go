package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 int8; F1 uint8; F2 bool }
type GenS5 struct { F0 uint32; F1 int16 }
type GenS6 struct { F0 uint16; F1 uint; F2 int64 }
type GenS7 struct { F0 bool; F1 int32; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{957, 0, 4294967295, 4294967295}
    v3 := make(map[uint32]uint32)
    var v4 int64 = -1
    v5 := []int{-9223372036854775808, 32767, 32767}
    v6 := make(map[string]uint32)
    v4 -= (v4 >> ((65535) & 63))
    switch 1 {
    case -1:
        var v7 uint16 = 42
        _ = v7
    default:
        fmt.Println(v4)
    }
    var v8 uint8 = 255
    var v9 int64 = 529
    if true {
        v9 += (-(v9))
    } else {
        if true {
            v10 := GenS1{F0: 65535, F1: -1}
            _ = v10
        } else {
            v11 := GenS1{F0: 0, F1: -263}
            var v12 int32 = -2147483648
            _ = v11
            _ = v12
        }
        var v13 float32 = 3.14159
        if true {
            delete(v6, "foo bar")
            if true {
                v14 := GenS2{F0: -1}
                var v15 int64 = v4
                _ = v14
                _ = v15
            }
        }
        _ = v13
    }
    v8 = 146
    switch -2147483648 {
    case 1:
        for fi16 := 0; fi16 < 19; fi16++ {
            if (("a" + "a") >= "\n\t") {
                if (!true) {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        v18 := GenS2{F0: -131}
                        v9 = (v9 ^ (v4 + -9223372036854775808))
                        _ = fi17
                        _ = v18
                    }
                    if true {
                        switch ("a" + "") {
                        case "foo bar":
                            v8 = (v8 >> ((42) & 63))
                        }
                        fmt.Println(v4)
                    }
                    var v19 int64 = v9
                    _ = v19
                }
                _ = len(v3)
                var v20 uint32 = 65535
                _ = v20
            }
            if (!true) {
                if true {
                    for fi21 := 0; fi21 < 9; fi21++ {
                        fmt.Println(v8, v9, v8)
                        _ = fi21
                    }
                } else {
                    var v22 int64 = v9
                    var v23 uint32 = 65535
                    var v24 int64 = v4
                    _ = v22
                    _ = v23
                    _ = v24
                }
                if false {
                    var v25 string = ""
                    v26 := GenS3{F0: 42, F1: 1, F2: true}
                    _ = v25
                    _ = v26
                }
                switch 32767 {
                case -679:
                    switch ("a" + "\n\t") {
                    case "foo bar":
                        v27 := GenS4{F0: 100, F1: v8, F2: (v9 == (v4 * 1))}
                        _ = v27
                    default:
                        v28 := GenS3{F0: -1, F1: 171, F2: false}
                        _ = v28
                    }
                case -531:
                    for fi29 := 0; fi29 < 1; fi29++ {
                        v8 = (v8 / (((v8 / ((v8) | 1))) | 1))
                        var v30 int = (fi29 % ((fi29) | 1))
                        v4 = v4
                        var v31 bool = (0.1 >= 3.14159)
                        if (!v31) {
                            var v32 string = ("\n\t" + "hello")
                            var v33 float64 = (1e10 - (0.1 + 0.1))
                            if true {
                                var v34 float64 = 1e10
                                switch "hello" {
                                case "\n\t":
                                    v31 = false
                                case "hello":
                                    var v35 uint32 = 673
                                    _ = v35
                                case "":
                                    var v36 int8 = 42
                                    _ = v36
                                }
                                v8 = v8
                                _ = v34
                            }
                            _ = v32
                            _ = v33
                        } else {
                            v37 := GenS2{F0: int32(fi29)}
                            v38 := GenS5{F0: 42, F1: -32768}
                            var v39 int16 = 1
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        if v31 {
                            fmt.Println(v4, fi29, v8)
                            fmt.Println(v4)
                            if (!(-10 > -4)) {
                                v40 := GenS5{F0: 0, F1: 32767}
                                _ = v40
                            }
                        }
                        v8 -= (v8 >> ((2) & 63))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                case 742:
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                default:
                    v41 := GenS5{F0: 281, F1: 1}
                    _ = v41
                }
            }
            fmt.Println(v4)
            _ = fi16
        }
    case 32768:
        for fi42 := 0; fi42 < 15; fi42++ {
            fmt.Println(v9, v9, v4)
            switch -54 {
            case 139:
                v6["\n\t"] = 435
            case 98:
                for fi43 := 0; fi43 < 11; fi43++ {
                    var v44 uint32 = 2147483647
                    if (1 != 30) {
                        var v45 int16 = -1
                        var v46 bool = (-128 > -98)
                        if v46 {
                            if v46 {
                                v9 -= v9
                            } else {
                                v44 = v44
                                if (-1 < (v9 / (((^(v4))) | 1))) {
                                    v47 := GenS6{F0: 818, F1: 42, F2: (-(v9))}
                                    _ = v47
                                } else {
                                    v48 := GenS7{F0: true, F1: -2147483648, F2: (true && false)}
                                    if v46 {
                                        if v46 {
                                            _ = len(v1)
                                        } else {
                                            var v49 uint = 1
                                            fmt.Println(v49, v49, v46)
                                            var v50 float64 = 1.5
                                            _ = v49
                                            _ = v50
                                        }
                                    } else {
                                        fmt.Println(v46, v46)
                                        v45 = (v45 + v45)
                                        var v51 uint = 0
                                        _ = v51
                                    }
                                    delete(v3, v44)
                                    v8 += v8
                                    _ = len(v6)
                                    switch "hello" {
                                    case "":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "hello":
                                        v4 = (v9 | -246)
                                        v52 := GenS7{F0: (true && (4294967295 == 18446744073709551615)), F1: 42, F2: false}
                                        _ = v52
                                    default:
                                        var v53 uint = 4294967295
                                        _ = v53
                                    }
                                    v8 -= v8
                                    _ = v48
                                }
                                var v54 bool = v46
                                _ = v54
                            }
                            v45 = v45
                            fmt.Println(fi42, v46)
                        }
                        _ = v45
                        _ = v46
                    }
                    _ = fi43
                    _ = v44
                }
            case -702:
                var v55 int64 = v9
                _ = v55
            default:
                v9 = v9
            }
            _ = fi42
        }
    case 42:
        for fi56 := 0; fi56 < 2; fi56++ {
            if (0 >= -2147483648) {
                v57 := GenS2{F0: 949}
                v9 = (^(v4))
                var v58 float32 = 1e10
                _ = v57
                _ = v58
            } else {
                var v59 int = (fi56 << ((223) & 63))
                v8 -= 29
                _ = v59
            }
            var v60 uint16 = 498
            _ = fi56
            _ = v60
        }
    default:
        var v61 float64 = 0.1
        _ = v61
    }
    var v62 int16 = -14
    v8 = (v8 / ((v8) | 1))
    v9 = (v4 % (((-(v9))) | 1))
    if ((!false) || false) {
        var v63 uint64 = 1
        _ = v63
    } else {
        if (v8 >= (v8 ^ 56)) {
            for fi64 := 0; fi64 < 8; fi64++ {
                var v65 int64 = v9
                _ = fi64
                _ = v65
            }
            v66 := GenS2{F0: 1}
            _ = v66
        } else {
            for fi67 := 0; fi67 < 12; fi67++ {
                v8 = (v8 ^ v8)
                var v68 int8 = 42
                _ = fi67
                _ = v68
            }
            var v69 uint32 = 4294967295
            _ = v69
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v8
    _ = v9
}
