package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "a"
    var v3 int8 = -1
    var v4 bool = true
    var v5 int8 = 1
    v6 := make(map[string]string)
    if v4 {
        var v7 string = "a"
        v8 := GenS1{F0: 0}
        v2 = "hello"
        _ = v7
        _ = v8
    }
    if (42 != 2) {
        v4 = (!(!true))
        for fi9 := 0; fi9 < 10; fi9++ {
            if (42 != -466) {
                v4 = (3.14159 > 1.5)
                switch -2147483648 {
                case 128:
                    for fi10 := 0; fi10 < 6; fi10++ {
                        fmt.Println(v4, v4, v3)
                        _ = fi10
                    }
                default:
                    v2 = "hello"
                }
                v11 := GenS1{F0: 0}
                _ = v11
            }
            var v12 int64 = 9223372036854775807
            _ = fi9
            _ = v12
        }
    }
    v13 := v6[v2]
    v13 = (v2 + v2)
    if v4 {
        var v14 uint16 = 0
        fmt.Println(v14, v14)
        var v15 uint64 = 18446744073709551615
        _ = v14
        _ = v15
    } else {
        var v16 int16 = 1000
        switch 127 {
        case 0:
            v3 = -33
        default:
            if (188 > 127) {
                v4 = (v4 && v4)
            }
        }
        v5 = -1
        _ = v16
    }
    var v17 int = 32767
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := []int{(v17 + fi18)}
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 uint16 = 823
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    _ = len(v6)
    v13 = ""
    switch v2 {
    case "a":
        var v22 uint32 = 4294967295
        _ = v22
    case "hello":
        if v4 {
            v23 := GenS1{F0: -239}
            v3 = 1
            fmt.Println(v4)
            _ = v23
        } else {
            fmt.Println(v4, v17, v5)
            v24 := []int{-2147483648, -699, v17}
            for fi25 := 0; fi25 < 2; fi25++ {
                v17 += v17
                v1[138] = (-948 < 59)
                fmt.Println(v17, v13, v17)
                _ = fi25
            }
            _ = v24
        }
    case "":
        v5 -= (v5 + v3)
    }
    var v26 bool = true
    v27 := v1[339]
    if v4 {
        fmt.Println(v17)
        var v28 float32 = 1.5
        var v29 float64 = (0.0 * 1e10)
        _ = v28
        _ = v29
    }
    fmt.Println(v4, v13, v17)
    v30 := GenS2{F0: ((true && (1.5 > ((-2.25 - float64(1e10)) * 1e10))) && (-914 != 42)), F1: 255}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
