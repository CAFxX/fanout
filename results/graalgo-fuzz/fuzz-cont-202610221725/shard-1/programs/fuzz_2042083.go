package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 string, p1 int8, p2 int32) uint64 {
    var v5 uint32 = 2
    v6 := GenS1{F0: "hello"}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 0
}

func main() {
    var v1 float32 = (3.14159 / float32(0.1))
    v2 := make(map[int]bool)
    var v3 string = ("hello" + ("foo bar" + ("a" + "\n\t")))
    var v4 int16 = 42
    for fi7 := 0; fi7 < 7; fi7++ {
        v4 -= v4
        v4 += v4
        _ = fi7
    }
    _ = len(v2)
    v8 := GenS1{F0: ""}
    var v9 int8 = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            v4 = 0
            if false {
                v11 := GenS2{F0: -2147483648, F1: 38, F2: 342}
                var v12 int8 = v9
                _ = v11
                _ = v12
            }
        }
        v1 -= (-2.25 / -2.25)
        _ = fi10
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            v9 -= -128
            if (-1 >= v4) {
                v15 := GenS1{F0: ""}
                if (genFn0(v3, v9, -1) <= 18446744073709551615) {
                    if false {
                        if true {
                            if (1 <= genFn0("a", v9, -1)) {
                                var v16 string = v3
                                fmt.Println(v9, v4)
                                _ = v16
                            } else {
                                var v17 uint = 18446744073709551615
                                v18 := GenS2{F0: 555, F1: (v9 * (v9 + (v9 + (v9 | v9)))), F2: v17}
                                _ = v17
                                _ = v18
                            }
                        } else {
                            var v19 int32 = 1
                            var v20 uint32 = 2147483647
                            v1 = (v1 * 3.14159)
                            _ = v19
                            _ = v20
                        }
                    }
                    var v21 uint16 = 255
                    _ = v21
                } else {
                    v22 := GenS2{F0: 42, F1: v9, F2: 255}
                    if false {
                        v23 := v2[fi14]
                        if (int32(v4) < 216) {
                            fmt.Println(v4, fi14, v9)
                            v2[246] = (0 <= 18446744073709551615)
                        }
                        v1 = (v1 + 1e10)
                        v1 -= (0.1 / 3.14159)
                        _ = v23
                    } else {
                        v24 := GenS2{F0: 2147483647, F1: (^(v9)), F2: 950}
                        v25 := GenS2{F0: 1, F1: -4, F2: 462}
                        v1 -= v1
                        _ = v24
                        _ = v25
                    }
                    _ = v22
                }
                var v26 float64 = -2.25
                _ = v15
                _ = v26
            }
            _ = fi14
        }
        v1 -= (0.0 + float32(0.1))
        _ = fi13
    }
    v27 := []int{4294967295}
    for fi28 := 0; fi28 < 12; fi28++ {
        v9 = (v9 << ((0) & 63))
        v1 = v1
        _ = fi28
    }
    switch 9223372036854775807 {
    case 42:
        v9 -= -128
    case 0:
        var v29 uint32 = 4294967295
        _ = v29
    default:
        var v30 int = -986
        _ = v30
    }
    fmt.Println(v9, v3)
    if (false && (v4 == v4)) {
        v9 -= (v9 % ((v9) | 1))
    }
    var v31 uint64 = 18446744073709551615
    v9 = -128
    v3 = ("hello" + "\n\t")
    switch -832 {
    case -214:
        v32 := GenS2{F0: 1, F1: v9, F2: 1}
        _ = v32
    case 32768:
        var v33 uint64 = (v31 * 746)
        _ = v33
    default:
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
    }
    for fi34 := 0; fi34 < 10; fi34++ {
        if ((v31 >> ((250) & 63)) < 0) {
            v35 := GenS1{F0: "hello"}
            fmt.Println(fi34)
            v1 = (v1 / (1.5 / v1))
            _ = v35
        }
        _ = fi34
    }
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v8
    _ = v9
}
