package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int16; F1 uint16 }

func main() {
    var v1 int = 26
    var v2 int = (v1 & (v1 << ((6) & 63)))
    var v3 uint32 = 1
    var v4 bool = false
    if (v1 == (v1 * v1)) {
        if v4 {
            v2 += -328
        } else {
            fmt.Println(v2, v4, v2)
            fmt.Println(v1, v4)
        }
        if (!v4) {
            switch "foo bar" {
            case "hello":
                var v5 int16 = -32768
                _ = v5
            case "foo bar":
                for fi6 := 0; fi6 < 5; fi6++ {
                    var v7 uint64 = 4294967295
                    if ((42 <= 100) || (0.0 > 0.1)) {
                        switch "a" {
                        case "hello":
                            if true {
                                v8 := GenS1{F0: "foo bar"}
                                _ = v8
                            } else {
                                fmt.Println(v3)
                                v1 -= v2
                                fmt.Println(v2, v3, v1)
                            }
                        }
                        for fi9 := 0; fi9 < 17; fi9++ {
                            switch 0 {
                            case 42:
                                var v10 uint8 = 128
                                _ = v10
                            case 0:
                                var v11 float64 = (1.5 * ((((1e10 + -2.25) + (0.1 + (0.1 / 1e10))) - 0.0) + (3.14159 + -2.25)))
                                _ = v11
                            case -1:
                                v4 = false
                            default:
                                v7 += (v7 / ((v7) | 1))
                            }
                            _ = fi9
                        }
                        v12 := []int{-2147483648, fi6, (fi6 % ((v1) | 1)), v1}
                        _ = v12
                    } else {
                        v3 -= v3
                    }
                    _ = fi6
                    _ = v7
                }
            default:
                if false {
                    switch 128 {
                    case 1:
                        fmt.Println(v3)
                    default:
                        var v13 int = v1
                        _ = v13
                    }
                    for fi14 := 0; fi14 < 5; fi14++ {
                        switch "foo bar" {
                        case "hello":
                            if v4 {
                                v15 := []int{}
                                v1 += (-(v2))
                                v16 := GenS2{F0: 1, F1: -1}
                                _ = v15
                                _ = v16
                            }
                        case "foo bar":
                            var v17 int32 = 1
                            _ = v17
                        }
                        v18 := []int{32767, (^(fi14))}
                        v18 = append(v18, 809)
                        _ = fi14
                        _ = v18
                    }
                }
            }
            v3 -= 0
        } else {
            switch v1 {
            case -401:
                v19 := []int{v1, (^(v1)), (v2 | (v1 << ((65535) & 63)))}
                _ = v19
            case 32767:
                v4 = false
            case 0:
                v3 -= v3
            }
        }
    }
    var v20 int = (v2 / (((v1 << ((42) & 63))) | 1))
    switch (-(v2)) {
    case 0:
        for fi21 := 0; fi21 < 2; fi21++ {
            v22 := GenS1{F0: ""}
            _ = fi21
            _ = v22
        }
    default:
        var v23 uint64 = 1
        _ = v23
    }
    fmt.Println(v3, v2)
    v20 = 1
    v24 := []int{}
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS3{F0: -1, F1: 255}
        _ = fi25
        _ = v26
    }
    _ = len(v24)
    if len(v24) > 0 {
        v24[0] = v1
    }
    _ = len(v24)
    _ = len(v24)
    fmt.Println(v1, v2)
    switch (v2 >> ((65535) & 63)) {
    case 1:
        v1 = 0
    case -400:
        var v27 uint = 42
        _ = v27
    case 0:
        for fi28 := 0; fi28 < 1; fi28++ {
            v20 += (v2 & v1)
            v3 += 4294967295
            _ = fi28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
}
