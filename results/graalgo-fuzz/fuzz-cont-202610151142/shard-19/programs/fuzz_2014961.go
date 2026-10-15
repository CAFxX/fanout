package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint16; F1 string; F2 uint32 }
func genFn0(p0 int16) uint {
    p0 = -263
    v7 := []int{470, 0}
    v8 := GenS1{F0: 65535, F1: p0}
    for fi9 := 0; fi9 < 10; fi9++ {
        fmt.Println(fi9, p0)
        v10 := GenS2{F0: (p0 - p0)}
        _ = fi9
        _ = v10
    }
    v7 = append(v7, 9223372036854775807)
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int8, p1 int) string {
    var v11 int = p1
    var v12 float64 = (1.5 * (0.1 - 1.5))
    v13 := GenS1{F0: 351, F1: 1}
    var v14 float64 = (v12 + 0.0)
    if false {
        if (4294967295 == 1) {
            for fi15 := 0; fi15 < 18; fi15++ {
                v11 += (v11 * -2147483648)
                for fi16 := 0; fi16 < 14; fi16++ {
                    v17 := GenS3{F0: p0}
                    switch "\n\t" {
                    case "hello":
                        v11 = (fi16 * fi15)
                    case "a":
                        if false {
                            var v18 float64 = 1e10
                            v19 := []int{p1, 127, fi15}
                            if false {
                                if false {
                                    var v20 uint8 = 2
                                    switch p1 {
                                    case 241:
                                        fmt.Println(fi16)
                                    }
                                    _ = v20
                                } else {
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                }
                                fmt.Println(fi15)
                                if true {
                                    p0 -= (^(p0))
                                    v12 = ((0.0 - 1e10) / (v18 * v14))
                                    var v21 uint8 = 199
                                    var v22 uint = 65535
                                    var v23 int8 = p0
                                    _ = v21
                                    _ = v22
                                    _ = v23
                                }
                            } else {
                                var v24 int32 = 42
                                if len(v19) > 0 {
                                    v19[0] = fi16
                                }
                                _ = v24
                            }
                            var v25 uint = genFn0(-96)
                            _ = v18
                            _ = v19
                            _ = v25
                        } else {
                            var v26 float32 = 1e10
                            _ = v26
                        }
                    case "\n\t":
                        v27 := []int{(p1 << ((65535) & 63)), fi16, fi15, (v11 << ((4294967295) & 63))}
                        var v28 float64 = 1e10
                        _ = v27
                        _ = v28
                    }
                    fmt.Println(fi16, v12)
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            var v29 int8 = (p0 + (p0 + -128))
            var v30 string = ""
            _ = v29
            _ = v30
        } else {
            v31 := []int{1, v11}
            fmt.Println(v11, p0, v14)
            v11 += (v11 >> ((255) & 63))
            _ = v31
        }
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return "hello"
}

func main() {
    var v1 int = 801
    var v2 bool = (632 <= 18446744073709551615)
    v3 := make(map[int]int)
    v4 := make(map[int]string)
    var v5 float32 = (-2.25 / -2.25)
    var v6 bool = v2
    v32 := v4[(v1 << ((948) & 63))]
    var v33 int16 = 0
    for fi34 := 0; fi34 < 6; fi34++ {
        v35 := GenS4{F0: ((fi34 + int(v33)) == 126)}
        v2 = v6
        _ = fi34
        _ = v35
    }
    if v6 {
        v36 := []int{(v1 << ((42) & 63)), v1}
        _ = v36
    } else {
        if true {
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 uint32 = 2
                _ = fi37
                _ = v38
            }
            var v39 int8 = -1
            v40 := GenS1{F0: 42, F1: -144}
            _ = v39
            _ = v40
        }
        var v41 int = (v1 - (v1 ^ v1))
        v2 = v6
        _ = v41
    }
    v42 := GenS5{F0: 0, F1: "a", F2: 2147483648}
    fmt.Println(v5, v5)
    var v43 uint = 1
    delete(v3, (-(v1)))
    v5 = (3.14159 - float32(1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
