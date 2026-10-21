package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
func genFn0() int64 {
    var v7 uint16 = 65535
    var v8 bool = true
    v7 += (v7 / ((v7) | 1))
    if false {
        v9 := []int{0}
        for fi10 := 0; fi10 < 19; fi10++ {
            fmt.Println(v7)
            var v11 bool = v8
            _ = fi10
            _ = v11
        }
        _ = v9
    }
    fmt.Println(v7, v8)
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() bool {
    if ((true && (!true)) && (true || (1000 < 32767))) {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint8 = 247
            v14 := GenS1{F0: 51, F1: 68, F2: 42}
            _ = fi12
            _ = v13
            _ = v14
        }
    } else {
        var v15 string = "\n\t"
        v15 = (((("foo bar" + v15) + v15) + v15) + "a")
        _ = v15
    }
    v16 := GenS2{F0: -647}
    v17 := []int{}
    var v18 uint = 65535
    v19 := GenS1{F0: 105, F1: 255, F2: 1}
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return (-1 >= -9223372036854775808)
}

func main() {
    var v1 int32 = -2147483648
    var v2 int16 = 653
    v3 := []int{32768, 9223372036854775807, 1}
    var v4 uint32 = 0
    var v5 uint = uint(v4)
    var v6 int = -2147483648
    for fi20 := 0; fi20 < 6; fi20++ {
        v2 = (v2 ^ (-(v2)))
        var v21 int8 = int8(v4)
        if false {
            if false {
                v2 -= (^(v2))
            }
            v6 = v6
            var v22 uint64 = 643
            _ = v22
        } else {
            fmt.Println(v21)
        }
        _ = fi20
        _ = v21
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        switch v6 {
        case 1:
            v3 = append(v3, v6)
        case 4294967295:
            switch (fi23 | fi23) {
            case 533:
                v4 = v4
            default:
                fmt.Println(v5)
            }
        case -9223372036854775808:
            if true {
                v24 := GenS2{F0: (v6 & v6)}
                fmt.Println(v4)
                _ = v24
            } else {
                for fi25 := 0; fi25 < 7; fi25++ {
                    v3 = append(v3, v6)
                    v4 -= v4
                    _ = fi25
                }
            }
        }
        if len(v3) > 0 {
            v3[0] = -2147483648
        }
        v1 = v1
        _ = fi23
    }
    var v26 float32 = 3.14159
    fmt.Println(v2)
    v27 := GenS2{F0: v6}
    var v28 bool = true
    var v29 uint8 = 2
    if len(v3) > 0 {
        v3[0] = v6
    }
    var v30 int16 = (v2 * 874)
    v31 := GenS3{F0: -2147483648}
    _ = len(v3)
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
