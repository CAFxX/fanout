package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 bool; F1 int64; F2 int }
type GenS4 struct { F0 uint8; F1 int16 }
func genFn0(p0 int8, p1 int64, p2 int64) int8 {
    var v4 int32 = 0
    switch ((((("foo bar" + "hello") + ("foo bar" + ("foo bar" + "hello"))) + "\n\t") + "a") + "\n\t") {
    case "foo bar":
        switch "\n\t" {
        case "hello":
            if (1 != 69) {
                v5 := []int{-2147483648}
                _ = v5
            } else {
                p0 -= (p0 / ((p0) | 1))
            }
        }
    default:
        var v6 uint = 2
        _ = v6
    }
    fmt.Println(p2, p0)
    if false {
        var v7 uint = 4294967295
        var v8 bool = (true && (-1 < -1))
        var v9 int = 242
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p0
}
func genFn1(p0 bool) string {
    v10 := []int{}
    if p0 {
        v11 := GenS1{F0: 255, F1: 42}
        _ = v11
    } else {
        if (p0 && p0) {
            v12 := GenS1{F0: 2, F1: 881}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v12
        } else {
            v13 := GenS1{F0: 1, F1: 2}
            var v14 int8 = 1
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 4; fi15++ {
            v16 := GenS1{F0: 2, F1: 0}
            p0 = (((((0.1 + 1e10) + -2.25) - 3.14159) - 0.0) == 3.14159)
            fmt.Println(fi15)
            _ = fi15
            _ = v16
        }
    }
    _ = p0
    _ = v10
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{32767, 1}
    v3 := make(map[string]uint32)
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 15; fi18++ {
            var v19 float32 = ((((1e10 * 1.5) - (0.1 - (float32(-2.25) * float32(-2.25)))) * (3.14159 - 1e10)) + 1e10)
            v19 -= 0.1
            switch fi17 {
            case 42:
                v20 := v1["hello"]
                _ = v20
            }
            _ = fi18
            _ = v19
        }
        var v21 int64 = 408
        for fi22 := 0; fi22 < 19; fi22++ {
            v21 += (v21 << ((1) & 63))
            _ = fi22
        }
        _ = fi17
        _ = v21
    }
    var v23 int64 = 541
    var v24 float32 = 0.1
    var v25 int8 = 4
    v24 = -2.25
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS1{F0: 0, F1: 2}
        if ((v24 + v24) != (v24 * 0.0)) {
            for fi28 := 0; fi28 < 8; fi28++ {
                if (7 > 898) {
                    if false {
                        v25 += (v25 - genFn0(v25, v23, -1))
                    } else {
                        var v29 uint64 = 396
                        _ = v29
                    }
                    var v30 float32 = ((v24 - (0.1 / v24)) - 1e10)
                    if len(v2) > 0 {
                        v2[0] = fi26
                    }
                    var v31 bool = true
                    _ = v30
                    _ = v31
                } else {
                    fmt.Println(v23)
                    switch "" {
                    case "":
                        v23 += v23
                    default:
                        fmt.Println(v25)
                    }
                }
                fmt.Println(v24, v23, v23)
                fmt.Println(v24, v25, fi28)
                var v32 int32 = -781
                v24 = v24
                _ = fi28
                _ = v32
            }
            if (false && true) {
                var v33 uint64 = 18446744073709551615
                switch "\n\t" {
                case "foo bar":
                    if false {
                        v34 := GenS1{F0: 65535, F1: (v33 & 42)}
                        var v35 float64 = 1.5
                        _ = v34
                        _ = v35
                    } else {
                        var v36 int32 = 0
                        delete(v1, ("" + "a"))
                        switch "hello" {
                        case "hello":
                            _ = len(v2)
                        }
                        _ = v36
                    }
                case "a":
                    if true {
                        var v37 uint64 = 494
                        var v38 uint32 = 2147483648
                        _ = v37
                        _ = v38
                    } else {
                        v33 -= (v33 - v33)
                        v3["hello"] = 42
                    }
                case "\n\t":
                    if true {
                        if len(v2) > 0 {
                            v2[0] = (fi26 & fi26)
                        }
                        fmt.Println(v23, v24, v33)
                    }
                default:
                    v39 := GenS1{F0: 0, F1: v33}
                    _ = v39
                }
                _ = v33
            } else {
                v24 = v24
                v40 := GenS1{F0: 944, F1: 70}
                v41 := GenS2{F0: 1, F1: v23}
                _ = v40
                _ = v41
            }
        } else {
            var v42 uint16 = 122
            for fi43 := 0; fi43 < 13; fi43++ {
                v44 := GenS3{F0: (4294967295 == 97), F1: 42, F2: fi43}
                v45 := GenS1{F0: v42, F1: 4294967295}
                fmt.Println(fi26, fi43)
                v46 := GenS4{F0: 2, F1: -1}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
            v1[("a" + "foo bar")] = ((100 > (v25 * (v25 * (^(v25))))) || true)
            _ = v42
        }
        var v47 uint16 = 373
        _ = fi26
        _ = v27
        _ = v47
    }
    for fi48 := 0; fi48 < 3; fi48++ {
        var v49 uint8 = 127
        if (false && true) {
            v50 := GenS2{F0: 1000, F1: (v23 * v23)}
            for fi51 := 0; fi51 < 2; fi51++ {
                var v52 float64 = 1.5
                if len(v2) > 0 {
                    v2[0] = (fi48 + (fi48 ^ (-(fi48))))
                }
                _ = fi51
                _ = v52
            }
            _ = v50
        } else {
            for fi53 := 0; fi53 < 6; fi53++ {
                if false {
                    var v54 int = -424
                    v25 = v25
                    var v55 uint8 = v49
                    _ = v54
                    _ = v55
                } else {
                    fmt.Println(v24, fi48)
                }
                var v56 int16 = 32767
                _ = fi53
                _ = v56
            }
        }
        _ = fi48
        _ = v49
    }
    v24 = 1.5
    v23 += v23
    switch "a" {
    case "":
        for fi57 := 0; fi57 < 5; fi57++ {
            if (-1 != -32768) {
                v25 = genFn0(v25, (v23 * (^(v23))), int64(fi57))
                _ = len(v3)
            } else {
                var v58 float32 = 3.14159
                v58 = v24
                _ = v58
            }
            _ = fi57
        }
    case "\n\t":
        var v59 bool = ((0.1 - float32(((1e10 / (float64(v24) * 1.5)) / 1.0))) == v24)
        _ = v59
    case "hello":
        var v60 int32 = -2147483648
        _ = v60
    default:
        var v61 float32 = v24
        _ = v61
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
