package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("foo bar" + ("a" + "foo bar")), "", ("a" + "\n\t")}
    var v3 int8 = -128
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := GenS1{F0: 0, F1: ((4294967295 <= 1) || ((((((v3 >> ((698) & 63)) < (-(v3))) || (124 == 213)) || false) && true) || false))}
        _ = fi4
        _ = v5
    }
    switch -422 {
    case 4294967295:
        fmt.Println(v3)
    }
    if false {
        v6 := GenS1{F0: 73, F1: ((1e10 * (0.1 + 1.5)) > 3.14159)}
        fmt.Println(v3)
        _ = v6
    }
    v7 := GenS1{F0: 0, F1: (((!(!true)) || (false && false)) || (42 >= 435))}
    fmt.Println(v3)
    for fi8 := 0; fi8 < 16; fi8++ {
        var v9 uint = 1
        if true {
            v3 = (v3 | (v3 | -127))
            fmt.Println(fi8, v3)
        }
        fmt.Println(v9)
        _ = fi8
        _ = v9
    }
    var v10 int = 2147483647
    v3 -= v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
}
