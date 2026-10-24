package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{0, 1, 65535}
    var v2 int32 = -776
    var v3 uint = 2
    var v4 string = "hello"
    v4 = (("" + "foo bar") + v4)
    switch v4 {
    case "\n\t":
        fmt.Println(v2)
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 773}
        _ = fi5
        _ = v6
    }
    var v7 uint16 = 65535
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 -= 252
    v8 := GenS1{F0: 1}
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: 1}
        _ = len(v1)
        _ = len(v1)
        _ = fi9
        _ = v10
    }
    v7 += (^(v7))
    v3 = v3
    var v11 int32 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v7, v11, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
