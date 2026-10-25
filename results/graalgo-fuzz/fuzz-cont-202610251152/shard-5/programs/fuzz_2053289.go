package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int16 }

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 255
    var v3 string = (v1 + "hello")
    var v4 int = -1
    var v5 string = ""
    var v6 int = (v4 % ((0) | 1))
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 float32 = 0.0
        _ = fi7
        _ = v8
    }
    v6 += v4
    fmt.Println(v2, v5, v1)
    fmt.Println(v3, v3)
    if ((true && (0 == 994)) || true) {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (834 >= 344) {
                v10 := []int{681}
                for fi11 := 0; fi11 < 3; fi11++ {
                    switch v6 {
                    case 2147483647:
                        var v12 uint64 = 0
                        _ = v12
                    }
                    v10 = append(v10, (v6 | v6))
                    _ = fi11
                }
                v6 += 565
                _ = v10
            }
            if (!false) {
                if (57 <= 255) {
                    v13 := []int{}
                    v5 = "foo bar"
                    _ = len(v13)
                    _ = v13
                } else {
                    var v14 uint16 = 2
                    if false {
                        v15 := GenS1{F0: 0}
                        _ = v15
                    } else {
                        v16 := GenS2{F0: 0, F1: v4}
                        v17 := GenS3{F0: 0, F1: -201}
                        var v18 bool = true
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    v19 := GenS1{F0: 4294967295}
                    _ = v14
                    _ = v19
                }
            }
            v5 = v3
            _ = fi9
        }
        v5 = v5
    } else {
        v2 -= (v2 << ((65535) & 63))
    }
    v6 = (v4 % ((0) | 1))
    fmt.Println(v6, v5, v5)
    for fi20 := 0; fi20 < 4; fi20++ {
        if true {
            var v21 int32 = 2147483647
            _ = v21
        } else {
            var v22 uint = 0
            if false {
                switch v4 {
                case 2147483647:
                    if (2 > 42) {
                        for fi23 := 0; fi23 < 13; fi23++ {
                            v2 = 90
                            if true {
                                v24 := GenS3{F0: 32767, F1: 42}
                                var v25 int64 = -1
                                v26 := GenS2{F0: v22, F1: fi20}
                                var v27 bool = ((v25 < v25) && (1 == -128))
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                            v5 = ((v1 + "a") + v5)
                            var v28 float64 = (((0.0 - 1e10) + -2.25) / ((float64(0.0) - 1e10) + float64(-2.25)))
                            _ = fi23
                            _ = v28
                        }
                        var v29 uint16 = 2
                        _ = v29
                    }
                default:
                    fmt.Println(v1)
                }
                v30 := []int{-9223372036854775808}
                _ = v30
            } else {
                fmt.Println(v5, v6)
                for fi31 := 0; fi31 < 16; fi31++ {
                    var v32 int32 = 728
                    v33 := GenS3{F0: 255, F1: 42}
                    v34 := GenS2{F0: uint(v2), F1: 1}
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
                if true {
                    v35 := GenS4{F0: -32768}
                    v6 += 183
                    _ = v35
                }
            }
            _ = v22
        }
        v6 = v6
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
