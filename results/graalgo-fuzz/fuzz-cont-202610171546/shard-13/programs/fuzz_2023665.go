package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint16; F1 string }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]bool)
    v3 := []uint32{42}
    var v4 string = ("" + (("" + (("hello" + "\n\t") + ("foo bar" + "hello"))) + "a"))
    var v5 uint64 = 18446744073709551615
    switch v4 {
    case "":
        v6 := GenS1{F0: 835, F1: 1, F2: 2147483648}
        _ = v6
    case "foo bar":
        v4 = "\n\t"
    }
    var v7 int16 = 32767
    var v8 uint32 = 1
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 uint16 = 65535
            v12 := GenS1{F0: fi10, F1: v7, F2: v8}
            fmt.Println(v8)
            if false {
                v13 := GenS1{F0: fi9, F1: (v7 / ((v7) | 1)), F2: 1}
                v7 += -32768
                var v14 int8 = 28
                _ = v13
                _ = v14
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    if false {
        v15 := GenS1{F0: -1, F1: v7, F2: (v8 << ((867) & 63))}
        var v16 int32 = -1
        v4 = "a"
        _ = v15
        _ = v16
    } else {
        delete(v2, "foo bar")
        var v17 bool = true
        _ = v17
    }
    var v18 uint32 = v8
    for fi19 := 0; fi19 < 12; fi19++ {
        if false {
            fmt.Println(v18, v8)
        }
        for fi20 := 0; fi20 < 16; fi20++ {
            var v21 string = v4
            switch 32768 {
            case 353:
                v22 := GenS2{F0: 134, F1: "\n\t"}
                _ = v22
            case 904:
                v23 := GenS1{F0: fi19, F1: 0, F2: v8}
                var v24 bool = true
                _ = v23
                _ = v24
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(v8, v7, v5)
        _ = fi19
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 13; fi26++ {
            var v27 uint = 2
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    var v28 uint = 18446744073709551615
    for fi29 := 0; fi29 < 19; fi29++ {
        v18 = 2
        _ = fi29
    }
    if false {
        fmt.Println(v28, v8)
    } else {
        _ = len(v3)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
