package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
func genFn0() int8 {
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 int = (fi7 << ((2) & 63))
        _ = fi7
        _ = v8
    }
    switch 385 {
    case -2147483648:
        v9 := []int{32768}
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 int8 = 127
        v11 += v11
        var v12 float64 = 0.0
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        v14 := []int{-9223372036854775808, -9223372036854775808, -9223372036854775808}
        _ = fi13
        _ = v14
    }
    return 42
}

func main() {
    var v1 float32 = -2.25
    v2 := []string{"\n\t", "", ("\n\t" + "foo bar"), "hello"}
    v3 := make(map[uint32]string)
    var v4 uint8 = 255
    v5 := make(map[string]uint32)
    v6 := []int{32768}
    if (true && false) {
        v15 := GenS1{F0: -1, F1: 0, F2: (62 == 127)}
        _ = v15
    }
    v1 += 0.1
    var v16 uint32 = 2147483647
    _ = len(v3)
    fmt.Println(v16)
    var v17 uint8 = 55
    var v18 int32 = 42
    var v19 uint16 = 255
    fmt.Println(v18, v17, v18)
    v16 += v16
    v19 -= (v19 & v19)
    var v20 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
