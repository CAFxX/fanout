package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint32 }
type GenS2 struct { F0 int }
func genFn0(p0 int, p1 int64) uint {
    v4 := GenS1{F0: 4294967295, F1: ("" + "a"), F2: 975}
    switch ("a" + "") {
    case "\n\t":
        if true {
            for fi5 := 0; fi5 < 3; fi5++ {
                fmt.Println(fi5, p0, fi5)
                p1 = p1
                var v6 string = "hello"
                _ = fi5
                _ = v6
            }
        }
    }
    var v7 uint64 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        p0 += fi8
        v9 := GenS2{F0: fi8}
        for fi10 := 0; fi10 < 6; fi10++ {
            v11 := []int{-65, (fi8 - (p0 & fi8))}
            _ = fi10
            _ = v11
        }
        _ = fi8
        _ = v9
    }
    v12 := []int{(p0 + p0), 32768}
    _ = p0
    _ = p1
    _ = v12
    _ = v4
    _ = v7
    return 978
}

func main() {
    var v1 bool = true
    var v2 float32 = -2.25
    v3 := []int{}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 int16 = 32767
        var v15 int = fi13
        _ = fi13
        _ = v14
        _ = v15
    }
    v2 -= ((0.1 + (0.1 - 3.14159)) + v2)
    if len(v3) > 0 {
        v3[0] = -9223372036854775808
    }
    v2 = v2
    fmt.Println(v1)
    v16 := GenS2{F0: 738}
    v1 = v1
    for fi17 := 0; fi17 < 2; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
    }
    if len(v3) > 0 {
        v3[0] = 691
    }
    v18 := GenS1{F0: 256, F1: "", F2: 2147483647}
    if true {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            v2 -= v2
            v2 += v2
        }
        var v19 uint32 = 42
        _ = v19
    }
    var v20 uint8 = 93
    var v21 float64 = (0.0 * 3.14159)
    fmt.Println(v2, v1)
    fmt.Println(v2, v2)
    v20 += 131
    var v22 uint32 = 2147483648
    var v23 uint8 = (v20 + (v20 / ((uint8(v22)) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
