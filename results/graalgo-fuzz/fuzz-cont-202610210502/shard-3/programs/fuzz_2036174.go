package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 uint) uint16 {
    p0 = (p0 - p0)
    var v4 uint16 = 517
    var v5 int8 = 126
    v4 = 386
    var v6 bool = (1 > -2147483648)
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (v4 + v4)
}

func main() {
    var v1 bool = (!(255 == 4294967295))
    v2 := make(map[string]string)
    v3 := make(map[uint32]uint32)
    v1 = true
    v1 = (!(v1 && v1))
    v7 := []int{32767, -2147483648, 4294967295}
    if v1 {
        v1 = (781 >= 277)
        var v8 float32 = 0.0
        v8 = (0.0 / 0.1)
        _ = v8
    } else {
        v1 = ((756 >= 4294967295) || false)
        var v9 int16 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        v1 = (v1 && false)
        if (1 != 118) {
            var v11 int8 = -54
            v11 = (^(v11))
            _ = v11
        }
        _ = fi10
    }
    fmt.Println(v1)
    if v1 {
        if len(v7) > 0 {
            v7[0] = -2147483648
        }
        v1 = v1
    } else {
        fmt.Println(v1)
    }
    fmt.Println(v1)
    for fi12 := 0; fi12 < 12; fi12++ {
        fmt.Println(v1, fi12)
        if (18446744073709551615 >= 944) {
            delete(v2, "\n\t")
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            for fi13 := 0; fi13 < 10; fi13++ {
                v1 = v1
                _ = fi13
            }
        } else {
            for fi14 := 0; fi14 < 4; fi14++ {
                v1 = v1
                v1 = true
                v15 := GenS1{F0: 827}
                _ = fi14
                _ = v15
            }
            v1 = (9223372036854775807 == 0)
        }
        _ = fi12
    }
    v2[""] = "\n\t"
    v16 := GenS2{F0: -448, F1: 42}
    var v17 int32 = -2147483648
    v18 := GenS1{F0: 0}
    if false {
        switch -9223372036854775808 {
        case 127:
            if v1 {
                for fi19 := 0; fi19 < 11; fi19++ {
                    switch fi19 {
                    case -268:
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v17 -= (-(v17))
                            v21 := GenS2{F0: -838, F1: 2}
                            _ = fi20
                            _ = v21
                        }
                    default:
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                    }
                    fmt.Println(fi19, fi19)
                    delete(v2, ("a" + ""))
                    _ = fi19
                }
            } else {
                v17 -= -2147483648
            }
        case -1:
            v22 := GenS2{F0: -622, F1: 42}
            _ = v22
        }
    } else {
        var v23 string = "hello"
        _ = v23
    }
    var v24 string = "hello"
    var v25 uint16 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v7
}
