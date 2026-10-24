package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 uint8 }
func genFn0() int {
    v5 := GenS1{F0: 0, F1: 42, F2: false}
    if true {
        var v6 float32 = 0.0
        var v7 int8 = 1
        _ = v6
        _ = v7
    }
    _ = v5
    return -9223372036854775808
}
func genFn1(p0 int64, p1 int32, p2 int8) uint64 {
    fmt.Println(p1, p0, p1)
    switch 0 {
    case 32768:
        v8 := GenS1{F0: 42, F1: 1, F2: (42 != 128)}
        _ = v8
    case 952:
        if ((p2 / (((^(p2))) | 1)) < p2) {
            v9 := []int{-890}
            var v10 float64 = (0.1 + (float64(3.14159) + (1e10 * 3.14159)))
            _ = v10
            _ = v9
        } else {
            var v11 int64 = p0
            var v12 uint32 = 1
            _ = v11
            _ = v12
        }
    case 1:
        if false {
            for fi13 := 0; fi13 < 6; fi13++ {
                p1 -= p1
                _ = fi13
            }
            p1 = p1
            v14 := GenS1{F0: (p2 - p2), F1: 18446744073709551615, F2: (true || true)}
            _ = v14
        } else {
            var v15 int16 = 42
            _ = v15
        }
    }
    if true {
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 10; fi17++ {
                switch fi17 {
                case -633:
                    v18 := []int{int(p0), fi17}
                    p1 = (p1 * p1)
                    _ = v18
                case -2147483648:
                    if true {
                        fmt.Println(fi16)
                    } else {
                        fmt.Println(fi17)
                        if (32 < 2) {
                            v19 := []int{}
                            _ = v19
                        } else {
                            if (true || (true && (9223372036854775807 >= (p0 ^ p0)))) {
                                v20 := []int{fi17, fi17}
                                _ = v20
                            } else {
                                v21 := []int{(fi17 & -105), genFn0()}
                                _ = v21
                            }
                            var v22 int16 = 0
                            _ = v22
                        }
                    }
                }
                _ = fi17
            }
            _ = fi16
        }
        p0 -= p0
    }
    _ = p0
    _ = p1
    _ = p2
    return 685
}

func main() {
    var v1 int8 = -128
    var v2 uint32 = 42
    var v3 bool = (!(false && (13 <= 40)))
    var v4 float64 = 1.5
    v23 := GenS2{F0: 2}
    fmt.Println(v1, v2, v1)
    v24 := GenS1{F0: v1, F1: 383, F2: v3}
    v4 += (v4 * v4)
    v3 = ("foo bar" < "hello")
    for fi25 := 0; fi25 < 16; fi25++ {
        fmt.Println(v4, v2)
        var v26 uint16 = 65535
        fmt.Println(v1)
        _ = fi25
        _ = v26
    }
    if false {
        v4 -= (v4 * (((((v4 + (-2.25 / (0.0 - 1e10))) * v4) - (v4 + v4)) - 1e10) - 1.5))
        var v27 uint = 255
        v28 := []int{}
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 uint8 = 1
        _ = fi29
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
