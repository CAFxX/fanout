package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 string, p1 int32, p2 int) int32 {
    p0 = p0
    switch (^(p2)) {
    case 1:
        var v6 int8 = 127
        _ = v6
    case 32767:
        var v7 uint8 = 128
        _ = v7
    default:
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint32 = 42
            p0 = p0
            _ = fi8
            _ = v9
        }
    }
    p0 = p0
    if (!((p1 - -2147483648) == (p1 * p1))) {
        for fi10 := 0; fi10 < 3; fi10++ {
            fmt.Println(p2, fi10, p2)
            fmt.Println(fi10)
            v11 := GenS1{F0: p0, F1: p0}
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"", "a"}
    var v3 bool = false
    var v4 int16 = -1
    var v5 float64 = 1e10
    v3 = (!true)
    fmt.Println(v3, v3)
    fmt.Println(v4, v4, v5)
    v4 = v4
    v12 := GenS1{F0: "hello", F1: ("" + "hello")}
    if (0 == 127) {
        if v3 {
            delete(v1, "foo bar")
            switch 101 {
            case -736:
                var v13 uint64 = 19
                _ = v13
            }
        }
        v2 = append(v2, ("\n\t" + ""))
    } else {
        switch 42 {
        case 2147483647:
            var v14 uint8 = 128
            _ = v14
        case 888:
            v3 = v3
        case -868:
            v15 := GenS1{F0: "a", F1: ("hello" + ("foo bar" + "hello"))}
            _ = v15
        }
    }
    switch "foo bar" {
    case "":
        v16 := GenS1{F0: ("\n\t" + "foo bar"), F1: "a"}
        _ = v16
    default:
        v4 -= v4
    }
    if v3 {
        v17 := GenS1{F0: "", F1: ("\n\t" + "\n\t")}
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v4 = 0
        var v19 string = ("\n\t" + "foo bar")
        fmt.Println(fi18, v19, v4)
        _ = fi18
        _ = v19
    }
    v20 := GenS1{F0: "foo bar", F1: "foo bar"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
