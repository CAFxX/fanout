package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int32; F1 uint }
type GenS5 struct { F0 uint64; F1 int32; F2 int16 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{1, 42, -9223372036854775808}
    var v3 int32 = -47
    var v4 int32 = v3
    v5 := make(map[string]bool)
    fmt.Println(v3, v3)
    v4 += 358
    v4 = v4
    if ((620 >= 32767) && (-937 != int(v4))) {
        for fi6 := 0; fi6 < 18; fi6++ {
            v7 := GenS1{F0: 4294967295, F1: 2, F2: -728}
            if ((0.1 <= 0.0) && false) {
                if (3.14159 > ((float64(0.1) - -2.25) + (0.1 * (0.0 / -2.25)))) {
                    if len(v2) > 0 {
                        v2[0] = fi6
                    }
                    if (32767 <= 65535) {
                        if (false && ((v3 | (v3 / ((v3) | 1))) >= v4)) {
                            var v8 uint = 2
                            for fi9 := 0; fi9 < 4; fi9++ {
                                fmt.Println(v8, fi9, fi9)
                                var v10 uint32 = 138
                                var v11 uint8 = 250
                                _ = fi9
                                _ = v10
                                _ = v11
                            }
                            _ = len(v5)
                            _ = v8
                        }
                        v3 = (v3 % (((v4 << ((255) & 63))) | 1))
                        var v12 uint32 = 704
                        _ = v12
                    }
                }
                var v13 int16 = 0
                _ = v13
            } else {
                var v14 uint = 65535
                v14 -= (^(v14))
                _ = v14
            }
            v4 = -317
            _ = fi6
            _ = v7
        }
        var v15 float64 = (0.0 / 1.0)
        _ = v15
    }
    fmt.Println(v3, v3)
    if false {
        for fi16 := 0; fi16 < 7; fi16++ {
            if (!(1 != 2)) {
                var v17 string = "foo bar"
                for fi18 := 0; fi18 < 6; fi18++ {
                    if (1 >= 756) {
                        var v19 string = (v17 + (v17 + "a"))
                        switch v17 {
                        case "":
                            v20 := GenS2{F0: 1000, F1: (v4 ^ v4)}
                            _ = v20
                        }
                        _ = len(v5)
                        _ = v19
                    }
                    v17 = v17
                    _ = fi18
                }
                for fi21 := 0; fi21 < 20; fi21++ {
                    fmt.Println(v3)
                    var v22 uint64 = 559
                    _ = fi21
                    _ = v22
                }
                _ = v17
            }
            _ = fi16
        }
        fmt.Println(v4)
    } else {
        var v23 uint8 = 127
        v3 = v3
        _ = v23
    }
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 2; fi25++ {
        if (1 <= uint16(v24)) {
            fmt.Println(v3, v4)
        } else {
            v4 -= (v3 >> ((0) & 63))
        }
        v3 = (-(v3))
        switch (("hello" + ("hello" + "foo bar")) + ((("" + "a") + "a") + "")) {
        case "hello":
            fmt.Println(fi25, v4)
        }
        _ = fi25
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        switch 4294967295 {
        case -9223372036854775808:
            for fi27 := 0; fi27 < 3; fi27++ {
                v28 := GenS3{F0: 32768}
                switch "a" {
                case "\n\t":
                    var v29 uint16 = 65535
                    _ = v29
                case "foo bar":
                    if false {
                        _ = len(v2)
                        fmt.Println(v3, fi26, fi26)
                    }
                default:
                    if len(v2) > 0 {
                        v2[0] = fi27
                    }
                }
                var v30 uint8 = 63
                _ = fi27
                _ = v28
                _ = v30
            }
        default:
            if true {
                delete(v1, ("hello" + "foo bar"))
                v31 := GenS1{F0: 0, F1: 0, F2: -331}
                _ = v31
            } else {
                var v32 float32 = (float32(((float64((1e10 - 3.14159)) - -2.25) + 1e10)) + ((((0.1 / 1e10) + 0.0) + -2.25) + (((-2.25 / (((float32(float64((1.5 - 0.1))) - 3.14159) + (0.1 - -2.25)) + ((0.0 / 1e10) / float32(3.14159)))) * ((1e10 + 3.14159) * ((-2.25 + -2.25) * 0.1))) / -2.25)))
                for fi33 := 0; fi33 < 13; fi33++ {
                    v3 -= v4
                    fmt.Println(fi33)
                    _ = fi33
                }
                for fi34 := 0; fi34 < 20; fi34++ {
                    if false {
                        v35 := GenS4{F0: (v3 >> ((18446744073709551615) & 63)), F1: 2}
                        v24 -= 949
                        v24 += 18446744073709551615
                        var v36 uint32 = 4294967295
                        var v37 int8 = 0
                        _ = v35
                        _ = v36
                        _ = v37
                    } else {
                        var v38 int64 = 544
                        _ = len(v2)
                        var v39 int64 = v38
                        _ = v38
                        _ = v39
                    }
                    var v40 uint32 = 1
                    if true {
                        v4 += (v3 / (((v3 + 2147483647)) | 1))
                    }
                    _ = fi34
                    _ = v40
                }
                _ = v32
            }
        }
        v3 = v4
        v4 -= (v3 >> ((255) & 63))
        _ = fi26
    }
    if (1 < 0) {
        var v41 uint = 42
        switch 128 {
        case 32768:
            var v42 uint16 = 32767
            _ = v42
        case 32767:
            var v43 float64 = 3.14159
            _ = v43
        case 464:
            for fi44 := 0; fi44 < 4; fi44++ {
                if (0 < 65535) {
                    v3 = v4
                } else {
                    if true {
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                var v45 uint16 = 314
                                _ = v45
                            case "":
                                v24 += (v24 + v24)
                            case "foo bar":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v46 float32 = 3.14159
                                _ = v46
                            }
                        }
                        v24 = v24
                        var v47 bool = false
                        _ = v47
                    }
                }
                fmt.Println(v4)
                v48 := v1[""]
                _ = fi44
                _ = v48
            }
        }
        _ = v41
    } else {
        fmt.Println(v4, v3, v3)
    }
    for fi49 := 0; fi49 < 7; fi49++ {
        v50 := GenS5{F0: (v24 & 776), F1: -1, F2: 1}
        _ = fi49
        _ = v50
    }
    v24 += v24
    switch 4294967295 {
    case 42:
        if (127 < -1) {
            fmt.Println(v4, v3, v3)
            var v51 string = "hello"
            _ = v51
        }
    case 0:
        var v52 uint16 = 0
        _ = v52
    }
    fmt.Println(v3, v24, v3)
    v24 = (v24 >> ((255) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
