package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int }

func main() {
    var v1 uint16 = 0
    var v2 int16 = 0
    var v3 int32 = -875
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 uint16 = (v1 * (v1 + (v1 ^ (v1 | v1))))
        v3 -= v3
        var v6 string = "foo bar"
        _ = fi4
        _ = v5
        _ = v6
    }
    fmt.Println(v1, v3, v3)
    v7 := GenS1{F0: 65535, F1: 128, F2: 0}
    var v8 int16 = v2
    v2 = (v2 | v2)
    switch 128 {
    case -597:
        if true {
            v2 = -174
        } else {
            v9 := GenS1{F0: 785, F1: 42, F2: -2147483648}
            var v10 uint16 = v1
            v11 := []int{4294967295, 42, 9223372036854775807}
            _ = v10
            _ = v11
            _ = v9
        }
    case -252:
        fmt.Println(v8)
    default:
        if false {
            v2 = 297
            var v12 uint = 2
            _ = v12
        } else {
            var v13 int32 = 1
            v14 := []int{32767, 2147483647}
            switch 42 {
            case -9223372036854775808:
                for fi15 := 0; fi15 < 1; fi15++ {
                    var v16 uint16 = v1
                    v3 = v13
                    if false {
                        switch "hello" {
                        case "foo bar":
                            _ = len(v14)
                        case "a":
                            if false {
                                var v17 float64 = (-2.25 / float64((-2.25 + (0.0 + 1e10))))
                                v14 = append(v14, (^(fi15)))
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 19; fi18++ {
                                    var v19 uint64 = 2
                                    v14 = append(v14, fi18)
                                    _ = fi18
                                    _ = v19
                                }
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v13 -= (v3 & 42)
                                    _ = fi20
                                }
                            }
                        default:
                            var v21 int64 = -9223372036854775808
                            _ = v21
                        }
                        var v22 uint32 = 4294967295
                        var v23 uint = 65535
                        _ = v22
                        _ = v23
                    }
                    _ = fi15
                    _ = v16
                }
            case 930:
                var v24 float64 = ((-2.25 - 1e10) + 3.14159)
                _ = v24
            }
            _ = v13
            _ = v14
        }
    }
    if (!true) {
        var v25 int8 = 127
        v26 := []int{4294967295}
        _ = v25
        _ = v26
    } else {
        v2 = (v2 / ((v2) | 1))
    }
    v27 := []int{42, 42, -9223372036854775808}
    _ = len(v27)
    v27 = append(v27, -224)
    fmt.Println(v3, v1)
    if true {
        v3 -= (v3 ^ (-(v3)))
        v8 += v8
    } else {
        if len(v27) >= 2 {
            v27 = v27[1:len(v27)]
        }
    }
    _ = len(v27)
    var v28 uint64 = 1
    v8 -= int16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v7
    _ = v8
}
