package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8) string {
    v6 := GenS1{F0: 127, F1: 42}
    p0 = p0
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 int = (fi7 << ((720) & 63))
        _ = fi7
        _ = v8
    }
    fmt.Println(p0)
    if (!("foo bar" >= "a")) {
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    return ""
}

func main() {
    var v1 bool = true
    var v2 int16 = -1
    v3 := make(map[string]int)
    v4 := []string{(("\n\t" + ("\n\t" + "foo bar")) + "hello")}
    var v5 int = 32767
    var v9 uint = 2
    fmt.Println(v5, v9, v5)
    switch 0 {
    case 842:
        delete(v3, ("a" + "a"))
    default:
        v10 := GenS2{F0: ("a" + "hello")}
        _ = v10
    }
    if (108 != 42) {
        if (true || v1) {
            v1 = v1
            v11 := GenS1{F0: 32767, F1: 0}
            _ = v11
        }
    } else {
        v1 = (!(v1 || true))
        v12 := GenS1{F0: (v5 ^ 9223372036854775807), F1: 1}
        v13 := GenS2{F0: ""}
        _ = v12
        _ = v13
    }
    v5 = v5
    var v14 uint16 = 65535
    v2 = (v2 | 1000)
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
