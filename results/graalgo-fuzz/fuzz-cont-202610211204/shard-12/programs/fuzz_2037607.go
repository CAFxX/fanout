package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint64 }
func genFn0() uint {
    if ((((42 >= 42) && false) || (!(-32768 <= 1000))) && (18446744073709551615 < 18446744073709551615)) {
        var v5 uint16 = 175
        _ = v5
    }
    var v6 int16 = 0
    fmt.Println(v6)
    _ = v6
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 uint16 = 65535
    v4 -= 2
    var v7 float32 = ((-2.25 / (float32(((1e10 - float64(1.5)) * (-2.25 + float64(1e10)))) + 1e10)) / (3.14159 / (1e10 + float32(((0.1 + (3.14159 / 1e10)) - ((3.14159 / (0.0 + (0.1 * 3.14159))) - 1e10))))))
    if (42 >= 0) {
        switch 4294967295 {
        case 2147483647:
            v8 := []int{}
            _ = v8
        default:
            v9 := GenS1{F0: 42, F1: (v1 % ((42) | 1)), F2: 42}
            _ = v9
        }
        v10 := []int{-625, 4294967295, 175, 1}
        _ = v10
    } else {
        v7 = float32(((-2.25 + 0.1) / 1e10))
        fmt.Println(v1, v1, v1)
        var v11 int16 = (v1 + v1)
        _ = v11
    }
    v2 = (v2 << ((genFn0()) & 63))
    v7 = float32(3.14159)
    v12 := GenS1{F0: 1, F1: v1, F2: 4294967295}
    v3[0] = false
    var v13 float32 = v7
    switch "a" {
    case "a":
        fmt.Println(v1, v1, v13)
    case "hello":
        var v14 uint16 = 173
        _ = v14
    case "\n\t":
        if (!(128 != 127)) {
            for fi15 := 0; fi15 < 12; fi15++ {
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := []int{fi16}
                    var v18 bool = false
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                fmt.Println(fi15)
                _ = fi15
            }
        }
    }
    fmt.Println(v1, v7)
    switch 32768 {
    case -1:
        switch -377 {
        case -9223372036854775808:
            for fi19 := 0; fi19 < 3; fi19++ {
                v4 += v4
                switch (fi19 << ((278) & 63)) {
                case 128:
                    delete(v3, 715)
                case 32768:
                    fmt.Println(v1, v2, v13)
                default:
                    for fi20 := 0; fi20 < 17; fi20++ {
                        if false {
                            var v21 int = fi20
                            switch v21 {
                            case 4294967295:
                                switch 128 {
                                case -2147483648:
                                    v22 := GenS1{F0: 2147483647, F1: (v1 << ((255) & 63)), F2: 459}
                                    _ = v22
                                case -1:
                                    fmt.Println(v21)
                                default:
                                    var v23 int16 = v1
                                    _ = v23
                                }
                            case 32768:
                                if (!(659 >= -577)) {
                                    fmt.Println(v21, v4, v7)
                                } else {
                                    v1 += v1
                                    var v24 int64 = v2
                                    var v25 float64 = float64(3.14159)
                                    _ = v24
                                    _ = v25
                                }
                            case 128:
                                var v26 bool = (!(132 >= 18446744073709551615))
                                _ = v26
                            default:
                                v2 = v2
                            }
                            _ = v21
                        }
                        fmt.Println(v7)
                        fmt.Println(v4, v7, v4)
                        _ = fi20
                    }
                }
                _ = fi19
            }
        case 127:
            var v27 int64 = (v2 / ((553) | 1))
            _ = v27
        default:
            fmt.Println(v13)
        }
    default:
        v2 = (v2 % ((v2) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
