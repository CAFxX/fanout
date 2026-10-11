package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
func genFn0(p0 uint32, p1 int, p2 string) uint32 {
    if (false && (false || (2 >= 42))) {
        v4 := []int{p1, p1}
        fmt.Println(p1, p2)
        _ = v4
    }
    switch p2 {
    case "a":
        v5 := []int{p1, (^(p1))}
        _ = v5
    case "foo bar":
        var v6 uint = 0
        _ = v6
    }
    v7 := GenS1{F0: -1, F1: 0, F2: 163}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    var v1 int16 = 1
    v2 := []uint32{42}
    var v3 int8 = 127
    if (!false) {
        _ = len(v2)
        var v8 uint64 = 1
        var v9 string = "hello"
        _ = v8
        _ = v9
    } else {
        var v10 bool = (false || false)
        for fi11 := 0; fi11 < 4; fi11++ {
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = fi11
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v10
    }
    if len(v2) > 0 {
        v2[0] = 889
    }
    var v12 int16 = int16(v3)
    for fi13 := 0; fi13 < 15; fi13++ {
        _ = len(v2)
        _ = fi13
    }
    var v14 uint8 = 127
    if (1 != uint64(v14)) {
        v3 = v3
        v14 = v14
        if true {
            v3 += v3
            v3 -= 30
            var v15 uint8 = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        v12 -= v1
        for fi17 := 0; fi17 < 19; fi17++ {
            v12 -= v1
            _ = fi17
        }
        _ = fi16
    }
    var v18 uint64 = 18446744073709551615
    v18 += 206
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
