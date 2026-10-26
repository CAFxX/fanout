package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 string; F2 int }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS8 struct { F0 int32; F1 int16 }
func genFn0() int32 {
    v7 := []int{0, 1, -9223372036854775808, 4294967295}
    var v8 int32 = 42
    v9 := GenS1{F0: 0}
    v10 := GenS2{F0: 0, F1: "\n\t", F2: -9223372036854775808}
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint) uint32 {
    p0 = 4294967295
    var v11 int8 = 127
    if ((^(p0)) >= p0) {
        p0 += (p0 - (p0 / (((p0 >> ((p0) & 63))) | 1)))
        v11 = -29
        if false {
            var v12 float64 = (0.1 + float64(1.5))
            fmt.Println(v12, v11)
            switch "hello" {
            case "foo bar":
                switch (("" + ("a" + ("foo bar" + "\n\t"))) + "") {
                case "hello":
                    var v13 int64 = 42
                    _ = v13
                case "foo bar":
                    fmt.Println(v11)
                }
            case "hello":
                v14 := []int{-604, 4294967295, -1, int(p0)}
                _ = v14
            case "a":
                var v15 bool = (255 != 1)
                _ = v15
            }
            _ = v12
        }
    } else {
        v11 = 0
        v11 -= 42
    }
    var v16 string = ("" + "")
    _ = p0
    _ = v11
    _ = v16
    return 2
}

func main() {
    var v1 uint = 92
    var v2 int16 = -787
    v3 := make(map[int]bool)
    v4 := make(map[uint32]bool)
    var v5 float64 = 3.14159
    v6 := make(map[string]bool)
    v5 = v5
    var v17 float64 = v5
    var v18 uint16 = 32767
    switch 0 {
    case 2147483647:
        if true {
            var v19 uint = (v1 ^ 1)
            var v20 uint64 = uint64(v19)
            v21 := GenS1{F0: v19}
            _ = v19
            _ = v20
            _ = v21
        }
    case -956:
        var v22 int8 = -42
        _ = v22
    case -1:
        if (true || true) {
            var v23 uint32 = genFn1((v1 >> ((v1) & 63)))
            fmt.Println(v17, v23, v1)
            _ = v23
        } else {
            v2 = (v2 - 32767)
            _ = len(v4)
            if true {
                if true {
                    var v24 int16 = v2
                    _ = v24
                }
                v25 := GenS2{F0: 148, F1: ("a" + ("" + "hello")), F2: 32767}
                _ = v25
            } else {
                delete(v3, 127)
                if true {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        switch (("\n\t" + "hello") + ("foo bar" + ("\n\t" + "\n\t"))) {
                        case "":
                            var v27 float32 = float32(v17)
                            _ = v27
                        case "hello":
                            v1 = (v1 >> (((v1 - (v1 & 0))) & 63))
                        case "foo bar":
                            v17 += 0.1
                        default:
                            if true {
                                for fi28 := 0; fi28 < 14; fi28++ {
                                    if ((!(!true)) || true) {
                                        fmt.Println(v2)
                                    } else {
                                        fmt.Println(v1)
                                    }
                                    if ((!false) && false) {
                                        var v29 int64 = -1
                                        _ = v29
                                    }
                                    _ = fi28
                                }
                                switch (fi26 & (^(fi26))) {
                                case -1:
                                    if ((^(v2)) >= (v2 / ((v2) | 1))) {
                                        var v30 float32 = 3.14159
                                        for fi31 := 0; fi31 < 4; fi31++ {
                                            v18 = v18
                                            _ = fi31
                                        }
                                        _ = v30
                                    }
                                case 309:
                                    fmt.Println(v5, fi26, v2)
                                case 0:
                                    v32 := GenS3{F0: ((0 == 0) || (!(!true)))}
                                    _ = v32
                                }
                                for fi33 := 0; fi33 < 5; fi33++ {
                                    if false {
                                        v34 := []int{1, fi26, fi33, fi26}
                                        v17 += (0.1 / v5)
                                        _ = v34
                                    } else {
                                        if false {
                                            var v35 string = "\n\t"
                                            v1 = v1
                                            v36 := GenS2{F0: 149, F1: v35, F2: 32767}
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    var v37 uint32 = 4294967295
                                    _ = fi33
                                    _ = v37
                                }
                            }
                        }
                        var v38 bool = (v17 == v17)
                        if (18446744073709551615 != 2) {
                            v3[(fi26 + (fi26 & (fi26 & fi26)))] = v38
                            _ = len(v6)
                            v39 := GenS4{F0: (v1 >> ((v1) & 63)), F1: 2, F2: 255}
                            _ = v39
                        } else {
                            v18 = (v18 * (^(v18)))
                        }
                        _ = fi26
                        _ = v38
                    }
                    v40 := GenS4{F0: v1, F1: 42, F2: 55}
                    if (float64(float32(float64(1e10))) == v5) {
                        var v41 uint16 = v18
                        _ = v41
                    }
                    _ = v40
                }
                if false {
                    v42 := GenS5{F0: (v2 - 744), F1: -1}
                    v43 := GenS6{F0: true}
                    _ = v42
                    _ = v43
                }
            }
        }
    default:
        if true {
            v1 = (v1 / ((602) | 1))
            var v44 float64 = ((v17 + v17) / 1.0)
            v45 := GenS7{F0: uint16(v1), F1: v2, F2: 475}
            _ = v44
            _ = v45
        }
    }
    v46 := GenS1{F0: (v1 | (v1 * (v1 % ((1) | 1))))}
    v47 := []int{-221, 128, 0, 32768}
    fmt.Println(v1)
    switch 32768 {
    case 32768:
        if false {
            if len(v47) > 0 {
                v47[0] = 609
            }
            var v48 int8 = 54
            switch -9223372036854775808 {
            case 532:
                var v49 uint32 = 85
                _ = v49
            default:
                var v50 int16 = 1
                _ = v50
            }
            _ = v48
        }
    case -9223372036854775808:
        switch 1 {
        case 0:
            v1 -= (v1 * (v1 & 0))
        case 9223372036854775807:
            for fi51 := 0; fi51 < 9; fi51++ {
                v6["\n\t"] = (713 != 2)
                v52 := GenS8{F0: -2147483648, F1: (v2 / ((v2) | 1))}
                _ = len(v47)
                _ = fi51
                _ = v52
            }
        case 4294967295:
            v53 := GenS7{F0: 65535, F1: int16(v1), F2: 19}
            _ = v53
        default:
            v1 = 0
        }
    case -2147483648:
        var v54 uint32 = 2147483648
        _ = v54
    default:
        switch 2147483647 {
        case -2147483648:
            var v55 uint = (v1 / (((v1 | v1)) | 1))
            _ = v55
        case 830:
            _ = len(v4)
        case 128:
            v18 = 83
        default:
            if len(v47) > 0 {
                v47[0] = -702
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
