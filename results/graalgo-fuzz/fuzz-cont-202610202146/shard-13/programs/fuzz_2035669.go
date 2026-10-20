package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 255
    v3 := make(map[string]string)
    var v4 int64 = 1
    v5 := []string{}
    if (32767 != 309) {
        v6 := GenS1{F0: 32768}
        v2 = (v2 - v2)
        _ = v6
    }
    _ = len(v5)
    v2 -= (v2 << ((v2) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v7 := v3[("foo bar" + "a")]
    v3[v7] = (v7 + v7)
    v1[((((("" + v7) + "hello") + (v7 + ("foo bar" + (v7 + ("hello" + v7))))) + v7) + ((v7 + (v7 + "hello")) + "a"))] = v7
    if true {
        v4 = v4
        v8 := GenS1{F0: 971}
        _ = v8
    } else {
        v9 := GenS1{F0: 9223372036854775807}
        _ = v9
    }
    v7 = (("foo bar" + v7) + "hello")
    fmt.Println(v2)
    fmt.Println(v2)
    var v10 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
