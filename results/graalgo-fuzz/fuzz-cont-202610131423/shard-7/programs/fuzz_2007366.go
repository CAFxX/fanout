package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 uint; F2 uint32 }

func main() {
    v1 := []int{4294967295, 42, 42}
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v4 := []uint32{483, 42, 2}
    var v5 int8 = -58
    var v6 string = "hello"
    switch 277 {
    case 2147483647:
        fmt.Println(v6)
    default:
        switch 128 {
        case -9223372036854775808:
            v5 -= (v5 / (((v5 & 1)) | 1))
        case -1:
            v5 = (-(v5))
        }
    }
    v3 = append(v3, v6)
    var v7 uint8 = 128
    if ((!true) || (-1 > 9223372036854775807)) {
        switch -431 {
        case 127:
            _ = len(v2)
        case 2147483647:
            for fi8 := 0; fi8 < 16; fi8++ {
                for fi9 := 0; fi9 < 4; fi9++ {
                    delete(v2, 42)
                    var v10 int16 = 0
                    v10 = (-(v10))
                    var v11 uint8 = 42
                    if false {
                        if false {
                            var v12 int64 = 906
                            switch (-(fi9)) {
                            case 32768:
                                v2[1] = 2147483648
                            case -9223372036854775808:
                                v4 = append(v4, 783)
                            }
                            var v13 uint64 = 4294967295
                            _ = v12
                            _ = v13
                        }
                    } else {
                        v14 := v2[2147483647]
                        var v15 int16 = 0
                        _ = v14
                        _ = v15
                    }
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                if (false || true) {
                    v7 -= 1
                    if true {
                        for fi16 := 0; fi16 < 9; fi16++ {
                            v17 := GenS1{F0: v7, F1: 9223372036854775807}
                            v7 = v7
                            var v18 bool = (!(!(!false)))
                            _ = fi16
                            _ = v17
                            _ = v18
                        }
                    }
                }
                for fi19 := 0; fi19 < 9; fi19++ {
                    if (-1 != -437) {
                        if (-2.25 > 1.5) {
                            if ((v7 << ((1) & 63)) >= 187) {
                                var v20 bool = (9223372036854775807 < -1)
                                fmt.Println(fi8)
                                _ = v20
                            }
                            switch (v6 + (((v6 + v6) + (v6 + v6)) + (v6 + ((v6 + v6) + ("a" + v6))))) {
                            case "foo bar":
                                v21 := GenS2{F0: 42, F1: 0, F2: 65535}
                                v2[2147483648] = 0
                                _ = v21
                            }
                            v22 := GenS3{F0: 0}
                            _ = v22
                        }
                    }
                    v7 -= v7
                    _ = fi19
                }
                _ = fi8
            }
        case -1:
            for fi23 := 0; fi23 < 5; fi23++ {
                v24 := GenS2{F0: 42, F1: 1, F2: 23}
                v7 -= 255
                _ = fi23
                _ = v24
            }
        }
        var v25 string = ("" + v6)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v25
    } else {
        var v26 int64 = 9223372036854775807
        delete(v2, 638)
        _ = v26
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := GenS2{F0: 2, F1: 0, F2: 255}
        var v29 bool = false
        v6 = ("hello" + "\n\t")
        _ = fi27
        _ = v28
        _ = v29
    }
    switch 127 {
    case -912:
        for fi30 := 0; fi30 < 8; fi30++ {
            if (65535 != 65535) {
                v31 := v2[417]
                if (false && (!(fi30 != (fi30 | (fi30 % (((-(fi30))) | 1)))))) {
                    for fi32 := 0; fi32 < 7; fi32++ {
                        v33 := GenS4{F0: 42}
                        var v34 int = (fi30 | (fi32 - (fi32 << ((255) & 63))))
                        v6 = (("a" + v6) + (v6 + v6))
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 15; fi35++ {
                        v7 -= v7
                        fmt.Println(v31, v6)
                        v31 = 0
                        if false {
                            var v36 uint = 255
                            v5 = v5
                            _ = v36
                        }
                        if (!true) {
                            fmt.Println(fi35, v5, v7)
                        } else {
                            v6 = "foo bar"
                        }
                        if true {
                            var v37 int32 = 0
                            var v38 string = v6
                            v39 := GenS5{F0: -620}
                            v2[(v31 / ((4294967295) | 1))] = uint32(v7)
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        if true {
                            if (223 >= 2) {
                                var v40 bool = (65535 == 65535)
                                _ = v40
                            } else {
                                fmt.Println(v5)
                            }
                            v7 = v7
                            if true {
                                v41 := GenS5{F0: (fi30 << ((65535) & 63))}
                                _ = v41
                            }
                            var v42 uint8 = 250
                            var v43 uint8 = (v7 % ((v42) | 1))
                            _ = v42
                            _ = v43
                        }
                        var v44 uint8 = (v7 + 0)
                        var v45 int8 = (^(v5))
                        var v46 int16 = 0
                        v47 := GenS6{F0: int32(v31), F1: 2, F2: 2147483647}
                        v5 = v45
                        if (v45 == v45) {
                            if len(v1) > 0 {
                                v1[0] = fi35
                            }
                            v48 := GenS3{F0: int8(v46)}
                            fmt.Println(v6, v44, v6)
                            _ = v48
                        } else {
                            v44 = (v7 ^ 216)
                        }
                        fmt.Println(v5)
                        _ = fi35
                        _ = v44
                        _ = v45
                        _ = v46
                        _ = v47
                    }
                    fmt.Println(v5, v31)
                } else {
                    v6 = v6
                    var v49 int8 = v5
                    var v50 uint32 = (v31 / ((539) | 1))
                    _ = v49
                    _ = v50
                }
                _ = v31
            } else {
                v2[65535] = 2147483648
                var v51 uint8 = v7
                _ = v51
            }
            _ = fi30
        }
    case 128:
        v7 = (v7 ^ v7)
    case -671:
        var v52 int16 = 1
        _ = v52
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v5, v5, v6)
    var v53 float64 = (-2.25 * (0.0 * 1.5))
    var v54 int = 602
    switch -578 {
    case 9223372036854775807:
        var v55 bool = (v6 <= v6)
        _ = v55
    case 4294967295:
        for fi56 := 0; fi56 < 14; fi56++ {
            v6 = "\n\t"
            _ = fi56
        }
    default:
        for fi57 := 0; fi57 < 19; fi57++ {
            v54 -= -179
            v53 = -2.25
            delete(v2, 2147483648)
            _ = fi57
        }
    }
    var v58 uint32 = 95
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v58
    _ = v6
    _ = v7
}
