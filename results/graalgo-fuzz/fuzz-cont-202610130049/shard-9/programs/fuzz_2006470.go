package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint32; F2 int }
func genFn0(p0 int32) uint32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        if (true || false) {
            v7 := GenS1{F0: (fi6 << ((175) & 63)), F1: 1}
            v8 := GenS1{F0: fi6, F1: 2147483648}
            p0 += p0
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    if true {
        var v9 uint32 = 2
        v10 := GenS2{F0: v9}
        _ = v10
        _ = v9
    } else {
        p0 -= (p0 >> ((115) & 63))
    }
    var v11 uint64 = 1
    v12 := []int{32767, 0}
    v13 := GenS1{F0: 0, F1: 2147483647}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return 878
}
func genFn1(p0 string) uint8 {
    if (647 > 42) {
        v14 := []int{32767, 562, 9223372036854775807}
        var v15 uint16 = 1
        _ = v14
        _ = v15
    }
    p0 = ""
    _ = p0
    return 2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = -2.25
    var v3 uint16 = 719
    var v4 bool = (!(2 >= 42))
    var v5 uint8 = 255
    var v16 int8 = 42
    v5 += (v5 ^ v5)
    var v17 uint16 = 65535
    fmt.Println(v17, v17, v3)
    fmt.Println(v17, v3)
    if (v17 <= 1000) {
        switch "foo bar" {
        case "":
            var v18 int8 = (v16 << ((282) & 63))
            _ = v18
        case "hello":
            v17 = v17
        case "\n\t":
            fmt.Println(v16, v16)
        }
        v2 += v2
        var v19 uint = 65535
        _ = v19
    }
    v4 = (!(float64((v2 / (float32((0.1 + -2.25)) + (3.14159 * (v2 - v2))))) <= 0.0))
    _ = len(v1)
    var v20 int32 = -1
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 bool = false
        if true {
            var v23 uint = 759
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    v20 = (v20 % ((v20) | 1))
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS3{F0: (fi24 & (fi24 & fi24)), F1: 10, F2: (-(fi24))}
        fmt.Println(v5, v17)
        var v26 float64 = (3.14159 * 3.14159)
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
