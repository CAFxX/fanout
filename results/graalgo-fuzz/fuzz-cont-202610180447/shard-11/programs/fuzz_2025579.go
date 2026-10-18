package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 uint8; F2 int32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint64) string {
    for fi7 := 0; fi7 < 17; fi7++ {
        if (((768 > 73) || false) || (true || ((true || (1 >= 0)) || (1 >= 42)))) {
            fmt.Println(p0, p0)
            var v8 uint = 978
            v8 = v8
            _ = v8
        }
        fmt.Println(fi7, fi7)
        _ = fi7
    }
    var v9 float64 = 1e10
    var v10 uint32 = 2147483647
    v11 := GenS1{F0: 128}
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return "hello"
}

func main() {
    v1 := []string{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := make(map[int]uint32)
    v5 := []int{-9223372036854775808, 9223372036854775807, 32767}
    var v6 int64 = -1
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 uint8 = 98
        _ = fi12
        _ = v13
    }
    v1 = append(v1, "\n\t")
    switch 9223372036854775807 {
    case 128:
        v6 -= 862
    case 42:
        v6 -= 452
    }
    if (false || ("foo bar" >= ("\n\t" + ("" + "a")))) {
        for fi14 := 0; fi14 < 1; fi14++ {
            v6 = v6
            if (32767 > 42) {
                v6 += 0
            } else {
                var v15 uint64 = 4294967295
                v16 := GenS2{F0: "a", F1: 42, F2: 0}
                _ = v15
                _ = v16
            }
            _ = fi14
        }
        var v17 string = ("a" + "hello")
        _ = v17
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := GenS3{F0: fi18}
        var v20 float64 = 0.1
        _ = fi18
        _ = v19
        _ = v20
    }
    v6 = v6
    var v21 float64 = -2.25
    v2[42] = 1
    var v22 uint64 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
