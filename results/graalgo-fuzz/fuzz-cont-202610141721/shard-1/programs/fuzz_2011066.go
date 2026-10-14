package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 0
    var v3 float32 = (((0.0 * float32((0.0 - (1e10 + -2.25)))) / (-2.25 * (1e10 * 3.14159))) / (-2.25 / -2.25))
    v4 := make(map[string]bool)
    v5 := make(map[int]int)
    var v6 int16 = -960
    for fi7 := 0; fi7 < 7; fi7++ {
        if false {
            var v8 uint32 = 0
            _ = len(v1)
            v9 := []int{583}
            _ = v8
            _ = v9
        }
        _ = fi7
    }
    fmt.Println(v6, v3, v6)
    switch 4294967295 {
    case 32767:
        var v10 bool = true
        _ = v10
    default:
        v11 := []int{}
        _ = v11
    }
    var v12 uint = 42
    v5[-837] = 0
    v13 := v5[4294967295]
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 uint64 = (v2 | v2)
        v16 := GenS1{F0: (((^(v6)) <= (v6 << (((v12 & (v12 ^ v12))) & 63))) || false), F1: -376, F2: ("hello" + "\n\t")}
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 -= 4294967295
    var v17 int64 = -190
    fmt.Println(v2, v12)
    v12 = (v12 << ((v12) & 63))
    v1["hello"] = false
    v18 := []int{v13, (v13 % (((v13 >> ((v12) & 63))) | 1)), (^(v13))}
    fmt.Println(v3, v12)
    if (false && ((^(v2)) == 42)) {
        var v19 string = ("" + "")
        _ = v19
    } else {
        v17 -= (^(v17))
        v20 := GenS2{F0: ("\n\t" + ""), F1: 4294967295, F2: 42}
        v21 := GenS3{F0: (-(v17)), F1: 517}
        _ = v20
        _ = v21
    }
    var v22 int = (v13 + v13)
    var v23 int16 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
