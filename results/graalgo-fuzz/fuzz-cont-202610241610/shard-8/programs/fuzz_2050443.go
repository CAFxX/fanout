package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint }
type GenS2 struct { F0 uint16; F1 uint16; F2 int16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint8 }
func genFn0(p0 int64) int32 {
    fmt.Println(p0)
    v7 := []int{279, -302}
    _ = p0
    _ = v7
    return 1
}
func genFn1() uint {
    var v8 uint8 = 98
    if (1 == -9223372036854775808) {
        for fi9 := 0; fi9 < 9; fi9++ {
            v10 := GenS1{F0: -914, F1: -9223372036854775808, F2: 42}
            var v11 int = fi9
            v12 := []int{v11, 2147483647, (v11 | (-(v11)))}
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    _ = v8
    return 1
}

func main() {
    var v1 int32 = 0
    v2 := []uint32{}
    var v3 uint16 = 1
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    if (false || false) {
        var v13 uint = 160
        _ = v13
    } else {
        switch 384 {
        case 0:
            var v14 int64 = -267
            _ = v14
        default:
            v15 := GenS2{F0: v3, F1: 65535, F2: 330}
            _ = v15
        }
        if true {
            var v16 int8 = 1
            _ = v16
        }
    }
    v1 += 42
    for fi17 := 0; fi17 < 5; fi17++ {
        v3 += v3
        var v18 uint64 = 612
        if false {
            delete(v4, 84)
            var v19 int16 = -305
            var v20 int = 0
            _ = v19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if false {
            var v22 int8 = -128
            v23 := GenS3{F0: 1, F1: (fi21 - fi21)}
            _ = v22
            _ = v23
        } else {
            var v24 int32 = (v1 >> ((2) & 63))
            v25 := GenS4{F0: uint64(v3), F1: 354}
            _ = v24
            _ = v25
        }
        if (697 != 1) {
            var v26 int32 = v1
            _ = v26
        } else {
            var v27 uint8 = 42
            switch -1 {
            case 9223372036854775807:
                if true {
                    _ = len(v2)
                } else {
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if true {
                            if (-128 >= -125) {
                                v1 = (v1 / ((genFn0(9223372036854775807)) | 1))
                                var v29 int32 = v1
                                _ = v29
                            } else {
                                switch (fi28 | fi28) {
                                case 2147483647:
                                    var v30 string = "foo bar"
                                    _ = v30
                                default:
                                    if (((1.5 - 3.14159) * 0.1) == 0.1) {
                                        var v31 uint16 = 576
                                        var v32 uint8 = (v27 & v27)
                                        if ((-9223372036854775808 > 42) || false) {
                                            _ = len(v2)
                                            switch int(v1) {
                                            case -2147483648:
                                                v33 := GenS4{F0: 2, F1: 0}
                                                _ = v33
                                            case 42:
                                                switch "foo bar" {
                                                case "\n\t":
                                                    fmt.Println(v32)
                                                case "foo bar":
                                                    v1 -= (v1 | genFn0(1))
                                                default:
                                                    v1 = 42
                                                }
                                            case 128:
                                                fmt.Println(v27, fi21, v27)
                                            }
                                            var v34 float32 = 0.0
                                            _ = v34
                                        }
                                        _ = v31
                                        _ = v32
                                    } else {
                                        var v35 int64 = 849
                                        _ = v35
                                    }
                                }
                            }
                            v1 += 880
                            switch (fi21 - fi28) {
                            case 1:
                                fmt.Println(v1)
                                var v36 int8 = 42
                                v3 -= (v3 & (^(v3)))
                                _ = v36
                            case -2147483648:
                                v37 := GenS5{F0: 927, F1: false}
                                _ = v37
                            case 9223372036854775807:
                                v38 := GenS4{F0: 549, F1: 42}
                                _ = v38
                            default:
                                _ = len(v5)
                            }
                        } else {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v39 int8 = int8(v1)
                            v3 += v3
                            if true {
                                var v40 float64 = (1e10 * (-2.25 - 0.1))
                                var v41 int64 = 9223372036854775807
                                if (0.1 == (-2.25 * 3.14159)) {
                                    v41 -= int64(v39)
                                    var v42 uint8 = 128
                                    _ = v42
                                } else {
                                    v43 := GenS6{F0: 128}
                                    _ = v43
                                }
                                _ = v40
                                _ = v41
                            } else {
                                switch "foo bar" {
                                case "":
                                    var v44 bool = false
                                    _ = v44
                                case "\n\t":
                                    fmt.Println(v39)
                                case "hello":
                                    switch ("foo bar" + (("foo bar" + "") + ("\n\t" + "hello"))) {
                                    case "":
                                        _ = len(v2)
                                    case "a":
                                        if (3.14159 == 3.14159) {
                                            v27 += v27
                                            v45 := GenS2{F0: v3, F1: (^(v3)), F2: -32768}
                                            _ = v45
                                        }
                                    }
                                default:
                                    if false {
                                        v46 := GenS2{F0: (v3 / (((^(v3))) | 1)), F1: 32767, F2: -1}
                                        _ = v46
                                    } else {
                                        if true {
                                            if (!false) {
                                                var v47 int16 = 0
                                                _ = v47
                                            } else {
                                                v1 = -2147483648
                                            }
                                            v48 := v5[1]
                                            _ = v48
                                        }
                                    }
                                }
                                var v49 float32 = -2.25
                                var v50 float64 = 0.1
                                _ = v49
                                _ = v50
                            }
                            v51 := GenS4{F0: 42, F1: 65535}
                            _ = v39
                            _ = v51
                        }
                        v3 += v3
                        switch (-(fi21)) {
                        case 730:
                            var v52 uint32 = 42
                            _ = v52
                        case 376:
                            fmt.Println(v3)
                        case -623:
                            var v53 uint32 = 98
                            _ = v53
                        }
                        _ = fi28
                    }
                    v1 -= v1
                }
            case -221:
                v3 -= v3
            case 2147483647:
                if ((1 < 4294967295) && true) {
                    v3 = (v3 - v3)
                    var v54 int16 = -406
                    _ = v54
                } else {
                    v1 = (v1 & (v1 >> ((65535) & 63)))
                    var v55 int64 = 9223372036854775807
                    v56 := v4[185]
                    _ = v55
                    _ = v56
                }
            }
            v1 -= (v1 & (-(v1)))
            _ = v27
        }
        var v57 int32 = v1
        _ = fi21
        _ = v57
    }
    v58 := v6[947]
    v3 += v3
    for fi59 := 0; fi59 < 6; fi59++ {
        v1 = (-(v1))
        _ = fi59
    }
    fmt.Println(v58, v58)
    v58 = v58
    var v60 bool = true
    for fi61 := 0; fi61 < 7; fi61++ {
        v62 := GenS6{F0: 164}
        _ = fi61
        _ = v62
    }
    v58 = "foo bar"
    switch "a" {
    case "hello":
        delete(v4, 2147483648)
    case "":
        var v63 int = 32767
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v60
}
