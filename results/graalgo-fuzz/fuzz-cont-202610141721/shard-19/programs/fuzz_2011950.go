package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint; F2 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16; F1 uint16 }
func genFn0(p0 int, p1 int32, p2 uint16) uint {
    var v7 uint64 = 70
    if true {
        v8 := GenS1{F0: 0}
        v9 := GenS1{F0: 255}
        var v10 uint32 = 4294967295
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]uint32)
    v3 := []uint32{4294967295, 2147483648, 4294967295}
    var v4 uint8 = 255
    var v5 uint = v1
    var v6 int = 0
    var v11 int = (v6 * v6)
    var v12 int64 = -9223372036854775808
    delete(v2, 0)
    v6 = -395
    if false {
        fmt.Println(v1, v4)
    }
    v1 = 780
    for fi13 := 0; fi13 < 17; fi13++ {
        if true {
            if true {
                var v14 float32 = (0.0 + (1.5 - 1.5))
                v6 -= (fi13 & 9223372036854775807)
                fmt.Println(v14, v5)
                _ = v14
            } else {
                v15 := GenS1{F0: (v4 ^ 42)}
                var v16 float32 = 1.5
                _ = v15
                _ = v16
            }
        }
        for fi17 := 0; fi17 < 3; fi17++ {
            switch "foo bar" {
            case "":
                v4 += v4
                fmt.Println(v4, v12)
                v11 = (^(fi13))
                switch fi17 {
                case -370:
                    _ = len(v2)
                case 0:
                    var v18 int32 = 0
                    _ = v18
                case 4294967295:
                    v6 -= fi17
                }
                v12 += (v12 | v12)
            case "\n\t":
                var v19 uint16 = 1
                _ = v19
            case "foo bar":
                _ = len(v3)
            default:
                var v20 int = (fi17 >> ((42) & 63))
                _ = v20
            }
            _ = fi17
        }
        var v21 int64 = (v12 / ((v12) | 1))
        _ = fi13
        _ = v21
    }
    if ((v4 * v4) > (v4 & (v4 * v4))) {
        _ = len(v3)
    }
    var v22 float64 = (0.1 / 1e10)
    for fi23 := 0; fi23 < 15; fi23++ {
        switch (v6 | 32767) {
        case 0:
            var v24 int = 682
            _ = v24
        case -414:
            v3 = append(v3, 65535)
        }
        switch ("\n\t" + "") {
        case "":
            for fi25 := 0; fi25 < 17; fi25++ {
                v6 -= (fi25 - -9223372036854775808)
                var v26 uint32 = 608
                var v27 int8 = 42
                switch 32768 {
                case -256:
                    var v28 int = fi23
                    if ((v12 % (((v12 * v12)) | 1)) < 9223372036854775807) {
                        fmt.Println(v4, v4)
                        fmt.Println(v4)
                        if false {
                            if (v22 < -2.25) {
                                var v29 int32 = -1
                                _ = v29
                            }
                            var v30 string = "hello"
                            v31 := GenS2{F0: uint16(v26), F1: 2, F2: v4}
                            _ = v30
                            _ = v31
                        } else {
                            switch "hello" {
                            case "a":
                                if true {
                                    v28 = 951
                                    if true {
                                        v32 := GenS1{F0: (v4 + 255)}
                                        _ = v32
                                    } else {
                                        switch 120 {
                                        case 950:
                                            v33 := GenS3{F0: ""}
                                            _ = v33
                                        case 128:
                                            switch (("foo bar" + ("\n\t" + "hello")) + "\n\t") {
                                            case "foo bar":
                                                v34 := GenS2{F0: 65535, F1: genFn0((v28 / ((v28) | 1)), 42, 32767), F2: v4}
                                                _ = v34
                                            default:
                                                var v35 uint16 = 32767
                                                _ = v35
                                            }
                                        default:
                                            delete(v2, 32767)
                                        }
                                        var v36 int64 = 261
                                        _ = v36
                                    }
                                    var v37 float64 = v22
                                    _ = v37
                                } else {
                                    v1 = 613
                                    v1 = (v1 << (((v1 & (^(v1)))) & 63))
                                }
                            default:
                                v2[fi25] = v26
                            }
                        }
                    }
                    v3 = append(v3, v26)
                    v38 := v2[v28]
                    var v39 float32 = 1.5
                    _ = v28
                    _ = v38
                    _ = v39
                case 721:
                    v40 := GenS3{F0: ("\n\t" + "hello")}
                    _ = v40
                }
                if true {
                    var v41 int32 = int32(v26)
                    v26 = (v26 + 2147483648)
                    _ = v41
                }
                _ = fi25
                _ = v26
                _ = v27
            }
        case "a":
            for fi42 := 0; fi42 < 17; fi42++ {
                v22 -= 0.1
                var v43 string = "hello"
                if false {
                    if false {
                        v44 := GenS1{F0: 128}
                        v22 = 1e10
                        delete(v2, (fi42 ^ v11))
                        _ = v44
                    } else {
                        if (!(v11 != 4294967295)) {
                            v45 := GenS1{F0: v4}
                            var v46 string = ""
                            _ = v45
                            _ = v46
                        } else {
                            switch v43 {
                            case "foo bar":
                                v47 := v2[(v11 << ((genFn0((^(v6)), 2147483647, 591)) & 63))]
                                _ = v47
                            default:
                                v48 := GenS1{F0: (v4 >> ((v5) & 63))}
                                _ = v48
                            }
                        }
                    }
                    fmt.Println(v6)
                    v2[fi23] = 65535
                    v22 += float64((-2.25 + (((0.0 + ((0.1 / -2.25) + 1e10)) + ((0.1 - (1.5 + -2.25)) + 3.14159)) - 0.0)))
                } else {
                    _ = len(v2)
                    v6 -= v11
                    var v49 string = "foo bar"
                    _ = v49
                }
                _ = fi42
                _ = v43
            }
        case "\n\t":
            for fi50 := 0; fi50 < 16; fi50++ {
                v51 := GenS3{F0: "foo bar"}
                _ = fi50
                _ = v51
            }
        }
        _ = fi23
    }
    switch v6 {
    case 405:
        for fi52 := 0; fi52 < 8; fi52++ {
            switch (fi52 | 2147483647) {
            case -220:
                delete(v2, fi52)
            case 0:
                switch 1 {
                case 32768:
                    if true {
                        var v53 uint8 = v4
                        v54 := GenS4{F0: 42, F1: 2}
                        _ = v53
                        _ = v54
                    }
                case 42:
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                case 127:
                    switch "hello" {
                    case "a":
                        var v55 int8 = 59
                        _ = v55
                    default:
                        v11 += (v11 | (fi52 >> ((1) & 63)))
                    }
                }
            default:
                switch fi52 {
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case 1:
                    var v56 int64 = (v12 >> ((v5) & 63))
                    _ = v56
                }
            }
            _ = fi52
        }
    case -306:
        var v57 int16 = 118
        _ = v57
    }
    if (((1.5 - 0.1) + float32(((3.14159 / (v22 / (0.0 - (1.5 * v22)))) - 3.14159))) > (0.0 - (0.1 - ((0.0 + (0.0 / (0.1 / (float32(float64(((3.14159 * 0.1) * 0.1))) + (float32(1e10) / float32(v22)))))) + 1.5)))) {
        v1 += uint(v4)
        var v58 float32 = ((float32(0.0) + -2.25) - -2.25)
        _ = v58
    } else {
        if (3.14159 >= 3.14159) {
            switch (^(v6)) {
            case 127:
                v5 = v1
            default:
                var v59 int8 = 0
                _ = v59
            }
            var v60 uint64 = 683
            var v61 uint16 = 710
            _ = v60
            _ = v61
        }
        var v62 string = ""
        _ = v62
    }
    fmt.Println(v6, v12, v6)
    switch (("foo bar" + (("" + ("" + "\n\t")) + "\n\t")) + "\n\t") {
    case "foo bar":
        v1 += v1
    case "":
        for fi63 := 0; fi63 < 6; fi63++ {
            v5 += (v1 >> ((v1) & 63))
            _ = fi63
        }
    case "\n\t":
        v5 += genFn0(317, -2147483648, 42)
    default:
        var v64 float32 = ((3.14159 + 0.1) / -2.25)
        _ = v64
    }
    for fi65 := 0; fi65 < 12; fi65++ {
        v11 = fi65
        _ = fi65
    }
    fmt.Println(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
