package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 uint32 }
func genFn0(p0 bool, p1 uint16, p2 string) bool {
    v6 := GenS1{F0: 22}
    fmt.Println(p1, p2)
    p2 = p2
    var v7 uint32 = 216
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 bool = (42 <= 1)
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (p2 < p2)
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 int16 = -32768
    v4 := make(map[uint32]bool)
    var v5 bool = (681 < 65535)
    v10 := GenS1{F0: -128}
    v11 := []int{}
    v5 = true
    v3 += (^(v3))
    if v5 {
        v12 := GenS2{F0: v3, F1: 2147483648}
        v5 = ((-102 < 42) || v5)
        if ((v3 << ((3) & 63)) > (v3 % ((v3) | 1))) {
            delete(v2, 42)
        }
        _ = v12
    } else {
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 string = "foo bar"
            _ = fi13
            _ = v14
        }
    }
    v3 = -1
    var v15 uint = 1
    v15 -= v15
    var v16 int = 943
    var v17 int8 = 60
    _ = len(v4)
    if len(v11) > 0 {
        v11[0] = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
