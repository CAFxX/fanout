package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 uint16 }
func genFn0(p0 int64) int32 {
    p0 = p0
    p0 = p0
    v5 := GenS1{F0: 20, F1: -2147483648, F2: false}
    p0 = p0
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 string = "\n\t"
    v2 := []string{((v1 + v1) + "a"), "foo bar", v1, v1}
    v3 := make(map[string]string)
    var v4 int8 = 0
    if (4294967295 >= 1) {
        _ = len(v2)
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        v4 = v4
        v4 = v4
        _ = fi6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        v2 = append(v2, v1)
        if true {
            fmt.Println(fi7, fi7)
        } else {
            var v8 int16 = -619
            v9 := GenS1{F0: 13, F1: (fi7 % (((fi7 >> ((1) & 63))) | 1)), F2: false}
            var v10 int64 = 1
            _ = v10
            _ = v8
            _ = v9
        }
        if false {
            var v11 float32 = 3.14159
            _ = v11
        }
        _ = fi7
    }
    fmt.Println(v4, v4)
    v12 := GenS1{F0: -521, F1: 9223372036854775807, F2: false}
    v13 := GenS1{F0: 0, F1: 32767, F2: (!(0 <= 2147483647))}
    var v14 int8 = v4
    var v15 uint = 442
    switch v1 {
    case "":
        v1 = v1
    default:
        if true {
            v16 := GenS1{F0: 898, F1: 273, F2: false}
            fmt.Println(v15)
            switch "\n\t" {
            case "foo bar":
                switch -962 {
                case 2147483647:
                    switch "hello" {
                    case "hello":
                        for fi17 := 0; fi17 < 4; fi17++ {
                            fmt.Println(v15, v4, v4)
                            if (645 < 42) {
                                var v18 int = fi17
                                _ = len(v2)
                                _ = v18
                            } else {
                                v2 = append(v2, v1)
                                v14 += v4
                            }
                            var v19 int64 = 333
                            _ = fi17
                            _ = v19
                        }
                    case "a":
                        var v20 string = ("" + (("\n\t" + v1) + "\n\t"))
                        _ = v20
                    default:
                        var v21 float32 = 3.14159
                        _ = v21
                    }
                default:
                    v15 = v15
                }
            case "a":
                fmt.Println(v4)
            case "hello":
                for fi22 := 0; fi22 < 13; fi22++ {
                    var v23 float64 = (0.1 / 0.1)
                    var v24 int64 = 301
                    v15 = v15
                    _ = fi22
                    _ = v23
                    _ = v24
                }
            default:
                var v25 float32 = (-2.25 - (0.1 - (-2.25 + 1e10)))
                _ = v25
            }
            _ = v16
        } else {
            if true {
                var v26 bool = false
                fmt.Println(v26)
                v14 = v4
                _ = v26
            } else {
                var v27 int = 1
                v1 = "a"
                var v28 int16 = int16(v15)
                _ = v27
                _ = v28
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                switch "foo bar" {
                case "foo bar":
                    v30 := GenS1{F0: -9223372036854775808, F1: 127, F2: (!true)}
                    _ = v30
                case "":
                    var v31 uint16 = 65535
                    _ = v31
                case "hello":
                    if false {
                        v3[v1] = ((("foo bar" + ("foo bar" + "\n\t")) + "hello") + (v1 + "a"))
                        v32 := GenS2{F0: (!(true && true)), F1: fi29, F2: 2}
                        var v33 int32 = 1
                        _ = v32
                        _ = v33
                    }
                default:
                    v1 = (((v1 + "\n\t") + ("foo bar" + ("a" + ("" + ("" + "hello"))))) + (v1 + v1))
                }
                var v34 int16 = 32767
                for fi35 := 0; fi35 < 5; fi35++ {
                    switch "" {
                    case "a":
                        var v36 int16 = v34
                        _ = v36
                    default:
                        var v37 float64 = 3.14159
                        _ = v37
                    }
                    v15 -= v15
                    v38 := GenS1{F0: 1, F1: (^(fi29)), F2: true}
                    _ = len(v3)
                    _ = fi35
                    _ = v38
                }
                _ = fi29
                _ = v34
            }
            if (103 >= 4294967295) {
                v39 := GenS1{F0: 0, F1: 2147483647, F2: true}
                fmt.Println(v4, v14)
                delete(v3, v1)
                _ = v39
            }
        }
    }
    v40 := GenS3{F0: int32(v4), F1: -713, F2: 32767}
    if false {
        var v41 uint32 = 65535
        v42 := GenS3{F0: 42, F1: 0, F2: 65535}
        switch "hello" {
        case "a":
            var v43 int64 = 9223372036854775807
            _ = v43
        case "foo bar":
            for fi44 := 0; fi44 < 16; fi44++ {
                var v45 uint8 = 191
                _ = fi44
                _ = v45
            }
        case "":
            var v46 int8 = 88
            _ = v46
        }
        _ = v41
        _ = v42
    } else {
        fmt.Println(v14, v14, v1)
    }
    var v47 bool = (false || (false && true))
    v2 = append(v2, ((v1 + ("a" + "foo bar")) + "foo bar"))
    for fi48 := 0; fi48 < 8; fi48++ {
        for fi49 := 0; fi49 < 18; fi49++ {
            v1 = v1
            var v50 int8 = (v14 % ((v4) | 1))
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        v14 += v4
        v4 = 59
        fmt.Println(v4, v1)
        _ = fi51
    }
    for fi52 := 0; fi52 < 14; fi52++ {
        for fi53 := 0; fi53 < 19; fi53++ {
            if v47 {
                v2 = append(v2, "\n\t")
            }
            if ((65535 < 2) && (!(!(v15 > (^(v15)))))) {
                var v54 bool = true
                _ = v54
            }
            _ = fi53
        }
        _ = fi52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v47
}
