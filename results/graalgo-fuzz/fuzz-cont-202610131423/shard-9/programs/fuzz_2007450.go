package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []int{}
    var v2 uint16 = 65535
    v3 := make(map[int]string)
    v4 := []uint32{}
    var v5 int16 = -1
    v6 := GenS1{F0: 0, F1: (v5 | v5), F2: 724}
    v2 = v2
    v7 := GenS2{F0: 2}
    for fi8 := 0; fi8 < 15; fi8++ {
        if ("" >= "foo bar") {
            v9 := GenS3{F0: 583}
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            var v11 string = ("hello" + "a")
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 15; fi12++ {
            v5 -= -32768
            _ = fi12
        }
        _ = fi8
    }
    var v13 float64 = float64(-2.25)
    var v14 string = ("hello" + "hello")
    var v15 uint8 = 126
    v15 = (^(v15))
    if false {
        v16 := GenS1{F0: (^(v2)), F1: (v5 | v5), F2: 1}
        var v17 uint8 = (v15 & (v15 | v15))
        _ = v16
        _ = v17
    } else {
        for fi18 := 0; fi18 < 18; fi18++ {
            fmt.Println(v15, v14, v5)
            var v19 uint32 = 823
            _ = fi18
            _ = v19
        }
        fmt.Println(v15)
        v5 = (v5 >> ((50) & 63))
    }
    var v20 uint = 5
    var v21 uint32 = 243
    fmt.Println(v5, v15)
    var v22 uint16 = v2
    var v23 string = "\n\t"
    if false {
        v5 -= v5
    }
    fmt.Println(v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
