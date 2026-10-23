package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int32, p1 int32) uint {
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 float32 = 3.14159
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: -1}
    fmt.Println(p0, p0)
    for fi9 := 0; fi9 < 19; fi9++ {
        switch -796 {
        case 303:
            v10 := GenS2{F0: 18446744073709551615}
            _ = v10
        case -9223372036854775808:
            switch (fi9 - (fi9 ^ -254)) {
            case 0:
                p0 = (p0 | (-(p1)))
            }
        }
        _ = fi9
    }
    v11 := GenS1{F0: 0, F1: 42}
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 4294967295
}
func genFn1(p0 uint8) uint16 {
    if true {
        var v12 uint = 4294967295
        v13 := []int{-607, 127, -248}
        _ = v12
        _ = v13
    }
    if true {
        p0 -= (p0 & 2)
        var v14 string = "\n\t"
        _ = v14
    }
    v15 := GenS2{F0: 18446744073709551615}
    fmt.Println(p0)
    var v16 uint8 = (p0 << ((596) & 63))
    _ = p0
    _ = v15
    _ = v16
    return 668
}

func main() {
    var v1 float32 = 0.0
    var v2 int16 = 32767
    v3 := []int{1, -9223372036854775808, 57}
    var v4 float32 = v1
    var v5 uint32 = 65535
    switch -9223372036854775808 {
    case -1:
        if true {
            v1 = 0.0
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
        }
    case -146:
        v3 = append(v3, 42)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 uint64 = 34
    if len(v3) > 0 {
        v3[0] = 0
    }
    var v18 int8 = -1
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(fi19, v18)
        _ = len(v3)
        _ = fi19
    }
    fmt.Println(v2)
    v5 = (^(v5))
    var v20 uint = genFn0(-629, -752)
    var v21 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
