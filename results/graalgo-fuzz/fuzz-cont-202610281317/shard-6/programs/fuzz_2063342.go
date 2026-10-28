package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint32, p1 uint64, p2 uint32) int32 {
    if (true || false) {
        v7 := GenS1{F0: -1, F1: 1}
        v8 := GenS1{F0: -896, F1: uint16(p1)}
        var v9 uint = 749
        _ = v7
        _ = v8
        _ = v9
    } else {
        if true {
            var v10 uint32 = 994
            var v11 string = ""
            var v12 uint16 = 42
            _ = v10
            _ = v11
            _ = v12
        } else {
            var v13 float64 = ((0.0 * (float64(-2.25) + float64(0.0))) * 0.1)
            _ = v13
        }
        var v14 float64 = 0.0
        _ = v14
    }
    var v15 uint32 = (p2 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return -601
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[string]uint32)
    var v3 uint = 65535
    v4 := make(map[string]int)
    var v5 int16 = 42
    var v6 uint32 = 655
    if (2 >= (v3 % (((v3 & v3)) | 1))) {
        fmt.Println(v6, v5, v6)
        var v16 uint64 = 0
        _ = v16
    }
    v17 := GenS1{F0: 273, F1: 32767}
    v3 -= (v3 * v3)
    fmt.Println(v6)
    var v18 int = 695
    for fi19 := 0; fi19 < 8; fi19++ {
        fmt.Println(v6, v6)
        if false {
            for fi20 := 0; fi20 < 14; fi20++ {
                if true {
                    var v21 int8 = -1
                    if ("" == ("foo bar" + "a")) {
                        fmt.Println(v6)
                        v22 := []int{(fi19 - fi20), v18, 32767, fi20}
                        _ = v22
                    } else {
                        v23 := GenS2{F0: 2, F1: 1}
                        fmt.Println(v3, fi19)
                        _ = v23
                    }
                    _ = v21
                }
                var v24 uint = 91
                _ = fi20
                _ = v24
            }
        } else {
            v25 := []int{v18}
            v6 = v6
            var v26 int64 = 261
            _ = v25
            _ = v26
        }
        _ = fi19
    }
    var v27 int8 = 127
    var v28 int8 = 1
    v27 -= (v28 - (-(v28)))
    fmt.Println(v3, v18, v3)
    v29 := []int{v18}
    if (!(0 <= v27)) {
        v30 := GenS1{F0: -497, F1: 1}
        v29 = append(v29, 371)
        _ = v30
    }
    v18 -= v18
    if true {
        var v31 uint32 = v6
        _ = v31
    }
    fmt.Println(v5)
    if true {
        v32 := GenS1{F0: -818, F1: 194}
        for fi33 := 0; fi33 < 12; fi33++ {
            if false {
                if false {
                    if false {
                        v34 := GenS3{F0: -9223372036854775808}
                        for fi35 := 0; fi35 < 12; fi35++ {
                            v27 += (v27 & (v28 ^ (^(v27))))
                            _ = fi35
                        }
                        _ = v34
                    }
                } else {
                    if len(v29) > 0 {
                        v29[0] = (v18 * (fi33 + v18))
                    }
                    var v36 string = ""
                    _ = v36
                }
            }
            fmt.Println(v28)
            var v37 float64 = -2.25
            _ = fi33
            _ = v37
        }
        switch "" {
        case "a":
            fmt.Println(v27)
        case "":
            v5 = v5
        case "foo bar":
            for fi38 := 0; fi38 < 14; fi38++ {
                v3 += (v3 / (((v3 << ((4294967295) & 63))) | 1))
                v39 := v1["foo bar"]
                switch v39 {
                case "hello":
                    for fi40 := 0; fi40 < 6; fi40++ {
                        fmt.Println(v6, v5)
                        _ = fi40
                    }
                case "a":
                    var v41 float64 = -2.25
                    _ = v41
                case "foo bar":
                    _ = len(v29)
                default:
                    fmt.Println(v39)
                }
                _ = fi38
                _ = v39
            }
        default:
            fmt.Println(v3, v3)
        }
        _ = v32
    }
    _ = len(v29)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
