package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 int8 }
func genFn0() string {
    var v7 uint64 = 1
    switch "hello" {
    case "a":
        fmt.Println(v7)
    case "hello":
        var v8 int = -1
        _ = v8
    case "\n\t":
        if (!(2 > (^(v7)))) {
            var v9 bool = false
            v10 := GenS1{F0: "hello", F1: 1, F2: 1}
            for fi11 := 0; fi11 < 3; fi11++ {
                for fi12 := 0; fi12 < 3; fi12++ {
                    v13 := []int{9223372036854775807, fi11, fi11}
                    v13 = append(v13, (fi12 % ((fi11) | 1)))
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: "a", F1: 32767, F2: 20}
                _ = fi11
                _ = v14
            }
            _ = v10
            _ = v9
        }
    }
    v15 := []int{}
    switch "" {
    case "a":
        fmt.Println(v7)
    case "\n\t":
        if (255 != 195) {
            fmt.Println(v7)
            fmt.Println(v7)
        }
    case "hello":
        v7 -= (v7 / ((289) | 1))
    default:
        if false {
            var v16 int16 = 798
            _ = v16
        }
    }
    _ = v15
    _ = v7
    return "foo bar"
}
func genFn1(p0 int32, p1 string, p2 int16) uint64 {
    if false {
        p1 = ("a" + "")
    }
    if true {
        fmt.Println(p2)
        p1 = p1
    }
    v17 := GenS1{F0: p1, F1: -32768, F2: 128}
    var v18 float64 = 0.0
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 100
    var v3 uint32 = 579
    var v4 uint32 = (v3 | v3)
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]int)
    for fi19 := 0; fi19 < 4; fi19++ {
        _ = len(v5)
        _ = fi19
    }
    v20 := GenS2{F0: -9223372036854775808}
    var v21 uint = 177
    for fi22 := 0; fi22 < 7; fi22++ {
        v3 = v3
        v3 = (v3 ^ (v3 / ((0) | 1)))
        _ = fi22
    }
    v23 := GenS2{F0: -475}
    switch 756 {
    case -1:
        v2 -= 1
    }
    var v24 uint8 = 0
    switch -494 {
    case -2147483648:
        var v25 int8 = v2
        _ = v25
    }
    var v26 bool = (uint8(v3) == v24)
    v27 := GenS2{F0: 42}
    v28 := GenS3{F0: 9223372036854775807, F1: 119}
    var v29 uint32 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
