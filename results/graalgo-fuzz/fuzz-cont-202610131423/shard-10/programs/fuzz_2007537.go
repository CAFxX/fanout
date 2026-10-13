package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }

func main() {
    var v1 int16 = -921
    v2 := make(map[uint32]int)
    var v3 int32 = -2147483648
    v4 := []string{}
    var v5 uint32 = 263
    for fi6 := 0; fi6 < 18; fi6++ {
        v1 = (v1 & v1)
        if len(v4) > 0 {
            v4[0] = ""
        }
        _ = fi6
    }
    var v7 uint8 = 43
    fmt.Println(v1, v1)
    v3 = (-(v3))
    if false {
        var v8 bool = true
        _ = v8
    } else {
        v9 := GenS1{F0: -281, F1: (^(v1))}
        _ = v9
    }
    v10 := GenS1{F0: 128, F1: -466}
    v11 := v2[2]
    v12 := GenS1{F0: (v11 % (((v11 / ((v11) | 1))) | 1)), F1: -1}
    if (!true) {
        var v13 int64 = 1
        v5 = v5
        _ = v13
    }
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 7; fi15++ {
        if true {
            _ = len(v2)
            v4 = append(v4, "foo bar")
            if len(v4) > 0 {
                v4[0] = "a"
            }
        } else {
            fmt.Println(v14)
        }
        _ = fi15
    }
    var v16 string = ("\n\t" + "a")
    var v17 string = v16
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
