package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 uint8; F2 bool }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 bool; F1 int8 }

func main() {
    var v1 bool = (0.0 < 0.1)
    var v2 float32 = (1.5 * (0.1 + float32(0.0)))
    var v3 uint64 = 188
    fmt.Println(v3)
    var v4 uint32 = 2147483647
    if (("\n\t" + ("hello" + "\n\t")) > "\n\t") {
        if (!v1) {
            v5 := GenS1{F0: 261}
            var v6 int8 = 1
            _ = v5
            _ = v6
        } else {
            v4 -= 821
            v7 := GenS1{F0: 2}
            _ = v7
        }
        var v8 uint = 404
        v9 := GenS2{F0: v8}
        _ = v8
        _ = v9
    }
    v10 := []int{}
    v2 -= 1.5
    for fi11 := 0; fi11 < 1; fi11++ {
        switch ("" + "hello") {
        case "foo bar":
            for fi12 := 0; fi12 < 17; fi12++ {
                v2 = 0.0
                v3 += (v3 ^ v3)
                _ = fi12
            }
        case "":
            v13 := GenS3{F0: 103}
            _ = v13
        }
        v1 = v1
        _ = fi11
    }
    if v1 {
        fmt.Println(v4, v1)
    }
    if (!true) {
        if (((0.0 + ((((1e10 + 1.5) - (-2.25 * -2.25)) / 1.0) / ((((1.5 * 3.14159) / 1.0) * 1.5) * float64(float32(((((float64((v2 - v2)) + 1.5) * -2.25) / ((((((0.0 / (0.0 * float64((v2 - v2)))) - 0.0) - (0.0 / -2.25)) + 3.14159) * (((3.14159 / (3.14159 * 1.5)) * (1e10 * 1e10)) / float64(v2))) / (0.0 / float64((1e10 - (1.5 * v2)))))) + 1.5)))))) > 0.1) && false) {
            for fi14 := 0; fi14 < 4; fi14++ {
                if len(v10) > 0 {
                    v10[0] = 1
                }
                _ = fi14
            }
            for fi15 := 0; fi15 < 16; fi15++ {
                fmt.Println(fi15)
                _ = fi15
            }
        } else {
            var v16 string = "hello"
            if v1 {
                var v17 uint16 = uint16(v4)
                fmt.Println(v3)
                var v18 uint16 = (v17 + 0)
                _ = v17
                _ = v18
            }
            switch "hello" {
            case "":
                if (32768 != 32767) {
                    for fi19 := 0; fi19 < 2; fi19++ {
                        v4 -= v4
                        _ = fi19
                    }
                    v20 := GenS4{F0: 0, F1: 128, F2: ((v3 % ((v3) | 1)) > (v3 * (v3 & 0)))}
                    v21 := GenS1{F0: 32767}
                    _ = v20
                    _ = v21
                }
            case "a":
                v16 = (v16 + "\n\t")
            default:
                v22 := GenS2{F0: 255}
                _ = v22
            }
            _ = v16
        }
        switch "a" {
        case "a":
            switch 32767 {
            case 128:
                var v23 uint = 42
                _ = v23
            case 767:
                v24 := GenS2{F0: 255}
                _ = v24
            }
        default:
            for fi25 := 0; fi25 < 15; fi25++ {
                fmt.Println(v1)
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi25
            }
        }
    }
    v26 := GenS3{F0: 243}
    v27 := GenS2{F0: 4294967295}
    v1 = false
    fmt.Println(v2, v2, v3)
    fmt.Println(v3, v4, v3)
    if (1 == 9223372036854775807) {
        switch "" {
        case "hello":
            v10 = append(v10, -651)
        case "a":
            v4 += (v4 - (^(v4)))
        case "\n\t":
            var v28 string = "foo bar"
            _ = v28
        }
        if true {
            switch -9223372036854775808 {
            case -2147483648:
                var v29 uint8 = 217
                _ = v29
            case 734:
                for fi30 := 0; fi30 < 11; fi30++ {
                    if v1 {
                        if v1 {
                            v10 = append(v10, (fi30 >> ((827) & 63)))
                            if (0 >= -501) {
                                var v31 int16 = 1000
                                var v32 uint = 172
                                var v33 uint = v32
                                _ = v31
                                _ = v32
                                _ = v33
                            } else {
                                fmt.Println(v4)
                                v2 = float32((((3.14159 - (((1e10 / 1.5) - (0.0 * 0.1)) + ((0.0 + ((1e10 - 0.1) - 3.14159)) - 0.0))) * (1e10 / 0.1)) * 0.1))
                                v34 := GenS1{F0: 0}
                                _ = v34
                            }
                        }
                    } else {
                        for fi35 := 0; fi35 < 7; fi35++ {
                            var v36 bool = v1
                            if (!v36) {
                                fmt.Println(v36, v4, fi35)
                                switch "hello" {
                                case "hello":
                                    v37 := GenS3{F0: 9}
                                    _ = v37
                                case "foo bar":
                                    v3 += (v3 / ((42) | 1))
                                }
                            }
                            _ = len(v10)
                            _ = fi35
                            _ = v36
                        }
                        v4 += 2
                        if (!v1) {
                            if len(v10) >= 2 {
                                v10 = v10[1:len(v10)]
                            }
                            v1 = false
                            var v38 int = (fi30 << ((255) & 63))
                            _ = v38
                        } else {
                            v39 := GenS5{F0: v3, F1: -136}
                            var v40 uint32 = v4
                            _ = v39
                            _ = v40
                        }
                    }
                    _ = fi30
                }
            }
        }
        for fi41 := 0; fi41 < 1; fi41++ {
            var v42 string = "\n\t"
            _ = len(v10)
            switch v42 {
            case "a":
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            case "hello":
                if (!((1 == 42) && false)) {
                    var v43 int64 = -1
                    for fi44 := 0; fi44 < 12; fi44++ {
                        v43 = (v43 - (v43 | (v43 ^ (-(v43)))))
                        if (v1 || v1) {
                            if (0 != 1) {
                                var v45 int32 = 175
                                _ = v45
                            } else {
                                fmt.Println(fi44)
                                v43 = v43
                                v46 := GenS2{F0: 854}
                                if (v42 > (v42 + "")) {
                                    var v47 float64 = (1e10 - -2.25)
                                    var v48 float32 = 0.1
                                    _ = len(v10)
                                    _ = v47
                                    _ = v48
                                } else {
                                    v49 := GenS2{F0: 4294967295}
                                    var v50 int32 = -211
                                    _ = v49
                                    _ = v50
                                }
                                v51 := GenS6{F0: (!(1 != -2147483648)), F1: int8(fi41)}
                                if len(v10) > 0 {
                                    v10[0] = (-(fi44))
                                }
                                _ = v46
                                _ = v51
                            }
                            var v52 int = (fi41 >> ((312) & 63))
                            _ = v52
                        }
                        _ = fi44
                    }
                    _ = v43
                } else {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    v10 = append(v10, (fi41 | (fi41 & 0)))
                    v42 = "hello"
                }
            }
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v3)
    switch -543 {
    case 0:
        v3 = v3
    case -858:
        v4 -= v4
    case 4294967295:
        for fi53 := 0; fi53 < 18; fi53++ {
            var v54 int = (fi53 ^ fi53)
            _ = fi53
            _ = v54
        }
    default:
        v3 -= v3
    }
    switch 2147483647 {
    case -9223372036854775808:
        switch "foo bar" {
        case "foo bar":
            if len(v10) > 0 {
                v10[0] = 32767
            }
        case "":
            v2 = (v2 * v2)
        }
    default:
        var v55 uint8 = 255
        _ = v55
    }
    for fi56 := 0; fi56 < 20; fi56++ {
        var v57 uint8 = 255
        fmt.Println(v4, v2, v57)
        _ = fi56
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
