package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int64 }
type GenS2 struct { F0 int32; F1 int; F2 int64 }
type GenS3 struct { F0 int }
func genFn0() int64 {
    var v4 int8 = 0
    for fi5 := 0; fi5 < 10; fi5++ {
        v4 = (v4 + (v4 & (v4 >> ((739) & 63))))
        _ = fi5
    }
    v4 = 84
    v6 := []int{-794, 42, 32767}
    fmt.Println(v4)
    _ = v4
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 float64 = (((0.1 - (0.0 + float64((-2.25 + 1e10)))) + 0.0) * (0.0 * 1.5))
    v2 := make(map[string]string)
    v3 := make(map[int]uint32)
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := GenS1{F0: (fi7 << ((0) & 63)), F1: 42, F2: 149}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int8 = -1
        v1 -= (3.14159 * -2.25)
        var v11 uint32 = 4294967295
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 += (v1 - v1)
    var v12 int = 127
    delete(v3, (v12 % (((v12 ^ (v12 ^ (-(v12))))) | 1)))
    var v13 float64 = float64(1.5)
    if false {
        for fi14 := 0; fi14 < 9; fi14++ {
            if (!(-441 > -1)) {
                v12 = 32767
            }
            v15 := GenS1{F0: fi14, F1: 233, F2: 79}
            var v16 uint16 = 892
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        var v17 int16 = 509
        var v18 bool = true
        _ = v17
        _ = v18
    }
    if (!true) {
        var v19 float32 = (0.1 * 0.1)
        for fi20 := 0; fi20 < 18; fi20++ {
            switch ("\n\t" + "\n\t") {
            case "hello":
                var v21 int = (v12 & v12)
                _ = v21
            case "foo bar":
                for fi22 := 0; fi22 < 7; fi22++ {
                    _ = len(v3)
                    _ = fi22
                }
            case "a":
                var v23 string = ""
                _ = v23
            }
            _ = fi20
        }
        _ = v19
    } else {
        fmt.Println(v13, v1)
        var v24 int16 = 926
        _ = v24
    }
    _ = len(v3)
    var v25 uint16 = 0
    var v26 string = "\n\t"
    for fi27 := 0; fi27 < 15; fi27++ {
        for fi28 := 0; fi28 < 6; fi28++ {
            v29 := GenS2{F0: -1, F1: fi28, F2: 9223372036854775807}
            var v30 int16 = -1
            _ = fi28
            _ = v29
            _ = v30
        }
        v31 := []int{-330, v12}
        _ = fi27
        _ = v31
    }
    var v32 float64 = (0.1 * v1)
    v33 := GenS3{F0: v12}
    var v34 int16 = 101
    switch "" {
    case "\n\t":
        fmt.Println(v34)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
