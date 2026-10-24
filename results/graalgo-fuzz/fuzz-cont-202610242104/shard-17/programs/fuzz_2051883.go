package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16) bool {
    var v5 uint = 4294967295
    v5 += v5
    p0 -= (p0 << ((v5) & 63))
    if true {
        var v6 uint64 = 4294967295
        p0 += (p0 * p0)
        _ = v6
    }
    fmt.Println(v5, v5)
    _ = p0
    _ = v5
    return (v5 != v5)
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    var v4 uint8 = 255
    var v7 uint8 = 189
    if (false && true) {
        switch -160 {
        case 9223372036854775807:
            var v8 int32 = -1
            _ = v8
        }
        var v9 int = 698
        _ = v9
    } else {
        v10 := GenS1{F0: 1000}
        _ = v10
    }
    v11 := v3[32767]
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(fi12, v4)
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(fi12, v11)
            _ = fi13
        }
        delete(v2, "")
        _ = fi12
    }
    if false {
        var v14 int64 = 1
        v11 += v11
        _ = v14
    }
    v15 := GenS1{F0: 1000}
    v16 := GenS1{F0: 1}
    v4 = 1
    v17 := GenS1{F0: 1000}
    for fi18 := 0; fi18 < 2; fi18++ {
        switch fi18 {
        case 120:
            v19 := GenS1{F0: -577}
            _ = v19
        default:
            fmt.Println(v1, v7)
        }
        _ = fi18
    }
    v20 := GenS1{F0: 42}
    v7 += 127
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}
