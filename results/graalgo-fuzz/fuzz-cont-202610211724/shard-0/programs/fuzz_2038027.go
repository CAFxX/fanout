package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64; F1 string }
func genFn0() uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: fi7, F1: 407}
        _ = fi7
        _ = v8
    }
    var v9 int16 = 32767
    v9 = (^(v9))
    _ = v9
    return 255
}
func genFn1(p0 string, p1 string, p2 string) uint {
    v10 := []int{42}
    var v11 string = p2
    p2 = "a"
    for fi12 := 0; fi12 < 13; fi12++ {
        p0 = "hello"
        if true {
            var v13 uint16 = genFn0()
            _ = v13
        } else {
            v11 = p0
        }
        if len(v10) > 0 {
            v10[0] = fi12
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 255
}

func main() {
    v1 := []int{4294967295, -42, 4294967295, -2147483648}
    var v2 uint8 = 99
    var v3 uint16 = 42
    var v4 int64 = 42
    v5 := []uint32{}
    var v6 uint16 = 891
    v6 = genFn0()
    var v14 string = "\n\t"
    v15 := GenS1{F0: -837, F1: 32767}
    var v16 float64 = (0.0 - ((0.1 / 1.0) / ((3.14159 + ((0.0 * (1e10 + ((1e10 * 1.5) * -2.25))) - 1.5)) + ((((-2.25 * (float64((-2.25 * 1e10)) * 3.14159)) - float64((1.5 * (-2.25 / ((0.1 / (float32(0.1) / 1.0)) / 3.14159))))) * ((1e10 / 1.0) / 1e10)) - float64((1e10 - ((float32((-2.25 + 3.14159)) * (3.14159 / ((0.0 / 1.0) - -2.25))) / -2.25)))))))
    if (104 != 255) {
        var v17 int16 = 66
        for fi18 := 0; fi18 < 19; fi18++ {
            var v19 uint64 = 18446744073709551615
            switch "" {
            case "hello":
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 uint = 948
                    v16 += 1.5
                    fmt.Println(v21, v2, v19)
                    v5 = append(v5, 959)
                    var v22 int16 = int16(v4)
                    var v23 int64 = (v4 & (v4 % (((v4 + v4)) | 1)))
                    var v24 uint8 = (v2 * (^(v2)))
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                }
            default:
                switch 261 {
                case -2147483648:
                    switch v14 {
                    case "a":
                        v2 = v2
                    case "":
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                case 127:
                    v1 = append(v1, 352)
                }
            }
            _ = fi18
            _ = v19
        }
        v25 := GenS1{F0: 32768, F1: 0}
        _ = v17
        _ = v25
    }
    v16 += float64((0.1 / (float32((v16 - -2.25)) + 0.0)))
    if (65535 != 2) {
        if len(v5) > 0 {
            v5[0] = 2
        }
    }
    var v26 float64 = (((v16 * -2.25) * v16) - v16)
    if false {
        var v27 string = v14
        _ = v27
    } else {
        v28 := GenS2{F0: 2, F1: v14}
        v5 = append(v5, 4294967295)
        v16 = (v26 + v16)
        _ = v28
    }
    v6 = 65535
    v16 = ((v26 / 3.14159) * 1e10)
    v1 = append(v1, -1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
