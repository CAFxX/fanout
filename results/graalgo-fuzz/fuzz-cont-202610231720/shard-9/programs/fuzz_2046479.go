package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 uint16, p1 uint16, p2 int) uint32 {
    var v7 uint32 = 0
    var v8 float64 = 3.14159
    v9 := GenS1{F0: p1, F1: -128, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1() int32 {
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 uint32 = genFn0(2, 255, fi10)
        _ = fi10
        _ = v11
    }
    if (128 < 226) {
        var v12 uint16 = 316
        _ = v12
    } else {
        if (false || (true || false)) {
            if true {
                if true {
                    switch ("a" + "a") {
                    case "a":
                        v13 := []int{32767, 42, 488, -2147483648}
                        _ = v13
                    default:
                        var v14 uint32 = 42
                        _ = v14
                    }
                    var v15 float32 = 0.0
                    _ = v15
                }
            } else {
                if (!(0.0 <= -2.25)) {
                    if true {
                        if (true || false) {
                            v16 := []int{127, 9223372036854775807, 2147483647}
                            _ = v16
                        }
                    } else {
                        var v17 uint = 1
                        var v18 int = -1
                        _ = v17
                        _ = v18
                    }
                } else {
                    v19 := GenS2{F0: 2147483647, F1: -456, F2: -631}
                    switch 0 {
                    case 128:
                        var v20 string = ("\n\t" + "hello")
                        _ = v20
                    }
                    _ = v19
                }
                var v21 int8 = -4
                v22 := GenS1{F0: 0, F1: (v21 / ((v21) | 1)), F2: 1}
                _ = v21
                _ = v22
            }
            if true {
                var v23 int = 876
                v24 := GenS3{F0: 127, F1: 255, F2: 0}
                _ = v23
                _ = v24
            }
            v25 := GenS2{F0: 42, F1: 906, F2: 9223372036854775807}
            _ = v25
        } else {
            for fi26 := 0; fi26 < 1; fi26++ {
                v27 := GenS4{F0: -32768, F1: 387}
                _ = fi26
                _ = v27
            }
            if (!false) {
                v28 := GenS4{F0: -1, F1: -1}
                _ = v28
            }
        }
        var v29 uint64 = 2
        v30 := GenS2{F0: 42, F1: 32768, F2: 9223372036854775807}
        _ = v29
        _ = v30
    }
    v31 := GenS2{F0: -205, F1: 127, F2: -9223372036854775808}
    _ = v31
    return 42
}

func main() {
    var v1 float64 = -2.25
    var v2 uint32 = 159
    var v3 int = -550
    v4 := make(map[string]int)
    v5 := []string{"hello", "hello"}
    v6 := make(map[uint32]bool)
    v1 = v1
    v32 := GenS5{F0: (^(v2))}
    var v33 uint16 = 32767
    v1 = v1
    v4["hello"] = -2147483648
    var v34 int16 = 42
    var v35 float64 = (3.14159 / (1e10 * v1))
    switch (("hello" + ("a" + "foo bar")) + "a") {
    case "a":
        v36 := GenS5{F0: (^(v2))}
        _ = v36
    case "foo bar":
        _ = len(v6)
    case "hello":
        v3 -= 0
    default:
        var v37 uint = 356
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
