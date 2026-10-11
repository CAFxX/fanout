package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int32; F1 string }
type GenS6 struct { F0 uint8; F1 uint16; F2 string }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint }
func genFn0(p0 int32, p1 int8) uint {
    switch 1 {
    case 9223372036854775807:
        var v5 bool = (true || (32767 >= 822))
        _ = v5
    case -59:
        p0 = (p0 - p0)
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{"a", "\n\t", "foo bar", "hello"}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 13; fi6++ {
        if (false || (0.1 < 0.1)) {
            var v7 bool = false
            _ = v7
        } else {
            v8 := GenS1{F0: -572}
            var v9 uint = 42
            _ = v8
            _ = v9
        }
        if (-9223372036854775808 != 381) {
            switch fi6 {
            case 32767:
                var v10 uint32 = 568
                _ = v10
            case 42:
                var v11 int = (fi6 | -2147483648)
                _ = v11
            case 32768:
                v12 := GenS2{F0: false}
                _ = v12
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v13 int64 = 42
            v13 += int64(fi6)
            _ = v13
        }
        _ = fi6
    }
    v14 := GenS3{F0: 62, F1: 42}
    for fi15 := 0; fi15 < 10; fi15++ {
        v16 := GenS3{F0: 65535, F1: 1}
        for fi17 := 0; fi17 < 15; fi17++ {
            if true {
                v18 := GenS3{F0: 2147483647, F1: -32768}
                if false {
                    v1[42] = "hello"
                    if (!true) {
                        var v19 uint = 888
                        _ = v19
                    } else {
                        fmt.Println(fi17)
                        var v20 float32 = (3.14159 - 3.14159)
                        fmt.Println(fi15, v20)
                        _ = v20
                    }
                }
                fmt.Println(fi15, fi15)
                _ = v18
            }
            var v21 int8 = 42
            v1[176] = ""
            _ = fi17
            _ = v21
        }
        if len(v2) > 0 {
            v2[0] = ("" + "hello")
        }
        _ = fi15
        _ = v16
    }
    var v22 int16 = 713
    fmt.Println(v22)
    for fi23 := 0; fi23 < 18; fi23++ {
        if (false && ((^(fi23)) <= 128)) {
            var v24 uint64 = 4294967295
            for fi25 := 0; fi25 < 4; fi25++ {
                delete(v3, "foo bar")
                _ = fi25
            }
            switch ("" + (("" + "") + ("foo bar" + "hello"))) {
            case "foo bar":
                v24 -= 18446744073709551615
            case "hello":
                var v26 int = 9223372036854775807
                _ = v26
            case "\n\t":
                var v27 uint8 = 128
                _ = v27
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = v24
        } else {
            if (!false) {
                var v28 uint8 = 96
                fmt.Println(fi23, fi23, v28)
                v22 -= (v22 + v22)
                _ = v28
            }
            switch (fi23 % ((fi23) | 1)) {
            case 774:
                v22 += v22
            }
        }
        var v29 float64 = (float64(1.5) / 3.14159)
        _ = fi23
        _ = v29
    }
    switch 128 {
    case -308:
        switch ("\n\t" + "foo bar") {
        case "a":
            if false {
                switch 32767 {
                case 1:
                    if (979 >= (v22 << ((2) & 63))) {
                        var v30 float64 = 0.0
                        v22 -= (v22 + (v22 - (v22 * (v22 >> ((255) & 63)))))
                        _ = v30
                    } else {
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 uint32 = 65535
                            v33 := GenS4{F0: 9223372036854775807, F1: uint16(v32)}
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS5{F0: 2147483647, F1: "hello"}
                        if false {
                            for fi35 := 0; fi35 < 11; fi35++ {
                                switch fi35 {
                                case 4294967295:
                                    if (true || true) {
                                        v36 := GenS6{F0: 178, F1: 1, F2: "\n\t"}
                                        var v37 uint32 = 2
                                        _ = v36
                                        _ = v37
                                    }
                                case 127:
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                case 128:
                                    for fi38 := 0; fi38 < 1; fi38++ {
                                        var v39 bool = (!(!((true || true) && false)))
                                        v39 = v39
                                        var v40 int = 32768
                                        _ = fi38
                                        _ = v39
                                        _ = v40
                                    }
                                default:
                                    fmt.Println(v22, v22)
                                }
                                v22 = 1000
                                v41 := GenS7{F0: "", F1: ("foo bar" + "\n\t")}
                                _ = fi35
                                _ = v41
                            }
                            switch "a" {
                            case "hello":
                                v42 := v4[2147483647]
                                _ = v42
                            case "\n\t":
                                fmt.Println(v22)
                            default:
                                var v43 uint16 = 32767
                                _ = v43
                            }
                            var v44 float64 = -2.25
                            _ = v44
                        } else {
                            if false {
                                if false {
                                    var v45 int64 = 42
                                    if ((v22 - (-(v22))) <= v22) {
                                        v22 += (^(v22))
                                    } else {
                                        for fi46 := 0; fi46 < 5; fi46++ {
                                            var v47 int8 = int8(v45)
                                            _ = fi46
                                            _ = v47
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        v45 -= -572
                                    }
                                    v45 = v45
                                    _ = v45
                                }
                                v48 := GenS8{F0: 65535}
                                _ = v48
                            }
                        }
                        _ = v34
                    }
                }
            } else {
                v1[2147483647] = "hello"
                var v49 uint8 = 127
                _ = v49
            }
        }
    case 127:
        var v50 int16 = v22
        _ = v50
    case -684:
        switch "a" {
        case "foo bar":
            var v51 uint16 = 65535
            _ = v51
        }
    default:
        switch 0 {
        case 383:
            fmt.Println(v22)
        case 0:
            switch "" {
            case "\n\t":
                v22 = (-(v22))
            case "foo bar":
                var v52 int32 = 0
                _ = v52
            case "":
                v53 := GenS5{F0: -759, F1: "a"}
                _ = v53
            }
        }
    }
    if ((51 != 18446744073709551615) || true) {
        v22 += v22
        var v54 uint = 276
        var v55 float32 = (1.5 / (float32((float64(((float32((0.1 + 1e10)) / (1.5 / 1.0)) - (3.14159 - 0.1))) + 0.0)) - 1e10))
        _ = v54
        _ = v55
    }
    v22 = v22
    if (true || false) {
        var v56 uint64 = 2
        v22 -= 586
        _ = v56
    }
    fmt.Println(v22)
    if (1 < 2147483648) {
        if false {
            if (false && false) {
                v22 = (v22 * v22)
            }
        }
        var v57 uint32 = 194
        v58 := GenS4{F0: 9223372036854775807, F1: 255}
        _ = v57
        _ = v58
    } else {
        if true {
            var v59 uint8 = 143
            v60 := GenS7{F0: "a", F1: "foo bar"}
            _ = v59
            _ = v60
        } else {
            if false {
                v61 := GenS6{F0: 219, F1: 874, F2: (("foo bar" + ("foo bar" + "foo bar")) + "a")}
                v22 = (v22 + (v22 % (((v22 << ((1) & 63))) | 1)))
                v22 = v22
                _ = v61
            }
        }
    }
    fmt.Println(v22)
    switch 42 {
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println(v22)
    var v62 uint = 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v62:", v62)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v62
}
