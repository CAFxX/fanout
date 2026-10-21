package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 uint = 6
    var v2 int = 9223372036854775807
    v3 := []uint32{0}
    v4 := make(map[int]string)
    var v5 uint64 = 865
    var v6 int8 = 78
    switch "foo bar" {
    case "foo bar":
        switch (v2 - (-(v2))) {
        case 1:
            var v7 string = "a"
            _ = v7
        default:
            var v8 int = 127
            _ = v8
        }
    case "":
        var v9 string = ""
        _ = v9
    case "a":
        if (-538 >= -1) {
            v10 := GenS1{F0: "foo bar"}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        fmt.Println(fi11, v5, v5)
        _ = fi11
    }
    if len(v3) > 0 {
        v3[0] = 0
    }
    switch ("hello" + "foo bar") {
    case "hello":
        v6 = v6
    default:
        _ = len(v4)
    }
    if ((v2 * 0) >= v2) {
        fmt.Println(v6)
        v2 = v2
    }
    if false {
        var v12 int16 = 0
        v1 -= (v1 / ((v1) | 1))
        _ = len(v3)
        _ = v12
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v6, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
