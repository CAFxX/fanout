package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
func genFn0() string {
    var v4 bool = true
    for fi5 := 0; fi5 < 4; fi5++ {
        v4 = true
        _ = fi5
    }
    v4 = (2147483647 >= -997)
    _ = v4
    return "hello"
}
func genFn1(p0 int16) int32 {
    switch -8 {
    case -451:
        fmt.Println(p0)
    case 42:
        p0 += 517
    }
    switch 258 {
    case 9223372036854775807:
        fmt.Println(p0)
    case 95:
        var v6 bool = false
        _ = v6
    }
    _ = p0
    return 845
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v2[921] = 2147483647
    v1 = (v1 + v1)
    var v7 float32 = 1e10
    if (genFn1(-1) < 1) {
        v8 := GenS1{F0: 255, F1: -9223372036854775808}
        var v9 float32 = 0.0
        _ = v8
        _ = v9
    } else {
        var v10 int64 = -1
        v7 += (v7 / float32(float64(((((v7 - v7) * float32(1.5)) + v7) + 1.5))))
        _ = v10
    }
    var v11 bool = (!true)
    fmt.Println(v1, v11, v1)
    var v12 uint8 = 110
    _ = len(v3)
    for fi13 := 0; fi13 < 19; fi13++ {
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := []int{(fi14 << ((65535) & 63)), (fi13 | fi13)}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: 1, F1: 311, F2: (v12 >> ((4294967295) & 63))}
            switch (fi16 << ((18446744073709551615) & 63)) {
            case 32767:
                v18 := GenS1{F0: 1, F1: 0}
                _ = v18
            }
            v12 = (v12 ^ (v12 ^ v12))
            _ = fi16
            _ = v17
        }
        _ = fi13
    }
    if v11 {
        v19 := []int{-9223372036854775808, -779}
        v1 = ("" + "foo bar")
        fmt.Println(v11, v12)
        _ = v19
    } else {
        v7 = v7
        var v20 string = "a"
        fmt.Println(v1)
        _ = v20
    }
    v21 := v3[-2147483648]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v7
}
