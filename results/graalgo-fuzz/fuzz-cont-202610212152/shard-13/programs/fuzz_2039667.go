package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int16; F1 int8; F2 uint32 }
type GenS3 struct { F0 uint8; F1 int16 }
type GenS4 struct { F0 int16; F1 bool; F2 int }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 string; F1 uint64; F2 uint32 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]string)
    var v3 bool = (false && (true || true))
    var v4 uint = 42
    var v5 uint8 = 1
    if (uint64(v5) > 18446744073709551615) {
        switch "hello" {
        case "":
            fmt.Println(v3, v3)
        case "foo bar":
            _ = len(v2)
        case "hello":
            v4 += 106
        default:
            if (!(2147483647 == 2147483647)) {
                fmt.Println(v4)
                for fi6 := 0; fi6 < 2; fi6++ {
                    for fi7 := 0; fi7 < 17; fi7++ {
                        if (!(0.0 == 3.14159)) {
                            if false {
                                var v8 int64 = 0
                                v9 := []int{}
                                _ = v8
                                _ = v9
                            }
                            var v10 int64 = 0
                            _ = v10
                        } else {
                            var v11 int8 = 42
                            var v12 uint8 = 1
                            _ = v11
                            _ = v12
                        }
                        v13 := GenS1{F0: (v5 + (v5 ^ v5)), F1: int(v5)}
                        var v14 float64 = ((0.1 + 0.0) + (1e10 - 1e10))
                        var v15 int = fi7
                        _ = fi7
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                    _ = fi6
                }
            }
        }
        v16 := []int{int(v5), 347, 9223372036854775807}
        var v17 uint = v4
        _ = v16
        _ = v17
    } else {
        fmt.Println(v5)
        switch 504 {
        case 388:
            var v18 float64 = 1.5
            _ = v18
        }
    }
    v4 = (v4 | (v4 / (((v4 / (((v4 + 0)) | 1))) | 1)))
    switch ("hello" + "foo bar") {
    case "a":
        if (-9223372036854775808 < 1) {
            if ("" > "foo bar") {
                switch "a" {
                case "foo bar":
                    v5 -= 0
                case "\n\t":
                    for fi19 := 0; fi19 < 2; fi19++ {
                        var v20 float64 = (((float64((0.0 - -2.25)) + 0.1) + 1e10) - 1.5)
                        switch "\n\t" {
                        case "":
                            if true {
                                var v21 int16 = 138
                                v5 += 2
                                v5 -= (^(v5))
                                _ = v21
                            } else {
                                var v22 float64 = 0.1
                                _ = v22
                            }
                        case "a":
                            v23 := GenS1{F0: 121, F1: fi19}
                            _ = v23
                        case "hello":
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v20 = v20
                                v25 := GenS1{F0: v5, F1: fi19}
                                _ = fi24
                                _ = v25
                            }
                        default:
                            switch fi19 {
                            case 2147483647:
                                var v26 int64 = -9223372036854775808
                                _ = v26
                            }
                        }
                        _ = len(v2)
                        _ = fi19
                        _ = v20
                    }
                case "hello":
                    v5 -= v5
                }
            } else {
                v4 = 0
            }
        } else {
            if v3 {
                if v3 {
                    v1 -= 1
                    v27 := GenS2{F0: 0, F1: 126, F2: (v1 | v1)}
                    _ = v27
                } else {
                    for fi28 := 0; fi28 < 3; fi28++ {
                        if (0 <= 2147483647) {
                            v1 -= (v1 / ((v1) | 1))
                            v29 := GenS1{F0: (v5 - (^(v5))), F1: 1}
                            _ = v29
                        }
                        if (42 >= 0) {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                if v3 {
                                    v1 = uint32(v5)
                                    var v31 float64 = (((((0.0 - float64(float32(((0.0 - 3.14159) + 1.5)))) / (1.5 / (1.5 + 1e10))) / 3.14159) * float64((1e10 / (1e10 / (float32(3.14159) - (((0.0 - float32((1e10 + 0.0))) + 1.5) - 0.0)))))) / (3.14159 / (0.0 + float64(0.1))))
                                    _ = v31
                                } else {
                                    if (!v3) {
                                        switch "a" {
                                        case "\n\t":
                                            v2[fi28] = "hello"
                                        case "foo bar":
                                            v4 -= (v4 >> (((v4 / ((v4) | 1))) & 63))
                                            v3 = (!(!false))
                                            v2[(fi28 - 2147483647)] = "hello"
                                            v32 := []int{(fi30 | fi28)}
                                            fmt.Println(fi30, fi28)
                                            _ = v32
                                        case "a":
                                            var v33 uint16 = uint16(v5)
                                            _ = v33
                                        default:
                                            fmt.Println(fi30, fi30, fi28)
                                        }
                                        switch ("foo bar" + "hello") {
                                        case "\n\t":
                                            var v34 string = "a"
                                            _ = v34
                                        }
                                        var v35 uint64 = 18446744073709551615
                                        if (v5 <= v5) {
                                            fmt.Println(fi30, v4, fi30)
                                            v36 := GenS2{F0: 42, F1: 100, F2: (v1 << ((v4) & 63))}
                                            _ = v36
                                        }
                                        _ = v35
                                    }
                                    v37 := []int{(^(fi30)), 2147483647, (fi28 ^ fi28)}
                                    _ = v37
                                }
                                _ = fi30
                            }
                        } else {
                            v4 -= 255
                            if (v3 || true) {
                                v4 -= (v4 >> ((v4) & 63))
                            } else {
                                var v38 uint32 = v1
                                if v3 {
                                    fmt.Println(v3, fi28)
                                }
                                if true {
                                    if v3 {
                                        v39 := GenS3{F0: 2, F1: -32768}
                                        v40 := GenS1{F0: uint8(v4), F1: -9223372036854775808}
                                        var v41 uint16 = 2
                                        _ = v39
                                        _ = v40
                                        _ = v41
                                    } else {
                                        v42 := []int{(fi28 - -725), (^(fi28)), fi28}
                                        _ = v42
                                    }
                                    switch fi28 {
                                    case 32767:
                                        if ((v5 | (v5 ^ (v5 - 2))) >= v5) {
                                            var v43 uint32 = (v38 ^ v38)
                                            for fi44 := 0; fi44 < 17; fi44++ {
                                                switch (fi44 % ((int(v5)) | 1)) {
                                                case -786:
                                                    v43 = uint32(v4)
                                                    if ((^(fi44)) == 294) {
                                                        switch (("" + "") + ("" + ("a" + "\n\t"))) {
                                                        case "":
                                                            v45 := GenS4{F0: 1, F1: v3, F2: (fi44 << (((v4 * (v4 + (^(v4))))) & 63))}
                                                            _ = v45
                                                        case "\n\t":
                                                            if (!true) {
                                                                if (580 != 42) {
                                                                    fmt.Println(v1, v3, v4)
                                                                }
                                                                switch fi44 {
                                                                case 880:
                                                                    delete(v2, 32767)
                                                                case 2147483647:
                                                                    fmt.Println(fi44)
                                                                case 128:
                                                                    var v46 float32 = 1e10
                                                                    _ = v46
                                                                default:
                                                                    v43 += (^(v1))
                                                                    fmt.Println(v43, fi28)
                                                                    v2[(^(fi44))] = "hello"
                                                                }
                                                                v4 -= (v4 << (((^(v4))) & 63))
                                                            }
                                                        case "foo bar":
                                                            delete(v2, (fi28 << (((v4 % ((0) | 1))) & 63)))
                                                        default:
                                                            v1 = (^(v1))
                                                        }
                                                    }
                                                case 718:
                                                    v1 += v1
                                                case 547:
                                                    fmt.Println(v4, fi28, fi28)
                                                }
                                                var v47 int32 = 0
                                                fmt.Println(fi44, v3, v47)
                                                _ = fi44
                                                _ = v47
                                            }
                                            _ = v43
                                        }
                                    case -9223372036854775808:
                                        var v48 uint = 255
                                        _ = v48
                                    case -2147483648:
                                        for fi49 := 0; fi49 < 13; fi49++ {
                                            v50 := v2[fi28]
                                            _ = fi49
                                            _ = v50
                                        }
                                    }
                                    v51 := GenS1{F0: v5, F1: fi28}
                                    _ = v51
                                } else {
                                    fmt.Println(v38)
                                    fmt.Println(v4)
                                }
                                _ = v38
                            }
                            for fi52 := 0; fi52 < 2; fi52++ {
                                var v53 float32 = 3.14159
                                v4 += (v4 | (v4 >> ((v4) & 63)))
                                _ = fi52
                                _ = v53
                            }
                        }
                        _ = fi28
                    }
                }
            }
            var v54 uint64 = 1
            v55 := []int{4294967295, 4294967295, 4294967295, 128}
            _ = v54
            _ = v55
        }
    }
    var v56 int64 = -637
    v57 := GenS5{F0: 1, F1: ("a" < "\n\t")}
    var v58 bool = v3
    fmt.Println(v1, v58)
    var v59 uint64 = 0
    v58 = true
    v4 -= 4294967295
    switch 4294967295 {
    case -2147483648:
        var v60 uint16 = uint16(v1)
        _ = v60
    case 128:
        fmt.Println(v4)
    case 127:
        if (0.0 <= 3.14159) {
            v61 := GenS6{F0: "foo bar", F1: 42, F2: 886}
            _ = v61
        } else {
            v2[0] = ("\n\t" + (("foo bar" + "a") + ""))
        }
    }
    for fi62 := 0; fi62 < 13; fi62++ {
        for fi63 := 0; fi63 < 5; fi63++ {
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v59)
            case "\n\t":
                v5 -= v5
            case "hello":
                fmt.Println(v58, v59)
            }
            v59 -= v59
            switch ("a" + "\n\t") {
            case "foo bar":
                var v64 int = 128
                _ = v64
            }
            _ = fi63
        }
        v2[(fi62 & (fi62 / ((fi62) | 1)))] = "hello"
        _ = fi62
    }
    var v65 uint16 = 2
    v65 = uint16(v5)
    for fi66 := 0; fi66 < 9; fi66++ {
        for fi67 := 0; fi67 < 10; fi67++ {
            v3 = v3
            _ = fi67
        }
        _ = fi66
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v65
}
