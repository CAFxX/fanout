package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int; F2 int16 }
type GenS3 struct { F0 int16; F1 int32; F2 uint8 }
func genFn0(p0 uint32, p1 uint16, p2 int64) bool {
    p2 = (p2 | (p2 % (((p2 | (p2 >> ((255) & 63)))) | 1)))
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int8 = -128
        v7 += (v7 - v7)
        _ = fi6
        _ = v7
    }
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return (!(!true))
}
func genFn1(p0 bool, p1 uint) uint16 {
    v8 := []int{}
    v8 = append(v8, 4)
    p0 = (153 == 2)
    _ = p0
    _ = p1
    _ = v8
    return 1
}

func main() {
    var v1 float64 = ((1.5 - 0.1) - (1e10 - 0.1))
    var v2 int = -2147483648
    var v3 int = v2
    var v4 uint32 = 2147483647
    v5 := []int{int(v4), v2, (v2 * v2), (v2 + (v2 % ((-1) | 1)))}
    switch "hello" {
    case "\n\t":
        fmt.Println(v3, v3, v2)
    case "hello":
        if false {
            v9 := GenS1{F0: 32767}
            fmt.Println(v2, v1)
            v10 := GenS1{F0: 2}
            _ = v10
            _ = v9
        } else {
            var v11 int64 = 0
            _ = v11
        }
    }
    var v12 float32 = -2.25
    var v13 float32 = v12
    v3 += int(v4)
    if (32767 > 609) {
        v12 = float32(float64(float32(float64(((0.0 - (-2.25 - float32(1e10))) * v12)))))
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := GenS1{F0: genFn1(false, 65535)}
            v16 := GenS2{F0: 42, F1: 59, F2: -818}
            v12 += (0.1 - v13)
            _ = fi14
            _ = v15
            _ = v16
        }
        v12 -= v12
    }
    switch ("foo bar" + "\n\t") {
    case "a":
        v4 += v4
    }
    v3 = (v2 / ((v3) | 1))
    if true {
        fmt.Println(v1, v4)
    } else {
        v13 = 1.5
        var v17 uint64 = 1
        _ = v17
    }
    if len(v5) > 0 {
        v5[0] = v3
    }
    var v18 int = (v3 ^ (v3 + v2))
    v13 += float32(v1)
    var v19 uint8 = 124
    v20 := GenS1{F0: 32767}
    var v21 int16 = 1
    for fi22 := 0; fi22 < 6; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            switch ("a" + "foo bar") {
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = v2
                }
            case "hello":
                fmt.Println(v21)
            case "":
                if true {
                    if (2147483647 > -1) {
                        fmt.Println(v4, v13, fi23)
                        if false {
                            v21 -= v21
                            if true {
                                if (!("\n\t" != ("a" + "hello"))) {
                                    switch (^(fi23)) {
                                    case 2147483647:
                                        v12 = (((v13 + float32(v1)) / v12) * float32(v1))
                                    case -1:
                                        var v24 float32 = (0.0 + v12)
                                        _ = v24
                                    default:
                                        switch v18 {
                                        case -50:
                                            fmt.Println(v19, v3, v1)
                                        default:
                                            v25 := GenS1{F0: genFn1(true, 4294967295)}
                                            _ = v25
                                        }
                                    }
                                    var v26 int8 = -128
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if (2 >= 0) {
                                        if true {
                                            var v27 uint64 = 1
                                            _ = v27
                                        }
                                        switch v18 {
                                        case 66:
                                            v28 := GenS2{F0: v19, F1: (^(v18)), F2: v21}
                                            fmt.Println(v26)
                                            v2 += v3
                                            _ = v28
                                        }
                                        v5 = append(v5, -9223372036854775808)
                                    } else {
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        fmt.Println(v21)
                                    }
                                    if true {
                                        fmt.Println(fi23, v21, fi22)
                                        switch 2147483647 {
                                        case -696:
                                            _ = len(v5)
                                        case -1:
                                            v26 = v26
                                        case 878:
                                            var v29 string = "hello"
                                            _ = v29
                                        default:
                                            switch (^(v2)) {
                                            case -1:
                                                fmt.Println(v4)
                                            default:
                                                v5 = append(v5, v18)
                                            }
                                        }
                                    }
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v21 = v21
                                    fmt.Println(v12)
                                    _ = v26
                                } else {
                                    var v30 uint = 873
                                    _ = v30
                                }
                                v2 = (fi22 / (((v2 * (fi23 * (v2 | -715)))) | 1))
                            } else {
                                switch ("hello" + "foo bar") {
                                case "foo bar":
                                    fmt.Println(v21, v21)
                                }
                                var v31 uint64 = 414
                                _ = v31
                            }
                        }
                    } else {
                        switch v2 {
                        case 861:
                            v2 += (fi23 | v2)
                        case -678:
                            var v32 uint = 896
                            _ = v32
                        case -1:
                            v13 = (v13 / (((float32(((((v1 * v1) * 1.5) - v1) * v1)) / v12) + (v13 + v13)) - v13))
                        }
                        v12 -= v12
                        v5 = append(v5, fi22)
                    }
                }
            default:
                v2 -= v2
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v19 += v19
            v5 = append(v5, (fi23 + (fi22 - fi23)))
            _ = fi23
        }
        v13 = 3.14159
        _ = fi22
    }
    v4 += 110
    v33 := GenS3{F0: (v21 * v21), F1: -1, F2: v19}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
