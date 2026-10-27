package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 bool; F2 uint32 }
func genFn0(p0 uint64) uint64 {
    var v5 float32 = 0.0
    var v6 uint8 = 104
    _ = p0
    _ = v5
    _ = v6
    return p0
}

func main() {
    var v1 string = "hello"
    v2 := []int{-1, 128, 128, -484}
    v3 := make(map[int]string)
    var v4 float64 = ((3.14159 * float64(1.5)) * 3.14159)
    var v7 uint = 255
    fmt.Println(v7, v1, v4)
    switch 0 {
    case -9223372036854775808:
        v4 = 0.1
    case 9223372036854775807:
        var v8 int8 = -32
        _ = v8
    case 4294967295:
        switch -560 {
        case -2147483648:
            var v9 int64 = 1
            _ = v9
        default:
            var v10 uint32 = 2
            _ = v10
        }
    }
    if (0.1 >= 1e10) {
        var v11 string = ""
        _ = v11
    } else {
        if (-9223372036854775808 != 1) {
            fmt.Println(v1)
        }
    }
    v7 += (v7 % (((v7 * v7)) | 1))
    v12 := GenS1{F0: 225}
    switch "foo bar" {
    case "hello":
        v7 = v7
    case "\n\t":
        var v13 float64 = v4
        _ = v13
    case "foo bar":
        for fi14 := 0; fi14 < 13; fi14++ {
            v15 := GenS1{F0: 202}
            var v16 uint64 = genFn0(18446744073709551615)
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    v17 := GenS2{F0: -1}
    switch ((v1 + "hello") + v1) {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        switch 2147483647 {
        case 21:
            switch 0 {
            case 0:
                fmt.Println(v7)
            default:
                if false {
                    for fi18 := 0; fi18 < 2; fi18++ {
                        v3[fi18] = v1
                        for fi19 := 0; fi19 < 12; fi19++ {
                            v1 = v1
                            switch fi19 {
                            case 781:
                                var v20 uint32 = 681
                                _ = v20
                            default:
                                var v21 uint = (v7 << (((v7 ^ (v7 / ((648) | 1)))) & 63))
                                _ = v21
                            }
                            _ = fi19
                        }
                        _ = fi18
                    }
                    v4 = float64((1e10 / ((((1.5 - 1e10) - ((0.1 + (float32(float64(1e10)) * 0.1)) * (0.1 / 1.0))) * 0.0) / float32(v4))))
                } else {
                    var v22 uint = (v7 + (v7 | v7))
                    v23 := GenS2{F0: -9223372036854775808}
                    for fi24 := 0; fi24 < 20; fi24++ {
                        v25 := GenS1{F0: 211}
                        _ = fi24
                        _ = v25
                    }
                    _ = v22
                    _ = v23
                }
            }
        case 32768:
            for fi26 := 0; fi26 < 8; fi26++ {
                for fi27 := 0; fi27 < 12; fi27++ {
                    _ = len(v2)
                    var v28 uint64 = 2
                    _ = fi27
                    _ = v28
                }
                _ = fi26
            }
        case 4294967295:
            var v29 float32 = 3.14159
            _ = v29
        default:
            v2 = append(v2, 1)
        }
    default:
        if (!true) {
            for fi30 := 0; fi30 < 13; fi30++ {
                v4 = (v4 * 0.1)
                for fi31 := 0; fi31 < 6; fi31++ {
                    fmt.Println(fi31, v4, fi31)
                    v7 = 255
                    _ = fi31
                }
                _ = fi30
            }
            switch v1 {
            case "":
                switch v1 {
                case "foo bar":
                    var v32 int64 = 9223372036854775807
                    _ = v32
                default:
                    switch 4294967295 {
                    case 32767:
                        switch -97 {
                        case -2147483648:
                            if (!true) {
                                _ = len(v2)
                            }
                        default:
                            v1 = "foo bar"
                        }
                    }
                }
            case "a":
                delete(v3, 32767)
            case "\n\t":
                var v33 uint64 = 694
                _ = v33
            default:
                switch 128 {
                case 32767:
                    fmt.Println(v4, v4)
                case 127:
                    switch -937 {
                    case 9223372036854775807:
                        var v34 uint = v7
                        _ = v34
                    case 32767:
                        var v35 int64 = 9223372036854775807
                        _ = v35
                    default:
                        v36 := GenS3{F0: 351, F1: (((!true) && false) && true), F2: 4294967295}
                        _ = v36
                    }
                }
            }
            var v37 float32 = (1.5 / float32((v4 * 3.14159)))
            _ = v37
        } else {
            var v38 bool = true
            _ = v38
        }
    }
    if true {
        for fi39 := 0; fi39 < 17; fi39++ {
            var v40 uint = (^(v7))
            v41 := v3[(fi39 << ((v7) & 63))]
            fmt.Println(v7)
            _ = fi39
            _ = v40
            _ = v41
        }
        var v42 float64 = 0.1
        _ = v42
    } else {
        var v43 uint32 = 42
        var v44 int64 = 1
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v43
        _ = v44
    }
    if (true && ((true || true) && true)) {
        v45 := GenS2{F0: -73}
        _ = v45
    } else {
        v1 = (v1 + v1)
    }
    v2 = append(v2, 964)
    var v46 uint64 = 0
    v3[42] = v1
    v1 = "\n\t"
    var v47 float64 = ((v4 / 1.0) / v4)
    switch ((("hello" + v1) + (v1 + "hello")) + "hello") {
    case "foo bar":
        v48 := GenS2{F0: 9223372036854775807}
        _ = v48
    }
    if (((("" + "\n\t") + "\n\t") + ((v1 + v1) + v1)) == "a") {
        v49 := GenS3{F0: 284, F1: false, F2: 1}
        _ = v49
    } else {
        v50 := GenS3{F0: 65535, F1: (-264 < 2147483647), F2: 942}
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v7
}
