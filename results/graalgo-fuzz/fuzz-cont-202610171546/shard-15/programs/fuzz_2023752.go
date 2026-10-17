package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = 1
    var v2 int8 = 42
    v3 := make(map[int]uint32)
    var v4 int = 127
    var v5 bool = true
    switch v4 {
    case 4294967295:
        switch ("\n\t" + "hello") {
        case "":
            var v6 float32 = (0.1 * 3.14159)
            _ = v6
        case "hello":
            var v7 uint = 0
            _ = v7
        }
    case 127:
        v8 := GenS1{F0: ((("\n\t" + ("hello" + "hello")) + "foo bar") + "")}
        _ = v8
    }
    switch 714 {
    case -858:
        v1 = v1
    case -271:
        v1 = (v1 | v1)
    }
    v4 += 9223372036854775807
    var v9 string = ""
    if (v5 || ((0.1 * 0.0) > (0.0 - 0.1))) {
        v10 := GenS1{F0: ("foo bar" + ("" + ((v9 + "hello") + "hello")))}
        _ = v10
    }
    v11 := v3[v4]
    var v12 int = v4
    switch ("\n\t" + "a") {
    case "a":
        v5 = ((!false) || v5)
    case "foo bar":
        var v13 uint64 = 42
        _ = v13
    default:
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := []int{(fi14 >> ((944) & 63)), (v4 >> ((42) & 63))}
            _ = len(v15)
            _ = fi14
            _ = v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
