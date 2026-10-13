package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int8 }
type GenS2 struct { F0 int; F1 int8; F2 bool }
type GenS3 struct { F0 uint8 }
func genFn0() int {
    v7 := []int{127}
    v8 := GenS1{F0: 255, F1: 0, F2: -1}
    if (true || true) {
        var v9 float32 = -2.25
        _ = v9
    } else {
        v10 := GenS1{F0: 0, F1: 2, F2: 100}
        if false {
            var v11 float64 = float64(0.0)
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 1, F1: 101, F2: -128}
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = v10
    }
    var v14 int8 = 87
    _ = v14
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]int)
    var v3 int = 425
    v4 := []uint32{1, 1, 0, 630}
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]string)
    if len(v4) > 0 {
        v4[0] = 2
    }
    var v15 float32 = 0.1
    if ((("foo bar" + "a") <= "a") && false) {
        switch "\n\t" {
        case "":
            v16 := GenS2{F0: (v3 + genFn0()), F1: -128, F2: (32767 > 1)}
            _ = v16
        case "a":
            var v17 bool = false
            _ = v17
        }
    }
    var v18 uint = 425
    v18 = v18
    _ = len(v4)
    if true {
        switch genFn0() {
        case -2147483648:
            switch "foo bar" {
            case "hello":
                if true {
                    var v19 uint32 = 2
                    _ = v19
                }
            default:
                v18 -= (v18 % ((v18) | 1))
            }
        }
    } else {
        v3 -= v3
    }
    v1 += v1
    delete(v6, "a")
    if false {
        var v20 uint8 = 2
        _ = v20
    }
    v21 := GenS3{F0: 128}
    v22 := GenS2{F0: -597, F1: 6, F2: (!true)}
    for fi23 := 0; fi23 < 15; fi23++ {
        v6[""] = "foo bar"
        fmt.Println(v18, v1, v18)
        _ = fi23
    }
    var v24 uint = v18
    switch v3 {
    case 127:
        v3 -= genFn0()
    case 9223372036854775807:
        var v25 float32 = (v15 + (v15 - 0.1))
        _ = v25
    case 42:
        var v26 uint32 = 42
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
