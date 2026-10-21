package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 bool }
func genFn0() uint64 {
    v6 := GenS1{F0: -1, F1: 1, F2: true}
    var v7 int32 = 1
    v8 := GenS1{F0: -9223372036854775808, F1: 766, F2: true}
    _ = v6
    _ = v7
    _ = v8
    return 126
}

func main() {
    var v1 int32 = 2147483647
    var v2 int16 = -631
    var v3 uint16 = 693
    var v4 uint16 = 1
    var v5 int = 32768
    var v9 int = (v5 * (v5 + (v5 & (-(v5)))))
    v1 = (-(v1))
    var v10 int8 = -28
    var v11 int32 = (v1 * (-(v1)))
    fmt.Println(v9, v9, v4)
    switch 127 {
    case 128:
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := []int{0, (v5 + fi12)}
            _ = fi12
            _ = v13
        }
    case 127:
        fmt.Println(v10)
    default:
        if (!((v2 ^ (v2 ^ (v2 ^ v2))) != v2)) {
            for fi14 := 0; fi14 < 4; fi14++ {
                fmt.Println(v3)
                if (1000 <= (-(v2))) {
                    for fi15 := 0; fi15 < 7; fi15++ {
                        if true {
                            var v16 string = "a"
                            var v17 uint8 = 58
                            fmt.Println(v17, v11, fi14)
                            _ = v16
                            _ = v17
                        }
                        _ = fi15
                    }
                }
                switch "" {
                case "a":
                    if ((!(!false)) || (0 > uint(v4))) {
                        v3 = v4
                        fmt.Println(v4, v9)
                    } else {
                        v10 = (v10 % ((int8(v4)) | 1))
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v1 = int32(v10)
                            var v19 uint64 = 973
                            _ = fi18
                            _ = v19
                        }
                    }
                default:
                    v11 -= v1
                }
                _ = fi14
            }
            v5 -= (v9 & v5)
        }
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        v10 += v10
        switch "" {
        case "":
            v21 := GenS1{F0: (v5 | (v9 & (-(v5)))), F1: 2, F2: (!false)}
            _ = v21
        default:
            var v22 int32 = -1
            _ = v22
        }
        _ = fi20
    }
    v23 := []int{479}
    fmt.Println(v5, v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
