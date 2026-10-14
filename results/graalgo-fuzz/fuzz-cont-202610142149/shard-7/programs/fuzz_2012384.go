package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 uint8 }

func main() {
    v1 := []string{"", "a", (("foo bar" + (("foo bar" + "hello") + ("foo bar" + ("foo bar" + "hello")))) + "foo bar")}
    v2 := make(map[uint32]string)
    var v3 int32 = 42
    var v4 uint16 = 982
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(v4, fi5)
        _ = fi5
    }
    v3 -= v3
    if true {
        v6 := GenS1{F0: 1, F1: v3, F2: 1}
        fmt.Println(v3)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v6
    } else {
        var v7 float64 = float64(1.5)
        v8 := GenS1{F0: 227, F1: v3, F2: 255}
        _ = v7
        _ = v8
    }
    var v9 float32 = (1e10 - 1.5)
    for fi10 := 0; fi10 < 3; fi10++ {
        if (!true) {
            v9 = 1.5
        } else {
            var v11 bool = true
            var v12 string = ("hello" + "\n\t")
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    var v13 uint64 = 2
    if false {
        v3 = (-(v3))
        var v14 bool = ("\n\t" > "hello")
        _ = v14
    } else {
        var v15 float32 = (3.14159 * float32(3.14159))
        _ = v15
    }
    var v16 uint8 = 1
    v17 := GenS1{F0: -1, F1: 42, F2: (v16 - (v16 ^ (v16 % ((v16) | 1))))}
    fmt.Println(v9, v13, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
