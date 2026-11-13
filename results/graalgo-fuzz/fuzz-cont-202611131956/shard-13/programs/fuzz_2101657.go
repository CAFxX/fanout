package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int32 }
func genFn0() uint8 {
    var v4 int32 = 2147483647
    fmt.Println(v4)
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 32767
    var v3 bool = false
    for fi5 := 0; fi5 < 19; fi5++ {
        v3 = true
        _ = fi5
    }
    v6 := GenS1{F0: "", F1: 2, F2: -732}
    fmt.Println(v3, v2)
    var v7 string = ""
    var v8 int16 = v2
    _ = len(v1)
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: (v7 + "hello"), F1: 42, F2: 2147483647}
        fmt.Println(v2)
        _ = fi9
        _ = v10
    }
    var v11 uint8 = 255
    fmt.Println(v7)
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 bool = (1 >= 42)
        var v14 uint = 2
        var v15 float64 = (((((0.1 + 3.14159) * ((0.1 / 1.5) * 1e10)) * 3.14159) + 3.14159) - 0.0)
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    if v3 {
        v1[127] = ((v7 + "") + (v7 + "foo bar"))
        var v16 int16 = (v8 ^ v2)
        _ = v16
    }
    v7 = "foo bar"
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(v8)
        var v18 uint16 = 255
        _ = fi17
        _ = v18
    }
    fmt.Println(v11, v11, v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
