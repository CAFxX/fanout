package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool, p1 int8) uint {
    for fi4 := 0; fi4 < 10; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: false, F1: (!false), F2: 2}
            var v7 uint = 910
            fmt.Println(fi4, p1, v7)
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := []int{}
        _ = len(v8)
        _ = fi4
        _ = v8
    }
    p1 += p1
    p0 = false
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS2{F0: 42, F1: p1}
        v11 := []int{(fi9 >> ((615) & 63)), (fi9 - fi9)}
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    return 912
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint16 = 42
    v3 := []uint32{4294967295, 0, 2147483648}
    var v12 bool = ((v2 != (^(v2))) && true)
    var v13 uint32 = 768
    v14 := GenS2{F0: (v13 - (v13 >> ((644) & 63))), F1: 127}
    v15 := GenS3{F0: 83}
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 string = "a"
        _ = fi16
        _ = v17
    }
    var v18 int16 = 42
    v18 = 0
    if v12 {
        v2 += v2
    }
    v19 := GenS1{F0: (!true), F1: v12, F2: 0}
    v2 += v2
    v1 = (^(v1))
    v18 -= 1000
    v18 = (v18 - v18)
    v13 = (v13 - v13)
    if v12 {
        v2 = (v2 ^ 1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
