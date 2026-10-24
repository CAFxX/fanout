package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int }

func main() {
    var v1 int32 = 0
    var v2 int32 = (-(v1))
    v3 := []int{0, -2147483648, 32768}
    var v4 float32 = (-2.25 / ((1e10 + 3.14159) - (0.0 / 3.14159)))
    v5 := make(map[string]bool)
    v6 := make(map[string]uint32)
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v4)
        _ = fi7
    }
    v8 := GenS1{F0: 0, F1: 41, F2: 127}
    var v9 uint64 = 624
    v10 := v6["hello"]
    if (!true) {
        var v11 string = "a"
        v4 = v4
        _ = v11
    }
    fmt.Println(v2, v9, v2)
    if len(v3) > 0 {
        v3[0] = 32768
    }
    switch "\n\t" {
    case "a":
        v12 := GenS1{F0: 65535, F1: 0, F2: 1}
        _ = v12
    case "foo bar":
        var v13 string = (("a" + "a") + ("" + "foo bar"))
        _ = v13
    }
    fmt.Println(v2, v10)
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 int64 = -1
        fmt.Println(fi14, fi14, v9)
        if false {
            delete(v5, "foo bar")
            for fi16 := 0; fi16 < 9; fi16++ {
                v17 := GenS2{F0: 0, F1: 0, F2: (-(fi16))}
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 3; fi18++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v2, v1)
        }
        var v19 int8 = 1
        v20 := v5["foo bar"]
        _ = fi18
        _ = v19
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            fmt.Println(v2, v10, fi22)
            _ = fi22
        }
        var v23 int64 = -1
        _ = fi21
        _ = v23
    }
    v24 := GenS1{F0: 65535, F1: 2, F2: 127}
    var v25 uint16 = 32767
    fmt.Println(v9, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
