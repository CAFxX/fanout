package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 string }
func genFn0(p0 string) int32 {
    p0 = ((p0 + p0) + (p0 + p0))
    v5 := GenS1{F0: 975, F1: 255}
    p0 = p0
    var v6 int16 = 32767
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(p0, p0, p0)
        _ = fi7
    }
    _ = p0
    _ = v5
    _ = v6
    return 151
}

func main() {
    v1 := []uint32{976, 4294967295, 0, 2}
    var v2 int64 = 9223372036854775807
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    if true {
        var v8 uint64 = 42
        fmt.Println(v2)
        if (false && (!true)) {
            fmt.Println(v2, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            v8 = v8
        }
        _ = v8
    }
    fmt.Println(v2)
    if (!true) {
        switch "" {
        case "foo bar":
            var v9 uint = 2
            _ = v9
        case "\n\t":
            fmt.Println(v2)
        case "":
            var v10 int8 = 0
            _ = v10
        }
    } else {
        var v11 bool = false
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    if true {
        v12 := GenS2{F0: "\n\t"}
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 uint32 = 4294967295
            _ = fi13
            _ = v14
        }
        var v15 int32 = genFn0("a")
        _ = v12
        _ = v15
    }
    v2 = v2
    v4[4294967295] = 65535
    v16 := GenS2{F0: "hello"}
    switch -120 {
    case 0:
        var v17 int = -1
        _ = v17
    case 833:
        var v18 string = "hello"
        _ = v18
    }
    v2 += (v2 | 655)
    var v19 int = 127
    var v20 uint8 = 0
    v19 -= (-(v19))
    var v21 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
