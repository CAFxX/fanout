package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 uint64; F1 uint32; F2 uint16 }
func genFn0(p0 int8, p1 uint8, p2 int) uint64 {
    p1 += 249
    var v4 uint32 = 42
    if false {
        v5 := GenS1{F0: 283, F1: 42, F2: (v4 & 928)}
        var v6 int8 = (p0 - (-(p0)))
        _ = v5
        _ = v6
    }
    if false {
        v7 := GenS1{F0: 18446744073709551615, F1: -334, F2: v4}
        var v8 int32 = 1
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 25
}
func genFn1(p0 int64) uint64 {
    switch "foo bar" {
    case "\n\t":
        v9 := GenS2{F0: 42}
        _ = v9
    case "":
        var v10 uint32 = 2147483648
        _ = v10
    default:
        var v11 int16 = -1
        _ = v11
    }
    var v12 uint64 = 18446744073709551615
    fmt.Println(p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = v12
    return (^(v12))
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = -2.25
    var v3 int16 = 42
    switch -889 {
    case 31:
        v13 := []int{128, 558}
        _ = v13
    }
    v14 := GenS2{F0: 1}
    var v15 bool = (!(false || false))
    v16 := []int{127, -2147483648, 32767, 32767}
    for fi17 := 0; fi17 < 6; fi17++ {
        if v15 {
            switch fi17 {
            case -1:
                for fi18 := 0; fi18 < 18; fi18++ {
                    v15 = v15
                    var v19 float32 = 1e10
                    v1 -= v1
                    _ = fi18
                    _ = v19
                }
            case -730:
                var v20 uint64 = 2
                _ = v20
            case 32768:
                if true {
                    v3 -= v3
                    for fi21 := 0; fi21 < 15; fi21++ {
                        if v15 {
                            v2 = ((v2 / ((3.14159 + (0.1 - (v2 * 1.5))) * v2)) / 1.0)
                            var v22 uint8 = 127
                            var v23 uint = 106
                            _ = v22
                            _ = v23
                        } else {
                            if (((^(v3)) <= (v3 % (((v3 - v3)) | 1))) && v15) {
                                v15 = (fi17 != 32768)
                                fmt.Println(v3, v3, v15)
                            }
                            if true {
                                var v24 float32 = (3.14159 / 3.14159)
                                var v25 uint = 2
                                _ = v24
                                _ = v25
                            } else {
                                var v26 int = (fi17 & 2147483647)
                                var v27 float32 = (v2 + ((v2 - (((-2.25 - v2) / 1.0) * 1.5)) - (1.5 - v2)))
                                var v28 string = "a"
                                var v29 uint16 = 641
                                if len(v16) > 0 {
                                    v16[0] = (v26 ^ 127)
                                }
                                _ = v26
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        }
                        if (42 > -58) {
                            var v30 bool = v15
                            if (v30 || (((0.1 - 1.5) / (v1 / 1.5)) < float64(v2))) {
                                var v31 int16 = (-(v3))
                                fmt.Println(v30, fi21, v3)
                                _ = v31
                            }
                            _ = v30
                        } else {
                            fmt.Println(v2, v3, v1)
                            if (152 == 1) {
                                if false {
                                    if (4294967295 > 45) {
                                        var v32 bool = v15
                                        _ = v32
                                    } else {
                                        v33 := GenS1{F0: 42, F1: 1, F2: 811}
                                        v3 -= 0
                                        if false {
                                            v1 += (((0.0 * (0.1 - (1.5 / v1))) * 0.0) / ((v1 / (v1 + (float64((float32((1.5 / 3.14159)) / 0.1)) + (-2.25 - 0.0)))) / (1e10 * v1)))
                                        }
                                        _ = v33
                                    }
                                    var v34 int16 = (v3 - v3)
                                    _ = v34
                                } else {
                                    v1 -= v1
                                    v16 = append(v16, (-(fi17)))
                                    switch "foo bar" {
                                    case "foo bar":
                                        if v15 {
                                            v2 = 1.5
                                            var v35 int = 128
                                            _ = v35
                                        } else {
                                            v15 = v15
                                            switch 42 {
                                            case -2147483648:
                                                v3 += v3
                                                if len(v16) > 0 {
                                                    v16[0] = fi17
                                                }
                                            }
                                        }
                                    case "\n\t":
                                        fmt.Println(fi17, fi21)
                                    case "hello":
                                        v1 += -2.25
                                    default:
                                        v36 := GenS2{F0: genFn1(9223372036854775807)}
                                        _ = v36
                                    }
                                }
                                var v37 int32 = 1
                                _ = v37
                            }
                            if false {
                                var v38 float32 = v2
                                _ = v38
                            }
                        }
                        _ = fi21
                    }
                }
            }
            _ = len(v16)
        }
        var v39 string = "foo bar"
        v15 = (65535 != 700)
        _ = fi17
        _ = v39
    }
    v3 = 261
    if (2147483647 == 652) {
        fmt.Println(v1)
    } else {
        fmt.Println(v1)
    }
    v40 := GenS3{F0: 0, F1: ("a" + "a")}
    var v41 bool = (v3 <= (v3 * (v3 & (v3 | v3))))
    v42 := GenS2{F0: genFn0(-128, 42, -1)}
    v1 += v1
    if v41 {
        v43 := GenS4{F0: 18446744073709551615, F1: 2147483648, F2: 255}
        _ = v43
    }
    _ = len(v16)
    var v44 int32 = -769
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v40
    _ = v41
    _ = v42
    _ = v44
}
