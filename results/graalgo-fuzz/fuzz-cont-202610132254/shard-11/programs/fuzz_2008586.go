package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
func genFn0(p0 int) uint64 {
    var v5 bool = true
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            fmt.Println(fi6, fi7, v5)
            _ = fi7
        }
        _ = fi6
    }
    var v8 int64 = -1
    _ = p0
    _ = v5
    _ = v8
    return 2
}
func genFn1(p0 string, p1 uint32) string {
    for fi9 := 0; fi9 < 11; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v11 := []int{(fi9 * (-(fi9))), fi9}
            var v12 uint32 = (^(p1))
            v13 := GenS1{F0: p1, F1: ((^(v12)) < p1), F2: 32767}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    var v14 int8 = 127
    _ = p0
    _ = p1
    _ = v14
    return ""
}

func main() {
    var v1 float32 = (-2.25 / (1.5 + 1e10))
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    var v4 float32 = v1
    fmt.Println(v1)
    var v15 uint8 = 249
    v1 = 3.14159
    v4 -= 0.1
    if (true && ((false || (4294967295 < 2)) || (!(!false)))) {
        var v16 uint8 = (^(v15))
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int64 = 420
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    if false {
        v19 := []int{-640, -31, -9223372036854775808}
        if false {
            v15 = v15
        } else {
            if ((-2.25 >= (((((float64(v1) * 3.14159) - 1.5) + (3.14159 * ((0.0 - 0.1) / 1.5))) + ((3.14159 / 1.0) - 0.1)) - 1.5)) && false) {
                for fi20 := 0; fi20 < 17; fi20++ {
                    for fi21 := 0; fi21 < 2; fi21++ {
                        v15 += 42
                        var v22 uint64 = 42
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        _ = fi21
                        _ = v22
                    }
                    v15 = (v15 & (v15 | v15))
                    _ = fi20
                }
                v15 = (v15 % (((v15 * v15)) | 1))
            }
        }
        switch 1 {
        case 9223372036854775807:
            v23 := GenS2{F0: 42, F1: 255, F2: "hello"}
            _ = v23
        case -1:
            var v24 uint16 = 734
            _ = v24
        case 128:
            var v25 uint8 = 2
            _ = v25
        }
        _ = v19
    }
    var v26 uint64 = 0
    v1 += (float32(0.1) + (3.14159 + ((3.14159 - ((1.5 * ((0.1 + (v4 / -2.25)) + v1)) / (v1 + (float32(0.1) * 1e10)))) + 0.1)))
    switch 2147483647 {
    case 42:
        switch -1 {
        case 128:
            for fi27 := 0; fi27 < 1; fi27++ {
                if false {
                    if (v26 <= v26) {
                        v4 += v1
                        if (1e10 > 1.5) {
                            v28 := GenS2{F0: -1, F1: 4294967295, F2: ("\n\t" + "foo bar")}
                            for fi29 := 0; fi29 < 8; fi29++ {
                                fmt.Println(v26, fi29, v1)
                                v30 := []int{42, fi29}
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        } else {
                            v1 -= -2.25
                            v1 -= 1.5
                            v4 += (float32(1.5) + float32(-2.25))
                        }
                        var v31 int16 = 1000
                        _ = v31
                    }
                }
                v1 -= v4
                _ = fi27
            }
        case 378:
            var v32 uint = 2
            _ = v32
        }
    }
    v33 := GenS3{F0: (v26 / (((v26 << ((uint(v15)) & 63))) | 1)), F1: 1}
    var v34 bool = false
    for fi35 := 0; fi35 < 15; fi35++ {
        for fi36 := 0; fi36 < 4; fi36++ {
            v37 := GenS4{F0: -58, F1: -503, F2: 797}
            _ = fi36
            _ = v37
        }
        v3[""] = (0 < 0)
        var v38 uint8 = v15
        _ = fi35
        _ = v38
    }
    if v34 {
        v26 = v26
        if false {
            v39 := []int{}
            _ = v39
        }
    } else {
        v1 = float32((3.14159 / (((0.1 / 1.0) * 1e10) + 1.5)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
