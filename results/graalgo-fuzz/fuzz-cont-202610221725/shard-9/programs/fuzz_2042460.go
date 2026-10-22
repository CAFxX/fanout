package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint, p1 int64, p2 uint16) uint32 {
    switch ("foo bar" + ("a" + "foo bar")) {
    case "":
        var v4 uint = p0
        _ = v4
    case "hello":
        fmt.Println(p2)
    default:
        p0 += p0
    }
    p1 = (p1 >> (((^(p0))) & 63))
    fmt.Println(p0, p2, p2)
    var v5 int16 = 1
    v6 := GenS1{F0: -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 9223372036854775807
    var v3 int64 = 165
    for fi7 := 0; fi7 < 12; fi7++ {
        var v8 int64 = v2
        v2 -= 42
        _ = fi7
        _ = v8
    }
    var v9 int = 32768
    var v10 uint64 = 42
    v11 := []int{(v9 + (-(v9)))}
    _ = len(v1)
    v9 -= 2147483647
    v10 = (v10 | v10)
    for fi12 := 0; fi12 < 8; fi12++ {
        v9 += v9
        fmt.Println(v2)
        _ = fi12
    }
    var v13 uint64 = (^(v10))
    var v14 uint = 149
    v15 := GenS1{F0: -9223372036854775808}
    var v16 bool = true
    var v17 uint = (v14 - (v14 / (((v14 << ((v14) & 63))) | 1)))
    var v18 uint16 = 2
    for fi19 := 0; fi19 < 18; fi19++ {
        v17 = (v14 + 4294967295)
        switch "\n\t" {
        case "hello":
            v17 += 42
        default:
            var v20 uint64 = (v10 * 42)
            _ = v20
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v9
}
