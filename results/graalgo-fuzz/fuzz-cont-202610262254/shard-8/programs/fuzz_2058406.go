package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    var v3 bool = ((381 <= 18446744073709551615) || (0.0 <= 0.1))
    var v4 int64 = 113
    var v5 int = 1
    v4 = v4
    var v6 uint64 = 0
    if false {
        v7 := GenS1{F0: 864, F1: 655}
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(v5, v6, v3)
        var v9 uint8 = 0
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int8 = 127
    var v12 float64 = 1e10
    switch ("\n\t" + "\n\t") {
    case "":
        fmt.Println(v11, v5)
    default:
        v12 += v12
    }
    v13 := GenS2{F0: 32767, F1: 42, F2: 649}
    _ = len(v1)
    v14 := []int{v5}
    for fi15 := 0; fi15 < 19; fi15++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi15
    }
    delete(v1, "foo bar")
    for fi16 := 0; fi16 < 13; fi16++ {
        var v17 float64 = v12
        var v18 uint64 = 18446744073709551615
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v4)
    if ("hello" >= "foo bar") {
        var v19 int = 0
        _ = v19
    }
    v3 = v3
    for fi20 := 0; fi20 < 16; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            switch (-(v5)) {
            case 1:
                v12 -= float64(0.1)
            case 32768:
                delete(v1, "a")
            default:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            }
            _ = fi21
        }
        v5 += (-(v5))
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
