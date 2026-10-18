package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint32 }
type GenS4 struct { F0 int64; F1 int64 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint32; F1 int64; F2 uint64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint; F1 uint16 }
type GenS12 struct { F0 int32; F1 uint8; F2 uint8 }
func genFn0(p0 uint64) int {
    p0 += 0
    var v6 int = 2147483647
    fmt.Println(v6, v6)
    v7 := GenS1{F0: -766, F1: uint32(p0), F2: uint8(p0)}
    switch "a" {
    case "foo bar":
        var v8 int8 = 0
        _ = v8
    case "a":
        var v9 int = (v6 + 2147483647)
        _ = v9
    case "":
        v6 = -1
    }
    _ = p0
    _ = v6
    _ = v7
    return v6
}
func genFn1(p0 uint8, p1 bool, p2 int8) int {
    if (9223372036854775807 >= -1) {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 uint32 = 2
            v12 := GenS2{F0: 0, F1: -1, F2: 1}
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 int = (fi13 | fi13)
            v15 := []int{genFn0(4294967295), fi13, (^(fi13)), (v14 / ((int(p2)) | 1))}
            _ = fi13
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            for fi17 := 0; fi17 < 7; fi17++ {
                p1 = p1
                _ = fi17
            }
            if p1 {
                v18 := GenS1{F0: -9223372036854775808, F1: 555, F2: (^(p0))}
                if true {
                    v19 := GenS3{F0: 425, F1: 86, F2: 18}
                    p2 = 42
                    _ = v19
                }
                _ = v18
            }
            _ = fi16
        }
    } else {
        v20 := GenS2{F0: 0, F1: 1000, F2: 42}
        p0 = 117
        p0 += p0
        _ = v20
    }
    if (9223372036854775807 < 2147483647) {
        if (uint(p0) >= 0) {
            if p1 {
                if (!(false && (-2.25 > float64(0.0)))) {
                    p0 = (^(p0))
                    var v21 int = 128
                    _ = v21
                }
                v22 := []int{32768}
                switch 127 {
                case 645:
                    v22 = append(v22, -725)
                case 127:
                    switch ("a" + ("\n\t" + "hello")) {
                    case "\n\t":
                        for fi23 := 0; fi23 < 2; fi23++ {
                            for fi24 := 0; fi24 < 4; fi24++ {
                                var v25 float32 = 3.14159
                                v26 := GenS4{F0: 42, F1: 9223372036854775807}
                                p1 = (127 > 0)
                                if len(v22) >= 2 {
                                    v22 = v22[1:len(v22)]
                                }
                                if p1 {
                                    v27 := GenS5{F0: 42, F1: 1}
                                    switch -428 {
                                    case -454:
                                        fmt.Println(fi23)
                                    case 376:
                                        var v28 int64 = 0
                                        _ = v28
                                    case 128:
                                        p1 = p1
                                    }
                                    _ = v27
                                }
                                v29 := GenS6{F0: (p1 && p1)}
                                _ = fi24
                                _ = v25
                                _ = v26
                                _ = v29
                            }
                            var v30 int32 = 1
                            _ = fi23
                            _ = v30
                        }
                    case "a":
                        p1 = (516 >= 555)
                    }
                }
                _ = v22
            }
        } else {
            fmt.Println(p1, p2, p0)
            var v31 int = -232
            _ = v31
        }
        if (!true) {
            var v32 bool = p1
            v33 := GenS4{F0: 42, F1: -632}
            v34 := []int{128, -765, 127, genFn0(1)}
            _ = v32
            _ = v33
            _ = v34
        } else {
            if true {
                if false {
                    p1 = false
                } else {
                    if (215 == 1) {
                        var v35 bool = (1.5 <= 0.1)
                        for fi36 := 0; fi36 < 13; fi36++ {
                            p0 -= 127
                            fmt.Println(v35, p2)
                            p0 -= 128
                            _ = fi36
                        }
                        _ = v35
                    } else {
                        v37 := []int{1, genFn0(42), -9223372036854775808}
                        v38 := GenS7{F0: 42}
                        v39 := GenS8{F0: 802}
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                }
                p0 += p0
            } else {
                p1 = ((p2 << ((0) & 63)) <= 100)
                var v40 bool = p1
                _ = v40
            }
            v41 := GenS9{F0: 2, F1: -1, F2: 601}
            _ = v41
        }
    } else {
        v42 := []int{-1, 2147483647, -168}
        _ = v42
    }
    v43 := GenS10{F0: 0}
    var v44 int = 32767
    switch 42 {
    case 128:
        v45 := GenS7{F0: 18446744073709551615}
        _ = v45
    case -54:
        if (p1 && true) {
            p0 -= p0
            p0 = (p0 | p0)
        } else {
            var v46 float64 = 1.5
            v46 = v46
            _ = v46
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v43
    _ = v44
    return (v44 + v44)
}

func main() {
    v1 := []uint32{2147483647, 802, 2}
    var v2 float64 = -2.25
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 float32 = (0.1 + 0.0)
    var v47 uint8 = 128
    for fi48 := 0; fi48 < 15; fi48++ {
        var v49 string = "foo bar"
        var v50 int8 = 1
        switch v49 {
        case "":
            for fi51 := 0; fi51 < 20; fi51++ {
                switch (fi48 >> ((0) & 63)) {
                case -9223372036854775808:
                    v49 = v49
                case -931:
                    v49 = "\n\t"
                case 127:
                    v5 = 0.1
                }
                var v52 bool = true
                switch ((v49 + v49) + v49) {
                case "":
                    v47 += 0
                }
                var v53 float32 = float32(3.14159)
                _ = fi51
                _ = v52
                _ = v53
            }
        case "foo bar":
            var v54 int8 = -128
            _ = v54
        case "hello":
            v1 = append(v1, 870)
        default:
            var v55 uint16 = 0
            _ = v55
        }
        _ = fi48
        _ = v49
        _ = v50
    }
    v56 := GenS11{F0: 0, F1: 800}
    v57 := GenS6{F0: true}
    v47 -= 0
    v5 -= -2.25
    v47 = (v47 << ((4294967295) & 63))
    if false {
        delete(v3, 2)
    } else {
        v58 := GenS10{F0: -9223372036854775808}
        for fi59 := 0; fi59 < 12; fi59++ {
            var v60 uint32 = 2147483647
            var v61 int = (fi59 & 32767)
            _ = fi59
            _ = v60
            _ = v61
        }
        fmt.Println(v47, v5, v5)
        _ = v58
    }
    var v62 string = "\n\t"
    v62 = v62
    var v63 uint32 = 2147483648
    if false {
        v64 := GenS12{F0: -560, F1: v47, F2: (^(v47))}
        _ = v64
    } else {
        for fi65 := 0; fi65 < 3; fi65++ {
            if false {
                v47 = 255
                v62 = v62
                var v66 uint = 416
                _ = v66
            } else {
                if false {
                    switch ((v62 + "a") + v62) {
                    case "\n\t":
                        if len(v1) > 0 {
                            v1[0] = 42
                        }
                    case "":
                        v5 -= float32(v2)
                    case "a":
                        var v67 int32 = 42
                        _ = v67
                    }
                    v62 = v62
                } else {
                    for fi68 := 0; fi68 < 20; fi68++ {
                        v63 += (v63 >> ((2) & 63))
                        if (!(!(false || false))) {
                            v69 := GenS3{F0: (v63 % (((^(v63))) | 1)), F1: 718, F2: v63}
                            var v70 uint64 = 1
                            _ = v69
                            _ = v70
                        }
                        _ = fi68
                    }
                    v5 += float32(v2)
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v71 := GenS12{F0: 42, F1: v47, F2: 2}
                _ = v71
            }
            _ = fi65
        }
        var v72 int64 = 9223372036854775807
        _ = v72
    }
    fmt.Println(v47)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v56
    _ = v57
    _ = v62
    _ = v63
}
