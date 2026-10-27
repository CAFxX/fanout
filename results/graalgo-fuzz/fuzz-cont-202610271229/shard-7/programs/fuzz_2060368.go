package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
func genFn0(p0 int64, p1 uint8) uint64 {
    fmt.Println(p0, p1)
    if true {
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := GenS1{F0: 947, F1: 0, F2: p1}
            _ = fi5
            _ = v6
        }
    } else {
        p0 = (^(p0))
        if true {
            var v7 int32 = int32(p1)
            _ = v7
        } else {
            var v8 uint8 = (p1 + p1)
            _ = v8
        }
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int32 = 2147483647
    v2 := []int{0, 32768}
    v3 := make(map[string]int)
    var v4 uint64 = 1
    for fi9 := 0; fi9 < 1; fi9++ {
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v1 += (v1 | 331)
        case "a":
            var v10 uint64 = (v4 >> ((399) & 63))
            _ = v10
        case "\n\t":
            if (false || (1 == 589)) {
                var v11 string = "\n\t"
                _ = v11
            } else {
                v4 = v4
                v4 += v4
                v12 := GenS1{F0: 1, F1: 0, F2: 128}
                _ = v12
            }
        }
        _ = fi9
    }
    fmt.Println(v4, v1)
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int16 = 1000
        v3[("\n\t" + "")] = (fi13 | (fi13 & 354))
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        var v16 uint8 = 127
        fmt.Println(fi15)
        v17 := GenS1{F0: int64(v16), F1: 4294967295, F2: 154}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 int32 = (v1 - v1)
    v19 := GenS1{F0: -1, F1: 4294967295, F2: 0}
    if true {
        v18 = -1
        var v20 float32 = (float32(1.5) - 3.14159)
        _ = v20
    }
    v21 := GenS1{F0: 9223372036854775807, F1: 440, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
