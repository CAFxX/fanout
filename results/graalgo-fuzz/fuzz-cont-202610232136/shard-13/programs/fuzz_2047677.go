package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }
type GenS2 struct { F0 string; F1 uint64; F2 bool }

func main() {
    var v1 uint64 = 42
    v2 := make(map[int]int)
    v3 := make(map[int]uint32)
    var v4 int32 = -489
    var v5 uint32 = 4294967295
    if true {
        if ("foo bar" <= "") {
            v5 = (^(v5))
            var v6 uint = 4294967295
            _ = v6
        } else {
            for fi7 := 0; fi7 < 3; fi7++ {
                v2[(fi7 % ((fi7) | 1))] = (fi7 ^ (fi7 * (-(fi7))))
                var v8 int64 = -9223372036854775808
                v9 := GenS1{F0: 922, F1: (v8 & v8), F2: 0}
                _ = fi7
                _ = v8
                _ = v9
            }
            if (!true) {
                var v10 int = 128
                _ = v10
            } else {
                for fi11 := 0; fi11 < 11; fi11++ {
                    if (841 == 0) {
                        switch fi11 {
                        case 42:
                            var v12 int32 = v4
                            _ = v12
                        case 32768:
                            var v13 float64 = (3.14159 + (((-2.25 * 1e10) + 3.14159) - ((((1e10 * (-2.25 - 0.1)) * 0.1) + (1e10 + (1.5 + 1.5))) * -2.25)))
                            _ = v13
                        case -9223372036854775808:
                            fmt.Println(v5, v1)
                        default:
                            if false {
                                for fi14 := 0; fi14 < 13; fi14++ {
                                    if true {
                                        v4 = -273
                                    }
                                    var v15 int16 = 1000
                                    v1 += 18446744073709551615
                                    _ = fi14
                                    _ = v15
                                }
                                var v16 int32 = 1
                                _ = v16
                            }
                        }
                        v5 = (^(v5))
                        v1 = v1
                    }
                    switch fi11 {
                    case -40:
                        v2[fi11] = -335
                    default:
                        var v17 uint32 = (v5 % (((v5 & (v5 & (v5 << ((30) & 63))))) | 1))
                        _ = v17
                    }
                    _ = fi11
                }
            }
            v18 := GenS1{F0: 439, F1: -128, F2: -568}
            _ = v18
        }
        if false {
            v1 = v1
            var v19 uint16 = 255
            _ = v19
        }
        if (761 >= v4) {
            fmt.Println(v4, v5, v5)
            var v20 uint16 = 42
            _ = v20
        } else {
            v21 := GenS1{F0: 251, F1: 215, F2: 1}
            _ = v21
        }
    }
    v22 := []int{42}
    var v23 bool = true
    var v24 bool = v23
    var v25 int64 = 9223372036854775807
    var v26 int = -56
    v27 := GenS2{F0: "", F1: v1, F2: ((int32(v5) < v4) && (!v24))}
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
