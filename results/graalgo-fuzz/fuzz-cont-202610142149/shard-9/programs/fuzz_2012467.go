package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int64; F2 int64 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint64; F1 int32; F2 uint }

func main() {
    var v1 bool = false
    v2 := []string{"hello", "a", "foo bar", (("\n\t" + "foo bar") + (("a" + "\n\t") + ("a" + "a")))}
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(fi4, v1)
        var v5 uint = 684
        if v1 {
            v1 = ("\n\t" > "foo bar")
            v3[(fi4 ^ fi4)] = "\n\t"
            v2 = append(v2, "a")
        } else {
            delete(v3, fi4)
            if (0.0 != 1e10) {
                for fi6 := 0; fi6 < 4; fi6++ {
                    if (258 < fi6) {
                        v7 := GenS1{F0: 160}
                        _ = v7
                    }
                    _ = fi6
                }
                switch fi4 {
                case 0:
                    for fi8 := 0; fi8 < 14; fi8++ {
                        switch "\n\t" {
                        case "foo bar":
                            v1 = (32767 < 641)
                        case "\n\t":
                            v3[(fi8 / (((fi8 & (fi8 >> ((42) & 63)))) | 1))] = ("a" + "foo bar")
                        default:
                            if v1 {
                                v9 := GenS1{F0: 42}
                                var v10 uint = 255
                                _ = v10
                                _ = v9
                            }
                        }
                        if ((127 <= 209) && (((!(42 > 2147483647)) && false) && false)) {
                            v11 := GenS1{F0: 632}
                            var v12 bool = true
                            if v1 {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                            } else {
                                fmt.Println(v12, fi8)
                            }
                            _ = v11
                            _ = v12
                        }
                        var v13 int16 = 42
                        _ = fi8
                        _ = v13
                    }
                }
            } else {
                v2 = append(v2, (((("" + "") + (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "")) + (("a" + "hello") + "\n\t")) + ""))
            }
            if v1 {
                for fi14 := 0; fi14 < 2; fi14++ {
                    v5 = (v5 & 0)
                    _ = fi14
                }
                var v15 int64 = 9223372036854775807
                _ = v15
            } else {
                if v1 {
                    v3[(fi4 >> ((v5) & 63))] = ("a" + (("\n\t" + "a") + "\n\t"))
                    v16 := GenS1{F0: 1}
                    fmt.Println(v5, v1)
                    _ = v16
                }
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 bool = v1
                    if v1 {
                        _ = len(v2)
                    }
                    _ = fi17
                    _ = v18
                }
                var v19 float32 = float32(3.14159)
                _ = v19
            }
        }
        _ = fi4
        _ = v5
    }
    v1 = (!false)
    v20 := GenS2{F0: -9223372036854775808, F1: 0, F2: -510}
    var v21 string = "\n\t"
    switch 710 {
    case -459:
        _ = len(v3)
    case 1:
        var v22 float64 = 1e10
        _ = v22
    case 32767:
        v1 = true
    default:
        fmt.Println(v1)
    }
    v21 = (v21 + "\n\t")
    v3[127] = ("\n\t" + ("\n\t" + (v21 + v21)))
    _ = len(v3)
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 4; fi24++ {
            if (1 > 4294967295) {
                var v25 uint32 = 295
                _ = v25
            } else {
                v26 := GenS3{F0: 609, F1: 2}
                fmt.Println(fi23, fi24)
                _ = v26
            }
            if v1 {
                v27 := GenS1{F0: 1}
                _ = v27
            } else {
                var v28 int32 = -571
                _ = v28
            }
            v1 = v1
            switch fi24 {
            case -1:
                switch v21 {
                case "":
                    var v29 bool = (!v1)
                    _ = v29
                default:
                    var v30 string = v21
                    _ = v30
                }
                if false {
                    switch fi24 {
                    case 127:
                        var v31 int8 = 127
                        _ = v31
                    case -541:
                        v1 = false
                    }
                }
            case 2147483647:
                var v32 uint8 = 12
                _ = v32
            default:
                v33 := GenS4{F0: 4294967295, F1: 0, F2: 2}
                _ = v33
            }
            var v34 uint16 = 42
            _ = fi24
            _ = v34
        }
        v21 = "hello"
        _ = fi23
    }
    switch 32768 {
    case -984:
        var v35 float64 = (3.14159 + 0.0)
        _ = v35
    case -1:
        var v36 float64 = (-2.25 * (-2.25 / (0.0 + 3.14159)))
        _ = v36
    default:
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
