package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 int64; F2 int32 }
type GenS7 struct { F0 int }
func genFn0() uint64 {
    var v7 float64 = 1.5
    v7 -= float64(float32(float64(1.5)))
    for fi8 := 0; fi8 < 5; fi8++ {
        v7 -= (3.14159 / v7)
        v7 += (v7 - 3.14159)
        v9 := []int{(fi8 ^ fi8), (fi8 % (((fi8 ^ fi8)) | 1)), (fi8 - 157)}
        _ = fi8
        _ = v9
    }
    var v10 int16 = 29
    _ = v10
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int16) uint {
    var v11 string = ("foo bar" + ("" + "\n\t"))
    v11 = (v11 + "a")
    v12 := []int{127}
    _ = p0
    _ = v11
    _ = v12
    return 961
}

func main() {
    v1 := []int{0}
    var v2 uint64 = 676
    v3 := []uint32{2147483648}
    var v4 uint32 = 539
    var v5 float32 = (-2.25 / -2.25)
    var v6 uint64 = (v2 / ((v2) | 1))
    v2 = (v2 + (^(v2)))
    var v13 int64 = 958
    if (((-2.25 + (-2.25 * 0.1)) + ((v5 - 0.1) * -2.25)) > v5) {
        fmt.Println(v13)
        fmt.Println(v2)
    }
    v6 = 19
    var v14 uint64 = (v2 << ((4294967295) & 63))
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 int8 = 36
        if true {
            if true {
                v2 = v14
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v3)
            var v17 float64 = 0.1
            _ = v17
        }
        _ = fi15
        _ = v16
    }
    v18 := GenS1{F0: 53, F1: 0, F2: true}
    var v19 string = ("\n\t" + "")
    var v20 uint8 = 204
    v21 := GenS2{F0: 4294967295}
    v4 += (^(v4))
    for fi22 := 0; fi22 < 18; fi22++ {
        _ = len(v1)
        v23 := GenS3{F0: (v13 % ((v13) | 1))}
        var v24 string = v19
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        fmt.Println(v2)
        _ = fi25
    }
    switch ("foo bar" + "") {
    case "":
        fmt.Println(v2, v19)
    case "a":
        if false {
            v26 := GenS4{F0: v19, F1: 18446744073709551615}
            switch 32768 {
            case 32767:
                for fi27 := 0; fi27 < 16; fi27++ {
                    fmt.Println(v14, v19, v19)
                    v3 = append(v3, v4)
                    switch 9223372036854775807 {
                    case -152:
                        if len(v3) > 0 {
                            v3[0] = v4
                        }
                    }
                    _ = fi27
                }
            case 128:
                var v28 int32 = -2147483648
                _ = v28
            }
            _ = v26
        } else {
            for fi29 := 0; fi29 < 3; fi29++ {
                if (!(-1 >= -1)) {
                    for fi30 := 0; fi30 < 7; fi30++ {
                        var v31 int32 = -2147483648
                        _ = len(v1)
                        v32 := GenS5{F0: 1}
                        v20 += (v20 >> ((18446744073709551615) & 63))
                        if (int32(v20) > 0) {
                            v33 := GenS5{F0: 838}
                            _ = v33
                        } else {
                            if (42 >= 0) {
                                var v34 int64 = v13
                                var v35 int64 = (v13 >> ((1) & 63))
                                _ = v34
                                _ = v35
                            } else {
                                v36 := GenS6{F0: (v31 + v31), F1: (-(v13)), F2: v31}
                                var v37 int16 = 0
                                fmt.Println(v14, v19)
                                _ = v36
                                _ = v37
                            }
                        }
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                } else {
                    switch (fi29 ^ 223) {
                    case 407:
                        v14 += (v2 ^ v6)
                    case -1:
                        var v38 int64 = (v13 ^ int64(fi29))
                        _ = v38
                    default:
                        v6 = v6
                    }
                }
                v6 -= uint64(v20)
                _ = fi29
            }
        }
    default:
        v39 := GenS7{F0: 864}
        _ = v39
    }
    v19 = v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
