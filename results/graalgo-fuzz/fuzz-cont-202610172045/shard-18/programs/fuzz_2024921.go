package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint; F1 bool; F2 int }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint32; F1 int }
func genFn0() uint {
    if (!false) {
        v5 := GenS1{F0: 1}
        v6 := GenS1{F0: 749}
        _ = v5
        _ = v6
    } else {
        var v7 int64 = 268
        _ = v7
    }
    var v8 uint32 = 2147483647
    v8 = (^(v8))
    _ = v8
    return 84
}
func genFn1(p0 int16, p1 bool, p2 int8) uint {
    v9 := []int{1, 1}
    v10 := GenS2{F0: 433, F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 int8 = -128
    v11 := GenS1{F0: 0}
    v4 -= (v4 >> ((123) & 63))
    switch (("" + "") + "a") {
    case "":
        if false {
            if ((!true) && (!(genFn1(0, (3.14159 == 1e10), 20) <= 4294967295))) {
                var v12 uint32 = v1
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 float64 = 1e10
                    var v15 float64 = ((0.1 + (float64(float32(0.1)) + v14)) / (float64(1e10) / 1.0))
                    for fi16 := 0; fi16 < 12; fi16++ {
                        var v17 uint = genFn0()
                        switch "" {
                        case "":
                            v18 := GenS3{F0: 128, F1: 97, F2: 42}
                            _ = v18
                        case "hello":
                            v4 += (-(v4))
                        }
                        v19 := v3[(^(v1))]
                        _ = fi16
                        _ = v17
                        _ = v19
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v12
            } else {
                if (-589 < -1) {
                    fmt.Println(v4, v4)
                    for fi20 := 0; fi20 < 16; fi20++ {
                        if false {
                            v21 := GenS4{F0: 255, F1: (-9223372036854775808 == -216), F2: fi20}
                            if false {
                                v22 := GenS2{F0: -1, F1: -604}
                                _ = v22
                            }
                            var v23 int8 = (-(v4))
                            _ = v21
                            _ = v23
                        } else {
                            var v24 uint64 = 4294967295
                            v1 -= v1
                            v25 := []int{}
                            _ = v24
                            _ = v25
                        }
                        _ = fi20
                    }
                } else {
                    for fi26 := 0; fi26 < 3; fi26++ {
                        v27 := GenS1{F0: 848}
                        switch "a" {
                        case "hello":
                            v28 := GenS5{F0: -1, F1: (fi26 ^ 42)}
                            _ = v28
                        default:
                            if true {
                                fmt.Println(v1, v1)
                                v1 = (v1 >> ((genFn0()) & 63))
                            }
                        }
                        v1 = v1
                        _ = fi26
                        _ = v27
                    }
                }
            }
        }
    case "hello":
        var v29 uint32 = (^(v1))
        _ = v29
    case "foo bar":
        var v30 int32 = -1
        _ = v30
    }
    v31 := GenS5{F0: 42, F1: 32767}
    v1 += (v1 % ((v1) | 1))
    v32 := GenS1{F0: -9223372036854775808}
    v33 := GenS4{F0: 394, F1: (-1 > 9223372036854775807), F2: -2147483648}
    fmt.Println(v4, v1)
    if false {
        v34 := GenS2{F0: 0, F1: 0}
        _ = v34
    }
    var v35 uint = 743
    var v36 int8 = -1
    for fi37 := 0; fi37 < 12; fi37++ {
        v38 := GenS6{F0: (v1 | (^(v1))), F1: fi37}
        _ = fi37
        _ = v38
    }
    if ("foo bar" >= ("foo bar" + ((("a" + ("\n\t" + (("" + ("" + "")) + ""))) + "foo bar") + "a"))) {
        v35 = 0
        v39 := []int{}
        _ = v39
    }
    v1 = (v1 - 2147483648)
    delete(v2, v1)
    v4 -= (v36 / ((v36) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v35
    _ = v36
    _ = v4
}
