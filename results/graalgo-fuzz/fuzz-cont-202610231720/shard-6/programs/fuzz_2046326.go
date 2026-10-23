package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 uint8 }
type GenS3 struct { F0 int16; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 int32; F1 uint32; F2 int }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 string }
func genFn0(p0 uint8, p1 uint16, p2 string) string {
    p1 = 2
    fmt.Println(p2, p0)
    var v5 uint = 577
    fmt.Println(p1, v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    var v1 float32 = (0.1 * 1e10)
    var v2 uint8 = 0
    v3 := make(map[int]uint32)
    var v4 bool = true
    var v6 uint8 = 225
    fmt.Println(v1)
    fmt.Println(v2, v4, v6)
    switch 895 {
    case 32767:
        v7 := GenS1{F0: 42, F1: 591, F2: 42}
        _ = v7
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS2{F0: 42, F1: 901, F2: v2}
            _ = fi8
            _ = v9
        }
    }
    v2 -= 42
    var v10 bool = v4
    if v10 {
        v11 := GenS3{F0: -32768, F1: 354, F2: 65535}
        for fi12 := 0; fi12 < 14; fi12++ {
            v6 -= (v2 + (^(v2)))
            for fi13 := 0; fi13 < 7; fi13++ {
                switch 128 {
                case 2147483647:
                    if (false || v10) {
                        v10 = (127 < 86)
                    } else {
                        fmt.Println(v10, v1, v6)
                        var v14 float64 = 3.14159
                        fmt.Println(v2)
                        _ = v14
                    }
                case 1:
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(v4, v2, v10)
                    case "foo bar":
                        var v15 int16 = 1000
                        _ = v15
                    default:
                        v16 := GenS2{F0: -1, F1: -400, F2: v2}
                        _ = v16
                    }
                default:
                    if v4 {
                        var v17 int32 = -731
                        v6 = (v6 - (v6 >> ((65535) & 63)))
                        _ = v17
                    }
                }
                var v18 float32 = v1
                _ = fi13
                _ = v18
            }
            var v19 int32 = 2147483647
            _ = fi12
            _ = v19
        }
        _ = v11
    }
    var v20 bool = (true && (708 != 0))
    v21 := GenS2{F0: 0, F1: -9223372036854775808, F2: (v6 >> ((18446744073709551615) & 63))}
    if v20 {
        for fi22 := 0; fi22 < 15; fi22++ {
            if v20 {
                for fi23 := 0; fi23 < 9; fi23++ {
                    if (0 == 42) {
                        var v24 int32 = -538
                        _ = v24
                    } else {
                        switch ("" + "\n\t") {
                        case "\n\t":
                            var v25 int8 = 100
                            _ = v25
                        case "":
                            var v26 uint16 = 42
                            _ = v26
                        case "a":
                            var v27 string = "foo bar"
                            _ = v27
                        default:
                            var v28 int16 = 753
                            _ = v28
                        }
                        v1 = float32(-2.25)
                    }
                    var v29 int16 = 1
                    _ = fi23
                    _ = v29
                }
                v30 := []int{}
                v31 := v3[4294967295]
                _ = v30
                _ = v31
            }
            _ = fi22
        }
        v2 -= 18
        v4 = (32767 < -767)
    } else {
        v32 := GenS3{F0: 89, F1: 413, F2: 621}
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        var v34 float32 = (v1 * ((v1 / (v1 * (v1 - v1))) + ((float32(3.14159) + 0.1) / v1)))
        v6 = v6
        if (1 < 2147483647) {
            for fi35 := 0; fi35 < 14; fi35++ {
                if ("" != ("a" + "foo bar")) {
                    v2 = (v2 * (v2 % (((v6 * v6)) | 1)))
                    v3[fi35] = 505
                    v36 := GenS2{F0: -1, F1: 825, F2: (v6 - v2)}
                    _ = v36
                } else {
                    var v37 int64 = -1
                    var v38 string = "\n\t"
                    var v39 uint8 = 251
                    _ = v37
                    _ = v38
                    _ = v39
                }
                _ = fi35
            }
            if v10 {
                v2 = (v2 % ((v6) | 1))
            }
            switch 4294967295 {
            case -1:
                if v10 {
                    v40 := GenS4{F0: 0, F1: -32768}
                    var v41 uint16 = 42
                    if v10 {
                        v10 = true
                        v2 -= v6
                        v42 := v3[(fi33 + (fi33 % ((fi33) | 1)))]
                        _ = v42
                    }
                    _ = v40
                    _ = v41
                }
            case 9223372036854775807:
                if v4 {
                    switch (fi33 | int(v6)) {
                    case 127:
                        v1 -= 3.14159
                    case 2147483647:
                        var v43 int16 = 523
                        _ = v43
                    case -9223372036854775808:
                        v44 := GenS5{F0: 42, F1: 2147483647, F2: (fi33 << ((2) & 63))}
                        _ = v44
                    default:
                        var v45 int32 = 2147483647
                        _ = v45
                    }
                    v4 = true
                    var v46 uint8 = v6
                    _ = v46
                }
            case 32767:
                var v47 int = (-(fi33))
                _ = v47
            default:
                v48 := GenS6{F0: 225, F1: 0}
                _ = v48
            }
        } else {
            fmt.Println(v4, v1)
            var v49 int8 = 124
            var v50 int32 = 595
            _ = v49
            _ = v50
        }
        _ = fi33
        _ = v34
    }
    fmt.Println(v6)
    for fi51 := 0; fi51 < 18; fi51++ {
        switch fi51 {
        case 127:
            var v52 int = (fi51 % ((4294967295) | 1))
            _ = v52
        case 555:
            for fi53 := 0; fi53 < 13; fi53++ {
                fmt.Println(v1)
                _ = fi53
            }
        case 32767:
            var v54 int64 = int64(fi51)
            _ = v54
        default:
            switch ("" + "") {
            case "foo bar":
                v4 = v4
            case "hello":
                fmt.Println(v20)
            case "\n\t":
                var v55 bool = true
                _ = v55
            }
        }
        for fi56 := 0; fi56 < 9; fi56++ {
            if (v10 && true) {
                fmt.Println(v6, v20)
            } else {
                v57 := []int{fi56, fi56}
                _ = v57
            }
            if false {
                v58 := []int{4294967295, (^(fi56)), (fi51 / (((fi51 - 2147483647)) | 1))}
                v59 := GenS7{F0: (("" + "foo bar") + "\n\t")}
                fmt.Println(v2, v20)
                _ = v58
                _ = v59
            }
            var v60 uint = 18446744073709551615
            _ = fi56
            _ = v60
        }
        v6 = (^(v6))
        _ = fi51
    }
    v61 := GenS6{F0: 65535, F1: 137}
    var v62 uint32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
    _ = v61
    _ = v62
}
