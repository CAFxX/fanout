package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
func genFn0(p0 uint32, p1 uint32) uint32 {
    for fi4 := 0; fi4 < 10; fi4++ {
        p1 -= 661
        _ = fi4
    }
    p1 = (p0 / ((806) | 1))
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := []int{fi5, (fi5 | int(p0))}
        _ = fi5
        _ = v6
    }
    p1 = p0
    var v7 int16 = 1000
    _ = p0
    _ = p1
    _ = v7
    return p1
}

func main() {
    var v1 int64 = -1
    var v2 uint32 = 136
    v3 := make(map[uint32]string)
    fmt.Println(v1, v1)
    var v8 uint64 = 2
    var v9 string = "a"
    if true {
        var v10 uint16 = 0
        if (true || false) {
            if true {
                v10 += (v10 << ((65535) & 63))
                var v11 int = 9223372036854775807
                _ = v11
            }
            if true {
                v8 += 962
                v12 := GenS1{F0: -228, F1: v10, F2: -1}
                v13 := []int{127}
                _ = v12
                _ = v13
            }
        }
        v14 := GenS1{F0: int16(v2), F1: 2, F2: 42}
        _ = v10
        _ = v14
    }
    v15 := GenS1{F0: 32767, F1: 1, F2: -1}
    var v16 int16 = 1000
    if true {
        for fi17 := 0; fi17 < 2; fi17++ {
            v2 += genFn0((v2 ^ v2), (^(v2)))
            fmt.Println(v16, v8, v9)
            var v18 int32 = 2147483647
            _ = fi17
            _ = v18
        }
        v1 -= (v1 << ((519) & 63))
    }
    v19 := GenS1{F0: 160, F1: 795, F2: -23}
    v20 := GenS1{F0: int16(v8), F1: 32767, F2: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
