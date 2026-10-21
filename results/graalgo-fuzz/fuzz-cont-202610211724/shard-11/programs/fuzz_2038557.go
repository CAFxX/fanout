package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int; F1 uint; F2 uint }
func genFn0(p0 int, p1 uint, p2 bool) int32 {
    p1 = 42
    v6 := []int{p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 769
}
func genFn1(p0 string, p1 int8) uint8 {
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            p0 = (p0 + ("" + (p0 + "foo bar")))
            _ = fi8
        }
        fmt.Println(p1, p0, fi7)
        v9 := []int{}
        _ = fi7
        _ = v9
    }
    var v10 string = p0
    fmt.Println(p1, p0)
    v11 := []int{-2147483648}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 48
}

func main() {
    v1 := []int{570}
    v2 := make(map[string]bool)
    var v3 bool = (!false)
    var v4 int8 = 1
    var v5 uint64 = 42
    fmt.Println(v5)
    v5 = (^(v5))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v4 = v4
    v12 := GenS1{F0: -35, F1: 0, F2: (!v3)}
    v13 := GenS2{F0: -1, F1: 435, F2: 933}
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint16 = 65535
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        if v3 {
            v3 = v3
            for fi17 := 0; fi17 < 12; fi17++ {
                v4 = (^(v4))
                _ = fi17
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = len(v2)
        }
        _ = fi16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
