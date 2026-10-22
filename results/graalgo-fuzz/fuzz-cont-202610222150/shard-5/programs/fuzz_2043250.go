package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
func genFn0(p0 string, p1 bool, p2 uint) uint32 {
    if p1 {
        var v6 float64 = (float64(1.5) + (((1.5 - 1.5) * 0.1) / (1e10 - (0.0 * 1.5))))
        v7 := []int{-9223372036854775808, 4294967295}
        _ = v6
        _ = v7
    } else {
        p1 = p1
        var v8 float32 = 3.14159
        _ = v8
    }
    for fi9 := 0; fi9 < 11; fi9++ {
        var v10 float64 = 0.0
        if (716 >= p2) {
            if (!p1) {
                v11 := GenS1{F0: p0, F1: 0}
                v10 -= (1e10 + 3.14159)
                _ = v11
            }
            v12 := []int{}
            _ = v12
        } else {
            if false {
                v13 := []int{}
                if (532 <= -1) {
                    v14 := GenS1{F0: p0, F1: 0}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 12; fi15++ {
                        fmt.Println(p1, p1)
                        if (false || p1) {
                            if (930 != 378) {
                                v16 := GenS1{F0: (p0 + p0), F1: 503}
                                _ = v16
                            }
                            v17 := GenS1{F0: p0, F1: -1}
                            v18 := GenS1{F0: ("\n\t" + "\n\t"), F1: 9223372036854775807}
                            _ = v17
                            _ = v18
                        }
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi15
                    }
                    fmt.Println(fi9, p0)
                }
                _ = v13
            }
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}
func genFn1() int16 {
    var v19 uint = 799
    if true {
        var v20 uint32 = 1
        fmt.Println(v19)
        _ = v20
    }
    v19 = v19
    var v21 bool = false
    _ = v19
    _ = v21
    return 0
}

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 255
    var v3 uint16 = 42
    var v4 int = -2147483648
    var v5 int16 = 1000
    v2 = (v2 + v2)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 uint32 = 508
        _ = fi22
        _ = v23
    }
    fmt.Println(v5, v5)
    if false {
        if false {
            if true {
                var v24 uint = 65535
                var v25 float64 = -2.25
                _ = v24
                _ = v25
            }
            v2 = 127
            v3 -= v3
        } else {
            v4 = v4
            fmt.Println(v2, v5)
            var v26 uint64 = 542
            _ = v26
        }
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS1{F0: (("\n\t" + "a") + (("foo bar" + ("foo bar" + "hello")) + "\n\t")), F1: 42}
            v29 := []int{(v4 - 2147483647)}
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    v2 -= v2
    var v30 int32 = 2147483647
    fmt.Println(v4)
    v31 := []int{}
    var v32 int32 = (v30 | v30)
    v33 := GenS2{F0: (v3 ^ v3), F1: 133, F2: 42}
    var v34 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
