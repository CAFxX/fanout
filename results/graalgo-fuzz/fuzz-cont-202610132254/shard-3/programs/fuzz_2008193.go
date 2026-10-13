package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int16 }
func genFn0(p0 bool, p1 uint) int64 {
    if (false && p0) {
        var v7 int = 861
        if p0 {
            v8 := GenS1{F0: p1, F1: 42}
            _ = v8
        } else {
            var v9 int16 = -44
            _ = v9
        }
        _ = v7
    } else {
        switch ((("foo bar" + (("\n\t" + "") + ("" + ("a" + (("\n\t" + "a") + ("a" + "a")))))) + "\n\t") + "foo bar") {
        case "hello":
            var v10 int16 = 1000
            _ = v10
        case "\n\t":
            fmt.Println(p1, p0)
        case "a":
            p1 += p1
        default:
            if p0 {
                if p0 {
                    var v11 uint16 = 461
                    fmt.Println(p1, v11)
                    _ = v11
                } else {
                    p0 = true
                }
                v12 := []int{32767, -457, -350}
                _ = v12
            }
        }
        v13 := []int{128, 9223372036854775807, 127, 4294967295}
        p1 -= (p1 >> ((183) & 63))
        _ = v13
    }
    p0 = false
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{42, 903}
    var v3 int16 = -32768
    var v4 uint8 = 1
    var v5 uint32 = 627
    v6 := make(map[string]bool)
    if (!false) {
        var v14 string = ("a" + ((("hello" + "foo bar") + (("\n\t" + "hello") + "")) + ""))
        _ = v14
    } else {
        if (v5 != (v5 ^ v5)) {
            v15 := GenS1{F0: 2, F1: (v4 * (v4 >> ((4294967295) & 63)))}
            if ((v5 * v5) > (v5 * uint32(v4))) {
                v3 = v3
                var v16 uint32 = 59
                _ = v16
            } else {
                v4 = v4
                v1[v5] = ((("a" + ("" + "\n\t")) + ((("\n\t" + "\n\t") + "hello") + "hello")) + "")
            }
            _ = v15
        } else {
            switch (("" + ("hello" + ("foo bar" + "foo bar"))) + "foo bar") {
            case "foo bar":
                if false {
                    v3 = v3
                } else {
                    fmt.Println(v4)
                    for fi17 := 0; fi17 < 15; fi17++ {
                        switch ("" + ("hello" + "foo bar")) {
                        case "\n\t":
                            v5 -= (v5 - v5)
                        }
                        if (354 < 89) {
                            var v18 float64 = (float64(0.0) - 0.0)
                            fmt.Println(v5, v3)
                            for fi19 := 0; fi19 < 20; fi19++ {
                                var v20 int8 = -10
                                _ = fi19
                                _ = v20
                            }
                            _ = v18
                        } else {
                            var v21 uint8 = (v4 / ((v4) | 1))
                            _ = v21
                        }
                        _ = fi17
                    }
                }
            default:
                for fi22 := 0; fi22 < 13; fi22++ {
                    fmt.Println(v3, v3)
                    v3 += (v3 << ((uint(v4)) & 63))
                    v23 := GenS2{F0: (^(v5)), F1: int32(v3), F2: 2147483647}
                    _ = fi22
                    _ = v23
                }
            }
            for fi24 := 0; fi24 < 2; fi24++ {
                if false {
                    v25 := GenS2{F0: (v5 - (v5 % ((v5) | 1))), F1: 42, F2: 42}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 11; fi26++ {
                        switch (fi26 - (fi24 / ((fi24) | 1))) {
                        case 4294967295:
                            var v27 uint32 = v5
                            _ = v27
                        case 127:
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case 862:
                            if (2 > 293) {
                                var v28 int = (fi24 - 127)
                                v4 = (v4 ^ v4)
                                _ = v28
                            } else {
                                v3 = -1
                                v29 := GenS1{F0: 42, F1: (v4 ^ v4)}
                                _ = v29
                            }
                        }
                        _ = fi26
                    }
                    if false {
                        fmt.Println(fi24)
                        if false {
                            var v30 int8 = -42
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v30
                        }
                        v4 -= (v4 + (v4 * v4))
                    } else {
                        v4 -= 128
                        delete(v6, "\n\t")
                        var v31 float32 = 3.14159
                        _ = v31
                    }
                    switch "" {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "a":
                        var v32 uint32 = v5
                        _ = v32
                    }
                }
                _ = fi24
            }
        }
    }
    for fi33 := 0; fi33 < 17; fi33++ {
        v2 = append(v2, v5)
        _ = fi33
    }
    v34 := GenS1{F0: 65535, F1: 42}
    var v35 bool = (-889 < 42)
    fmt.Println(v4)
    var v36 int16 = v3
    if ((!v35) || (159 >= 4294967295)) {
        var v37 string = ("a" + "")
        v4 += 128
        v5 = v5
        _ = v37
    } else {
        _ = len(v2)
        fmt.Println(v36)
    }
    if (!v35) {
        var v38 float32 = float32(0.0)
        _ = v38
    }
    fmt.Println(v35, v35)
    for fi39 := 0; fi39 < 16; fi39++ {
        switch fi39 {
        case 1:
            for fi40 := 0; fi40 < 13; fi40++ {
                var v41 int = (fi40 * -1)
                v4 = (v4 ^ (^(v4)))
                v42 := GenS2{F0: 0, F1: 693, F2: 2147483647}
                var v43 uint64 = 18446744073709551615
                v41 = 4294967295
                _ = fi40
                _ = v41
                _ = v42
                _ = v43
            }
        case 843:
            for fi44 := 0; fi44 < 11; fi44++ {
                switch fi39 {
                case 0:
                    if ((v36 == (v3 & (^(v3)))) && true) {
                        if (!(!v35)) {
                            v4 = v4
                            var v45 uint16 = 255
                            _ = v45
                        } else {
                            v3 = -208
                            if v35 {
                                v35 = (2 > 152)
                                delete(v1, (^(v5)))
                            }
                            v36 = v3
                        }
                    } else {
                        switch fi44 {
                        case 1:
                            if (!v35) {
                                var v46 int32 = 659
                                v47 := GenS3{F0: 1, F1: (v46 - (^(v46)))}
                                if v35 {
                                    switch 2147483647 {
                                    case 142:
                                        switch (fi39 + (^(fi44))) {
                                        case 317:
                                            fmt.Println(v46)
                                        default:
                                            v6["foo bar"] = (!(v35 && (4294967295 <= 4294967295)))
                                        }
                                    case 9223372036854775807:
                                        var v48 uint8 = v4
                                        _ = v48
                                    }
                                    if v35 {
                                        _ = len(v1)
                                    } else {
                                        v35 = v35
                                        switch (fi44 >> ((1) & 63)) {
                                        case 0:
                                            var v49 float64 = 1.5
                                            _ = v49
                                        case -708:
                                            v50 := GenS1{F0: 962, F1: v4}
                                            _ = v50
                                        case 1:
                                            fmt.Println(v35, v3, v46)
                                        default:
                                            v36 -= int16(v46)
                                        }
                                        switch (fi44 % (((-(fi44))) | 1)) {
                                        case 128:
                                            v5 = (v5 - 939)
                                        case 32767:
                                            v51 := GenS4{F0: v5, F1: 42}
                                            _ = v51
                                        case -594:
                                            var v52 int64 = -1
                                            _ = v52
                                        default:
                                            v36 = 349
                                        }
                                    }
                                } else {
                                    v53 := GenS3{F0: 42, F1: v46}
                                    var v54 string = ("" + "")
                                    _ = v53
                                    _ = v54
                                }
                                _ = v46
                                _ = v47
                            }
                        default:
                            v4 = (v4 - v4)
                        }
                        var v55 uint32 = (v5 * v5)
                        var v56 uint = 417
                        var v57 float64 = -2.25
                        fmt.Println(v4, v57, fi39)
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                default:
                    if (!(0 != 539)) {
                        var v58 uint16 = 0
                        var v59 int32 = 167
                        _ = v58
                        _ = v59
                    }
                }
                _ = fi44
            }
        }
        _ = fi39
    }
    var v60 int = -525
    v36 = (^(v3))
    v61 := GenS5{F0: int16(v5)}
    v4 -= (v4 % (((^(v4))) | 1))
    if false {
        v5 += (v5 + v5)
        v3 = v36
        for fi62 := 0; fi62 < 17; fi62++ {
            v63 := GenS5{F0: v3}
            _ = fi62
            _ = v63
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
