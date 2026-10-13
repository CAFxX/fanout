package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 string) uint16 {
    v5 := GenS1{F0: 119}
    if (!true) {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint = 318
            _ = fi6
            _ = v7
        }
        v8 := GenS1{F0: 194}
        var v9 uint = 155
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v5
    return 42
}
func genFn1(p0 uint, p1 uint8, p2 uint16) uint32 {
    v10 := GenS2{F0: (961 >= 42), F1: 42}
    p0 = p0
    v11 := []int{2147483647, 684, 9223372036854775807, -1}
    p0 = p0
    p2 -= 2
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    v1 := []int{4294967295, 128, 9223372036854775807, -55}
    var v2 int8 = 127
    var v3 uint64 = 2
    var v4 int8 = 42
    if (1 < 906) {
        v12 := GenS2{F0: true, F1: -2147483648}
        _ = v12
    } else {
        var v13 uint16 = 471
        v4 += v2
        _ = v13
    }
    switch ("hello" + ("hello" + "")) {
    case "a":
        v4 -= 1
    case "":
        fmt.Println(v2, v4)
    default:
        var v14 uint8 = uint8(v3)
        _ = v14
    }
    if (true || (65535 == 65535)) {
        for fi15 := 0; fi15 < 17; fi15++ {
            for fi16 := 0; fi16 < 3; fi16++ {
                var v17 int8 = (-(v4))
                v18 := GenS2{F0: false, F1: -2147483648}
                if true {
                    v19 := GenS1{F0: 65535}
                    var v20 uint64 = 2
                    _ = v19
                    _ = v20
                } else {
                    fmt.Println(v3, v2)
                    v21 := GenS1{F0: uint32(v3)}
                    _ = v21
                }
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
        var v22 uint32 = 1
        _ = v22
    }
    v4 -= v4
    v3 += 320
    if len(v1) > 0 {
        v1[0] = 32767
    }
    var v23 uint8 = 2
    if (1 != 32767) {
        switch "a" {
        case "a":
            if true {
                var v24 float32 = ((float32(((((((float64((float32(((0.1 / (float64(1.5) * float64((1.5 / 3.14159)))) / 0.1)) * ((-2.25 - 3.14159) * float32(1e10)))) - 1.5) + float64(0.0)) - (0.0 / 0.1)) * 0.0) * ((-2.25 * 3.14159) - ((0.1 * 3.14159) / -2.25))) / (3.14159 / -2.25))) / (float32(1.5) - (1e10 + (3.14159 - 0.0)))) / float32(0.1))
                if true {
                    var v25 int16 = -32768
                    var v26 int16 = v25
                    _ = v25
                    _ = v26
                }
                v27 := GenS1{F0: 452}
                _ = v24
                _ = v27
            }
        case "hello":
            var v28 float64 = 0.1
            _ = v28
        case "\n\t":
            for fi29 := 0; fi29 < 15; fi29++ {
                switch fi29 {
                case 599:
                    v3 = (v3 + v3)
                case -829:
                    var v30 string = "a"
                    _ = v30
                default:
                    switch "hello" {
                    case "":
                        if false {
                            var v31 float64 = -2.25
                            _ = v31
                        } else {
                            for fi32 := 0; fi32 < 9; fi32++ {
                                v4 -= 112
                                v33 := GenS1{F0: 813}
                                if true {
                                    v23 = v23
                                    fmt.Println(v2)
                                    if len(v1) > 0 {
                                        v1[0] = (fi29 / ((297) | 1))
                                    }
                                }
                                _ = fi32
                                _ = v33
                            }
                            var v34 uint = 41
                            if false {
                                var v35 int8 = 54
                                fmt.Println(v2, v34)
                                var v36 uint64 = (v3 / ((v3) | 1))
                                _ = v35
                                _ = v36
                            }
                            _ = v34
                        }
                    case "hello":
                        v1 = append(v1, fi29)
                    }
                }
                for fi37 := 0; fi37 < 3; fi37++ {
                    var v38 uint32 = 4294967295
                    if (0.1 > 1.5) {
                        fmt.Println(v3, v23, fi37)
                    } else {
                        v39 := GenS3{F0: 12}
                        var v40 bool = false
                        _ = v39
                        _ = v40
                    }
                    _ = fi37
                    _ = v38
                }
                _ = fi29
            }
        }
    } else {
        var v41 int32 = -2147483648
        _ = v41
    }
    if len(v1) > 0 {
        v1[0] = -314
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
