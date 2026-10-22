package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool) uint64 {
    var v5 int64 = 9223372036854775807
    p0 += (^(p0))
    v6 := GenS1{F0: (255 <= 2), F1: "foo bar"}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []int{-1, 887, 42}
    v2 := make(map[uint32]int)
    v3 := make(map[string]bool)
    var v4 float32 = (1e10 + 0.0)
    if (0 > 0) {
        var v7 int32 = 0
        _ = v7
    } else {
        v8 := GenS1{F0: true, F1: "\n\t"}
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 int8 = 42
                _ = fi10
                _ = v11
            }
            if ((-71 >= 9223372036854775807) && false) {
                v4 = v4
            }
            _ = fi9
        }
        _ = v8
    }
    _ = len(v3)
    var v12 int = 128
    if true {
        fmt.Println(v4, v12)
        if len(v1) > 0 {
            v1[0] = v12
        }
        v4 = (((1.5 + v4) + v4) / v4)
    } else {
        fmt.Println(v12, v12)
        for fi13 := 0; fi13 < 2; fi13++ {
            v4 -= v4
            var v14 uint8 = 34
            _ = fi13
            _ = v14
        }
    }
    if false {
        for fi15 := 0; fi15 < 17; fi15++ {
            fmt.Println(fi15, fi15)
            _ = fi15
        }
    }
    switch ("\n\t" + "") {
    case "":
        v4 += -2.25
    }
    v16 := GenS2{F0: 42}
    var v17 string = "a"
    _ = len(v3)
    v17 = v17
    v18 := GenS2{F0: 0}
    v17 = "a"
    v19 := GenS3{F0: 122}
    delete(v3, "foo bar")
    v17 = "a"
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 int32 = -361
        _ = fi20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
