package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
func genFn0(p0 uint64, p1 bool, p2 int8) uint8 {
    fmt.Println(p0, p2)
    for fi7 := 0; fi7 < 1; fi7++ {
        fmt.Println(fi7)
        _ = fi7
    }
    p1 = p1
    v8 := GenS1{F0: (("a" + "") + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 127
}
func genFn1(p0 uint64, p1 uint, p2 uint64) bool {
    for fi9 := 0; fi9 < 7; fi9++ {
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                p2 = 0
                p2 -= (p2 + p2)
                _ = fi10
            }
        }
        v11 := []int{}
        for fi12 := 0; fi12 < 19; fi12++ {
            v11 = append(v11, 78)
            var v13 uint8 = 0
            if len(v11) > 0 {
                v11[0] = fi12
            }
            _ = fi12
            _ = v13
        }
        _ = fi9
        _ = v11
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{(fi14 << (((^(p1))) & 63)), 42}
        fmt.Println(p1, p2)
        _ = fi14
        _ = v15
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -1
    v3 := make(map[uint32]string)
    var v4 uint8 = 147
    var v5 uint8 = 0
    v6 := make(map[int]int)
    v16 := GenS1{F0: "hello"}
    v17 := GenS2{F0: 1, F1: "foo bar", F2: 4294967295}
    delete(v1, 2)
    var v18 bool = ((42 < 1) && ((!false) || true))
    v2 -= (^(v2))
    if (2147483647 < 58) {
        v19 := []int{42, -9223372036854775808, 9223372036854775807, 4294967295}
        _ = v19
    } else {
        fmt.Println(v5, v4)
        v2 = (v2 >> ((65535) & 63))
    }
    v4 = (v4 + 42)
    if true {
        v2 += v2
    }
    var v20 string = "\n\t"
    if (-111 <= -58) {
        fmt.Println(v4, v18)
        v4 = genFn0(18446744073709551615, false, 127)
    }
    var v21 float32 = float32(-2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
