package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int8; F1 int32; F2 uint32 }

func main() {
    v1 := make(map[int]int)
    var v2 bool = true
    var v3 int64 = 42
    var v4 int64 = v3
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: 42, F1: "\n\t"}
        _ = fi5
        _ = v6
    }
    var v7 uint = 42
    v8 := GenS1{F0: v7, F1: ("\n\t" + "\n\t")}
    v9 := GenS1{F0: (v7 >> ((1) & 63)), F1: "hello"}
    if true {
        v3 = v4
        var v10 uint64 = 18446744073709551615
        _ = v10
    } else {
        v3 = (^(v4))
        v3 = (-(v4))
    }
    v11 := []int{1, -1, 127, 0}
    if false {
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v12 uint64 = 4294967295
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 float32 = -2.25
            v15 := GenS1{F0: (v7 * (^(v7))), F1: (("foo bar" + ("foo bar" + "hello")) + "")}
            var v16 int = 4294967295
            _ = fi13
            _ = v14
            _ = v15
            _ = v16
        }
        _ = v12
    } else {
        if (1 < 1) {
            v3 = (v4 << (((v7 / ((v7) | 1))) & 63))
            fmt.Println(v4)
            var v17 uint8 = 131
            _ = v17
        }
        v18 := GenS2{F0: 100, F1: 0, F2: 65535}
        v19 := GenS2{F0: int8(v4), F1: 2147483647, F2: 4294967295}
        _ = v18
        _ = v19
    }
    var v20 uint = (v7 * (v7 | 255))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
