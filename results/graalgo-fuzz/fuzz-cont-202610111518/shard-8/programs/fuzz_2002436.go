package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64; F1 uint32 }
type GenS7 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint8, p1 uint) int32 {
    v5 := GenS1{F0: 42}
    fmt.Println(p0, p0)
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "a" {
        case "foo bar":
            v7 := []int{1, fi6, fi6, (^(fi6))}
            _ = v7
        case "a":
            var v8 int32 = -567
            _ = v8
        case "":
            var v9 uint = p1
            _ = v9
        }
        v10 := GenS1{F0: -434}
        _ = fi6
        _ = v10
    }
    _ = p0
    _ = p1
    _ = v5
    return -396
}
func genFn1() int64 {
    var v11 uint16 = 65535
    fmt.Println(v11)
    _ = v11
    return int64(v11)
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 float64 = ((float64((((float32(1e10) - ((0.0 - 3.14159) - -2.25)) * -2.25) + -2.25)) * (((0.1 + 1.5) / 0.1) / 1.5)) - 1e10)
    var v3 float32 = 1.5
    var v4 uint64 = 18446744073709551615
    v12 := []int{42, 111, 0, -325}
    fmt.Println(v2)
    fmt.Println(v1, v4, v1)
    if true {
        fmt.Println(v3, v2, v3)
        if true {
            var v13 float32 = (-2.25 + v3)
            _ = v13
        } else {
            if true {
                var v14 int32 = 7
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                v2 += (((((((float64(v3) * v2) - v2) + v2) / 1.0) * v2) + 1e10) - v2)
                _ = v14
            }
        }
    } else {
        v1 -= (v1 * (v1 + (v1 & (v1 ^ v1))))
        var v15 string = "foo bar"
        if false {
            var v16 int64 = -9223372036854775808
            for fi17 := 0; fi17 < 13; fi17++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                fmt.Println(v1)
                v16 = (v16 & v16)
                _ = fi17
            }
            if len(v12) > 0 {
                v12[0] = -2147483648
            }
            _ = v16
        }
        _ = v15
    }
    var v18 uint16 = 32767
    _ = len(v12)
    switch 433 {
    case -802:
        if true {
            var v19 int = 9223372036854775807
            _ = v19
        } else {
            for fi20 := 0; fi20 < 1; fi20++ {
                v12 = append(v12, fi20)
                v21 := GenS2{F0: v1, F1: true, F2: (v4 * v4)}
                _ = fi20
                _ = v21
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                var v23 int32 = -130
                v12 = append(v12, (fi22 / ((-461) | 1)))
                _ = fi22
                _ = v23
            }
        }
    case -1:
        fmt.Println(v1, v4)
    default:
        fmt.Println(v4, v4)
    }
    fmt.Println(v1, v3, v2)
    var v24 uint32 = 880
    v18 = (v18 & v18)
    v25 := GenS3{F0: uint(v4)}
    if true {
        v26 := GenS4{F0: 389}
        _ = v26
    } else {
        if false {
            var v27 int8 = 0
            _ = v27
        }
        v2 += (-2.25 + v2)
        for fi28 := 0; fi28 < 12; fi28++ {
            fmt.Println(v3, v24)
            if true {
                if (true && true) {
                    v3 += (v3 - 0.1)
                    var v29 float64 = 1.5
                    _ = v29
                }
                if ((^(v4)) > v4) {
                    switch -919 {
                    case 583:
                        for fi30 := 0; fi30 < 7; fi30++ {
                            var v31 float64 = float64((v3 + v3))
                            fmt.Println(v2)
                            _ = len(v12)
                            _ = fi30
                            _ = v31
                        }
                    }
                    if ((v2 == v2) && true) {
                        _ = len(v12)
                        v12 = append(v12, 4294967295)
                    }
                }
            }
            _ = fi28
        }
    }
    if (!true) {
        v1 += v1
        var v32 uint16 = v18
        v33 := GenS5{F0: (v1 + v1)}
        _ = v32
        _ = v33
    } else {
        v34 := GenS6{F0: v4, F1: (v24 * v24)}
        v35 := GenS7{F0: 2, F1: 0}
        _ = v34
        _ = v35
    }
    v36 := GenS3{F0: (v1 >> (((^(v1))) & 63))}
    var v37 bool = (true || false)
    if len(v12) > 0 {
        v12[0] = int(v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v4
}
