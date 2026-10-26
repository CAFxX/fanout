package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int; F1 int8 }
func genFn0() int {
    v5 := GenS1{F0: "a", F1: 0}
    var v6 float64 = 0.1
    v6 = float64(1e10)
    v6 = v6
    for fi7 := 0; fi7 < 20; fi7++ {
        v6 = v6
        if (1 < (fi7 << ((14) & 63))) {
            var v8 string = ("hello" + ("hello" + "hello"))
            var v9 int16 = 1000
            v6 += v6
            _ = v8
            _ = v9
        } else {
            v10 := []int{(fi7 & (fi7 & (fi7 + (fi7 % (((fi7 / ((32768) | 1))) | 1))))), (fi7 >> ((4294967295) & 63)), (fi7 & (^(fi7))), (-(fi7))}
            var v11 uint16 = 0
            v12 := GenS1{F0: ("" + "\n\t"), F1: -71}
            _ = v10
            _ = v11
            _ = v12
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(fi7)
        case "":
            v6 += v6
        }
        _ = fi7
    }
    _ = v5
    _ = v6
    return 32768
}

func main() {
    var v1 string = ""
    var v2 string = ""
    v3 := make(map[int]uint32)
    var v4 int = -9223372036854775808
    v13 := GenS2{F0: 42, F1: 1}
    v14 := GenS1{F0: (v2 + v2), F1: int16(v4)}
    v15 := GenS1{F0: v1, F1: -958}
    v16 := GenS1{F0: ("" + (v2 + v1)), F1: 1000}
    v4 += genFn0()
    v17 := GenS3{F0: -103}
    v18 := GenS4{F0: (v4 >> ((0) & 63)), F1: 42}
    fmt.Println(v2, v4, v1)
    var v19 uint64 = 535
    for fi20 := 0; fi20 < 4; fi20++ {
        var v21 float32 = (1.5 + (-2.25 / -2.25))
        v22 := []int{v4}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 14; fi23++ {
        delete(v3, v4)
        var v24 float32 = 0.0
        v19 = v19
        _ = fi23
        _ = v24
    }
    var v25 string = v2
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint8 = 128
        _ = fi26
        _ = v27
    }
    var v28 int = -461
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
}
