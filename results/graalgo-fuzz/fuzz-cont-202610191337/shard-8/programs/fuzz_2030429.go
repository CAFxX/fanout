package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint16 }
type GenS2 struct { F0 uint; F1 uint; F2 uint64 }
func genFn0(p0 int64, p1 int16) bool {
    for fi4 := 0; fi4 < 10; fi4++ {
        p1 += -396
        var v5 uint64 = 4294967295
        _ = fi4
        _ = v5
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 32767, F1: fi6, F2: 524}
        _ = fi6
        _ = v7
    }
    p0 += p0
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 uint16, p1 int64, p2 int32) int8 {
    v8 := []int{0, 2147483647}
    p0 -= 521
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 uint16 = 32767
    var v9 bool = false
    var v10 uint32 = 2147483647
    var v11 uint8 = 128
    if ("a" > "foo bar") {
        v3 = v3
        v12 := []int{0, 80}
        _ = v12
    } else {
        if false {
            v13 := GenS2{F0: 42, F1: 549, F2: uint64(v11)}
            _ = v13
        } else {
            v9 = v9
        }
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        v15 := GenS1{F0: (fi14 + (fi14 % ((fi14) | 1))), F1: fi14, F2: (v3 ^ 32767)}
        v9 = ("" <= ("a" + (("" + "hello") + "")))
        _ = fi14
        _ = v15
    }
    v16 := GenS2{F0: 655, F1: 65535, F2: 18446744073709551615}
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 int64 = -855
        v1["hello"] = (fi17 - (fi17 & -2147483648))
        v19 := GenS2{F0: 77, F1: 4294967295, F2: 0}
        _ = fi17
        _ = v18
        _ = v19
    }
    v9 = (!(!v9))
    v11 += v11
    v3 -= (v3 >> ((2) & 63))
    fmt.Println(v11)
    v11 += v11
    for fi20 := 0; fi20 < 10; fi20++ {
        v9 = v9
        _ = fi20
    }
    var v21 int64 = 42
    v11 += v11
    v21 = (v21 % ((v21) | 1))
    v22 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
