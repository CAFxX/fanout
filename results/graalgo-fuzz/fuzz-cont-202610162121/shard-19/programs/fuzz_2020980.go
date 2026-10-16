package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 string }
func genFn0(p0 int16) int64 {
    for fi4 := 0; fi4 < 15; fi4++ {
        var v5 int8 = 127
        var v6 int = fi4
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 1
    if (128 != 1) {
        fmt.Println(p0)
        v8 := GenS1{F0: 420}
        _ = v8
    }
    if (true && ((true || false) && true)) {
        v9 := GenS1{F0: 2147483648}
        _ = v9
    }
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint, p1 string, p2 int16) int32 {
    var v10 float32 = 0.1
    var v11 int16 = 956
    v12 := []int{32768, 16}
    fmt.Println(p1, v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]string)
    v3 := make(map[string]string)
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int32 = genFn1(65535, (("a" + "hello") + ("foo bar" + "")), 32767)
        if false {
            if (0.1 <= float64((1.5 / 1.0))) {
                var v15 uint32 = 0
                _ = v15
            } else {
                if true {
                    var v16 int8 = 0
                    var v17 int16 = -1
                    switch "\n\t" {
                    case "\n\t":
                        fmt.Println(fi13, v17)
                    case "foo bar":
                        if (((v17 << ((255) & 63)) <= 190) && true) {
                            for fi18 := 0; fi18 < 11; fi18++ {
                                v19 := GenS2{F0: int64(fi13), F1: 42}
                                _ = fi18
                                _ = v19
                            }
                            v20 := []int{(fi13 | (fi13 / (((^(fi13))) | 1)))}
                            _ = v20
                        } else {
                            switch "foo bar" {
                            case "a":
                                v21 := GenS3{F0: -32768, F1: (v17 | v17), F2: ""}
                                _ = v21
                            case "\n\t":
                                v22 := []int{(fi13 / ((-2147483648) | 1)), fi13, -2147483648}
                                _ = v22
                            default:
                                fmt.Println(fi13, v16)
                            }
                        }
                    case "":
                        v16 = (v16 + 1)
                    }
                    _ = v16
                    _ = v17
                } else {
                    switch fi13 {
                    case 9223372036854775807:
                        v14 = -947
                    default:
                        if true {
                            var v23 int16 = 32767
                            var v24 int32 = (v14 % ((v14) | 1))
                            var v25 float64 = (1e10 / -2.25)
                            _ = v23
                            _ = v24
                            _ = v25
                        } else {
                            v14 += v14
                            if (!true) {
                                fmt.Println(fi13)
                                v26 := GenS2{F0: -1, F1: 255}
                                _ = v26
                            }
                        }
                    }
                }
                switch "hello" {
                case "hello":
                    v14 -= v14
                }
            }
            var v27 uint32 = 75
            _ = v27
        } else {
            switch "\n\t" {
            case "":
                var v28 int32 = (v14 << ((65535) & 63))
                _ = v28
            case "a":
                var v29 uint = 1
                _ = v29
            case "hello":
                fmt.Println(fi13, fi13)
            }
            var v30 uint64 = 719
            v30 += (v30 << ((583) & 63))
            _ = v30
        }
        _ = fi13
        _ = v14
    }
    var v31 int = 2147483647
    v32 := v1[("foo bar" + ("" + (("a" + ("\n\t" + ("foo bar" + ("" + ("" + "\n\t"))))) + "a")))]
    v33 := v3[("hello" + "foo bar")]
    v32 += (^(v32))
    var v34 int = (v31 / (((v32 | v31)) | 1))
    if (42 == 42) {
        for fi35 := 0; fi35 < 2; fi35++ {
            switch fi35 {
            case 4294967295:
                v33 = v33
            case 127:
                v33 = "a"
            case 42:
                fmt.Println(v31)
            }
            switch fi35 {
            case -181:
                v31 = (v34 << ((18446744073709551615) & 63))
            }
            _ = fi35
        }
    } else {
        v33 = "\n\t"
        v36 := []int{9223372036854775807, (v32 / (((^(v31))) | 1)), v32}
        _ = v36
    }
    v37 := []int{v31, (-(v34)), 42}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v37
}
