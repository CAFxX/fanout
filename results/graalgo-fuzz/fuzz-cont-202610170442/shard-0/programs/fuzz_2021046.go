package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int8 }

func main() {
    v1 := []string{("foo bar" + "hello")}
    var v2 uint = 1
    var v3 uint32 = 2147483648
    var v4 uint = 804
    v5 := []uint32{771, (v3 - (^(v3)))}
    v6 := make(map[uint32]int)
    var v7 int8 = -128
    v8 := GenS1{F0: -44, F1: v7, F2: (v3 / (((^(v3))) | 1))}
    var v9 uint = v4
    fmt.Println(v7)
    v4 = (^(v9))
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int8 = v7
        _ = fi10
        _ = v11
    }
    v12 := GenS2{F0: 42, F1: 82}
    for fi13 := 0; fi13 < 6; fi13++ {
        switch "" {
        case "foo bar":
            if (fi13 > fi13) {
                fmt.Println(v7)
                v9 = (^(v9))
                var v14 int16 = -32768
                _ = v14
            } else {
                for fi15 := 0; fi15 < 18; fi15++ {
                    delete(v6, (v3 >> ((42) & 63)))
                    _ = fi15
                }
                v2 = (^(v9))
            }
        }
        _ = fi13
    }
    if (!(9223372036854775807 != 713)) {
        var v16 string = ("a" + "")
        v2 -= v2
        _ = v16
    }
    var v17 float32 = -2.25
    var v18 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
