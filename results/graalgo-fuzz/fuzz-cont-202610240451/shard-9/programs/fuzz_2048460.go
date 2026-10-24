package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint8 }
func genFn0(p0 int8, p1 int, p2 int) uint16 {
    var v6 int64 = 0
    fmt.Println(p2)
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 int32 = -20
        fmt.Println(p1, p0, fi7)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 255
}

func main() {
    v1 := []int{127, 42, 2147483647, 42}
    v2 := []uint32{2, 65535, 494}
    var v3 float64 = 0.1
    v4 := []uint32{1, 1}
    v5 := make(map[int]int)
    var v9 int = -1
    fmt.Println(v9, v9)
    if ((!(true || false)) || true) {
        var v10 uint32 = 2147483648
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint = 1
        v13 := GenS1{F0: 42, F1: -1, F2: 0}
        var v14 uint64 = 26
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    fmt.Println(v9)
    var v15 float64 = 0.0
    switch ((("\n\t" + ("" + (("hello" + "") + "hello"))) + "\n\t") + "\n\t") {
    case "foo bar":
        v16 := GenS1{F0: -1, F1: 42, F2: 127}
        _ = v16
    case "hello":
        for fi17 := 0; fi17 < 7; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                fmt.Println(fi18, v15)
                v5[9223372036854775807] = fi18
                v15 = v15
                _ = fi18
            }
            fmt.Println(fi17)
            fmt.Println(v15)
            _ = fi17
        }
    case "a":
        var v19 uint16 = 2
        _ = v19
    }
    v20 := GenS1{F0: 0, F1: -269, F2: 127}
    var v21 float32 = 1.5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
