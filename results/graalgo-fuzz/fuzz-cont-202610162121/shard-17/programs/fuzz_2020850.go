package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int8 = -1
    var v2 bool = (42 > 8)
    var v3 uint16 = 65535
    v4 := make(map[string]int)
    var v5 float64 = float64(((0.1 / 1.5) / 3.14159))
    switch "a" {
    case "foo bar":
        v6 := GenS1{F0: 722}
        _ = v6
    case "":
        var v7 int8 = v1
        _ = v7
    case "\n\t":
        var v8 int8 = 127
        _ = v8
    }
    fmt.Println(v3, v3)
    if (-991 < -143) {
        switch 9223372036854775807 {
        case -1:
            v9 := []int{32768}
            _ = v9
        }
    } else {
        v2 = v2
    }
    delete(v4, ("hello" + ("\n\t" + "")))
    var v10 uint8 = 127
    v11 := []int{}
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(v5, v3, v10)
        _ = fi12
    }
    v5 = 0.0
    var v13 float64 = 0.0
    v10 = (v10 >> ((216) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
