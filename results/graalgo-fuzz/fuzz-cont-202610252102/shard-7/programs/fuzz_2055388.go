package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }

func main() {
    v1 := []uint32{42, 2147483648}
    var v2 float64 = 1e10
    var v3 float64 = float64(1e10)
    var v4 bool = true
    v5 := []string{}
    var v6 int16 = -1
    if len(v5) > 0 {
        v5[0] = (("foo bar" + "\n\t") + ("hello" + "foo bar"))
    }
    if ("a" != "foo bar") {
        fmt.Println(v4, v2, v6)
    } else {
        v3 = v3
    }
    var v7 bool = (v3 == float64(1e10))
    var v8 bool = v7
    var v9 int = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if true {
        var v10 int = v9
        v11 := GenS1{F0: (v10 >> ((51) & 63)), F1: 42, F2: (0 > -864)}
        _ = v10
        _ = v11
    } else {
        var v12 uint8 = 166
        fmt.Println(v2, v8)
        switch "a" {
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        _ = v12
    }
    var v13 string = "a"
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v6 = -32768
        switch v9 {
        case -33:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        case -1:
            v8 = (!(65535 < 290))
        }
        _ = fi14
    }
    var v15 int64 = 1
    v13 = v13
    var v16 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
