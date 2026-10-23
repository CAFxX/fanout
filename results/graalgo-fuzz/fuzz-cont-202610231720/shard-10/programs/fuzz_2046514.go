package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int }
type GenS2 struct { F0 int16; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string; F2 int32 }
type GenS5 struct { F0 string; F1 bool; F2 bool }
type GenS6 struct { F0 uint; F1 uint }
func genFn0(p0 int16, p1 string, p2 uint64) int64 {
    var v7 uint32 = 65535
    v8 := []int{}
    switch 152 {
    case 128:
        for fi9 := 0; fi9 < 12; fi9++ {
            if true {
                if true {
                    if ((v7 & (v7 | 278)) >= v7) {
                        v10 := GenS1{F0: 2, F1: p0, F2: fi9}
                        _ = v10
                    } else {
                        var v11 uint64 = (p2 ^ p2)
                        p1 = p1
                        _ = v11
                    }
                    if ((v7 ^ v7) > v7) {
                        if true {
                            if true {
                                if len(v8) > 0 {
                                    v8[0] = fi9
                                }
                                if false {
                                    p0 = p0
                                }
                                var v12 float64 = (0.0 / (-2.25 / 1e10))
                                _ = v12
                            }
                            var v13 uint32 = v7
                            switch (fi9 - (fi9 * (fi9 << ((65535) & 63)))) {
                            case -592:
                                switch p1 {
                                case "hello":
                                    if len(v8) > 0 {
                                        v8[0] = (fi9 >> ((65535) & 63))
                                    }
                                case "":
                                    if true {
                                        for fi14 := 0; fi14 < 11; fi14++ {
                                            v13 = 65535
                                            v15 := GenS2{F0: p0, F1: 2147483647, F2: 576}
                                            _ = fi14
                                            _ = v15
                                        }
                                        var v16 uint = 18446744073709551615
                                        var v17 float64 = 1.5
                                        _ = v16
                                        _ = v17
                                    }
                                default:
                                    var v18 int64 = -1
                                    _ = v18
                                }
                            case -2147483648:
                                var v19 uint8 = 128
                                _ = v19
                            }
                            _ = v13
                        } else {
                            var v20 int32 = 886
                            var v21 uint = 65535
                            fmt.Println(p0, v7)
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS2{F0: p0, F1: -1, F2: 26}
                        _ = v22
                    } else {
                        v8 = append(v8, -110)
                        p0 += (-(p0))
                    }
                }
                v7 = v7
                p1 = p1
            } else {
                v23 := GenS1{F0: 42, F1: 662, F2: (fi9 & fi9)}
                p1 = p1
                _ = v23
            }
            p2 = p2
            for fi24 := 0; fi24 < 2; fi24++ {
                switch ("a" + p1) {
                case "":
                    var v25 int = -540
                    _ = v25
                case "hello":
                    var v26 int16 = p0
                    if false {
                        var v27 bool = false
                        switch fi9 {
                        case -2147483648:
                            var v28 string = p1
                            _ = v28
                        }
                        _ = v27
                    }
                    _ = v26
                }
                v29 := GenS1{F0: 285, F1: (p0 % ((0) | 1)), F2: (^(fi9))}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi24
                _ = v29
            }
            _ = fi9
        }
    }
    var v30 int32 = 24
    var v31 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    _ = v31
    _ = v7
    _ = v8
    return -983
}

func main() {
    var v1 uint16 = 2
    var v2 float64 = 3.14159
    var v3 uint8 = 42
    var v4 float64 = v2
    v5 := []string{"\n\t", ("a" + ("a" + "a"))}
    var v6 uint = 537
    for fi32 := 0; fi32 < 7; fi32++ {
        v4 += 0.1
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi32
    }
    var v33 string = ("a" + ("a" + "foo bar"))
    for fi34 := 0; fi34 < 6; fi34++ {
        var v35 float64 = v2
        v36 := GenS1{F0: 65535, F1: 0, F2: fi34}
        fmt.Println(v35, v3, v4)
        _ = fi34
        _ = v35
        _ = v36
    }
    v37 := GenS3{F0: (v3 / (((^(v3))) | 1))}
    v38 := GenS4{F0: (v1 + v1), F1: ((v33 + ((v33 + v33) + (v33 + (((v33 + (((v33 + v33) + v33) + v33)) + (v33 + v33)) + "hello")))) + v33), F2: 1}
    fmt.Println(v4, v2, v4)
    fmt.Println(v2, v3, v2)
    v39 := GenS5{F0: v33, F1: false, F2: true}
    var v40 uint64 = 2
    _ = len(v5)
    if (116 <= 57) {
        fmt.Println(v3)
        v41 := GenS6{F0: (v6 - (v6 >> ((0) & 63))), F1: 0}
        var v42 uint32 = 82
        _ = v41
        _ = v42
    }
    v4 = ((0.0 + (((v4 / (v4 * ((((v4 / 1.0) - 0.0) - float64((-2.25 / (float32((v2 / -2.25)) * 0.0)))) / (float64(0.1) / 1.5)))) * (v2 * v4)) + v4)) - 3.14159)
    var v43 float64 = v2
    for fi44 := 0; fi44 < 1; fi44++ {
        var v45 int64 = 1
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v43
    _ = v5
    _ = v6
}
