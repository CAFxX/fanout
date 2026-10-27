package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    var v1 int16 = 1
    var v2 uint64 = 0
    v3 := make(map[string]bool)
    var v4 float64 = ((3.14159 - 0.1) + 1e10)
    var v5 int = -484
    switch (^(v5)) {
    case 1:
        var v6 string = ("a" + ((("foo bar" + "") + "") + "\n\t"))
        _ = v6
    case 127:
        switch ("a" + "a") {
        case "hello":
            var v7 uint = 65535
            _ = v7
        case "foo bar":
            var v8 int = (v5 * 32768)
            _ = v8
        case "":
            v9 := GenS1{F0: 163}
            _ = v9
        default:
            for fi10 := 0; fi10 < 19; fi10++ {
                if false {
                    var v11 int32 = 42
                    var v12 uint32 = 2147483647
                    _ = v11
                    _ = v12
                } else {
                    v4 += v4
                }
                v13 := GenS1{F0: 0}
                var v14 int64 = -9223372036854775808
                _ = fi10
                _ = v13
                _ = v14
            }
        }
    default:
        v15 := GenS2{F0: 733, F1: int64(v2)}
        _ = v15
    }
    v5 = (v5 >> ((42) & 63))
    var v16 int = (v5 & (v5 << ((2) & 63)))
    v17 := v3[((("foo bar" + "\n\t") + ("\n\t" + ("" + ("hello" + (("" + ("hello" + "hello")) + ("foo bar" + "hello")))))) + "hello")]
    v18 := GenS2{F0: -1, F1: 9223372036854775807}
    var v19 int64 = 1
    v20 := GenS2{F0: v19, F1: v19}
    v21 := GenS2{F0: v19, F1: v19}
    var v22 float32 = 0.0
    var v23 float32 = v22
    var v24 int16 = (v1 / (((^(v1))) | 1))
    v25 := []int{(v16 ^ 9223372036854775807), (v5 - v16), v5}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
