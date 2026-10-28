package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint) uint8 {
    switch "\n\t" {
    case "a":
        var v7 int = 42
        _ = v7
    case "":
        v8 := GenS1{F0: 0, F1: false, F2: -32768}
        _ = v8
    case "foo bar":
        if false {
            fmt.Println(p0)
        } else {
            switch -1 {
            case 127:
                fmt.Println(p0)
            case -2147483648:
                p0 = p0
            case 4294967295:
                v9 := []int{127}
                _ = v9
            }
        }
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 127, F1: 127}
        v12 := []int{(fi10 ^ fi10), 32767}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    return 255
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{-534, -225, 0}
    v3 := []string{v1, "foo bar"}
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    var v6 int8 = 0
    var v13 uint8 = genFn0(42)
    v14 := GenS1{F0: 65535, F1: ((true && (!(0.1 <= -2.25))) && (0 > int16(v6))), F2: -32768}
    if true {
        if true {
            var v15 float32 = float32((1.5 - 1.5))
            switch 4294967295 {
            case -710:
                var v16 float64 = 0.0
                _ = v16
            case -1:
                if (!(2147483647 >= 0)) {
                    v17 := GenS3{F0: (v13 + 0), F1: 32768}
                    _ = v17
                }
            case -168:
                v18 := GenS4{F0: v13}
                _ = v18
            default:
                v1 = ""
            }
            _ = len(v2)
            _ = v15
        } else {
            var v19 uint32 = 422
            var v20 bool = (32767 != 673)
            var v21 float64 = (float64((((0.1 + ((1.5 * 1.5) - (1e10 * 1.5))) + -2.25) - (1.5 / ((1.5 + float32(-2.25)) / float32(1e10))))) - (0.0 - (3.14159 + 1e10)))
            _ = v19
            _ = v20
            _ = v21
        }
        var v22 string = "a"
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        v24 := GenS5{F0: 65535, F1: 42}
        for fi25 := 0; fi25 < 6; fi25++ {
            v6 = (^(v6))
            _ = fi25
        }
        _ = fi23
        _ = v24
    }
    v1 = v1
    switch v1 {
    case "a":
        v26 := GenS5{F0: 2, F1: 2}
        _ = v26
    }
    fmt.Println(v6, v13, v13)
    fmt.Println(v13, v6)
    _ = len(v5)
    for fi27 := 0; fi27 < 11; fi27++ {
        var v28 bool = true
        var v29 uint64 = 0
        v6 += v6
        _ = fi27
        _ = v28
        _ = v29
    }
    v6 = 100
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
