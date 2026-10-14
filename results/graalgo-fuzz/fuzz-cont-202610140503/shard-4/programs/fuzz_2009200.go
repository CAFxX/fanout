package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int }

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = (0.1 / (1.5 - 0.1))
    v3 := make(map[uint32]int)
    switch -9223372036854775808 {
    case -9223372036854775808:
        var v4 int = 32767
        _ = v4
    case 9223372036854775807:
        var v5 int8 = -128
        _ = v5
    case 4294967295:
        switch "" {
        case "hello":
            _ = len(v3)
        case "foo bar":
            v6 := []int{-358, -856}
            _ = v6
        case "":
            v7 := GenS1{F0: -243, F1: 0}
            _ = v7
        default:
            if true {
                switch -923 {
                case 0:
                    v8 := GenS2{F0: -1}
                    _ = v8
                case 4294967295:
                    for fi9 := 0; fi9 < 3; fi9++ {
                        v2 = float32(-2.25)
                        _ = fi9
                    }
                case 128:
                    v2 += (v2 * v2)
                }
            }
        }
    }
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 9; fi11++ {
            if ((("hello" + ("\n\t" + ((("" + "") + "foo bar") + "\n\t"))) + (("a" + "") + ("hello" + ("hello" + ("hello" + "foo bar"))))) == "hello") {
                if (1 < -2147483648) {
                    delete(v1, (fi11 ^ fi10))
                } else {
                    if true {
                        v2 = -2.25
                        v12 := []int{127, fi10}
                        switch fi11 {
                        case 4294967295:
                            v13 := GenS3{F0: 100}
                            _ = v13
                        case 42:
                            _ = len(v1)
                        default:
                            var v14 uint32 = 53
                            _ = v14
                        }
                        _ = v12
                    } else {
                        v2 = ((v2 + 1e10) - (float32(1e10) / v2))
                        var v15 float32 = (1e10 / v2)
                        var v16 uint = 722
                        if (v2 != v15) {
                            var v17 int16 = 42
                            fmt.Println(v16, v15, v15)
                            var v18 uint64 = 0
                            v15 = (0.0 - 0.0)
                            _ = v17
                            _ = v18
                        } else {
                            var v19 int64 = -675
                            switch (fi11 * fi11) {
                            case 277:
                                var v20 uint32 = 1
                                _ = v20
                            default:
                                delete(v3, 4294967295)
                            }
                            var v21 uint = 764
                            _ = v19
                            _ = v21
                        }
                        _ = v15
                        _ = v16
                    }
                }
                switch fi10 {
                case 4294967295:
                    var v22 float32 = v2
                    _ = v22
                case -181:
                    if false {
                        v2 -= 0.0
                    } else {
                        if (2 < 0) {
                            v2 += v2
                            var v23 int64 = -1
                            _ = v23
                        } else {
                            if true {
                                fmt.Println(v2, v2, v2)
                                v24 := GenS2{F0: 1}
                                _ = v24
                            }
                            v25 := GenS1{F0: 0, F1: 32767}
                            var v26 int16 = 32767
                            _ = v25
                            _ = v26
                        }
                    }
                case -563:
                    if true {
                        var v27 uint = 987
                        var v28 uint = 4294967295
                        v29 := GenS1{F0: 32767, F1: 65535}
                        _ = v27
                        _ = v28
                        _ = v29
                    } else {
                        var v30 int64 = int64(fi10)
                        if false {
                            var v31 uint64 = 0
                            _ = v31
                        } else {
                            if (!(2 == 2147483647)) {
                                switch ((("foo bar" + "hello") + "foo bar") + "foo bar") {
                                case "":
                                    v32 := GenS4{F0: 521}
                                    _ = v32
                                case "a":
                                    v33 := GenS5{F0: fi11}
                                    _ = v33
                                default:
                                    var v34 uint = 910
                                    _ = v34
                                }
                                v30 = 9223372036854775807
                                fmt.Println(v30)
                            }
                            var v35 uint32 = 699
                            _ = v35
                        }
                        var v36 uint8 = 1
                        _ = v30
                        _ = v36
                    }
                default:
                    v2 = 3.14159
                }
                switch (fi10 / (((fi10 | -1)) | 1)) {
                case 2147483647:
                    fmt.Println(fi10)
                case 32767:
                    if true {
                        v2 -= v2
                        if true {
                            var v37 int64 = 42
                            _ = v37
                        } else {
                            var v38 uint16 = 0
                            _ = v38
                        }
                    }
                }
            } else {
                v1[(fi10 & 0)] = (654 <= 720)
                switch "hello" {
                case "hello":
                    v39 := []int{(fi10 - 4294967295), (fi11 / ((fi10) | 1))}
                    _ = v39
                case "":
                    var v40 int32 = -2147483648
                    _ = v40
                case "foo bar":
                    var v41 uint32 = 2147483647
                    _ = v41
                default:
                    fmt.Println(v2, fi11)
                }
                var v42 uint16 = 2
                _ = v42
            }
            var v43 uint32 = 0
            _ = fi11
            _ = v43
        }
        for fi44 := 0; fi44 < 20; fi44++ {
            fmt.Println(fi10)
            v45 := GenS4{F0: 1}
            v2 -= -2.25
            _ = fi44
            _ = v45
        }
        switch "a" {
        case "hello":
            v2 = (float32(-2.25) + v2)
        case "":
            for fi46 := 0; fi46 < 7; fi46++ {
                if (!false) {
                    var v47 int16 = 701
                    fmt.Println(fi46)
                    _ = v47
                }
                v2 += ((3.14159 + (v2 + v2)) * v2)
                switch (fi46 >> ((0) & 63)) {
                case 32768:
                    v2 -= v2
                case 9223372036854775807:
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        v48 := GenS4{F0: 1}
                        _ = v48
                    case "hello":
                        v49 := []int{-9223372036854775808, -436, (fi46 ^ fi46)}
                        _ = v49
                    case "":
                        v50 := []int{640, (fi10 % ((fi46) | 1)), (fi46 / ((fi46) | 1))}
                        _ = v50
                    default:
                        var v51 float64 = 1.5
                        _ = v51
                    }
                case -2147483648:
                    fmt.Println(fi46)
                }
                _ = fi46
            }
        default:
            fmt.Println(v2, v2)
        }
        _ = fi10
    }
    v2 += 3.14159
    v52 := GenS3{F0: 100}
    v53 := []int{127, -839, -667}
    var v54 uint64 = 18446744073709551615
    for fi55 := 0; fi55 < 1; fi55++ {
        for fi56 := 0; fi56 < 12; fi56++ {
            _ = len(v53)
            v57 := GenS3{F0: -1}
            v54 += v54
            _ = fi56
            _ = v57
        }
        var v58 int16 = -1
        fmt.Println(fi55, v58, v54)
        _ = fi55
        _ = v58
    }
    switch "\n\t" {
    case "a":
        var v59 uint16 = 65535
        _ = v59
    default:
        v60 := GenS4{F0: 4294967295}
        _ = v60
    }
    switch "hello" {
    case "a":
        fmt.Println(v2, v2)
    default:
        switch -120 {
        case 32768:
            for fi61 := 0; fi61 < 7; fi61++ {
                var v62 int16 = 1000
                var v63 int8 = 0
                _ = fi61
                _ = v62
                _ = v63
            }
        case 902:
            v54 -= v54
        }
    }
    v64 := GenS3{F0: 20}
    v65 := GenS2{F0: 963}
    v2 += v2
    v54 -= 4294967295
    if len(v53) > 0 {
        v53[0] = 579
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v53:", len(v53), v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v54
    _ = v64
    _ = v65
}
