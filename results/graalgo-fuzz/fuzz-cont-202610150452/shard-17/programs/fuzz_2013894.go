package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 int }
type GenS4 struct { F0 int; F1 uint16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool }
func genFn0(p0 uint, p1 int64, p2 int16) bool {
    for fi6 := 0; fi6 < 9; fi6++ {
        p0 = p0
        _ = fi6
    }
    p1 = (-(p1))
    _ = p0
    _ = p1
    _ = p2
    return ((p1 >> (((p0 - (^(p0)))) & 63)) < 9223372036854775807)
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 557
    var v3 uint64 = 0
    v4 := []string{"foo bar", "\n\t"}
    var v5 float64 = (((1e10 / -2.25) + 3.14159) * (1.5 + -2.25))
    v5 = v5
    var v7 int = int(v3)
    v8 := GenS1{F0: v3, F1: (v2 / ((2) | 1)), F2: v3}
    var v9 bool = (v2 != v2)
    var v10 int = v7
    fmt.Println(v7, v7)
    v10 = (v7 & (v10 * v7))
    fmt.Println(v10, v2, v9)
    fmt.Println(v7)
    var v11 int32 = 838
    var v12 float32 = ((-2.25 + 3.14159) + 0.1)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: 255}
        var v15 int16 = 0
        for fi16 := 0; fi16 < 4; fi16++ {
            switch "" {
            case "\n\t":
                switch 2147483647 {
                case 1:
                    var v17 uint8 = 72
                    _ = v17
                case 2147483647:
                    fmt.Println(fi16, v2, fi13)
                case 9223372036854775807:
                    var v18 float32 = (1e10 - 0.0)
                    _ = v18
                }
            case "":
                fmt.Println(v10)
            case "foo bar":
                if (1 < 1) {
                    v5 = 0.1
                }
            }
            v19 := GenS3{F0: v2, F1: (v3 >> ((v2) & 63)), F2: 517}
            _ = len(v1)
            _ = fi16
            _ = v19
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    if (142 > 1) {
        v20 := GenS4{F0: (v7 - v7), F1: 83}
        v7 += -9223372036854775808
        for fi21 := 0; fi21 < 11; fi21++ {
            for fi22 := 0; fi22 < 18; fi22++ {
                switch (("" + ("foo bar" + ("\n\t" + ""))) + "hello") {
                case "a":
                    if (((float32(1e10) - (((v12 * 3.14159) / v12) - float32(v5))) > v12) || false) {
                        switch v7 {
                        case -979:
                            if (255 == 2) {
                                v9 = v9
                                if false {
                                    _ = len(v4)
                                    switch "foo bar" {
                                    case "\n\t":
                                        switch (-(v10)) {
                                        case 4294967295:
                                            v23 := GenS3{F0: (v2 >> ((0) & 63)), F1: v3, F2: v7}
                                            _ = v23
                                        }
                                    case "hello":
                                        var v24 int8 = -1
                                        _ = v24
                                    }
                                    if len(v4) > 0 {
                                        v4[0] = ("hello" + (("\n\t" + "hello") + "foo bar"))
                                    }
                                } else {
                                    fmt.Println(v3)
                                }
                                v5 += 1e10
                            } else {
                                if false {
                                    v9 = (1 >= 42)
                                    v25 := GenS3{F0: 18446744073709551615, F1: v3, F2: (v10 % ((-760) | 1))}
                                    var v26 uint8 = 0
                                    if v9 {
                                        if (v9 && (0 == 0)) {
                                            v1[""] = int(v2)
                                            v12 -= ((v12 / v12) + float32(3.14159))
                                            v27 := GenS5{F0: 1}
                                            _ = v27
                                        }
                                        if v9 {
                                            fmt.Println(v2, v12)
                                            v5 -= float64(v12)
                                        }
                                        var v28 uint8 = 2
                                        _ = v28
                                    }
                                    switch "a" {
                                    case "hello":
                                        v29 := GenS3{F0: v2, F1: (^(v3)), F2: (v10 & 4294967295)}
                                        _ = v29
                                    default:
                                        delete(v1, (("foo bar" + "") + (("" + "hello") + "\n\t")))
                                    }
                                    _ = v25
                                    _ = v26
                                }
                                fmt.Println(v2, fi22)
                            }
                        default:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    } else {
                        switch (v7 + fi22) {
                        case 32768:
                            if (32767 >= 65535) {
                                switch 2147483647 {
                                case 0:
                                    var v30 uint32 = 668
                                    _ = v30
                                case 127:
                                    var v31 int16 = 1000
                                    _ = v31
                                case 128:
                                    v12 -= v12
                                }
                                v2 -= v2
                            }
                        case 128:
                            v2 += v2
                        }
                    }
                case "hello":
                    var v32 int = int(v11)
                    _ = v32
                case "\n\t":
                    var v33 bool = v9
                    _ = v33
                }
                var v34 float64 = (v5 * v5)
                _ = fi22
                _ = v34
            }
            v10 = fi21
            for fi35 := 0; fi35 < 7; fi35++ {
                if len(v4) > 0 {
                    v4[0] = ("" + "hello")
                }
                _ = fi35
            }
            _ = fi21
        }
        _ = v20
    }
    v2 += v2
    v36 := GenS6{F0: ((-74 <= 42) && (((v2 ^ (^(v2))) != (v2 & v2)) || v9))}
    for fi37 := 0; fi37 < 18; fi37++ {
        var v38 string = "hello"
        _ = fi37
        _ = v38
    }
    v39 := GenS3{F0: (v2 | (v2 + v2)), F1: 922, F2: v7}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
