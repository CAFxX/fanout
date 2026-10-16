package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 int; F1 int8; F2 string }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]int)
    var v3 float64 = 1.5
    var v4 uint = 34
    var v5 int8 = 127
    for fi6 := 0; fi6 < 14; fi6++ {
        if false {
            v7 := GenS1{F0: "\n\t", F1: 743}
            _ = v7
        }
        v5 += v5
        _ = fi6
    }
    if true {
        switch -1 {
        case 127:
            if true {
                v8 := GenS1{F0: ("a" + (("foo bar" + "\n\t") + "hello")), F1: 42}
                _ = v8
            }
        default:
            var v9 uint8 = 0
            _ = v9
        }
        fmt.Println(v5, v4)
    }
    _ = len(v2)
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v4 = 28
            v3 += 3.14159
            _ = fi11
        }
        _ = fi10
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 int = (fi12 >> (((v4 << (((v4 * v4)) & 63))) & 63))
        v4 += (v4 - (v4 - v4))
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 uint64 = 878
            fmt.Println(fi12, v13, fi14)
            var v16 int8 = v5
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint8 = 1
    fmt.Println(v4, v5, v1)
    if true {
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS2{F0: fi18, F1: (v5 << (((v4 << ((v4) & 63))) & 63)), F2: ((("foo bar" + "\n\t") + "\n\t") + "a")}
            _ = fi18
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            if false {
                v2["\n\t"] = (fi20 % ((127) | 1))
                v17 += (v17 * (v17 + 1))
                v4 += v4
            }
            v21 := GenS3{F0: 686}
            _ = fi20
            _ = v21
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
