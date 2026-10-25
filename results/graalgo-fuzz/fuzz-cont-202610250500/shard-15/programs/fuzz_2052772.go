package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint16 = 2
    var v2 string = "a"
    v3 := make(map[int]uint32)
    var v4 int = 127
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 uint32 = 65535
        _ = fi5
        _ = v6
    }
    var v7 int8 = -128
    v8 := GenS1{F0: -487}
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int8 = (v7 / ((int8(v4)) | 1))
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int32 = -825
            _ = fi11
            _ = v12
        }
        fmt.Println(v4, v7)
        _ = fi9
        _ = v10
    }
    v7 = 122
    var v13 uint16 = v1
    fmt.Println(v4, v4, v13)
    var v14 uint16 = 830
    var v15 uint = 2
    v13 += v1
    v1 += (v14 / (((v13 + (v13 + v13))) | 1))
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
