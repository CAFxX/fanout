package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string) int8 {
    var v6 uint = 924
    v7 := []int{32767}
    switch p0 {
    case "a":
        switch 661 {
        case 0:
            var v8 float32 = (float32((0.0 + 0.1)) + (1.5 + 1e10))
            _ = v8
        }
    default:
        var v9 uint = v6
        _ = v9
    }
    p0 = "a"
    _ = p0
    _ = v6
    _ = v7
    return 57
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 uint64 = 175
    v4 := make(map[uint32]string)
    v5 := make(map[string]int)
    var v10 float64 = (-2.25 + ((1e10 - 1.5) + (1.5 * 1.5)))
    v3 = (v3 / ((v3) | 1))
    v3 -= (^(v3))
    switch "hello" {
    case "foo bar":
        v10 += v10
    case "\n\t":
        var v11 float32 = 0.0
        _ = v11
    default:
        v3 += v3
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        fmt.Println(fi12, v3, v10)
        _ = fi12
    }
    v13 := GenS1{F0: "\n\t"}
    v14 := GenS1{F0: ("\n\t" + (("" + ("\n\t" + ((("hello" + "hello") + "foo bar") + ("\n\t" + (("\n\t" + "a") + ("a" + ("\n\t" + "a"))))))) + "foo bar"))}
    fmt.Println(v3, v3)
    fmt.Println(v3, v10)
    for fi15 := 0; fi15 < 1; fi15++ {
        v16 := GenS2{F0: -763, F1: 32767}
        var v17 int32 = -1
        v17 += v17
        _ = fi15
        _ = v16
        _ = v17
    }
    if (!(1.5 > 1.5)) {
        v10 -= float64(-2.25)
        v10 = v10
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v10 -= (0.0 * ((v10 * 1.5) * (float64((0.1 + 1e10)) + (float64((0.1 - (float32(v10) * (3.14159 / 1.5)))) + 1e10))))
        var v19 uint8 = 255
        var v20 float32 = ((-2.25 * ((-2.25 - -2.25) - 0.1)) / 1e10)
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: -9223372036854775808}
    fmt.Println(v10, v3)
    var v22 float32 = ((-2.25 + 1.5) + 1e10)
    _ = len(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
