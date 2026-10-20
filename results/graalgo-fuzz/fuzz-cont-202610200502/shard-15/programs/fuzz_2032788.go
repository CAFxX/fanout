package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 string; F1 int64; F2 int16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 int8; F1 uint64 }
func genFn0(p0 int, p1 bool) string {
    var v4 float64 = 1.5
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch "foo bar" {
            case "a":
                p0 = (^(fi7))
            case "\n\t":
                var v8 int = -1
                _ = v8
            default:
                var v9 int8 = 120
                _ = v9
            }
            _ = fi7
        }
        if p1 {
            v5 = (v5 & v5)
            switch (fi6 & fi6) {
            case 1:
                v10 := []int{-1, 128}
                _ = v10
            default:
                var v11 int64 = -1
                _ = v11
            }
            v12 := []int{fi6, (fi6 | (p0 * (fi6 ^ 1))), fi6}
            _ = v12
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("" + "a")
}
func genFn1(p0 uint16, p1 uint8, p2 int16) uint64 {
    p2 = p2
    if (false && ((0.0 - 3.14159) >= ((-2.25 - -2.25) + ((0.1 * (float64(3.14159) / 0.1)) / (float64(-2.25) / 1e10))))) {
        v13 := GenS1{F0: true, F1: (-(p2)), F2: 125}
        _ = v13
    } else {
        var v14 uint = uint(p1)
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 int8 = 100
            switch fi15 {
            case 203:
                fmt.Println(p1)
            case 42:
                var v17 int8 = -128
                _ = v17
            case -2147483648:
                for fi18 := 0; fi18 < 5; fi18++ {
                    p0 = p0
                    v19 := GenS2{F0: 2}
                    if false {
                        p2 = (p2 & p2)
                    } else {
                        var v20 uint16 = p0
                        var v21 bool = false
                        _ = v20
                        _ = v21
                    }
                    p2 = int16(p0)
                    v16 -= (v16 - 19)
                    v22 := GenS3{F0: (p0 << ((42) & 63)), F1: (p1 - 233)}
                    _ = fi18
                    _ = v19
                    _ = v22
                }
            }
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := []int{(-(fi23))}
                p2 = p2
                v14 = v14
                _ = fi23
                _ = v24
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v25 int16 = -1
    v26 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return 2
}

func main() {
    var v1 float32 = 0.1
    var v2 bool = false
    v3 := []uint32{269, 818, 4294967295, 1}
    var v27 bool = (0.0 >= (1e10 - 0.1))
    v27 = v27
    var v28 bool = true
    var v29 int16 = -819
    for fi30 := 0; fi30 < 10; fi30++ {
        v31 := GenS3{F0: 0, F1: 249}
        _ = fi30
        _ = v31
    }
    fmt.Println(v2, v1)
    v32 := GenS4{F0: "foo bar", F1: -1, F2: (-(v29))}
    if (v29 == 42) {
        var v33 string = ("" + "foo bar")
        _ = v33
    } else {
        v29 = (v29 | v29)
        switch "foo bar" {
        case "\n\t":
            fmt.Println(v27, v27)
        case "":
            v28 = (!true)
        }
    }
    v1 -= 0.0
    var v34 int32 = 0
    v35 := GenS5{F0: 106}
    v1 += v1
    v36 := GenS6{F0: 185, F1: 2}
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    v27 = (42 <= 2147483648)
    if (704 >= 358) {
        for fi37 := 0; fi37 < 13; fi37++ {
            v2 = ((!v27) && (0 > 705))
            for fi38 := 0; fi38 < 4; fi38++ {
                _ = len(v3)
                v39 := GenS6{F0: 42, F1: 65535}
                v40 := GenS1{F0: v28, F1: v29, F2: 42}
                var v41 uint = 18446744073709551615
                v2 = (v41 <= v41)
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
            var v42 uint8 = 201
            _ = fi37
            _ = v42
        }
        switch 127 {
        case -713:
            if v27 {
                fmt.Println(v34, v1, v29)
            }
        case 32767:
            for fi43 := 0; fi43 < 8; fi43++ {
                if true {
                    v44 := GenS7{F0: -128, F1: 1}
                    var v45 int = int(v34)
                    v3 = append(v3, 410)
                    _ = v44
                    _ = v45
                } else {
                    fmt.Println(v1)
                    v46 := GenS4{F0: "\n\t", F1: 9223372036854775807, F2: v29}
                    _ = v46
                }
                _ = fi43
            }
        case 697:
            if (-175 <= 42) {
                if v27 {
                    if v2 {
                        if true {
                            var v47 float64 = 0.0
                            _ = v47
                        } else {
                            if ((!true) && v2) {
                                fmt.Println(v27, v1, v2)
                                v27 = false
                                if v27 {
                                    if len(v3) > 0 {
                                        v3[0] = 494
                                    }
                                    v3 = append(v3, 603)
                                } else {
                                    var v48 string = "hello"
                                    _ = len(v3)
                                    v2 = false
                                    _ = v48
                                }
                            }
                        }
                    } else {
                        switch -2147483648 {
                        case 1:
                            var v49 int16 = (-(v29))
                            _ = v49
                        case 446:
                            for fi50 := 0; fi50 < 1; fi50++ {
                                var v51 int32 = (v34 / ((-1) | 1))
                                fmt.Println(v29)
                                v52 := GenS4{F0: "foo bar", F1: 42, F2: v29}
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        case 2147483647:
                            v53 := GenS3{F0: 32767, F1: 42}
                            _ = v53
                        default:
                            for fi54 := 0; fi54 < 17; fi54++ {
                                _ = len(v3)
                                var v55 int32 = v34
                                if len(v3) > 0 {
                                    v3[0] = 438
                                }
                                _ = fi54
                                _ = v55
                            }
                        }
                    }
                    v56 := GenS2{F0: 255}
                    _ = v56
                }
            }
        default:
            fmt.Println(v1, v29, v2)
        }
    } else {
        v28 = v2
    }
    if false {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v36
}
