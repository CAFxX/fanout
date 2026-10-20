package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 uint8; F1 uint64; F2 int64 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]bool)
    v1 -= v1
    for fi5 := 0; fi5 < 12; fi5++ {
        v1 = (v1 ^ 65535)
        fmt.Println(v1, fi5)
        _ = fi5
    }
    if ((3.14159 * 3.14159) <= 1e10) {
        var v6 uint32 = v1
        _ = v6
    } else {
        v7 := GenS1{F0: (false || (false || false)), F1: 4294967295}
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{fi8, fi8}
            var v10 int32 = -2147483648
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: 241, F1: 4294967295, F2: -844}
        _ = v11
        _ = v7
    }
    var v12 uint64 = 875
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 float32 = (1e10 / -2.25)
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    v3[(("foo bar" + ("foo bar" + "foo bar")) + "a")] = "\n\t"
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 int16 = 1000
        v16 += (v16 * (v16 | v16))
        _ = fi15
        _ = v16
    }
    v1 += 4294967295
    var v17 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
