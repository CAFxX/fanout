package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 string; F1 int8; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint16; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint8; F1 int16 }
type GenS8 struct { F0 uint64; F1 int32 }
func genFn0() string {
    var v4 int16 = 32767
    v4 = 42
    var v5 uint = 18446744073709551615
    fmt.Println(v5, v4)
    var v6 string = "\n\t"
    _ = v4
    _ = v5
    _ = v6
    return v6
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 int64 = 643
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 float32 = (1.5 * (1e10 * 0.1))
        if true {
            v9 := GenS1{F0: "", F1: uint32(v2), F2: 4294967295}
            var v10 uint8 = 21
            if false {
                switch (fi7 - 127) {
                case -2147483648:
                    v11 := []int{fi7, fi7, (fi7 / ((-155) | 1))}
                    _ = v11
                case 2147483647:
                    var v12 uint8 = (v10 / ((v10) | 1))
                    _ = v12
                default:
                    for fi13 := 0; fi13 < 9; fi13++ {
                        v1[4294967295] = fi13
                        _ = fi13
                    }
                }
                if (((float64((v8 - float32(-2.25))) * (1e10 + -2.25)) * -2.25) < 0.1) {
                    v3 = v3
                }
                v3 -= v3
            }
            _ = v10
            _ = v9
        } else {
            var v14 string = "a"
            v3 += v3
            for fi15 := 0; fi15 < 16; fi15++ {
                var v16 uint8 = 2
                v2 = v2
                switch v14 {
                case "a":
                    v3 += (v3 | v3)
                case "\n\t":
                    if false {
                        fmt.Println(v2, v14)
                        var v17 uint32 = 2147483648
                        _ = v17
                    } else {
                        switch v14 {
                        case "hello":
                            if (!false) {
                                var v18 uint8 = (v16 ^ v16)
                                v18 += 25
                                _ = v18
                            } else {
                                var v19 float64 = 1.5
                                v20 := GenS2{F0: -1, F1: 18446744073709551615}
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS3{F0: 2147483647}
                            v8 -= float32(float64((v8 - (((0.0 / (-2.25 * v8)) * 0.0) * 1.5))))
                            _ = v21
                        }
                        if (18446744073709551615 > 1) {
                            if true {
                                if (true && false) {
                                    var v22 bool = false
                                    var v23 float32 = 3.14159
                                    _ = v22
                                    _ = v23
                                }
                                v24 := []int{(fi15 | (fi15 + fi7)), fi15}
                                _ = len(v24)
                                var v25 int16 = 42
                                if len(v24) >= 2 {
                                    v24 = v24[1:len(v24)]
                                }
                                v26 := GenS2{F0: -128, F1: 0}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                var v27 int32 = 1
                                var v28 uint = 0
                                _ = v27
                                _ = v28
                            }
                            switch v14 {
                            case "foo bar":
                                var v29 int = fi7
                                v30 := GenS3{F0: 42}
                                v31 := GenS2{F0: -128, F1: 856}
                                v8 += v8
                                var v32 string = (v14 + v14)
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                            case "a":
                                var v33 int64 = 9223372036854775807
                                _ = v33
                            default:
                                if true {
                                    v14 = v14
                                    if (!true) {
                                        var v34 int16 = -32768
                                        var v35 uint64 = 780
                                        var v36 int16 = (v34 >> ((255) & 63))
                                        _ = v34
                                        _ = v35
                                        _ = v36
                                    }
                                    var v37 float32 = ((v8 * 1.5) + 0.0)
                                    v37 -= v8
                                    var v38 int8 = 23
                                    _ = v37
                                    _ = v38
                                } else {
                                    v39 := GenS1{F0: (v14 + v14), F1: 80, F2: 1}
                                    _ = v39
                                }
                            }
                            var v40 float32 = float32(float64(v8))
                            _ = v40
                        }
                    }
                case "hello":
                    v3 = v3
                default:
                    v41 := GenS4{F0: v14, F1: 100, F2: -1}
                    _ = v41
                }
                _ = fi15
                _ = v16
            }
            _ = v14
        }
        v1[2147483648] = (fi7 >> ((977) & 63))
        _ = fi7
        _ = v8
    }
    var v42 string = "a"
    v43 := []int{int(v3), -512}
    if (true && true) {
        for fi44 := 0; fi44 < 14; fi44++ {
            if (!true) {
                fmt.Println(fi44, v3, v42)
                if len(v43) >= 2 {
                    v43 = v43[1:len(v43)]
                }
            }
            if ((42 < 278) || false) {
                for fi45 := 0; fi45 < 9; fi45++ {
                    v42 = "a"
                    v2 = v2
                    _ = fi45
                }
                for fi46 := 0; fi46 < 13; fi46++ {
                    v42 = ((v42 + "a") + ("foo bar" + v42))
                    _ = fi46
                }
                if (!true) {
                    v47 := GenS1{F0: v42, F1: 0, F2: 4294967295}
                    _ = v47
                } else {
                    fmt.Println(v3, v42)
                    for fi48 := 0; fi48 < 17; fi48++ {
                        delete(v1, 2147483648)
                        v49 := GenS5{F0: 184, F1: (v2 & v2), F2: (fi48 / (((fi44 - fi48)) | 1))}
                        _ = fi48
                        _ = v49
                    }
                    if true {
                        if false {
                            v50 := GenS6{F0: (fi44 | 32767)}
                            for fi51 := 0; fi51 < 18; fi51++ {
                                var v52 int = int(v3)
                                v53 := GenS3{F0: 0}
                                v3 = (v3 | (v3 + v3))
                                _ = fi51
                                _ = v52
                                _ = v53
                            }
                            _ = v50
                        } else {
                            v54 := GenS7{F0: 0, F1: -831}
                            _ = v54
                        }
                    } else {
                        switch (fi44 % (((^(fi44))) | 1)) {
                        case 42:
                            var v55 int32 = 752
                            _ = v55
                        case 9223372036854775807:
                            _ = len(v43)
                        case 0:
                            v56 := GenS8{F0: 1, F1: 0}
                            _ = v56
                        }
                    }
                }
            }
            v57 := v1[2147483647]
            _ = fi44
            _ = v57
        }
    }
    var v58 uint64 = 750
    v59 := GenS5{F0: (v58 & 0), F1: v2, F2: -9223372036854775808}
    if true {
        switch v42 {
        case "":
            v58 = (v58 % (((v58 >> ((2) & 63))) | 1))
        default:
            v2 += 1
        }
        var v60 float64 = 1e10
        _ = v60
    }
    var v61 int8 = 0
    fmt.Println(v61)
    if ((42 <= 1) || (false || true)) {
        var v62 bool = (!false)
        v61 = (v61 >> ((255) & 63))
        _ = v62
    } else {
        _ = len(v43)
        if (false || (2147483647 >= 0)) {
            var v63 int32 = -1
            var v64 int16 = 184
            switch v42 {
            case "hello":
                var v65 uint8 = 0
                _ = v65
            }
            _ = v63
            _ = v64
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v58:", v58)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v42
    _ = v43
    _ = v58
    _ = v59
    _ = v61
}
