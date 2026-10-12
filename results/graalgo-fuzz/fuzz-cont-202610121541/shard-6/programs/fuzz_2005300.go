package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 int32; F2 uint }

func main() {
    var v1 uint8 = 1
    var v2 uint64 = 4294967295
    var v3 int64 = -1
    v4 := make(map[int]int)
    v5 := make(map[uint32]bool)
    v6 := make(map[string]string)
    v5[0] = (!(2 > 2))
    v2 = v2
    fmt.Println(v1, v2, v2)
    v7 := []int{127, 9223372036854775807}
    fmt.Println(v3, v2)
    var v8 uint64 = v2
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v1)
        var v10 string = (("" + "a") + "\n\t")
        _ = fi9
        _ = v10
    }
    v4[32768] = 32767
    var v11 int16 = 1
    var v12 bool = (-1 > 64)
    v13 := v5[575]
    switch 9223372036854775807 {
    case 127:
        if v13 {
            var v14 uint16 = 0
            if (v11 <= v11) {
                for fi15 := 0; fi15 < 15; fi15++ {
                    var v16 int8 = 100
                    _ = fi15
                    _ = v16
                }
                if v12 {
                    v13 = (v12 || v12)
                } else {
                    switch ("hello" + ("a" + ("foo bar" + "foo bar"))) {
                    case "a":
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v18 := GenS1{F0: v13}
                            v12 = v12
                            _ = fi17
                            _ = v18
                        }
                    default:
                        delete(v4, 128)
                    }
                    var v19 string = ("" + "")
                    _ = v19
                }
            } else {
                v13 = (v3 != v3)
            }
            v20 := GenS1{F0: v12}
            _ = v14
            _ = v20
        } else {
            v8 = (v8 & 2)
            v21 := GenS2{F0: v1, F1: 1, F2: 42}
            var v22 string = "\n\t"
            _ = v21
            _ = v22
        }
    default:
        fmt.Println(v8)
    }
    var v23 float32 = ((0.1 / ((-2.25 - 1e10) - 3.14159)) - 0.0)
    if v13 {
        v11 = -32768
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            fmt.Println(v23, v1)
            var v25 float64 = 3.14159
            _ = fi24
            _ = v25
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
