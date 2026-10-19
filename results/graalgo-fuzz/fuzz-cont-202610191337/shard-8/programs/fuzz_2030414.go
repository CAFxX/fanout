package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 int }
type GenS4 struct { F0 int8; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool }
func genFn0(p0 string) uint8 {
    v4 := GenS1{F0: 604, F1: 42, F2: -32768}
    v5 := GenS2{F0: 1}
    v6 := GenS3{F0: 84, F1: 634}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 2
}
func genFn1() uint32 {
    for fi7 := 0; fi7 < 5; fi7++ {
        var v8 int = fi7
        _ = fi7
        _ = v8
    }
    var v9 bool = (true && true)
    for fi10 := 0; fi10 < 6; fi10++ {
        if v9 {
            if (((65535 <= 2) || (9223372036854775807 == -32)) && v9) {
                var v11 uint = 18446744073709551615
                switch (^(fi10)) {
                case -453:
                    if v9 {
                        var v12 uint16 = 880
                        v9 = v9
                        _ = v12
                    } else {
                        for fi13 := 0; fi13 < 6; fi13++ {
                            v14 := []int{fi13}
                            v11 = (v11 - (v11 % ((v11) | 1)))
                            if (-1 >= -82) {
                                if true {
                                    if (9223372036854775807 >= 9223372036854775807) {
                                        v9 = v9
                                        v11 -= v11
                                    } else {
                                        v9 = (false && v9)
                                        v15 := GenS4{F0: 42, F1: 2, F2: 42}
                                        if len(v14) > 0 {
                                            v14[0] = 611
                                        }
                                        _ = v15
                                    }
                                }
                                if len(v14) > 0 {
                                    v14[0] = 270
                                }
                                fmt.Println(v11)
                            }
                            v11 -= v11
                            v11 = 255
                            _ = fi13
                            _ = v14
                        }
                        switch fi10 {
                        case 215:
                            v16 := GenS5{F0: v9}
                            _ = v16
                        case -9223372036854775808:
                            var v17 uint8 = 128
                            _ = v17
                        }
                        fmt.Println(v11, v11, v9)
                    }
                case -818:
                    v18 := []int{(fi10 * (fi10 ^ fi10))}
                    _ = v18
                }
                fmt.Println(v9)
                _ = v11
            }
            var v19 bool = (!(!((((!(-32768 != 685)) || (-1 <= 9223372036854775807)) || (0 != 472)) || (("a" + "foo bar") != "\n\t"))))
            _ = v19
        }
        for fi20 := 0; fi20 < 11; fi20++ {
            var v21 string = "\n\t"
            _ = fi20
            _ = v21
        }
        var v22 uint = 505
        _ = fi10
        _ = v22
    }
    var v23 uint = 4294967295
    _ = v23
    _ = v9
    return 253
}

func main() {
    v1 := []int{0, 9223372036854775807, 503, 743}
    v2 := make(map[int]int)
    var v3 int16 = 505
    v24 := GenS2{F0: 2}
    var v25 uint64 = 706
    var v26 int = 367
    var v27 int16 = v3
    _ = len(v2)
    if false {
        fmt.Println(v3, v27)
        var v28 bool = (967 < 2147483648)
        _ = v28
    }
    if false {
        v3 = (v3 + (v27 / ((v3) | 1)))
    }
    fmt.Println(v26, v3, v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
}
