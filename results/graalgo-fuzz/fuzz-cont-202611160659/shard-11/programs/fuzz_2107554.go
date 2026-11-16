package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 uint32) uint32 {
    v4 := []int{}
    _ = len(v4)
    v4 = append(v4, 9223372036854775807)
    v5 := GenS1{F0: (false && false), F1: 1}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[uint32]bool)
    var v3 int32 = 1
    var v6 uint16 = (v1 ^ (v1 + v1))
    var v7 int64 = -445
    switch "foo bar" {
    case "a":
        v8 := GenS1{F0: (((((1e10 - 0.1) / 1.0) + (float32(1e10) - 0.1)) + ((1.5 / 3.14159) + (1.5 * 0.1))) >= (((((((0.1 - ((1e10 + -2.25) * 0.1)) / ((1.5 * (1e10 * ((1.5 - (0.1 + ((0.0 + 0.1) * (1.5 * 0.1)))) + -2.25))) + 1e10)) / (0.1 - 0.0)) + 1.5) - (0.0 - ((0.0 - 3.14159) / 0.1))) + 0.1) / -2.25)), F1: 100}
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: false, F1: -128}
        _ = v9
    default:
        v1 = v6
    }
    var v10 uint32 = genFn0(435, 871)
    v1 = v1
    if (!(864 == 457)) {
        v3 = v3
    }
    if true {
        v10 += (v10 / ((v10) | 1))
        if (42 <= 127) {
            switch -118 {
            case 834:
                var v11 float64 = (3.14159 + (-2.25 / 1.0))
                _ = v11
            }
        }
        v10 = (v10 - (v10 % (((v10 << ((1) & 63))) | 1)))
    }
    var v12 float64 = (3.14159 + 0.1)
    v2[(^(v10))] = (128 >= 128)
    v13 := GenS1{F0: true, F1: 0}
    var v14 int64 = (v7 | 382)
    if (false || true) {
        var v15 uint32 = (v10 ^ v10)
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if true {
            v3 = v3
            v17 := GenS1{F0: true, F1: 1}
            if (83 >= 1) {
                v2[genFn0(2, (v10 ^ v10))] = (173 <= 166)
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 int = (fi16 | (^(fi18)))
                    switch ("foo bar" + "hello") {
                    case "\n\t":
                        v20 := GenS2{F0: 76}
                        _ = v20
                    case "hello":
                        v10 -= v10
                    case "":
                        var v21 int64 = (v7 << ((1) & 63))
                        _ = v21
                    }
                    var v22 int64 = 9223372036854775807
                    delete(v2, 1)
                    v1 -= (v1 * (^(v1)))
                    _ = fi18
                    _ = v19
                    _ = v22
                }
                v23 := GenS3{F0: false, F1: 422, F2: 0}
                _ = v23
            } else {
                v12 = (v12 * (v12 * (v12 - v12)))
                for fi24 := 0; fi24 < 15; fi24++ {
                    if true {
                        var v25 uint = 42
                        v26 := []int{fi24, fi16, 2147483647, fi16}
                        _ = v25
                        _ = v26
                    }
                    _ = fi24
                }
                v27 := []int{32767, 402, fi16}
                _ = v27
            }
            _ = v17
        }
        v28 := GenS2{F0: 100}
        _ = fi16
        _ = v28
    }
    delete(v2, (v10 >> ((4294967295) & 63)))
    for fi29 := 0; fi29 < 13; fi29++ {
        for fi30 := 0; fi30 < 9; fi30++ {
            var v31 int8 = 127
            v1 += 255
            if ((!false) || (!(float32((0.1 - 1e10)) <= 3.14159))) {
                v3 -= v3
            } else {
                v10 -= (v10 & (v10 + (v10 & (v10 & (v10 >> ((18446744073709551615) & 63))))))
                if false {
                    v32 := []int{fi30, 2147483647}
                    if (2 == 706) {
                        v2[(v10 >> ((717) & 63))] = ((^(v1)) == (v1 / ((uint16(v10)) | 1)))
                        fmt.Println(v6, fi29)
                        fmt.Println(v1)
                    } else {
                        fmt.Println(fi30)
                        v33 := GenS4{F0: true}
                        var v34 string = "foo bar"
                        _ = v33
                        _ = v34
                    }
                    var v35 int32 = v3
                    _ = v32
                    _ = v35
                } else {
                    v7 -= (^(v7))
                    v7 = v14
                }
            }
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
