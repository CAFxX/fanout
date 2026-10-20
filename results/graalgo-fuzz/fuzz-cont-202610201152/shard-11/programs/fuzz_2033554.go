package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    v1 := []string{}
    var v2 int8 = 1
    var v3 int8 = v2
    var v4 uint32 = 42
    var v5 string = ""
    var v6 uint = 1
    v1 = append(v1, ("\n\t" + ("\n\t" + "\n\t")))
    fmt.Println(v6, v2, v4)
    v4 += uint32(v6)
    var v7 float32 = (1e10 - (3.14159 + 1e10))
    fmt.Println(v2)
    var v8 uint64 = 1
    var v9 uint16 = 960
    var v10 uint32 = v4
    switch 4294967295 {
    case 760:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch 4294967295 {
        case 865:
            var v12 string = (v5 + ((v5 + "foo bar") + "\n\t"))
            _ = v12
        }
        v4 = (v10 & (v4 << ((v6) & 63)))
        _ = fi11
    }
    _ = len(v1)
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v9, fi13, v8)
        _ = fi13
    }
    v14 := GenS1{F0: uint16(v6), F1: 765}
    switch v5 {
    case "hello":
        for fi15 := 0; fi15 < 10; fi15++ {
            v5 = v5
            v6 += v6
            var v16 uint16 = (v9 - 65535)
            _ = fi15
            _ = v16
        }
    }
    fmt.Println(v4, v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
