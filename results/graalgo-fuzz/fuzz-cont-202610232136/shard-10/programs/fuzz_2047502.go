package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
func genFn0() uint64 {
    switch -2147483648 {
    case 2147483647:
        var v4 int32 = 2147483647
        _ = v4
    case -1:
        var v5 float64 = 3.14159
        _ = v5
    default:
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 uint16 = 42
            v7 = 708
            for fi8 := 0; fi8 < 19; fi8++ {
                v7 -= (v7 | v7)
                var v9 uint32 = 915
                _ = fi8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int8 = -1
        _ = fi10
        _ = v11
    }
    var v12 string = ((((((("\n\t" + "hello") + ("foo bar" + "\n\t")) + "") + "\n\t") + "\n\t") + "foo bar") + ("\n\t" + "a"))
    v13 := GenS1{F0: -855, F1: v12}
    v14 := GenS1{F0: 9223372036854775807, F1: v12}
    _ = v12
    _ = v13
    _ = v14
    return 709
}

func main() {
    v1 := []int{}
    var v2 float32 = -2.25
    var v3 int16 = -539
    var v15 uint64 = 535
    switch "hello" {
    case "foo bar":
        var v16 int8 = 1
        _ = v16
    case "a":
        v15 -= v15
    default:
        if (true || (!true)) {
            var v17 uint64 = 4294967295
            fmt.Println(v3)
            _ = v17
        }
    }
    v1 = append(v1, 969)
    v18 := GenS2{F0: "hello", F1: 42}
    switch "" {
    case "":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        v3 = v3
    }
    fmt.Println(v2, v3, v15)
    var v19 int8 = 0
    var v20 int16 = (^(v3))
    fmt.Println(v20, v19)
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            var v22 int16 = (v20 - (v20 | v3))
            if true {
                v23 := GenS1{F0: -1, F1: ""}
                _ = v23
            } else {
                if ((fi21 << ((2) & 63)) <= 4294967295) {
                    if (((("" + ("\n\t" + "a")) + "\n\t") + "foo bar") == "a") {
                        v20 += -777
                        if true {
                            fmt.Println(v20, v22)
                            if false {
                                if len(v1) > 0 {
                                    v1[0] = -927
                                }
                                fmt.Println(v15)
                                for fi24 := 0; fi24 < 2; fi24++ {
                                    var v25 uint32 = 42
                                    v2 += 1e10
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if (42 < 42) {
                        v22 -= (v22 / (((v3 + (v3 - (v20 << ((0) & 63))))) | 1))
                    } else {
                        fmt.Println(fi21)
                        v15 -= v15
                    }
                    fmt.Println(v19, v19, v22)
                }
            }
            _ = fi21
            _ = v22
        }
        if true {
            var v26 uint32 = 1
            _ = v26
        }
    }
    v3 += v20
    for fi27 := 0; fi27 < 2; fi27++ {
        _ = len(v1)
        _ = fi27
    }
    var v28 float64 = 0.0
    v29 := GenS3{F0: 4294967295, F1: 42}
    switch "foo bar" {
    case "hello":
        fmt.Println(v3, v3)
    case "a":
        var v30 float32 = (0.0 / v2)
        _ = v30
    }
    var v31 int = -956
    v19 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v29
    _ = v3
    _ = v31
}
