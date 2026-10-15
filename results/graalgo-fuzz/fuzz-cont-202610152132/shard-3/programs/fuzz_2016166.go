package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint }
type GenS2 struct { F0 int64 }
func genFn0() uint64 {
    var v6 string = "foo bar"
    if (!(((32767 < 867) && (703 != 574)) || true)) {
        fmt.Println(v6)
        v6 = v6
    } else {
        var v7 int = -9223372036854775808
        _ = v7
    }
    _ = v6
    return 18446744073709551615
}

func main() {
    v1 := []uint32{64}
    v2 := []string{"\n\t", "a", "foo bar", ("foo bar" + "foo bar")}
    var v3 string = ""
    v4 := make(map[string]int)
    v5 := make(map[int]string)
    for fi8 := 0; fi8 < 3; fi8++ {
        fmt.Println(fi8, fi8)
        _ = fi8
    }
    v9 := GenS1{F0: v3, F1: 65535, F2: 2}
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(fi10, fi10)
        v3 = v3
        _ = fi10
    }
    v11 := GenS2{F0: 42}
    v5[4294967295] = (v3 + (((v3 + v3) + "a") + v3))
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        switch fi13 {
        case 2147483647:
            var v14 bool = (!true)
            _ = v14
        default:
            v3 = v3
        }
        _ = fi13
    }
    fmt.Println(v3)
    v15 := v5[42]
    if false {
        _ = len(v5)
    }
    v16 := GenS2{F0: -936}
    fmt.Println(v3, v3)
    v15 = ((v15 + ("hello" + ("hello" + v3))) + v15)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
