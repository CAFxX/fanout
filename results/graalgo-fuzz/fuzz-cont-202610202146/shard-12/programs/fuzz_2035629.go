package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 int = 526
    var v2 uint8 = 255
    v3 := make(map[string]int)
    v4 := []uint32{42, 42, 4294967295}
    v2 = (v2 % ((0) | 1))
    var v5 uint = 4294967295
    v5 -= (v5 | (v5 % ((v5) | 1)))
    v5 += v5
    v6 := GenS1{F0: 0}
    if true {
        v7 := GenS2{F0: 683, F1: 65535}
        v5 -= (v5 & 1)
        var v8 string = ""
        _ = v7
        _ = v8
    } else {
        v9 := GenS2{F0: -330, F1: 37}
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        if (4294967295 < 4294967295) {
            fmt.Println(v1, v1)
        } else {
            var v11 uint8 = (v2 << (((v5 | 65535)) & 63))
            _ = v11
        }
        var v12 uint64 = 2
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 bool = ((v1 % ((fi10) | 1)) != v1)
            v14 = v14
            _ = fi13
            _ = v14
        }
        _ = fi10
        _ = v12
    }
    for fi15 := 0; fi15 < 1; fi15++ {
        var v16 float64 = 1.5
        v16 -= v16
        if (2147483647 <= 2147483647) {
            fmt.Println(fi15)
        } else {
            switch ("foo bar" + ("foo bar" + ("foo bar" + "foo bar"))) {
            case "hello":
                if true {
                    v17 := v3["foo bar"]
                    fmt.Println(fi15, fi15, fi15)
                    _ = v17
                } else {
                    v4 = append(v4, 2147483648)
                    fmt.Println(v2, v2, v1)
                }
            default:
                if false {
                    var v18 int32 = 878
                    var v19 uint32 = 2147483648
                    _ = v18
                    _ = v19
                }
            }
        }
        _ = fi15
        _ = v16
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
