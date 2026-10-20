package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 string; F1 int; F2 int32 }
type GenS6 struct { F0 int64; F1 uint8 }
func genFn0(p0 int) uint64 {
    switch (p0 >> ((2) & 63)) {
    case -9223372036854775808:
        var v5 float32 = 1.5
        _ = v5
    case 972:
        if true {
            for fi6 := 0; fi6 < 7; fi6++ {
                v7 := GenS1{F0: 11, F1: 2147483648, F2: 7}
                var v8 uint64 = 0
                if ("\n\t" <= "foo bar") {
                    v9 := GenS1{F0: -128, F1: 2147483648, F2: -285}
                    p0 = fi6
                    if (v8 >= v8) {
                        var v10 bool = true
                        if (((((("hello" + ("hello" + ("" + "foo bar"))) + ("a" + ("" + "hello"))) + ("hello" + "foo bar")) + ("hello" + "hello")) + ("a" + "a")) < "hello") {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                var v12 bool = ((!true) && v10)
                                var v13 int32 = -1
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            v14 := GenS1{F0: 0, F1: 42, F2: -32768}
                            _ = v14
                        }
                        _ = v10
                    } else {
                        v15 := []int{fi6, p0, 42}
                        _ = v15
                    }
                    _ = v9
                }
                _ = fi6
                _ = v7
                _ = v8
            }
        }
    case 32768:
        var v16 bool = true
        _ = v16
    default:
        v17 := GenS1{F0: 100, F1: 4294967295, F2: -260}
        _ = v17
    }
    if (9223372036854775807 >= 9223372036854775807) {
        if (false && true) {
            fmt.Println(p0)
            v18 := GenS2{F0: 42, F1: 42}
            p0 += (p0 | (p0 & (p0 + (-(p0)))))
            _ = v18
        } else {
            fmt.Println(p0)
        }
        fmt.Println(p0)
    } else {
        if false {
            p0 = (p0 | 1)
        }
    }
    v19 := GenS1{F0: 100, F1: 65535, F2: 252}
    _ = p0
    _ = v19
    return 2
}

func main() {
    v1 := []uint32{42, 2147483647}
    var v2 float32 = float32(0.1)
    var v3 uint32 = 1
    var v4 uint64 = 0
    v20 := GenS2{F0: v4, F1: 789}
    if false {
        var v21 string = "a"
        fmt.Println(v2, v21)
        _ = v21
    }
    v22 := GenS3{F0: false}
    var v23 int32 = -2147483648
    var v24 uint64 = genFn0(-2147483648)
    v2 += ((1.5 + v2) + v2)
    v25 := GenS4{F0: (v23 & (v23 / ((42) | 1)))}
    var v26 int16 = -32768
    var v27 string = "a"
    var v28 uint = uint(v24)
    var v29 uint = v28
    v28 -= v28
    if false {
        if ((false || false) || false) {
            if (!false) {
                v30 := GenS5{F0: v27, F1: -1, F2: -83}
                for fi31 := 0; fi31 < 15; fi31++ {
                    v32 := GenS6{F0: -721, F1: 0}
                    v2 = ((float32(1.5) / v2) / v2)
                    var v33 int16 = v26
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = v30
            }
        }
        for fi34 := 0; fi34 < 1; fi34++ {
            if len(v1) > 0 {
                v1[0] = v3
            }
            _ = fi34
        }
    }
    var v35 int8 = 22
    for fi36 := 0; fi36 < 3; fi36++ {
        v26 = v26
        _ = fi36
    }
    var v37 int32 = (^(v23))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v37
    _ = v4
}
