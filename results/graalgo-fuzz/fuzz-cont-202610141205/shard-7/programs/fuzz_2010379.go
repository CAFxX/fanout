package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 uint32; F2 int64 }
type GenS5 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 bool, p1 uint64) uint64 {
    switch "hello" {
    case "hello":
        p0 = ((42 == 18446744073709551615) || false)
    case "\n\t":
        v7 := GenS1{F0: 100, F1: 2}
        _ = v7
    case "":
        v8 := GenS1{F0: 70, F1: 972}
        _ = v8
    }
    v9 := GenS2{F0: 32767}
    for fi10 := 0; fi10 < 12; fi10++ {
        var v11 float64 = ((1e10 / ((3.14159 / 1e10) - 0.0)) + (1.5 * ((float64(float32(0.1)) / -2.25) - 3.14159)))
        _ = fi10
        _ = v11
    }
    v12 := []int{}
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return p1
}
func genFn1(p0 uint64, p1 uint16) int {
    p0 = (p0 >> ((0) & 63))
    var v13 float32 = 1e10
    _ = p0
    _ = p1
    _ = v13
    return -909
}

func main() {
    v1 := []int{-1}
    var v2 string = ((("" + "a") + ("foo bar" + "hello")) + "foo bar")
    v3 := []string{}
    var v4 int8 = -35
    v5 := make(map[uint32]uint32)
    v6 := []uint32{}
    fmt.Println(v2)
    v14 := GenS1{F0: -73, F1: 42}
    fmt.Println(v4, v4)
    v2 = v2
    var v15 bool = (832 >= genFn0(true, 18446744073709551615))
    v2 = v2
    for fi16 := 0; fi16 < 15; fi16++ {
        v15 = (!(!v15))
        _ = fi16
    }
    switch 0 {
    case 1:
        v15 = true
    default:
        v2 = (("foo bar" + ("" + "foo bar")) + (v2 + (v2 + "hello")))
    }
    for fi17 := 0; fi17 < 8; fi17++ {
        if false {
            switch v2 {
            case "":
                var v18 uint = 611
                _ = v18
            default:
                if v15 {
                    v2 = ((v2 + ((v2 + ("\n\t" + (v2 + "hello"))) + v2)) + v2)
                    if (127 != 42) {
                        _ = len(v5)
                        _ = len(v5)
                    } else {
                        if (float64((((float32(((float64(0.1) + ((1e10 * 1e10) * 1e10)) - float64((0.1 + 0.0)))) + 1.5) / (1e10 / float32(1e10))) / 1.0)) > ((1e10 + (float64(((1.5 - 0.0) + (1.5 / ((3.14159 - 1.5) + 0.0)))) / (3.14159 * float64(1e10)))) * (1e10 * float64(((0.1 * 3.14159) + (3.14159 - 0.1)))))) {
                            v19 := GenS3{F0: 2}
                            var v20 float64 = 0.1
                            _ = v19
                            _ = v20
                        } else {
                            v4 -= v4
                            v4 -= (^(v4))
                            v21 := GenS4{F0: 608, F1: 42, F2: 1}
                            _ = v21
                        }
                        var v22 uint16 = 0
                        for fi23 := 0; fi23 < 20; fi23++ {
                            switch (fi17 | -9223372036854775808) {
                            case 32767:
                                var v24 float64 = (0.1 * 3.14159)
                                _ = v24
                            case -705:
                                v4 = v4
                            }
                            var v25 bool = (v22 < (v22 % (((v22 >> ((18446744073709551615) & 63))) | 1)))
                            _ = fi23
                            _ = v25
                        }
                        _ = v22
                    }
                }
            }
            v26 := GenS4{F0: int32(v4), F1: 477, F2: 212}
            _ = v26
        } else {
            var v27 bool = v15
            _ = v27
        }
        _ = fi17
    }
    v15 = false
    if (32768 == -2147483648) {
        var v28 int32 = 42
        switch 9223372036854775807 {
        case 1:
            v28 = (v28 % (((v28 % (((v28 * (v28 - (v28 * (v28 + v28))))) | 1))) | 1))
        }
        _ = v28
    } else {
        v4 -= v4
        if (18446744073709551615 > 232) {
            fmt.Println(v2, v15, v2)
        }
    }
    v29 := GenS1{F0: (v4 * (v4 & -128)), F1: 566}
    v6 = append(v6, 1)
    for fi30 := 0; fi30 < 7; fi30++ {
        v2 = ""
        if v15 {
            if v15 {
                v31 := v5[265]
                v32 := GenS5{F0: v4, F1: "foo bar", F2: 42}
                if (-766 > 157) {
                    var v33 string = (v2 + "foo bar")
                    _ = v33
                }
                _ = v31
                _ = v32
            } else {
                _ = len(v1)
                var v34 string = v2
                fmt.Println(v2, v34, v34)
                _ = v34
            }
            v35 := GenS1{F0: 1, F1: 65535}
            _ = v35
        }
        _ = fi30
    }
    if (!(v15 || v15)) {
        v36 := GenS3{F0: 200}
        v37 := GenS2{F0: -774}
        for fi38 := 0; fi38 < 7; fi38++ {
            for fi39 := 0; fi39 < 12; fi39++ {
                fmt.Println(v15, v2, v15)
                _ = fi39
            }
            _ = fi38
        }
        _ = v36
        _ = v37
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
