package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint; F2 int8 }
type GenS3 struct { F0 int; F1 int64; F2 string }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint; F1 int8; F2 int8 }
func genFn0() uint32 {
    var v6 uint64 = 0
    v7 := GenS1{F0: 526, F1: -661}
    _ = v6
    _ = v7
    return 623
}
func genFn1(p0 string, p1 uint8) uint16 {
    v8 := []int{128, 4294967295, 0}
    v9 := GenS1{F0: 18446744073709551615, F1: 0}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := []int{-629, 1, -986}
    var v2 int16 = 42
    var v3 uint64 = 2
    v4 := make(map[string]int)
    var v5 float64 = (float64(float32((-2.25 - (0.1 / (3.14159 * (0.1 + (0.0 * 0.1))))))) + -2.25)
    var v10 uint = 255
    _ = len(v4)
    v11 := GenS1{F0: v3, F1: 1}
    v12 := GenS1{F0: (v3 * 370), F1: (v2 << (((v10 % (((v10 >> (((v10 ^ (v10 * 255))) & 63))) | 1))) & 63))}
    v13 := GenS1{F0: v3, F1: (-(v2))}
    var v14 string = "\n\t"
    v3 += (v3 * v3)
    for fi15 := 0; fi15 < 17; fi15++ {
        delete(v4, (v14 + v14))
        v10 -= 523
        if (!((fi15 / (((^(fi15))) | 1)) >= (fi15 | (fi15 & 32768)))) {
            var v16 bool = false
            _ = v16
        }
        _ = fi15
    }
    if ((32767 <= (v2 ^ (v2 / ((-32768) | 1)))) || (42 < 1)) {
        for fi17 := 0; fi17 < 11; fi17++ {
            if (2147483647 != 437) {
                var v18 bool = false
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 uint8 = 255
                    v20 -= 235
                    if v18 {
                        if len(v1) > 0 {
                            v1[0] = fi19
                        }
                        _ = len(v4)
                    }
                    var v21 uint64 = v3
                    v22 := GenS1{F0: 4294967295, F1: v2}
                    var v23 string = ("a" + v14)
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 float64 = (v5 - v5)
                _ = v18
                _ = v24
            }
            _ = fi17
        }
        var v25 int16 = (v2 + v2)
        _ = v25
    } else {
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS2{F0: 200, F1: v10, F2: 42}
            _ = fi26
            _ = v27
        }
        var v28 float64 = ((((v5 * 0.1) / v5) / (float64(1e10) * v5)) * (1.5 + 0.0))
        _ = v28
    }
    var v29 uint16 = 2
    v29 -= genFn1(("a" + ("a" + ((v14 + v14) + "foo bar"))), 255)
    fmt.Println(v3, v3, v5)
    v30 := v4[v14]
    v4[""] = 32767
    var v31 bool = ((3.14159 - -2.25) <= (1.5 / (((0.0 * (float32((0.1 + 0.1)) + (((3.14159 - (1e10 * (0.1 * 1e10))) - (1.5 / (0.1 - (1.5 - ((1e10 / float32(float64(3.14159))) / (((0.1 - 1e10) - float32(v5)) * 1e10)))))) / (3.14159 * (3.14159 - float32(1.5)))))) / float32(v5)) / 1.0)))
    switch (v14 + v14) {
    case "foo bar":
        v32 := GenS2{F0: 127, F1: 140, F2: 57}
        _ = v32
    default:
        if true {
            if v31 {
                for fi33 := 0; fi33 < 13; fi33++ {
                    for fi34 := 0; fi34 < 9; fi34++ {
                        var v35 uint8 = 205
                        _ = fi34
                        _ = v35
                    }
                    v3 = v3
                    v3 -= 1
                    _ = fi33
                }
            } else {
                v36 := GenS1{F0: v3, F1: v2}
                fmt.Println(v10, v29)
                switch (v14 + ("\n\t" + (v14 + "foo bar"))) {
                case "\n\t":
                    var v37 uint = uint(v3)
                    _ = v37
                case "hello":
                    if false {
                        v2 = 869
                        var v38 int8 = 100
                        _ = v38
                    } else {
                        if v31 {
                            v4[("foo bar" + v14)] = v30
                            v1 = append(v1, (^(v30)))
                            fmt.Println(v14, v31, v14)
                        } else {
                            if v31 {
                                var v39 int8 = -128
                                v31 = false
                                _ = v39
                            } else {
                                if len(v1) > 0 {
                                    v1[0] = v30
                                }
                                v40 := GenS3{F0: v30, F1: -356, F2: v14}
                                _ = v40
                            }
                            var v41 float32 = float32(v5)
                            _ = v41
                        }
                    }
                default:
                    v42 := GenS3{F0: (v30 * int(v3)), F1: 9223372036854775807, F2: (v14 + v14)}
                    _ = v42
                }
                _ = v36
            }
            if (372 >= -2147483648) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v29 -= v29
                v43 := GenS3{F0: v30, F1: -625, F2: "hello"}
                _ = v43
            } else {
                _ = len(v4)
            }
            for fi44 := 0; fi44 < 4; fi44++ {
                _ = len(v4)
                var v45 uint16 = v29
                _ = fi44
                _ = v45
            }
        } else {
            switch "hello" {
            case "foo bar":
                fmt.Println(v30)
            case "hello":
                for fi46 := 0; fi46 < 1; fi46++ {
                    for fi47 := 0; fi47 < 4; fi47++ {
                        var v48 string = (v14 + "\n\t")
                        if v31 {
                            switch (fi46 + (fi47 | 644)) {
                            case 563:
                                v29 = genFn1(("a" + "a"), 2)
                            case -1:
                                var v49 int = (-(v30))
                                _ = v49
                            case 627:
                                v10 = v10
                            default:
                                if ((0.0 + (((0.1 / 0.1) + (((float64(0.0) - v5) - -2.25) - ((-2.25 + -2.25) - ((v5 + v5) * v5)))) * (float64((float32(0.0) + 0.0)) / ((0.1 / 3.14159) * v5)))) > v5) {
                                    switch "hello" {
                                    case "":
                                        fmt.Println(v10)
                                    default:
                                        switch (v30 | 42) {
                                        case -1:
                                            if ((v48 + (v14 + v14)) > (((((v14 + v48) + ((v48 + (v48 + v14)) + v48)) + v14) + v48) + (v48 + "foo bar"))) {
                                                v50 := GenS2{F0: 42, F1: 394, F2: 127}
                                                if (-1 == 100) {
                                                    v29 += v29
                                                }
                                                var v51 float32 = 0.0
                                                v52 := GenS2{F0: 1, F1: 65535, F2: 1}
                                                fmt.Println(v51, fi46)
                                                _ = v50
                                                _ = v51
                                                _ = v52
                                            } else {
                                                var v53 int64 = -1
                                                if ((-(v53)) >= -1) {
                                                    var v54 float32 = 1e10
                                                    v53 += v53
                                                    v55 := GenS3{F0: v30, F1: (v53 % ((int64(v2)) | 1)), F2: v14}
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if v31 {
                                                    v14 = v48
                                                    _ = len(v1)
                                                }
                                                _ = v53
                                            }
                                        }
                                    }
                                    v48 = ""
                                    v29 += genFn1(v14, 42)
                                    v56 := GenS4{F0: 127, F1: 1}
                                    var v57 int = (fi46 ^ fi46)
                                    v5 -= 0.0
                                    _ = v56
                                    _ = v57
                                }
                            }
                            v58 := GenS3{F0: (-(fi46)), F1: -1, F2: (v14 + v48)}
                            _ = len(v4)
                            v59 := GenS5{F0: v10, F1: 1, F2: 127}
                            _ = v58
                            _ = v59
                        } else {
                            v3 = (v3 & (v3 >> ((1) & 63)))
                        }
                        var v60 string = v14
                        _ = fi47
                        _ = v48
                        _ = v60
                    }
                    fmt.Println(fi46, v30, v3)
                    _ = fi46
                }
            case "\n\t":
                var v61 uint32 = 65535
                _ = v61
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(v14)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
