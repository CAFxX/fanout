package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 bool }
func genFn0(p0 uint32) int64 {
    p0 = p0
    if false {
        v5 := []int{98}
        var v6 int = 42
        _ = v5
        _ = v6
    }
    _ = p0
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 42
    var v3 int32 = -1
    var v4 uint = 65535
    fmt.Println(v3)
    v4 += (v4 + v4)
    for fi7 := 0; fi7 < 12; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    v8 := []int{-734, 42}
    switch "hello" {
    case "a":
        v4 = v4
    case "hello":
        var v9 uint16 = 65535
        _ = v9
    default:
        var v10 string = "hello"
        _ = v10
    }
    switch "hello" {
    case "foo bar":
        v4 = v4
    default:
        for fi11 := 0; fi11 < 1; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 bool = (false || true)
                var v14 string = "hello"
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    v15 := GenS1{F0: 42, F1: "a", F2: ("foo bar" >= "\n\t")}
    var v16 uint32 = 4294967295
    v17 := GenS1{F0: 42, F1: "a", F2: true}
    var v18 int16 = -708
    v4 = v4
    if false {
        for fi19 := 0; fi19 < 11; fi19++ {
            fmt.Println(v3, v2)
            _ = fi19
        }
    } else {
        v4 = (v4 - (v4 << ((v4) & 63)))
        if true {
            fmt.Println(v4)
        }
    }
    var v20 uint = (v4 / (((v4 & (v4 / ((718) | 1)))) | 1))
    switch 809 {
    case -830:
        var v21 int8 = 127
        _ = v21
    case 127:
        for fi22 := 0; fi22 < 20; fi22++ {
            for fi23 := 0; fi23 < 11; fi23++ {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = fi23
            }
            _ = fi22
        }
    default:
        v20 -= 0
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
}
