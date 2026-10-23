package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32; F2 bool }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0() uint {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{42}
        fmt.Println(fi7)
        if false {
            fmt.Println(fi7)
            for fi9 := 0; fi9 < 5; fi9++ {
                switch ("" + ((("foo bar" + "a") + "foo bar") + "hello")) {
                case "a":
                    if (!(false || (255 <= 4294967295))) {
                        var v10 int32 = 0
                        _ = v10
                    } else {
                        var v11 int = 9223372036854775807
                        if (("hello" == "\n\t") && false) {
                            if len(v8) > 0 {
                                v8[0] = (^(v11))
                            }
                            v12 := GenS1{F0: (18446744073709551615 > 2), F1: 122, F2: false}
                            var v13 float32 = 0.0
                            _ = v12
                            _ = v13
                        } else {
                            var v14 uint = 18446744073709551615
                            _ = v14
                        }
                        _ = v11
                    }
                }
                if ((0.0 <= 1e10) || true) {
                    if len(v8) > 0 {
                        v8[0] = fi9
                    }
                    fmt.Println(fi7)
                    var v15 uint32 = 322
                    _ = v15
                }
                var v16 uint64 = 4294967295
                _ = fi9
                _ = v16
            }
            for fi17 := 0; fi17 < 7; fi17++ {
                var v18 uint32 = 2
                var v19 int16 = 32767
                v18 += v18
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v20 int = -2147483648
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
        }
        _ = fi7
        _ = v8
    }
    var v21 int32 = 42
    v21 += v21
    v21 -= (v21 | v21)
    v22 := GenS2{F0: 938, F1: 134}
    _ = v21
    _ = v22
    return 255
}
func genFn1(p0 uint) uint {
    p0 += genFn0()
    var v23 int = 9223372036854775807
    if (9223372036854775807 > -1) {
        var v24 int = v23
        _ = v24
    }
    v25 := []int{v23, v23}
    p0 = (p0 | (^(p0)))
    _ = p0
    _ = v23
    _ = v25
    return genFn0()
}

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{2, 426}
    var v3 float64 = v1
    v4 := make(map[string]int)
    v5 := make(map[int]int)
    var v6 int32 = 0
    v6 -= v6
    fmt.Println(v3, v1, v1)
    fmt.Println(v1, v6)
    v3 += v1
    var v26 uint64 = 0
    var v27 int = -1
    fmt.Println(v6, v26)
    v2 = append(v2, 2)
    v1 += (-2.25 - (v3 / v1))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1)
    var v28 bool = true
    for fi29 := 0; fi29 < 2; fi29++ {
        v3 += v3
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
