package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS6 struct { F0 string }
func genFn0(p0 uint64, p1 int8) uint8 {
    v5 := []int{-9223372036854775808}
    v5 = append(v5, 499)
    p1 += (p1 - 83)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            switch fi6 {
            case 0:
                p1 += (p1 << ((829) & 63))
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v8 uint32 = 42
            _ = fi7
            _ = v8
        }
        if ((^(p0)) == p0) {
            for fi9 := 0; fi9 < 6; fi9++ {
                v10 := GenS1{F0: (p0 / ((4294967295) | 1)), F1: 901}
                _ = fi9
                _ = v10
            }
            p0 = (p0 ^ 41)
        } else {
            _ = len(v5)
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 2
}

func main() {
    var v1 float64 = (((0.1 / 1e10) / 0.1) + -2.25)
    v2 := make(map[string]int)
    v3 := []uint32{4294967295, 65535}
    v4 := make(map[string]int)
    v11 := GenS1{F0: 4294967295, F1: 1}
    fmt.Println(v1)
    v3 = append(v3, 2147483648)
    var v12 string = (("a" + ("" + "hello")) + "foo bar")
    if true {
        v13 := GenS1{F0: 4294967295, F1: 494}
        _ = v13
    } else {
        fmt.Println(v12)
        for fi14 := 0; fi14 < 2; fi14++ {
            _ = len(v4)
            v15 := v2["\n\t"]
            _ = fi14
            _ = v15
        }
    }
    v16 := GenS1{F0: 18446744073709551615, F1: 42}
    switch "hello" {
    case "foo bar":
        v3 = append(v3, 4294967295)
    case "":
        if (false || false) {
            if (true || true) {
                var v17 uint32 = 42
                _ = v17
            } else {
                for fi18 := 0; fi18 < 9; fi18++ {
                    _ = len(v2)
                    v1 += (1e10 * (v1 + v1))
                    var v19 float64 = 1e10
                    _ = fi18
                    _ = v19
                }
            }
            v20 := GenS1{F0: 42, F1: 0}
            if (!false) {
                var v21 int64 = 776
                v1 -= (float64(1e10) * (float64(-2.25) / ((3.14159 * ((float64(1e10) - v1) / (0.0 + v1))) + (v1 / (v1 / v1)))))
                _ = v21
            }
            _ = v20
        } else {
            v1 += v1
            v1 -= 0.1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case "hello":
        if (0 <= -9223372036854775808) {
            v22 := GenS1{F0: 2, F1: 582}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi23 := 0; fi23 < 17; fi23++ {
                v24 := GenS2{F0: 0}
                var v25 uint8 = 217
                var v26 int64 = 42
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            _ = v22
        } else {
            var v27 uint64 = 462
            var v28 int64 = 42
            var v29 string = (v12 + v12)
            _ = v27
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        fmt.Println(fi30, v1, fi30)
        switch fi30 {
        case 267:
            v31 := GenS2{F0: 0}
            _ = v31
        }
        if (fi30 <= fi30) {
            var v32 uint32 = 4294967295
            if (true && ("\n\t" <= "\n\t")) {
                for fi33 := 0; fi33 < 2; fi33++ {
                    v34 := GenS1{F0: 59, F1: 210}
                    v1 = v1
                    _ = fi33
                    _ = v34
                }
                v32 = v32
                var v35 uint8 = 0
                _ = v35
            }
            fmt.Println(v1, v1)
            _ = v32
        } else {
            var v36 uint32 = 2147483647
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 float64 = ((v1 + (1.5 - v1)) + 3.14159)
                _ = fi37
                _ = v38
            }
            _ = v36
        }
        _ = fi30
    }
    for fi39 := 0; fi39 < 6; fi39++ {
        var v40 int64 = 42
        var v41 int64 = v40
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println(v12, v1)
    switch -389 {
    case -2147483648:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case 4294967295:
        for fi42 := 0; fi42 < 8; fi42++ {
            switch (fi42 - (fi42 / ((fi42) | 1))) {
            case -9223372036854775808:
                for fi43 := 0; fi43 < 8; fi43++ {
                    fmt.Println(fi43, fi43)
                    v12 = "foo bar"
                    if (true || (true && (false || (false || (false && (2147483647 > -1)))))) {
                        switch ("hello" + v12) {
                        case "hello":
                            v44 := GenS2{F0: 42}
                            _ = v44
                        default:
                            switch v12 {
                            case "a":
                                fmt.Println(fi42)
                            default:
                                v45 := v2[((("" + ("" + v12)) + v12) + v12)]
                                _ = v45
                            }
                        }
                    } else {
                        if true {
                            switch ("a" + v12) {
                            case "a":
                                v46 := GenS3{F0: false, F1: 43, F2: 24}
                                _ = v46
                            case "\n\t":
                                var v47 int8 = -82
                                _ = v47
                            }
                        }
                        v12 = (v12 + ((v12 + v12) + "hello"))
                        if false {
                            if (32767 == -32768) {
                                v48 := GenS4{F0: 42}
                                _ = v48
                            }
                            fmt.Println(fi43, v1)
                        }
                    }
                    if false {
                        v1 = float64(((0.1 + (0.1 + 1e10)) * 0.1))
                        fmt.Println(v1, fi42, v1)
                        v1 -= 0.0
                        v49 := GenS2{F0: -644}
                        _ = v49
                    }
                    var v50 int32 = int32(fi43)
                    if ((v50 < -784) && false) {
                        delete(v4, "\n\t")
                    }
                    v51 := GenS1{F0: 0, F1: 255}
                    _ = fi43
                    _ = v50
                    _ = v51
                }
            case 9223372036854775807:
                v1 = v1
            case 127:
                if true {
                    var v52 bool = false
                    fmt.Println(fi42)
                    _ = v52
                } else {
                    if true {
                        var v53 string = v12
                        if false {
                            v54 := GenS2{F0: 1}
                            v55 := GenS5{F0: -1, F1: 773, F2: 2}
                            _ = v54
                            _ = v55
                        }
                        v56 := v4[(v53 + ("foo bar" + (("\n\t" + ((v53 + v12) + v12)) + v12)))]
                        _ = v53
                        _ = v56
                    } else {
                        var v57 bool = (-1 < 1)
                        _ = v57
                    }
                    v58 := GenS4{F0: 0}
                    fmt.Println(v1, v1)
                    _ = v58
                }
            default:
                v59 := v2["a"]
                _ = v59
            }
            _ = len(v3)
            var v60 int64 = 9223372036854775807
            _ = fi42
            _ = v60
        }
    case 32768:
        var v61 uint32 = 966
        _ = v61
    default:
        var v62 uint8 = 1
        _ = v62
    }
    v63 := GenS5{F0: 9223372036854775807, F1: 0, F2: 2}
    var v64 uint16 = 1
    for fi65 := 0; fi65 < 18; fi65++ {
        if true {
            v1 = (((0.0 * 0.0) + (float64(0.1) - -2.25)) + (v1 - float64(((3.14159 * (-2.25 + 1e10)) - 1e10))))
            var v66 uint8 = 2
            if len(v3) > 0 {
                v3[0] = 300
            }
            _ = v66
        } else {
            var v67 uint16 = v64
            var v68 uint16 = 255
            _ = v67
            _ = v68
        }
        for fi69 := 0; fi69 < 13; fi69++ {
            var v70 float32 = 0.1
            v71 := GenS6{F0: ""}
            v64 -= 187
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = fi65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v63
    _ = v64
}
