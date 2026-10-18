package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int64; F1 uint64; F2 int64 }
func genFn0(p0 int8, p1 uint8) bool {
    p1 += p1
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int16 = 1
        fmt.Println(p0, v8, v8)
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 475
    var v10 int32 = 0
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return false
}
func genFn1(p0 int16) int16 {
    v11 := []int{128, 32767, -98}
    var v12 float32 = (0.1 / float32(1.5))
    var v13 uint32 = 2
    fmt.Println(p0, v12, v12)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 1000
}

func main() {
    v1 := []int{32767, -2147483648, 42, -336}
    var v2 uint8 = 101
    v3 := make(map[int]int)
    var v4 int8 = 100
    var v5 bool = true
    var v6 bool = (v5 && false)
    fmt.Println(v5, v2)
    var v14 int64 = 9223372036854775807
    var v15 int8 = v4
    delete(v3, -2147483648)
    var v16 string = ("a" + (("\n\t" + "hello") + ("" + "\n\t")))
    var v17 int64 = v14
    if false {
        for fi18 := 0; fi18 < 13; fi18++ {
            if (!false) {
                _ = len(v1)
            }
            if (true || ((v15 > 126) && (((0 > fi18) || v5) && v6))) {
                for fi19 := 0; fi19 < 1; fi19++ {
                    var v20 int16 = 551
                    _ = fi19
                    _ = v20
                }
                if len(v1) > 0 {
                    v1[0] = fi18
                }
                v21 := GenS1{F0: (v5 || (586 <= 1)), F1: v14}
                _ = v21
            }
            v6 = v6
            _ = fi18
        }
        v6 = (255 != 18446744073709551615)
    } else {
        var v22 int16 = -380
        _ = v22
    }
    fmt.Println(v15)
    for fi23 := 0; fi23 < 8; fi23++ {
        for fi24 := 0; fi24 < 20; fi24++ {
            v16 = "hello"
            var v25 int32 = -2147483648
            if (1 <= 4294967295) {
                if true {
                    var v26 uint = 2
                    _ = v26
                } else {
                    switch (fi24 % (((fi23 + (fi24 * (fi24 * fi23)))) | 1)) {
                    case 1:
                        var v27 uint32 = 4294967295
                        _ = v27
                    case 42:
                        v25 -= (v25 >> ((0) & 63))
                    case 791:
                        var v28 uint16 = 867
                        _ = v28
                    }
                    v6 = (v5 && (v6 || true))
                }
                var v29 float32 = (1e10 - 3.14159)
                _ = v29
            } else {
                if ((v2 ^ v2) < 0) {
                    fmt.Println(v15, v14, v16)
                    fmt.Println(v5)
                    switch (fi24 % ((fi24) | 1)) {
                    case -623:
                        switch fi23 {
                        case -993:
                            if len(v1) > 0 {
                                v1[0] = fi23
                            }
                        default:
                            switch (fi24 | int(v2)) {
                            case 32768:
                                var v30 uint64 = 42
                                _ = v30
                            case 4294967295:
                                fmt.Println(v2, v5)
                            default:
                                v31 := GenS2{F0: (-(v25)), F1: 65535}
                                _ = v31
                            }
                        }
                    default:
                        v2 = (^(v2))
                        var v32 int8 = v4
                        v33 := GenS1{F0: false, F1: v17}
                        _ = v32
                        _ = v33
                    }
                    _ = len(v1)
                    fmt.Println(v14, v25)
                    _ = len(v3)
                } else {
                    var v34 int32 = 42
                    var v35 float64 = -2.25
                    v36 := GenS2{F0: v34, F1: 981}
                    if v6 {
                        v6 = (false && (-32768 >= -124))
                        var v37 bool = (v5 && ((v25 >> ((0) & 63)) != 317))
                        v34 = -422
                        var v38 uint8 = (v2 / (((v2 * v2)) | 1))
                        var v39 uint = 0
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    v6 = true
                    _ = v34
                    _ = v35
                    _ = v36
                }
                v15 += v4
            }
            if v6 {
                v2 = (v2 & (v2 << ((1) & 63)))
                v1 = append(v1, (fi24 >> ((2) & 63)))
            }
            _ = fi24
            _ = v25
        }
        if (v6 || (!(!false))) {
            switch fi23 {
            case 4294967295:
                fmt.Println(fi23, v6, v14)
            case 127:
                switch 1 {
                case 872:
                    fmt.Println(v17, v2, fi23)
                case -12:
                    switch (^(fi23)) {
                    case -1:
                        for fi40 := 0; fi40 < 14; fi40++ {
                            fmt.Println(v4, v16, v4)
                            var v41 int16 = genFn1(374)
                            var v42 uint32 = 1
                            fmt.Println(v16, v6, v15)
                            _ = fi40
                            _ = v41
                            _ = v42
                        }
                    case 9223372036854775807:
                        v43 := GenS3{F0: 1}
                        _ = v43
                    default:
                        fmt.Println(fi23)
                    }
                case -2147483648:
                    if len(v1) > 0 {
                        v1[0] = (fi23 | int(v14))
                    }
                }
            case 0:
                if ((v5 && (float32(-2.25) < 1.5)) || v6) {
                    var v44 bool = (!(1 > -32768))
                    _ = v44
                } else {
                    delete(v3, (^(fi23)))
                    for fi45 := 0; fi45 < 12; fi45++ {
                        var v46 uint16 = 255
                        var v47 uint = 255
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                }
            }
            v48 := GenS2{F0: -736, F1: uint16(v2)}
            v49 := GenS4{F0: 127, F1: v6}
            _ = v48
            _ = v49
        } else {
            v50 := GenS4{F0: (-(fi23)), F1: v6}
            v2 = (^(v2))
            v51 := GenS5{F0: -1, F1: 4294967295, F2: (v14 * v17)}
            _ = v50
            _ = v51
        }
        _ = fi23
    }
    _ = len(v1)
    v5 = ((!(!(!(813 >= 42)))) || v5)
    var v52 uint8 = 40
    delete(v3, 42)
    v5 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
}
