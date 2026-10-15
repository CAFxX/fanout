package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 string }

func main() {
    var v1 int64 = 0
    v2 := []uint32{4294967295, 982}
    var v3 int32 = 0
    var v4 uint8 = 1
    v5 := make(map[int]string)
    v1 = (v1 ^ v1)
    switch "foo bar" {
    case "a":
        v1 -= v1
    case "hello":
        var v6 uint32 = 0
        _ = v6
    default:
        v3 -= v3
    }
    if true {
        var v7 uint32 = 555
        _ = len(v5)
        _ = v7
    } else {
        var v8 int = 42
        var v9 uint64 = uint64(v4)
        _ = v8
        _ = v9
    }
    var v10 uint16 = 958
    fmt.Println(v4, v4)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            if false {
                var v13 uint64 = 617
                _ = v13
            }
            var v14 string = ("\n\t" + "foo bar")
            v10 = (v10 >> ((769) & 63))
            _ = fi12
            _ = v14
        }
        _ = fi11
    }
    if true {
        if len(v2) > 0 {
            v2[0] = 425
        }
        v1 = v1
        var v15 float32 = 3.14159
        _ = v15
    } else {
        v10 = 930
        _ = len(v2)
        v4 = (v4 >> ((18446744073709551615) & 63))
    }
    v16 := GenS1{F0: "hello", F1: 65535}
    for fi17 := 0; fi17 < 5; fi17++ {
        _ = len(v2)
        var v18 bool = ((0 >= 203) || true)
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
