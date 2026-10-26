package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 string }
func genFn0(p0 int8, p1 uint8) uint32 {
    switch "hello" {
    case "a":
        var v6 int16 = 1
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 uint64 = 0
        var v9 string = ""
        p1 -= p1
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return 42
}

func main() {
    var v1 float32 = 1e10
    var v2 string = ("a" + "hello")
    v3 := make(map[uint32]uint32)
    var v4 uint = 839
    var v5 float64 = (1.5 - 1.5)
    switch v2 {
    case "\n\t":
        if (true && (v1 >= 0.1)) {
            v11 := GenS1{F0: v4, F1: true}
            _ = v11
        }
    case "a":
        var v12 string = v2
        _ = v12
    case "foo bar":
        v13 := GenS1{F0: v4, F1: false}
        _ = v13
    default:
        v14 := []int{}
        _ = v14
    }
    v15 := v3[758]
    var v16 float64 = v5
    for fi17 := 0; fi17 < 1; fi17++ {
        _ = len(v3)
        switch v2 {
        case "":
            v15 = genFn0(127, 78)
        case "a":
            for fi18 := 0; fi18 < 5; fi18++ {
                fmt.Println(v2, v15, fi17)
                v19 := GenS2{F0: v2}
                _ = fi18
                _ = v19
            }
        case "foo bar":
            v5 += float64(-2.25)
        }
        _ = fi17
    }
    var v20 uint64 = 4294967295
    if ((!true) && (!true)) {
        delete(v3, v15)
    }
    fmt.Println(v20)
    var v21 int = 0
    v21 = (v21 << ((v4) & 63))
    var v22 string = v2
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 2; fi24++ {
            var v25 float64 = float64(v1)
            v4 = (^(v4))
            _ = fi24
            _ = v25
        }
        var v26 int8 = -62
        var v27 int = (-(fi23))
        _ = fi23
        _ = v26
        _ = v27
    }
    v28 := v3[uint32(v4)]
    var v29 float64 = v16
    for fi30 := 0; fi30 < 2; fi30++ {
        fmt.Println(v4, v15, v2)
        v4 = (v4 >> (((^(v4))) & 63))
        _ = fi30
    }
    v31 := []int{(v21 | v21), (v21 ^ v21), (v21 ^ (v21 + v21))}
    v32 := GenS2{F0: (v2 + "a")}
    var v33 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
}
