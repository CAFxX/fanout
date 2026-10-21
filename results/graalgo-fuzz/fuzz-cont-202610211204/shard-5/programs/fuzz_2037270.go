package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint8; F2 int64 }
type GenS7 struct { F0 uint8; F1 int; F2 int16 }
func genFn0(p0 string, p1 int16, p2 uint) int32 {
    switch "a" {
    case "foo bar":
        v7 := GenS1{F0: -2147483648}
        _ = v7
    case "hello":
        switch p0 {
        case "a":
            var v8 int = -966
            _ = v8
        case "\n\t":
            p0 = p0
        case "foo bar":
            p2 -= 1
        }
    }
    if true {
        var v9 float32 = -2.25
        if (true && true) {
            var v10 int64 = -74
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int16 = p1
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        _ = v9
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 bool = false
        p0 = p0
        if true {
            for fi15 := 0; fi15 < 19; fi15++ {
                var v16 uint8 = 157
                _ = fi15
                _ = v16
            }
            p0 = p0
        }
        _ = fi13
        _ = v14
    }
    var v17 int8 = 72
    for fi18 := 0; fi18 < 11; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := []int{fi18, (fi18 << (((p2 | p2)) & 63)), (fi18 % (((-(fi19))) | 1)), (fi19 - fi18)}
            v17 = v17
            if true {
                fmt.Println(fi19, fi18)
            }
            fmt.Println(fi18)
            fmt.Println(v17, v17, p1)
            switch (fi18 << (((p2 * (p2 % ((p2) | 1)))) & 63)) {
            case 300:
                fmt.Println(fi18)
            case 1:
                switch "foo bar" {
                case "\n\t":
                    var v21 float64 = 0.1
                    _ = v21
                case "foo bar":
                    v22 := GenS1{F0: 1}
                    _ = v22
                case "":
                    if (((fi19 & (fi19 / (((fi19 << (((p2 + p2)) & 63))) | 1))) < fi19) || (!(32767 < 42))) {
                        v17 = v17
                        v23 := GenS2{F0: 168}
                        if false {
                            switch (fi19 ^ fi19) {
                            case -2147483648:
                                if len(v20) > 0 {
                                    v20[0] = fi19
                                }
                            case 128:
                                var v24 int8 = v17
                                _ = v24
                            case 874:
                                switch fi19 {
                                case 489:
                                    v20 = append(v20, fi18)
                                case -1:
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                case 9223372036854775807:
                                    _ = len(v20)
                                default:
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                }
                            default:
                                switch fi19 {
                                case -9223372036854775808:
                                    v25 := GenS2{F0: 2147483648}
                                    _ = v25
                                case 32767:
                                    if false {
                                        if (255 >= p2) {
                                            v17 = v17
                                            if len(v20) > 0 {
                                                v20[0] = fi19
                                            }
                                            var v26 int16 = p1
                                            _ = v26
                                        } else {
                                            var v27 bool = (!(2 >= 116))
                                            _ = v27
                                        }
                                    } else {
                                        v28 := GenS3{F0: p2, F1: true, F2: 255}
                                        p1 += 1000
                                        _ = v28
                                    }
                                    fmt.Println(fi18, fi19, fi19)
                                default:
                                    var v29 int = (fi19 % (((fi18 & -1)) | 1))
                                    _ = v29
                                }
                            }
                            switch fi18 {
                            case 32767:
                                v17 += 0
                            case 884:
                                v17 = -128
                            case -9223372036854775808:
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                            default:
                                _ = len(v20)
                            }
                        }
                        _ = v23
                    }
                }
            case 127:
                p2 += (p2 << ((2) & 63))
            }
            p0 = (p0 + p0)
            _ = fi19
            _ = v20
        }
        var v30 uint64 = 483
        _ = fi18
        _ = v30
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 588
}
func genFn1() bool {
    for fi31 := 0; fi31 < 20; fi31++ {
        switch (fi31 * (fi31 - (fi31 >> ((255) & 63)))) {
        case -1:
            var v32 uint16 = 557
            _ = v32
        case -552:
            var v33 float32 = 0.0
            _ = v33
        case 2147483647:
            if (!false) {
                v34 := GenS4{F0: 4294967295, F1: 0}
                _ = v34
            }
        default:
            var v35 uint16 = 1
            _ = v35
        }
        _ = fi31
    }
    v36 := GenS4{F0: 4294967295, F1: 18446744073709551615}
    v37 := []int{-296, -77, -720}
    for fi38 := 0; fi38 < 6; fi38++ {
        switch "" {
        case "hello":
            var v39 int32 = 42
            _ = v39
        default:
            fmt.Println(fi38)
        }
        _ = fi38
    }
    _ = v36
    _ = v37
    return (0 < -1)
}

func main() {
    var v1 uint = 697
    var v2 uint8 = 0
    v3 := []string{""}
    var v4 int32 = 0
    var v5 int32 = 0
    var v6 int32 = 2147483647
    v40 := GenS5{F0: 185}
    if false {
        v3 = append(v3, "hello")
        v3 = append(v3, (("hello" + "hello") + (("\n\t" + "hello") + "foo bar")))
        fmt.Println(v4)
    }
    if (!(!false)) {
        for fi41 := 0; fi41 < 9; fi41++ {
            if false {
                var v42 bool = (v1 != 4294967295)
                _ = v42
            } else {
                var v43 uint8 = (v2 - uint8(v1))
                if false {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if ((false || (float32((-2.25 + (0.1 - 0.1))) == (0.0 / 1e10))) || false) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        if false {
                            switch (("" + ("\n\t" + ("hello" + ""))) + ("hello" + "\n\t")) {
                            case "a":
                                _ = len(v3)
                            }
                            for fi44 := 0; fi44 < 6; fi44++ {
                                if true {
                                    v6 += v5
                                }
                                v4 += genFn0(("\n\t" + ("\n\t" + "")), 1000, 1)
                                v3 = append(v3, ("a" + ("hello" + "foo bar")))
                                _ = fi44
                            }
                            if ("hello" <= ("hello" + ("\n\t" + ("foo bar" + (("" + (("a" + "a") + "foo bar")) + ("hello" + "a")))))) {
                                fmt.Println(v6)
                                if ((128 >= (fi41 * fi41)) || ((324 == 2) || true)) {
                                    var v45 int64 = 679
                                    fmt.Println(v2, fi41, fi41)
                                    _ = v45
                                }
                                v43 += v43
                            } else {
                                for fi46 := 0; fi46 < 7; fi46++ {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    _ = fi46
                                }
                                var v47 uint = (v1 / ((18446744073709551615) | 1))
                                _ = v47
                            }
                        }
                        switch fi41 {
                        case 9223372036854775807:
                            var v48 int32 = -1
                            _ = v48
                        case 32768:
                            v43 -= (v43 | (^(v2)))
                        case 292:
                            v43 = (v43 | v43)
                        default:
                            for fi49 := 0; fi49 < 17; fi49++ {
                                var v50 uint = (^(v1))
                                _ = fi49
                                _ = v50
                            }
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v51 uint = v1
                        v52 := GenS6{F0: 1, F1: 154, F2: 1}
                        _ = v51
                        _ = v52
                    }
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v2 = (v2 - v43)
                    _ = len(v3)
                }
                _ = v43
            }
            for fi53 := 0; fi53 < 13; fi53++ {
                fmt.Println(fi53)
                v3 = append(v3, ((("hello" + "foo bar") + "a") + ("" + "")))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi53
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi41
        }
        fmt.Println(v1)
        v1 = 416
    } else {
        var v54 uint32 = 423
        v1 += (v1 & 181)
        _ = v54
    }
    var v55 int16 = -1
    if (!(!false)) {
        switch int(v1) {
        case 42:
            v5 += (v4 * 1)
        case -1:
            v56 := GenS7{F0: v2, F1: 2147483647, F2: (v55 + v55)}
            _ = v56
        case -448:
            fmt.Println(v2, v6)
        }
    }
    v2 += v2
    v55 = v55
    if len(v3) > 0 {
        v3[0] = ""
    }
    var v57 int64 = 1
    var v58 uint16 = 715
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v55
    _ = v57
    _ = v58
    _ = v6
}
