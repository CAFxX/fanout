package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 bool; F1 int16; F2 int }
func genFn0(p0 uint32) int32 {
    p0 += (p0 & (p0 << ((42) & 63)))
    v4 := []int{9223372036854775807}
    if true {
        for fi5 := 0; fi5 < 7; fi5++ {
            p0 += p0
            v4 = append(v4, 23)
            _ = fi5
        }
        v4 = append(v4, 128)
        var v6 uint16 = 32767
        _ = v6
    }
    _ = p0
    _ = v4
    return -2147483648
}

func main() {
    var v1 float64 = -2.25
    var v2 uint = 976
    v3 := make(map[int]string)
    v1 += (float64(0.1) + (v1 / v1))
    v7 := []int{2147483647, -1, 9223372036854775807, 1}
    fmt.Println(v2, v2)
    for fi8 := 0; fi8 < 4; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 int = fi9
            delete(v3, v10)
            if true {
                var v11 int = fi9
                fmt.Println(fi9, v2, v2)
                _ = len(v3)
                fmt.Println(v10, v1, v10)
                _ = v11
            }
            v1 += (0.0 * 1e10)
            v12 := GenS1{F0: 1, F1: 65535}
            _ = fi9
            _ = v10
            _ = v12
        }
        fmt.Println(v1, v2, fi8)
        var v13 uint = (^(v2))
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v2 = v2
        v1 -= float64(1e10)
        _ = fi14
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 bool = (v2 < v2)
            if v17 {
                if v17 {
                    var v18 string = "a"
                    v2 += (^(v2))
                    _ = v18
                }
                v19 := GenS2{F0: "", F1: 100}
                _ = len(v3)
                _ = v19
            } else {
                if (-269 != 518) {
                    switch fi16 {
                    case -488:
                        v1 -= -2.25
                    default:
                        _ = len(v3)
                    }
                    if v17 {
                        var v20 int32 = 2147483647
                        v2 = (v2 / ((255) | 1))
                        _ = v20
                    } else {
                        var v21 uint16 = 42
                        switch (-(fi15)) {
                        case 9223372036854775807:
                            if len(v7) > 0 {
                                v7[0] = fi16
                            }
                        }
                        switch (fi15 << (((^(v2))) & 63)) {
                        case 42:
                            fmt.Println(v2, v17)
                        case 32767:
                            _ = len(v7)
                        default:
                            v22 := GenS1{F0: 1, F1: (v21 % (((v21 - 32767)) | 1))}
                            _ = v22
                        }
                        v23 := GenS3{F0: (-168 < 2147483647), F1: 1000, F2: fi15}
                        if (-91 == -1) {
                            delete(v3, fi16)
                            v3[-2147483648] = "\n\t"
                            var v24 bool = v17
                            _ = v24
                        } else {
                            v17 = (-1 == 0)
                            switch (("a" + "") + (("foo bar" + "a") + "foo bar")) {
                            case "":
                                switch ("" + "a") {
                                case "a":
                                    v25 := GenS2{F0: (("" + ("foo bar" + "foo bar")) + "\n\t"), F1: 0}
                                    _ = v25
                                case "\n\t":
                                    switch (fi15 / ((fi15) | 1)) {
                                    case 0:
                                        switch "a" {
                                        case "\n\t":
                                            if v17 {
                                                var v26 uint = 4294967295
                                                _ = v26
                                            } else {
                                                _ = len(v3)
                                                _ = len(v3)
                                            }
                                        case "foo bar":
                                            fmt.Println(v2, v21)
                                        case "a":
                                            var v27 string = ("\n\t" + "a")
                                            _ = v27
                                        default:
                                            var v28 float64 = (-2.25 + v1)
                                            _ = v28
                                        }
                                    default:
                                        fmt.Println(v2, v21)
                                    }
                                default:
                                    v17 = (!(-1 > -32768))
                                }
                            default:
                                if v17 {
                                    v2 += (v2 % ((uint(v21)) | 1))
                                }
                            }
                        }
                        _ = v21
                        _ = v23
                    }
                } else {
                    v17 = v17
                    v29 := GenS1{F0: -344, F1: 42}
                    _ = v29
                }
            }
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v30 := GenS1{F0: 0, F1: 255}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if len(v7) > 0 {
        v7[0] = 186
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v7
}
