package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 int; F2 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint, p1 uint32) uint32 {
    var v4 uint32 = p1
    v4 = (v4 << ((2) & 63))
    v4 = 4294967295
    if false {
        v5 := []int{128}
        fmt.Println(p0, v4)
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    }
    p0 -= 489
    _ = p0
    _ = p1
    _ = v4
    return (p1 & (p1 ^ 961))
}
func genFn1() int16 {
    v6 := GenS1{F0: 2147483647}
    v7 := []int{}
    var v8 uint = 1
    v9 := GenS2{F0: "hello", F1: 1000}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    if false {
        v1 += (v1 / v1)
        v1 -= -2.25
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int = 4294967295
        _ = fi10
        _ = v11
    }
    if false {
        v12 := GenS1{F0: -1}
        if (300 == 69) {
            v1 = v1
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            if false {
                fmt.Println(fi13, fi13)
                v14 := []int{(fi13 / ((-1) | 1)), fi13}
                _ = v14
            } else {
                var v15 uint8 = 127
                v15 = v15
                v16 := GenS3{F0: (v15 % ((v15) | 1)), F1: (fi13 << ((856) & 63)), F2: 1}
                _ = v15
                _ = v16
            }
            _ = fi13
        }
        _ = v12
    }
    v1 += 3.14159
    v17 := GenS4{F0: 4294967295}
    v18 := GenS5{F0: 1}
    v1 += v1
    var v19 uint16 = 65535
    var v20 int = 2147483647
    v21 := GenS5{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
