package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int; F1 int16 }
func genFn0(p0 uint16) int16 {
    p0 -= p0
    var v5 float32 = -2.25
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint16, p1 uint8) int8 {
    var v6 uint16 = (p0 ^ (p0 / ((65535) | 1)))
    p1 = p1
    v7 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -1
}

func main() {
    v1 := []string{"hello", ("a" + "a"), "hello", ("hello" + "foo bar")}
    var v2 uint8 = 128
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]string)
    fmt.Println(v2)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    v2 = (v2 << ((193) & 63))
    v8 := GenS1{F0: 635, F1: 782}
    for fi9 := 0; fi9 < 3; fi9++ {
        _ = len(v3)
        if (!(621 != 1)) {
            if (false && (!false)) {
                var v10 int16 = -175
                _ = v10
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        } else {
            v2 -= v2
        }
        v2 = 0
        _ = fi9
    }
    var v11 uint = 42
    if (v2 >= v2) {
        v11 += v11
        if (127 <= 42) {
            var v12 uint32 = 2147483648
            _ = v12
        }
        var v13 int8 = 100
        _ = v13
    }
    v14 := GenS2{F0: 0, F1: 64}
    if (v11 <= 255) {
        if len(v1) > 0 {
            v1[0] = "a"
        }
    } else {
        v2 += (v2 & (^(v2)))
    }
    v15 := GenS2{F0: 42, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
