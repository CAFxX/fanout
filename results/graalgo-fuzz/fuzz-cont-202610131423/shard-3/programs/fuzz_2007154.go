package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 string }
type GenS2 struct { F0 int8; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 uint16) int32 {
    p0 += uint64(p1)
    switch -434 {
    case 32767:
        v5 := []int{-151, 741, 689}
        _ = v5
    case 2147483647:
        fmt.Println(p1)
    case 42:
        var v6 int16 = -1
        _ = v6
    default:
        var v7 int = 2147483647
        _ = v7
    }
    p0 = (p0 ^ 883)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    v1 := []int{-33}
    var v2 int32 = -1
    var v3 int16 = -32768
    v4 := make(map[string]string)
    var v8 string = (("foo bar" + "a") + (("foo bar" + "\n\t") + "\n\t"))
    v8 = v8
    var v9 int64 = 1
    if (2 >= 39) {
        _ = len(v4)
        v8 = v8
    } else {
        var v10 string = "\n\t"
        v10 = ("" + v10)
        _ = v10
    }
    switch "foo bar" {
    case "":
        if (!(974 <= 0)) {
            v9 += (^(v9))
            if true {
                switch "" {
                case "a":
                    v1 = append(v1, -420)
                case "":
                    fmt.Println(v2, v8)
                }
                v11 := GenS1{F0: v2, F1: (v3 + (v3 + v3)), F2: ""}
                _ = v11
            }
            var v12 int64 = -1
            _ = v12
        } else {
            var v13 string = "\n\t"
            fmt.Println(v9, v13, v2)
            _ = v13
        }
    }
    switch 127 {
    case -611:
        if true {
            fmt.Println(v9, v2)
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 int64 = (v9 % ((v9) | 1))
                switch fi14 {
                case 0:
                    for fi16 := 0; fi16 < 9; fi16++ {
                        if true {
                            var v17 int16 = (^(v3))
                            fmt.Println(v17, v2, fi14)
                            if true {
                                v2 += (^(v2))
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            _ = v17
                        } else {
                            var v18 uint = 2
                            _ = v18
                        }
                        var v19 float64 = ((1.5 + 3.14159) + 0.1)
                        _ = len(v4)
                        if (!false) {
                            var v20 float32 = 1.5
                            fmt.Println(v3, v9, fi16)
                            _ = v20
                        }
                        var v21 uint = 424
                        v21 = 65535
                        _ = fi16
                        _ = v19
                        _ = v21
                    }
                case -2147483648:
                    if false {
                        var v22 float32 = 3.14159
                        _ = v22
                    }
                }
                _ = fi14
                _ = v15
            }
        }
    case 1:
        v9 -= (-(v9))
    case -9223372036854775808:
        v23 := GenS1{F0: v2, F1: (v3 + v3), F2: "a"}
        _ = v23
    default:
        var v24 uint8 = 78
        _ = v24
    }
    var v25 int = 128
    for fi26 := 0; fi26 < 14; fi26++ {
        v9 = 0
        var v27 string = (v8 + "a")
        _ = fi26
        _ = v27
    }
    switch v8 {
    case "a":
        var v28 uint8 = 128
        _ = v28
    case "foo bar":
        v29 := GenS1{F0: (v2 << ((124) & 63)), F1: (v3 * v3), F2: ""}
        _ = v29
    default:
        if true {
            var v30 float32 = (3.14159 - 0.0)
            for fi31 := 0; fi31 < 12; fi31++ {
                v8 = v8
                _ = fi31
            }
            _ = v30
        } else {
            if (!false) {
                v32 := GenS1{F0: 1, F1: v3, F2: (("" + v8) + v8)}
                _ = v32
            }
            v33 := GenS2{F0: 1, F1: -81, F2: 13}
            if ((!((int(v9) >= int(v2)) && false)) || false) {
                var v34 float64 = 0.1
                _ = v34
            } else {
                var v35 float32 = ((-2.25 * -2.25) * 0.1)
                switch ("hello" + v8) {
                case "a":
                    fmt.Println(v35, v3, v2)
                default:
                    var v36 float64 = (0.0 - 3.14159)
                    _ = v36
                }
                _ = v35
            }
            _ = v33
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
