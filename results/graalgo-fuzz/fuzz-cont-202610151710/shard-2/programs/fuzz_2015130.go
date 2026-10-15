package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 uint16) uint8 {
    var v4 uint64 = 4294967295
    v4 = (v4 ^ v4)
    v5 := GenS1{F0: 2147483647, F1: -1}
    fmt.Println(v4, v4)
    var v6 int32 = 42
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 128
}
func genFn1(p0 uint) uint {
    if (2147483647 > 499) {
        switch "\n\t" {
        case "foo bar":
            var v7 string = ("a" + "foo bar")
            _ = v7
        case "":
            switch 1 {
            case -373:
                if (2 > 255) {
                    p0 += 1
                    switch ("foo bar" + "") {
                    case "\n\t":
                        p0 = (p0 % ((p0) | 1))
                    case "a":
                        switch -9223372036854775808 {
                        case 32768:
                            v8 := []int{42, 128, 365}
                            _ = v8
                        case -43:
                            if (1 >= -845) {
                                for fi9 := 0; fi9 < 3; fi9++ {
                                    switch "\n\t" {
                                    case "\n\t":
                                        var v10 int16 = 1000
                                        _ = v10
                                    case "foo bar":
                                        v11 := []int{fi9, -728}
                                        _ = v11
                                    default:
                                        v12 := GenS2{F0: 253}
                                        _ = v12
                                    }
                                    _ = fi9
                                }
                                var v13 int32 = 287
                                var v14 bool = true
                                _ = v13
                                _ = v14
                            } else {
                                switch 740 {
                                case 32767:
                                    if (299 <= int64(p0)) {
                                        p0 += p0
                                        p0 -= (p0 & p0)
                                        for fi15 := 0; fi15 < 5; fi15++ {
                                            switch ("a" + (("" + "\n\t") + "\n\t")) {
                                            case "\n\t":
                                                var v16 string = "hello"
                                                _ = v16
                                            default:
                                                v17 := GenS2{F0: 1}
                                                _ = v17
                                            }
                                            _ = fi15
                                        }
                                    }
                                case 1:
                                    for fi18 := 0; fi18 < 5; fi18++ {
                                        var v19 string = ("\n\t" + "\n\t")
                                        _ = fi18
                                        _ = v19
                                    }
                                }
                                p0 = p0
                                p0 += p0
                            }
                        default:
                            if true {
                                if true {
                                    v20 := GenS2{F0: 127}
                                    if (float32(float64(0.0)) < -2.25) {
                                        p0 = 18446744073709551615
                                    }
                                    _ = v20
                                }
                            }
                        }
                    default:
                        var v21 int8 = 1
                        _ = v21
                    }
                    p0 += p0
                } else {
                    if (467 == 2147483647) {
                        v22 := []int{-9223372036854775808}
                        var v23 float64 = 0.0
                        _ = v22
                        _ = v23
                    }
                }
            }
        case "a":
            var v24 int = 42
            _ = v24
        }
        if true {
            v25 := []int{0}
            _ = v25
        }
        for fi26 := 0; fi26 < 8; fi26++ {
            var v27 int8 = 0
            _ = fi26
            _ = v27
        }
    }
    var v28 float32 = 3.14159
    v28 += v28
    var v29 bool = false
    _ = p0
    _ = v28
    _ = v29
    return p0
}

func main() {
    var v1 uint8 = 189
    var v2 int16 = 1000
    var v3 uint64 = 450
    v30 := GenS3{F0: (v3 << ((genFn1(273)) & 63)), F1: (v3 / ((v3) | 1))}
    var v31 float64 = ((((1.5 * 0.1) / ((((1.5 * 0.0) * 1e10) + 0.1) * 0.1)) / ((1.5 - 1e10) + (1.5 * float64(1.5)))) * ((3.14159 - 0.0) * float64((0.0 / 3.14159))))
    v32 := []int{127, 42}
    var v33 uint32 = 0
    v1 += v1
    v1 += uint8(v33)
    _ = len(v32)
    v34 := GenS1{F0: 127, F1: 42}
    v3 = 1
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    for fi35 := 0; fi35 < 3; fi35++ {
        var v36 uint32 = (v33 & 0)
        for fi37 := 0; fi37 < 17; fi37++ {
            v31 = (1.5 - (v31 / ((1.5 / ((v31 / ((v31 - ((((float64(-2.25) * ((((float64(0.1) / float64(1.5)) / v31) + (1e10 + (v31 * (v31 * (v31 / 3.14159))))) * float64(float32(float64(1.5))))) / ((v31 - v31) - (0.1 / v31))) + (((-2.25 / v31) / v31) / v31)) * 3.14159)) * v31)) * (v31 / v31))) + v31)))
            var v38 int16 = v2
            _ = fi37
            _ = v38
        }
        fmt.Println(v3)
        _ = fi35
        _ = v36
    }
    fmt.Println(v1, v33)
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    var v39 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v39
}
