package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 string; F2 string }
func genFn0(p0 int32, p1 uint16, p2 uint64) uint64 {
    v4 := []int{216}
    fmt.Println(p0)
    v5 := GenS1{F0: 18446744073709551615, F1: 142}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (p2 << ((1) & 63))
}
func genFn1(p0 int32) int32 {
    var v6 uint = 0
    p0 = -2147483648
    _ = p0
    _ = v6
    return p0
}

func main() {
    var v1 int32 = 587
    v2 := make(map[int]string)
    v3 := []string{"foo bar", ("a" + ""), ("a" + "a"), ("foo bar" + "foo bar")}
    v1 = v1
    v7 := GenS2{F0: 128}
    v1 += (v1 + genFn1(v1))
    v1 = genFn1(genFn1(v1))
    switch 4294967295 {
    case 270:
        var v8 uint32 = 0
        _ = v8
    case 806:
        v3 = append(v3, "foo bar")
    default:
        v9 := GenS1{F0: 65535, F1: 9223372036854775807}
        _ = v9
    }
    fmt.Println(v1)
    v1 -= 2147483647
    v10 := GenS3{F0: 2147483647, F1: (("foo bar" + "a") + ""), F2: "\n\t"}
    for fi11 := 0; fi11 < 18; fi11++ {
        switch (fi11 * -746) {
        case 32767:
            var v12 uint = 42
            _ = v12
        case 547:
            v1 -= (^(v1))
        }
        _ = fi11
    }
    var v13 int = 4294967295
    v1 = v1
    fmt.Println(v1, v1)
    v13 += 32767
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v7
}
