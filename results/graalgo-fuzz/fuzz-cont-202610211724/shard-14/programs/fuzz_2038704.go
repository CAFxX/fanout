package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 uint8, p1 int16) string {
    v4 := []int{128, -2147483648}
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 127
        if (!true) {
            switch "\n\t" {
            case "a":
                fmt.Println(p0)
            case "foo bar":
                v7 := GenS1{F0: 2147483647, F1: 42, F2: 127}
                _ = v7
            }
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 bool = true
                v10 := GenS1{F0: 2, F1: 42, F2: v6}
                v11 := GenS1{F0: 42, F1: 0, F2: v6}
                if len(v4) > 0 {
                    v4[0] = 127
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
        } else {
            fmt.Println(p1)
            v12 := GenS1{F0: 65535, F1: 4294967295, F2: (v6 << ((2) & 63))}
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return ("hello" + "foo bar")
}
func genFn1(p0 uint8, p1 uint16, p2 string) int {
    for fi13 := 0; fi13 < 20; fi13++ {
        p1 -= (p1 << ((2) & 63))
        v14 := []int{fi13}
        var v15 float64 = 1.5
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint64 = 797
    var v17 uint8 = (p0 * (p0 | p0))
    var v18 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return 128
}

func main() {
    var v1 float32 = ((0.1 / 0.1) * 0.1)
    v2 := make(map[uint32]bool)
    v3 := []int{-1, 158}
    for fi19 := 0; fi19 < 14; fi19++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi19
    }
    v1 += v1
    v2[42] = (2147483647 == -678)
    v1 = v1
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v1 += v1
    v20 := GenS2{F0: 864, F1: 65535, F2: 2}
    var v21 uint16 = 2
    v1 += v1
    v1 -= 3.14159
    var v22 uint64 = 42
    switch "" {
    case "hello":
        var v23 bool = true
        _ = v23
    }
    v21 = (v21 ^ v21)
    fmt.Println(v21, v1)
    v21 = 2
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
