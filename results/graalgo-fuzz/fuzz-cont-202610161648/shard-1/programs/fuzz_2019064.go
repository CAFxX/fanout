package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int64; F1 string }
func genFn0() int {
    var v5 int64 = 0
    switch "a" {
    case "a":
        var v6 uint8 = 2
        _ = v6
    default:
        v5 -= (v5 << ((1) & 63))
    }
    _ = v5
    return 199
}

func main() {
    v1 := []uint32{2147483648, 1, 0}
    v2 := make(map[uint32]string)
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]bool)
    switch ("hello" + "foo bar") {
    case "\n\t":
        var v7 int64 = 52
        _ = v7
    case "foo bar":
        var v8 int32 = 2147483647
        _ = v8
    case "hello":
        v1 = append(v1, 0)
    }
    var v9 int = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            v12 := GenS1{F0: 749, F1: 1}
            var v13 bool = (899 > 2)
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 11; fi14++ {
            _ = len(v1)
            delete(v4, 260)
            if (4294967295 <= 555) {
                fmt.Println(fi10)
                var v15 uint32 = 2147483647
                _ = v15
            }
            _ = fi14
        }
        _ = fi10
    }
    if (true || false) {
        fmt.Println(v9)
    } else {
        if false {
            v16 := GenS2{F0: -74, F1: "a"}
            _ = v16
        }
        v9 = (v9 + v9)
    }
    var v17 int8 = -1
    v3[(("foo bar" + (((("a" + ("hello" + "\n\t")) + (("" + "") + (("\n\t" + "foo bar") + "a"))) + "\n\t") + ("hello" + "a"))) + "hello")] = 2
    for fi18 := 0; fi18 < 4; fi18++ {
        v17 -= v17
        for fi19 := 0; fi19 < 15; fi19++ {
            v17 = v17
            _ = fi19
        }
        _ = fi18
    }
    switch "a" {
    case "hello":
        v1 = append(v1, 42)
    case "\n\t":
        var v20 float64 = float64(0.1)
        _ = v20
    default:
        for fi21 := 0; fi21 < 17; fi21++ {
            delete(v4, 754)
            _ = fi21
        }
    }
    var v22 float32 = ((-2.25 / (0.0 + ((-2.25 - float32((3.14159 * (1.5 / (-2.25 + 0.0))))) - 3.14159))) - (3.14159 / 1.0))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
