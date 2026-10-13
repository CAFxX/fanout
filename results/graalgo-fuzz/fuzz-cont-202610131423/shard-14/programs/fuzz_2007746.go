package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 int64 = -1
    var v2 int32 = 1
    v3 := make(map[string]bool)
    var v4 int32 = (v2 & v2)
    v5 := []string{(("hello" + ((("\n\t" + "a") + "") + "\n\t")) + "\n\t")}
    var v6 uint16 = 2
    var v7 uint = 375
    var v8 uint16 = (v6 / ((255) | 1))
    v7 = (v7 / ((252) | 1))
    var v9 uint8 = 0
    v7 += (v7 / (((^(v7))) | 1))
    if false {
        var v10 int8 = 127
        v10 = (v10 & 100)
        _ = v10
    } else {
        fmt.Println(v4)
    }
    v1 = v1
    if (v7 >= (v7 + (^(v7)))) {
        v2 += v2
        if false {
            switch 128 {
            case 4294967295:
                var v11 float32 = ((float32(1e10) / (-2.25 * float32(float64((float32(-2.25) + 3.14159))))) / 1e10)
                _ = v11
            case 223:
                if (!(!true)) {
                    for fi12 := 0; fi12 < 20; fi12++ {
                        var v13 int32 = (v2 % ((int32(v7)) | 1))
                        v14 := GenS1{F0: v4, F1: (-(v1))}
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    if false {
                        for fi15 := 0; fi15 < 12; fi15++ {
                            if false {
                                var v16 uint = 239
                                if ((147 < 2) || (!true)) {
                                    v17 := GenS1{F0: (v4 % (((v2 << ((v16) & 63))) | 1)), F1: (v1 * v1)}
                                    var v18 float32 = (float32(float64(3.14159)) + 1e10)
                                    _ = v17
                                    _ = v18
                                }
                                _ = v16
                            } else {
                                switch "hello" {
                                case "a":
                                    v6 = v6
                                case "foo bar":
                                    v5 = append(v5, ("" + ""))
                                case "\n\t":
                                    fmt.Println(v9)
                                default:
                                    v2 -= -2147483648
                                }
                                var v19 bool = false
                                _ = v19
                            }
                            _ = fi15
                        }
                        if (!(v1 > 737)) {
                            v9 = (v9 % (((v9 >> ((1) & 63))) | 1))
                        } else {
                            v8 -= (v6 << (((v7 % ((v7) | 1))) & 63))
                            var v20 int16 = -1
                            for fi21 := 0; fi21 < 17; fi21++ {
                                if (((("a" + "\n\t") + "") + "a") <= "\n\t") {
                                    fmt.Println(v7, fi21, v2)
                                    v8 = 1
                                    var v22 uint8 = 236
                                    _ = v22
                                } else {
                                    v23 := GenS1{F0: (v2 & (v2 | v2)), F1: 924}
                                    fmt.Println(v20, v2, v20)
                                    _ = len(v5)
                                    _ = v23
                                }
                                v2 = v2
                                _ = fi21
                            }
                            _ = v20
                        }
                        var v24 string = "foo bar"
                        _ = v24
                    } else {
                        switch 2147483647 {
                        case -1:
                            var v25 int = 289
                            _ = v25
                        case 569:
                            var v26 float32 = 1.5
                            _ = v26
                        default:
                            v3[("" + (("\n\t" + "hello") + "hello"))] = true
                        }
                    }
                }
            case 639:
                for fi27 := 0; fi27 < 5; fi27++ {
                    v28 := GenS1{F0: v4, F1: (v1 << (((v7 << ((v7) & 63))) & 63))}
                    _ = fi27
                    _ = v28
                }
            }
            v5 = append(v5, ("foo bar" + ("" + "")))
            if (728 == 0) {
                var v29 float32 = 1e10
                _ = v29
            }
        } else {
            v6 = v8
            var v30 string = ("a" + "a")
            for fi31 := 0; fi31 < 11; fi31++ {
                for fi32 := 0; fi32 < 5; fi32++ {
                    if false {
                        if len(v5) > 0 {
                            v5[0] = ((v30 + v30) + (v30 + v30))
                        }
                    } else {
                        if (true || (!true)) {
                            v7 = (^(v7))
                            var v33 float64 = (3.14159 + (float64(float32(((0.1 * ((1.5 - (-2.25 + 0.1)) / 0.1)) + 0.1))) + float64(0.1)))
                            var v34 uint64 = 237
                            _ = v33
                            _ = v34
                        } else {
                            if ((-128 == 42) && ((v30 + v30) >= v30)) {
                                v35 := GenS1{F0: v2, F1: (^(v1))}
                                var v36 int16 = -463
                                fmt.Println(v1, v2, v8)
                                v37 := GenS2{F0: (v36 / (((v36 << ((v7) & 63))) | 1)), F1: int64(v7)}
                                _ = v35
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = len(v5)
                        fmt.Println(v2, v30, v9)
                        var v38 string = ("foo bar" + "foo bar")
                        switch "foo bar" {
                        case "a":
                            v4 -= (v2 % (((v2 | (v2 * 42))) | 1))
                        }
                        _ = v38
                    }
                    _ = fi32
                }
                _ = len(v5)
                _ = fi31
            }
            _ = v30
        }
    } else {
        var v39 int8 = 0
        _ = v39
    }
    var v40 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
