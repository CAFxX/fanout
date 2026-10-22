package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 int16, p1 uint32, p2 int8) uint16 {
    var v4 int32 = 1
    p1 += (p1 ^ (p1 / ((65535) | 1)))
    fmt.Println(p0, v4, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 42
}

func main() {
    var v1 int32 = 42
    var v2 uint32 = 522
    var v3 int32 = v1
    switch ("" + ("\n\t" + ("hello" + "\n\t"))) {
    case "hello":
        var v5 int16 = -32768
        _ = v5
    default:
        if false {
            fmt.Println(v2, v3)
            for fi6 := 0; fi6 < 1; fi6++ {
                v7 := GenS1{F0: 0}
                _ = fi6
                _ = v7
            }
            if false {
                fmt.Println(v3, v3)
                for fi8 := 0; fi8 < 12; fi8++ {
                    v9 := []int{(fi8 << ((1) & 63))}
                    switch 771 {
                    case 0:
                        v2 = 807
                    default:
                        if false {
                            v9 = append(v9, fi8)
                        }
                    }
                    v3 = v1
                    _ = fi8
                    _ = v9
                }
                var v10 int32 = (-(v1))
                _ = v10
            }
        }
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        if ((v2 ^ (v2 ^ v2)) >= 0) {
            var v12 float32 = float32((1.5 - (3.14159 + (0.1 / (1.5 / 1e10)))))
            v13 := []int{528, fi11}
            for fi14 := 0; fi14 < 11; fi14++ {
                switch (fi11 * (fi11 >> ((18446744073709551615) & 63))) {
                case 128:
                    v3 = (v3 / ((v1) | 1))
                    fmt.Println(fi14)
                    v3 = v3
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v15 float64 = (-2.25 * (0.1 / (0.0 - (((float64(v12) - (3.14159 * 1e10)) / (0.0 + 1e10)) * (1e10 / ((0.0 * 1e10) + 1.5))))))
                    if true {
                        var v16 uint = 1
                        var v17 int8 = 12
                        _ = v16
                        _ = v17
                    } else {
                        if false {
                            v15 -= 3.14159
                            v1 = int32(v2)
                            if ((4294967295 < 42) || (true && (141 >= 0))) {
                                var v18 int = 42
                                _ = v18
                            } else {
                                v19 := GenS2{F0: 100, F1: 255}
                                _ = v19
                            }
                        }
                        v12 -= ((0.1 * v12) - v12)
                        var v20 uint64 = 4294967295
                        _ = v20
                    }
                    v21 := GenS3{F0: 0, F1: 213, F2: 1}
                    _ = v15
                    _ = v21
                case 32767:
                    v1 += (v3 % ((v3) | 1))
                default:
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                }
                if len(v13) > 0 {
                    v13[0] = (fi14 * fi11)
                }
                _ = fi14
            }
            _ = v12
            _ = v13
        } else {
            v22 := []int{}
            _ = v22
        }
        _ = fi11
    }
    var v23 uint16 = 32767
    var v24 uint16 = (v23 + 345)
    fmt.Println(v2, v1)
    var v25 int8 = -1
    for fi26 := 0; fi26 < 10; fi26++ {
        var v27 string = ("" + ("a" + "foo bar"))
        fmt.Println(v25, fi26)
        var v28 int8 = -128
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := []int{32767}
    _ = len(v29)
    v30 := GenS3{F0: 255, F1: 578, F2: (^(v23))}
    v31 := GenS1{F0: 323}
    var v32 uint8 = 2
    switch "\n\t" {
    case "a":
        _ = len(v29)
    case "\n\t":
        var v33 int16 = 32767
        _ = v33
    }
    fmt.Println(v25)
    switch "a" {
    case "":
        var v34 uint64 = 18446744073709551615
        _ = v34
    case "a":
        v2 += 2
    case "\n\t":
        switch -97 {
        case 9223372036854775807:
            v32 += v32
        default:
            switch 830 {
            case 399:
                for fi35 := 0; fi35 < 13; fi35++ {
                    if false {
                        switch ("" + ("a" + "\n\t")) {
                        case "foo bar":
                            var v36 int = 9223372036854775807
                            _ = v36
                        default:
                            for fi37 := 0; fi37 < 3; fi37++ {
                                var v38 bool = false
                                v24 += (v23 | (v24 + v24))
                                var v39 float64 = 1.5
                                _ = fi37
                                _ = v38
                                _ = v39
                            }
                        }
                    } else {
                        v40 := GenS2{F0: v25, F1: (v32 ^ v32)}
                        _ = v40
                    }
                    fmt.Println(v24, v24, v32)
                    _ = fi35
                }
            }
        }
    default:
        v25 -= v25
    }
    fmt.Println(v2, v25)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
