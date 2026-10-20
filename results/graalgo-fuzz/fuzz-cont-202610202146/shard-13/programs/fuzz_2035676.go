package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = -1
    var v3 int32 = -2147483648
    v4 := []string{(("hello" + "hello") + ("foo bar" + ("a" + "\n\t"))), "", "foo bar", "\n\t"}
    var v5 uint16 = 490
    fmt.Println(v3, v3, v3)
    v2 -= 114
    fmt.Println(v5)
    v2 = -1
    if (!((v3 >> ((42) & 63)) > -1)) {
        v6 := GenS1{F0: ""}
        _ = v6
    }
    v2 -= v2
    v7 := v1[772]
    var v8 uint8 = 76
    var v9 string = "hello"
    _ = len(v1)
    if true {
        var v10 float64 = 0.0
        _ = v10
    } else {
        for fi11 := 0; fi11 < 5; fi11++ {
            var v12 int8 = (v2 | -1)
            _ = fi11
            _ = v12
        }
    }
    var v13 int64 = 1
    switch (v7 & (v7 & -790)) {
    case 127:
        var v14 uint32 = 153
        _ = v14
    case 9223372036854775807:
        v15 := GenS1{F0: "\n\t"}
        _ = v15
    case -427:
        v2 = v2
    default:
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: v9}
            var v18 int16 = -32768
            v3 += v3
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    v13 += 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
