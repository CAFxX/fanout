package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{"", "", "hello", "foo bar"}
    var v2 uint8 = 47
    var v3 string = ("a" + "a")
    var v4 uint8 = v2
    v5 := []int{65}
    for fi6 := 0; fi6 < 1; fi6++ {
        fmt.Println(v4)
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 2147483647}
    switch v3 {
    case "a":
        fmt.Println(v4, v2, v2)
    case "hello":
        var v9 uint8 = 42
        _ = v9
    }
    switch (v3 + (v3 + v3)) {
    case "":
        v10 := GenS2{F0: 397}
        _ = v10
    case "foo bar":
        var v11 uint = 255
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if false {
            fmt.Println(v2)
        } else {
            v4 = (v4 ^ 0)
            for fi13 := 0; fi13 < 5; fi13++ {
                if (-2.25 >= float32(3.14159)) {
                    if (-115 == 127) {
                        v14 := GenS1{F0: 42, F1: -2147483648}
                        var v15 int64 = 9223372036854775807
                        v3 = v3
                        v16 := GenS3{F0: v4, F1: fi13, F2: 127}
                        _ = v14
                        _ = v15
                        _ = v16
                    } else {
                        fmt.Println(v2, v2, fi12)
                        fmt.Println(v4, v4, v4)
                        var v17 float64 = ((0.0 / -2.25) * 3.14159)
                        _ = v17
                    }
                    var v18 uint = 1
                    switch fi13 {
                    case 785:
                        v2 = (v4 + v4)
                    case 2147483647:
                        if true {
                            switch fi13 {
                            case 128:
                                var v19 int64 = 864
                                _ = v19
                            case 42:
                                if len(v1) > 0 {
                                    v1[0] = (("hello" + (v3 + "\n\t")) + v3)
                                }
                            case 32768:
                                v3 = ((v3 + v3) + "hello")
                            default:
                                fmt.Println(v3, v2)
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                    case -1:
                        var v20 uint64 = 1
                        _ = v20
                    default:
                        v3 = v3
                    }
                    _ = v18
                } else {
                    v4 = (v4 ^ (v4 / (((v2 ^ (v4 % ((193) | 1)))) | 1)))
                    v21 := GenS4{F0: -1, F1: 42}
                    if ((!(1000 != -837)) || false) {
                        var v22 uint32 = 450
                        var v23 int8 = 100
                        _ = v22
                        _ = v23
                    } else {
                        v24 := GenS5{F0: (fi13 ^ (fi13 | (fi12 & 4294967295)))}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v24
                    }
                    _ = v21
                }
                _ = fi13
            }
            var v25 int16 = 1000
            _ = v25
        }
        if false {
            if true {
                v2 += 206
            }
        } else {
            if true {
                v2 += v2
                v3 = v3
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((2 != 126) && false) {
                    fmt.Println(v3, v2, v3)
                    v4 -= (v4 - v2)
                    var v26 float64 = -2.25
                    _ = v26
                }
            }
            v4 = v2
        }
        _ = fi12
    }
    if (1 <= 2) {
        var v27 int = 9223372036854775807
        _ = v27
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v28 int = 127
        _ = v28
    }
    fmt.Println(v3, v2, v3)
    var v29 int = -852
    v30 := GenS2{F0: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v8
}
