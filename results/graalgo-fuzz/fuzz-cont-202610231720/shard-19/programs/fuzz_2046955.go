package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }

func main() {
    var v1 bool = false
    var v2 uint8 = 165
    v3 := []string{}
    v4 := []int{620, -642, 1, 4294967295}
    for fi5 := 0; fi5 < 9; fi5++ {
        fmt.Println(v1, v2)
        _ = fi5
    }
    if v1 {
        fmt.Println(v2)
        var v6 uint8 = v2
        var v7 string = "foo bar"
        _ = v6
        _ = v7
    } else {
        if v1 {
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 float32 = ((3.14159 / 1.0) - 1e10)
                _ = fi8
                _ = v9
            }
        }
        v1 = (1 != 42)
        v1 = (-1 <= 32767)
    }
    if v1 {
        fmt.Println(v1, v1)
    }
    v4 = append(v4, -501)
    if false {
        var v10 string = ("a" + ("a" + "foo bar"))
        v11 := GenS1{F0: 1, F1: -9223372036854775808}
        _ = v10
        _ = v11
    }
    if v1 {
        var v12 float32 = (1.5 / -2.25)
        _ = v12
    }
    fmt.Println(v2, v2)
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
