package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 int16; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    var v3 int8 = 127
    v4 := GenS1{F0: 32767, F1: true}
    v3 = -128
    v3 = v3
    var v5 uint8 = 127
    v1 = false
    if (("hello" + "foo bar") != (("\n\t" + "\n\t") + "a")) {
        v6 := GenS2{F0: 42, F1: 438, F2: 543}
        v7 := []int{int(v3), -9223372036854775808, -2147483648}
        if (((1e10 / -2.25) / (-2.25 + 0.1)) <= -2.25) {
            fmt.Println(v5, v5, v3)
            if v1 {
                if false {
                    if (!v1) {
                        if v1 {
                            if (v1 || false) {
                                switch "\n\t" {
                                case "foo bar":
                                    if ((false || (0 != int64(v5))) && v1) {
                                        v5 = v5
                                    }
                                }
                                v8 := v2[32767]
                                _ = v8
                            } else {
                                var v9 float32 = 0.0
                                _ = v9
                            }
                            var v10 uint8 = (v5 ^ 2)
                            fmt.Println(v5, v5)
                            _ = v10
                        } else {
                            fmt.Println(v5)
                            for fi11 := 0; fi11 < 19; fi11++ {
                                v7 = append(v7, (fi11 & fi11))
                                var v12 int = fi11
                                if (-1 == int16(v5)) {
                                    for fi13 := 0; fi13 < 16; fi13++ {
                                        var v14 int = fi11
                                        v12 = 351
                                        if true {
                                            fmt.Println(v3, v3, v14)
                                            v1 = v1
                                            if (4294967295 < 42) {
                                                v15 := GenS1{F0: 42, F1: v1}
                                                switch fi11 {
                                                case -2147483648:
                                                    _ = len(v7)
                                                case -589:
                                                    var v16 uint16 = 0
                                                    _ = v16
                                                }
                                                _ = v15
                                            } else {
                                                v17 := GenS2{F0: 42, F1: 65535, F2: 0}
                                                var v18 uint8 = (v5 | v5)
                                                v19 := v2[(fi11 ^ fi13)]
                                                _ = v17
                                                _ = v18
                                                _ = v19
                                            }
                                        } else {
                                            if false {
                                                if ((float64(-2.25) - ((-2.25 / 1.5) + 0.1)) <= 1e10) {
                                                    if (!((113 > 0) || v1)) {
                                                        fmt.Println(v12, v5)
                                                        _ = len(v7)
                                                        var v20 uint64 = 926
                                                        _ = v20
                                                    }
                                                    if (v1 && false) {
                                                        var v21 uint64 = 18446744073709551615
                                                        _ = v21
                                                    }
                                                    fmt.Println(v5, fi13, v3)
                                                }
                                                var v22 string = ((("" + (("a" + "") + (("\n\t" + "hello") + ("hello" + "hello")))) + "hello") + "foo bar")
                                                _ = v22
                                            } else {
                                                var v23 int = (^(fi13))
                                                v24 := GenS1{F0: -862, F1: v1}
                                                _ = v23
                                                _ = v24
                                            }
                                            v14 = (v12 * 2147483647)
                                            fmt.Println(fi11)
                                        }
                                        _ = fi13
                                        _ = v14
                                    }
                                    v25 := GenS3{F0: 1, F1: 926}
                                    _ = v25
                                }
                                _ = fi11
                                _ = v12
                            }
                        }
                        v3 += (v3 >> ((189) & 63))
                    } else {
                        switch ((("a" + "foo bar") + ("a" + ("a" + "a"))) + (("\n\t" + "hello") + ("hello" + "hello"))) {
                        case "a":
                            fmt.Println(v3, v5, v5)
                        case "hello":
                            v5 = (v5 - 128)
                        }
                    }
                    v7 = append(v7, 81)
                    fmt.Println(v3)
                } else {
                    var v26 uint32 = 0
                    _ = v26
                }
            }
            v3 = (v3 ^ v3)
        }
        _ = v6
        _ = v7
    }
    if v1 {
        v1 = (v1 && false)
        v5 = (^(v5))
    }
    var v27 float64 = 1e10
    v1 = v1
    v28 := GenS3{F0: 0, F1: 348}
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v1, fi29)
        _ = fi29
    }
    v27 -= 1e10
    _ = len(v2)
    if v1 {
        if v1 {
            v27 -= 1e10
            v1 = v1
            v3 = (v3 & -36)
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            delete(v2, fi30)
            _ = fi30
        }
    }
    if v1 {
        for fi31 := 0; fi31 < 10; fi31++ {
            v32 := []int{fi31}
            if (42 >= 2147483648) {
                fmt.Println(v1)
            }
            fmt.Println(v3, v1, v27)
            _ = fi31
            _ = v32
        }
        var v33 int16 = 411
        v1 = (257 < 339)
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
