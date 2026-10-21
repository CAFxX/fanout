package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 bool; F1 int32 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int; F1 int64 }
func genFn0() int {
    var v6 string = (("\n\t" + "foo bar") + ("\n\t" + "a"))
    if false {
        for fi7 := 0; fi7 < 4; fi7++ {
            v6 = ("\n\t" + "a")
            fmt.Println(v6, v6)
            v8 := GenS1{F0: 0}
            _ = fi7
            _ = v8
        }
    } else {
        v6 = v6
        v9 := []int{128}
        _ = v9
    }
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 uint64, p2 uint64) uint {
    var v10 uint32 = 1
    switch "\n\t" {
    case "\n\t":
        v11 := GenS2{F0: 127, F1: 32767}
        _ = v11
    }
    v12 := []int{9223372036854775807, 32768}
    p1 = p2
    v13 := GenS3{F0: true, F1: 100}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v12
    _ = v13
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 2
    var v3 int32 = 2147483647
    var v4 uint64 = 966
    var v5 uint16 = (^(v2))
    v4 = v4
    v14 := GenS1{F0: v5}
    for fi15 := 0; fi15 < 20; fi15++ {
        for fi16 := 0; fi16 < 18; fi16++ {
            v17 := GenS4{F0: (!(genFn1(-939, (v4 / (((v4 >> ((genFn1(42, v4, v4)) & 63))) | 1)), (v4 ^ v4)) >= 42)), F1: -128}
            _ = fi16
            _ = v17
        }
        var v18 string = "a"
        var v19 uint16 = 42
        _ = fi15
        _ = v18
        _ = v19
    }
    v20 := v1[2]
    if true {
        for fi21 := 0; fi21 < 18; fi21++ {
            var v22 uint8 = 2
            switch ("foo bar" + (("foo bar" + "hello") + "a")) {
            case "foo bar":
                var v23 int32 = v3
                _ = v23
            }
            _ = fi21
            _ = v22
        }
        v24 := GenS3{F0: true, F1: v3}
        if true {
            var v25 int32 = v3
            switch 128 {
            case 1:
                var v26 uint64 = uint64(v2)
                _ = v26
            case 380:
                if true {
                    fmt.Println(v4, v2, v4)
                    if (0 >= -32768) {
                        switch 128 {
                        case 32768:
                            v27 := GenS5{F0: genFn0(), F1: 42}
                            _ = v27
                        case -77:
                            switch 2147483647 {
                            case 1:
                                v28 := GenS3{F0: (1 < -1), F1: (v3 / (((^(v25))) | 1))}
                                _ = v28
                            case 0:
                                v20 -= uint32(v4)
                            case 4294967295:
                                v5 -= v2
                            }
                        case -101:
                            var v29 uint = 255
                            _ = v29
                        }
                    } else {
                        v4 += (^(v4))
                        if (1 != -195) {
                            fmt.Println(v20, v5)
                            var v30 int8 = 100
                            _ = v30
                        } else {
                            var v31 uint = genFn1(-32768, (v4 >> ((0) & 63)), v4)
                            var v32 int = 32767
                            _ = v31
                            _ = v32
                        }
                    }
                    switch 9223372036854775807 {
                    case 127:
                        fmt.Println(v3, v4, v4)
                    case 4294967295:
                        var v33 uint8 = 127
                        _ = v33
                    case 0:
                        var v34 int64 = 42
                        _ = v34
                    default:
                        if true {
                            for fi35 := 0; fi35 < 9; fi35++ {
                                var v36 float32 = ((-2.25 - 0.0) * 0.1)
                                var v37 int64 = 42
                                var v38 bool = false
                                _ = fi35
                                _ = v36
                                _ = v37
                                _ = v38
                            }
                            if (0 != (v4 + (v4 << ((genFn1(-745, 2, (v4 - (^(v4))))) & 63)))) {
                                var v39 int = 42
                                if ((^(v20)) > v20) {
                                    for fi40 := 0; fi40 < 19; fi40++ {
                                        if (!(2 == 42)) {
                                            var v41 uint64 = (^(v4))
                                            v25 = -1
                                            _ = v41
                                        } else {
                                            fmt.Println(v3, v20, v39)
                                            var v42 uint = 583
                                            fmt.Println(v39, v39, v5)
                                            _ = v42
                                        }
                                        var v43 float64 = (float64(1e10) / 1.0)
                                        v44 := v1[(v20 | (v20 - (v20 & 795)))]
                                        _ = fi40
                                        _ = v43
                                        _ = v44
                                    }
                                    v45 := []int{genFn0(), v39, genFn0(), (v39 * (v39 >> ((574) & 63)))}
                                    _ = v45
                                } else {
                                    var v46 float32 = (0.1 - 3.14159)
                                    var v47 string = ("\n\t" + "foo bar")
                                    fmt.Println(v2, v47)
                                    _ = v46
                                    _ = v47
                                }
                                _ = v39
                            } else {
                                if ((v4 & 464) == 4294967295) {
                                    v5 = v2
                                }
                                var v48 int16 = 0
                                switch ("" + "foo bar") {
                                case "a":
                                    v25 += -1
                                case "hello":
                                    fmt.Println(v3, v4, v48)
                                default:
                                    fmt.Println(v4, v2, v20)
                                }
                                _ = v48
                            }
                        } else {
                            for fi49 := 0; fi49 < 6; fi49++ {
                                if ("a" <= "foo bar") {
                                    switch 127 {
                                    case 32768:
                                        delete(v1, (v20 + 65535))
                                    case 42:
                                        switch "a" {
                                        case "a":
                                            var v50 uint16 = (v5 * 392)
                                            _ = v50
                                        }
                                    default:
                                        var v51 int = (fi49 - 9223372036854775807)
                                        _ = v51
                                    }
                                } else {
                                    for fi52 := 0; fi52 < 10; fi52++ {
                                        v2 += v5
                                        _ = fi52
                                    }
                                    var v53 float32 = float32(0.1)
                                    fmt.Println(v20, v5, fi49)
                                    _ = v53
                                }
                                var v54 uint64 = 0
                                for fi55 := 0; fi55 < 14; fi55++ {
                                    var v56 float32 = ((0.0 + (((1e10 / 3.14159) - 1.5) * 0.1)) + (3.14159 * 1e10))
                                    _ = fi55
                                    _ = v56
                                }
                                _ = fi49
                                _ = v54
                            }
                            v57 := GenS4{F0: false, F1: -1}
                            v58 := GenS5{F0: -1, F1: -9223372036854775808}
                            _ = v57
                            _ = v58
                        }
                    }
                } else {
                    v59 := GenS3{F0: false, F1: (v25 * -243)}
                    _ = v59
                }
            default:
                var v60 bool = false
                _ = v60
            }
            _ = v25
        } else {
            v1[v20] = v20
        }
        _ = v24
    }
    if true {
        switch int(v4) {
        case -9223372036854775808:
            fmt.Println(v2, v5, v5)
        case -981:
            v5 += (v2 & v5)
        case 9223372036854775807:
            v61 := []int{genFn0(), 4294967295, genFn0(), 128}
            _ = v61
        default:
            var v62 int64 = 9223372036854775807
            _ = v62
        }
    } else {
        v2 += 874
    }
    for fi63 := 0; fi63 < 18; fi63++ {
        v64 := GenS4{F0: false, F1: 74}
        _ = fi63
        _ = v64
    }
    if (32767 < 0) {
        var v65 float64 = 0.1
        v5 += v2
        for fi66 := 0; fi66 < 15; fi66++ {
            for fi67 := 0; fi67 < 13; fi67++ {
                v1[uint32(v4)] = v20
                _ = fi67
            }
            _ = fi66
        }
        _ = v65
    } else {
        var v68 uint = 65535
        _ = v68
    }
    if (("\n\t" + ("" + ((("\n\t" + (("hello" + ("hello" + "hello")) + ("foo bar" + ""))) + "hello") + "\n\t"))) < "hello") {
        v2 = v5
    }
    switch "foo bar" {
    case "a":
        var v69 float32 = 1e10
        _ = v69
    case "":
        var v70 string = "hello"
        _ = v70
    default:
        v20 = (v20 - (v20 + (^(v20))))
    }
    fmt.Println(v4)
    var v71 uint16 = v5
    fmt.Println(v3, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v71
}
