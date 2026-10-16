package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool; F1 int; F2 int32 }
type GenS5 struct { F0 int16; F1 int8; F2 uint }
func genFn0(p0 bool, p1 int32, p2 int64) uint16 {
    fmt.Println(p1)
    if p0 {
        switch "\n\t" {
        case "hello":
            var v5 uint = 1
            _ = v5
        default:
            var v6 int32 = p1
            _ = v6
        }
    } else {
        v7 := GenS1{F0: 2, F1: 128, F2: 0}
        _ = v7
    }
    v8 := GenS1{F0: 128, F1: -2147483648, F2: 2}
    v9 := GenS2{F0: -371, F1: 2147483648, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int16) uint64 {
    if (199 != 2147483648) {
        if (!(false && false)) {
            p0 = p0
        } else {
            var v10 string = "hello"
            v11 := []int{}
            _ = v10
            _ = v11
        }
        fmt.Println(p0)
    } else {
        v12 := []int{4294967295, 9223372036854775807, 9223372036854775807, 42}
        v12 = append(v12, 32768)
        _ = v12
    }
    v13 := GenS2{F0: 376, F1: 0, F2: 26}
    var v14 int32 = 42
    if ((0.1 != 1e10) || false) {
        v15 := []int{-2147483648, 0, 32767}
        if len(v15) > 0 {
            v15[0] = 9223372036854775807
        }
        p0 -= (p0 * (p0 & (p0 - 846)))
        _ = v15
    }
    _ = p0
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 uint32 = 4294967295
    var v3 float32 = 3.14159
    var v4 float32 = float32((0.0 / float64((v3 + float32(float64((v3 / (v3 / 0.1))))))))
    var v16 int64 = -1
    var v17 int64 = v16
    v4 = v4
    v2 += v2
    if true {
        if true {
            for fi18 := 0; fi18 < 15; fi18++ {
                fmt.Println(v16, v4)
                _ = fi18
            }
            for fi19 := 0; fi19 < 17; fi19++ {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v1 = v1
                    v3 = 1e10
                    _ = fi20
                }
                _ = fi19
            }
            for fi21 := 0; fi21 < 1; fi21++ {
                if (42 <= 1) {
                    var v22 uint64 = 1
                    if (42 >= -128) {
                        fmt.Println(v17, v16, v17)
                    }
                    _ = v22
                }
                for fi23 := 0; fi23 < 3; fi23++ {
                    switch 447 {
                    case 114:
                        v24 := GenS3{F0: 707, F1: 1}
                        _ = v24
                    case 9223372036854775807:
                        if false {
                            if false {
                                v25 := GenS4{F0: (v17 <= v17), F1: (fi23 + 4294967295), F2: -1}
                                _ = v25
                            } else {
                                if false {
                                    v2 = (v2 & (v2 | (v2 + 42)))
                                    if (1 != -89) {
                                        var v26 bool = true
                                        fmt.Println(v17, v17, fi21)
                                        fmt.Println(v1, fi21, v1)
                                        _ = v26
                                    } else {
                                        v2 = v2
                                        var v27 string = "\n\t"
                                        v17 = v17
                                        _ = v27
                                    }
                                }
                                v1 = ("hello" + "a")
                            }
                        } else {
                            v3 -= v4
                        }
                    }
                    _ = fi23
                }
                var v28 int = -832
                _ = fi21
                _ = v28
            }
        } else {
            var v29 string = v1
            _ = v29
        }
        fmt.Println(v1)
        var v30 int64 = (v16 >> ((708) & 63))
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint = 2
        var v33 float32 = v3
        _ = fi31
        _ = v32
        _ = v33
    }
    v34 := []int{32768}
    v16 -= (v16 * v16)
    for fi35 := 0; fi35 < 3; fi35++ {
        v4 = v4
        var v36 int8 = -128
        for fi37 := 0; fi37 < 20; fi37++ {
            v36 = v36
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    for fi38 := 0; fi38 < 3; fi38++ {
        var v39 float64 = float64(-2.25)
        var v40 string = v1
        v41 := GenS5{F0: 0, F1: -122, F2: 142}
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    var v42 float64 = (0.1 + 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
}
