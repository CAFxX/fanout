package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint }
type GenS2 struct { F0 int; F1 int8 }
func genFn0(p0 uint) uint64 {
    v5 := GenS1{F0: false, F1: true, F2: p0}
    p0 += (^(p0))
    p0 -= (^(p0))
    p0 += 2
    v6 := GenS2{F0: 342, F1: 127}
    _ = p0
    _ = v5
    _ = v6
    return 715
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    var v3 bool = (!(736 >= 42))
    v4 := []string{(("foo bar" + "") + ("" + "foo bar"))}
    v3 = (!false)
    if (((42 < 1) || (v3 || v3)) || (!false)) {
        if v3 {
            v1 = -636
        }
    }
    var v7 bool = v3
    if len(v4) > 0 {
        v4[0] = ("hello" + "a")
    }
    if true {
        var v8 uint16 = 2
        _ = v8
    } else {
        switch 325 {
        case 98:
            for fi9 := 0; fi9 < 15; fi9++ {
                for fi10 := 0; fi10 < 15; fi10++ {
                    _ = len(v2)
                    _ = fi10
                }
                v7 = (v7 || false)
                _ = len(v4)
                _ = fi9
            }
        case 315:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        default:
            var v11 int64 = 42
            _ = v11
        }
    }
    fmt.Println(v1, v1)
    var v12 int32 = 42
    v12 = (v1 | (v12 - v1))
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v14 int16 = -715
        _ = fi13
        _ = v14
    }
    if v7 {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint32 = 65535
            fmt.Println(v16, v12)
            for fi17 := 0; fi17 < 20; fi17++ {
                _ = len(v4)
                var v18 uint = 18446744073709551615
                _ = fi17
                _ = v18
            }
            _ = fi15
            _ = v16
        }
        fmt.Println(v12, v1)
    }
    v12 = 2147483647
    fmt.Println(v1, v12)
    v3 = (false && v3)
    fmt.Println(v7, v7)
    var v19 uint8 = 1
    fmt.Println(v12)
    var v20 int32 = v12
    v19 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v7
}
