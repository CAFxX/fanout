package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint16 }
type GenS4 struct { F0 uint; F1 string; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 bool; F2 uint32 }
type GenS7 struct { F0 uint64; F1 uint64 }
type GenS8 struct { F0 int64; F1 int32; F2 int64 }
type GenS9 struct { F0 uint; F1 uint }
type GenS10 struct { F0 int; F1 uint16 }
type GenS11 struct { F0 string }
type GenS12 struct { F0 int32; F1 string; F2 uint8 }
func genFn0() int8 {
    var v6 string = "hello"
    var v7 uint32 = 592
    v6 = (("" + v6) + "hello")
    switch 57 {
    case 9223372036854775807:
        var v8 uint8 = uint8(v7)
        _ = v8
    }
    _ = v6
    _ = v7
    return 90
}
func genFn1(p0 int8, p1 int32, p2 uint32) uint {
    switch 68 {
    case 2147483647:
        var v9 int = -9223372036854775808
        _ = v9
    default:
        if (!(uint8(p2) == 192)) {
            v10 := []int{4294967295, -694, int(p1), int(p2)}
            var v11 bool = true
            _ = v10
            _ = v11
        }
    }
    v12 := GenS1{F0: 567}
    v13 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 0
}

func main() {
    var v1 float32 = 1.5
    v2 := make(map[int]uint32)
    v3 := make(map[int]string)
    var v4 uint32 = 0
    var v5 bool = (true && false)
    v5 = false
    switch -872 {
    case -9223372036854775808:
        v14 := []int{79}
        _ = v14
    case 0:
        for fi15 := 0; fi15 < 19; fi15++ {
            if (255 > 1) {
                fmt.Println(v1)
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (9223372036854775807 >= -341) {
                        if ((!true) || v5) {
                            v5 = ((0 != -1) || v5)
                        }
                    }
                    fmt.Println(v4)
                    _ = fi16
                }
                v17 := []int{1, fi15, fi15, (fi15 >> ((271) & 63))}
                _ = v17
            } else {
                v1 = 0.1
            }
            var v18 bool = v5
            v19 := v3[(fi15 & (fi15 * (fi15 | fi15)))]
            _ = fi15
            _ = v18
            _ = v19
        }
    default:
        if (1e10 > 1e10) {
            for fi20 := 0; fi20 < 2; fi20++ {
                if ("" >= "hello") {
                    switch (fi20 ^ (fi20 - fi20)) {
                    case 2147483647:
                        v4 += (v4 << ((18446744073709551615) & 63))
                    case -9223372036854775808:
                        v21 := GenS1{F0: 1}
                        _ = v21
                    }
                    v22 := GenS2{F0: uint(v4)}
                    v5 = v5
                    _ = v22
                } else {
                    fmt.Println(v4)
                }
                v5 = (32767 > 32767)
                v5 = (v5 || v5)
                _ = fi20
            }
            if (252 > 1) {
                if false {
                    fmt.Println(v1)
                    var v23 uint32 = (v4 ^ (^(v4)))
                    _ = v23
                } else {
                    v24 := GenS3{F0: -77, F1: 2, F2: 739}
                    var v25 bool = (true && v5)
                    _ = v24
                    _ = v25
                }
                if v5 {
                    switch 32768 {
                    case -97:
                        switch "hello" {
                        case "a":
                            v4 -= v4
                        case "foo bar":
                            var v26 uint32 = v4
                            _ = v26
                        default:
                            delete(v3, 32767)
                        }
                    }
                    v5 = v5
                    v4 = (v4 << ((51) & 63))
                } else {
                    _ = len(v3)
                    var v27 int16 = -32768
                    _ = v27
                }
                for fi28 := 0; fi28 < 20; fi28++ {
                    var v29 int8 = 100
                    v30 := GenS4{F0: genFn1(100, 1, (v4 - (v4 >> ((2) & 63)))), F1: ("" + ("a" + "hello")), F2: 52}
                    if (42 != 1) {
                        fmt.Println(v1, v29)
                        v31 := GenS5{F0: 101}
                        v32 := GenS6{F0: 1, F1: v5, F2: 0}
                        _ = v31
                        _ = v32
                    }
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
            if false {
                if (v5 && v5) {
                    v1 += (v1 * 3.14159)
                    switch 456 {
                    case 1:
                        v1 = v1
                    case 808:
                        v4 += v4
                    }
                    v33 := GenS3{F0: 127, F1: 0, F2: 329}
                    _ = v33
                }
                v34 := []int{42, -1, 127, -1}
                if (v5 || ((!(128 < 94)) && true)) {
                    for fi35 := 0; fi35 < 14; fi35++ {
                        v36 := GenS1{F0: 0}
                        fmt.Println(v1, v1, fi35)
                        if true {
                            v1 += v1
                            var v37 int32 = -858
                            v38 := GenS7{F0: 677, F1: 2}
                            _ = v37
                            _ = v38
                        } else {
                            var v39 int16 = 0
                            fmt.Println(v1)
                            _ = v39
                        }
                        _ = fi35
                        _ = v36
                    }
                    for fi40 := 0; fi40 < 6; fi40++ {
                        v34 = append(v34, (fi40 - (fi40 * (fi40 + -9223372036854775808))))
                        if (0 == 32767) {
                            v41 := GenS8{F0: -828, F1: 774, F2: 0}
                            _ = v41
                        }
                        _ = fi40
                    }
                }
                _ = v34
            } else {
                v42 := GenS4{F0: 4294967295, F1: ((("a" + "") + "\n\t") + "a"), F2: 165}
                _ = v42
            }
        } else {
            for fi43 := 0; fi43 < 20; fi43++ {
                v44 := GenS4{F0: 4294967295, F1: "a", F2: 127}
                if ("foo bar" == (("hello" + ("hello" + "")) + "foo bar")) {
                    var v45 uint32 = (v4 ^ v4)
                    _ = v45
                } else {
                    var v46 uint8 = 147
                    v47 := GenS6{F0: -9223372036854775808, F1: ("\n\t" == (("foo bar" + "foo bar") + "\n\t")), F2: v4}
                    _ = v46
                    _ = v47
                }
                delete(v3, fi43)
                _ = fi43
                _ = v44
            }
            var v48 uint64 = 1
            v4 = 0
            _ = v48
        }
    }
    for fi49 := 0; fi49 < 6; fi49++ {
        switch (-(fi49)) {
        case -949:
            v50 := GenS6{F0: int64(v4), F1: v5, F2: v4}
            _ = v50
        }
        _ = fi49
    }
    v4 += 46
    for fi51 := 0; fi51 < 17; fi51++ {
        v5 = (!("" >= "foo bar"))
        _ = fi51
    }
    v52 := GenS9{F0: 61, F1: 1}
    v1 += v1
    if (!v5) {
        fmt.Println(v4)
        v53 := GenS5{F0: 42}
        v5 = (18446744073709551615 <= 42)
        _ = v53
    }
    v54 := []int{1}
    if (0 != 4294967295) {
        for fi55 := 0; fi55 < 6; fi55++ {
            var v56 float32 = (float32(((0.1 / float64(v1)) - (0.1 - 3.14159))) - v1)
            var v57 uint8 = 42
            _ = fi55
            _ = v56
            _ = v57
        }
        v4 -= v4
        if (v5 || true) {
            v58 := GenS9{F0: 1, F1: genFn1(-74, 0, (v4 >> ((664) & 63)))}
            v59 := GenS3{F0: -1, F1: 255, F2: 65535}
            _ = v58
            _ = v59
        } else {
            fmt.Println(v1)
            var v60 float32 = float32(((1e10 - 3.14159) + -2.25))
            if v5 {
                if (42 != -262) {
                    v61 := GenS10{F0: -2147483648, F1: 344}
                    _ = v61
                } else {
                    var v62 uint32 = v4
                    _ = v62
                }
                if false {
                    v5 = (!true)
                    v63 := GenS10{F0: 9223372036854775807, F1: 32767}
                    _ = v63
                } else {
                    fmt.Println(v5)
                    v64 := GenS11{F0: ("\n\t" + "\n\t")}
                    _ = v64
                }
                fmt.Println(v4)
            } else {
                for fi65 := 0; fi65 < 11; fi65++ {
                    var v66 uint = 0
                    var v67 int16 = 32767
                    v68 := GenS10{F0: (-(fi65)), F1: 255}
                    _ = fi65
                    _ = v66
                    _ = v67
                    _ = v68
                }
                fmt.Println(v5, v5, v60)
                for fi69 := 0; fi69 < 9; fi69++ {
                    if v5 {
                        v60 = v1
                    }
                    _ = fi69
                }
            }
            _ = v60
        }
    } else {
        if v5 {
            var v70 uint = 18446744073709551615
            v71 := GenS12{F0: 2147483647, F1: "", F2: 128}
            v1 -= v1
            _ = v70
            _ = v71
        }
    }
    for fi72 := 0; fi72 < 15; fi72++ {
        if len(v54) >= 2 {
            v54 = v54[1:len(v54)]
        }
        var v73 uint64 = 18446744073709551615
        _ = fi72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v54:", len(v54), v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v54
}
