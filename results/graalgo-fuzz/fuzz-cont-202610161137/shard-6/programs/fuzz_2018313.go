package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
func genFn0() uint32 {
    var v4 int16 = 32767
    v5 := GenS1{F0: -2147483648, F1: -943, F2: 100}
    switch "hello" {
    case "foo bar":
        v6 := GenS1{F0: 744, F1: -1, F2: -1}
        _ = v6
    case "hello":
        var v7 uint64 = 0
        _ = v7
    }
    var v8 bool = (!true)
    _ = v4
    _ = v5
    _ = v8
    return 2
}
func genFn1(p0 int) int8 {
    for fi9 := 0; fi9 < 3; fi9++ {
        switch -360 {
        case 2147483647:
            v10 := []int{(p0 << ((65535) & 63)), 2147483647, 32768, fi9}
            _ = v10
        default:
            for fi11 := 0; fi11 < 1; fi11++ {
                p0 -= 42
                _ = fi11
            }
        }
        _ = fi9
    }
    if true {
        fmt.Println(p0)
        switch "" {
        case "\n\t":
            var v12 uint64 = 4294967295
            _ = v12
        case "foo bar":
            for fi13 := 0; fi13 < 15; fi13++ {
                if true {
                    p0 = 997
                    p0 += (p0 ^ fi13)
                    switch "hello" {
                    case "hello":
                        var v14 uint8 = 1
                        _ = v14
                    case "a":
                        p0 -= (^(p0))
                    case "":
                        for fi15 := 0; fi15 < 8; fi15++ {
                            v16 := []int{fi13, 32767}
                            p0 = 1
                            v17 := GenS1{F0: -2147483648, F1: 42, F2: -41}
                            p0 = -2147483648
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                } else {
                    v18 := []int{128, (^(p0)), fi13}
                    if len(v18) > 0 {
                        v18[0] = fi13
                    }
                    p0 = (^(fi13))
                    _ = v18
                }
                _ = fi13
            }
        case "hello":
            v19 := GenS1{F0: -89, F1: -9223372036854775808, F2: 1}
            _ = v19
        }
    }
    if true {
        if false {
            var v20 int16 = -987
            v20 -= 1
            v21 := []int{(p0 / ((32768) | 1)), (p0 * (p0 << ((939) & 63))), (p0 / (((p0 % ((p0) | 1))) | 1))}
            _ = v20
            _ = v21
        }
        p0 = (p0 ^ (p0 * -567))
    } else {
        p0 = -363
        var v22 int8 = 1
        _ = v22
    }
    v23 := GenS1{F0: -1, F1: 744, F2: 1}
    if true {
        fmt.Println(p0)
    } else {
        switch p0 {
        case 1:
            p0 = (p0 + (p0 & p0))
        case 2147483647:
            p0 -= (p0 * p0)
        }
        p0 -= -9223372036854775808
        v24 := GenS1{F0: 1, F1: -1, F2: -128}
        _ = v24
    }
    _ = p0
    _ = v23
    return 0
}

func main() {
    var v1 uint8 = 2
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]int)
    var v25 int16 = 1
    v26 := []int{-85}
    v27 := GenS1{F0: 1, F1: 9223372036854775807, F2: 0}
    fmt.Println(v25)
    v1 = v1
    delete(v2, 2)
    v1 = v1
    v1 += v1
    v28 := GenS1{F0: -180, F1: -1, F2: 127}
    fmt.Println(v25, v1)
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 string = ("" + ("hello" + "foo bar"))
        var v31 uint64 = 42
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v25, v25)
    if true {
        var v32 int64 = 476
        switch ("a" + ("hello" + ("hello" + ("" + "a")))) {
        case "foo bar":
            v32 = v32
        case "":
            var v33 uint8 = 209
            _ = v33
        case "\n\t":
            v32 -= 782
        }
        if true {
            switch 1 {
            case 9223372036854775807:
                for fi34 := 0; fi34 < 9; fi34++ {
                    var v35 uint32 = genFn0()
                    v26 = append(v26, 21)
                    v36 := GenS1{F0: 1, F1: 9223372036854775807, F2: 100}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case 0:
                var v37 int64 = int64(v1)
                _ = v37
            case 42:
                v1 -= v1
            }
        } else {
            v26 = append(v26, 128)
            if len(v26) > 0 {
                v26[0] = 744
            }
            for fi38 := 0; fi38 < 1; fi38++ {
                v1 -= v1
                fmt.Println(v1, v32)
                _ = fi38
            }
        }
        _ = v32
    }
    fmt.Println(v1)
    if (v1 < (v1 / (((v1 / ((v1) | 1))) | 1))) {
        v26 = append(v26, 128)
    }
    v39 := GenS1{F0: -1, F1: 1, F2: 127}
    v25 += -32768
    var v40 int16 = v25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v40
}
