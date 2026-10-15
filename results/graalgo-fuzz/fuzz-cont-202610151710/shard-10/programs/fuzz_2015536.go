package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }

func main() {
    v1 := []uint32{0}
    v2 := []string{"\n\t", ""}
    var v3 uint32 = 44
    v2 = append(v2, "foo bar")
    var v4 float32 = (float32(0.1) + -2.25)
    if true {
        for fi5 := 0; fi5 < 11; fi5++ {
            fmt.Println(fi5, v3, v3)
            var v6 uint16 = 2
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 19; fi7++ {
            if (42 == 9223372036854775807) {
                var v8 int = (fi7 - fi7)
                if len(v2) > 0 {
                    v2[0] = ("" + "")
                }
                v8 = 2147483647
                _ = v8
            } else {
                var v9 uint16 = 176
                _ = v9
            }
            var v10 uint64 = 391
            _ = fi7
            _ = v10
        }
    }
    v11 := GenS1{F0: "\n\t", F1: 1}
    fmt.Println(v4)
    var v12 int = 2147483647
    switch 127 {
    case 128:
        for fi13 := 0; fi13 < 19; fi13++ {
            if true {
                _ = len(v1)
            }
            _ = len(v1)
            _ = fi13
        }
    }
    v12 = (v12 / (((v12 + (v12 - v12))) | 1))
    if true {
        fmt.Println(v3, v3)
        fmt.Println(v3)
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            if false {
                v12 = 9223372036854775807
                v12 = v12
            }
            v15 := GenS1{F0: "hello", F1: -234}
            fmt.Println(v3)
            _ = fi14
            _ = v15
        }
        var v16 uint8 = 253
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
