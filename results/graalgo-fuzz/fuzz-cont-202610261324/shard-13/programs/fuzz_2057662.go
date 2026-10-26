package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS3 struct { F0 int32; F1 int; F2 int8 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint16) int8 {
    var v7 int32 = 2147483647
    v8 := []int{32768}
    var v9 int8 = -1
    switch ("a" + "hello") {
    case "":
        p0 += p0
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (v9 / (((v9 % ((v9) | 1))) | 1))
}

func main() {
    var v1 int16 = 32767
    var v2 int = 42
    var v3 float32 = ((((-2.25 * (0.1 - -2.25)) / 0.1) - (1.5 * 1.5)) + 0.1)
    v4 := []uint32{2, 1, 2147483648, 2147483648}
    var v5 uint16 = 0
    var v6 int32 = -768
    v5 -= (^(v5))
    v10 := GenS1{F0: 0, F1: -9223372036854775808}
    if (!(1 >= 1)) {
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 int16 = v1
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = v5
                if true {
                    if (true && true) {
                        var v16 uint8 = 55
                        var v17 uint64 = 0
                        v4 = append(v4, 918)
                        _ = v16
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 0, F1: 1}
                        var v19 int16 = 1000
                        var v20 int = fi13
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    fmt.Println(v5, v2)
                    switch v2 {
                    case -159:
                        if false {
                            var v21 int8 = 32
                            v1 = (v1 * v1)
                            if len(v4) > 0 {
                                v4[0] = 492
                            }
                            _ = v21
                        } else {
                            v15 -= (v5 & (v15 | (v15 % (((v5 >> ((1) & 63))) | 1))))
                            v15 = 32767
                        }
                    default:
                        var v22 float32 = float32(1e10)
                        _ = v22
                    }
                    var v23 float32 = v3
                    var v24 int16 = 1000
                    _ = v23
                    _ = v24
                } else {
                    var v25 float64 = 1.5
                    var v26 float64 = v25
                    _ = v25
                    _ = v26
                }
                _ = fi14
                _ = v15
            }
            v1 -= -1
            v27 := GenS2{F0: 0, F1: 1, F2: 4294967295}
            _ = fi13
            _ = v27
        }
    }
    if (868 <= -9223372036854775808) {
        fmt.Println(v1, v2)
        for fi28 := 0; fi28 < 11; fi28++ {
            for fi29 := 0; fi29 < 16; fi29++ {
                switch "a" {
                case "foo bar":
                    var v30 float32 = float32(3.14159)
                    _ = v30
                case "":
                    v6 = v6
                default:
                    v31 := GenS2{F0: 42, F1: 952, F2: 436}
                    _ = v31
                }
                switch (fi29 >> ((0) & 63)) {
                case 1:
                    v6 = 307
                default:
                    if true {
                        v1 = (v1 ^ 116)
                        v32 := GenS1{F0: 0, F1: -9223372036854775808}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = v32
                    } else {
                        var v33 uint32 = 4294967295
                        _ = v33
                    }
                }
                v34 := GenS3{F0: v6, F1: (fi28 * 42), F2: -91}
                _ = fi29
                _ = v34
            }
            _ = fi28
        }
        var v35 int = 128
        _ = v35
    }
    v36 := GenS4{F0: (!(v5 <= v5))}
    var v37 bool = true
    var v38 int32 = v6
    fmt.Println(v37)
    if len(v4) > 0 {
        v4[0] = 42
    }
    var v39 int = v2
    switch "a" {
    case "foo bar":
        v3 += v3
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 65535
        }
    }
    if v37 {
        fmt.Println(v3, v37, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
