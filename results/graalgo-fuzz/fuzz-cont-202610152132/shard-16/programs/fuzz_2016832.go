package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0() int32 {
    v6 := []int{4294967295}
    var v7 int16 = -722
    v6 = append(v6, -2147483648)
    _ = v6
    _ = v7
    return -212
}

func main() {
    var v1 uint8 = 127
    v2 := []int{0, 4294967295}
    v3 := make(map[string]bool)
    v4 := make(map[uint32]string)
    v5 := make(map[int]string)
    var v8 int64 = 0
    if len(v2) > 0 {
        v2[0] = 32767
    }
    if (!false) {
        v3["\n\t"] = (9223372036854775807 == 1)
        if false {
            v9 := v3["a"]
            var v10 int32 = 1
            var v11 int32 = v10
            _ = v10
            _ = v11
            _ = v9
        } else {
            var v12 string = "\n\t"
            switch (("foo bar" + "foo bar") + "") {
            case "":
                v13 := v3["foo bar"]
                _ = v13
            case "a":
                switch v12 {
                case "":
                    for fi14 := 0; fi14 < 17; fi14++ {
                        var v15 bool = true
                        _ = fi14
                        _ = v15
                    }
                case "\n\t":
                    switch v12 {
                    case "foo bar":
                        fmt.Println(v12, v12)
                    case "\n\t":
                        var v16 int64 = (^(v8))
                        _ = v16
                    }
                case "foo bar":
                    v17 := v3[v12]
                    _ = v17
                default:
                    v12 = (v12 + ("" + v12))
                }
            case "foo bar":
                v1 = v1
            }
            _ = v12
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            v1 -= (v1 % ((2) | 1))
            v1 = 62
            _ = fi18
        }
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            for fi20 := 0; fi20 < 12; fi20++ {
                if (false && (224 <= -32768)) {
                    var v21 float32 = (0.1 * 1e10)
                    v8 = (-(v8))
                    _ = v21
                } else {
                    if false {
                        var v22 uint = 1
                        var v23 int = 607
                        if ((!(812 == v22)) || true) {
                            v1 = v1
                            v8 -= (-(v8))
                            var v24 uint32 = 627
                            _ = v24
                        }
                        v23 -= (^(fi19))
                        var v25 int32 = -206
                        _ = v22
                        _ = v23
                        _ = v25
                    }
                    v1 -= (v1 >> ((42) & 63))
                    v1 -= 197
                }
                v26 := GenS1{F0: 2147483647, F1: v8, F2: 194}
                v1 = (v1 | v1)
                fmt.Println(fi20, v8)
                _ = fi20
                _ = v26
            }
            switch (fi19 << ((950) & 63)) {
            case 42:
                var v27 uint32 = 2147483647
                _ = v27
            case -2147483648:
                var v28 uint32 = 0
                _ = v28
            case 1:
                var v29 float64 = 0.1
                _ = v29
            default:
                v30 := GenS1{F0: 0, F1: 42, F2: v1}
                _ = v30
            }
            _ = fi19
        }
    }
    var v31 float64 = float64(3.14159)
    v32 := GenS2{F0: 1, F1: 18446744073709551615}
    if (42 != 0) {
        v8 += -1
        var v33 uint8 = 127
        _ = v33
    }
    var v34 int16 = -1
    v35 := v5[-206]
    if (!true) {
        var v36 uint64 = 4294967295
        fmt.Println(v34, v36)
        var v37 int64 = (v8 ^ (v8 - 690))
        _ = v36
        _ = v37
    }
    v8 += (v8 - v8)
    for fi38 := 0; fi38 < 1; fi38++ {
        if true {
            delete(v5, (-(fi38)))
            fmt.Println(v34)
            if (true && ((!((v34 / ((v34) | 1)) <= v34)) || false)) {
                v8 = 0
                if true {
                    if ((false || true) || (54 >= 65535)) {
                        _ = len(v2)
                    } else {
                        if (!false) {
                            v8 = (v8 % ((v8) | 1))
                            var v39 int = (fi38 % (((fi38 - fi38)) | 1))
                            _ = v39
                        } else {
                            if false {
                                _ = len(v3)
                            }
                        }
                        var v40 uint = 42
                        _ = v40
                    }
                }
            }
        }
        var v41 uint = 623
        _ = fi38
        _ = v41
    }
    var v42 string = (v35 + v35)
    var v43 int64 = v8
    v43 = 42
    for fi44 := 0; fi44 < 14; fi44++ {
        v1 = (^(v1))
        _ = fi44
    }
    v1 += (v1 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
