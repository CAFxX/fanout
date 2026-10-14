package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 bool }
func genFn0(p0 uint) int64 {
    switch "\n\t" {
    case "hello":
        var v7 float32 = 0.0
        _ = v7
    }
    v8 := GenS1{F0: 315, F1: 1, F2: true}
    var v9 int32 = -1
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint8 = 89
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v8
    _ = v9
    return 31
}
func genFn1(p0 int8) int32 {
    v12 := []int{128, -2, 903}
    if len(v12) > 0 {
        v12[0] = -891
    }
    v13 := GenS1{F0: 0, F1: 42, F2: (834 == -972)}
    var v14 int32 = -1
    v12 = append(v12, -2147483648)
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return v14
}

func main() {
    v1 := []uint32{0, 65535}
    var v2 int32 = 2147483647
    v3 := []string{("a" + "\n\t")}
    var v4 int = -203
    var v5 int = v4
    var v6 bool = (2147483647 < 29)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v15 float32 = 1e10
    v1 = append(v1, 140)
    v15 = -2.25
    var v16 uint64 = 2
    if (42 >= v5) {
        v6 = true
        if len(v3) > 0 {
            v3[0] = ("hello" + "foo bar")
        }
    }
    var v17 float32 = 1.5
    v16 = (v16 * 587)
    fmt.Println(v16, v17)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
