package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 int32, p2 uint16) uint {
    if true {
        for fi5 := 0; fi5 < 4; fi5++ {
            v6 := GenS1{F0: "\n\t", F1: (p2 | p2), F2: 939}
            _ = fi5
            _ = v6
        }
        v7 := []int{-634, 1, -985}
        var v8 uint32 = 531
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: ("" + "foo bar")}
    var v10 string = "a"
    var v11 uint32 = 525
    var v12 uint32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    _ = v9
    return 415
}

func main() {
    v1 := []int{}
    v2 := make(map[string]uint32)
    var v3 int16 = -32768
    var v4 int8 = 127
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            fmt.Println(v4, v4, v3)
        } else {
            v1 = append(v1, int(v4))
        }
        _ = fi13
    }
    fmt.Println(v3)
    if false {
        v3 = v3
        _ = len(v1)
    }
    var v14 uint64 = 18446744073709551615
    v15 := GenS1{F0: "a", F1: 155, F2: (^(v3))}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v14, v4, v14)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 uint64 = v14
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: ("hello" + ("hello" + "a"))}
    switch 289 {
    case 4294967295:
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v4)
            _ = fi19
        }
    case 0:
        _ = len(v1)
    default:
        v14 += (v14 | v14)
    }
    var v20 int32 = 267
    var v21 string = ("hello" + "hello")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
