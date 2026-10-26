package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -1
    var v3 uint8 = 1
    var v4 int8 = 0
    var v5 uint = 255
    v6 := GenS1{F0: 2, F1: (!true), F2: ((((("\n\t" + "hello") + "a") + "\n\t") + "") + "a")}
    v7 := GenS2{F0: 493}
    if true {
        var v8 uint32 = 2147483647
        v2 -= -1
        _ = v8
    }
    v4 = v4
    v2 += int16(v5)
    v5 = v5
    for fi9 := 0; fi9 < 13; fi9++ {
        if true {
            v10 := []int{fi9, fi9, 32767, 9223372036854775807}
            _ = v10
        } else {
            var v11 uint16 = 255
            fmt.Println(v2, fi9)
            _ = v11
        }
        v12 := GenS3{F0: "", F1: 42}
        fmt.Println(v3)
        _ = fi9
        _ = v12
    }
    v3 += 42
    var v13 uint = (v5 % ((1) | 1))
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := []int{int(v3), -1, fi14}
        _ = fi14
        _ = v15
    }
    var v16 uint64 = 42
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := []int{1}
        v3 = (v3 + (v3 * v3))
        fmt.Println(fi17, v5)
        _ = fi17
        _ = v18
    }
    v4 -= v4
    var v19 uint = (v5 * uint(v3))
    if (!((0.1 * 3.14159) != 0.1)) {
        v16 = 935
    } else {
        var v20 uint64 = 0
        var v21 bool = false
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        v2 = 42
        _ = fi22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
