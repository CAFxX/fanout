package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0() string {
    var v7 uint8 = 255
    v7 = 0
    _ = v7
    return ""
}
func genFn1(p0 int64, p1 string) uint32 {
    fmt.Println(p1)
    if true {
        var v8 float64 = (3.14159 - 0.0)
        _ = v8
    }
    _ = p0
    _ = p1
    return 200
}

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (true || (-2.25 == 1.5))
    var v3 int32 = 42
    var v4 uint64 = 2
    var v5 uint32 = 353
    var v6 int16 = 166
    for fi9 := 0; fi9 < 4; fi9++ {
        v6 -= 1000
        v10 := GenS1{F0: uint8(v4), F1: 65535, F2: 2}
        v6 = v6
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        _ = len(v1)
        v12 := []int{(fi11 + (fi11 | (^(fi11)))), fi11}
        v5 -= v5
        _ = fi11
        _ = v12
    }
    v3 -= (v3 << ((1) & 63))
    var v13 bool = ("\n\t" <= ("a" + "\n\t"))
    v14 := GenS1{F0: 0, F1: 441, F2: 0}
    if (int(v6) == 4294967295) {
        var v15 int16 = (v6 >> ((18446744073709551615) & 63))
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v6, fi16, fi16)
        _ = fi16
    }
    fmt.Println(v13, v2, v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
