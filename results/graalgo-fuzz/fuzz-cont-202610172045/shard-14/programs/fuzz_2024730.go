package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 bool; F2 bool }
type GenS4 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 uint32; F1 int }
type GenS7 struct { F0 int16; F1 int; F2 int }
type GenS8 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int32; F1 uint32 }
func genFn0(p0 bool) uint {
    if false {
        for fi4 := 0; fi4 < 14; fi4++ {
            fmt.Println(fi4, p0)
            var v5 uint = 257
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(p0)
    switch "" {
    case "\n\t":
        fmt.Println(p0)
    case "hello":
        fmt.Println(p0)
    case "":
        fmt.Println(p0)
    }
    _ = p0
    return 2
}
func genFn1(p0 bool) uint {
    if (1 < 0) {
        fmt.Println(p0)
    } else {
        v6 := GenS1{F0: 1, F1: 42, F2: 127}
        var v7 uint = 65535
        _ = v6
        _ = v7
    }
    v8 := []int{2147483647, -619}
    _ = p0
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{}
    v3 := make(map[uint32]bool)
    v9 := GenS2{F0: 32767}
    v2 = append(v2, "hello")
    var v10 int = -812
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 = (v10 % (((v10 & v10)) | 1))
    var v11 int32 = 2147483647
    v11 = -398
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            switch 42 {
            case 4294967295:
                var v13 bool = (((186 >= 9223372036854775807) && false) && true)
                _ = v13
            }
            if (!true) {
                switch "hello" {
                case "foo bar":
                    switch ("\n\t" + "hello") {
                    case "foo bar":
                        if (-2147483648 == (-(v11))) {
                            var v14 uint16 = 65535
                            _ = v14
                        }
                    case "a":
                        var v15 string = ("\n\t" + "a")
                        _ = v15
                    case "hello":
                        v16 := GenS2{F0: 1}
                        _ = v16
                    default:
                        var v17 uint32 = 65535
                        _ = v17
                    }
                case "hello":
                    v18 := v3[2]
                    _ = v18
                }
            }
            _ = fi12
        }
    } else {
        v19 := GenS3{F0: -1, F1: false, F2: true}
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := GenS1{F0: -1, F1: 42, F2: 1}
        _ = fi20
        _ = v21
    }
    fmt.Println(v11, v11)
    v11 = v11
    var v22 uint8 = 1
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 uint = 255
        v22 = v22
        for fi25 := 0; fi25 < 1; fi25++ {
            v11 = 0
            if true {
                var v26 float64 = 3.14159
                v27 := GenS1{F0: 9223372036854775807, F1: 219, F2: v22}
                switch int(v22) {
                case 705:
                    v26 -= 0.0
                case 4294967295:
                    if false {
                        v22 += 128
                        delete(v3, uint32(v24))
                    }
                default:
                    switch ("foo bar" + ((("\n\t" + ("\n\t" + ("\n\t" + ("" + (("\n\t" + "hello") + ("hello" + "a")))))) + "") + "")) {
                    case "a":
                        var v28 uint32 = 65535
                        _ = v28
                    case "\n\t":
                        v29 := GenS3{F0: 0, F1: (false || (!((2 == 65535) && true))), F2: (v22 <= (v22 % (((v22 + 0)) | 1)))}
                        _ = v29
                    case "foo bar":
                        if false {
                            fmt.Println(fi25)
                            var v30 float32 = (-2.25 * (1e10 / (3.14159 - 0.1)))
                            v30 = 1e10
                            _ = v30
                        }
                    default:
                        v11 = v11
                    }
                }
                var v31 uint16 = 699
                _ = v26
                _ = v27
                _ = v31
            }
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v32 := GenS4{F0: (v11 & 340), F1: 194, F2: 42}
    if (int16(v10) < 1) {
        switch int(v22) {
        case 9223372036854775807:
            if ((0.1 == ((0.0 - 0.1) + 0.0)) || true) {
                var v33 float64 = 0.1
                v34 := v3[2147483647]
                _ = v33
                _ = v34
            }
        case -469:
            if (1000 == 1000) {
                fmt.Println(v11)
                v22 = v22
            }
        }
    }
    if true {
        for fi35 := 0; fi35 < 5; fi35++ {
            for fi36 := 0; fi36 < 14; fi36++ {
                v37 := GenS5{F0: 65535, F1: 0}
                _ = fi36
                _ = v37
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 uint16 = 297
                v39 += (v39 / (((^(v39))) | 1))
                _ = fi38
                _ = v39
            }
            v40 := GenS2{F0: -32768}
            _ = fi35
            _ = v40
        }
        switch "\n\t" {
        case "hello":
            var v41 uint64 = 42
            _ = v41
        }
    } else {
        switch "" {
        case "hello":
            for fi42 := 0; fi42 < 1; fi42++ {
                v10 -= (-(v10))
                _ = fi42
            }
        default:
            var v43 int8 = 34
            _ = v43
        }
        fmt.Println(v22, v10, v11)
        if true {
            for fi44 := 0; fi44 < 12; fi44++ {
                if false {
                    if (true && true) {
                        v45 := GenS6{F0: 836, F1: 0}
                        switch "foo bar" {
                        case "a":
                            if true {
                                if false {
                                    v46 := GenS7{F0: -1, F1: (fi44 - (fi44 ^ v10)), F2: (-(fi44))}
                                    v3[2147483647] = false
                                    _ = v46
                                } else {
                                    var v47 int16 = 0
                                    v47 = v47
                                    _ = v47
                                }
                                fmt.Println(v22)
                                var v48 bool = true
                                _ = v48
                            } else {
                                if true {
                                    if (!(!false)) {
                                        v11 = 2147483647
                                        v49 := GenS2{F0: 0}
                                        _ = v49
                                    } else {
                                        switch "\n\t" {
                                        case "hello":
                                            if false {
                                                fmt.Println(v10, v22)
                                                if (255 < 255) {
                                                    v50 := v1[fi44]
                                                    v50 = (!((v22 | v22) == v22))
                                                    var v51 bool = (!true)
                                                    _ = v50
                                                    _ = v51
                                                }
                                            }
                                        case "foo bar":
                                            v10 += fi44
                                        case "":
                                            fmt.Println(v22, v22)
                                        }
                                        for fi52 := 0; fi52 < 6; fi52++ {
                                            var v53 int = v10
                                            var v54 float32 = (float32(((1.5 + (1.5 - (float64(1e10) / (float64(((((1.5 - 1e10) - float32((float64(3.14159) / (((0.0 / 1.5) - -2.25) + float64(1.5))))) * (float32((1.5 - (float64(3.14159) + (3.14159 - 1.5)))) + ((float32((3.14159 * float64((1e10 * 1.5)))) - 0.1) + 1e10))) / 1.0)) - -2.25)))) / 0.1)) * 0.0)
                                            _ = fi52
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                    v1[v10] = false
                                    for fi55 := 0; fi55 < 8; fi55++ {
                                        v22 = (v22 % ((v22) | 1))
                                        v56 := GenS8{F0: 255, F1: 0, F2: 42}
                                        v22 = v22
                                        v57 := GenS2{F0: 42}
                                        var v58 int64 = 1
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    }
                                } else {
                                    v59 := GenS9{F0: (v11 & (-(v11))), F1: 1}
                                    v60 := GenS4{F0: (^(v11)), F1: 32767, F2: 65535}
                                    _ = v59
                                    _ = v60
                                }
                                v61 := GenS8{F0: 0, F1: 778, F2: 0}
                                v22 -= 255
                                _ = v61
                            }
                        }
                        v22 += v22
                        _ = v45
                    }
                } else {
                    var v62 int16 = -32768
                    for fi63 := 0; fi63 < 5; fi63++ {
                        v10 = v10
                        if (451 <= 737) {
                            var v64 int32 = v11
                            v64 = 394
                            _ = v64
                        }
                        var v65 bool = true
                        _ = fi63
                        _ = v65
                    }
                    v62 += v62
                    _ = v62
                }
                if len(v2) > 0 {
                    v2[0] = ""
                }
                v66 := v1[(fi44 & (fi44 | fi44))]
                _ = fi44
                _ = v66
            }
            switch 0 {
            case 32768:
                v22 = (v22 * 42)
            case 128:
                v67 := GenS4{F0: -165, F1: 65535, F2: 32767}
                _ = v67
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v9
}
