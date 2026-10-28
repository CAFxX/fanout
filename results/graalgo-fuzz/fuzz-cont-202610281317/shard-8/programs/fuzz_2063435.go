package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 bool }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 string, p1 int, p2 uint64) uint16 {
    switch (p1 % (((^(p1))) | 1)) {
    case 4294967295:
        for fi4 := 0; fi4 < 18; fi4++ {
            if (1 == -1) {
                var v5 int16 = 1000
                var v6 int8 = 1
                fmt.Println(fi4, v6, v6)
                _ = v5
                _ = v6
            } else {
                fmt.Println(p2, p2)
                fmt.Println(p0, p0)
            }
            _ = fi4
        }
    default:
        v7 := GenS1{F0: -242, F1: (-(p1)), F2: true}
        _ = v7
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    return 775
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]uint32)
    var v3 string = "a"
    v1 = (^(v1))
    fmt.Println(v1)
    v1 += v1
    var v8 uint64 = (v1 - (v1 + v1))
    switch 1 {
    case -2147483648:
        var v9 int = -2147483648
        _ = v9
    case -1:
        var v10 int64 = 0
        _ = v10
    default:
        var v11 uint32 = 856
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v3, fi12, v1)
        _ = fi12
    }
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 15; fi14++ {
        v8 = (v1 & (v8 | 2))
        if (4294967295 > 1) {
            v8 = 2
        } else {
            var v15 uint32 = 4294967295
            var v16 bool = (42 > 255)
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    v17 := []int{}
    v13 = (v13 >> (((^(v13))) & 63))
    for fi18 := 0; fi18 < 20; fi18++ {
        fmt.Println(v13, v8, v1)
        for fi19 := 0; fi19 < 7; fi19++ {
            var v20 string = "foo bar"
            delete(v2, (((v3 + v3) + "foo bar") + ""))
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println(v3)
    var v21 uint32 = 1
    var v22 float32 = (0.1 - (-2.25 - -2.25))
    v2[v3] = (v21 - 42)
    v23 := GenS2{F0: 18446744073709551615, F1: ""}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v8
}
