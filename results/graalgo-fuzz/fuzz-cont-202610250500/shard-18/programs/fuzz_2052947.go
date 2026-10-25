package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS3 struct { F0 int; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS5 struct { F0 int; F1 int32 }
func genFn0(p0 uint64) int16 {
    var v6 int8 = 34
    switch -1 {
    case -1:
        v7 := []int{-611, 4294967295}
        _ = v7
    case 0:
        fmt.Println(p0, v6)
    default:
        if false {
            fmt.Println(p0, p0)
            var v8 uint16 = 1
            if false {
                var v9 bool = true
                _ = v9
            }
            _ = v8
        }
    }
    v6 = -1
    v10 := GenS1{F0: 42, F1: 42, F2: 127}
    _ = p0
    _ = v10
    _ = v6
    return -1
}
func genFn1() int8 {
    v11 := GenS2{F0: 0, F1: 2147483647, F2: 42}
    switch 127 {
    case 313:
        var v12 bool = (false && false)
        _ = v12
    }
    _ = v11
    return -11
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 float32 = 3.14159
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 455
    for fi13 := 0; fi13 < 2; fi13++ {
        if (4294967295 > 18446744073709551615) {
            switch ((((("a" + ((("a" + "foo bar") + ("" + "hello")) + "")) + "\n\t") + "") + "hello") + "foo bar") {
            case "a":
                if (v3 >= 1.5) {
                    v3 = float32(0.1)
                }
            case "hello":
                var v14 string = "\n\t"
                _ = v14
            case "foo bar":
                if (1 != 194) {
                    v15 := GenS2{F0: 1, F1: 2147483647, F2: genFn1()}
                    v3 -= ((1.5 - 1e10) - 0.1)
                    _ = v15
                } else {
                    v16 := GenS2{F0: 743, F1: 2, F2: 42}
                    for fi17 := 0; fi17 < 14; fi17++ {
                        var v18 uint64 = 0
                        v3 = v3
                        _ = fi17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 9; fi19++ {
                        fmt.Println(v5)
                        _ = len(v4)
                        fmt.Println(fi13, v3)
                        _ = fi19
                    }
                    _ = v16
                }
            }
            v20 := GenS3{F0: (-(fi13)), F1: 4294967295, F2: -9223372036854775808}
            if (-185 <= 0) {
                var v21 int32 = 576
                v3 = (float32(float64(1e10)) - (v3 + v3))
                v21 -= (^(v21))
                _ = v21
            } else {
                fmt.Println(fi13)
                v5 -= 837
            }
            _ = v20
        }
        v22 := GenS3{F0: fi13, F1: 422, F2: 9223372036854775807}
        var v23 string = (("foo bar" + "hello") + ((("" + "\n\t") + "foo bar") + "hello"))
        _ = fi13
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        if ((-1 <= genFn0(42)) && false) {
            v25 := GenS4{F0: 18446744073709551615, F1: genFn1(), F2: v5}
            _ = v25
        } else {
            v3 = 0.0
            var v26 uint = 777
            fmt.Println(fi24)
            _ = v26
        }
        v27 := GenS5{F0: fi24, F1: 2147483647}
        _ = fi24
        _ = v27
    }
    v28 := []int{938, 9223372036854775807}
    v29 := GenS3{F0: 127, F1: 0, F2: 1}
    v5 += (^(v5))
    if (false || true) {
        v2[422] = (!(v3 != -2.25))
    } else {
        var v30 int64 = 927
        _ = len(v2)
        _ = v30
    }
    v3 -= (3.14159 * (v3 + -2.25))
    fmt.Println(v3, v3)
    switch -2147483648 {
    case 42:
        fmt.Println(v3, v5)
    case 127:
        if (113 <= -786) {
            v31 := GenS3{F0: -414, F1: 1, F2: 24}
            _ = v31
        }
    case -765:
        var v32 float64 = ((3.14159 / (1e10 / 1e10)) * (float64(v3) + 3.14159))
        _ = v32
    }
    v2[2147483648] = true
    v33 := v1[("foo bar" + "a")]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
