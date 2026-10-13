package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 int16, p1 uint) uint {
    p0 = (p0 * (p0 ^ (-(p0))))
    v5 := GenS1{F0: 255}
    v6 := GenS1{F0: 334}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 uint = 1
    var v2 uint32 = 979
    v3 := []int{}
    var v4 uint8 = 195
    for fi7 := 0; fi7 < 18; fi7++ {
        v2 -= (v2 * v2)
        _ = fi7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := GenS1{F0: 42}
        _ = fi8
        _ = v9
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = (v1 - genFn0(32767, v1))
    fmt.Println(v2)
    switch ("foo bar" + ("" + "")) {
    case "\n\t":
        v1 = genFn0(0, 18446744073709551615)
    case "foo bar":
        switch "hello" {
        case "hello":
            fmt.Println(v1, v1, v1)
        case "\n\t":
            v10 := GenS1{F0: 255}
            _ = v10
        case "":
            fmt.Println(v1)
        }
    case "":
        if (!true) {
            _ = len(v3)
        }
    }
    if true {
        var v11 uint64 = 18446744073709551615
        if false {
            v4 += v4
        }
        fmt.Println(v2)
        _ = v11
    } else {
        for fi12 := 0; fi12 < 2; fi12++ {
            fmt.Println(v1, v2, fi12)
            var v13 int32 = -312
            var v14 uint16 = 42
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    fmt.Println(v1)
    var v15 uint16 = 65535
    var v16 int = 972
    if false {
        var v17 bool = ((v2 >> ((genFn0(32767, 920)) & 63)) < v2)
        _ = v17
    }
    var v18 uint32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
