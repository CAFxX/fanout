package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
func genFn0() int32 {
    v6 := []int{9223372036854775807, 153, 128}
    var v7 float64 = (3.14159 + 0.0)
    v8 := GenS1{F0: 1, F1: 1}
    _ = v6
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []int{}
    var v3 uint64 = 0
    v4 := []uint32{}
    var v5 int = 1
    var v9 int16 = 42
    var v10 int32 = 42
    for fi11 := 0; fi11 < 9; fi11++ {
        var v12 int = 2147483647
        switch "hello" {
        case "hello":
            v12 -= v5
        case "":
            v4 = append(v4, 482)
        }
        v13 := GenS1{F0: 32767, F1: 127}
        _ = fi11
        _ = v12
        _ = v13
    }
    switch "hello" {
    case "a":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        if (false || true) {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v5 = 127
            var v14 string = ""
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        v9 -= v9
        if true {
            v16 := GenS1{F0: 0, F1: 127}
            switch "a" {
            case "":
                var v17 int16 = -4
                _ = v17
            case "a":
                v5 += int(v3)
            case "\n\t":
                v3 += v3
            }
            _ = v16
        }
        var v18 uint = 1
        _ = fi15
        _ = v18
    }
    v9 += (v9 | -1)
    for fi19 := 0; fi19 < 1; fi19++ {
        switch v5 {
        case 42:
            fmt.Println(fi19, v3)
        default:
            var v20 float32 = 1.5
            _ = v20
        }
        _ = fi19
    }
    if false {
        v3 += 42
        v10 -= 1
        if false {
            var v21 float32 = 0.1
            v22 := v1[(("foo bar" + "hello") + "\n\t")]
            _ = v21
            _ = v22
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
