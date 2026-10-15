package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }

func main() {
    var v1 int8 = -24
    v2 := []string{("\n\t" + ("\n\t" + ""))}
    var v3 uint = 42
    v4 := make(map[string]string)
    v4[("foo bar" + "")] = "a"
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            if false {
                var v6 int8 = v1
                _ = v6
            }
            v1 -= (-(v1))
            var v7 int16 = 1000
            _ = fi5
            _ = v7
        }
        fmt.Println(v3)
        var v8 uint16 = 221
        _ = v8
    }
    var v9 string = "foo bar"
    if false {
        fmt.Println(v1)
        v10 := GenS1{F0: 1, F1: 235}
        var v11 string = ("\n\t" + (v9 + "hello"))
        _ = v10
        _ = v11
    }
    _ = len(v2)
    fmt.Println(v9)
    v12 := GenS1{F0: 42, F1: -32768}
    fmt.Println(v9, v9)
    v1 += (v1 << ((v3) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
