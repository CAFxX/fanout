package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64; F2 uint32 }
func genFn0(p0 uint32, p1 string, p2 int8) int64 {
    p0 += (p0 << ((1) & 63))
    if true {
        if true {
            v5 := GenS1{F0: (("hello" + p1) + (p1 + (p1 + "a")))}
            var v6 uint64 = 154
            _ = v5
            _ = v6
        } else {
            p1 = (p1 + p1)
            var v7 uint = 18446744073709551615
            _ = v7
        }
    }
    p1 = (p1 + ("foo bar" + (p1 + "")))
    p0 += (p0 | 2147483648)
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int8 = (p2 ^ (p2 ^ (p2 % ((p2) | 1))))
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return -9223372036854775808
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 uint8 = 95
    var v10 bool = false
    for fi11 := 0; fi11 < 8; fi11++ {
        v10 = true
        v12 := []int{(-(fi11))}
        _ = fi11
        _ = v12
    }
    if ((!v10) || (1 > -1)) {
        var v13 int8 = -29
        v10 = v10
        _ = v13
    } else {
        v14 := []int{}
        _ = len(v1)
        _ = v14
    }
    v15 := GenS2{F0: 1, F1: 694, F2: 674}
    var v16 int16 = 0
    v17 := []int{2147483647, 73, 42, -2147483648}
    v16 -= v16
    var v18 bool = false
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint8 = 2
        v21 := GenS1{F0: "hello"}
        if v10 {
            var v22 int64 = 0
            for fi23 := 0; fi23 < 20; fi23++ {
                v20 = 127
                v24 := GenS2{F0: 0, F1: 42, F2: 4294967295}
                v22 = (v22 - (v22 - v22))
                switch 32768 {
                case -9223372036854775808:
                    var v25 float32 = 0.0
                    _ = v25
                case -137:
                    v16 += v16
                case 0:
                    var v26 uint = 2
                    _ = v26
                }
                var v27 int = fi19
                _ = fi23
                _ = v24
                _ = v27
            }
            _ = v22
        } else {
            v4 += (^(v4))
            fmt.Println(fi19)
        }
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
