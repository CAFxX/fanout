package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 int32; F1 int8; F2 int64 }
type GenS5 struct { F0 int8; F1 bool }
type GenS6 struct { F0 uint16; F1 int16 }
func genFn0(p0 int16, p1 uint16, p2 uint8) int32 {
    v6 := []int{128, 32768, 128, 1}
    var v7 uint = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 1
}

func main() {
    var v1 int32 = -566
    var v2 bool = (65535 < 1)
    v3 := []uint32{4294967295, 0, 2147483647}
    v4 := make(map[uint32]bool)
    var v5 uint16 = 291
    v2 = ((!true) && v2)
    var v8 uint = 65535
    for fi9 := 0; fi9 < 18; fi9++ {
        var v10 int64 = 9223372036854775807
        switch "" {
        case "":
            var v11 uint8 = 2
            _ = v11
        case "foo bar":
            v5 += v5
        case "hello":
            fmt.Println(fi9, v5, fi9)
        }
        v8 += v8
        _ = fi9
        _ = v10
    }
    var v12 string = ((("\n\t" + ((("a" + "a") + "foo bar") + "a")) + "foo bar") + "\n\t")
    if (((((v1 / (((v1 ^ (v1 ^ int32(v5)))) | 1)) == v1) && v2) || true) || v2) {
        v13 := v4[2]
        fmt.Println(v1, v2, v2)
        _ = v13
    } else {
        var v14 uint64 = 843
        v12 = v12
        fmt.Println(v2)
        _ = v14
    }
    if (v1 <= 42) {
        v15 := GenS1{F0: 32767, F1: 128, F2: 5}
        _ = v15
    }
    v12 = v12
    fmt.Println(v12, v5, v2)
    fmt.Println(v5, v12)
    if ((v8 >> (((^(v8))) & 63)) > v8) {
        v5 += (v5 * v5)
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 string = v12
            for fi18 := 0; fi18 < 8; fi18++ {
                if (v17 >= v12) {
                    switch fi16 {
                    case -9223372036854775808:
                        var v19 int = -177
                        _ = v19
                    case 1:
                        fmt.Println(v1, v17)
                        fmt.Println(v5, fi16)
                        v1 -= genFn0(177, (^(v5)), 127)
                        if v2 {
                            if v2 {
                                var v20 string = v12
                                var v21 uint64 = 0
                                if ((v8 / (((^(v8))) | 1)) <= v8) {
                                    switch v20 {
                                    case "\n\t":
                                        fmt.Println(v21, v5)
                                    case "":
                                        v20 = (v17 + v17)
                                    default:
                                        v8 -= (v8 & uint(v21))
                                    }
                                    if (1 > 42) {
                                        v3 = append(v3, 4294967295)
                                        fmt.Println(v1)
                                        var v22 int16 = -32768
                                        v5 += (v5 + (^(v5)))
                                        if v2 {
                                            if v2 {
                                                var v23 int = (fi18 | (fi16 / (((fi16 | (fi18 / ((fi16) | 1)))) | 1)))
                                                v24 := GenS2{F0: 112}
                                                v3 = append(v3, 2147483648)
                                                var v25 float32 = ((float32(0.1) / (1e10 * 3.14159)) / 3.14159)
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            v8 = (v8 / ((2) | 1))
                                        } else {
                                            v26 := GenS3{F0: int8(v1), F1: 2}
                                            _ = len(v4)
                                            _ = v26
                                        }
                                        if (uint(v21) == v8) {
                                            var v27 uint = 1
                                            if (false && ((2 < 691) || false)) {
                                                v27 = 310
                                                var v28 string = "a"
                                                _ = v28
                                            }
                                            var v29 float32 = float32(0.1)
                                            _ = v27
                                            _ = v29
                                        } else {
                                            if v2 {
                                                v30 := GenS3{F0: 100, F1: v21}
                                                _ = v30
                                            }
                                            var v31 string = ((v12 + (v12 + v12)) + v12)
                                            var v32 uint = v8
                                            _ = v31
                                            _ = v32
                                        }
                                        var v33 int32 = (v1 - v1)
                                        fmt.Println(v17, v17, v2)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = v22
                                        _ = v33
                                    } else {
                                        v2 = (1 >= 32767)
                                    }
                                    var v34 uint64 = v21
                                    _ = v34
                                }
                                _ = v20
                                _ = v21
                            }
                            v35 := GenS2{F0: 33}
                            var v36 int8 = 6
                            _ = v35
                            _ = v36
                        }
                        fmt.Println(v17, v12, v2)
                    case -372:
                        v17 = ("" + (((v17 + "a") + v12) + ("" + v17)))
                    }
                } else {
                    v37 := GenS1{F0: -514, F1: fi16, F2: 0}
                    if true {
                        v38 := GenS4{F0: (v1 >> ((v8) & 63)), F1: -128, F2: 1}
                        if false {
                            v12 = ("a" + v12)
                        }
                        var v39 int64 = -9223372036854775808
                        _ = v38
                        _ = v39
                    }
                    var v40 bool = ((-1 > 0) || (221 >= 38))
                    _ = v37
                    _ = v40
                }
                v5 = (^(v5))
                _ = fi18
            }
            var v41 uint16 = v5
            _ = fi16
            _ = v17
            _ = v41
        }
        if true {
            v42 := GenS4{F0: v1, F1: 0, F2: -830}
            v2 = (!(!false))
            for fi43 := 0; fi43 < 8; fi43++ {
                v44 := v4[2147483648]
                switch (fi43 % ((-2147483648) | 1)) {
                case 2147483647:
                    v45 := GenS5{F0: 0, F1: v2}
                    _ = v45
                case 127:
                    var v46 uint16 = (v5 | 255)
                    _ = v46
                case 1:
                    var v47 uint32 = 2147483647
                    _ = v47
                default:
                    var v48 bool = v2
                    _ = v48
                }
                _ = fi43
                _ = v44
            }
            _ = v42
        } else {
            fmt.Println(v1, v8)
            if (false && v2) {
                v49 := GenS4{F0: (v1 * v1), F1: 127, F2: -9223372036854775808}
                _ = v49
            } else {
                v50 := GenS6{F0: v5, F1: 1}
                for fi51 := 0; fi51 < 16; fi51++ {
                    if false {
                        var v52 int16 = 1
                        _ = v52
                    }
                    _ = fi51
                }
                _ = len(v4)
                _ = v50
            }
        }
    }
    for fi53 := 0; fi53 < 2; fi53++ {
        v2 = v2
        _ = fi53
    }
    if (257 == 521) {
        v54 := GenS2{F0: 0}
        if v2 {
            v55 := GenS4{F0: 628, F1: -123, F2: 1}
            _ = v55
        } else {
            for fi56 := 0; fi56 < 5; fi56++ {
                var v57 int64 = -76
                if (!v2) {
                    if v2 {
                        var v58 uint64 = 862
                        v59 := GenS6{F0: 946, F1: 741}
                        _ = v58
                        _ = v59
                    } else {
                        for fi60 := 0; fi60 < 3; fi60++ {
                            if ((v1 << ((255) & 63)) != -599) {
                                switch v12 {
                                case "foo bar":
                                    fmt.Println(v57, v1)
                                }
                            }
                            v61 := GenS5{F0: 0, F1: true}
                            if len(v3) > 0 {
                                v3[0] = 1
                            }
                            _ = fi60
                            _ = v61
                        }
                        for fi62 := 0; fi62 < 18; fi62++ {
                            var v63 int = 42
                            v12 = ((v12 + "") + ((v12 + "a") + (((v12 + ("" + v12)) + ("a" + ("foo bar" + "a"))) + (v12 + v12))))
                            _ = fi62
                            _ = v63
                        }
                    }
                }
                v12 = v12
                _ = fi56
                _ = v57
            }
            for fi64 := 0; fi64 < 3; fi64++ {
                for fi65 := 0; fi65 < 17; fi65++ {
                    switch ((v12 + "a") + v12) {
                    case "":
                        if false {
                            var v66 float32 = ((float32(1e10) + float32(float64(-2.25))) + (3.14159 / 3.14159))
                            v2 = v2
                            var v67 int32 = (v1 ^ genFn0(-1, (v5 | 65535), 1))
                            _ = v66
                            _ = v67
                        } else {
                            v12 = v12
                        }
                    }
                    v8 = 255
                    v2 = false
                    _ = fi65
                }
                _ = fi64
            }
        }
        _ = v54
    } else {
        delete(v4, 377)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
