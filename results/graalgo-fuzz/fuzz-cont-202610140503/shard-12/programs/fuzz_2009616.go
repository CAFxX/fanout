package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint32 }

func main() {
    var v1 uint = 42
    var v2 int32 = -1
    v3 := make(map[string]string)
    var v4 int32 = (^(v2))
    delete(v3, "")
    if (2 == 0) {
        if (v1 == v1) {
            v4 -= v4
            var v5 int = 32768
            v5 -= 32767
            _ = v5
        }
        var v6 uint64 = 2
        for fi7 := 0; fi7 < 7; fi7++ {
            v6 = (^(v6))
            _ = fi7
        }
        _ = v6
    } else {
        fmt.Println(v4)
        if false {
            if (true && (false && false)) {
                var v8 uint64 = 0
                _ = v8
            }
            v9 := GenS1{F0: 42, F1: -1, F2: 2147483647}
            _ = v9
        } else {
            var v10 string = ("" + "a")
            fmt.Println(v2, v10)
            var v11 uint8 = 161
            _ = v10
            _ = v11
        }
        var v12 uint64 = 427
        _ = v12
    }
    var v13 int16 = 42
    v3["a"] = ("\n\t" + "")
    var v14 uint32 = 622
    v15 := GenS1{F0: 0, F1: v13, F2: 326}
    v2 = v2
    v16 := GenS1{F0: 61, F1: v13, F2: v14}
    if false {
        var v17 uint64 = 1
        _ = v17
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 int8 = 100
            var v20 uint8 = 127
            fmt.Println(v2, v13, v19)
            _ = fi18
            _ = v19
            _ = v20
        }
        v3["a"] = "foo bar"
        v21 := GenS1{F0: int8(v4), F1: v13, F2: (v14 ^ 0)}
        _ = v21
    }
    _ = len(v3)
    if false {
        switch "hello" {
        case "\n\t":
            v4 = -157
        case "foo bar":
            var v22 int64 = 0
            _ = v22
        }
        v23 := GenS1{F0: 42, F1: v13, F2: (v14 / ((v14) | 1))}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
