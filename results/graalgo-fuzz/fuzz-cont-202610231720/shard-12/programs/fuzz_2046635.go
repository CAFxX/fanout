package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 bool; F2 bool }

func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 18446744073709551615
    var v3 uint16 = 305
    var v4 float32 = v1
    var v5 string = ""
    switch 0 {
    case 966:
        v6 := GenS1{F0: 275, F1: 18446744073709551615}
        _ = v6
    case 889:
        v7 := []int{}
        _ = v7
    case 0:
        fmt.Println(v2, v1)
    default:
        v8 := GenS1{F0: 986, F1: 454}
        _ = v8
    }
    v9 := GenS2{F0: 0, F1: v3, F2: 4294967295}
    fmt.Println(v4, v1, v3)
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: fi10, F1: 65535, F2: 207}
        switch v5 {
        case "a":
            v2 = (v2 * v2)
        case "foo bar":
            switch v5 {
            case "hello":
                v5 = v5
            case "\n\t":
                v12 := GenS3{F0: v2}
                _ = v12
            case "a":
                v4 -= v4
            }
        }
        _ = fi10
        _ = v11
    }
    v4 += 0.1
    fmt.Println(v4, v1)
    v3 = v3
    fmt.Println(v5)
    var v13 bool = (((v5 + v5) + "") >= v5)
    var v14 string = ("" + v5)
    var v15 uint = 385
    for fi16 := 0; fi16 < 8; fi16++ {
        switch fi16 {
        case 127:
            var v17 uint32 = 42
            _ = v17
        }
        v5 = v14
        _ = fi16
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float32 = (v1 + ((1e10 + (float32((1.5 / 1.0)) / (v4 - v1))) - 1.5))
        _ = fi18
        _ = v19
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        switch fi20 {
        case 4294967295:
            var v21 int64 = 42
            _ = v21
        case 128:
            switch v5 {
            case "\n\t":
                for fi22 := 0; fi22 < 16; fi22++ {
                    switch -2147483648 {
                    case 128:
                        fmt.Println(fi22, v15, v14)
                    case -599:
                        if (127 != 14) {
                            switch fi20 {
                            case 4294967295:
                                v23 := []int{}
                                _ = v23
                            case -1:
                                var v24 float64 = 3.14159
                                if true {
                                    v4 = (float32(v24) + ((v1 / (v4 * 0.0)) / float32(((-2.25 - v24) - v24))))
                                } else {
                                    v25 := []int{fi22, (fi22 | (^(fi22)))}
                                    v15 -= v15
                                    v25 = append(v25, -607)
                                    _ = v25
                                }
                                fmt.Println(v1, fi22, fi20)
                                v26 := []int{817, (fi20 % ((fi22) | 1)), -69, fi22}
                                _ = v24
                                _ = v26
                            default:
                                v2 -= v2
                            }
                        }
                    default:
                        switch "hello" {
                        case "foo bar":
                            v27 := GenS2{F0: (^(fi22)), F1: (^(v3)), F2: 42}
                            _ = v27
                        default:
                            var v28 float64 = -2.25
                            _ = v28
                        }
                    }
                    fmt.Println(v1, v5, v13)
                    _ = fi22
                }
            case "hello":
                var v29 int16 = 0
                _ = v29
            default:
                v30 := GenS2{F0: (fi20 << ((272) & 63)), F1: 32767, F2: 915}
                _ = v30
            }
        case 190:
            if (9223372036854775807 <= -70) {
                v31 := GenS3{F0: v2}
                _ = v31
            }
        default:
            v32 := GenS4{F0: (v2 | (v2 | 4294967295)), F1: v13, F2: (89 > -128)}
            _ = v32
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
