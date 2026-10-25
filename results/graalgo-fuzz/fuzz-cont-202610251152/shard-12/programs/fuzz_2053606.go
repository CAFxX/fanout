package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0(p0 int8, p1 int64) string {
    if true {
        p0 = (p0 | (p0 * 127))
        var v5 uint32 = 486
        _ = v5
    } else {
        v6 := []int{-539, 2147483647, 560, 2147483647}
        _ = v6
    }
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int64 = p1
            v10 := []int{fi7, fi8}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return "foo bar"
}
func genFn1(p0 uint16, p1 uint32, p2 int16) bool {
    fmt.Println(p1)
    v11 := GenS1{F0: p1, F1: -2147483648, F2: (p2 << ((2) & 63))}
    if false {
        v12 := []int{-2147483648, 127, 32768, 42}
        _ = v12
    }
    v13 := GenS1{F0: p1, F1: 1, F2: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v13
    return true
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 44
    var v3 uint16 = 23
    var v4 int = -2147483648
    var v14 float64 = 1.5
    for fi15 := 0; fi15 < 16; fi15++ {
        for fi16 := 0; fi16 < 20; fi16++ {
            v17 := GenS1{F0: uint32(v3), F1: 1, F2: 0}
            var v18 int32 = 165
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 uint8 = 0
            switch "foo bar" {
            case "\n\t":
                var v21 float32 = (1.5 / 1e10)
                _ = v21
            case "hello":
                v3 = (v3 << ((4294967295) & 63))
            case "foo bar":
                v22 := GenS1{F0: 824, F1: -2147483648, F2: 727}
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        if false {
            v23 := []int{v4, (fi15 - 774), fi15, fi15}
            var v24 uint8 = 220
            if (!true) {
                delete(v1, "hello")
                var v25 uint32 = 2147483648
                if true {
                    switch ("\n\t" + ("" + "a")) {
                    case "foo bar":
                        if len(v23) > 0 {
                            v23[0] = v4
                        }
                    case "a":
                        v14 = v14
                    case "":
                        for fi26 := 0; fi26 < 4; fi26++ {
                            if len(v23) > 0 {
                                v23[0] = (v4 | v4)
                            }
                            if (2 < (v24 ^ (v24 << ((588) & 63)))) {
                                if true {
                                    v27 := GenS1{F0: (v25 / (((v25 | v25)) | 1)), F1: 0, F2: int16(v4)}
                                    _ = v27
                                }
                                v24 = 128
                            } else {
                                fmt.Println(fi15, fi15)
                            }
                            var v28 bool = (int32(v2) < -574)
                            v29 := GenS2{F0: 65535, F1: (v24 * v24), F2: v3}
                            if len(v23) > 0 {
                                v23[0] = (fi15 & fi15)
                            }
                            _ = fi26
                            _ = v28
                            _ = v29
                        }
                    }
                    v2 = v2
                    if false {
                        var v30 uint64 = v2
                        _ = v30
                    } else {
                        if (v14 < v14) {
                            if true {
                                for fi31 := 0; fi31 < 20; fi31++ {
                                    if ((v14 <= -2.25) || (2147483647 < 42)) {
                                        if false {
                                            if false {
                                                v4 -= (fi15 ^ fi15)
                                                v24 += 2
                                                var v32 bool = false
                                                if v32 {
                                                    v4 = (v4 >> ((4294967295) & 63))
                                                    v32 = (true || v32)
                                                    var v33 int32 = -2147483648
                                                    _ = v33
                                                }
                                                var v34 int64 = -9223372036854775808
                                                _ = v32
                                                _ = v34
                                            }
                                        }
                                    }
                                    switch ("foo bar" + ("" + "a")) {
                                    case "hello":
                                        if true {
                                            fmt.Println(v2, fi31, v4)
                                            v4 = (-(fi31))
                                            var v35 bool = (535 != 311)
                                            var v36 uint64 = (v2 % (((v2 % (((v2 - 4294967295)) | 1))) | 1))
                                            _ = v35
                                            _ = v36
                                        }
                                    default:
                                        var v37 int = fi15
                                        _ = v37
                                    }
                                    v38 := GenS3{F0: 788, F1: (v2 - (v2 | (v2 & v2))), F2: 4294967295}
                                    if false {
                                        if len(v23) > 0 {
                                            v23[0] = fi31
                                        }
                                        fmt.Println(v4, v4)
                                    } else {
                                        var v39 int16 = -468
                                        v23 = append(v23, (v4 & fi15))
                                        v2 += v2
                                        _ = v39
                                    }
                                    var v40 uint8 = 128
                                    _ = fi31
                                    _ = v38
                                    _ = v40
                                }
                                var v41 float64 = v14
                                _ = v41
                            }
                        } else {
                            v4 = v4
                            v3 = v3
                            v42 := GenS3{F0: 0, F1: v2, F2: 0}
                            _ = v42
                        }
                        for fi43 := 0; fi43 < 4; fi43++ {
                            v14 = float64(0.1)
                            v24 += 1
                            _ = fi43
                        }
                    }
                }
                _ = v25
            } else {
                switch fi15 {
                case 128:
                    var v44 int16 = 0
                    _ = v44
                case 885:
                    _ = len(v1)
                case 1:
                    v3 += v3
                default:
                    fmt.Println(v14)
                }
                for fi45 := 0; fi45 < 5; fi45++ {
                    v14 -= (v14 + -2.25)
                    v24 += v24
                    _ = fi45
                }
                v24 -= v24
            }
            _ = v23
            _ = v24
        }
        _ = fi15
    }
    v2 = v2
    v46 := v1[(("hello" + ("" + "\n\t")) + "")]
    if false {
        v47 := []int{v46, (v46 & (^(v4))), (^(v46))}
        var v48 int64 = 0
        _ = v47
        _ = v48
    } else {
        v49 := GenS1{F0: 2, F1: 0, F2: 42}
        _ = v49
    }
    v3 += (v3 << ((2) & 63))
    for fi50 := 0; fi50 < 15; fi50++ {
        if (661 == 42) {
            delete(v1, "\n\t")
        } else {
            v4 -= v46
        }
        if (-925 <= -781) {
            _ = len(v1)
            _ = len(v1)
        } else {
            delete(v1, "foo bar")
        }
        _ = fi50
    }
    v51 := GenS3{F0: 4294967295, F1: (^(v2)), F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v51
}
