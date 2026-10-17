package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 bool }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 16; fi4++ {
        var v5 uint64 = 0
        _ = fi4
        _ = v5
    }
    var v6 int64 = -527
    var v7 int = 32768
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := GenS1{F0: 100}
        _ = fi8
        _ = v9
    }
    if true {
        fmt.Println(v7, v6)
        v7 -= 2147483647
    } else {
        v10 := GenS1{F0: 0}
        if false {
            v11 := GenS1{F0: 255}
            var v12 uint32 = 65535
            switch "foo bar" {
            case "a":
                if (1 <= 42) {
                    v13 := []int{42}
                    if true {
                        v14 := GenS1{F0: 215}
                        for fi15 := 0; fi15 < 6; fi15++ {
                            var v16 int8 = 42
                            for fi17 := 0; fi17 < 7; fi17++ {
                                v13 = append(v13, (v7 * (fi15 + fi17)))
                                var v18 int16 = -378
                                var v19 string = "hello"
                                switch fi15 {
                                case 1:
                                    var v20 bool = false
                                    _ = v20
                                case -668:
                                    v19 = v19
                                case 32767:
                                    fmt.Println(v7, v16)
                                }
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            _ = fi15
                            _ = v16
                        }
                        v21 := GenS2{F0: v6, F1: (v12 > 4294967295)}
                        _ = v14
                        _ = v21
                    } else {
                        var v22 float32 = (1.5 - 0.0)
                        fmt.Println(v12)
                        if true {
                            if (v12 < 74) {
                                var v23 int16 = -1
                                v22 = (v22 / v22)
                                _ = v23
                            } else {
                                v7 -= (v7 >> ((2) & 63))
                                var v24 float32 = v22
                                _ = v24
                            }
                        }
                        _ = v22
                    }
                    _ = v13
                }
            }
            _ = v11
            _ = v12
        }
        _ = v10
    }
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool, p1 int8, p2 int64) int32 {
    fmt.Println(p1, p2)
    fmt.Println(p0)
    var v25 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 42
}

func main() {
    v1 := []string{"foo bar", "foo bar"}
    var v2 uint64 = 18446744073709551615
    v3 := []int{-2147483648, 2147483647}
    v26 := GenS1{F0: 42}
    v1 = append(v1, "a")
    var v27 string = "hello"
    var v28 int16 = 1
    for fi29 := 0; fi29 < 19; fi29++ {
        v2 = (v2 - 18446744073709551615)
        if true {
            if ((true || ((!(1e10 == 1.5)) || false)) && true) {
                _ = len(v3)
            } else {
                switch (fi29 ^ (fi29 / (((^(fi29))) | 1))) {
                case -1:
                    fmt.Println(v28)
                }
            }
            var v30 float32 = float32(0.0)
            switch v27 {
            case "":
                var v31 int16 = v28
                _ = v31
            default:
                var v32 int16 = (^(v28))
                _ = v32
            }
            _ = v30
        } else {
            v2 -= 0
            _ = len(v1)
            if false {
                v27 = v27
            }
        }
        _ = fi29
    }
    var v33 int64 = -751
    var v34 int = 1
    var v35 int = (v34 - (v34 & v34))
    var v36 int16 = v28
    var v37 int32 = genFn1((-1 != (v36 << ((255) & 63))), -128, 1)
    v33 = (v33 & (v33 ^ 0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
