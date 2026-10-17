package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 string }

func main() {
    var v1 int = -140
    var v2 bool = true
    v3 := []string{"\n\t"}
    v4 := make(map[uint32]string)
    v1 += (v1 + v1)
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: -93}
            v2 = v2
            _ = len(v3)
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 uint16 = 1
    v9 := GenS2{F0: -1, F1: ""}
    if true {
        fmt.Println(v8, v8)
        var v10 bool = v2
        _ = v10
    } else {
        fmt.Println(v2)
    }
    if (1.5 == -2.25) {
        v1 = v1
    }
    if v2 {
        v2 = v2
    }
    fmt.Println(v1, v8)
    for fi11 := 0; fi11 < 1; fi11++ {
        v3 = append(v3, ("\n\t" + "foo bar"))
        var v12 string = "foo bar"
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 78
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
