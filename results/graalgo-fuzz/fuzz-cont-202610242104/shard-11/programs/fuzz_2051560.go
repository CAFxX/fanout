package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32; F1 uint; F2 int16 }
func genFn0() uint {
    v7 := []int{-490, 1}
    var v8 string = ("hello" + "")
    _ = len(v7)
    var v9 uint16 = 1
    if true {
        var v10 string = v8
        fmt.Println(v8)
        _ = v10
    }
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 2147483648
    v3 := make(map[int]int)
    v4 := []int{-9223372036854775808, -54, 42, 42}
    v5 := make(map[string]string)
    v6 := make(map[uint32]string)
    fmt.Println(v2)
    switch ("\n\t" + ("\n\t" + "")) {
    case "hello":
        v2 -= (^(v2))
    case "a":
        var v11 int = -161
        _ = v11
    }
    var v12 int16 = 1000
    v4 = append(v4, -89)
    v4 = append(v4, 127)
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS1{F0: 1, F1: 2, F2: 793}
        var v15 int = (^(fi13))
        if false {
            v2 = v2
        } else {
            v12 = v12
            v16 := GenS1{F0: -53, F1: 42, F2: v12}
            for fi17 := 0; fi17 < 3; fi17++ {
                if (false && ((v2 >= (v2 & 2147483647)) && false)) {
                    var v18 int = (fi13 & fi17)
                    v19 := GenS1{F0: 100, F1: 677, F2: 0}
                    v20 := GenS1{F0: 1, F1: 2, F2: (v12 << ((18446744073709551615) & 63))}
                    var v21 float32 = (0.0 * 0.1)
                    _ = v18
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    v22 := GenS1{F0: -128, F1: 18446744073709551615, F2: 1}
                    var v23 int = (fi17 | -563)
                    _ = v22
                    _ = v23
                }
                _ = fi17
            }
            _ = v16
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v12, v12)
    if true {
        v24 := GenS1{F0: 0, F1: 18446744073709551615, F2: 32767}
        _ = v24
    }
    var v25 uint16 = 65535
    var v26 uint64 = 4294967295
    v2 = uint32(v25)
    v25 = (v25 - (v25 | uint16(v2)))
    v27 := GenS2{F0: v2, F1: 2, F2: (^(v12))}
    fmt.Println(v2, v2, v25)
    var v28 uint8 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
