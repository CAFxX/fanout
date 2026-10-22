package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint8, p1 bool, p2 int8) int8 {
    p2 = p2
    p2 += 0
    _ = p0
    _ = p1
    _ = p2
    return (p2 ^ 42)
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    var v5 int32 = -452
    fmt.Println(v5, v5)
    fmt.Println(v1, v5)
    if (true || (true || (!true))) {
        delete(v2, -980)
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(fi6, v1)
            var v7 int64 = -1
            v5 += v5
            _ = fi6
            _ = v7
        }
        v1 = v1
    }
    v1 = (v1 & (v1 ^ v1))
    v8 := GenS1{F0: (42 > 58), F1: 1}
    fmt.Println(v5, v1)
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case 4294967295:
                var v11 int16 = (-(v1))
                _ = v11
            default:
                var v12 int32 = (v5 & -2147483648)
                _ = v12
            }
            if true {
                var v13 int64 = 435
                _ = v13
            }
            v1 -= (v1 % ((1) | 1))
            _ = fi10
        }
        v5 = -288
        v14 := []int{2147483647}
        _ = fi9
        _ = v14
    }
    v15 := GenS1{F0: true, F1: 65535}
    delete(v4, 4294967295)
    v1 = (v1 - (v1 & v1))
    var v16 uint = 1
    var v17 string = ("\n\t" + (("\n\t" + "\n\t") + (("" + "hello") + "\n\t")))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
