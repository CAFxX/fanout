package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 bool; F2 uint }

func main() {
    var v1 int8 = 42
    var v2 float64 = 1.5
    v3 := make(map[uint32]string)
    var v4 uint = 18446744073709551615
    var v5 uint64 = 2
    v6 := []string{"\n\t", "\n\t", "a", "a"}
    v1 = v1
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 float64 = float64((3.14159 * float32(v2)))
    v5 += 1
    fmt.Println(v1, v5)
    fmt.Println(v2)
    v8 := GenS1{F0: 0}
    for fi9 := 0; fi9 < 2; fi9++ {
        switch "a" {
        case "foo bar":
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := v3[4294967295]
                var v12 uint = uint(v5)
                var v13 uint32 = 42
                if true {
                    v6 = append(v6, (v11 + "foo bar"))
                    v14 := GenS2{F0: (-1 >= 42), F1: false, F2: v4}
                    _ = v14
                } else {
                    _ = len(v3)
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
        }
        if false {
            v4 = v4
            var v15 int16 = 32767
            fmt.Println(v1, v7, v4)
            _ = v15
        } else {
            v16 := GenS1{F0: 128}
            for fi17 := 0; fi17 < 2; fi17++ {
                fmt.Println(v4, v2, v4)
                _ = fi17
            }
            _ = v16
        }
        v18 := GenS2{F0: (!("foo bar" != ("" + ""))), F1: true, F2: v4}
        _ = fi9
        _ = v18
    }
    if false {
        var v19 float64 = float64(0.0)
        _ = v19
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
