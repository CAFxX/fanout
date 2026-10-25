package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint {
    var v7 int64 = 42
    v8 := GenS1{F0: 572}
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 int32 = 1
    var v2 int64 = -1
    var v3 uint32 = 65535
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]int)
    v6 := make(map[int]string)
    switch "\n\t" {
    case "foo bar":
        v9 := v6[int(v2)]
        _ = v9
    }
    if (63 != 255) {
        v2 = 797
        v1 += (v1 / ((v1) | 1))
    }
    switch 395 {
    case 639:
        for fi10 := 0; fi10 < 1; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(v1, v2, v2)
                v2 -= v2
                _ = fi11
            }
            fmt.Println(v3)
            _ = fi10
        }
    case 1:
        var v12 uint = 235
        _ = v12
    default:
        var v13 int64 = v2
        _ = v13
    }
    for fi14 := 0; fi14 < 2; fi14++ {
        switch "hello" {
        case "":
            _ = len(v4)
        case "\n\t":
            fmt.Println(v2, fi14, v3)
        case "hello":
            v15 := GenS1{F0: fi14}
            _ = v15
        }
        _ = fi14
    }
    var v16 float32 = ((0.1 + float32((float64((float32(1e10) * (1e10 - 1e10))) * (1e10 * 0.0)))) + 1e10)
    switch "" {
    case "":
        v1 += (v1 & v1)
    case "\n\t":
        if (65535 != 42) {
            delete(v6, 127)
        } else {
            delete(v4, 456)
            switch "foo bar" {
            case "":
                v16 = 3.14159
            default:
                var v17 float32 = v16
                _ = v17
            }
            var v18 int64 = (v2 * v2)
            _ = v18
        }
    }
    var v19 bool = (!true)
    v19 = (v19 && (-115 != -1))
    switch -1 {
    case -250:
        v20 := GenS1{F0: -194}
        _ = v20
    case 128:
        var v21 bool = true
        _ = v21
    }
    var v22 uint16 = 163
    fmt.Println(v1)
    _ = len(v6)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
