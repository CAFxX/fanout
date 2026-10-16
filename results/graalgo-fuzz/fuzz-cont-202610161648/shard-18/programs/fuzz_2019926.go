package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 int }

func main() {
    v1 := make(map[string]bool)
    var v2 bool = true
    var v3 float32 = -2.25
    for fi4 := 0; fi4 < 18; fi4++ {
        var v5 int16 = 0
        var v6 int = fi4
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 uint16 = 887
    v1[""] = (v3 > 1.5)
    v3 += (v3 - v3)
    switch ((("" + ("foo bar" + "")) + "") + (("hello" + ((("" + "a") + "hello") + ("\n\t" + "hello"))) + "\n\t")) {
    case "a":
        fmt.Println(v3, v3, v2)
    default:
        if (v2 || (true && (24 == 1))) {
            v8 := GenS1{F0: 111, F1: v2, F2: int(v7)}
            _ = v8
        }
    }
    v2 = (!true)
    var v9 int32 = 2147483647
    var v10 int8 = 42
    v9 += v9
    for fi11 := 0; fi11 < 10; fi11++ {
        for fi12 := 0; fi12 < 9; fi12++ {
            if true {
                var v13 string = ("a" + "")
                _ = v13
            }
            v10 -= (^(v10))
            if (true || (!((!v2) || v2))) {
                fmt.Println(v9, v10, fi11)
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println(v7, v7, v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
