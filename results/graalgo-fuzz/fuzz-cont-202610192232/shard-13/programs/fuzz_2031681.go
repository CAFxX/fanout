package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 int16; F1 int32; F2 int64 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint; F1 uint8; F2 int32 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int64; F1 int }
type GenS12 struct { F0 uint32; F1 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 42
    var v3 float32 = (((0.1 / (3.14159 + (-2.25 - -2.25))) / (float32(((0.1 + (3.14159 * (1.5 / float64(float32(3.14159))))) - float64(0.1))) - ((((1.5 - float32(((1e10 / float64(1.5)) + (0.1 * 3.14159)))) + float32(0.1)) - -2.25) * (1e10 / (1e10 + ((-2.25 / -2.25) - 1.5)))))) + 0.1)
    var v4 uint = 18446744073709551615
    v5 := make(map[int]string)
    var v6 int16 = (v2 + -32768)
    v7 := GenS1{F0: -104, F1: 127}
    for fi8 := 0; fi8 < 8; fi8++ {
        if false {
            v3 -= v3
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS2{F0: (false || ((v4 | v4) <= (^(v4)))), F1: 18446744073709551615}
                v4 = (v4 - v4)
                fmt.Println(v4, v3, v6)
                var v11 uint32 = 0
                v4 = (^(v4))
                if (-128 == 127) {
                    var v12 int = (fi8 ^ fi8)
                    v12 += fi8
                    _ = v12
                } else {
                    if false {
                        v2 -= int16(fi9)
                        if (4294967295 == 833) {
                            v2 = (v2 % ((v6) | 1))
                            v13 := GenS2{F0: (((0.1 - 0.1) - 0.1) > 3.14159), F1: 42}
                            fmt.Println(v11, v2)
                            _ = v13
                        }
                        v14 := []int{fi8}
                        _ = v14
                    } else {
                        fmt.Println(v2, v11, v11)
                        if (false && (!false)) {
                            if ((128 < 77) && false) {
                                if (!((fi8 + (fi8 * 127)) < (fi8 >> ((v4) & 63)))) {
                                    var v15 uint64 = 1
                                    var v16 uint8 = 0
                                    fmt.Println(v4, fi9, v15)
                                    if (v4 != (v4 % ((uint(v15)) | 1))) {
                                        switch "" {
                                        case "hello":
                                            switch ("" + "hello") {
                                            case "":
                                                var v17 uint64 = 42
                                                _ = v17
                                            case "\n\t":
                                                v18 := GenS2{F0: false, F1: 18446744073709551615}
                                                _ = v18
                                            }
                                        default:
                                            var v19 float32 = v3
                                            _ = v19
                                        }
                                        var v20 int16 = (v2 + (v6 >> ((v4) & 63)))
                                        v21 := GenS2{F0: true, F1: (^(v15))}
                                        v11 += (v11 - (^(v11)))
                                        switch (fi8 + (fi8 ^ (fi9 | fi9))) {
                                        case -9223372036854775808:
                                            fmt.Println(v16, v15)
                                        case 128:
                                            v22 := GenS1{F0: (v6 ^ (v20 + v6)), F1: 1}
                                            _ = v22
                                        default:
                                            if true {
                                                v6 -= v20
                                                var v23 float32 = (v3 - v3)
                                                fmt.Println(fi9, v20)
                                                _ = v23
                                            }
                                            var v24 int = fi8
                                            _ = v24
                                        }
                                        _ = v20
                                        _ = v21
                                    } else {
                                        _ = len(v1)
                                        fmt.Println(v6, v15, v6)
                                    }
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v25 uint8 = 127
                                    v26 := []int{fi8, fi9}
                                    fmt.Println(v11, v4, v11)
                                    delete(v5, (fi9 >> ((uint(v25)) & 63)))
                                    _ = v25
                                    _ = v26
                                }
                                switch fi8 {
                                case 0:
                                    if false {
                                        v27 := GenS3{F0: 255, F1: (-(fi9))}
                                        if false {
                                            fmt.Println(v6)
                                            var v28 int = -1
                                            var v29 int32 = 0
                                            _ = v28
                                            _ = v29
                                        } else {
                                            var v30 int64 = -707
                                            _ = v30
                                        }
                                        v31 := []int{fi9}
                                        fmt.Println(v4, v11, v3)
                                        _ = v27
                                        _ = v31
                                    }
                                case -71:
                                    v11 -= v11
                                }
                                switch "" {
                                case "hello":
                                    switch "hello" {
                                    case "\n\t":
                                        var v32 string = ""
                                        _ = v32
                                    case "a":
                                        v33 := GenS4{F0: 177, F1: 32767}
                                        _ = v33
                                    }
                                }
                                v34 := GenS5{F0: 0, F1: -419, F2: -1}
                                _ = v34
                            } else {
                                v6 = v2
                            }
                            var v35 int8 = 127
                            v2 += (v2 ^ 0)
                            _ = v35
                        } else {
                            v36 := GenS6{F0: (fi9 & (fi8 / (((fi9 << (((v4 / ((v4) | 1))) & 63))) | 1))), F1: -1}
                            var v37 int64 = -1
                            if true {
                                var v38 uint16 = 65535
                                v39 := GenS7{F0: 127}
                                _ = v38
                                _ = v39
                            } else {
                                var v40 uint = (v4 % (((v4 << (((v4 / ((v4) | 1))) & 63))) | 1))
                                var v41 uint64 = 573
                                v42 := []int{fi8, fi8, 127}
                                _ = v40
                                _ = v41
                                _ = v42
                            }
                            _ = v36
                            _ = v37
                        }
                    }
                    if false {
                        var v43 int32 = int32(v4)
                        v11 -= (v11 << ((42) & 63))
                        if true {
                            v44 := GenS5{F0: (-(v2)), F1: 1, F2: 9223372036854775807}
                            _ = v44
                        }
                        _ = v43
                    } else {
                        v11 = v11
                    }
                    var v45 int64 = -9223372036854775808
                    fmt.Println(v4)
                    _ = v45
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            for fi46 := 0; fi46 < 20; fi46++ {
                if (!true) {
                    var v47 uint8 = 42
                    var v48 bool = false
                    _ = v47
                    _ = v48
                } else {
                    v2 = v2
                    if (v4 == v4) {
                        v49 := GenS8{F0: v4, F1: 255, F2: int32(v4)}
                        _ = v49
                    }
                }
                if true {
                    v2 = int16(fi46)
                    v50 := GenS9{F0: 1, F1: 0, F2: 42}
                    if false {
                        var v51 int32 = 42
                        var v52 uint32 = 277
                        if true {
                            fmt.Println(v4, fi46, v3)
                            var v53 uint8 = 189
                            _ = v53
                        } else {
                            v54 := GenS1{F0: 0, F1: -111}
                            var v55 bool = (!false)
                            _ = v54
                            _ = v55
                        }
                        _ = v51
                        _ = v52
                    }
                    _ = v50
                } else {
                    if true {
                        if true {
                            if (18446744073709551615 != 42) {
                                v2 += (^(v6))
                            }
                            v56 := []int{(fi8 & fi46)}
                            v57 := GenS4{F0: 978, F1: 42}
                            _ = v56
                            _ = v57
                        }
                        switch fi46 {
                        case 0:
                            switch fi46 {
                            case 4294967295:
                                v4 = 65535
                            default:
                                v58 := GenS6{F0: fi46, F1: 127}
                                _ = v58
                            }
                            if (65535 >= (v4 % ((v4) | 1))) {
                                v3 += v3
                            }
                        }
                        if ((v4 - v4) > v4) {
                            v59 := GenS10{F0: (v4 % ((v4) | 1))}
                            if true {
                                v4 = (v4 | (^(v4)))
                                _ = len(v5)
                            } else {
                                fmt.Println(v3, v4)
                            }
                            _ = v59
                        } else {
                            var v60 int16 = v2
                            switch int(v4) {
                            case 127:
                                v3 = ((((v3 / v3) / v3) * 0.0) + ((((v3 + v3) * v3) - (v3 + -2.25)) / (v3 / (v3 * (-2.25 / v3)))))
                            case 814:
                                var v61 float64 = 1.5
                                _ = v61
                            default:
                                v3 += v3
                            }
                            var v62 bool = true
                            _ = v60
                            _ = v62
                        }
                    } else {
                        var v63 uint32 = 2147483647
                        var v64 uint = 628
                        _ = v63
                        _ = v64
                    }
                }
                v3 -= v3
                _ = fi46
            }
        }
        _ = fi8
    }
    for fi65 := 0; fi65 < 6; fi65++ {
        if true {
            v2 += (v2 + (v2 + (v6 | 1000)))
            for fi66 := 0; fi66 < 1; fi66++ {
                fmt.Println(v3, v6, v4)
                _ = fi66
            }
        }
        v4 += (^(v4))
        v67 := []int{(fi65 - 1), (fi65 ^ 448)}
        _ = fi65
        _ = v67
    }
    v4 = 2
    v2 -= -298
    v68 := GenS3{F0: 255, F1: 71}
    v69 := GenS11{F0: -9223372036854775808, F1: 9223372036854775807}
    v70 := GenS12{F0: 4294967295, F1: v4}
    if false {
        for fi71 := 0; fi71 < 12; fi71++ {
            v6 = (-(v2))
            _ = fi71
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v68
    _ = v69
    _ = v7
    _ = v70
}
