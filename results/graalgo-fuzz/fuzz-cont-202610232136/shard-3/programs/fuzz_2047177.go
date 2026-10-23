package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int8 }

func main() {
    v1 := []string{"foo bar", "hello", "hello", "a"}
    var v2 int16 = 1
    var v3 uint8 = 216
    var v4 string = "a"
    for fi5 := 0; fi5 < 18; fi5++ {
        fmt.Println(v3)
        _ = fi5
    }
    if len(v1) > 0 {
        v1[0] = ((("foo bar" + (v4 + v4)) + v4) + (v4 + (v4 + v4)))
    }
    if len(v1) > 0 {
        v1[0] = (v4 + "foo bar")
    }
    var v6 string = v4
    if false {
        var v7 bool = false
        fmt.Println(v6, v2)
        _ = v7
    } else {
        var v8 float64 = float64(0.0)
        _ = v8
    }
    _ = len(v1)
    var v9 int16 = v2
    fmt.Println(v2)
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := GenS1{F0: 42, F1: (true && ((127 == fi10) || true)), F2: 42}
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 440, F1: false, F2: -58}
    var v13 uint = 799
    for fi14 := 0; fi14 < 9; fi14++ {
        if true {
            v1 = append(v1, v6)
        } else {
            var v15 int32 = -2147483648
            var v16 int16 = v2
            _ = v15
            _ = v16
        }
        var v17 bool = true
        _ = fi14
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v9
}
