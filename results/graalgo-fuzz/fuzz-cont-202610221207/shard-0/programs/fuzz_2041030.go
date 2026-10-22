package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 uint32 = 42
    var v2 int64 = -1
    var v3 int64 = v2
    v4 := []int{797}
    var v5 float64 = (-2.25 + 1e10)
    var v6 float32 = float32(v5)
    fmt.Println(v6, v2)
    var v7 int = 42
    var v8 int16 = 421
    v9 := GenS1{F0: (v2 % (((v2 | 9223372036854775807)) | 1))}
    v10 := GenS2{F0: 32767}
    var v11 uint64 = 2
    v1 = 335
    v12 := GenS2{F0: v8}
    fmt.Println(v3, v1, v3)
    v2 += 820
    for fi13 := 0; fi13 < 20; fi13++ {
        if len(v4) > 0 {
            v4[0] = (fi13 ^ fi13)
        }
        var v14 int = 683
        fmt.Println(v2, fi13, v7)
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 string = "hello"
        if len(v4) > 0 {
            v4[0] = fi15
        }
        v4 = append(v4, fi15)
        _ = fi15
        _ = v16
    }
    if true {
        var v17 int32 = -672
        _ = v17
    } else {
        fmt.Println(v3, v7)
    }
    var v18 int = v7
    fmt.Println(v7, v8)
    v18 = v18
    switch (("a" + (("" + "foo bar") + ("hello" + "hello"))) + "foo bar") {
    case "":
        v4 = append(v4, 32768)
    case "a":
        fmt.Println(v5, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
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
