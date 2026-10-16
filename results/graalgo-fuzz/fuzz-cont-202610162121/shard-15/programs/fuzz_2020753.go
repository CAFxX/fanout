package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 bool = false
    v2 := []string{"\n\t", (("\n\t" + "foo bar") + "")}
    v3 := make(map[uint32]bool)
    var v4 string = "\n\t"
    var v5 int16 = -140
    var v6 int16 = 0
    var v7 uint64 = 2
    var v8 uint16 = 255
    var v9 int16 = (v5 >> ((1) & 63))
    for fi10 := 0; fi10 < 1; fi10++ {
        switch (fi10 + fi10) {
        case 32768:
            fmt.Println(fi10, v1)
        case 2147483647:
            var v11 int32 = 42
            _ = v11
        default:
            var v12 float64 = ((-2.25 * 0.0) + -2.25)
            _ = v12
        }
        v9 = (v6 ^ v9)
        switch (v4 + v4) {
        case "\n\t":
            switch fi10 {
            case -1:
                switch (fi10 * (fi10 + fi10)) {
                case 32768:
                    var v13 int64 = -9223372036854775808
                    _ = v13
                }
            }
        default:
            var v14 string = "hello"
            _ = v14
        }
        _ = fi10
    }
    v6 = (v6 ^ (v9 % ((1000) | 1)))
    v15 := GenS1{F0: -1}
    v3[926] = (true && (42 != int8(v5)))
    v7 -= uint64(v8)
    if (31 == int8(v8)) {
        var v16 float32 = 0.1
        v17 := GenS1{F0: 1000}
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: v9}
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
