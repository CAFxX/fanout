package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 int8 = v1
    v3 = (^(v1))
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: ((true || true) || ((v3 * v3) != v1)), F1: "hello"}
        if (v1 != v3) {
            var v6 uint = 591
            for fi7 := 0; fi7 < 5; fi7++ {
                if (v3 == v3) {
                    v6 += v6
                }
                var v8 uint64 = 2
                _ = fi7
                _ = v8
            }
            v3 -= (v1 & v1)
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    if (42 >= 255) {
        v9 := GenS2{F0: 2}
        var v10 int16 = 776
        _ = v10
        _ = v9
    } else {
        v11 := GenS2{F0: 4294967295}
        if (3.14159 <= (float32(((-2.25 + (float64(0.1) + (0.0 + (0.0 - 0.1)))) + 1e10)) / (-2.25 * 1e10))) {
            var v12 uint64 = 141
            var v13 uint16 = 255
            if false {
                if true {
                    var v14 int = -473
                    v15 := GenS1{F0: true, F1: (("\n\t" + "a") + ("\n\t" + "hello"))}
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 14; fi16++ {
                    for fi17 := 0; fi17 < 12; fi17++ {
                        v18 := GenS3{F0: 645, F1: v13, F2: v12}
                        v19 := GenS3{F0: 281, F1: v13, F2: (v12 * 1)}
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if (292 <= v12) {
                        if (0 != -32768) {
                            var v20 int = 127
                            v21 := v2["\n\t"]
                            if true {
                                fmt.Println(v1, v1, v20)
                                _ = len(v2)
                            }
                            _ = v20
                            _ = v21
                        }
                    } else {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := []int{(fi16 ^ (fi22 / (((-(fi22))) | 1))), (fi22 + (fi16 / (((fi16 >> ((42) & 63))) | 1))), (fi16 * (fi16 / ((fi22) | 1)))}
                            _ = fi22
                            _ = v23
                        }
                        v12 -= 1
                    }
                    _ = fi16
                }
            } else {
                v24 := GenS4{F0: -439, F1: ((0 >= -32768) && true)}
                for fi25 := 0; fi25 < 12; fi25++ {
                    var v26 uint = 255
                    _ = fi25
                    _ = v26
                }
                if true {
                    var v27 string = ("a" + ("a" + "a"))
                    _ = v27
                } else {
                    switch 127 {
                    case -9223372036854775808:
                        v12 = (^(v12))
                    case 822:
                        v3 = -67
                    default:
                        var v28 int64 = -1
                        _ = v28
                    }
                    switch 4294967295 {
                    case 42:
                        v1 += v1
                    default:
                        switch "a" {
                        case "\n\t":
                            v29 := GenS2{F0: v12}
                            _ = v29
                        default:
                            fmt.Println(v13)
                        }
                    }
                    for fi30 := 0; fi30 < 13; fi30++ {
                        v2["foo bar"] = 2147483648
                        v31 := GenS3{F0: 408, F1: v13, F2: v12}
                        _ = len(v2)
                        _ = fi30
                        _ = v31
                    }
                }
                _ = v24
            }
            _ = v12
            _ = v13
        }
        for fi32 := 0; fi32 < 1; fi32++ {
            for fi33 := 0; fi33 < 3; fi33++ {
                if true {
                    var v34 uint64 = 1
                    _ = v34
                } else {
                    var v35 uint64 = 4294967295
                    var v36 float64 = 0.0
                    _ = v35
                    _ = v36
                }
                v37 := GenS1{F0: false, F1: (("foo bar" + "a") + "")}
                _ = fi33
                _ = v37
            }
            _ = fi32
        }
        _ = v11
    }
    v38 := v2[("a" + "foo bar")]
    v39 := []int{-84}
    var v40 uint16 = 42
    v41 := GenS5{F0: 226}
    v38 = (v38 % ((2147483647) | 1))
    v39 = append(v39, 4294967295)
    var v42 uint8 = 128
    switch 21 {
    case -146:
        var v43 bool = (true && (-464 > 834))
        _ = v43
    }
    var v44 uint8 = (v42 * (v42 ^ (v42 % ((v42) | 1))))
    var v45 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
    _ = v44
    _ = v45
}
