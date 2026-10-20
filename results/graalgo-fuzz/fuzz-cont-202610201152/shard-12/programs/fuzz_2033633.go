package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int8) int8 {
    p0 += (p0 % ((p0) | 1))
    var v5 bool = ((("foo bar" + (("hello" + ("foo bar" + "hello")) + ("a" + ("" + "foo bar")))) + ("a" + "hello")) != "hello")
    var v6 uint8 = 255
    v5 = v5
    v7 := GenS1{F0: 312, F1: 1, F2: 32767}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 int16) uint {
    p0 = p0
    if (2 < 0) {
        var v8 float64 = (-2.25 / 0.1)
        v9 := []int{1}
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: 2, F1: -800, F2: p0}
        _ = v10
    }
    p0 += p0
    p0 = (p0 & p0)
    if false {
        switch "\n\t" {
        case "":
            var v11 int = 32767
            _ = v11
        }
    }
    _ = p0
    return 0
}

func main() {
    v1 := []uint32{769, 402, 2, 139}
    var v2 uint16 = 42
    var v3 float32 = (3.14159 / ((3.14159 + (1.5 - 0.0)) + ((1.5 * 1.5) / 0.1)))
    var v4 int = -983
    v1 = append(v1, 65535)
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 18; fi13++ {
        if false {
            v1 = append(v1, 2147483648)
        }
        _ = fi13
    }
    var v14 int = 0
    if (-1 < 116) {
        var v15 uint64 = 18446744073709551615
        _ = v15
    }
    var v16 uint32 = 0
    for fi17 := 0; fi17 < 8; fi17++ {
        var v18 uint = 2
        if (!(-21 < -1)) {
            var v19 int32 = 2147483647
            for fi20 := 0; fi20 < 14; fi20++ {
                var v21 float64 = ((0.0 + (1.5 - (1.5 - 3.14159))) - 0.0)
                v22 := GenS1{F0: 445, F1: (v19 - (v19 & v19)), F2: 42}
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v14, v14, v2)
    if (0 < (v16 >> ((843) & 63))) {
        v12 = ""
        var v23 string = (v12 + v12)
        _ = v23
    } else {
        var v24 uint64 = 0
        _ = v24
    }
    if true {
        switch "\n\t" {
        case "foo bar":
            v14 += -449
        case "a":
            if len(v1) > 0 {
                v1[0] = (v16 & (^(v16)))
            }
        }
        for fi25 := 0; fi25 < 17; fi25++ {
            if (!true) {
                if (false && false) {
                    if (3.14159 <= v3) {
                        var v26 bool = true
                        fmt.Println(v12, v2, v26)
                        switch (v4 & 42) {
                        case 597:
                            v27 := GenS1{F0: (^(v2)), F1: 42, F2: -32768}
                            _ = v27
                        default:
                            if (42 >= 2147483647) {
                                if len(v1) > 0 {
                                    v1[0] = v16
                                }
                            } else {
                                var v28 uint = 0
                                v26 = v26
                                fmt.Println(v26, fi25, v26)
                                _ = v28
                            }
                        }
                        _ = v26
                    }
                    fmt.Println(v12, v3)
                } else {
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v29 := GenS2{F0: 32767}
                _ = v29
            } else {
                v30 := GenS1{F0: (^(v2)), F1: -898, F2: 1}
                v16 += 472
                _ = v30
            }
            var v31 uint = genFn1(990)
            v16 = (^(v16))
            _ = fi25
            _ = v31
        }
    } else {
        if (false || (!true)) {
            if len(v1) > 0 {
                v1[0] = 42
            }
        }
    }
    switch v12 {
    case "hello":
        fmt.Println(v2, v14)
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    var v32 int8 = genFn0(genFn0(42))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
