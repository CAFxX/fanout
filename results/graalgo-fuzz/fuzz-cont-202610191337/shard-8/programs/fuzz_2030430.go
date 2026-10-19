package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 uint16) uint {
    p0 -= p0
    v4 := GenS1{F0: false, F1: "hello", F2: 42}
    v5 := []int{}
    p0 = p0
    v6 := GenS1{F0: (!false), F1: ((("foo bar" + ("hello" + (("" + "\n\t") + "hello"))) + "\n\t") + ""), F2: -101}
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return 255
}
func genFn1() uint {
    v7 := []int{143, 2147483647, 42, -2147483648}
    if len(v7) > 0 {
        v7[0] = -1
    }
    v8 := GenS2{F0: 982, F1: 32767, F2: 859}
    _ = v7
    _ = v8
    return genFn0(32767)
}

func main() {
    var v1 uint = 42
    var v2 uint16 = 255
    var v3 uint32 = 0
    var v9 int32 = 1
    var v10 uint32 = v3
    v11 := GenS2{F0: int(v9), F1: 611, F2: v3}
    fmt.Println(v10, v3)
    v12 := GenS2{F0: -1, F1: 4, F2: (v10 ^ uint32(v2))}
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{2147483647, fi13}
        _ = fi13
        _ = v14
    }
    fmt.Println(v10, v10, v2)
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v9
}
