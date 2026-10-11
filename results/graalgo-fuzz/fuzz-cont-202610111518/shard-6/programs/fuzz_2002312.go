package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 string; F1 int; F2 uint }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint16; F1 int8 }

func main() {
    var v1 bool = false
    var v2 float32 = -2.25
    var v3 int16 = -1
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := []int{(fi4 + (fi4 * fi4))}
        v5 = append(v5, (fi4 * fi4))
        _ = fi4
        _ = v5
    }
    v6 := []int{}
    fmt.Println(v2)
    var v7 int8 = -128
    v8 := GenS1{F0: 2, F1: 18446744073709551615, F2: 57}
    if (1e10 == 0.1) {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        for fi9 := 0; fi9 < 17; fi9++ {
            v1 = v1
            var v10 bool = ((fi9 != 32768) || v1)
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v3, v2, v1)
    v1 = (2147483647 == 80)
    v2 = v2
    v7 += (v7 >> ((926) & 63))
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 int8 = (v7 / ((v7) | 1))
        v1 = false
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi11
        _ = v12
    }
    _ = len(v6)
    if ((v3 % (((v3 ^ (-(v3)))) | 1)) >= v3) {
        var v13 uint32 = 297
        _ = v13
    }
    var v14 int = 128
    if (0 == 360) {
        if false {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            if len(v6) > 0 {
                v6[0] = (v14 >> ((18446744073709551615) & 63))
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        } else {
            var v15 uint8 = 127
            if v1 {
                switch "hello" {
                case "":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        if (-859 <= v3) {
                            for fi17 := 0; fi17 < 15; fi17++ {
                                if false {
                                    if v1 {
                                        var v18 int = (v14 ^ fi16)
                                        _ = v18
                                    } else {
                                        v14 = 198
                                    }
                                    fmt.Println(fi17, fi16)
                                }
                                _ = fi17
                            }
                            v19 := GenS1{F0: 624, F1: 2, F2: 2}
                            _ = len(v6)
                            _ = v19
                        } else {
                            fmt.Println(v3, v14, fi16)
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(fi16)
                            case "a":
                                v6 = append(v6, (v14 / ((-288) | 1)))
                            default:
                                if (true && v1) {
                                    v15 -= 128
                                    var v20 int8 = (-(v7))
                                    var v21 string = "a"
                                    _ = v20
                                    _ = v21
                                } else {
                                    v22 := GenS2{F0: (("" + ("a" + "")) + ""), F1: fi16, F2: 2}
                                    _ = v22
                                }
                            }
                            for fi23 := 0; fi23 < 15; fi23++ {
                                switch 0 {
                                case -773:
                                    v7 += 100
                                default:
                                    var v24 int32 = 2147483647
                                    _ = v24
                                }
                                if false {
                                    v1 = ((v15 != v15) || (v1 && false))
                                    var v25 string = ("a" + ("foo bar" + "foo bar"))
                                    if (((v1 && (false || v1)) && v1) && v1) {
                                        v3 = 42
                                        if len(v6) > 0 {
                                            v6[0] = (fi16 % (((fi16 << ((255) & 63))) | 1))
                                        }
                                    } else {
                                        fmt.Println(fi23, v14)
                                        switch "a" {
                                        case "foo bar":
                                            if v1 {
                                                v26 := GenS3{F0: (fi23 | fi16), F1: ("a" + "a")}
                                                var v27 uint16 = 255
                                                var v28 int16 = (v3 ^ (v3 * v3))
                                                _ = v26
                                                _ = v27
                                                _ = v28
                                            } else {
                                                fmt.Println(v3, v2, fi23)
                                            }
                                        case "\n\t":
                                            if v1 {
                                                var v29 int = (^(v14))
                                                if len(v6) > 0 {
                                                    v6[0] = (v29 & int(v3))
                                                }
                                                _ = len(v6)
                                                _ = v29
                                            }
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                            v1 = (-813 < 9223372036854775807)
                                        case "":
                                            fmt.Println(v3)
                                        }
                                        if (347 < 464) {
                                            var v30 uint8 = v15
                                            v31 := GenS4{F0: v3, F1: (v3 + (v3 - v3))}
                                            v6 = append(v6, (fi16 % (((fi23 / (((^(fi16))) | 1))) | 1)))
                                            _ = v30
                                            _ = v31
                                        }
                                    }
                                    _ = v25
                                } else {
                                    fmt.Println(v14, v1, v15)
                                }
                                _ = fi23
                            }
                        }
                        if (!v1) {
                            fmt.Println(v15, v14, fi16)
                            v2 = ((v2 - (v2 - 1.5)) * v2)
                            v3 = v3
                        }
                        var v32 float32 = v2
                        _ = fi16
                        _ = v32
                    }
                }
                _ = len(v6)
            } else {
                v15 += 127
            }
            _ = v15
        }
        v6 = append(v6, v14)
    } else {
        if ("" > "\n\t") {
            var v33 uint = 1
            _ = v33
        }
        if true {
            var v34 float64 = 1e10
            v35 := GenS3{F0: v14, F1: "foo bar"}
            if (!v1) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                fmt.Println(v14, v3, v14)
                _ = len(v6)
            }
            _ = v34
            _ = v35
        }
    }
    var v36 int32 = -70
    v37 := GenS5{F0: 65535, F1: v7}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v6
    _ = v7
    _ = v8
}
