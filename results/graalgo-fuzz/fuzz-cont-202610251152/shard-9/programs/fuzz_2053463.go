package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int; F1 string; F2 int64 }
func genFn0(p0 bool, p1 string) uint8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := GenS1{F0: 32767}
        _ = fi7
        _ = v8
    }
    p1 = p1
    v9 := GenS1{F0: 32767}
    _ = p0
    _ = p1
    _ = v9
    return 139
}

func main() {
    var v1 bool = (32767 >= -395)
    var v2 uint64 = 4294967295
    v3 := make(map[int]bool)
    v4 := []uint32{uint32(v2), 652}
    var v5 int16 = -32768
    v6 := make(map[string]int)
    var v10 int = 288
    if v1 {
        var v11 uint16 = 0
        v12 := GenS1{F0: uint16(v2)}
        _ = len(v6)
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 string = "foo bar"
            _ = fi14
            _ = v15
        }
        v16 := GenS2{F0: (fi13 & 2147483647), F1: ("\n\t" + ("\n\t" + ("foo bar" + "a"))), F2: -737}
        _ = fi13
        _ = v16
    }
    v17 := GenS2{F0: (^(v10)), F1: "foo bar", F2: 1}
    if (v1 && v1) {
        v2 -= (^(v2))
        var v18 uint8 = 103
        v1 = (!v1)
        _ = v18
    } else {
        var v19 uint32 = 0
        _ = v19
    }
    var v20 int32 = 728
    var v21 uint16 = 607
    var v22 uint16 = (v21 << ((42) & 63))
    v20 -= (v20 % ((int32(v21)) | 1))
    var v23 uint16 = 65535
    v24 := v6[(("foo bar" + "") + "foo bar")]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
