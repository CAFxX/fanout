package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 string, p1 int32, p2 int) uint {
    if false {
        var v6 string = p0
        v7 := GenS1{F0: 4294967295, F1: "foo bar"}
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{(fi8 & fi8), -770}
            _ = fi8
            _ = v9
        }
        _ = v6
        _ = v7
    }
    p1 = -1
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() int32 {
    var v10 uint16 = 255
    v11 := GenS1{F0: 2, F1: "\n\t"}
    for fi12 := 0; fi12 < 20; fi12++ {
        var v13 uint8 = 255
        v13 += (^(v13))
        fmt.Println(v13, v13, v13)
        _ = fi12
        _ = v13
    }
    var v14 uint = 255
    fmt.Println(v10, v10)
    _ = v10
    _ = v11
    _ = v14
    return 1
}

func main() {
    var v1 int8 = 1
    var v2 float64 = 0.1
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 1
    v5 := make(map[uint32]bool)
    switch "foo bar" {
    case "hello":
        var v15 float64 = v2
        _ = v15
    }
    var v16 int16 = 770
    var v17 uint = 1
    for fi18 := 0; fi18 < 3; fi18++ {
        for fi19 := 0; fi19 < 1; fi19++ {
            v4 -= (v4 | v4)
            v4 += v4
            delete(v3, 130)
            if true {
                var v20 uint = v17
                v21 := GenS1{F0: 182, F1: ("\n\t" + "")}
                _ = v20
                _ = v21
            }
            _ = fi19
        }
        _ = fi18
    }
    fmt.Println(v4, v2)
    v2 -= v2
    var v22 uint32 = 2147483647
    fmt.Println(v22, v4, v4)
    v23 := GenS2{F0: v1, F1: (v1 & -27)}
    for fi24 := 0; fi24 < 1; fi24++ {
        v2 = v2
        _ = fi24
    }
    var v25 uint16 = 806
    var v26 float32 = (0.1 - (0.1 + 3.14159))
    var v27 bool = false
    switch 690 {
    case 128:
        switch (("a" + "foo bar") + "hello") {
        case "":
            fmt.Println(v22)
        default:
            for fi28 := 0; fi28 < 1; fi28++ {
                var v29 uint8 = v4
                _ = fi28
                _ = v29
            }
        }
    default:
        v4 += v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
