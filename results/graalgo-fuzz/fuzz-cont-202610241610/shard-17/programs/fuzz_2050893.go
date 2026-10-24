package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 string; F1 uint8; F2 uint64 }
func genFn0(p0 int32) uint16 {
    v6 := []int{}
    p0 = p0
    if true {
        var v7 uint16 = 0
        if true {
            for fi8 := 0; fi8 < 5; fi8++ {
                _ = len(v6)
                for fi9 := 0; fi9 < 16; fi9++ {
                    if false {
                        var v10 int32 = -1
                        if true {
                            switch 4294967295 {
                            case -831:
                                if len(v6) > 0 {
                                    v6[0] = (fi9 % ((fi9) | 1))
                                }
                            default:
                                fmt.Println(fi8)
                            }
                            v11 := GenS1{F0: -1, F1: 2}
                            v7 = 852
                            _ = v11
                        } else {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        var v12 int8 = 24
                        _ = v10
                        _ = v12
                    } else {
                        v7 = (v7 >> ((42) & 63))
                        v13 := GenS2{F0: -760, F1: 32767, F2: 260}
                        v14 := GenS1{F0: -32768, F1: 2147483648}
                        _ = v13
                        _ = v14
                    }
                    var v15 uint32 = 1
                    v7 += v7
                    _ = fi9
                    _ = v15
                }
                _ = fi8
            }
        } else {
            v16 := GenS3{F0: 65535, F1: 32767}
            fmt.Println(v7, p0)
            var v17 int16 = -1
            _ = v16
            _ = v17
        }
        v18 := GenS3{F0: 2147483648, F1: 819}
        _ = v18
        _ = v7
    } else {
        for fi19 := 0; fi19 < 12; fi19++ {
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 float64 = -2.25
                _ = fi20
                _ = v21
            }
            p0 += (p0 << ((42) & 63))
            var v22 int16 = 726
            _ = fi19
            _ = v22
        }
        var v23 float32 = 0.0
        _ = v23
    }
    var v24 string = "\n\t"
    var v25 uint64 = 4294967295
    _ = p0
    _ = v24
    _ = v25
    _ = v6
    return 0
}
func genFn1(p0 int16, p1 int32, p2 int32) uint32 {
    var v26 float32 = 3.14159
    var v27 int32 = 2147483647
    switch ("foo bar" + "a") {
    case "hello":
        if (1 <= 42) {
            p1 -= p1
            var v28 bool = true
            _ = v28
        }
    }
    var v29 float64 = -2.25
    for fi30 := 0; fi30 < 19; fi30++ {
        var v31 float64 = float64(v26)
        fmt.Println(p1, v31)
        _ = fi30
        _ = v31
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v27
    _ = v29
    return 4294967295
}

func main() {
    var v1 uint32 = 1
    var v2 uint8 = 2
    v3 := make(map[string]uint32)
    v4 := []uint32{(v1 % (((v1 * v1)) | 1)), 309, (v1 >> ((65535) & 63)), 155}
    var v5 int8 = -122
    for fi32 := 0; fi32 < 12; fi32++ {
        for fi33 := 0; fi33 < 6; fi33++ {
            v1 -= (v1 ^ v1)
            _ = fi33
        }
        _ = fi32
    }
    v34 := GenS4{F0: 0}
    var v35 float64 = -2.25
    var v36 uint16 = 255
    v35 -= (v35 - 1e10)
    v37 := GenS5{F0: 23, F1: (!true)}
    v38 := GenS4{F0: (v36 % (((v36 / ((genFn0(2147483647)) | 1))) | 1))}
    var v39 float32 = float32(1.5)
    switch ("\n\t" + ("" + "\n\t")) {
    case "a":
        v3["\n\t"] = v1
    }
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int32 = -357
        delete(v3, "hello")
        _ = fi40
        _ = v41
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v42 := GenS6{F0: "", F1: v2, F2: 4294967295}
    v35 += v35
    if false {
        switch -731 {
        case 2147483647:
            var v43 float32 = v39
            _ = v43
        case 128:
            fmt.Println(v35, v2, v5)
        case 127:
            if true {
                for fi44 := 0; fi44 < 12; fi44++ {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        var v45 string = "foo bar"
                        _ = v45
                    case "":
                        if false {
                            if true {
                                fmt.Println(v36, v1, v36)
                                switch -2147483648 {
                                case 709:
                                    var v46 uint64 = 1
                                    _ = v46
                                }
                            } else {
                                fmt.Println(v35, v2)
                                var v47 int = fi44
                                _ = v47
                            }
                            switch ("foo bar" + ((("a" + "foo bar") + (("a" + "") + "")) + "a")) {
                            case "\n\t":
                                var v48 uint32 = 2147483647
                                _ = v48
                            }
                        } else {
                            var v49 uint8 = 1
                            _ = v49
                        }
                    case "hello":
                        if (65535 < 4294967295) {
                            v5 += -67
                        }
                    default:
                        for fi50 := 0; fi50 < 7; fi50++ {
                            v35 = v35
                            _ = fi50
                        }
                    }
                    for fi51 := 0; fi51 < 12; fi51++ {
                        delete(v3, "")
                        _ = len(v3)
                        _ = fi51
                    }
                    _ = fi44
                }
                for fi52 := 0; fi52 < 2; fi52++ {
                    if (true || false) {
                        var v53 int64 = 529
                        _ = v53
                    }
                    var v54 int16 = -32768
                    fmt.Println(v39, v39)
                    _ = fi52
                    _ = v54
                }
                fmt.Println(v39, v2, v39)
            } else {
                if (1 != 255) {
                    var v55 uint64 = 2
                    v39 = v39
                    _ = v55
                } else {
                    v35 -= float64(float32((v35 * 0.1)))
                    var v56 uint8 = v2
                    v4 = append(v4, v1)
                    _ = v56
                }
            }
        }
        v4 = append(v4, (v1 / ((genFn1(0, -1, -2147483648)) | 1)))
        v4 = append(v4, v1)
    } else {
        for fi57 := 0; fi57 < 9; fi57++ {
            v35 -= v35
            var v58 int32 = 728
            if false {
                var v59 uint64 = 42
                _ = v59
            }
            _ = fi57
            _ = v58
        }
        v4 = append(v4, (^(v1)))
    }
    v1 -= (v1 | genFn1(409, -1, 0))
    if (768 != 841) {
        var v60 uint8 = v2
        var v61 int64 = -9223372036854775808
        _ = v60
        _ = v61
    } else {
        var v62 int = 604
        v39 = float32(v35)
        v63 := v3["hello"]
        _ = v62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
