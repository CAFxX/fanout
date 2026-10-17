package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint32; F1 int16 }

func main() {
    var v1 int16 = -1
    v2 := []int{-7}
    v3 := make(map[string]string)
    var v4 int = -180
    v4 += v4
    var v5 uint64 = 2
    if false {
        v6 := GenS1{F0: false, F1: "\n\t"}
        _ = len(v2)
        var v7 uint16 = 1
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = 0
            _ = fi8
            _ = v9
        }
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int64 = -650
    switch (-(v4)) {
    case 1:
        v3["hello"] = ("" + "")
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        fmt.Println(v5, fi12)
        var v13 bool = (((!false) && true) && (-2147483648 <= 2147483647))
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: ((v4 & (v4 + (v4 - (v4 * (-(v4)))))) == v4), F1: ""}
    switch "foo bar" {
    case "hello":
        var v15 string = ""
        _ = v15
    case "a":
        delete(v3, "\n\t")
    case "foo bar":
        fmt.Println(v1, v4, v4)
    }
    switch (v4 | -217) {
    case 46:
        switch v4 {
        case -703:
            switch (v4 * v4) {
            case -312:
                var v16 uint64 = v5
                _ = v16
            default:
                v17 := GenS1{F0: false, F1: ("foo bar" + ("\n\t" + ("hello" + "\n\t")))}
                _ = v17
            }
        }
    case 9223372036854775807:
        v18 := GenS2{F0: 4294967295, F1: (^(v1))}
        _ = v18
    }
    v4 = 128
    var v19 uint = 365
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
