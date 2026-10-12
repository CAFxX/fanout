package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool; F1 string; F2 bool }
type GenS4 struct { F0 string; F1 uint }
func genFn0() uint8 {
    var v6 int = -421
    v7 := GenS1{F0: -448, F1: 200}
    if ((((v6 | (v6 | v6)) <= -2147483648) && (2 > 255)) && true) {
        v8 := GenS1{F0: -1, F1: 32767}
        var v9 int32 = -63
        _ = v8
        _ = v9
    } else {
        fmt.Println(v6)
        v10 := []int{-389, 42, 0}
        _ = v10
    }
    fmt.Println(v6)
    _ = v6
    _ = v7
    return 255
}

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 67
    v3 := make(map[int]string)
    v4 := make(map[string]int)
    v5 := make(map[uint32]string)
    if true {
        v11 := []int{1, 127, -2147483648, 405}
        _ = v11
    }
    var v12 uint32 = 1
    v2 -= 42
    v13 := GenS2{F0: 0, F1: -1}
    var v14 int = 32768
    v15 := GenS3{F0: (false || true), F1: "hello", F2: true}
    fmt.Println(v2)
    v16 := GenS4{F0: "\n\t", F1: 2}
    for fi17 := 0; fi17 < 20; fi17++ {
        fmt.Println(fi17)
        _ = fi17
    }
    var v18 uint16 = 13
    v5[1] = "hello"
    v12 = uint32(v18)
    for fi19 := 0; fi19 < 18; fi19++ {
        v14 = v14
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
