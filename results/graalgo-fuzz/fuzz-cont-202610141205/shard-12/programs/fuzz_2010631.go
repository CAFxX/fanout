package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int32, p1 int64, p2 uint16) string {
    var v4 uint = 754
    v5 := GenS1{F0: ("foo bar" + "\n\t"), F1: 406, F2: 0}
    v6 := []int{-397, -9223372036854775808, 128, 78}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return (("hello" + ("\n\t" + "")) + "hello")
}
func genFn1(p0 uint64, p1 int8, p2 uint8) int32 {
    v7 := []int{32767, 42, 0, 458}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -670
}

func main() {
    var v1 int8 = 109
    var v2 uint32 = 65535
    var v3 int = -526
    if (v2 < v2) {
        var v8 int16 = -1
        switch "a" {
        case "foo bar":
            v9 := []int{v3, v3, (^(v3))}
            _ = v9
        case "":
            v3 = (v3 / (((-(v3))) | 1))
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                v1 = -1
                _ = fi10
            }
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        fmt.Println(v3, v1, v2)
        _ = fi11
    }
    v3 = v3
    var v12 int8 = 0
    for fi13 := 0; fi13 < 7; fi13++ {
        if true {
            for fi14 := 0; fi14 < 17; fi14++ {
                switch "" {
                case "\n\t":
                    v15 := GenS2{F0: (false && false), F1: 0}
                    _ = v15
                case "a":
                    if (3.14159 != 3.14159) {
                        var v16 int32 = int32(v12)
                        v17 := []int{}
                        if (344 < 42) {
                            if false {
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                            } else {
                                fmt.Println(fi14, v2)
                                var v18 bool = (genFn1(393, -19, 2) <= (v16 & (v16 / ((v16) | 1))))
                                var v19 int64 = 1
                                v19 -= (v19 / ((-9223372036854775808) | 1))
                                var v20 float64 = -2.25
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                            v16 = (v16 * (v16 & v16))
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                        } else {
                            v21 := GenS1{F0: "hello", F1: -9223372036854775808, F2: v2}
                            fmt.Println(v12, v16, v3)
                            var v22 string = ""
                            _ = v21
                            _ = v22
                        }
                        _ = v16
                        _ = v17
                    }
                }
                var v23 uint16 = 474
                v24 := []int{(^(fi13)), int(v23), fi13}
                _ = fi14
                _ = v23
                _ = v24
            }
            v25 := GenS3{F0: 255, F1: 0}
            _ = v25
        }
        v12 -= v12
        _ = fi13
    }
    v12 = v12
    var v26 bool = ((0.1 + ((float64(1e10) - 0.1) + (3.14159 / float64(0.1)))) > (float64(0.1) - 0.0))
    for fi27 := 0; fi27 < 9; fi27++ {
        if v26 {
            if v26 {
                v28 := []int{}
                v12 = (v1 ^ (v12 / ((v1) | 1)))
                _ = v28
            }
        } else {
            v29 := []int{(fi27 * -2147483648), (v3 | (fi27 * v3)), (-(fi27))}
            for fi30 := 0; fi30 < 19; fi30++ {
                v31 := GenS3{F0: 32767, F1: 42}
                var v32 float64 = 0.0
                var v33 uint32 = v2
                var v34 int16 = -1
                v26 = v26
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
                _ = v34
            }
            v3 = fi27
            _ = v29
        }
        _ = fi27
    }
    var v35 int64 = -999
    if v26 {
        v36 := []int{v3, v3}
        _ = v36
    } else {
        var v37 bool = (!((0 <= 545) || (0.1 == 1.5)))
        _ = v37
    }
    v38 := []int{(^(v3)), (v3 / ((v3) | 1))}
    switch (v3 / ((v3) | 1)) {
    case -698:
        var v39 uint64 = 4294967295
        _ = v39
    case 42:
        v35 = (^(v35))
    case -916:
        var v40 int64 = (v35 * (v35 - (v35 ^ v35)))
        _ = v40
    default:
        switch v3 {
        case -819:
            var v41 float32 = (1.5 + ((-2.25 + 0.0) / 1.0))
            _ = v41
        default:
            var v42 uint8 = 0
            _ = v42
        }
    }
    var v43 uint64 = 2
    v38 = append(v38, int(v35))
    v1 += (v1 / ((1) | 1))
    v1 = v1
    fmt.Println(v35)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v12
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v38
    _ = v43
}
