package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = -378
    var v3 uint16 = 2
    v2 = 0
    v2 = 42
    for fi4 := 0; fi4 < 5; fi4++ {
        v5 := []int{}
        if ("\n\t" == (((("a" + "") + "a") + (("hello" + "foo bar") + ("hello" + ("\n\t" + "foo bar")))) + "")) {
            var v6 string = ("foo bar" + "a")
            _ = v6
        } else {
            var v7 string = "foo bar"
            v2 = (v2 * (v2 & v2))
            _ = v7
        }
        if true {
            if false {
                var v8 uint = 0
                _ = v8
            }
        } else {
            var v9 uint32 = 240
            v3 = v3
            _ = v9
        }
        _ = fi4
        _ = v5
    }
    v10 := GenS1{F0: 578, F1: 42}
    var v11 float32 = 1e10
    v2 = v2
    var v12 float64 = float64(3.14159)
    var v13 bool = (95 < -103)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
