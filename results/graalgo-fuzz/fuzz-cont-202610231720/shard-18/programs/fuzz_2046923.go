package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS4 struct { F0 int; F1 int8 }
type GenS5 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 int16) string {
    p0 = (-(p0))
    for fi7 := 0; fi7 < 4; fi7++ {
        p0 = p0
        v8 := GenS1{F0: -9223372036854775808}
        _ = fi7
        _ = v8
    }
    v9 := []int{9223372036854775807, 143, 0, 9223372036854775807}
    var v10 float64 = 0.0
    _ = p0
    _ = v10
    _ = v9
    return ((("foo bar" + ("hello" + ("\n\t" + (("" + "hello") + "")))) + ((("hello" + "") + ("\n\t" + "")) + "foo bar")) + ("" + "foo bar"))
}
func genFn1(p0 uint16, p1 int16, p2 bool) int32 {
    fmt.Println(p2)
    var v11 bool = p2
    v11 = p2
    p1 = p1
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return -30
}

func main() {
    var v1 uint = 255
    v2 := make(map[string]bool)
    var v3 int16 = 42
    v4 := []string{"\n\t", "hello"}
    var v5 string = "foo bar"
    v6 := []uint32{2147483648, 165, 400, 1}
    v12 := v2[((v5 + "hello") + v5)]
    for fi13 := 0; fi13 < 14; fi13++ {
        fmt.Println(v5, v1)
        _ = len(v2)
        _ = fi13
    }
    fmt.Println(v1, v5)
    switch (v5 + ("a" + v5)) {
    case "a":
        v3 = v3
    case "hello":
        v1 = (v1 % (((v1 % ((v1) | 1))) | 1))
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            if (281 < -180) {
                _ = len(v6)
                if v12 {
                    fmt.Println(v3)
                }
            }
            _ = fi15
        }
        fmt.Println(v12)
        v16 := GenS2{F0: 255, F1: -9223372036854775808, F2: 255}
        _ = fi14
        _ = v16
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 11; fi18++ {
            v1 = v1
            _ = fi18
        }
        for fi19 := 0; fi19 < 2; fi19++ {
            switch (v5 + (("" + (v5 + "")) + v5)) {
            case "\n\t":
                var v20 uint = 65535
                _ = v20
            case "a":
                var v21 int8 = 0
                _ = v21
            case "":
                var v22 uint64 = 4294967295
                _ = v22
            }
            v3 = -540
            _ = fi19
        }
        var v23 int16 = (-(v3))
        _ = fi17
        _ = v23
    }
    if true {
        var v24 uint = (v1 + v1)
        v3 = (v3 ^ v3)
        for fi25 := 0; fi25 < 15; fi25++ {
            fmt.Println(v24, v5, v5)
            fmt.Println(v3, v1, v5)
            _ = fi25
        }
        _ = v24
    }
    switch 128 {
    case 1:
        switch "foo bar" {
        case "foo bar":
            v26 := GenS1{F0: 0}
            _ = v26
        default:
            switch v5 {
            case "":
                if (-4 <= -128) {
                    v27 := GenS3{F0: 0, F1: 18446744073709551615, F2: v3}
                    v28 := v2["foo bar"]
                    _ = v27
                    _ = v28
                } else {
                    v12 = false
                }
            case "hello":
                v1 -= (^(v1))
            default:
                var v29 float64 = 1.5
                _ = v29
            }
        }
    case -9223372036854775808:
        for fi30 := 0; fi30 < 5; fi30++ {
            v12 = (-2.25 < 0.0)
            _ = fi30
        }
    }
    var v31 uint = (v1 | v1)
    var v32 int8 = 42
    if true {
        switch 32768 {
        case 1:
            fmt.Println(v31, v1)
        }
    } else {
        var v33 int = 127
        if v12 {
            v32 += (v32 * -12)
            var v34 string = (v5 + "")
            for fi35 := 0; fi35 < 6; fi35++ {
                v33 += fi35
                if ((((float64(1e10) + 0.1) / (((float64(((1.5 - (0.0 / 0.1)) + 1.5)) + 1e10) + 1.5) - 0.1)) - 1e10) == (((((3.14159 / 1.0) * 0.1) * 0.0) + float64(((0.0 + 3.14159) + float32(1e10)))) / (0.0 - -2.25))) {
                    var v36 float64 = float64(1e10)
                    var v37 int = v33
                    _ = v36
                    _ = v37
                }
                _ = fi35
            }
            _ = v34
        } else {
            if true {
                v38 := GenS4{F0: -463, F1: (v32 / ((42) | 1))}
                _ = v38
            }
            v31 = (v31 ^ 277)
        }
        var v39 uint = (^(v1))
        _ = v33
        _ = v39
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int8 = (^(v32))
        for fi42 := 0; fi42 < 17; fi42++ {
            v41 += (v32 >> (((v1 + v1)) & 63))
            _ = fi42
        }
        _ = fi40
        _ = v41
    }
    for fi43 := 0; fi43 < 8; fi43++ {
        for fi44 := 0; fi44 < 9; fi44++ {
            fmt.Println(v12, v12)
            _ = fi44
        }
        _ = fi43
    }
    v45 := v2[""]
    if (1 == 255) {
        v3 = (^(v3))
        v45 = v45
        var v46 uint16 = 32767
        _ = v46
    } else {
        v1 = v31
        if (!v45) {
            for fi47 := 0; fi47 < 5; fi47++ {
                var v48 bool = v45
                var v49 int32 = -62
                _ = fi47
                _ = v48
                _ = v49
            }
            switch (v5 + v5) {
            case "":
                var v50 uint = (v1 << ((v1) & 63))
                _ = v50
            case "\n\t":
                if v45 {
                    if v12 {
                        var v51 string = "a"
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 5; fi52++ {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        if (498 == 768) {
                            var v53 int32 = -267
                            var v54 string = ("foo bar" + "")
                            _ = v53
                            _ = v54
                        } else {
                            v31 -= v31
                            if ((2147483647 > 2) && (!(1 >= 0))) {
                                v5 = ""
                            }
                        }
                        var v55 bool = (!true)
                        _ = fi52
                        _ = v55
                    }
                    v2[v5] = v45
                } else {
                    switch "hello" {
                    case "\n\t":
                        for fi56 := 0; fi56 < 18; fi56++ {
                            var v57 float32 = ((((1.5 / ((-2.25 + 1.5) + 0.0)) * ((3.14159 + (1e10 / 1.5)) / (3.14159 - 0.0))) + 3.14159) + 3.14159)
                            v1 -= (^(v31))
                            _ = fi56
                            _ = v57
                        }
                    default:
                        switch 42 {
                        case 99:
                            var v58 int32 = -448
                            _ = v58
                        case 128:
                            for fi59 := 0; fi59 < 8; fi59++ {
                                switch "a" {
                                case "\n\t":
                                    var v60 bool = (1 > 1)
                                    _ = v60
                                case "foo bar":
                                    v61 := GenS1{F0: -1}
                                    _ = v61
                                }
                                _ = len(v2)
                                var v62 uint64 = 106
                                _ = fi59
                                _ = v62
                            }
                        default:
                            if false {
                                var v63 uint = (v1 | v31)
                                var v64 int = 0
                                _ = v63
                                _ = v64
                            }
                        }
                    }
                }
            default:
                if false {
                    for fi65 := 0; fi65 < 4; fi65++ {
                        v45 = true
                        for fi66 := 0; fi66 < 11; fi66++ {
                            var v67 int16 = (v3 * (v3 * (v3 >> ((v31) & 63))))
                            v1 += v1
                            if (true && ((4294967295 == 204) || (!v12))) {
                                fmt.Println(v31, fi66)
                            }
                            _ = fi66
                            _ = v67
                        }
                        v31 = v31
                        _ = fi65
                    }
                    for fi68 := 0; fi68 < 17; fi68++ {
                        v31 = v31
                        fmt.Println(v32, v45)
                        switch fi68 {
                        case 127:
                            v69 := GenS5{F0: (v32 << (((v31 + v31)) & 63)), F1: v5, F2: -258}
                            _ = v69
                        default:
                            for fi70 := 0; fi70 < 12; fi70++ {
                                switch v5 {
                                case "a":
                                    v6 = append(v6, 1)
                                }
                                _ = fi70
                            }
                        }
                        _ = fi68
                    }
                    v71 := GenS2{F0: 0, F1: 921, F2: 2}
                    _ = v71
                }
            }
            _ = len(v2)
        }
        switch (((v5 + v5) + "foo bar") + "hello") {
        case "a":
            fmt.Println(v5, v45)
        case "hello":
            fmt.Println(v3, v32)
        default:
            if (!(0 == -1)) {
                var v72 uint64 = 970
                _ = v72
            } else {
                var v73 int8 = (v32 % (((v32 << (((v31 - (v31 - v1))) & 63))) | 1))
                var v74 int = 0
                _ = v73
                _ = v74
            }
        }
    }
    var v75 int64 = 0
    var v76 int16 = v3
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v75:", v75)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v75
    _ = v76
}
