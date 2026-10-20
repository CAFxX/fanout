package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 bool }
type GenS2 struct { F0 uint16; F1 int16; F2 int }
func genFn0(p0 uint, p1 int16, p2 uint8) int {
    if true {
        switch 1 {
        case -113:
            v4 := []int{}
            _ = v4
        case 9223372036854775807:
            var v5 uint32 = uint32(p2)
            _ = v5
        }
        var v6 int8 = int8(p2)
        _ = v6
    }
    if false {
        var v7 bool = ((^(p0)) != (p0 | p0))
        _ = v7
    }
    fmt.Println(p1, p0, p0)
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return 32767
}
func genFn1() int32 {
    var v8 int64 = -9223372036854775808
    fmt.Println(v8)
    v8 += (v8 << ((18446744073709551615) & 63))
    v8 -= (v8 * v8)
    var v9 int32 = -1
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 uint16 = 780
    var v2 uint8 = 197
    v3 := make(map[uint32]bool)
    v2 = (v2 | v2)
    var v10 int32 = -530
    v10 = v10
    v10 -= (v10 * v10)
    if (3.14159 >= 0.0) {
        fmt.Println(v1)
        v1 -= (^(v1))
    } else {
        v2 -= v2
    }
    v11 := GenS1{F0: int(v10), F1: 436, F2: ((!true) && (!false))}
    _ = len(v3)
    for fi12 := 0; fi12 < 2; fi12++ {
        v13 := []int{genFn0(461, 0, (v2 / (((v2 - (v2 >> ((797) & 63)))) | 1))), (fi12 % ((fi12) | 1))}
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int16 = -32768
            if false {
                v1 = v1
                var v16 float64 = 3.14159
                _ = v16
            } else {
                fmt.Println(v10, v10)
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    switch genFn0(4294967295, 42, (v2 & v2)) {
    case -2147483648:
        v17 := GenS1{F0: 32767, F1: 0, F2: (true || false)}
        _ = v17
    case -843:
        var v18 string = "a"
        _ = v18
    case -9223372036854775808:
        for fi19 := 0; fi19 < 5; fi19++ {
            v10 = v10
            switch -2147483648 {
            case -553:
                fmt.Println(fi19, v1)
            case -520:
                fmt.Println(fi19, v2)
            default:
                v10 -= int32(v2)
            }
            v2 = (^(v2))
            _ = fi19
        }
    default:
        v20 := GenS2{F0: v1, F1: 1000, F2: -940}
        _ = v20
    }
    v2 = (^(v2))
    for fi21 := 0; fi21 < 20; fi21++ {
        var v22 int64 = 42
        v23 := []int{fi21}
        for fi24 := 0; fi24 < 9; fi24++ {
            v2 += v2
            if ((v22 >> ((18446744073709551615) & 63)) < (v22 / ((894) | 1))) {
                if (42 < 1) {
                    switch (^(fi24)) {
                    case 216:
                        v2 = 42
                    default:
                        v3[42] = (421 == 741)
                    }
                    if len(v23) > 0 {
                        v23[0] = -2147483648
                    }
                } else {
                    fmt.Println(v1)
                    var v25 uint = 4294967295
                    _ = v25
                }
            } else {
                _ = len(v3)
                var v26 float32 = (0.1 * 3.14159)
                v2 = 127
                _ = v26
            }
            _ = fi24
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    var v27 int64 = 269
    delete(v3, 4294967295)
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v3
}
