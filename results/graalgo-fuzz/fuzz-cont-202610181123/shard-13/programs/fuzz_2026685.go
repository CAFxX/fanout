package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int64 }
func genFn0(p0 int, p1 string, p2 int32) int64 {
    for fi4 := 0; fi4 < 12; fi4++ {
        v5 := []int{fi4}
        p2 += (p2 << ((42) & 63))
        _ = fi4
        _ = v5
    }
    p1 = (p1 + p1)
    _ = p0
    _ = p1
    _ = p2
    return -845
}

func main() {
    v1 := []uint32{}
    var v2 float32 = 3.14159
    v3 := []int{79, 9223372036854775807, 128}
    v6 := GenS1{F0: (!(1e10 < 0.1)), F1: 317, F2: genFn0(32767, ("\n\t" + "a"), -2147483648)}
    v7 := GenS1{F0: false, F1: 128, F2: 42}
    for fi8 := 0; fi8 < 11; fi8++ {
        var v9 uint = 18446744073709551615
        v2 = (3.14159 - ((v2 + (1e10 + v2)) / ((v2 * (((float32(((1.5 / 1.0) * 1e10)) - ((3.14159 / -2.25) + v2)) + v2) * (1.5 - (v2 / 1.0)))) - (3.14159 + v2))))
        _ = fi8
        _ = v9
    }
    v1 = append(v1, 555)
    switch "\n\t" {
    case "hello":
        for fi10 := 0; fi10 < 1; fi10++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
        }
    case "a":
        var v11 uint8 = 128
        _ = v11
    }
    fmt.Println(v2)
    v2 = (v2 - (((1e10 * 0.0) - (1.5 / 1.5)) * v2))
    var v12 int = 0
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
