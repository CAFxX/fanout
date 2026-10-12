package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]int)
    var v2 int64 = -1
    var v3 int32 = 1
    v4 := []string{"\n\t", ("" + "foo bar")}
    v3 = v3
    v2 += v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi5 := 0; fi5 < 20; fi5++ {
        v3 = 0
        if true {
            _ = len(v1)
            fmt.Println(fi5, v3)
            v2 -= v2
        } else {
            var v6 float64 = -2.25
            var v7 uint64 = 102
            _ = v6
            _ = v7
        }
        var v8 bool = true
        _ = fi5
        _ = v8
    }
    delete(v1, 779)
    v9 := GenS1{F0: -128}
    v2 = v2
    var v10 string = ("hello" + "hello")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
