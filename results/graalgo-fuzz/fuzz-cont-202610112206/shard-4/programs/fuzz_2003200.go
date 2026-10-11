package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int32, p1 bool, p2 int16) int16 {
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 20; fi5++ {
            v6 := GenS1{F0: (p2 << ((4294967295) & 63))}
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    v7 := []int{279, 803, -9223372036854775808, 859}
    var v8 uint64 = 0
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 10; fi10++ {
            var v11 float32 = 1e10
            var v12 float64 = float64(3.14159)
            v13 := GenS1{F0: 32767}
            p1 = p1
            if (197 > 19) {
                fmt.Println(p2, v12, v11)
                _ = len(v7)
            } else {
                if true {
                    switch (fi10 * fi9) {
                    case -2147483648:
                        fmt.Println(v8, p0, fi10)
                    case -114:
                        if p1 {
                            p2 -= (p2 / ((p2) | 1))
                            var v14 int16 = p2
                            _ = v14
                        } else {
                            p1 = (!p1)
                        }
                    case -1:
                        p1 = (p1 && (!p1))
                    }
                    v7 = append(v7, 32768)
                }
            }
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(p0, fi9, p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 int32 = 0
    v2 := make(map[int]uint32)
    v3 := []int{}
    if true {
        v15 := GenS1{F0: -32768}
        if true {
            _ = len(v3)
            v16 := GenS2{F0: v1}
            var v17 int8 = 6
            _ = v16
            _ = v17
        }
        _ = v15
    }
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint16 = 65535
        var v20 uint64 = 18446744073709551615
        if len(v3) > 0 {
            v3[0] = (^(fi18))
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        _ = len(v3)
        _ = fi21
    }
    v1 = v1
    v1 += v1
    if (-1 != 9223372036854775807) {
        v1 += (v1 >> ((65535) & 63))
        v22 := GenS1{F0: 0}
        _ = v22
    }
    if (!true) {
        v23 := GenS2{F0: v1}
        _ = v23
    }
    var v24 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
}
