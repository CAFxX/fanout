package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32, p1 int) uint {
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 -= -2147483648
        var v5 uint32 = 4294967295
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                var v7 int16 = 617
                _ = fi6
                _ = v7
            }
        } else {
            p0 -= (-(p0))
            fmt.Println(p1)
        }
        _ = fi4
        _ = v5
    }
    v8 := GenS1{F0: 18446744073709551615, F1: 591}
    if true {
        var v9 int8 = 100
        v10 := GenS2{F0: 0}
        _ = v10
        _ = v9
    }
    p0 = 42
    _ = p0
    _ = p1
    _ = v8
    return 2
}
func genFn1() uint {
    var v11 float32 = (3.14159 - 3.14159)
    fmt.Println(v11)
    _ = v11
    return 4294967295
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{-767}
    v3 := make(map[string]int)
    var v12 uint = 4294967295
    for fi13 := 0; fi13 < 9; fi13++ {
        if false {
            fmt.Println(v12)
        } else {
            switch 0 {
            case 128:
                fmt.Println(fi13, fi13)
            case 9223372036854775807:
                for fi14 := 0; fi14 < 1; fi14++ {
                    switch 1 {
                    case 915:
                        v12 = genFn1()
                    case -1:
                        if false {
                            var v15 uint = v12
                            _ = len(v2)
                            if (!true) {
                                v16 := GenS1{F0: 0, F1: 688}
                                _ = v16
                            } else {
                                switch ("foo bar" + "foo bar") {
                                case "foo bar":
                                    fmt.Println(fi14, fi13, v12)
                                default:
                                    var v17 float64 = 3.14159
                                    _ = v17
                                }
                                _ = len(v1)
                                var v18 bool = true
                                _ = v18
                            }
                            _ = v15
                        } else {
                            var v19 uint8 = 255
                            _ = v19
                        }
                    case 127:
                        v12 = v12
                        v1[2147483648] = "foo bar"
                    }
                    _ = fi14
                }
            case 1:
                fmt.Println(v12)
            default:
                if (-1 <= 0) {
                    v2 = append(v2, (fi13 >> ((v12) & 63)))
                    _ = len(v2)
                    var v20 uint16 = 0
                    _ = v20
                }
            }
        }
        if ((("" + "hello") + "foo bar") < "foo bar") {
            fmt.Println(v12, fi13)
            v12 = 2
            var v21 string = "a"
            _ = v21
        }
        fmt.Println(fi13, v12)
        _ = fi13
    }
    v22 := GenS1{F0: 0, F1: 42}
    v12 = (v12 / ((genFn1()) | 1))
    v12 += 542
    if ((-2.25 - 1e10) < 0.1) {
        _ = len(v2)
        v12 -= v12
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        v12 -= (v12 / ((v12) | 1))
        v24 := GenS1{F0: 42, F1: -1}
        if (!false) {
            v25 := GenS1{F0: 4294967295, F1: 0}
            var v26 uint16 = 1
            var v27 float64 = (-2.25 * float64(0.0))
            _ = v25
            _ = v26
            _ = v27
        }
        _ = fi23
        _ = v24
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS3{F0: ((false && false) || false), F1: 2}
        var v30 int = fi28
        _ = fi28
        _ = v29
        _ = v30
    }
    if false {
        fmt.Println(v12)
        var v31 uint32 = 65535
        var v32 uint8 = 127
        _ = v31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v3
}
