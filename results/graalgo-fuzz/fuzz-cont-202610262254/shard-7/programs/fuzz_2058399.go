package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 uint16 = 42
    var v5 int32 = 1
    var v6 uint16 = (v4 << ((65535) & 63))
    if true {
        fmt.Println(v5)
    }
    fmt.Println(v4, v6)
    fmt.Println(v6, v4)
    v7 := GenS1{F0: 482}
    v6 = v6
    if false {
        var v8 uint8 = 0
        if false {
            delete(v3, -1)
            var v9 bool = ("foo bar" > (("foo bar" + "foo bar") + ""))
            _ = v9
        } else {
            v10 := GenS1{F0: 42}
            fmt.Println(v8, v4)
            var v11 float64 = (0.0 * ((1.5 + 3.14159) - ((-2.25 * 3.14159) - (3.14159 + 3.14159))))
            _ = v10
            _ = v11
        }
        var v12 string = "foo bar"
        _ = v12
        _ = v8
    } else {
        v4 = 42
    }
    v13 := GenS1{F0: 0}
    switch -730 {
    case 4294967295:
        v6 = (v6 & v4)
    }
    v14 := GenS2{F0: "foo bar", F1: ("" + "")}
    for fi15 := 0; fi15 < 18; fi15++ {
        fmt.Println(fi15, fi15, v4)
        _ = fi15
    }
    v5 = v5
    var v16 int64 = 1
    v17 := GenS2{F0: "a", F1: "\n\t"}
    var v18 uint = 211
    v18 -= 0
    var v19 int64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
