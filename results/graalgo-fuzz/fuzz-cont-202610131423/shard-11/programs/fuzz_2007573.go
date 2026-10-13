package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 uint64; F1 uint64; F2 string }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0() int32 {
    var v7 int16 = 1000
    v8 := GenS1{F0: 503}
    fmt.Println(v7)
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int = 428
    v2 := []int{v1, (v1 * (v1 << ((2) & 63))), (v1 + (v1 >> ((952) & 63)))}
    var v3 uint8 = 170
    v4 := []uint32{}
    v5 := make(map[string]string)
    v6 := make(map[int]bool)
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint32 = 65535
        _ = fi9
        _ = v10
    }
    var v11 string = ("\n\t" + "\n\t")
    var v12 uint = 483
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 bool = (!(-1 == 2147483647))
        switch v11 {
        case "hello":
            for fi15 := 0; fi15 < 7; fi15++ {
                if (4294967295 >= 2) {
                    switch 351 {
                    case 0:
                        if len(v2) > 0 {
                            v2[0] = (fi13 % (((fi15 >> (((v12 - v12)) & 63))) | 1))
                        }
                    case 128:
                        v3 = v3
                    }
                    switch v11 {
                    case "foo bar":
                        var v16 string = v11
                        var v17 int = 9223372036854775807
                        v18 := GenS1{F0: 119}
                        v12 = (v12 % (((v12 % (((v12 % (((v12 | 702)) | 1))) | 1))) | 1))
                        v14 = v14
                        var v19 uint = v12
                        if v14 {
                            fmt.Println(v12, v19, v19)
                        }
                        _ = v16
                        _ = v17
                        _ = v18
                        _ = v19
                    case "a":
                        var v20 int8 = -39
                        _ = v20
                    }
                    var v21 int32 = -2147483648
                    _ = v21
                } else {
                    if (!(false && (float64(((1.5 - (1.5 / (float32(0.0) - 1e10))) * 0.0)) < (((1.5 / 1.5) + float64((1e10 / ((0.1 + -2.25) - 3.14159)))) + 3.14159)))) {
                        var v22 string = v11
                        v23 := GenS2{F0: -65, F1: v1}
                        _ = v22
                        _ = v23
                    }
                    v3 = v3
                }
                v2 = append(v2, (fi13 ^ 32768))
                v4 = append(v4, 1)
                _ = fi15
            }
        case "a":
            for fi24 := 0; fi24 < 1; fi24++ {
                v25 := GenS2{F0: 0, F1: fi24}
                var v26 int8 = 127
                if (783 <= 2147483647) {
                    switch (v1 * 32768) {
                    case -112:
                        switch "a" {
                        case "\n\t":
                            var v27 uint64 = 4294967295
                            _ = v27
                        case "":
                            switch (v11 + (("" + v11) + "a")) {
                            case "hello":
                                var v28 int64 = 42
                                _ = v28
                            }
                        default:
                            if v14 {
                                var v29 uint8 = (v3 ^ v3)
                                v30 := GenS3{F0: ("\n\t" + v11), F1: ("" + "foo bar")}
                                _ = v29
                                _ = v30
                            }
                        }
                    case 525:
                        v14 = (false || ((255 >= 65535) && (v11 == v11)))
                    }
                    if ((1.5 + 3.14159) <= 1.5) {
                        v31 := GenS4{F0: 658, F1: 2, F2: "hello"}
                        v5[(v11 + v11)] = (v11 + ("a" + v11))
                        switch v1 {
                        case 59:
                            var v32 int16 = 401
                            _ = v32
                        case -2147483648:
                            delete(v5, v11)
                        case -655:
                            if v14 {
                                var v33 string = v11
                                v34 := GenS3{F0: v33, F1: (v33 + v33)}
                                v33 = "hello"
                                _ = v33
                                _ = v34
                            }
                        default:
                            var v35 string = ("" + "")
                            _ = v35
                        }
                        _ = v31
                    } else {
                        var v36 float32 = 3.14159
                        fmt.Println(v11, v12, fi24)
                        v37 := GenS5{F0: 42, F1: -898}
                        if (-9223372036854775808 >= -1) {
                            _ = len(v4)
                            _ = len(v5)
                        }
                        _ = v36
                        _ = v37
                    }
                    switch "foo bar" {
                    case "foo bar":
                        switch 1 {
                        case -1:
                            v4 = append(v4, 2)
                        case -9223372036854775808:
                            v38 := GenS5{F0: 0, F1: -1}
                            _ = v38
                        default:
                            fmt.Println(v1, fi13)
                        }
                    case "":
                        fmt.Println(v1, v3, fi24)
                    case "a":
                        switch 42 {
                        case -9223372036854775808:
                            v1 += v1
                        }
                    default:
                        switch v11 {
                        case "a":
                            var v39 uint8 = v3
                            _ = v39
                        case "":
                            v3 = uint8(v12)
                        case "foo bar":
                            v1 = (-(fi13))
                        }
                    }
                    var v40 int8 = 0
                    v1 -= (fi13 >> ((1) & 63))
                    fmt.Println(fi13)
                    v5[v11] = "\n\t"
                    _ = v40
                }
                _ = fi24
                _ = v25
                _ = v26
            }
        default:
            fmt.Println(v1)
        }
        v12 = (v12 << ((v12) & 63))
        _ = fi13
        _ = v14
    }
    v1 = (v1 - (^(v1)))
    v4 = append(v4, 4294967295)
    v12 = (v12 + v12)
    delete(v5, (v11 + (("a" + (v11 + "a")) + (("" + ("" + v11)) + ""))))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
