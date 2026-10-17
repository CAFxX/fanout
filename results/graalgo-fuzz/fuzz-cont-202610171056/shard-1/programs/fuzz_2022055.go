package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int16; F1 bool; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int32 = 0
    var v2 int32 = 1
    v3 := []int{1, -117}
    var v4 uint16 = 32767
    v5 := []uint32{}
    fmt.Println(v2)
    v1 -= (v2 % (((v2 - (v2 & -1))) | 1))
    var v6 int32 = (v2 * v1)
    var v7 uint8 = 0
    v6 -= 986
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            if false {
                var v9 string = "\n\t"
                for fi10 := 0; fi10 < 10; fi10++ {
                    if true {
                        _ = len(v5)
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        var v11 uint16 = v4
                        if false {
                            v7 = (v7 + (v7 % (((v7 * (v7 >> ((582) & 63)))) | 1)))
                        } else {
                            var v12 uint64 = 42
                            _ = v12
                        }
                        _ = v11
                    }
                    switch "\n\t" {
                    case "hello":
                        var v13 uint16 = (v4 / ((2) | 1))
                        _ = v13
                    case "a":
                        if true {
                            if (v2 == (v2 >> ((311) & 63))) {
                                if false {
                                    _ = len(v3)
                                }
                                var v14 float32 = 1e10
                                v9 = ((((((v9 + (v9 + (("foo bar" + "foo bar") + "a"))) + "") + v9) + (v9 + (((v9 + v9) + "foo bar") + (((v9 + v9) + (v9 + v9)) + v9)))) + "foo bar") + "foo bar")
                                _ = v14
                            } else {
                                v4 -= 149
                                var v15 int32 = v6
                                var v16 int32 = (^(v2))
                                _ = v15
                                _ = v16
                            }
                            v9 = v9
                            var v17 int = -9223372036854775808
                            _ = v17
                        } else {
                            var v18 uint64 = 42
                            var v19 float32 = ((1.5 + (3.14159 - ((3.14159 * float32((1.5 / (((0.1 + 0.0) + -2.25) * float64(-2.25))))) / 3.14159))) + 0.1)
                            if (0.0 >= 3.14159) {
                                v9 = v9
                                if len(v3) > 0 {
                                    v3[0] = (fi10 % (((fi10 % ((-32) | 1))) | 1))
                                }
                                _ = len(v3)
                            } else {
                                if (false && (!(0 > 42))) {
                                    v2 += -158
                                    if (!(2147483648 != uint32(v18))) {
                                        var v20 int8 = 0
                                        var v21 int = fi10
                                        _ = v20
                                        _ = v21
                                    }
                                    switch fi8 {
                                    case 4294967295:
                                        var v22 float64 = 0.0
                                        _ = v22
                                    }
                                }
                                v9 = v9
                            }
                            _ = v18
                            _ = v19
                        }
                    case "\n\t":
                        v1 = (v1 << ((946) & 63))
                    default:
                        switch "hello" {
                        case "\n\t":
                            if len(v5) > 0 {
                                v5[0] = 992
                            }
                        case "a":
                            if false {
                                v2 = (^(v6))
                            }
                        }
                    }
                    v23 := GenS1{F0: v4, F1: 544}
                    v24 := GenS2{F0: int16(v6), F1: true, F2: -486}
                    _ = fi10
                    _ = v23
                    _ = v24
                }
                if true {
                    v5 = append(v5, 431)
                } else {
                    var v25 int32 = v1
                    var v26 int64 = 336
                    if true {
                        var v27 bool = true
                        var v28 uint16 = (v4 >> ((776) & 63))
                        v29 := GenS3{F0: false}
                        _ = v27
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                }
                _ = v9
            } else {
                fmt.Println(v4, fi8)
            }
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 uint32 = 42
                v4 = v4
                _ = fi30
                _ = v31
            }
            var v32 float32 = 3.14159
            _ = fi8
            _ = v32
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            v4 = (v4 << ((4294967295) & 63))
            v34 := GenS4{F0: -128}
            var v35 uint64 = 0
            _ = fi33
            _ = v34
            _ = v35
        }
        v36 := GenS3{F0: true}
        _ = v36
    }
    var v37 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
