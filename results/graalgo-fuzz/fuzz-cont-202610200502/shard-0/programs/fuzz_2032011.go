package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = ("" < ("hello" + ""))
    var v3 string = (("" + ("\n\t" + ("a" + (("hello" + "\n\t") + "foo bar")))) + ("foo bar" + "foo bar"))
    var v4 int32 = 0
    v4 -= v4
    for fi5 := 0; fi5 < 7; fi5++ {
        delete(v1, 557)
        v4 += v4
        _ = fi5
    }
    switch 4294967295 {
    case 42:
        var v6 uint16 = 61
        _ = v6
    case 32767:
        var v7 int16 = 42
        _ = v7
    case -247:
        v3 = (("\n\t" + (v3 + v3)) + "foo bar")
    default:
        var v8 uint64 = 142
        _ = v8
    }
    switch v3 {
    case "foo bar":
        var v9 float32 = float32(1e10)
        _ = v9
    default:
        var v10 uint = 1
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v12 := GenS1{F0: -1}
        var v13 int16 = -1
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float32 = (0.0 - -2.25)
    var v15 uint32 = 54
    _ = len(v1)
    for fi16 := 0; fi16 < 18; fi16++ {
        v17 := GenS2{F0: 596, F1: fi16}
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v15 = 65535
        _ = fi18
    }
    v19 := []int{0, -883, 32767}
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
