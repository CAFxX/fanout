package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int32; F2 int32 }
type GenS5 struct { F0 uint16; F1 string; F2 bool }
type GenS6 struct { F0 uint; F1 uint32; F2 bool }
func genFn0(p0 string) int64 {
    v7 := GenS1{F0: 1, F1: 1}
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            p0 = "hello"
            v9 := []int{}
            _ = fi8
            _ = v9
        }
        switch p0 {
        case "foo bar":
            p0 = p0
        case "":
            p0 = ""
        case "a":
            if false {
                v10 := GenS2{F0: 65535, F1: 0}
                var v11 int8 = -24
                var v12 int32 = 418
                _ = v10
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: 32767, F1: 0}
                if (2 < 255) {
                    p0 = (p0 + p0)
                    for fi14 := 0; fi14 < 17; fi14++ {
                        v15 := []int{fi14}
                        v15 = append(v15, -761)
                        _ = fi14
                        _ = v15
                    }
                    v16 := GenS3{F0: 994}
                    _ = v16
                } else {
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v18 := GenS1{F0: fi17, F1: 713}
                        var v19 int8 = 42
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    if (442 < 502) {
                        if false {
                            var v20 uint = 319
                            fmt.Println(v20, v20)
                            _ = v20
                        }
                        v21 := []int{32768}
                        var v22 uint32 = 2147483647
                        _ = v21
                        _ = v22
                    }
                }
                if false {
                    fmt.Println(p0)
                    var v23 uint8 = 0
                    var v24 int32 = 161
                    _ = v23
                    _ = v24
                } else {
                    switch p0 {
                    case "a":
                        v25 := []int{9223372036854775807}
                        _ = v25
                    }
                }
                _ = v13
            }
        }
        var v26 uint8 = 128
        _ = v26
    }
    _ = p0
    _ = v7
    return -9223372036854775808
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -123
    var v3 uint16 = 1
    var v4 int16 = -32768
    v5 := make(map[uint32]int)
    var v6 uint = 42
    v27 := v1[2147483648]
    v27 = (v27 | v27)
    v28 := []int{}
    fmt.Println(v6)
    v2 = genFn0("a")
    for fi29 := 0; fi29 < 4; fi29++ {
        v4 -= int16(v2)
        var v30 float32 = -2.25
        _ = fi29
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := GenS4{F0: 2, F1: 0, F2: 42}
        if true {
            switch "a" {
            case "a":
                fmt.Println(v4, v27)
            }
        }
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        v34 := GenS5{F0: 0, F1: ("hello" + ("foo bar" + "foo bar")), F2: ((-2.25 * (1e10 / (3.14159 * 1e10))) == (float64((0.1 * 1e10)) / 0.1))}
        var v35 uint64 = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := GenS6{F0: (v6 / ((18446744073709551615) | 1)), F1: v27, F2: ((v3 >> (((v6 & (v6 >> (((v6 - (^(v6)))) & 63)))) & 63)) == 0)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
