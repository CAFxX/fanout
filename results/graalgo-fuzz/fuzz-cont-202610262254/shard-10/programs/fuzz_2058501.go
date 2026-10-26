package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16; F1 bool; F2 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 uint; F2 int }
type GenS7 struct { F0 string; F1 string; F2 bool }
type GenS8 struct { F0 int16; F1 uint64 }
type GenS9 struct { F0 int16; F1 int16 }
type GenS10 struct { F0 uint32; F1 uint32 }
type GenS11 struct { F0 int64 }

func main() {
    v1 := []uint32{2147483647, 2147483648, 33}
    var v2 float64 = (0.1 + ((1.5 * (float64((0.1 - 1e10)) + 3.14159)) + 1e10))
    v3 := make(map[uint32]uint32)
    var v4 int32 = -772
    if (1 >= 757) {
        v5 := GenS1{F0: -850}
        var v6 uint = 255
        _ = v5
        _ = v6
    } else {
        if false {
            var v7 uint8 = 2
            _ = v7
        } else {
            v4 = (v4 >> ((2) & 63))
            if (0.0 < 0.1) {
                v2 = 0.0
            }
            for fi8 := 0; fi8 < 3; fi8++ {
                switch 128 {
                case 32768:
                    if true {
                        var v9 float64 = 0.1
                        if false {
                            var v10 uint8 = 2
                            _ = v10
                        }
                        _ = v9
                    } else {
                        var v11 int = (fi8 - fi8)
                        _ = v11
                    }
                default:
                    fmt.Println(fi8, v4, v2)
                }
                v12 := GenS2{F0: ((!(360 < 2147483648)) && (true && false))}
                _ = fi8
                _ = v12
            }
        }
        var v13 float32 = 0.0
        _ = v13
    }
    if len(v1) > 0 {
        v1[0] = 0
    }
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint64 = 18446744073709551615
        for fi17 := 0; fi17 < 11; fi17++ {
            var v18 uint64 = 18446744073709551615
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    _ = len(v3)
    v2 -= (1e10 * float64(1.5))
    v2 -= v2
    v4 += (v4 | v4)
    var v19 int = -32
    var v20 float64 = (v2 + v2)
    for fi21 := 0; fi21 < 5; fi21++ {
        v22 := GenS2{F0: (4294967295 > 2147483648)}
        for fi23 := 0; fi23 < 11; fi23++ {
            switch (fi21 * fi23) {
            case 491:
                _ = len(v1)
            case 9223372036854775807:
                v24 := GenS1{F0: 42}
                var v25 bool = false
                v26 := GenS3{F0: 255, F1: 448}
                var v27 float32 = -2.25
                v28 := v3[43]
                if ((627 != 818) && v25) {
                    fmt.Println(fi21, v19)
                    if (255 <= 79) {
                        v19 -= 32768
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v29 := GenS1{F0: 1}
                        _ = v29
                    } else {
                        switch (v19 - 42) {
                        case -177:
                            v28 = v28
                        case 42:
                            if true {
                                switch v19 {
                                case -2147483648:
                                    switch (fi21 & fi21) {
                                    case 510:
                                        var v30 uint = 1
                                        _ = v30
                                    case -934:
                                        var v31 uint64 = 2
                                        _ = v31
                                    case 161:
                                        if (!((1 >= -11) && (v25 || (((v20 + 0.0) != v20) && v25)))) {
                                            switch v19 {
                                            case 2147483647:
                                                var v32 string = "a"
                                                _ = v32
                                            case 32767:
                                                v20 = v2
                                            default:
                                                fmt.Println(v28, v19)
                                            }
                                            v20 += (v20 + (v2 + 3.14159))
                                        } else {
                                            v19 = fi21
                                            if (626 == 88) {
                                                if (127 < uint8(v28)) {
                                                    if v25 {
                                                        v33 := GenS4{F0: 42, F1: v25, F2: -128}
                                                        _ = len(v3)
                                                        if (-2.25 < 1e10) {
                                                            switch (fi21 % (((fi21 % ((42) | 1))) | 1)) {
                                                            case 418:
                                                                v34 := GenS3{F0: 65535, F1: int64(v4)}
                                                                _ = v34
                                                            case 32767:
                                                                v25 = v25
                                                            }
                                                            if (!(1 > 65535)) {
                                                                v28 += v28
                                                                _ = len(v3)
                                                            } else {
                                                                if true {
                                                                    v35 := GenS3{F0: 845, F1: -20}
                                                                    _ = v35
                                                                } else {
                                                                    if ((100 > 11) && v25) {
                                                                        v25 = v25
                                                                    } else {
                                                                        switch (v19 & fi21) {
                                                                        case -502:
                                                                            switch fi21 {
                                                                            case 127:
                                                                                var v36 int64 = -614
                                                                                _ = v36
                                                                            case 32768:
                                                                                v4 = v4
                                                                            case 527:
                                                                                v2 += v20
                                                                            default:
                                                                                v28 -= (v28 % ((v28) | 1))
                                                                            }
                                                                        case -1:
                                                                            v37 := GenS4{F0: -403, F1: v25, F2: 0}
                                                                            _ = v37
                                                                        case -452:
                                                                            switch v19 {
                                                                            case -2147483648:
                                                                                _ = len(v3)
                                                                            }
                                                                        }
                                                                        var v38 float64 = float64((((v27 + v27) / float32(v2)) / (v27 - 0.1)))
                                                                        _ = v38
                                                                    }
                                                                }
                                                                switch "\n\t" {
                                                                case "":
                                                                    v39 := GenS5{F0: "\n\t"}
                                                                    _ = v39
                                                                }
                                                            }
                                                        }
                                                        _ = v33
                                                    }
                                                    var v40 uint16 = 144
                                                    v27 = (1.5 - (v27 * (v27 / 1.5)))
                                                    _ = v40
                                                } else {
                                                    fmt.Println(v27, v27, v20)
                                                    v19 += fi21
                                                    var v41 float32 = ((v27 + v27) / (float32(v20) + 1e10))
                                                    _ = v41
                                                }
                                            } else {
                                                v27 = 0.1
                                                fmt.Println(v28)
                                                v3[(v28 << ((2) & 63))] = 1
                                                if true {
                                                    if false {
                                                        if (true || v25) {
                                                            v28 -= (v28 * (v28 << ((1) & 63)))
                                                        }
                                                        v3[(^(v28))] = v28
                                                    } else {
                                                        var v42 uint16 = 65535
                                                        fmt.Println(v27, fi21, fi23)
                                                        v43 := GenS2{F0: v25}
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    fmt.Println(v2)
                                                } else {
                                                    var v44 uint8 = 255
                                                    fmt.Println(v27, v19)
                                                    if len(v1) > 0 {
                                                        v1[0] = v28
                                                    }
                                                    _ = v44
                                                }
                                            }
                                        }
                                    default:
                                        if (42 > 91) {
                                            var v45 uint8 = uint8(v28)
                                            _ = v45
                                        }
                                    }
                                case -1:
                                    if (255 != 192) {
                                        switch (fi23 & v19) {
                                        case 127:
                                            var v46 int64 = -9223372036854775808
                                            _ = v46
                                        case 344:
                                            var v47 int8 = -35
                                            _ = v47
                                        case -2147483648:
                                            v48 := v3[v28]
                                            fmt.Println(v25)
                                            v4 = v4
                                            _ = v48
                                        }
                                        var v49 string = "foo bar"
                                        _ = v49
                                    } else {
                                        var v50 uint32 = 2147483648
                                        v2 = v20
                                        _ = v50
                                    }
                                case 127:
                                    v28 = v28
                                }
                                var v51 int16 = 42
                                fmt.Println(v28, v19, v19)
                                _ = v51
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                var v52 uint64 = 91
                                _ = v52
                            }
                        case 128:
                            if len(v1) > 0 {
                                v1[0] = v28
                            }
                        }
                    }
                    v27 = (-2.25 * v27)
                    v53 := GenS6{F0: 1, F1: 65535, F2: 692}
                    v54 := GenS7{F0: ("\n\t" + "foo bar"), F1: "", F2: ((true && (-1 != 1)) || v25)}
                    _ = v53
                    _ = v54
                } else {
                    fmt.Println(v4, v25, v2)
                    v55 := GenS8{F0: -150, F1: 0}
                    v56 := v3[v28]
                    _ = v55
                    _ = v56
                }
                _ = v24
                _ = v25
                _ = v26
                _ = v27
                _ = v28
            case 900:
                v57 := GenS3{F0: 4294967295, F1: -9223372036854775808}
                _ = v57
            }
            v58 := GenS5{F0: ("a" + (("foo bar" + "hello") + "\n\t"))}
            _ = fi23
            _ = v58
        }
        v59 := GenS9{F0: 32767, F1: -531}
        _ = fi21
        _ = v22
        _ = v59
    }
    var v60 int16 = 583
    if ((-1 >= 42) || true) {
        var v61 uint16 = 65535
        for fi62 := 0; fi62 < 18; fi62++ {
            var v63 uint16 = v61
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi62
            _ = v63
        }
        switch v19 {
        case -954:
            if true {
                v4 = (v4 & v4)
                for fi64 := 0; fi64 < 16; fi64++ {
                    v60 += (^(v60))
                    var v65 uint64 = 2
                    v66 := GenS10{F0: 0, F1: 338}
                    _ = fi64
                    _ = v65
                    _ = v66
                }
            } else {
                var v67 uint8 = 42
                for fi68 := 0; fi68 < 5; fi68++ {
                    v61 = v61
                    fmt.Println(v19, v67, v61)
                    _ = fi68
                }
                _ = v67
            }
        default:
            if (!false) {
                v61 += (v61 * (v61 - v61))
            }
        }
        _ = v61
    }
    var v69 string = "a"
    v70 := GenS11{F0: -9223372036854775808}
    var v71 uint16 = 283
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v60:", v60)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v60
    _ = v69
    _ = v70
    _ = v71
}
