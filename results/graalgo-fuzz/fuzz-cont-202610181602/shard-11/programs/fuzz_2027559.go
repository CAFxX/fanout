package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint16, p1 uint64, p2 uint8) bool {
    var v7 float64 = (0.0 * 0.0)
    v8 := GenS1{F0: p1, F1: p2}
    p2 = p2
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (((-231 >= 0) && (32767 >= -1)) || (!(0 <= 659)))
}
func genFn1() int32 {
    var v9 float32 = 0.1
    if (!((float32((1.5 + 0.0)) > (1e10 / float32(1e10))) && (true && false))) {
        v9 += v9
        var v10 uint8 = 115
        _ = v10
    }
    if false {
        v9 = (1.5 - ((v9 + 0.0) - 3.14159))
        var v11 uint8 = 128
        _ = v11
    } else {
        if (!true) {
            v9 = (v9 + float32(0.1))
        } else {
            v9 = (v9 + (-2.25 * -2.25))
        }
        v12 := []int{127}
        _ = v12
    }
    var v13 uint8 = 2
    _ = v13
    _ = v9
    return 1
}

func main() {
    v1 := []string{"\n\t"}
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    var v4 int32 = -894
    v5 := []string{"", ""}
    var v6 bool = (!false)
    v4 = (^(v4))
    switch 0 {
    case 0:
        v3[668] = ((("" + "\n\t") + "") + ("hello" + ("hello" + "\n\t")))
    case 1:
        var v14 float64 = float64(3.14159)
        _ = v14
    default:
        v4 = v4
    }
    v6 = v6
    v4 = (v4 >> ((959) & 63))
    switch "" {
    case "hello":
        v15 := GenS1{F0: 42, F1: 80}
        _ = v15
    }
    v16 := GenS1{F0: 42, F1: 42}
    v6 = v6
    var v17 uint = 4294967295
    var v18 uint = v17
    v19 := GenS1{F0: 674, F1: 60}
    var v20 uint = 0
    v21 := GenS1{F0: 851, F1: 255}
    v20 -= (v17 & v17)
    for fi22 := 0; fi22 < 8; fi22++ {
        fmt.Println(v6)
        fmt.Println(fi22, v4, v6)
        for fi23 := 0; fi23 < 11; fi23++ {
            _ = len(v5)
            v18 = v18
            if v6 {
                v24 := GenS2{F0: 610}
                v25 := GenS1{F0: 2, F1: 2}
                _ = v24
                _ = v25
            } else {
                var v26 float64 = (-2.25 * 3.14159)
                _ = v26
            }
            _ = fi23
        }
        _ = fi22
    }
    v27 := GenS1{F0: 4294967295, F1: 127}
    switch ("" + "a") {
    case "":
        switch ("\n\t" + "hello") {
        case "\n\t":
            v18 -= (v20 | v20)
        case "foo bar":
            _ = len(v2)
        }
    }
    var v28 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
