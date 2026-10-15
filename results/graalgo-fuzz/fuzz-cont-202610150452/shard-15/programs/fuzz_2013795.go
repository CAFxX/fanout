package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 int32 }
func genFn0() int64 {
    var v5 uint32 = 993
    if (523 == 42) {
        var v6 uint8 = 128
        _ = v6
    }
    v5 += v5
    v5 -= v5
    fmt.Println(v5)
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    var v3 uint = 4294967295
    v4 := []int{732}
    switch "" {
    case "hello":
        if false {
            switch "foo bar" {
            case "foo bar":
                for fi7 := 0; fi7 < 14; fi7++ {
                    if false {
                        if false {
                            var v8 int32 = -459
                            _ = v8
                        } else {
                            if (1 != -1) {
                                v3 = 2
                                for fi9 := 0; fi9 < 12; fi9++ {
                                    _ = len(v4)
                                    v3 = v3
                                    v3 = v3
                                    _ = fi9
                                }
                            }
                            v3 = v3
                        }
                        v3 -= v3
                    }
                    if true {
                        var v10 string = ""
                        v10 = "hello"
                        _ = v10
                    }
                    _ = fi7
                }
            case "\n\t":
                if (!true) {
                    v11 := GenS1{F0: 568, F1: -9223372036854775808, F2: 128}
                    _ = v11
                } else {
                    v12 := GenS1{F0: genFn0(), F1: 1, F2: 234}
                    for fi13 := 0; fi13 < 19; fi13++ {
                        v3 += v3
                        v3 += v3
                        v14 := GenS1{F0: -1, F1: 295, F2: 127}
                        _ = fi13
                        _ = v14
                    }
                    v15 := v2[32768]
                    _ = v12
                    _ = v15
                }
            }
            switch 127 {
            case 32768:
                fmt.Println(v3)
            case 42:
                fmt.Println(v3)
            }
        } else {
            if (1 > 0) {
                var v16 string = "hello"
                var v17 float64 = 1.5
                v18 := GenS1{F0: -1, F1: -452, F2: 2}
                _ = v16
                _ = v17
                _ = v18
            } else {
                v3 = (v3 & 0)
                for fi19 := 0; fi19 < 14; fi19++ {
                    fmt.Println(fi19)
                    _ = fi19
                }
            }
        }
    default:
        v3 = (v3 % ((256) | 1))
    }
    v20 := GenS2{F0: genFn0(), F1: 127}
    for fi21 := 0; fi21 < 6; fi21++ {
        v22 := GenS3{F0: 2147483647}
        var v23 float64 = 3.14159
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 float32 = 1e10
            _ = fi24
            _ = v25
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if ((1 >= -1) && false) {
        var v26 int16 = -1
        v27 := v1[("a" + ("" + "hello"))]
        _ = v26
        _ = v27
    }
    v1[""] = (32767 >= uint16(v3))
    v28 := GenS4{F0: 1, F1: 2147483647}
    fmt.Println(v3)
    if (1e10 > 0.0) {
        v3 = v3
        v3 = 46
        v29 := GenS3{F0: 42}
        _ = v29
    }
    var v30 uint16 = 715
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v30
    _ = v4
}
