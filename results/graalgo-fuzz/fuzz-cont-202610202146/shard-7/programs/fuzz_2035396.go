package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint8 }

func main() {
    var v1 bool = true
    var v2 int32 = -224
    var v3 string = ""
    if (!(!(1 < 969))) {
        if false {
            if v1 {
                if (0 >= 65535) {
                    for fi4 := 0; fi4 < 10; fi4++ {
                        for fi5 := 0; fi5 < 15; fi5++ {
                            var v6 uint8 = 128
                            _ = fi5
                            _ = v6
                        }
                        v3 = v3
                        v1 = (!false)
                        _ = fi4
                    }
                }
            } else {
                fmt.Println(v3, v2, v3)
                v7 := []int{42, 1, 1}
                fmt.Println(v1, v1, v1)
                _ = v7
            }
        }
        switch 2147483647 {
        case 1:
            for fi8 := 0; fi8 < 13; fi8++ {
                v9 := []int{(fi8 << ((2) & 63)), 32768}
                _ = fi8
                _ = v9
            }
        case 32768:
            for fi10 := 0; fi10 < 15; fi10++ {
                v1 = (false && (0 > 2))
                fmt.Println(v1, v1, v2)
                switch fi10 {
                case 343:
                    switch (-(fi10)) {
                    case 151:
                        for fi11 := 0; fi11 < 15; fi11++ {
                            v2 -= v2
                            var v12 string = ""
                            v3 = v3
                            _ = fi11
                            _ = v12
                        }
                    case -2147483648:
                        for fi13 := 0; fi13 < 16; fi13++ {
                            var v14 bool = v1
                            v15 := GenS1{F0: 2147483648}
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    default:
                        v1 = ((-128 == -1) && v1)
                    }
                case 560:
                    v2 += (v2 | int32(fi10))
                case 596:
                    var v16 int8 = 42
                    _ = v16
                default:
                    var v17 int = fi10
                    _ = v17
                }
                _ = fi10
            }
        }
        v18 := GenS1{F0: 4294967295}
        _ = v18
    } else {
        v19 := GenS2{F0: v3, F1: -753}
        var v20 int8 = -1
        fmt.Println(v20, v3)
        _ = v19
        _ = v20
    }
    var v21 float32 = (0.0 / 1.0)
    var v22 float64 = 3.14159
    v23 := []int{32767, -779}
    v22 += (v22 + -2.25)
    for fi24 := 0; fi24 < 13; fi24++ {
        var v25 int64 = -1
        v21 -= float32(1e10)
        _ = fi24
        _ = v25
    }
    var v26 float32 = (float32(((v22 - v22) * v22)) / (3.14159 + 0.0))
    v23 = append(v23, 779)
    var v27 int8 = 1
    if len(v23) > 0 {
        v23[0] = 378
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        if (!((1 <= 18446744073709551615) && false)) {
            var v29 string = "hello"
            _ = v29
        }
        _ = fi28
    }
    v30 := GenS3{F0: 42, F1: v1, F2: 1}
    v31 := GenS2{F0: "", F1: (v2 >> ((673) & 63))}
    v26 = v21
    if v1 {
        var v32 int16 = -32768
        _ = v32
    }
    if v1 {
        var v33 uint16 = 1
        fmt.Println(v22, v27)
        for fi34 := 0; fi34 < 14; fi34++ {
            var v35 uint64 = 936
            for fi36 := 0; fi36 < 19; fi36++ {
                v23 = append(v23, fi34)
                v37 := GenS2{F0: (((v3 + ("a" + v3)) + "a") + (v3 + (((("a" + ((("" + v3) + (v3 + ((v3 + ((v3 + "foo bar") + "")) + v3))) + "a")) + "hello") + "\n\t") + (v3 + v3)))), F1: (v2 % (((v2 & 42)) | 1))}
                v38 := GenS4{F0: true}
                _ = fi36
                _ = v37
                _ = v38
            }
            for fi39 := 0; fi39 < 6; fi39++ {
                v40 := GenS5{F0: (v2 ^ int32(v27)), F1: 0}
                if v1 {
                    fmt.Println(v22, v35, v27)
                    var v41 uint64 = v35
                    _ = v41
                } else {
                    var v42 uint64 = v35
                    v21 = v21
                    v3 = ("\n\t" + ("" + "hello"))
                    if len(v23) > 0 {
                        v23[0] = 128
                    }
                    v3 = v3
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
            _ = fi34
            _ = v35
        }
        _ = v33
    } else {
        var v43 float64 = (v22 / v22)
        if len(v23) > 0 {
            v23[0] = 678
        }
        v44 := GenS5{F0: 991, F1: 203}
        _ = v43
        _ = v44
    }
    v27 -= v27
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
}
