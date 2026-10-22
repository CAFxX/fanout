package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 string; F1 uint32; F2 uint64 }
func genFn0() uint16 {
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 136, F1: 255}
        var v7 uint32 = 0
        var v8 uint8 = 255
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{128, -819, 32767, 4294967295}
    _ = v9
    return 863
}
func genFn1(p0 int32, p1 uint64) int64 {
    fmt.Println(p1)
    for fi10 := 0; fi10 < 16; fi10++ {
        p0 += p0
        for fi11 := 0; fi11 < 14; fi11++ {
            v12 := GenS2{F0: 256, F1: -32768}
            _ = fi11
            _ = v12
        }
        v13 := GenS3{F0: (!true), F1: 42}
        _ = fi10
        _ = v13
    }
    fmt.Println(p0)
    switch 2147483647 {
    case 1:
        for fi14 := 0; fi14 < 13; fi14++ {
            if ((2147483647 <= 935) || ((fi14 + int(p1)) < fi14)) {
                for fi15 := 0; fi15 < 7; fi15++ {
                    p1 = p1
                    v16 := []int{fi15, fi15, (^(fi15)), (fi15 / ((int(p0)) | 1))}
                    if (106 > 255) {
                        v17 := GenS3{F0: false, F1: -37}
                        if true {
                            if (174 > 128) {
                                p1 = 1
                            } else {
                                var v18 int = fi15
                                _ = v18
                            }
                            if false {
                                p1 = (p1 ^ p1)
                                var v19 uint64 = (^(p1))
                                if true {
                                    fmt.Println(v19, v19, v19)
                                    var v20 int16 = -381
                                    _ = v20
                                }
                                _ = v19
                            }
                            switch "foo bar" {
                            case "a":
                                v21 := GenS1{F0: 2, F1: 32767}
                                _ = v21
                            default:
                                var v22 uint32 = uint32(p1)
                                if false {
                                    if (false && true) {
                                        if (32767 < 2) {
                                            var v23 uint32 = v22
                                            var v24 int32 = -1
                                            _ = v23
                                            _ = v24
                                        }
                                        var v25 uint8 = uint8(p1)
                                        _ = v25
                                    }
                                    if false {
                                        p1 = p1
                                    } else {
                                        v26 := GenS1{F0: (v22 >> ((662) & 63)), F1: 65535}
                                        p0 += 42
                                        var v27 uint8 = 2
                                        _ = v26
                                        _ = v27
                                    }
                                    v22 = (v22 - v22)
                                }
                                if true {
                                    if (!(("\n\t" + ("foo bar" + "\n\t")) >= "foo bar")) {
                                        if (164 < 127) {
                                            _ = len(v16)
                                            fmt.Println(p0, p0)
                                        } else {
                                            var v28 bool = true
                                            fmt.Println(v22, v22)
                                            _ = v28
                                        }
                                    } else {
                                        if true {
                                            var v29 uint = 428
                                            _ = v29
                                        } else {
                                            fmt.Println(p0, fi14, fi14)
                                            var v30 bool = (!((215 > genFn0()) && ((v22 | 0) > (^(v22)))))
                                            _ = v30
                                        }
                                        var v31 string = "hello"
                                        _ = v31
                                    }
                                } else {
                                    var v32 float32 = 3.14159
                                    var v33 bool = ("foo bar" == "hello")
                                    fmt.Println(v32, fi14)
                                    _ = v32
                                    _ = v33
                                }
                                _ = v22
                            }
                        }
                        p0 = (p0 / ((p0) | 1))
                        _ = v17
                    } else {
                        p1 = p1
                        var v34 uint64 = p1
                        _ = v34
                    }
                    _ = fi15
                    _ = v16
                }
                v35 := GenS4{F0: ("\n\t" + "\n\t"), F1: 931, F2: p1}
                _ = v35
            } else {
                var v36 int16 = -1
                v37 := GenS2{F0: 171, F1: (v36 * (v36 - (v36 ^ (v36 << ((42) & 63)))))}
                var v38 uint8 = 0
                _ = v36
                _ = v37
                _ = v38
            }
            var v39 float32 = (0.1 - (-2.25 - (0.1 + 1e10)))
            for fi40 := 0; fi40 < 13; fi40++ {
                v41 := []int{(fi14 | int(p1)), (fi40 << ((1) & 63)), fi40}
                fmt.Println(p1, v39, fi40)
                _ = fi40
                _ = v41
            }
            _ = fi14
            _ = v39
        }
    case -1:
        for fi42 := 0; fi42 < 2; fi42++ {
            switch -602 {
            case 32768:
                v43 := []int{-1}
                _ = v43
            case 128:
                fmt.Println(p1, p1)
            default:
                p0 = 2147483647
            }
            for fi44 := 0; fi44 < 5; fi44++ {
                switch 9223372036854775807 {
                case 389:
                    fmt.Println(fi42)
                default:
                    p1 = p1
                }
                p0 = (p0 / (((p0 % (((p0 % ((0) | 1))) | 1))) | 1))
                _ = fi44
            }
            var v45 int32 = p0
            _ = fi42
            _ = v45
        }
    }
    var v46 int32 = p0
    _ = p0
    _ = p1
    _ = v46
    return 9223372036854775807
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]string)
    var v3 float64 = (float64(1e10) * 0.1)
    var v4 uint64 = 2
    v4 += (^(v4))
    v4 += (^(v4))
    if true {
        if false {
            v3 = v3
        }
    }
    v47 := []int{-965, -1}
    var v48 uint64 = v4
    for fi49 := 0; fi49 < 17; fi49++ {
        var v50 int = (-(fi49))
        _ = fi49
        _ = v50
    }
    var v51 uint = 39
    var v52 uint32 = 0
    fmt.Println(v3, v48)
    v4 -= v48
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v47:", len(v47), v47)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v51
    _ = v52
}
