package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 uint8; F1 uint64; F2 uint }
type GenS6 struct { F0 uint16 }
func genFn0() int {
    var v7 float64 = float64(((0.0 * 1.5) * (float32(1e10) / 1.0)))
    v7 = v7
    _ = v7
    return 127
}
func genFn1(p0 bool) int16 {
    v8 := []int{580}
    for fi9 := 0; fi9 < 19; fi9++ {
        for fi10 := 0; fi10 < 12; fi10++ {
            p0 = p0
            _ = len(v8)
            switch "foo bar" {
            case "a":
                var v11 uint64 = 0
                _ = v11
            case "\n\t":
                v12 := GenS1{F0: 0}
                _ = v12
            case "hello":
                var v13 uint16 = 122
                _ = v13
            }
            fmt.Println(fi10)
            _ = fi10
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            _ = len(v8)
            fmt.Println(fi14, p0)
            p0 = false
            _ = fi14
        }
        _ = len(v8)
        _ = fi9
    }
    var v15 uint64 = 2
    _ = p0
    _ = v15
    _ = v8
    return -1
}

func main() {
    var v1 int8 = -5
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 bool = false
    v5 := []uint32{2, 42}
    var v6 int = 32768
    v6 = v6
    if (150 < 0) {
        for fi16 := 0; fi16 < 11; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
    }
    fmt.Println(v6, v4, v1)
    if (!((!(760 >= 42)) && (float32(0.0) > 3.14159))) {
        delete(v2, 42)
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 3; fi18++ {
                fmt.Println(fi18)
                _ = len(v5)
                _ = fi18
            }
            _ = fi17
        }
    } else {
        for fi19 := 0; fi19 < 19; fi19++ {
            v4 = ((v4 || v4) && false)
            if ("foo bar" != "a") {
                v6 -= genFn0()
                var v20 uint32 = 42
                _ = v20
            } else {
                switch "\n\t" {
                case "hello":
                    for fi21 := 0; fi21 < 7; fi21++ {
                        v22 := GenS1{F0: 42}
                        _ = len(v2)
                        switch (v6 % (((v6 | genFn0())) | 1)) {
                        case 1:
                            var v23 int8 = -128
                            _ = v23
                        case 0:
                            _ = len(v2)
                        case 9223372036854775807:
                            var v24 int32 = 1
                            _ = v24
                        }
                        _ = fi21
                        _ = v22
                    }
                default:
                    if v4 {
                        for fi25 := 0; fi25 < 15; fi25++ {
                            if v4 {
                                v26 := GenS1{F0: 188}
                                _ = v26
                            } else {
                                _ = len(v3)
                                var v27 uint8 = 213
                                var v28 uint32 = 65535
                                _ = v27
                                _ = v28
                            }
                            if (255 != 127) {
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                            } else {
                                v29 := GenS1{F0: 0}
                                switch "hello" {
                                case "foo bar":
                                    fmt.Println(v4, fi25)
                                }
                                delete(v3, (("" + "foo bar") + "hello"))
                                fmt.Println(v1, fi19)
                                if (0 >= -1) {
                                    if (v4 || (v4 && true)) {
                                        v2[2147483647] = 42
                                        var v30 int = genFn0()
                                        var v31 float64 = -2.25
                                        _ = v30
                                        _ = v31
                                    } else {
                                        var v32 int8 = v1
                                        var v33 bool = (v4 || ((genFn0() != genFn0()) || v4))
                                        _ = v32
                                        _ = v33
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v6 += fi19
                                    v2[2147483648] = 1
                                    var v34 int8 = (v1 ^ (^(v1)))
                                    var v35 int64 = -816
                                    if ((3.14159 * (float64(0.1) / 1.5)) <= (-2.25 - 0.1)) {
                                        delete(v2, 65535)
                                    } else {
                                        fmt.Println(v1, v34, v6)
                                        if len(v5) > 0 {
                                            v5[0] = 65535
                                        }
                                    }
                                    _ = v34
                                    _ = v35
                                }
                                var v36 float32 = 0.0
                                if true {
                                    v3[("hello" + ("\n\t" + "a"))] = v4
                                    var v37 uint64 = 90
                                    if v4 {
                                        v38 := GenS2{F0: 334, F1: 0, F2: 345}
                                        _ = v38
                                    } else {
                                        v36 = v36
                                    }
                                    _ = v37
                                }
                                v3["foo bar"] = ((!false) || (!false))
                                _ = v29
                                _ = v36
                            }
                            v1 += v1
                            _ = fi25
                        }
                        fmt.Println(v6)
                    } else {
                        for fi39 := 0; fi39 < 9; fi39++ {
                            if v4 {
                                v40 := GenS3{F0: v1}
                                var v41 uint64 = 1
                                switch (fi39 * (v6 & (fi19 - -876))) {
                                case 346:
                                    if v4 {
                                        var v42 uint64 = v41
                                        var v43 int = 4294967295
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v44 := GenS1{F0: v41}
                                        _ = v44
                                    }
                                    if (v1 == 100) {
                                        var v45 uint64 = (v41 >> ((0) & 63))
                                        fmt.Println(fi39, fi19)
                                        v4 = v4
                                        switch fi19 {
                                        case 128:
                                            _ = len(v3)
                                        case 9223372036854775807:
                                            fmt.Println(fi39, v1, v4)
                                        case 645:
                                            v46 := GenS3{F0: v1}
                                            _ = v46
                                        default:
                                            v47 := GenS2{F0: 0, F1: -269, F2: -1}
                                            _ = v47
                                        }
                                        if v4 {
                                            var v48 int8 = v1
                                            v1 = (^(v1))
                                            _ = v48
                                        }
                                        _ = v45
                                    }
                                    if v4 {
                                        if len(v5) > 0 {
                                            v5[0] = 2147483647
                                        }
                                    } else {
                                        fmt.Println(fi39, fi19, v1)
                                        var v49 int16 = genFn1((475 != -771))
                                        if v4 {
                                            v50 := GenS4{F0: "a", F1: (v41 - (v41 << ((688) & 63)))}
                                            _ = v50
                                        }
                                        _ = v49
                                    }
                                }
                                _ = v40
                                _ = v41
                            } else {
                                var v51 float64 = 0.0
                                switch (("" + ("\n\t" + "")) + ("foo bar" + ("foo bar" + "foo bar"))) {
                                case "foo bar":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "":
                                    v4 = (2 == 72)
                                }
                                fmt.Println(fi39, fi39)
                                if v4 {
                                    if len(v5) > 0 {
                                        v5[0] = 2
                                    }
                                } else {
                                    v52 := GenS2{F0: 424, F1: 106, F2: 32767}
                                    _ = v52
                                }
                                _ = v51
                            }
                            v53 := v2[65535]
                            _ = fi39
                            _ = v53
                        }
                    }
                }
                for fi54 := 0; fi54 < 6; fi54++ {
                    _ = len(v5)
                    _ = fi54
                }
                switch "hello" {
                case "hello":
                    fmt.Println(fi19)
                case "\n\t":
                    v55 := GenS1{F0: 327}
                    _ = v55
                default:
                    fmt.Println(fi19)
                }
            }
            v6 = v6
            _ = fi19
        }
        var v56 int64 = 1
        if v4 {
            fmt.Println(v6)
            if (!(1 >= 4294967295)) {
                var v57 float64 = 1.5
                fmt.Println(v56)
                _ = v57
            } else {
                if (v56 != v56) {
                    var v58 float32 = (1e10 / float32((1e10 + ((((float64(0.1) + 1e10) + 1.5) - -2.25) - (0.1 / 0.1)))))
                    v58 = 0.1
                    _ = v58
                }
            }
        } else {
            v1 = v1
            for fi59 := 0; fi59 < 3; fi59++ {
                if (-27 <= 1) {
                    var v60 int64 = -9223372036854775808
                    if false {
                        if v4 {
                            _ = len(v5)
                            v61 := v3[""]
                            var v62 uint8 = 127
                            _ = v61
                            _ = v62
                        } else {
                            if (-2.25 >= (float32(3.14159) / 1.5)) {
                                v63 := GenS5{F0: 128, F1: 931, F2: 255}
                                _ = v63
                            }
                            v64 := v2[70]
                            for fi65 := 0; fi65 < 2; fi65++ {
                                fmt.Println(v60, fi59)
                                switch "hello" {
                                case "":
                                    v2[v64] = (^(v64))
                                case "a":
                                    var v66 uint = 18446744073709551615
                                    _ = v66
                                default:
                                    v3["foo bar"] = v4
                                }
                                v2[(v64 << ((4294967295) & 63))] = 2147483648
                                _ = fi65
                            }
                            _ = v64
                        }
                    }
                    _ = v60
                }
                switch "" {
                case "":
                    if v4 {
                        v67 := GenS1{F0: 1}
                        v68 := GenS6{F0: 255}
                        _ = v67
                        _ = v68
                    } else {
                        v69 := GenS5{F0: 128, F1: 237, F2: 0}
                        v2[456] = 2147483647
                        _ = v69
                    }
                case "\n\t":
                    var v70 int64 = v56
                    _ = v70
                case "hello":
                    v1 -= v1
                }
                _ = fi59
            }
            v4 = v4
        }
        _ = v56
    }
    fmt.Println(v4)
    fmt.Println(v4, v6, v6)
    v6 += (v6 << ((1) & 63))
    fmt.Println(v6)
    fmt.Println(v1, v1, v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
