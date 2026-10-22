package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
type GenS2 struct { F0 int; F1 int }

func main() {
    v1 := []uint32{65535, 1, 472}
    var v2 uint32 = 784
    var v3 bool = true
    v4 := GenS1{F0: 42, F1: "", F2: ("\n\t" + ("a" + "hello"))}
    var v5 uint8 = 127
    var v6 int16 = -1
    var v7 string = "foo bar"
    fmt.Println(v6)
    var v8 int = 342
    v9 := GenS1{F0: 94, F1: v7, F2: v7}
    for fi10 := 0; fi10 < 19; fi10++ {
        if true {
            v11 := GenS1{F0: -62, F1: v7, F2: ("foo bar" + v7)}
            v5 += (v5 ^ (^(v5)))
            _ = v11
        } else {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint = 812
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v14 uint8 = (v5 >> ((42) & 63))
                if (42 == 490) {
                    var v15 int = (fi12 >> (((v13 & 255)) & 63))
                    switch "foo bar" {
                    case "hello":
                        v16 := GenS1{F0: 100, F1: (("foo bar" + v7) + v7), F2: v7}
                        _ = v16
                    default:
                        v8 -= 0
                    }
                    _ = v15
                } else {
                    v2 -= 0
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            var v17 uint64 = 4294967295
            var v18 uint16 = 255
            _ = v17
            _ = v18
        }
        _ = fi10
    }
    var v19 float32 = (3.14159 - 0.1)
    if v3 {
        if false {
            v20 := GenS2{F0: (v8 * v8), F1: (v8 - v8)}
            _ = v20
        }
    } else {
        var v21 int32 = 0
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 int64 = 0
            _ = fi22
            _ = v23
        }
        fmt.Println(v8)
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
