package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 uint, p2 uint8) int8 {
    for fi5 := 0; fi5 < 20; fi5++ {
        p0 = -9223372036854775808
        _ = fi5
    }
    switch "" {
    case "foo bar":
        var v6 int = -976
        _ = v6
    case "a":
        switch "a" {
        case "hello":
            fmt.Println(p1, p2, p2)
        default:
            var v7 uint64 = 1
            _ = v7
        }
    case "\n\t":
        p1 = (^(p1))
    default:
        p2 = (^(p2))
    }
    v8 := GenS1{F0: -9223372036854775808, F1: 42}
    var v9 int32 = 891
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 61
}
func genFn1(p0 uint16) int {
    var v10 bool = false
    switch "a" {
    case "\n\t":
        if v10 {
            v10 = v10
        } else {
            v10 = false
            switch -296 {
            case 1:
                var v11 float32 = float32(1e10)
                _ = v11
            }
        }
    default:
        v12 := []int{32767, 32768, 181}
        _ = v12
    }
    _ = p0
    _ = v10
    return 0
}

func main() {
    var v1 uint64 = 2
    var v2 bool = (-77 < -128)
    var v3 uint8 = 57
    v4 := make(map[string]bool)
    fmt.Println(v2, v2)
    v13 := GenS1{F0: 32767, F1: (v1 & (^(v1)))}
    switch "\n\t" {
    case "":
        if true {
            var v14 string = ("a" + ("a" + "\n\t"))
            v1 = 18446744073709551615
            fmt.Println(v14)
            _ = v14
        }
    }
    v3 = (v3 + 2)
    var v15 uint = 4294967295
    v16 := v4["hello"]
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 uint16 = 42
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v16 = true
        fmt.Println(v3, v1, v16)
        v20 := []int{(fi19 - -1), fi19, fi19}
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
