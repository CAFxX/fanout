package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 uint64 }

func main() {
    var v1 int16 = 0
    var v2 uint64 = 309
    var v3 int32 = 740
    var v4 int = 32768
    var v5 uint8 = 42
    var v6 string = (("" + "foo bar") + "")
    if true {
        var v7 uint64 = (v2 * 0)
        v5 += (v5 >> ((42) & 63))
        _ = v7
    }
    if (v4 < -622) {
        v2 = v2
        if (true || false) {
            v5 -= v5
        }
    } else {
        v8 := GenS1{F0: false}
        fmt.Println(v2, v4)
        switch v4 {
        case 32768:
            var v9 float32 = float32((1e10 * ((-2.25 + (float64(0.0) - float64(0.1))) * ((((1.5 / 3.14159) / float64(3.14159)) + ((1e10 - (((-2.25 + (1.5 - 0.0)) / 1.0) / 1.0)) / 3.14159)) * (float64(((1.5 - 0.0) - 3.14159)) + ((0.1 / ((((0.0 + 1e10) + -2.25) / -2.25) - -2.25)) * 3.14159))))))
            _ = v9
        case 128:
            var v10 int16 = int16(v3)
            _ = v10
        }
        _ = v8
    }
    if (1 == uint64(v5)) {
        var v11 float64 = (-2.25 / float64((1e10 + 1.5)))
        v11 += ((v11 * (v11 + ((-2.25 / v11) + v11))) * 1.5)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 11; fi12++ {
            v1 += (-(v1))
            _ = fi12
        }
        if true {
            var v13 float64 = 3.14159
            fmt.Println(v1)
            _ = v13
        } else {
            v3 = (^(v3))
        }
    }
    v4 = v4
    v1 -= (^(v1))
    v3 = v3
    var v14 uint8 = (v5 | (v5 * (v5 >> ((220) & 63))))
    if true {
        if false {
            if ((0.0 <= 3.14159) || (v2 < v2)) {
                if true {
                    v15 := GenS1{F0: (154 > 13)}
                    _ = v15
                } else {
                    var v16 int16 = v1
                    v17 := []int{-501, (v4 & (-(v4)))}
                    _ = v16
                    _ = v17
                }
                var v18 uint32 = 42
                _ = v18
            } else {
                for fi19 := 0; fi19 < 3; fi19++ {
                    v20 := GenS1{F0: (!((fi19 & int(v5)) < fi19))}
                    v21 := []int{v4, v4}
                    if true {
                        switch -2147483648 {
                        case 537:
                            switch "" {
                            case "a":
                                for fi22 := 0; fi22 < 5; fi22++ {
                                    v23 := GenS1{F0: (true || (false && ((!false) || (v3 == int32(v2)))))}
                                    _ = fi22
                                    _ = v23
                                }
                            case "\n\t":
                                fmt.Println(v1, v1)
                            default:
                                v6 = v6
                            }
                        default:
                            fmt.Println(v6, v1)
                        }
                    }
                    _ = fi19
                    _ = v20
                    _ = v21
                }
                if false {
                    v24 := []int{32767, 1, 1}
                    v5 += 0
                    _ = v24
                }
            }
            v25 := GenS2{F0: 223, F1: 1}
            _ = v25
        }
    } else {
        fmt.Println(v3, v5)
    }
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int32 = v3
        var v28 int16 = (^(v1))
        _ = fi26
        _ = v27
        _ = v28
    }
    v2 = v2
    for fi29 := 0; fi29 < 2; fi29++ {
        fmt.Println(v3, v4, v4)
        var v30 uint32 = 24
        _ = fi29
        _ = v30
    }
    var v31 bool = (v6 < (((v6 + v6) + (("a" + v6) + "\n\t")) + v6))
    var v32 float64 = 1e10
    var v33 uint32 = 261
    for fi34 := 0; fi34 < 15; fi34++ {
        v35 := []int{(fi34 ^ int(v14)), (v4 | (v4 & 42)), v4}
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
