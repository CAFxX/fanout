package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    var v1 float64 = (0.1 - -2.25)
    var v2 int = -421
    var v3 float64 = v1
    var v4 uint8 = 42
    var v5 int16 = 272
    v6 := make(map[uint32]int)
    v1 -= 1.5
    fmt.Println(v5, v4)
    v7 := v6[0]
    if (!(-9223372036854775808 < 9223372036854775807)) {
        if false {
            var v8 int32 = 0
            v8 -= int32(v5)
            v9 := GenS1{F0: v7, F1: v4, F2: int8(v8)}
            _ = v8
            _ = v9
        }
        v7 -= v2
    }
    fmt.Println(v4)
    if (-1 > -1) {
        if true {
            var v10 bool = ((("hello" + "foo bar") + "\n\t") != ("" + "foo bar"))
            _ = v10
        } else {
            var v11 int64 = 0
            var v12 int8 = 0
            _ = len(v6)
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 2; fi13++ {
            v4 = v4
            switch (("a" + "") + "hello") {
            case "hello":
                fmt.Println(v7)
            }
            _ = fi13
        }
        if (false || false) {
            delete(v6, 0)
        }
    } else {
        var v14 float32 = ((1e10 - float32(((1.5 * v3) - (v1 - ((-2.25 * ((v3 - v3) - 3.14159)) - ((0.0 * ((v3 + 1e10) * v1)) * 1.5)))))) / 1e10)
        switch -289 {
        case 32767:
            v15 := GenS1{F0: (-(v2)), F1: 0, F2: 42}
            _ = v15
        case 241:
            var v16 float64 = 0.1
            _ = v16
        case -1:
            var v17 int64 = 42
            _ = v17
        default:
            v7 = v2
        }
        var v18 uint8 = (v4 ^ 36)
        _ = v14
        _ = v18
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 uint64 = 18446744073709551615
        var v21 int64 = 42
        v22 := GenS2{F0: 18446744073709551615, F1: 873, F2: false}
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
