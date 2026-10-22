package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int }
type GenS2 struct { F0 int; F1 int8 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 int16; F2 bool }
type GenS6 struct { F0 uint32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{128, 4294967295, -635}
    v3 := []string{"a"}
    v4 := []int{377, 872, 42}
    v5 := make(map[uint32]int)
    var v6 int = 700
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            v8 := GenS1{F0: 552, F1: 2, F2: v6}
            for fi9 := 0; fi9 < 19; fi9++ {
                var v10 uint32 = 0
                if true {
                    var v11 int8 = 100
                    switch (fi7 * (v6 * fi9)) {
                    case 850:
                        switch v6 {
                        case 0:
                            var v12 float64 = 1e10
                            _ = v12
                        case 2147483647:
                            fmt.Println(v10, fi9)
                            var v13 int = (v6 - 42)
                            _ = v13
                        default:
                            if (("\n\t" + "\n\t") == "\n\t") {
                                v14 := GenS1{F0: 0, F1: 64, F2: (v6 / (((v6 + 42)) | 1))}
                                switch (v6 % (((fi9 % (((fi7 | v6)) | 1))) | 1)) {
                                case 732:
                                    v10 += (v10 & v10)
                                }
                                fmt.Println(fi7, v11, v11)
                                v11 = v11
                                var v15 uint64 = 720
                                _ = v14
                                _ = v15
                            }
                        }
                    case -161:
                        fmt.Println(v6, fi9, fi9)
                    default:
                        v16 := GenS1{F0: 488, F1: 128, F2: -2147483648}
                        _ = v16
                    }
                    _ = v11
                } else {
                    v17 := GenS1{F0: 0, F1: 0, F2: (fi7 / ((fi9) | 1))}
                    var v18 bool = false
                    v19 := GenS1{F0: 1, F1: 1, F2: (-(fi7))}
                    if v18 {
                        v18 = false
                        v20 := GenS1{F0: -1, F1: 3, F2: fi9}
                        _ = v20
                    } else {
                        var v21 int = (fi9 >> ((0) & 63))
                        v6 = fi9
                        _ = v21
                    }
                    v6 = (fi7 ^ (-(fi9)))
                    var v22 uint8 = 128
                    if true {
                        v22 -= (^(v22))
                        _ = len(v2)
                    }
                    var v23 uint8 = (v22 - 255)
                    v24 := GenS2{F0: 570, F1: -1}
                    if false {
                        var v25 uint = 255
                        _ = v25
                    }
                    fmt.Println(v10)
                    if true {
                        v26 := GenS2{F0: 32767, F1: 100}
                        if false {
                            v27 := GenS3{F0: -2147483648, F1: 0}
                            var v28 int = v6
                            _ = v27
                            _ = v28
                        }
                        var v29 float32 = ((1e10 * (((1e10 / -2.25) + float32(0.1)) - ((float32(0.0) / 1.5) / float32(-2.25)))) - 1.5)
                        var v30 string = "\n\t"
                        _ = v26
                        _ = v29
                        _ = v30
                    } else {
                        var v31 uint8 = (v23 << ((0) & 63))
                        var v32 uint16 = 1
                        var v33 int8 = -128
                        if (true || false) {
                            var v34 bool = (!((v32 << ((381) & 63)) < v32))
                            if v18 {
                                fmt.Println(v32, v6)
                            } else {
                                if (!((((v33 * -128) != int8(v32)) || true) || (v23 != v31))) {
                                    fmt.Println(v33, v23)
                                    v35 := GenS2{F0: (v6 << ((1) & 63)), F1: v33}
                                    _ = len(v1)
                                    var v36 int = (fi7 + v6)
                                    var v37 int8 = (v33 % (((v33 ^ (v33 >> ((434) & 63)))) | 1))
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                } else {
                                    v38 := GenS4{F0: -9223372036854775808, F1: (v31 & (^(v22)))}
                                    if v34 {
                                        var v39 float32 = 3.14159
                                        _ = v39
                                    } else {
                                        var v40 uint8 = (v31 / (((v31 - (v23 | v22))) | 1))
                                        var v41 int8 = 21
                                        _ = v40
                                        _ = v41
                                    }
                                    var v42 uint16 = v32
                                    _ = v38
                                    _ = v42
                                }
                                fmt.Println(fi7)
                                var v43 string = ""
                                _ = v43
                            }
                            v44 := GenS5{F0: (v32 + v32), F1: -32768, F2: v34}
                            var v45 float32 = 0.1
                            fmt.Println(v34)
                            _ = v34
                            _ = v44
                            _ = v45
                        } else {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v6 += (fi9 << ((2) & 63))
                            var v46 uint64 = 18446744073709551615
                            fmt.Println(fi9)
                            switch "" {
                            case "foo bar":
                                var v47 uint64 = (v46 ^ uint64(v10))
                                _ = v47
                            }
                            _ = v46
                        }
                        var v48 uint32 = v10
                        var v49 float64 = 1.5
                        _ = v31
                        _ = v32
                        _ = v33
                        _ = v48
                        _ = v49
                    }
                    _ = v17
                    _ = v18
                    _ = v19
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    }
    v50 := GenS4{F0: -9223372036854775808, F1: 2}
    var v51 int = v6
    v51 = v51
    v51 = 162
    var v52 int16 = 42
    for fi53 := 0; fi53 < 19; fi53++ {
        var v54 float64 = 1.5
        v1[("" + ("foo bar" + "foo bar"))] = true
        var v55 bool = ((((!(42 < 499)) || ((2147483647 > 42) || false)) && false) || ((float32(((-2.25 - (((v54 / -2.25) + float64(0.0)) * (1e10 * v54))) + (((v54 + v54) / 1e10) + 1.5))) * -2.25) == (3.14159 / 1.0)))
        _ = fi53
        _ = v54
        _ = v55
    }
    if (32767 > 2) {
        switch "hello" {
        case "a":
            if false {
                if false {
                    for fi56 := 0; fi56 < 14; fi56++ {
                        fmt.Println(v6, v6, v51)
                        var v57 uint8 = 2
                        delete(v1, (("hello" + "hello") + ""))
                        _ = fi56
                        _ = v57
                    }
                    if (32768 > 32767) {
                        var v58 float32 = 1e10
                        v59 := GenS3{F0: 0, F1: -1}
                        _ = v58
                        _ = v59
                    }
                    if false {
                        v60 := GenS2{F0: -1, F1: 100}
                        v52 += (v52 + 1000)
                        _ = v60
                    }
                } else {
                    v6 = 0
                    v61 := GenS6{F0: 2}
                    _ = v61
                }
            }
        default:
            v52 -= (^(v52))
        }
    }
    var v62 bool = false
    v51 = v51
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v6
    _ = v62
}
