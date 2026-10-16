package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64) int16 {
    v7 := GenS1{F0: (p0 - (p0 << ((0) & 63)))}
    p0 += (p0 << ((929) & 63))
    var v8 uint = 2
    _ = p0
    _ = v7
    _ = v8
    return -30
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]int)
    var v3 int32 = -1
    v4 := make(map[int]uint32)
    v5 := make(map[int]bool)
    var v6 float32 = (-2.25 - float32(3.14159))
    v3 -= -785
    var v9 float32 = 1e10
    v10 := []int{-787, 127}
    switch ("hello" + "hello") {
    case "foo bar":
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    case "":
        v3 += (v3 >> ((0) & 63))
    case "\n\t":
        v6 = v6
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        var v12 int8 = int8(fi11)
        _ = fi11
        _ = v12
    }
    var v13 int16 = -42
    v14 := GenS1{F0: 42}
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 int32 = v3
        _ = fi15
        _ = v16
    }
    var v17 uint = 0
    var v18 uint8 = 2
    _ = len(v1)
    for fi19 := 0; fi19 < 14; fi19++ {
        fmt.Println(v18, v17)
        _ = fi19
    }
    v18 -= v18
    var v20 int64 = -672
    var v21 float64 = 0.0
    v21 = (v21 * v21)
    if (v18 == 1) {
        _ = len(v10)
    } else {
        v17 = v17
    }
    var v22 float64 = v21
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
