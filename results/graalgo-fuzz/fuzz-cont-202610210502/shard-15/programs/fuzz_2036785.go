package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int64; F1 uint16; F2 uint }
type GenS3 struct { F0 uint32; F1 string; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]int)
    var v3 string = ("a" + "")
    v4 := []string{(v3 + (v3 + v3))}
    v5 := make(map[string]bool)
    var v6 int16 = 1000
    v3 = ((("foo bar" + v3) + v3) + ((((v3 + v3) + v3) + (("a" + "foo bar") + ((v3 + "hello") + v3))) + (v3 + "a")))
    v3 = v3
    if (!(!(127 <= 2))) {
        v4 = append(v4, "foo bar")
        if false {
            v3 = ("hello" + v3)
            v6 -= v6
            v4 = append(v4, "foo bar")
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v3 = v3
            if (!(9223372036854775807 != 0)) {
                for fi7 := 0; fi7 < 8; fi7++ {
                    v8 := GenS1{F0: 42, F1: 627}
                    v4 = append(v4, ("hello" + v3))
                    _ = fi7
                    _ = v8
                }
                v3 = (v3 + (v3 + v3))
            }
        }
    } else {
        v6 = 0
        v9 := GenS2{F0: 42, F1: 0, F2: 1}
        _ = v9
    }
    switch 820 {
    case -849:
        if false {
            v1[609] = (2147483647 > 4294967295)
            v3 = "\n\t"
        }
    default:
        v10 := GenS3{F0: 0, F1: v3, F2: 999}
        _ = v10
    }
    var v11 bool = false
    fmt.Println(v3, v3, v11)
    var v12 int = 1
    for fi13 := 0; fi13 < 9; fi13++ {
        switch v3 {
        case "hello":
            v12 += (v12 * v12)
        default:
            v3 = ("hello" + (((("hello" + "") + v3) + (v3 + v3)) + ""))
        }
        v14 := v5[(v3 + ("\n\t" + "foo bar"))]
        for fi15 := 0; fi15 < 15; fi15++ {
            fmt.Println(fi15, fi13, v6)
            v12 += 32768
            if v11 {
                if true {
                    var v16 int8 = 1
                    var v17 int8 = 127
                    var v18 string = "\n\t"
                    _ = v16
                    _ = v17
                    _ = v18
                }
            } else {
                fmt.Println(v11, fi13, fi15)
                switch ("" + v3) {
                case "hello":
                    var v19 uint = 255
                    _ = v19
                case "\n\t":
                    v12 += 4294967295
                }
                fmt.Println(v12, v6)
                v20 := GenS4{F0: (^(v12))}
                fmt.Println(fi13)
                v6 = v6
                _ = v20
            }
            _ = fi15
        }
        _ = fi13
        _ = v14
    }
    if (474 < 862) {
        v21 := GenS4{F0: 128}
        _ = v21
    }
    var v22 float64 = ((float64(1.5) + float64(0.1)) + 1e10)
    _ = len(v4)
    for fi23 := 0; fi23 < 8; fi23++ {
        var v24 int = 32767
        _ = fi23
        _ = v24
    }
    if (!v11) {
        if v11 {
            switch v3 {
            case "\n\t":
                v3 = "a"
            case "foo bar":
                v22 -= (0.0 / 1.5)
            default:
                v12 = v12
            }
            for fi25 := 0; fi25 < 11; fi25++ {
                _ = len(v1)
                switch v12 {
                case -1:
                    switch "" {
                    case "foo bar":
                        if len(v4) > 0 {
                            v4[0] = ("" + (v3 + ""))
                        }
                    case "":
                        v22 -= v22
                    case "\n\t":
                        for fi26 := 0; fi26 < 9; fi26++ {
                            v4 = append(v4, v3)
                            var v27 float32 = 3.14159
                            v11 = ((^(v6)) != 42)
                            _ = fi26
                            _ = v27
                        }
                    }
                case 2147483647:
                    fmt.Println(v3, v22, v6)
                }
                var v28 int16 = 1
                _ = fi25
                _ = v28
            }
        } else {
            if v11 {
                v22 += 0.1
                if (!(1e10 != 0.1)) {
                    if len(v4) > 0 {
                        v4[0] = (v3 + (v3 + ("" + "a")))
                    }
                    for fi29 := 0; fi29 < 2; fi29++ {
                        _ = len(v5)
                        fmt.Println(fi29)
                        _ = fi29
                    }
                    switch (v12 * v12) {
                    case 127:
                        for fi30 := 0; fi30 < 17; fi30++ {
                            v3 = v3
                            v4 = append(v4, ("foo bar" + v3))
                            if (v11 || (42 < 2147483647)) {
                                switch v12 {
                                case -1:
                                    v3 = (("hello" + v3) + "hello")
                                case -2147483648:
                                    v6 = v6
                                case 0:
                                    if ((v12 | fi30) != fi30) {
                                        var v31 int = v12
                                        _ = v31
                                    } else {
                                        var v32 int = 1
                                        _ = v32
                                    }
                                }
                            }
                            _ = fi30
                        }
                    case 32768:
                        for fi33 := 0; fi33 < 4; fi33++ {
                            _ = len(v5)
                            switch -907 {
                            case 566:
                                delete(v2, 42)
                            }
                            for fi34 := 0; fi34 < 12; fi34++ {
                                v6 -= v6
                                var v35 int64 = -1
                                v36 := GenS3{F0: 2, F1: v3, F2: fi33}
                                _ = fi34
                                _ = v35
                                _ = v36
                            }
                            _ = fi33
                        }
                    case -1:
                        var v37 uint32 = 528
                        _ = v37
                    }
                }
            } else {
                v3 = v3
                var v38 int32 = 713
                if v11 {
                    var v39 uint16 = 0
                    v40 := GenS5{F0: false, F1: ""}
                    var v41 uint64 = 1
                    _ = v39
                    _ = v40
                    _ = v41
                }
                _ = v38
            }
            delete(v5, v3)
            fmt.Println(v11, v6, v3)
        }
        var v42 int8 = 42
        v42 = v42
        _ = v42
    } else {
        v43 := GenS6{F0: v12, F1: -982}
        _ = v43
    }
    var v44 uint8 = 2
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v45 float32 = 0.1
        _ = v45
    case 0:
        var v46 uint64 = 588
        _ = v46
    }
    for fi47 := 0; fi47 < 20; fi47++ {
        var v48 uint64 = 263
        var v49 int16 = (v6 % (((v6 / (((v6 << ((255) & 63))) | 1))) | 1))
        fmt.Println(v22, v11, v49)
        _ = fi47
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
