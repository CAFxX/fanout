package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 int32; F1 int64 }

func main() {
    v1 := []uint32{65535, 329, 0}
    var v2 float32 = 3.14159
    var v3 int8 = 100
    v4 := make(map[int]uint32)
    v5 := GenS1{F0: "foo bar"}
    for fi6 := 0; fi6 < 12; fi6++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi6
    }
    v3 += -32
    _ = len(v4)
    v7 := GenS1{F0: "\n\t"}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 float64 = (1e10 / float64(v2))
        v9 += 0.1
        _ = fi8
        _ = v9
    }
    if true {
        switch -1 {
        case 456:
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        default:
            _ = len(v1)
        }
    } else {
        if true {
            v2 -= 1.5
            _ = len(v1)
        } else {
            for fi10 := 0; fi10 < 3; fi10++ {
                v2 = v2
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 int64 = 328
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v13 := GenS1{F0: ("hello" + "hello")}
                    v2 += (v2 + 1.5)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                if (-1 != 42) {
                    switch (fi10 % ((971) | 1)) {
                    case -1:
                        for fi14 := 0; fi14 < 13; fi14++ {
                            switch "foo bar" {
                            case "hello":
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case "foo bar":
                                v15 := GenS1{F0: ("" + ("\n\t" + "foo bar"))}
                                _ = v15
                            }
                            switch int(v3) {
                            case 2147483647:
                                var v16 string = "a"
                                _ = v16
                            case 32768:
                                var v17 int8 = v3
                                _ = v17
                            default:
                                v3 += -1
                                if (true || ((fi10 & (-(fi14))) >= 4294967295)) {
                                    var v18 int = (fi10 + (fi10 << ((273) & 63)))
                                    v4[fi10] = 916
                                    v18 = (fi14 + (v18 - 777))
                                    _ = v18
                                }
                                var v19 uint32 = 2
                                _ = v19
                            }
                            _ = fi14
                        }
                    }
                    v20 := GenS2{F0: 660, F1: (v3 << ((497) & 63))}
                    v3 -= (v3 >> ((541) & 63))
                    _ = v20
                } else {
                    var v21 int8 = -1
                    for fi22 := 0; fi22 < 4; fi22++ {
                        switch (("" + "") + "hello") {
                        case "":
                            var v23 int32 = 370
                            _ = v23
                        case "a":
                            if (418 >= 0) {
                                var v24 bool = (195 == 522)
                                if v24 {
                                    var v25 float32 = (0.1 + (v2 + v2))
                                    _ = v25
                                } else {
                                    v26 := GenS3{F0: 1000}
                                    v3 = v3
                                    var v27 int = (fi22 >> ((448) & 63))
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                            } else {
                                var v28 int = (fi10 % ((-1) | 1))
                                v29 := GenS2{F0: 42, F1: int8(v28)}
                                switch ("foo bar" + "foo bar") {
                                case "hello":
                                    v3 -= v3
                                }
                                var v30 uint8 = 226
                                var v31 float64 = 1e10
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        case "hello":
                            v21 = 42
                        default:
                            if true {
                                switch (fi22 - fi22) {
                                case 428:
                                    var v32 uint8 = 249
                                    _ = v32
                                case 803:
                                    if (!false) {
                                        var v33 float64 = 3.14159
                                        if ((198 <= 1) || ((float64(0.0) + 1.5) < (v33 + (1.5 - float64(0.1))))) {
                                            v2 = v2
                                            _ = len(v1)
                                            if (true && (!(v33 != 1.5))) {
                                                if false {
                                                    _ = len(v1)
                                                    if true {
                                                        if true {
                                                            v2 = v2
                                                        } else {
                                                            fmt.Println(v3, v21)
                                                            switch fi10 {
                                                            case -796:
                                                                if false {
                                                                    var v34 int = (fi22 % (((fi22 - 127)) | 1))
                                                                    var v35 string = "a"
                                                                    _ = v34
                                                                    _ = v35
                                                                }
                                                            case 32768:
                                                                var v36 float32 = float32(v33)
                                                                _ = v36
                                                            }
                                                            fmt.Println(v33, fi10, v3)
                                                        }
                                                        v4[(-(fi22))] = 0
                                                    } else {
                                                        var v37 int64 = -721
                                                        if (2147483647 >= -1) {
                                                            var v38 float32 = (float32((v33 * v33)) - v2)
                                                            _ = v38
                                                        }
                                                        _ = v37
                                                    }
                                                } else {
                                                    switch fi22 {
                                                    case -9223372036854775808:
                                                        _ = len(v4)
                                                    case 42:
                                                        var v39 string = ("a" + "hello")
                                                        _ = v39
                                                    case 2147483647:
                                                        v40 := GenS1{F0: "a"}
                                                        _ = v40
                                                    }
                                                }
                                                if true {
                                                    v41 := GenS1{F0: "a"}
                                                    var v42 float32 = 0.1
                                                    fmt.Println(fi10)
                                                    if (true && (!(!((false || (18446744073709551615 == 2)) || (351 <= -1))))) {
                                                        if true {
                                                            var v43 int8 = 0
                                                            v21 = v21
                                                            var v44 string = "hello"
                                                            _ = v43
                                                            _ = v44
                                                        }
                                                        v2 = float32(float64(((0.1 * 1e10) + v42)))
                                                        v21 = (-(v3))
                                                    }
                                                    v3 = (v3 / ((v21) | 1))
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = v41
                                                    _ = v42
                                                }
                                                if ("" > ("\n\t" + "")) {
                                                    if (42 > 42) {
                                                        v45 := GenS4{F0: 32768}
                                                        _ = v45
                                                    }
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 2147483648
                                                }
                                            }
                                        }
                                        _ = v33
                                    }
                                }
                                v46 := GenS2{F0: 671, F1: (v21 - v21)}
                                switch (fi10 & -9223372036854775808) {
                                case 127:
                                    switch "" {
                                    case "a":
                                        fmt.Println(v2, v2)
                                    case "foo bar":
                                        var v47 uint64 = 4294967295
                                        _ = v47
                                    case "hello":
                                        switch "a" {
                                        case "foo bar":
                                            v3 += v3
                                        default:
                                            var v48 uint = 4294967295
                                            _ = v48
                                        }
                                    }
                                case -1:
                                    _ = len(v1)
                                default:
                                    var v49 string = "foo bar"
                                    _ = v49
                                }
                                _ = v46
                            } else {
                                var v50 float64 = 1e10
                                delete(v4, (fi10 << ((2) & 63)))
                                var v51 uint64 = 4294967295
                                _ = v50
                                _ = v51
                            }
                        }
                        _ = fi22
                    }
                    _ = v21
                }
                _ = fi10
            }
            var v52 bool = true
            _ = v52
        }
        if (0.0 >= (-2.25 * float32(0.0))) {
            v2 += v2
        }
        switch -2147483648 {
        case 32767:
            var v53 uint32 = 0
            _ = v53
        case 9223372036854775807:
            v54 := GenS5{F0: 32767, F1: -1}
            _ = v54
        case 770:
            var v55 uint32 = 2
            _ = v55
        }
    }
    if (!(-1 == 2147483647)) {
        if false {
            fmt.Println(v2, v3)
        } else {
            switch 4294967295 {
            case 0:
                fmt.Println(v3, v2)
            case -1:
                fmt.Println(v2)
            case -647:
                if (0.1 == 0.0) {
                    _ = len(v1)
                    var v56 bool = true
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 11; fi57++ {
                        var v58 int64 = -1
                        _ = fi57
                        _ = v58
                    }
                    if ((true && (-208 != -1)) && true) {
                        v59 := GenS6{F0: -557, F1: 1}
                        _ = v59
                    }
                    v3 = -33
                }
            default:
                v2 += (-2.25 - v2)
            }
            var v60 bool = true
            for fi61 := 0; fi61 < 13; fi61++ {
                v1 = append(v1, 305)
                _ = fi61
            }
            _ = v60
        }
    } else {
        v3 = v3
        for fi62 := 0; fi62 < 19; fi62++ {
            for fi63 := 0; fi63 < 1; fi63++ {
                var v64 int8 = v3
                _ = fi63
                _ = v64
            }
            v65 := GenS3{F0: -1}
            var v66 uint64 = 18446744073709551615
            _ = fi62
            _ = v65
            _ = v66
        }
        fmt.Println(v2, v2)
    }
    v2 = (v2 * 0.1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
