package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (false || true)
    var v3 float64 = 0.1
    var v4 int16 = -32768
    var v5 uint16 = 460
    v6 := make(map[string]string)
    fmt.Println(v3, v3, v3)
    var v7 int32 = 2147483647
    fmt.Println(v2)
    var v8 int16 = (v4 + v4)
    if (!(v2 || true)) {
        v9 := GenS1{F0: 255, F1: -32768, F2: ("\n\t" + "foo bar")}
        switch ("\n\t" + "hello") {
        case "\n\t":
            fmt.Println(v2)
        case "a":
            v10 := GenS1{F0: 65535, F1: (v4 % ((v4) | 1)), F2: "a"}
            _ = v10
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            fmt.Println(v3, v5)
            var v12 uint32 = 65535
            _ = fi11
            _ = v12
        }
        _ = v9
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(v2)
        v1[0] = (fi13 * (fi13 - (-(fi13))))
        var v14 uint64 = 276
        _ = fi13
        _ = v14
    }
    fmt.Println(v7, v3)
    var v15 int8 = 85
    for fi16 := 0; fi16 < 14; fi16++ {
        if v2 {
            if v2 {
                var v17 int16 = 1000
                v18 := []int{}
                v8 = v8
                _ = v17
                _ = v18
            } else {
                v3 -= 0.1
                var v19 int64 = 0
                v19 += (^(v19))
                _ = v19
            }
            v15 -= int8(v7)
            v15 = v15
        }
        if v2 {
            v20 := GenS1{F0: 255, F1: (-(v8)), F2: "hello"}
            _ = v20
        } else {
            var v21 bool = true
            _ = v21
        }
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
