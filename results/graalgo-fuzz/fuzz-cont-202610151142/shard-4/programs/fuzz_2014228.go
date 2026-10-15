package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 uint32, p1 string) int {
    v7 := []int{2147483647}
    v8 := GenS1{F0: 65535, F1: 6, F2: 1}
    fmt.Println(p1)
    fmt.Println(p0)
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(p1)
    case 4294967295:
        p1 = ("a" + p1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 275
}

func main() {
    var v1 uint64 = 328
    v2 := make(map[int]string)
    v3 := make(map[int]bool)
    v4 := []uint32{2147483647, 2147483648, 676}
    v5 := []string{("hello" + ((("a" + "\n\t") + "hello") + "a")), "hello"}
    v6 := []uint32{}
    v9 := GenS1{F0: 216, F1: -614, F2: -705}
    fmt.Println(v1)
    fmt.Println(v1)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 627
        v1 -= (v1 << ((65535) & 63))
        _ = fi10
        _ = v11
    }
    switch -925 {
    case -797:
        v12 := GenS1{F0: 271, F1: 2147483647, F2: 42}
        _ = v12
    case 127:
        v1 -= v1
    case -566:
        v1 += 56
    default:
        for fi13 := 0; fi13 < 2; fi13++ {
            v1 = 1
            _ = fi13
        }
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(fi14, fi14, v1)
            if true {
                if false {
                    var v16 bool = false
                    _ = v16
                }
                v17 := GenS1{F0: 42, F1: 2147483647, F2: 0}
                fmt.Println(v1, fi15)
                _ = v17
            }
            var v18 uint32 = 0
            _ = fi15
            _ = v18
        }
        _ = fi14
    }
    var v19 float64 = 1e10
    fmt.Println(v19, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
