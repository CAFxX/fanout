package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
func genFn0() string {
    v5 := []int{-496, -380, 1, 32767}
    _ = len(v5)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v6 bool = (true || (true || (false && (18446744073709551615 > 42))))
    v5 = append(v5, -833)
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 100
    var v3 bool = ((("\n\t" + "a") + "foo bar") != "\n\t")
    var v4 uint64 = 2
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := v1[158]
        if (-1 != (fi7 / (((v8 | fi7)) | 1))) {
            var v9 uint64 = (v4 >> ((265) & 63))
            _ = v9
        }
        var v10 int64 = -351
        _ = fi7
        _ = v10
        _ = v8
    }
    delete(v1, 42)
    v3 = v3
    var v11 int8 = (v2 - 42)
    v12 := GenS1{F0: (v2 * (-(v2)))}
    if v3 {
        delete(v1, 2147483647)
        var v13 uint16 = 32767
        v14 := []int{9223372036854775807, 127, 4294967295, -1}
        _ = v13
        _ = v14
    }
    var v15 uint8 = 2
    if true {
        v16 := GenS1{F0: (v11 % ((v11) | 1))}
        _ = v16
    }
    var v17 bool = v3
    v17 = v3
    v18 := GenS2{F0: (v11 & (^(v11)))}
    var v19 uint64 = (v4 & v4)
    v3 = true
    for fi20 := 0; fi20 < 16; fi20++ {
        fmt.Println(v11)
        _ = fi20
    }
    v21 := []int{-58, 557}
    v22 := GenS1{F0: 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
