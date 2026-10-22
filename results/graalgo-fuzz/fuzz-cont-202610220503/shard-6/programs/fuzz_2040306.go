package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
func genFn0(p0 uint16, p1 int64, p2 int) int32 {
    for fi5 := 0; fi5 < 10; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 int64 = 9223372036854775807
            _ = fi6
            _ = v7
        }
        v8 := []int{p2, p2, fi5, fi5}
        if false {
            p0 -= p0
            if len(v8) > 0 {
                v8[0] = (fi5 - -2147483648)
            }
            _ = len(v8)
        }
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        p1 = (p1 << ((788) & 63))
        p0 -= (p0 | p0)
        fmt.Println(p0)
        _ = fi9
    }
    v10 := []int{(p2 * p2), p2, p2, int(p0)}
    var v11 uint = 256
    if len(v10) > 0 {
        v10[0] = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return -2147483648
}
func genFn1(p0 int32, p1 uint16) uint64 {
    p0 = -1
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 150
}

func main() {
    var v1 int32 = 2147483647
    v2 := make(map[string]bool)
    var v3 uint8 = 255
    v4 := make(map[int]int)
    v12 := GenS1{F0: 0, F1: 0}
    var v13 bool = ((926 == -2147483648) || false)
    v3 = 255
    fmt.Println(v1, v1)
    if v13 {
        fmt.Println(v13)
        delete(v2, ((("a" + ("" + ("\n\t" + ""))) + "foo bar") + ""))
        if v13 {
            v4[-94] = -1
            fmt.Println(v1, v3)
            fmt.Println(v1)
        }
    } else {
        v1 = 0
        v3 += 128
        var v14 float32 = (-2.25 + (3.14159 / (-2.25 / float32(((-2.25 - ((0.0 / ((((3.14159 * 3.14159) + 1e10) / 1.0) - float64((-2.25 * (0.0 * 0.1))))) - ((float64(float32(1.5)) + float64(float32(0.0))) * 1e10))) - (3.14159 * (1e10 + 0.0)))))))
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 int32 = (v1 >> ((18446744073709551615) & 63))
        fmt.Println(v13)
        _ = fi15
        _ = v16
    }
    v2["a"] = v13
    var v17 uint8 = (v3 | (v3 - (v3 + (v3 << ((65535) & 63)))))
    v18 := GenS1{F0: 9223372036854775807, F1: (v1 | (v1 >> ((390) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
