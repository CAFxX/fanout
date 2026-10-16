package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"a", (("a" + "a") + ("" + ("a" + "foo bar"))), ("\n\t" + "hello"), ("a" + "foo bar")}
    var v3 bool = (801 <= 1000)
    fmt.Println(v3)
    if v3 {
        v3 = v3
    }
    v3 = false
    v3 = ((!(!(385 <= 42))) || v3)
    var v4 float64 = ((-2.25 * (float64(-2.25) + -2.25)) / (1e10 + (3.14159 - (3.14159 + 1.5))))
    fmt.Println(v4)
    v5 := GenS1{F0: 1, F1: 100, F2: 4294967295}
    v3 = v3
    if false {
        var v6 bool = v3
        var v7 int = 0
        _ = len(v2)
        _ = v6
        _ = v7
    }
    var v8 bool = v3
    v9 := GenS1{F0: 215, F1: 125, F2: 793}
    if (2147483648 > 0) {
        switch "a" {
        case "\n\t":
            switch "hello" {
            case "":
                var v10 uint = 376
                _ = v10
            case "a":
                fmt.Println(v8, v4, v3)
            case "hello":
                v11 := GenS1{F0: -44, F1: -1, F2: -9223372036854775808}
                _ = v11
            }
        case "a":
            var v12 uint64 = 4294967295
            _ = v12
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            fmt.Println(v8)
            fmt.Println(v3, v3, v8)
            fmt.Println(fi13, v4, v3)
            _ = fi13
        }
        var v14 int64 = 451
        _ = v14
    } else {
        var v15 uint = 2
        if len(v2) > 0 {
            v2[0] = (("" + "hello") + ("hello" + ""))
        }
        _ = v15
    }
    var v16 bool = v8
    switch "hello" {
    case "\n\t":
        v8 = ("\n\t" < "hello")
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
