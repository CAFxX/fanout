package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint }
type GenS5 struct { F0 int64; F1 uint; F2 uint16 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0(p0 uint32) int8 {
    p0 = p0
    p0 = p0
    if true {
        p0 -= p0
        if (!(-1 == -423)) {
            var v7 uint16 = 543
            for fi8 := 0; fi8 < 20; fi8++ {
                v7 -= (v7 | (v7 % (((^(v7))) | 1)))
                var v9 int16 = 235
                switch "" {
                case "hello":
                    v10 := []int{(fi8 / ((fi8) | 1)), (fi8 + fi8)}
                    _ = v10
                default:
                    if false {
                        if ((!false) && false) {
                            var v11 uint32 = 1
                            v12 := []int{(fi8 * (fi8 << ((461) & 63))), (fi8 / (((fi8 * fi8)) | 1)), fi8, fi8}
                            _ = v11
                            _ = v12
                        } else {
                            v13 := GenS1{F0: "foo bar", F1: 1}
                            fmt.Println(fi8, fi8, v7)
                            v14 := GenS2{F0: ("a" + (("foo bar" + "foo bar") + ("\n\t" + "a")))}
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS1{F0: ("\n\t" + ""), F1: 55}
                        _ = v15
                    }
                }
                _ = fi8
                _ = v9
            }
            switch "a" {
            case "a":
                v16 := GenS3{F0: -9223372036854775808}
                _ = v16
            case "foo bar":
                switch "\n\t" {
                case "foo bar":
                    var v17 string = "a"
                    _ = v17
                case "\n\t":
                    switch 42 {
                    case 32767:
                        var v18 string = "foo bar"
                        _ = v18
                    case 4294967295:
                        v19 := GenS4{F0: uint32(v7), F1: 104, F2: uint(p0)}
                        _ = v19
                    default:
                        var v20 uint32 = (p0 * 65535)
                        _ = v20
                    }
                case "hello":
                    v21 := GenS2{F0: "a"}
                    _ = v21
                }
            default:
                v22 := GenS5{F0: 9223372036854775807, F1: 42, F2: v7}
                _ = v22
            }
            _ = v7
        } else {
            p0 += p0
            for fi23 := 0; fi23 < 17; fi23++ {
                var v24 uint64 = 936
                if true {
                    var v25 int16 = 42
                    switch fi23 {
                    case 2147483647:
                        fmt.Println(fi23, fi23)
                    default:
                        var v26 float64 = float64((float32((-2.25 * float64(0.0))) / -2.25))
                        _ = v26
                    }
                    v27 := GenS2{F0: "foo bar"}
                    _ = v25
                    _ = v27
                } else {
                    v28 := GenS5{F0: int64(fi23), F1: 18446744073709551615, F2: 308}
                    if ((-2.25 / 3.14159) < 0.1) {
                        v24 -= (v24 | 4294967295)
                        fmt.Println(p0, v24, p0)
                    }
                    if (!((v24 <= v24) && false)) {
                        var v29 int64 = 317
                        if (40 > 1) {
                            for fi30 := 0; fi30 < 16; fi30++ {
                                var v31 float64 = 3.14159
                                var v32 int = 446
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            v24 = 978
                            p0 = p0
                        } else {
                            fmt.Println(v24, v29)
                            v24 = v24
                            if false {
                                for fi33 := 0; fi33 < 3; fi33++ {
                                    fmt.Println(v24, v29, fi33)
                                    v34 := GenS6{F0: -1, F1: 1, F2: 65535}
                                    _ = fi33
                                    _ = v34
                                }
                                var v35 float32 = ((3.14159 * float32((1.5 + (0.1 + float64((-2.25 + ((1.5 + -2.25) + ((1.5 * (0.0 / 1.0)) - 0.0)))))))) + 1e10)
                                _ = v35
                            } else {
                                var v36 uint16 = 255
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    var v38 int16 = int16(p0)
                                    switch fi37 {
                                    case 4294967295:
                                        if true {
                                            p0 -= (p0 - p0)
                                        } else {
                                            v39 := []int{fi23, fi37, -1, 32768}
                                            var v40 uint = 191
                                            _ = v39
                                            _ = v40
                                        }
                                    default:
                                        v41 := GenS6{F0: (v29 - 1), F1: v36, F2: v36}
                                        _ = v41
                                    }
                                    v42 := GenS3{F0: (v29 ^ (v29 >> ((65535) & 63)))}
                                    fmt.Println(p0, fi37, fi23)
                                    _ = fi37
                                    _ = v38
                                    _ = v42
                                }
                                _ = v36
                            }
                        }
                        p0 += p0
                        _ = v29
                    }
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
        }
    }
    v43 := GenS6{F0: -9223372036854775808, F1: 204, F2: 194}
    _ = p0
    _ = v43
    return 127
}

func main() {
    var v1 string = "hello"
    var v2 int8 = -1
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    v5 := []string{v1, ((v1 + ((v1 + v1) + v1)) + v1)}
    v6 := make(map[int]string)
    v2 += (v2 % ((127) | 1))
    var v44 int16 = 0
    v3 = v3
    if (0.1 <= float32(3.14159)) {
        if len(v5) > 0 {
            v5[0] = v1
        }
        var v45 string = "a"
        _ = v45
    }
    switch 0 {
    case 1:
        if len(v5) > 0 {
            v5[0] = ("a" + "\n\t")
        }
    default:
        fmt.Println(v3, v44)
    }
    var v46 float64 = (-2.25 + 3.14159)
    var v47 int16 = v44
    switch 744 {
    case -699:
        v48 := GenS5{F0: -9223372036854775808, F1: 18446744073709551615, F2: 982}
        _ = v48
    case 42:
        var v49 bool = ((!(v2 != genFn0(2147483648))) && true)
        _ = v49
    default:
        fmt.Println(v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
