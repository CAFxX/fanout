package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int }
func genFn0(p0 uint, p1 uint, p2 uint32) uint {
    fmt.Println(p1, p2)
    fmt.Println(p0)
    fmt.Println(p1)
    var v5 int64 = 42
    if true {
        fmt.Println(p0, p0, v5)
        if true {
            var v6 int32 = -1
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 uint32 = (p2 - 1)
                if (2 != 65535) {
                    v9 := []int{fi7, 4294967295, (fi7 ^ -818), 694}
                    v6 = (v6 + v6)
                    _ = v9
                }
                v10 := GenS1{F0: 128, F1: p2, F2: 42}
                _ = fi7
                _ = v10
                _ = v8
            }
            _ = v6
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v5)
                fmt.Println(fi11, fi11)
                _ = fi11
            }
            fmt.Println(p1)
            v12 := []int{487}
            _ = v12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 15
}

func main() {
    var v1 uint = 65535
    v2 := make(map[int]bool)
    var v3 int8 = 100
    v4 := []uint32{65535, 2147483648}
    v1 = v1
    var v13 bool = (("" + "\n\t") <= ("a" + "hello"))
    _ = len(v4)
    v14 := GenS2{F0: 255, F1: v13, F2: -9223372036854775808}
    _ = len(v2)
    v15 := GenS1{F0: 1, F1: 65535, F2: v3}
    if v13 {
        _ = len(v2)
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
