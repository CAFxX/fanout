package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32; F1 int32 }
type GenS5 struct { F0 uint; F1 bool }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]bool)
    var v3 float32 = (-2.25 + 0.1)
    if true {
        fmt.Println(v1, v3)
    } else {
        for fi4 := 0; fi4 < 6; fi4++ {
            for fi5 := 0; fi5 < 9; fi5++ {
                delete(v2, fi4)
                if true {
                    var v6 uint64 = 1
                    v3 = (3.14159 / (((3.14159 + (v3 * float32(3.14159))) + v3) + (-2.25 + (-2.25 * v3))))
                    _ = v6
                }
                _ = fi5
            }
            _ = fi4
        }
        fmt.Println(v1)
    }
    var v7 int64 = (v1 ^ -9223372036854775808)
    if (-178 >= 42) {
        if true {
            var v8 int64 = v7
            v9 := []int{32767, 4294967295, -2147483648, 128}
            _ = v8
            _ = v9
        }
        v10 := []int{32767, -2147483648}
        _ = v10
    } else {
        fmt.Println(v7, v1, v3)
        v11 := []int{32768, 4294967295}
        if true {
            var v12 int8 = 42
            _ = v12
        }
        _ = v11
    }
    if true {
        v13 := []int{32768, 32767, 4294967295, 127}
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 550
            _ = fi14
            _ = v15
        }
        _ = v13
    }
    v1 += (v1 * v1)
    if true {
        v3 += v3
        switch 127 {
        case 128:
            fmt.Println(v7, v7, v3)
        }
        var v16 float64 = float64(0.1)
        _ = v16
    } else {
        v17 := []int{264}
        _ = v17
    }
    v2[-1] = (-713 > 470)
    var v18 float64 = 0.0
    v19 := GenS1{F0: true, F1: -97, F2: 65535}
    v20 := GenS2{F0: 42, F1: 42, F2: 1}
    v21 := GenS3{F0: -1}
    _ = len(v2)
    v18 -= v18
    var v22 uint32 = 0
    for fi23 := 0; fi23 < 16; fi23++ {
        fmt.Println(v1, fi23, v7)
        _ = fi23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 2; fi25++ {
            var v26 uint16 = 42
            v27 := GenS4{F0: -135, F1: -215}
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi24
    }
    v28 := GenS5{F0: 4294967295, F1: (42 <= 0)}
    var v29 uint = 485
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v7
}
