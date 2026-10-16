package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 string }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0() int8 {
    var v4 int32 = -2147483648
    v5 := []int{2147483647, 2147483647, 0}
    fmt.Println(v4)
    _ = v4
    _ = v5
    return 42
}
func genFn1() uint {
    v6 := GenS1{F0: -44, F1: 127, F2: ((("hello" + ("\n\t" + ("a" + ""))) + "") + ("" + ""))}
    var v7 float64 = (1e10 - (1e10 * 1e10))
    v7 += (v7 + ((0.0 + (0.1 / v7)) - v7))
    switch "hello" {
    case "\n\t":
        v8 := []int{-2147483648, 365, -1, 42}
        _ = v8
    }
    var v9 uint16 = 255
    _ = v6
    _ = v7
    _ = v9
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 75
    var v3 int64 = -586
    _ = len(v1)
    var v10 int16 = -1
    if false {
        for fi11 := 0; fi11 < 2; fi11++ {
            if false {
                switch -1 {
                case 2147483647:
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 uint16 = 0
                        if true {
                            v14 := []int{fi12}
                            if false {
                                switch (fi11 & fi11) {
                                case 431:
                                    v3 -= v3
                                }
                                var v15 float32 = 0.0
                                _ = v15
                            }
                            _ = v14
                        } else {
                            if false {
                                switch fi12 {
                                case -1:
                                    v3 -= (^(v3))
                                default:
                                    if false {
                                        v13 = uint16(v2)
                                    } else {
                                        v10 += 880
                                        var v16 int16 = v10
                                        v3 = (v3 << ((934) & 63))
                                        _ = v16
                                    }
                                }
                            }
                            var v17 uint = 18446744073709551615
                            v18 := GenS2{F0: 18446744073709551615, F1: 4294967295}
                            _ = v17
                            _ = v18
                        }
                        _ = fi12
                        _ = v13
                    }
                }
            }
            _ = fi11
        }
    } else {
        var v19 uint64 = 823
        v19 += 2
        _ = v19
    }
    v2 = 184
    v2 = 110
    if true {
        if true {
            for fi20 := 0; fi20 < 15; fi20++ {
                for fi21 := 0; fi21 < 7; fi21++ {
                    fmt.Println(v10, fi21)
                    var v22 float64 = float64((0.0 * 3.14159))
                    _ = fi21
                    _ = v22
                }
                _ = fi20
            }
            v23 := v1["foo bar"]
            _ = v23
        }
        var v24 int32 = 1
        _ = v24
    }
    v2 += (^(v2))
    v3 = v3
    if (true && true) {
        v2 -= v2
    }
    var v25 uint16 = 2
    var v26 float32 = ((1e10 - 0.1) * 1e10)
    v27 := GenS1{F0: int32(v2), F1: 159, F2: "\n\t"}
    var v28 string = "hello"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
