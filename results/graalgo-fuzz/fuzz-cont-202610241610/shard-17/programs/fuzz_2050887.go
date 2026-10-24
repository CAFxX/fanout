package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 string; F2 int16 }
type GenS4 struct { F0 string; F1 uint; F2 int16 }
type GenS5 struct { F0 bool; F1 int; F2 uint32 }
func genFn0(p0 int64, p1 int64, p2 int) string {
    v4 := []int{}
    if true {
        var v5 int16 = -256
        v6 := GenS1{F0: 0, F1: 1}
        var v7 uint8 = 128
        _ = v5
        _ = v6
        _ = v7
    }
    p0 = (-(p0))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ""
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 2, 2147483647}
    v2 := []uint32{136, 376, 162}
    var v3 string = ""
    v3 = ""
    v8 := GenS2{F0: 76}
    v9 := GenS3{F0: -992, F1: "", F2: -1}
    fmt.Println(v3)
    v3 = ""
    v3 = v3
    for fi10 := 0; fi10 < 20; fi10++ {
        v3 = v3
        switch v3 {
        case "":
            v11 := GenS4{F0: ("" + ("a" + v3)), F1: 0, F2: 0}
            _ = v11
        case "a":
            v12 := GenS3{F0: fi10, F1: v3, F2: 42}
            _ = v12
        case "hello":
            for fi13 := 0; fi13 < 5; fi13++ {
                v3 = v3
                _ = fi13
            }
        default:
            var v14 string = ("" + v3)
            _ = v14
        }
        _ = fi10
    }
    var v15 bool = (false && (!true))
    for fi16 := 0; fi16 < 7; fi16++ {
        v3 = v3
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v17 uint8 = 127
        _ = fi16
        _ = v17
    }
    v18 := GenS5{F0: v15, F1: 1, F2: 2147483648}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint32 = 4294967295
        for fi21 := 0; fi21 < 20; fi21++ {
            var v22 uint16 = 2
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        v15 = v15
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
