package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint64, p1 uint32, p2 uint8) int16 {
    switch 0 {
    case -694:
        fmt.Println(p2, p0, p1)
    case 32768:
        v7 := GenS1{F0: (0.0 >= 1.5)}
        _ = v7
    case -540:
        p1 -= (p1 ^ 2147483648)
    }
    v8 := GenS2{F0: 1, F1: 823}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -32768
}
func genFn1() uint {
    var v9 int8 = 121
    for fi10 := 0; fi10 < 8; fi10++ {
        if false {
            for fi11 := 0; fi11 < 15; fi11++ {
                var v12 int = (^(fi11))
                _ = fi11
                _ = v12
            }
            var v13 float32 = 0.1
            if false {
                var v14 float32 = ((v13 - -2.25) + v13)
                _ = v14
            }
            _ = v13
        }
        var v15 int16 = 21
        v16 := []int{}
        _ = fi10
        _ = v15
        _ = v16
    }
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    var v3 uint32 = 2
    var v4 uint = 2
    var v5 string = "foo bar"
    v6 := make(map[string]bool)
    switch 4294967295 {
    case 9223372036854775807:
        fmt.Println(v4)
    default:
        v3 += (v3 + 65535)
    }
    v17 := GenS2{F0: int8(v3), F1: 42}
    v4 -= (v4 & (v4 / ((v4) | 1)))
    _ = len(v1)
    switch "hello" {
    case "foo bar":
        v18 := []int{int(v4), -2147483648}
        _ = v18
    }
    var v19 uint8 = 127
    fmt.Println(v3)
    v20 := GenS2{F0: 95, F1: 866}
    v5 = v5
    delete(v6, ("a" + "foo bar"))
    var v21 int64 = 190
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
