package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 uint8; F2 int16 }
func genFn0(p0 int64) uint32 {
    p0 -= (p0 / (((p0 & p0)) | 1))
    p0 += (p0 * p0)
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1(p0 uint64) uint64 {
    for fi7 := 0; fi7 < 19; fi7++ {
        if ((true || true) && (!true)) {
            var v8 float64 = 1e10
            if (p0 <= 617) {
                var v9 int32 = 42
                v8 -= -2.25
                v9 += (v9 | (^(v9)))
                _ = v9
            }
            v10 := []int{0, (fi7 / ((4294967295) | 1)), -584}
            _ = v10
            _ = v8
        } else {
            v11 := GenS1{F0: 2, F1: 128}
            if false {
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := []int{}
                    var v14 float64 = -2.25
                    switch (fi7 >> ((18446744073709551615) & 63)) {
                    case 127:
                        var v15 float32 = (0.1 - float32(-2.25))
                        _ = v15
                    case 32767:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := []int{fi16}
                    switch -1 {
                    case 32767:
                        if true {
                            p0 = (p0 / ((p0) | 1))
                            var v18 int64 = 0
                            _ = v18
                        }
                    case -2147483648:
                        var v19 float32 = -2.25
                        _ = v19
                    default:
                        _ = len(v17)
                    }
                    var v20 uint64 = 866
                    var v21 uint16 = 65535
                    _ = fi16
                    _ = v17
                    _ = v20
                    _ = v21
                }
            } else {
                for fi22 := 0; fi22 < 15; fi22++ {
                    if (("hello" + "") == ("foo bar" + "a")) {
                        var v23 uint = 26
                        var v24 uint64 = p0
                        _ = v23
                        _ = v24
                    } else {
                        v25 := GenS1{F0: 1, F1: (fi22 << ((4294967295) & 63))}
                        _ = v25
                    }
                    var v26 string = ""
                    _ = fi22
                    _ = v26
                }
                for fi27 := 0; fi27 < 20; fi27++ {
                    v28 := GenS1{F0: 0, F1: (fi27 + (-(fi7)))}
                    v29 := []int{42}
                    if (((3.14159 == 0.0) && (42 != 128)) || true) {
                        var v30 uint = 555
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v31 := []int{(fi7 * -1)}
                _ = v31
            }
            _ = v11
        }
        fmt.Println(fi7, fi7)
        p0 = p0
        _ = fi7
    }
    v32 := []int{-202, 4294967295, 1, 42}
    _ = p0
    _ = v32
    return (^(p0))
}

func main() {
    var v1 int32 = 198
    var v2 int32 = (v1 & v1)
    v3 := []string{"\n\t"}
    var v4 uint32 = 2
    var v5 float64 = (1.5 / 0.1)
    v6 := make(map[string]uint32)
    v6["foo bar"] = 497
    v5 -= (0.1 * v5)
    v2 -= v2
    fmt.Println(v5, v1, v1)
    v33 := GenS2{F0: -128, F1: 2, F2: 1000}
    fmt.Println(v2, v2)
    var v34 uint32 = (^(v4))
    if true {
        var v35 uint64 = 0
        _ = v35
    } else {
        v34 = (v4 ^ v4)
        fmt.Println(v34, v5, v1)
        fmt.Println(v5, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
