package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) int {
    p0 = (p0 >> ((uint(p0)) & 63))
    p0 += p0
    if false {
        p0 = (p0 * p0)
    } else {
        var v6 uint = 65535
        var v7 uint8 = 52
        fmt.Println(v7, v6)
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: 2147483647, F1: 346, F2: 111}
    v9 := []int{}
    _ = p0
    _ = v8
    _ = v9
    return 762
}
func genFn1(p0 bool, p1 uint) uint64 {
    fmt.Println(p1, p1)
    p0 = p0
    p1 += p1
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []string{("a" + "hello")}
    var v2 float32 = -2.25
    v3 := []int{1, -239, -734}
    var v4 int64 = 9223372036854775807
    v5 := make(map[uint32]string)
    for fi10 := 0; fi10 < 10; fi10++ {
        v2 -= 1.5
        if (true || (42 >= 2147483647)) {
            v2 = v2
        }
        v5[1] = "a"
        _ = fi10
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v2, v4)
    }
    fmt.Println(v2)
    delete(v5, 2)
    if (!true) {
        v4 -= (v4 / (((^(v4))) | 1))
    } else {
        v11 := GenS1{F0: -1, F1: -2147483648, F2: 0}
        if true {
            v4 = 423
            v1 = append(v1, ("" + "a"))
        } else {
            fmt.Println(v2, v2)
            v2 -= v2
            v2 -= (1.5 + v2)
        }
        v4 = v4
        _ = v11
    }
    var v12 int32 = -1
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS2{F0: 139}
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int8 = 1
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if (9223372036854775807 <= (v4 | 9223372036854775807)) {
            var v17 uint = 0
            _ = v17
        } else {
            for fi18 := 0; fi18 < 20; fi18++ {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                _ = fi18
            }
        }
        _ = fi15
        _ = v16
    }
    var v19 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
