package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint }
type GenS2 struct { F0 string; F1 uint16; F2 bool }
func genFn0() int8 {
    v7 := GenS1{F0: 492, F1: "a", F2: 18446744073709551615}
    for fi8 := 0; fi8 < 9; fi8++ {
        v9 := []int{fi8, fi8}
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 79
    _ = v10
    _ = v7
    return 66
}
func genFn1() int32 {
    v11 := []int{127, 4294967295, 42, 0}
    var v12 bool = (!false)
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 1.5
    var v3 int32 = 2147483647
    var v4 uint32 = 248
    var v5 int16 = int16(v3)
    var v6 bool = ((false || false) || false)
    v13 := GenS2{F0: ("hello" + "\n\t"), F1: 1, F2: true}
    v14 := []int{9223372036854775807, 32768, -336, 128}
    var v15 int8 = 0
    v4 = 2147483648
    if (!v6) {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v4 -= v4
    }
    v4 = 65535
    v5 = (v5 - v5)
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    v16 := GenS2{F0: ((("" + "hello") + "foo bar") + "a"), F1: 678, F2: v6}
    v15 += int8(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
