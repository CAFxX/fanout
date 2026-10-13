package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64 }
func genFn0(p0 string) int16 {
    p0 = p0
    var v7 uint16 = 2
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 uint) uint32 {
    for fi8 := 0; fi8 < 17; fi8++ {
        p0 = 1
        switch fi8 {
        case -482:
            v9 := GenS1{F0: 17, F1: 42}
            _ = v9
        case -121:
            fmt.Println(fi8)
        default:
            p0 += (p0 >> ((0) & 63))
        }
        var v10 int32 = -1
        _ = fi8
        _ = v10
    }
    v11 := []int{0}
    _ = p0
    _ = v11
    return 42
}

func main() {
    var v1 uint32 = 42
    var v2 int16 = -796
    v3 := []int{-9223372036854775808, 128, -284}
    var v4 uint = 4294967295
    var v5 bool = (121 != 42)
    var v6 int32 = -1
    v12 := GenS1{F0: 18446744073709551615, F1: 1}
    fmt.Println(v2, v1)
    fmt.Println(v4, v6)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint32 = genFn1(v4)
    v13 += (^(v1))
    for fi14 := 0; fi14 < 10; fi14++ {
        if v5 {
            if (v5 || (-9223372036854775808 >= 9223372036854775807)) {
                var v15 int8 = int8(v2)
                _ = v15
            } else {
                for fi16 := 0; fi16 < 4; fi16++ {
                    if false {
                        if len(v3) > 0 {
                            v3[0] = 128
                        }
                        v17 := GenS2{F0: 1, F1: 42}
                        _ = v17
                    }
                    fmt.Println(v5)
                    v4 += v4
                    _ = fi16
                }
                if v5 {
                    v1 += genFn1(65535)
                    v4 = v4
                    var v18 int = fi14
                    _ = v18
                }
            }
            v3 = append(v3, fi14)
            for fi19 := 0; fi19 < 20; fi19++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi19
            }
        } else {
            var v20 int = (fi14 % ((0) | 1))
            v3 = append(v3, fi14)
            var v21 float32 = -2.25
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    if len(v3) > 0 {
        v3[0] = 32768
    }
    var v22 uint64 = 360
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
