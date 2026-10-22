package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 bool; F2 int }
type GenS3 struct { F0 uint; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint = 67
    var v2 int32 = 42
    var v3 uint64 = 2
    fmt.Println(v1, v3, v2)
    if true {
        v2 -= (v2 + (v2 | (^(v2))))
        var v4 int64 = 42
        _ = v4
    }
    var v5 float64 = -2.25
    fmt.Println(v1, v2)
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 20; fi7++ {
            v8 := []int{}
            fmt.Println(v3, v1)
            fmt.Println(v5, v2)
            if true {
                switch "" {
                case "":
                    v5 = float64((3.14159 * 1.5))
                default:
                    v8 = append(v8, (fi7 * 4294967295))
                }
                v8 = append(v8, int(v3))
            }
            if len(v8) > 0 {
                v8[0] = 32767
            }
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: 0}
        if (!((true || ((float64(3.14159) <= v5) && (0 > 42))) && (!(((v5 + v5) * v5) > 1e10)))) {
            switch -207 {
            case 42:
                for fi10 := 0; fi10 < 18; fi10++ {
                    fmt.Println(fi10, v2)
                    if false {
                        v11 := GenS2{F0: 9223372036854775807, F1: true, F2: 338}
                        var v12 int64 = -9223372036854775808
                        _ = v11
                        _ = v12
                    } else {
                        if true {
                            var v13 uint16 = 0
                            v14 := []int{fi6, 32, (fi10 - int(v2)), 976}
                            if false {
                                v15 := GenS3{F0: (^(v1)), F1: 42, F2: 100}
                                _ = v15
                            }
                            _ = v13
                            _ = v14
                        } else {
                            v1 = v1
                            v1 += (v1 % ((v1) | 1))
                            fmt.Println(fi6, v5)
                        }
                        var v16 int16 = 769
                        var v17 bool = false
                        _ = v16
                        _ = v17
                    }
                    _ = fi10
                }
            }
            var v18 int64 = -9223372036854775808
            _ = v18
        } else {
            v2 += -226
            fmt.Println(fi6)
        }
        _ = fi6
        _ = v9
    }
    switch 127 {
    case 936:
        fmt.Println(v2, v3, v1)
    case 779:
        v19 := []int{}
        _ = v19
    case -9223372036854775808:
        v20 := GenS1{F0: 2}
        _ = v20
    default:
        v21 := GenS4{F0: 18446744073709551615, F1: 0}
        _ = v21
    }
    fmt.Println(v5, v3)
    v5 -= (v5 / 1.0)
    var v22 int64 = -157
    var v23 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v5
}
