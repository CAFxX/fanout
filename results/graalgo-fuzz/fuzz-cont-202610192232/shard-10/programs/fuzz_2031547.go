package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int64 }
func genFn0(p0 uint, p1 int) int32 {
    var v4 uint64 = 18446744073709551615
    v5 := GenS1{F0: 1, F1: (!(v4 >= (v4 << (((^(p0))) & 63)))), F2: 409}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return -2147483648
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{2147483648, 2, 2, 734}
    v3 := []uint32{2147483648, 348}
    for fi6 := 0; fi6 < 20; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int16 = 0
            _ = fi7
            _ = v8
        }
        switch "hello" {
        case "foo bar":
            var v9 string = ""
            _ = v9
        case "":
            var v10 bool = (!(true && true))
            _ = v10
        case "\n\t":
            fmt.Println(fi6)
        }
        v11 := GenS1{F0: -9223372036854775808, F1: true, F2: -470}
        _ = fi6
        _ = v11
    }
    var v12 int16 = 977
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 uint = 0
    switch (("a" + "foo bar") + ("a" + ("hello" + "foo bar"))) {
    case "a":
        var v14 uint64 = 18446744073709551615
        _ = v14
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "hello":
        v12 = v12
    default:
        v15 := GenS1{F0: -515, F1: false, F2: 42}
        _ = v15
    }
    v13 = (v13 | 255)
    fmt.Println(v12, v13)
    v12 += (v12 - -221)
    v2 = append(v2, 65535)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
