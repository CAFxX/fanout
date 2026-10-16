package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 int }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }
func genFn0() int32 {
    var v6 int16 = 1000
    v7 := []int{-1, -837, -260}
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 422
}

func main() {
    var v1 int = -129
    var v2 uint32 = 65535
    var v3 float32 = -2.25
    v4 := make(map[int]bool)
    v5 := make(map[int]bool)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint = 225
        if false {
            var v10 int16 = -1
            fmt.Println(fi8)
            if (-41 < 100) {
                var v11 uint32 = v2
                if ((0.0 - 3.14159) == (0.1 / 3.14159)) {
                    if false {
                        switch 9223372036854775807 {
                        case 1:
                            var v12 float64 = 1.5
                            _ = v12
                        case 4294967295:
                            var v13 uint32 = uint32(v9)
                            _ = v13
                        case 9223372036854775807:
                            fmt.Println(v10)
                        default:
                            fmt.Println(v3, fi8, v2)
                        }
                    } else {
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v15 := GenS1{F0: 18446744073709551615}
                            fmt.Println(v10)
                            fmt.Println(v2)
                            _ = fi14
                            _ = v15
                        }
                        v3 += 1e10
                    }
                }
                if true {
                    switch (fi8 >> ((v9) & 63)) {
                    case 0:
                        if (65535 < 811) {
                            var v16 int8 = 100
                            for fi17 := 0; fi17 < 13; fi17++ {
                                v10 = (v10 % ((v10) | 1))
                                var v18 uint16 = uint16(v11)
                                var v19 int64 = 0
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                            switch (v1 & (fi8 * fi8)) {
                            case 0:
                                var v20 int32 = 278
                                _ = v20
                            case 128:
                                fmt.Println(v1)
                            case 236:
                                var v21 int64 = -347
                                _ = v21
                            default:
                                v22 := GenS1{F0: 2}
                                _ = v22
                            }
                            _ = v16
                        }
                    case 42:
                        var v23 uint = (v9 / (((^(v9))) | 1))
                        _ = v23
                    }
                    if false {
                        v24 := []int{(fi8 | (-(fi8))), fi8}
                        var v25 uint32 = v2
                        _ = v24
                        _ = v25
                    } else {
                        if false {
                            switch "hello" {
                            case "\n\t":
                                var v26 string = ""
                                _ = v26
                            case "":
                                v27 := GenS2{F0: (199 != 54)}
                                _ = v27
                            }
                        } else {
                            v4[-1] = (!(true && true))
                            var v28 uint8 = 255
                            _ = v28
                        }
                        v29 := GenS3{F0: "", F1: fi8}
                        _ = v29
                    }
                }
                _ = v11
            } else {
                v30 := []int{-772, (^(v1)), fi8}
                var v31 int16 = 1
                _ = v30
                _ = v31
            }
            _ = v10
        } else {
            switch "foo bar" {
            case "a":
                if false {
                    v32 := []int{(v1 << (((v9 & (v9 | (^(v9))))) & 63)), (^(fi8)), (fi8 * int(v9)), 727}
                    _ = v32
                }
            default:
                if (-1 >= 42) {
                    var v33 uint32 = (^(v2))
                    _ = v33
                } else {
                    v3 -= 1e10
                }
            }
        }
        _ = fi8
        _ = v9
    }
    switch (-(v1)) {
    case 4294967295:
        _ = len(v5)
    default:
        fmt.Println(v3, v2)
    }
    v34 := GenS4{F0: 4294967295, F1: 2, F2: 100}
    v35 := []int{(v1 / ((-373) | 1))}
    v1 = (v1 - (v1 * int(v2)))
    for fi36 := 0; fi36 < 4; fi36++ {
        v2 -= v2
        v2 += 0
        _ = fi36
    }
    _ = len(v4)
    var v37 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
    _ = v4
    _ = v5
}
