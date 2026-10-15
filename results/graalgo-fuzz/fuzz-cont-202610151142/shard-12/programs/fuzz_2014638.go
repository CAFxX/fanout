package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }

func main() {
    var v1 float32 = 0.0
    v2 := []int{}
    var v3 int8 = -1
    v4 := make(map[string]int)
    var v5 int32 = 0
    var v6 string = "hello"
    v2 = append(v2, 32767)
    var v7 string = ((v6 + ("" + ("\n\t" + (v6 + ("" + v6))))) + v6)
    var v8 int64 = 42
    v5 += (v5 - (-(v5)))
    if false {
        v9 := GenS1{F0: 127, F1: 423}
        if len(v2) > 0 {
            v2[0] = 32768
        }
        var v10 uint8 = 255
        _ = v10
        _ = v9
    } else {
        var v11 uint8 = 1
        var v12 float32 = v1
        v13 := GenS1{F0: 459, F1: (-(v8))}
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 int64 = (v8 + v8)
    for fi15 := 0; fi15 < 16; fi15++ {
        if len(v2) > 0 {
            v2[0] = 941
        }
        var v16 uint16 = 459
        if true {
            for fi17 := 0; fi17 < 8; fi17++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v18 uint16 = 255
                var v19 float32 = 0.1
                switch 1 {
                case 32768:
                    v6 = "\n\t"
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint32 = 42
            _ = v20
        } else {
            v21 := GenS1{F0: 32768, F1: 1}
            fmt.Println(v16)
            _ = v21
        }
        _ = fi15
        _ = v16
    }
    _ = len(v2)
    delete(v4, ((v7 + "foo bar") + v7))
    var v22 int32 = 42
    fmt.Println(v1, v3, v22)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
