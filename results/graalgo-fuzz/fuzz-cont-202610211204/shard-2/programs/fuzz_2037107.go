package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint8, p1 uint32, p2 int) int {
    switch (p2 << ((65535) & 63)) {
    case 643:
        if true {
            v4 := GenS1{F0: (p0 + 77), F1: 1}
            v5 := []int{(^(p2)), 571, p2, (p2 + 32767)}
            _ = v4
            _ = v5
        } else {
            v6 := []int{}
            _ = len(v6)
            _ = v6
        }
    case 127:
        p2 -= p2
    }
    v7 := []int{(p2 + (p2 / ((2147483647) | 1))), (p2 * p2), (p2 >> ((42) & 63))}
    p1 = 298
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p2
}

func main() {
    var v1 uint = 0
    v2 := make(map[uint32]uint32)
    var v3 bool = (true || (!true))
    var v8 bool = true
    v9 := GenS1{F0: 0, F1: 255}
    v2[2147483648] = 2
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 20; fi11++ {
            v1 = (v1 - v1)
            v12 := GenS1{F0: uint8(v1), F1: 42}
            fmt.Println(v8, fi10, fi11)
            v3 = v8
            _ = fi11
            _ = v12
        }
        var v13 int8 = 1
        v14 := GenS1{F0: 248, F1: uint16(v1)}
        _ = fi10
        _ = v13
        _ = v14
    }
    fmt.Println(v3)
    switch "\n\t" {
    case "hello":
        v1 = (v1 + 1)
    case "a":
        for fi15 := 0; fi15 < 14; fi15++ {
            v8 = false
            _ = fi15
        }
    }
    if (!v3) {
        v16 := GenS1{F0: 50, F1: 0}
        if v8 {
            v8 = true
        }
        _ = v16
    }
    v8 = true
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
