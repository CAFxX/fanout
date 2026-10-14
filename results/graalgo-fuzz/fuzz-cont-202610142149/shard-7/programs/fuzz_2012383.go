package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint64 }
type GenS3 struct { F0 uint32; F1 uint32; F2 bool }
type GenS4 struct { F0 int64; F1 int32; F2 int64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int16, p1 int) uint {
    switch (p1 >> ((65535) & 63)) {
    case 0:
        v7 := GenS1{F0: 2, F1: 1, F2: 1}
        _ = v7
    default:
        for fi8 := 0; fi8 < 7; fi8++ {
            var v9 float32 = 0.0
            _ = fi8
            _ = v9
        }
    }
    fmt.Println(p0, p1)
    var v10 float64 = -2.25
    switch (-(p1)) {
    case -9223372036854775808:
        for fi11 := 0; fi11 < 4; fi11++ {
            fmt.Println(p1, p1, fi11)
            _ = fi11
        }
    case 320:
        for fi12 := 0; fi12 < 11; fi12++ {
            if false {
                var v13 uint16 = 465
                _ = v13
            } else {
                v14 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                var v15 string = "hello"
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    default:
        v16 := []int{p1}
        _ = v16
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = v10
    return 856
}

func main() {
    var v1 int = 353
    var v2 uint = 18446744073709551615
    var v3 string = "foo bar"
    v4 := []int{v1, 127, (v1 | (v1 + v1)), (v1 - v1)}
    v5 := make(map[int]uint32)
    var v6 int64 = -1
    var v17 bool = true
    v18 := GenS2{F0: 1, F1: 42}
    if v17 {
        if v17 {
            if ((v2 >> ((genFn0(0, int(v2))) & 63)) >= v2) {
                v19 := GenS1{F0: 42, F1: 4294967295, F2: 1}
                _ = v19
            } else {
                var v20 int8 = 1
                for fi21 := 0; fi21 < 9; fi21++ {
                    v4 = append(v4, fi21)
                    if false {
                        v3 = "hello"
                        switch "" {
                        case "\n\t":
                            if len(v4) > 0 {
                                v4[0] = -203
                            }
                        default:
                            var v22 bool = ((1 == 238) && (v17 && v17))
                            _ = v22
                        }
                        var v23 float32 = 1.5
                        _ = v23
                    }
                    v17 = ("foo bar" >= v3)
                    _ = fi21
                }
                if v17 {
                    var v24 float32 = 0.1
                    var v25 int32 = 530
                    _ = v24
                    _ = v25
                } else {
                    var v26 string = ""
                    v27 := GenS3{F0: 4294967295, F1: 2147483648, F2: true}
                    var v28 string = "hello"
                    _ = v26
                    _ = v27
                    _ = v28
                }
                _ = v20
            }
            v6 -= (-(v6))
            v29 := GenS1{F0: 161, F1: 4294967295, F2: 255}
            _ = v29
        }
    } else {
        for fi30 := 0; fi30 < 5; fi30++ {
            if (2147483647 != (fi30 - v1)) {
                var v31 uint64 = 4294967295
                _ = v31
            }
            fmt.Println(v17)
            _ = fi30
        }
    }
    var v32 float64 = (float64(((0.0 / (float32((((1.5 * 3.14159) + 3.14159) + 1.5)) - 0.1)) / 1e10)) * 3.14159)
    for fi33 := 0; fi33 < 6; fi33++ {
        fmt.Println(v2, v1)
        v17 = true
        if (((!((1 < 32767) || v17)) && (!(v17 || (!true)))) && ((1e10 / (v32 * ((v32 * (v32 + (float64((0.0 + 0.0)) + 1e10))) + float64(1.5)))) < float64(3.14159))) {
            v17 = (v17 || true)
        }
        _ = fi33
    }
    fmt.Println(v6, v1, v1)
    for fi34 := 0; fi34 < 8; fi34++ {
        switch fi34 {
        case 373:
            var v35 int32 = -1
            _ = v35
        case -2147483648:
            if (!false) {
                fmt.Println(v32, v17)
            }
        default:
            if ((2 >= 2) || ("" == v3)) {
                switch "\n\t" {
                case "\n\t":
                    var v36 float32 = float32((v32 + 0.1))
                    _ = v36
                case "":
                    var v37 int64 = (v6 / (((v6 * -28)) | 1))
                    _ = v37
                default:
                    v2 -= v2
                }
                for fi38 := 0; fi38 < 7; fi38++ {
                    v17 = (617 >= int32(fi38))
                    _ = fi38
                }
            } else {
                v32 -= 0.1
                if false {
                    delete(v5, fi34)
                    switch (v3 + "\n\t") {
                    case "foo bar":
                        switch fi34 {
                        case 0:
                            var v39 string = v3
                            _ = v39
                        case -1:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            var v40 int64 = v6
                            _ = v40
                        }
                    default:
                        if (65535 == 65535) {
                            fmt.Println(v32)
                            if len(v4) > 0 {
                                v4[0] = (^(v1))
                            }
                            v1 -= (fi34 | (-(v1)))
                        }
                    }
                    if (43 != 255) {
                        v32 += v32
                        if len(v4) > 0 {
                            v4[0] = v1
                        }
                    } else {
                        v41 := GenS2{F0: 944, F1: 538}
                        v42 := GenS4{F0: (v6 ^ v6), F1: 42, F2: v6}
                        _ = v41
                        _ = v42
                    }
                }
                v6 += 1
            }
        }
        _ = fi34
    }
    if ((v2 % (((v2 / (((v2 | 833)) | 1))) | 1)) != 255) {
        _ = len(v5)
        var v43 bool = true
        v6 = (v6 / (((v6 * v6)) | 1))
        _ = v43
    }
    v44 := GenS5{F0: 848}
    delete(v5, v1)
    if v17 {
        v6 = v6
        v45 := v5[(v1 - 2147483647)]
        v17 = ("a" != v3)
        _ = v45
    }
    for fi46 := 0; fi46 < 4; fi46++ {
        for fi47 := 0; fi47 < 14; fi47++ {
            v48 := GenS1{F0: 4294967295, F1: 65535, F2: 255}
            var v49 int64 = 9223372036854775807
            v3 = v3
            v50 := GenS2{F0: 2, F1: 4294967295}
            v51 := GenS2{F0: 294, F1: 872}
            _ = fi47
            _ = v48
            _ = v49
            _ = v50
            _ = v51
        }
        _ = fi46
    }
    v32 -= (v32 - float64(float32(float64(float32(v32)))))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
