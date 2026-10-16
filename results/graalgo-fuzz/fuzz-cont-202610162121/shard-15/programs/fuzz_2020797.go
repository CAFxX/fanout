package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 uint16; F1 int32; F2 string }
type GenS5 struct { F0 bool; F1 bool; F2 uint32 }
func genFn0() uint8 {
    var v4 int = 1
    v4 += (^(v4))
    var v5 uint = 4294967295
    v4 = v4
    _ = v4
    _ = v5
    return 136
}
func genFn1(p0 int) uint8 {
    if false {
        p0 -= p0
    }
    var v6 bool = (103 == 0)
    if (v6 && (4294967295 > 255)) {
        v7 := GenS1{F0: 536, F1: 100, F2: 2}
        _ = v7
    }
    fmt.Println(p0, v6)
    _ = p0
    _ = v6
    return 111
}

func main() {
    var v1 float32 = (float32(1e10) * (float32(float64(-2.25)) * 3.14159))
    var v2 uint16 = 497
    var v3 uint16 = 481
    v1 = v1
    if (128 >= genFn0()) {
        v8 := []int{42, -1, -9223372036854775808}
        _ = v8
    }
    v3 -= (v3 % (((v3 | (v2 % ((v3) | 1)))) | 1))
    v3 -= v3
    for fi9 := 0; fi9 < 3; fi9++ {
        if false {
            if ((v2 % ((v3) | 1)) != v2) {
                if (!true) {
                    for fi10 := 0; fi10 < 11; fi10++ {
                        v2 = v3
                        v11 := []int{(fi9 & (-(fi9))), fi9}
                        _ = fi10
                        _ = v11
                    }
                    if true {
                        switch "" {
                        case "hello":
                            v12 := GenS1{F0: 4294967295, F1: 1, F2: 42}
                            _ = v12
                        }
                        var v13 int32 = 0
                        _ = v13
                    }
                }
                fmt.Println(fi9)
            } else {
                v1 = -2.25
                v14 := GenS2{F0: (v3 * (^(v3))), F1: 1}
                v15 := GenS2{F0: v2, F1: 0}
                _ = v14
                _ = v15
            }
        } else {
            v2 -= 42
        }
        if true {
            v2 = v2
        } else {
            for fi16 := 0; fi16 < 17; fi16++ {
                if true {
                    if (true && false) {
                        var v17 int = (fi9 & (fi16 >> ((2) & 63)))
                        fmt.Println(v17)
                        _ = v17
                    }
                    v18 := []int{(fi16 - 1)}
                    _ = v18
                }
                _ = fi16
            }
        }
        _ = fi9
    }
    v19 := []int{128}
    v1 += (v1 * float32(0.1))
    switch "\n\t" {
    case "foo bar":
        if true {
            if (!false) {
                if (!(true || true)) {
                    fmt.Println(v1, v3)
                    if len(v19) > 0 {
                        v19[0] = -806
                    }
                }
                switch "hello" {
                case "foo bar":
                    v20 := GenS2{F0: v3, F1: 1}
                    _ = v20
                default:
                    if true {
                        v21 := GenS2{F0: (^(v2)), F1: -9223372036854775808}
                        if len(v19) > 0 {
                            v19[0] = 1
                        }
                        _ = v21
                    } else {
                        v1 += float32(3.14159)
                    }
                }
            } else {
                var v22 string = "a"
                switch v22 {
                case "\n\t":
                    var v23 uint16 = v3
                    _ = v23
                case "foo bar":
                    fmt.Println(v2, v2)
                case "a":
                    v3 = 1
                }
                var v24 uint8 = 1
                _ = v22
                _ = v24
            }
        } else {
            var v25 float64 = (1.5 - (0.0 * float64(v1)))
            for fi26 := 0; fi26 < 7; fi26++ {
                for fi27 := 0; fi27 < 14; fi27++ {
                    if true {
                        v19 = append(v19, 1)
                        if ((^(fi26)) >= 2147483647) {
                            if (!false) {
                                fmt.Println(v2, v2)
                                var v28 uint8 = 5
                                _ = v28
                            }
                        }
                    } else {
                        var v29 float32 = v1
                        fmt.Println(v25, v1)
                        _ = v29
                    }
                    fmt.Println(v25, v1)
                    var v30 float32 = v1
                    if (!true) {
                        v2 = (^(v2))
                        v31 := GenS3{F0: v3, F1: true}
                        _ = v31
                    }
                    fmt.Println(fi26)
                    _ = fi27
                    _ = v30
                }
                _ = fi26
            }
            _ = v25
        }
    case "":
        var v32 bool = (((0 > 428) && (2 > 1)) || false)
        _ = v32
    }
    if true {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        v19 = append(v19, 128)
        fmt.Println(v1, v2, v2)
    } else {
        var v33 int64 = -9223372036854775808
        v34 := GenS3{F0: 42, F1: false}
        _ = v33
        _ = v34
    }
    v35 := GenS4{F0: (v2 * 1), F1: 291, F2: "hello"}
    var v36 uint16 = (v2 >> ((42) & 63))
    v36 = v2
    if len(v19) > 0 {
        v19[0] = 32767
    }
    for fi37 := 0; fi37 < 19; fi37++ {
        if false {
            v3 = v3
            v36 = v36
        }
        if ((1 > 315) || true) {
            if len(v19) > 0 {
                v19[0] = fi37
            }
            for fi38 := 0; fi38 < 20; fi38++ {
                v3 -= v2
                fmt.Println(v1, fi38, v1)
                _ = len(v19)
                _ = fi38
            }
        }
        _ = fi37
    }
    fmt.Println(v1, v36, v36)
    v39 := GenS5{F0: (true && true), F1: (127 == 32767), F2: 4294967295}
    if len(v19) > 0 {
        v19[0] = 794
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
}
