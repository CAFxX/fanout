package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 uint) int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 - fi6), (fi6 << ((4294967295) & 63)), fi6, (fi6 / ((int(p0)) | 1))}
        _ = fi6
        _ = v7
    }
    v8 := []int{42, -159}
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v8
    return (p0 | (-(p0)))
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []string{"foo bar", "hello", "hello"}
    var v3 int = -2147483648
    v4 := make(map[string]string)
    var v5 uint8 = 128
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int64 = 629
        var v11 uint8 = v5
        _ = fi9
        _ = v10
        _ = v11
    }
    if false {
        if (!(false || (409 == 4294967295))) {
            if true {
                var v12 float32 = -2.25
                if false {
                    v13 := GenS1{F0: 355, F1: 255}
                    _ = v13
                } else {
                    v5 += (v5 >> ((42) & 63))
                    var v14 uint8 = v5
                    _ = v14
                }
                v5 = v5
                _ = v12
            }
        }
        var v15 float32 = -2.25
        _ = v15
    } else {
        var v16 bool = ((0 <= 18446744073709551615) || false)
        if v16 {
            delete(v4, ("\n\t" + ("foo bar" + ("" + "hello"))))
        }
        _ = v16
    }
    switch v3 {
    case -2147483648:
        for fi17 := 0; fi17 < 18; fi17++ {
            fmt.Println(v3, v5)
            if (!false) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    if false {
                        if true {
                            fmt.Println(fi18, v3)
                            var v19 uint16 = 805
                            _ = v19
                        }
                        fmt.Println(fi17, fi18)
                    }
                    v5 += v5
                    if false {
                        if (false || (-1 >= 32767)) {
                            switch "\n\t" {
                            case "\n\t":
                                var v20 float64 = -2.25
                                _ = v20
                            case "foo bar":
                                v21 := GenS1{F0: 255, F1: 32767}
                                _ = v21
                            default:
                                v22 := v4["hello"]
                                _ = v22
                            }
                            var v23 int64 = -9223372036854775808
                            var v24 float64 = 1.5
                            _ = v23
                            _ = v24
                        } else {
                            v5 = (v5 ^ v5)
                            v3 = v3
                        }
                        var v25 float32 = 1e10
                        var v26 uint64 = 693
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS2{F0: (255 == 42)}
                        v3 -= (-(fi18))
                        _ = v27
                    }
                    _ = fi18
                }
                v5 = v5
            }
            _ = fi17
        }
    case -400:
        v3 = 14
    case -519:
        v28 := GenS3{F0: 48, F1: -243, F2: 2147483648}
        _ = v28
    }
    var v29 int16 = -1
    if ((!true) || false) {
        for fi30 := 0; fi30 < 19; fi30++ {
            switch "hello" {
            case "hello":
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS1{F0: 42, F1: 424}
                    _ = fi31
                    _ = v32
                }
            default:
                v29 -= -32768
            }
            v5 += (v5 % ((2) | 1))
            _ = fi30
        }
        fmt.Println(v3)
    }
    v3 = v3
    var v33 uint16 = 255
    fmt.Println(v33)
    var v34 float32 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
