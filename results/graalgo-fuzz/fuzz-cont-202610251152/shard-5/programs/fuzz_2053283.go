package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
func genFn0() int8 {
    var v4 int = 32768
    fmt.Println(v4)
    _ = v4
    return 127
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 bool = ((-1 == -275) || (!(!(0.1 == 3.14159))))
    var v3 uint32 = 980
    var v5 int8 = genFn0()
    v1 = 521
    fmt.Println(v2, v1, v3)
    switch 278 {
    case 2147483647:
        v6 := []int{32768, -188, 32767, 996}
        _ = v6
    case -16:
        v7 := []int{0}
        _ = v7
    case 32768:
        if v2 {
            var v8 bool = (985 >= 708)
            _ = v8
        } else {
            if v2 {
                switch 0 {
                case -799:
                    v9 := []int{903, 0, int(v3), 32768}
                    _ = v9
                default:
                    if v2 {
                        v2 = (!true)
                        v5 -= (v5 - (v5 & (v5 * (v5 ^ (-(v5))))))
                    } else {
                        for fi10 := 0; fi10 < 11; fi10++ {
                            if (9223372036854775807 < 9223372036854775807) {
                                var v11 int64 = 371
                                v2 = (false || (v2 || (367 > 2147483648)))
                                _ = v11
                            }
                            v3 = v3
                            _ = fi10
                        }
                        var v12 uint8 = 2
                        switch -9223372036854775808 {
                        case 1:
                            var v13 int32 = 863
                            _ = v13
                        }
                        _ = v12
                    }
                }
            }
            v14 := []int{-795, 0, 2147483647, 2147483647}
            _ = v14
        }
    }
    v5 += v5
    v15 := []int{0, 4294967295, 42}
    v1 += (v1 ^ (v1 - v1))
    var v16 int64 = 42
    var v17 uint8 = 0
    for fi18 := 0; fi18 < 14; fi18++ {
        if true {
            for fi19 := 0; fi19 < 18; fi19++ {
                var v20 float64 = (-2.25 + 1.5)
                _ = fi19
                _ = v20
            }
            for fi21 := 0; fi21 < 10; fi21++ {
                v2 = true
                _ = fi21
            }
        } else {
            for fi22 := 0; fi22 < 7; fi22++ {
                if (!((^(fi18)) != (fi22 >> ((78) & 63)))) {
                    var v23 bool = v2
                    v24 := GenS1{F0: v17, F1: 2, F2: (v3 & v3)}
                    v25 := GenS1{F0: (v17 << ((759) & 63)), F1: 0, F2: uint32(v17)}
                    _ = v23
                    _ = v24
                    _ = v25
                }
                _ = fi22
            }
            if (32767 >= 1000) {
                v26 := GenS2{F0: v3, F1: 307}
                _ = v26
            } else {
                v3 -= (v3 + (v3 / (((^(v3))) | 1)))
            }
            fmt.Println(v5, v1)
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
}
