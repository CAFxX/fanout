package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 0.1
    var v3 float32 = 1e10
    var v4 float64 = 0.0
    var v5 int64 = 0
    v6 := GenS1{F0: 0}
    var v7 uint32 = 517
    var v8 uint32 = (^(v7))
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 int32 = -22
        var v11 uint = 335
        fmt.Println(v5)
        _ = fi9
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := GenS1{F0: 906}
        var v14 int32 = 42
        var v15 uint64 = 2
        _ = fi12
        _ = v13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        _ = len(v1)
        var v17 uint = 34
        _ = fi16
        _ = v17
    }
    v8 = v7
    switch "hello" {
    case "hello":
        v18 := []int{1, -402}
        _ = v18
    default:
        switch "foo bar" {
        case "\n\t":
            for fi19 := 0; fi19 < 10; fi19++ {
                v3 += v3
                v7 += (^(v7))
                _ = fi19
            }
        default:
            v8 = (v7 + v7)
        }
    }
    if false {
        v20 := GenS1{F0: -1}
        _ = v20
    }
    var v21 string = "a"
    var v22 string = ("\n\t" + "a")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
