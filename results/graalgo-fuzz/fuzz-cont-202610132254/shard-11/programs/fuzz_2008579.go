package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool; F1 int8; F2 int8 }
type GenS4 struct { F0 uint32; F1 string; F2 uint64 }
func genFn0(p0 bool, p1 uint8, p2 int) int {
    if (1 >= 0) {
        if true {
            v7 := GenS1{F0: (486 < 474), F1: (p1 % (((^(p1))) | 1))}
            v8 := []int{}
            if true {
                if (1 <= 0) {
                    for fi9 := 0; fi9 < 11; fi9++ {
                        v10 := GenS2{F0: 2147483648}
                        var v11 int8 = 100
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                } else {
                    var v12 uint = 1
                    p0 = p0
                    if true {
                        if p0 {
                            var v13 uint64 = 896
                            _ = v13
                        } else {
                            for fi14 := 0; fi14 < 12; fi14++ {
                                var v15 uint16 = 0
                                var v16 uint16 = 65535
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            v17 := GenS3{F0: (-667 < -1), F1: 127, F2: 127}
                            fmt.Println(p2, v12)
                            _ = v17
                        }
                    }
                    _ = v12
                }
                if (!(570 < 4294967295)) {
                    var v18 uint8 = (p1 % ((p1) | 1))
                    _ = v18
                }
            }
            _ = v7
            _ = v8
        } else {
            switch "" {
            case "foo bar":
                v19 := GenS4{F0: 2147483647, F1: "\n\t", F2: 4294967295}
                _ = v19
            case "\n\t":
                v20 := GenS3{F0: (false && p0), F1: 42, F2: -15}
                _ = v20
            case "hello":
                var v21 float32 = 1.5
                _ = v21
            }
            var v22 string = ""
            _ = v22
        }
        var v23 int64 = -1
        _ = v23
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 int8 = -128
        _ = fi24
        _ = v25
    }
    fmt.Println(p2, p0)
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    var v1 uint8 = 1
    var v2 uint8 = 127
    v3 := []uint32{65535, 713}
    var v4 string = (("hello" + "foo bar") + ("" + "foo bar"))
    var v5 uint = 1
    v6 := []uint32{1, 328, 65535}
    v26 := GenS4{F0: 4294967295, F1: "hello", F2: uint64(v2)}
    v2 = 33
    switch v4 {
    case "hello":
        var v27 string = "\n\t"
        _ = v27
    case "":
        if false {
            v28 := GenS1{F0: (-1 != 1), F1: (v2 * (v2 - (^(v2))))}
            v2 = 2
            var v29 float64 = 0.0
            _ = v28
            _ = v29
        }
    default:
        for fi30 := 0; fi30 < 3; fi30++ {
            v6 = append(v6, 2147483647)
            _ = fi30
        }
    }
    v5 = v5
    fmt.Println(v5, v5, v4)
    var v31 uint32 = 65535
    _ = len(v3)
    v4 = (v4 + ((("" + "a") + v4) + "hello"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
