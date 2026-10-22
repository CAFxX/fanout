package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = 919
    var v2 int32 = 681
    var v3 int = 32768
    var v4 uint64 = 956
    var v5 uint16 = 233
    switch "a" {
    case "foo bar":
        v1 = (v1 + v1)
    case "hello":
        var v6 int64 = (v1 >> ((65535) & 63))
        _ = v6
    }
    if ((3.14159 * 3.14159) > -2.25) {
        v7 := GenS1{F0: 178}
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint64 = uint64(v5)
            v1 += 610
            v10 := GenS2{F0: 2147483648}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v7
    }
    var v11 int64 = (v1 << ((646) & 63))
    var v12 uint32 = 1
    for fi13 := 0; fi13 < 6; fi13++ {
        v14 := []int{127, fi13, 4294967295, v3}
        _ = fi13
        _ = v14
    }
    var v15 int16 = 0
    v11 += v11
    v16 := []int{4294967295, 9223372036854775807}
    v16 = append(v16, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
