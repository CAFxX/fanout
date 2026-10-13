package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }

func main() {
    var v1 int = -199
    v2 := []string{"hello"}
    v3 := make(map[int]bool)
    var v4 uint = 2
    switch "foo bar" {
    case "a":
        var v5 uint16 = 1
        _ = v5
    case "foo bar":
        var v6 uint = v4
        _ = v6
    }
    v1 += (v1 >> ((v4) & 63))
    switch (v1 / ((32768) | 1)) {
    case 2147483647:
        var v7 int8 = 1
        _ = v7
    }
    v4 = (v4 % ((v4) | 1))
    v8 := GenS1{F0: 42, F1: -771, F2: ("hello" + "hello")}
    if false {
        v9 := GenS1{F0: 0, F1: 1, F2: "\n\t"}
        switch "foo bar" {
        case "foo bar":
            v10 := GenS1{F0: 1, F1: 32767, F2: ("foo bar" + ("foo bar" + "\n\t"))}
            _ = v10
        default:
            var v11 int64 = 1
            _ = v11
        }
        _ = v9
    }
    switch v1 {
    case 4294967295:
        for fi12 := 0; fi12 < 13; fi12++ {
            v1 = (v1 ^ (v1 / ((-787) | 1)))
            _ = fi12
        }
    default:
        var v13 int64 = -455
        _ = v13
    }
    if false {
        fmt.Println(v4, v4)
    } else {
        v4 = (v4 - (v4 | (v4 | 4294967295)))
        for fi14 := 0; fi14 < 6; fi14++ {
            v2 = append(v2, "")
            v2 = append(v2, "a")
            _ = fi14
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
