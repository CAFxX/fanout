package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint8) int64 {
    p0 += (p0 - 190)
    switch (("" + (((("foo bar" + ("\n\t" + "a")) + "") + "foo bar") + ("hello" + "\n\t"))) + "\n\t") {
    case "":
        for fi5 := 0; fi5 < 9; fi5++ {
            p0 = p0
            _ = fi5
        }
    }
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -438
}

func main() {
    var v1 int32 = 2147483647
    v2 := []string{}
    v3 := make(map[string]string)
    v4 := []int{312, 2147483647, 32768}
    v6 := v3["a"]
    v7 := GenS1{F0: 2, F1: 2}
    var v8 uint = 0
    for fi9 := 0; fi9 < 14; fi9++ {
        v8 -= 42
        v1 = (v1 - (v1 | (v1 - -2147483648)))
        _ = fi9
    }
    var v10 float32 = float32(0.0)
    v11 := GenS2{F0: 1, F1: 128}
    v10 = -2.25
    v10 = float32(0.1)
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v6, v10, v8)
    default:
        delete(v3, v6)
    }
    v12 := GenS3{F0: 1, F1: 0}
    if false {
        var v13 int32 = (v1 | 1)
        if false {
            v13 -= (v1 | 0)
            v2 = append(v2, "foo bar")
        } else {
            v14 := GenS2{F0: 0, F1: 19}
            _ = v14
        }
        _ = v13
    }
    fmt.Println(v8, v1)
    v1 -= v1
    var v15 uint8 = 145
    v8 = (v8 << (((v8 / ((v8) | 1))) & 63))
    var v16 bool = (0 == 444)
    var v17 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
