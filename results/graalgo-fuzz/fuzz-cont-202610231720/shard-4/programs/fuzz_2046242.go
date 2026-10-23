package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 bool = (!true)
    var v3 int = 786
    var v4 float32 = 1e10
    var v5 bool = (2147483647 >= 1)
    var v6 int32 = 0
    v7 := GenS1{F0: 179, F1: 18446744073709551615}
    switch "a" {
    case "hello":
        v6 = 42
    case "foo bar":
        switch (v3 % ((32767) | 1)) {
        case 487:
            switch v3 {
            case -328:
                delete(v1, v3)
            case 2147483647:
                fmt.Println(v3, v2)
            case -1:
                if (false || (132 > 4294967295)) {
                    var v8 int = (v3 + v3)
                    var v9 uint32 = 42
                    _ = v8
                    _ = v9
                }
            }
        }
    }
    var v10 float32 = (float32(1e10) - 0.1)
    v1[v3] = 891
    v3 -= 127
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 float32 = v4
        var v13 string = "a"
        _ = fi11
        _ = v12
        _ = v13
    }
    if (42 < 1) {
        switch "hello" {
        case "hello":
            var v14 int = (v3 | v3)
            _ = v14
        default:
            v15 := GenS1{F0: 0, F1: 1}
            _ = v15
        }
    } else {
        v16 := []int{0, (v3 - (v3 ^ 1))}
        if ((v6 + v6) >= 1) {
            fmt.Println(v5, v3)
            v17 := GenS1{F0: 42, F1: 2}
            _ = v17
        }
        _ = len(v1)
        _ = v16
    }
    var v18 int8 = int8(v6)
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint = 4294967295
            v10 -= (0.0 + v10)
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
