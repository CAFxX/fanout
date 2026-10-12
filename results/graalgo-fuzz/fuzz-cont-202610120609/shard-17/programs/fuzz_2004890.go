package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 uint8, p2 uint8) string {
    p0 -= p0
    var v7 bool = (2147483648 > uint32(p2))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return "a"
}

func main() {
    var v1 int8 = 100
    v2 := make(map[uint32]bool)
    var v3 int64 = int64(v1)
    var v4 float64 = 0.0
    v5 := make(map[uint32]string)
    v6 := make(map[int]bool)
    delete(v5, 815)
    var v8 int = 127
    if ((-(v3)) > (^(v3))) {
        v3 = (v3 - (v3 >> ((164) & 63)))
        switch "foo bar" {
        case "\n\t":
            v4 = v4
        case "":
            v8 -= v8
        case "hello":
            v9 := GenS1{F0: 4294967295, F1: v1}
            _ = v9
        }
    } else {
        fmt.Println(v4)
        delete(v6, (v8 ^ -53))
        var v10 uint64 = 1
        _ = v10
    }
    v4 -= v4
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 uint64 = 42
        fmt.Println(fi11)
        _ = fi11
        _ = v12
    }
    var v13 int8 = v1
    v14 := GenS2{F0: 737}
    var v15 float64 = (((v4 + v4) - 1.5) - 1e10)
    for fi16 := 0; fi16 < 1; fi16++ {
        v4 += v4
        fmt.Println(v3, fi16, v1)
        _ = fi16
    }
    if ((2147483647 <= 1) && (!((v1 >> ((30) & 63)) != (v13 / (((v1 / ((121) | 1))) | 1))))) {
        var v17 int32 = 1
        _ = v17
    } else {
        var v18 uint16 = 157
        _ = v18
    }
    fmt.Println(v4)
    var v19 int64 = v3
    v2[42] = true
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
