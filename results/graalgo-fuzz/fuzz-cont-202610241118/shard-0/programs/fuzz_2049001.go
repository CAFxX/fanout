package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int) uint {
    p0 += p0
    for fi4 := 0; fi4 < 3; fi4++ {
        fmt.Println(fi4, p0)
        _ = fi4
    }
    var v5 uint8 = 255
    var v6 int16 = 42
    _ = p0
    _ = v5
    _ = v6
    return 4294967295
}
func genFn1(p0 int8) uint64 {
    var v7 int16 = -32768
    v8 := []int{-859}
    for fi9 := 0; fi9 < 6; fi9++ {
        if (fi9 > 2147483647) {
            p0 += 98
            var v10 bool = (-989 == 1)
            _ = v10
        } else {
            v11 := GenS1{F0: 9223372036854775807}
            _ = v11
        }
        var v12 uint64 = 1
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    v1 := []int{-1}
    var v2 int64 = 9223372036854775807
    v3 := []uint32{531, 997}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v1)
    if true {
        v13 := GenS1{F0: -554}
        _ = v13
    }
    v2 = v2
    fmt.Println(v2)
    v2 = 42
    var v14 float32 = (1.5 * 1e10)
    v2 = v2
    v15 := GenS1{F0: -187}
    if true {
        v14 += (v14 + v14)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v16 bool = true
    v16 = v16
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
