package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int }
type GenS3 struct { F0 string; F1 int64; F2 int64 }

func main() {
    var v1 int64 = 1
    v2 := make(map[int]bool)
    var v3 float64 = 0.0
    var v4 uint8 = 32
    var v5 uint16 = 255
    v2[-1] = ((!true) || (0.0 != (3.14159 / ((((v3 * v3) - 0.1) / 1.5) / 3.14159))))
    if (uint32(v5) != 367) {
        var v6 int64 = 9223372036854775807
        _ = v6
    } else {
        v3 = float64(((float32((3.14159 - (v3 / v3))) + 3.14159) / 1e10))
        v7 := []int{4294967295, int(v5), 239}
        fmt.Println(v5, v3)
        _ = v7
    }
    var v8 int64 = v1
    v8 = (v8 << ((0) & 63))
    for fi9 := 0; fi9 < 8; fi9++ {
        v5 = v5
        fmt.Println(fi9, fi9, v5)
        _ = fi9
    }
    v8 = (v1 & int64(v5))
    var v10 int16 = 951
    v11 := GenS1{F0: 650}
    if (0 > 28) {
        for fi12 := 0; fi12 < 18; fi12++ {
            if true {
                switch (^(fi12)) {
                case 127:
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v14 := GenS2{F0: "hello", F1: (fi12 & (fi12 * fi13))}
                        v15 := GenS2{F0: "foo bar", F1: (^(fi13))}
                        v3 += v3
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                case 9223372036854775807:
                    if false {
                        var v16 uint32 = 841
                        if ("" < (("hello" + ("a" + (("a" + "a") + "a"))) + ((("a" + "\n\t") + ("hello" + "a")) + ""))) {
                            var v17 int16 = v10
                            _ = v17
                        }
                        _ = v16
                    } else {
                        v8 = v1
                        var v18 bool = false
                        _ = v18
                    }
                }
                var v19 uint16 = v5
                switch "foo bar" {
                case "foo bar":
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 int = (fi20 & fi20)
                        v8 = (-(v1))
                        fmt.Println(fi12)
                        _ = fi20
                        _ = v21
                    }
                case "a":
                    for fi22 := 0; fi22 < 14; fi22++ {
                        v4 += (v4 & (v4 / (((v4 & v4)) | 1)))
                        if true {
                            v3 = (0.0 / float64(3.14159))
                            v4 += 128
                            switch (fi12 % ((fi12) | 1)) {
                            case 32767:
                                var v23 int8 = -66
                                _ = v23
                            case -760:
                                if (true && (-858 >= 0)) {
                                    v24 := GenS2{F0: "a", F1: int(v1)}
                                    var v25 bool = ((v4 + v4) < 212)
                                    _ = v24
                                    _ = v25
                                }
                                v26 := []int{55, int(v5), fi22, fi12}
                                _ = v26
                            case -1:
                                if (!false) {
                                    _ = len(v2)
                                }
                            default:
                                v19 += (^(v5))
                            }
                        }
                        if true {
                            _ = len(v2)
                        }
                        _ = fi22
                    }
                }
                _ = v19
            } else {
                if (255 <= 434) {
                    v3 = v3
                    if true {
                        v3 = v3
                    } else {
                        var v27 bool = (688 == 42)
                        v28 := GenS1{F0: 517}
                        _ = v27
                        _ = v28
                    }
                }
            }
            var v29 uint16 = 32767
            _ = fi12
            _ = v29
        }
    }
    v4 += (v4 % ((v4) | 1))
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := []int{(fi30 >> ((875) & 63))}
        for fi32 := 0; fi32 < 8; fi32++ {
            switch "hello" {
            case "\n\t":
                v10 = v10
            case "":
                var v33 float64 = (-2.25 * (float64(0.0) + ((float64((-2.25 * 0.0)) + 0.0) - float64(0.1))))
                _ = v33
            default:
                v3 = (v3 + float64(((-2.25 * float32(float64(-2.25))) / 3.14159)))
            }
            v10 = v10
            v8 = int64(v10)
            _ = fi32
        }
        _ = fi30
        _ = v31
    }
    v34 := []int{128, -2147483648, 32768}
    switch "foo bar" {
    case "":
        v35 := GenS3{F0: "a", F1: (v1 % ((v1) | 1)), F2: v8}
        _ = v35
    case "hello":
        var v36 int32 = 0
        _ = v36
    default:
        for fi37 := 0; fi37 < 17; fi37++ {
            v10 += 707
            if ((!true) || (!((v8 | (^(v8))) != (v1 - v8)))) {
                var v38 int8 = 45
                _ = v38
            } else {
                var v39 float64 = 0.0
                delete(v2, fi37)
                _ = v39
            }
            _ = fi37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
}
