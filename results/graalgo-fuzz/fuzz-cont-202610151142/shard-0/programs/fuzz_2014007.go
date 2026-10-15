package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int32 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int64; F2 string }
type GenS5 struct { F0 bool }
func genFn0(p0 uint8, p1 int64, p2 uint64) int {
    p2 = p2
    if (!true) {
        fmt.Println(p0, p2)
        v7 := GenS1{F0: 1000, F1: 2147483647}
        for fi8 := 0; fi8 < 4; fi8++ {
            for fi9 := 0; fi9 < 12; fi9++ {
                var v10 int32 = 180
                v11 := GenS2{F0: uint32(p0), F1: 0, F2: int8(p2)}
                _ = fi9
                _ = v10
                _ = v11
            }
            var v12 int = 127
            _ = fi8
            _ = v12
        }
        _ = v7
    }
    v13 := GenS2{F0: 995, F1: 837, F2: -128}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 127
}
func genFn1(p0 int16, p1 int64) uint64 {
    for fi14 := 0; fi14 < 15; fi14++ {
        p1 = 896
        switch fi14 {
        case -1:
            var v15 float64 = 0.0
            _ = v15
        case 782:
            if (2147483647 == 65535) {
                switch (fi14 * (fi14 << ((451) & 63))) {
                case 2147483647:
                    v16 := []int{(fi14 | fi14), fi14}
                    _ = v16
                case -540:
                    var v17 uint32 = 1
                    _ = v17
                }
            }
        }
        if false {
            v18 := GenS3{F0: 121, F1: 18446744073709551615}
            _ = v18
        }
        _ = fi14
    }
    var v19 int8 = -128
    var v20 int16 = (p0 + (p0 / (((p0 | (p0 ^ 93))) | 1)))
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    return 900
}

func main() {
    var v1 int8 = 118
    var v2 uint32 = 4294967295
    v3 := []int{128, 9223372036854775807}
    var v4 int32 = 1
    v5 := make(map[int]bool)
    var v6 int64 = 9223372036854775807
    var v21 string = "a"
    if true {
        var v22 int64 = v6
        _ = v22
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            switch v21 {
            case "a":
                switch v21 {
                case "foo bar":
                    if (true || false) {
                        if (!(2 < 65535)) {
                            var v24 uint16 = 119
                            for fi25 := 0; fi25 < 19; fi25++ {
                                switch v21 {
                                case "":
                                    switch "" {
                                    case "a":
                                        v6 = v6
                                    case "\n\t":
                                        switch v21 {
                                        case "\n\t":
                                            fmt.Println(fi25)
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = 517
                                            }
                                        default:
                                            v4 = (v4 ^ v4)
                                        }
                                        if len(v3) > 0 {
                                            v3[0] = (fi25 | (fi23 & (-(fi25))))
                                        }
                                    default:
                                        v26 := GenS2{F0: (^(v2)), F1: (v24 * v24), F2: (-(v1))}
                                        _ = v26
                                    }
                                case "\n\t":
                                    v4 += v4
                                case "hello":
                                    if false {
                                        v3 = append(v3, (fi23 & fi25))
                                        v4 -= 1
                                    } else {
                                        v27 := GenS2{F0: v2, F1: v24, F2: (v1 * v1)}
                                        v1 = (v1 | v1)
                                        _ = v27
                                    }
                                }
                                if true {
                                    var v28 float32 = (0.0 * 0.1)
                                    var v29 uint8 = 179
                                    _ = v28
                                    _ = v29
                                }
                                var v30 int = fi25
                                _ = fi25
                                _ = v30
                            }
                            var v31 int32 = (^(v4))
                            _ = v24
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 7; fi32++ {
                                switch genFn0(255, v6, 0) {
                                case 906:
                                    var v33 float32 = 3.14159
                                    _ = v33
                                default:
                                    v2 = v2
                                }
                                var v34 int = fi32
                                fmt.Println(fi23, fi32)
                                _ = fi32
                                _ = v34
                            }
                        }
                        switch (fi23 & (fi23 ^ (fi23 % ((fi23) | 1)))) {
                        case -905:
                            var v35 uint32 = v2
                            _ = v35
                        case -324:
                            fmt.Println(fi23, v4)
                        case 2147483647:
                            v21 = ("a" + v21)
                        }
                        for fi36 := 0; fi36 < 12; fi36++ {
                            v21 = ""
                            _ = fi36
                        }
                    }
                case "a":
                    var v37 int32 = v4
                    _ = v37
                }
            case "":
                switch (fi23 * genFn0(0, -1, 4294967295)) {
                case 0:
                    var v38 int32 = (v4 & 426)
                    _ = v38
                case 42:
                    v21 = v21
                case 874:
                    fmt.Println(v2)
                }
            case "foo bar":
                v21 = "\n\t"
            default:
                v1 = v1
            }
            _ = fi23
        }
    }
    v4 = v4
    var v39 uint32 = v2
    if false {
        _ = len(v3)
    } else {
        v2 -= v39
    }
    if false {
        var v40 int32 = 42
        _ = v40
    } else {
        var v41 uint32 = (v39 << ((606) & 63))
        switch ("foo bar" + ((v21 + "a") + (v21 + ("\n\t" + v21)))) {
        case "a":
            var v42 float32 = 1e10
            _ = v42
        default:
            if false {
                var v43 int32 = 2147483647
                _ = v43
            }
        }
        _ = v41
    }
    if (v1 >= (^(v1))) {
        var v44 uint8 = 127
        v44 += 127
        _ = v44
    }
    v4 = v4
    _ = len(v5)
    v45 := GenS3{F0: (v1 / ((v1) | 1)), F1: 2}
    var v46 int32 = v4
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 uint64 = 4294967295
        v49 := GenS4{F0: (v39 + 65535), F1: (v6 ^ 0), F2: v21}
        _ = fi47
        _ = v48
        _ = v49
    }
    if (true && ("\n\t" != (v21 + (("\n\t" + ((v21 + v21) + v21)) + (((v21 + v21) + "foo bar") + "hello"))))) {
        if (!(!(1 <= 4294967295))) {
            switch v21 {
            case "hello":
                delete(v5, genFn0(79, 1, 18446744073709551615))
            default:
                v50 := GenS5{F0: (-735 >= 42)}
                _ = v50
            }
            v51 := GenS3{F0: v1, F1: genFn1(0, (v6 - 9223372036854775807))}
            _ = v51
        } else {
            for fi52 := 0; fi52 < 8; fi52++ {
                fmt.Println(v6, v21, v21)
                var v53 uint = 255
                _ = fi52
                _ = v53
            }
        }
    }
    var v54 int = 172
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v54
    _ = v6
}
