package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 bool; F1 int32 }
func genFn0(p0 uint8, p1 uint32) uint {
    if (42 == p1) {
        for fi4 := 0; fi4 < 4; fi4++ {
            for fi5 := 0; fi5 < 15; fi5++ {
                p1 -= p1
                var v6 bool = (65535 >= 65535)
                _ = fi5
                _ = v6
            }
            _ = fi4
        }
    }
    switch 42 {
    case -1:
        v7 := []int{769, 809, 128}
        _ = v7
    case 0:
        var v8 uint = 690
        _ = v8
    }
    if true {
        v9 := GenS1{F0: 65535, F1: -1}
        var v10 int32 = 2147483647
        fmt.Println(p1, v10, v10)
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int32 = -718
        var v13 string = ""
        _ = fi11
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int8, p1 uint16) int64 {
    v14 := GenS2{F0: 1, F1: -1, F2: 938}
    v15 := GenS1{F0: 2, F1: 1000}
    v16 := []int{-185, 2147483647, 32768}
    fmt.Println(p1, p1)
    v17 := GenS2{F0: -2147483648, F1: -111, F2: 1}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    return -991
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 751
    var v3 uint64 = 217
    fmt.Println(v2, v3)
    switch "a" {
    case "hello":
        v2 = v2
    case "foo bar":
        v18 := GenS1{F0: 281, F1: 32767}
        _ = v18
    case "":
        var v19 float32 = 0.0
        _ = v19
    }
    var v20 float32 = 0.1
    var v21 int32 = -306
    var v22 float32 = (3.14159 + 1.5)
    fmt.Println(v2, v2)
    for fi23 := 0; fi23 < 3; fi23++ {
        switch fi23 {
        case 42:
            v3 += v3
        case -552:
            v24 := GenS2{F0: -186, F1: -9223372036854775808, F2: 65535}
            _ = v24
        }
        _ = fi23
    }
    var v25 bool = (!(int(v3) != 0))
    for fi26 := 0; fi26 < 4; fi26++ {
        v3 = v3
        v27 := v1[0]
        v25 = (0.1 == 0.0)
        _ = fi26
        _ = v27
    }
    fmt.Println(v2)
    v21 = (v21 / ((v21) | 1))
    var v28 uint16 = 42
    switch "\n\t" {
    case "foo bar":
        v29 := GenS3{F0: 2147483648, F1: v25}
        _ = v29
    case "hello":
        v30 := GenS1{F0: 65535, F1: 1000}
        _ = v30
    case "a":
        var v31 float32 = 0.0
        _ = v31
    }
    v32 := GenS4{F0: (1000 != 1), F1: (v21 ^ (v21 - -981))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v28
    _ = v3
    _ = v32
}
