package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS2 struct { F0 int64; F1 int; F2 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int16; F2 uint16 }
type GenS7 struct { F0 uint }
func genFn0(p0 int8) int64 {
    switch -9223372036854775808 {
    case -1:
        var v7 int8 = (p0 * (p0 << ((18446744073709551615) & 63)))
        _ = v7
    case -749:
        v8 := GenS1{F0: 128, F1: 4294967295, F2: 42}
        _ = v8
    }
    var v9 int64 = -9223372036854775808
    _ = p0
    _ = v9
    return 9223372036854775807
}
func genFn1() uint {
    var v10 uint16 = 32767
    v10 -= v10
    _ = v10
    return 65535
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    var v3 uint64 = 2
    var v4 uint64 = (v3 & v3)
    v5 := make(map[int]uint32)
    var v6 uint64 = (v4 * (v4 * v3))
    v11 := GenS1{F0: 217, F1: (^(v4)), F2: v4}
    v4 -= 468
    var v12 int = 9223372036854775807
    if true {
        if (-9223372036854775808 != 1) {
            for fi13 := 0; fi13 < 7; fi13++ {
                v6 -= v6
                var v14 uint = 2
                v12 = 127
                _ = fi13
                _ = v14
            }
            if false {
                if (!false) {
                    var v15 bool = true
                    var v16 uint64 = (v4 << ((4294967295) & 63))
                    _ = v15
                    _ = v16
                }
                for fi17 := 0; fi17 < 1; fi17++ {
                    for fi18 := 0; fi18 < 15; fi18++ {
                        v3 = v4
                        _ = fi18
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v20 := GenS2{F0: -798, F1: 9223372036854775807, F2: (v1 / ((v1) | 1))}
                        _ = fi19
                        _ = v20
                    }
                    _ = fi17
                }
                v21 := []int{}
                _ = v21
            } else {
                if ((v4 >> ((genFn1()) & 63)) != (^(v3))) {
                    var v22 float64 = (0.0 - 3.14159)
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v12 = fi23
                        var v24 bool = false
                        if (v24 || v24) {
                            v12 = -446
                            switch "hello" {
                            case "a":
                                v24 = (896 >= 255)
                            case "\n\t":
                                var v25 uint = 42
                                _ = v25
                            default:
                                v22 += v22
                            }
                            v6 = (v3 & (v6 - v3))
                        } else {
                            fmt.Println(v22)
                        }
                        _ = fi23
                        _ = v24
                    }
                    if (!(0 <= -9223372036854775808)) {
                        v26 := []int{2147483647, 128, (-(v12)), (-(v12))}
                        v27 := GenS2{F0: -9223372036854775808, F1: (v12 / (((v12 ^ (v12 / ((-1) | 1)))) | 1)), F2: v1}
                        _ = v26
                        _ = v27
                    }
                    _ = v22
                } else {
                    v4 = v3
                    v1 -= 424
                }
            }
        } else {
            for fi28 := 0; fi28 < 13; fi28++ {
                v6 = v4
                v29 := []int{v12, 32768}
                if (3.14159 < (1.5 + -2.25)) {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v31 := GenS2{F0: genFn0(14), F1: fi30, F2: 2}
                        v3 = 2
                        _ = fi30
                        _ = v31
                    }
                    switch "" {
                    case "hello":
                        v6 += v4
                    case "foo bar":
                        if false {
                            if (!false) {
                                if (1 > -1) {
                                    v32 := GenS1{F0: 1, F1: v4, F2: (^(v3))}
                                    if (712 >= genFn1()) {
                                        v6 += (v6 % ((4294967295) | 1))
                                    } else {
                                        for fi33 := 0; fi33 < 20; fi33++ {
                                            var v34 int64 = 9223372036854775807
                                            _ = fi33
                                            _ = v34
                                        }
                                        v6 = (v4 ^ v4)
                                        v35 := GenS1{F0: 2, F1: (v3 << ((255) & 63)), F2: 18446744073709551615}
                                        _ = v35
                                    }
                                    switch fi28 {
                                    case 32768:
                                        v36 := GenS3{F0: (fi28 ^ int(v1)), F1: v12}
                                        _ = v36
                                    default:
                                        var v37 int16 = int16(v3)
                                        _ = v37
                                    }
                                    _ = v32
                                } else {
                                    for fi38 := 0; fi38 < 13; fi38++ {
                                        switch (v12 / (((fi28 - fi38)) | 1)) {
                                        case 286:
                                            if false {
                                                v39 := GenS4{F0: 32767, F1: 1, F2: (v1 | (v1 - v1))}
                                                v40 := GenS2{F0: 221, F1: fi38, F2: v1}
                                                var v41 bool = (true && false)
                                                _ = v39
                                                _ = v40
                                                _ = v41
                                            } else {
                                                var v42 int8 = 127
                                                fmt.Println(v1, v42)
                                                _ = v42
                                            }
                                        case 32768:
                                            v12 += (v12 ^ fi38)
                                        }
                                        switch fi28 {
                                        case -553:
                                            switch (fi28 / (((^(fi38))) | 1)) {
                                            case -626:
                                                if (!(-265 > 591)) {
                                                    v4 += (v4 >> ((65535) & 63))
                                                    var v43 uint16 = v1
                                                    _ = v43
                                                }
                                            }
                                        }
                                        _ = fi38
                                    }
                                }
                                v12 = fi28
                                if (((("a" + "foo bar") + "hello") + ("\n\t" + "a")) >= ("hello" + (("a" + (("" + "hello") + "foo bar")) + ""))) {
                                    if (1e10 == 1.5) {
                                        v44 := GenS2{F0: -1, F1: fi28, F2: (^(v1))}
                                        switch 1 {
                                        case -324:
                                            var v45 float32 = 1e10
                                            _ = v45
                                        case 42:
                                            v6 += (v6 / ((v4) | 1))
                                        case -92:
                                            var v46 uint8 = 127
                                            _ = v46
                                        }
                                        _ = v44
                                    }
                                } else {
                                    if false {
                                        var v47 int16 = -1
                                        _ = v47
                                    }
                                    v48 := GenS1{F0: 255, F1: 1, F2: v6}
                                    if (false && true) {
                                        for fi49 := 0; fi49 < 19; fi49++ {
                                            var v50 float32 = ((1e10 / ((1e10 + 0.1) - (3.14159 - 1.5))) - 1e10)
                                            _ = fi49
                                            _ = v50
                                        }
                                        fmt.Println(v1, fi28, v4)
                                        v51 := GenS4{F0: 32767, F1: -128, F2: uint16(v3)}
                                        _ = v51
                                    }
                                    _ = v48
                                }
                            } else {
                                var v52 float64 = 0.1
                                v53 := GenS5{F0: -676}
                                _ = v52
                                _ = v53
                            }
                        }
                    case "a":
                        v1 += v1
                    default:
                        v1 = 392
                    }
                }
                _ = fi28
                _ = v29
            }
            v54 := GenS6{F0: -86, F1: -32768, F2: (v1 * v1)}
            _ = v54
        }
        for fi55 := 0; fi55 < 15; fi55++ {
            var v56 uint = 303
            v6 += (v4 - 42)
            fmt.Println(v4, v4)
            _ = fi55
            _ = v56
        }
        var v57 uint64 = 684
        _ = v57
    }
    v58 := v2[2147483648]
    switch v12 {
    case 381:
        var v59 uint = 2
        _ = v59
    case -9223372036854775808:
        v60 := GenS2{F0: 9223372036854775807, F1: (v12 ^ (v12 & (v12 << ((genFn1()) & 63)))), F2: v1}
        _ = v60
    case 127:
        if v58 {
            if (-2.25 < 1e10) {
                delete(v5, (^(v12)))
                var v61 int64 = 1
                _ = v61
            }
            v12 = (-(v12))
        } else {
            for fi62 := 0; fi62 < 10; fi62++ {
                for fi63 := 0; fi63 < 19; fi63++ {
                    var v64 float64 = 1e10
                    v65 := GenS2{F0: 0, F1: v12, F2: (^(v1))}
                    _ = fi63
                    _ = v64
                    _ = v65
                }
                _ = fi62
            }
            v6 = 42
            v12 = -1
        }
    default:
        switch "a" {
        case "":
            v12 = (v12 ^ -490)
        case "foo bar":
            var v66 int16 = -32768
            _ = v66
        case "a":
            var v67 int = 32767
            _ = v67
        }
    }
    if true {
        fmt.Println(v58, v58)
    }
    switch (("foo bar" + "hello") + "a") {
    case "foo bar":
        v68 := GenS5{F0: 1}
        _ = v68
    case "hello":
        v69 := GenS7{F0: 601}
        _ = v69
    }
    v2[65535] = false
    fmt.Println(v58, v6)
    var v70 uint8 = 2
    var v71 int64 = genFn0(-126)
    for fi72 := 0; fi72 < 9; fi72++ {
        v71 = (^(v71))
        v73 := []int{(v12 * -9223372036854775808), v12}
        v71 += genFn0(42)
        _ = fi72
        _ = v73
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v70
    _ = v71
}
