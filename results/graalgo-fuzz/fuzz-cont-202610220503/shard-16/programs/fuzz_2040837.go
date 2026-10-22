package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint32 = 2
    v2 := make(map[uint32]int)
    var v3 float32 = ((1.5 + -2.25) - 0.0)
    switch -490 {
    case 32768:
        for fi4 := 0; fi4 < 7; fi4++ {
            for fi5 := 0; fi5 < 19; fi5++ {
                var v6 uint8 = 2
                fmt.Println(v1, fi4, v1)
                fmt.Println(v6)
                var v7 uint32 = 0
                _ = fi5
                _ = v6
                _ = v7
            }
            v1 = 65535
            _ = fi4
        }
    case 1:
        switch "a" {
        case "a":
            var v8 string = "a"
            _ = v8
        case "\n\t":
            delete(v2, 2147483648)
        case "foo bar":
            v3 = 1.5
        }
    }
    v1 += v1
    if true {
        v3 -= v3
        if (v1 == (^(v1))) {
            v9 := GenS1{F0: 24}
            v3 -= 0.1
            var v10 uint16 = 77
            _ = v10
            _ = v9
        }
        var v11 uint64 = 4294967295
        _ = v11
    }
    for fi12 := 0; fi12 < 20; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            fmt.Println(fi13, v3)
            _ = fi13
        }
        v14 := GenS1{F0: 127}
        _ = fi12
        _ = v14
    }
    var v15 uint8 = 0
    v16 := []int{-1, -2147483648, 850, -971}
    v3 -= v3
    v3 += ((1e10 * (3.14159 + v3)) + 0.0)
    v1 += 4294967295
    v3 = (-2.25 * 1e10)
    if len(v16) > 0 {
        v16[0] = -171
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
