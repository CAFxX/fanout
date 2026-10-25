package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint64, p1 uint8, p2 int16) uint {
    p0 += 4294967295
    v4 := GenS1{F0: (4294967295 <= 32768)}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 443
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]uint32)
    var v3 uint32 = 304
    switch 125 {
    case 32768:
        fmt.Println(v3)
    }
    fmt.Println(v3)
    v3 = 559
    fmt.Println(v3)
    var v5 bool = false
    v6 := v1[(^(v3))]
    for fi7 := 0; fi7 < 16; fi7++ {
        if v5 {
            for fi8 := 0; fi8 < 7; fi8++ {
                switch "\n\t" {
                case "":
                    v3 = (^(v3))
                case "a":
                    if (float32(float64(1.5)) > (3.14159 + (0.0 + 0.1))) {
                        var v9 int8 = 1
                        _ = v9
                    }
                case "hello":
                    switch "foo bar" {
                    case "":
                        var v10 float64 = 1e10
                        _ = v10
                    default:
                        v11 := GenS1{F0: v5}
                        fmt.Println(v6)
                        var v12 uint64 = 4294967295
                        _ = v11
                        _ = v12
                    }
                }
                v5 = (v5 && v5)
                var v13 uint8 = 185
                _ = fi8
                _ = v13
            }
        }
        v14 := GenS2{F0: -1}
        var v15 int64 = -1
        _ = fi7
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        if (v5 || v5) {
            var v17 uint = 462
            _ = v17
        }
        _ = fi16
    }
    v5 = (520 < 127)
    if (v5 && v5) {
        switch ("foo bar" + "hello") {
        case "foo bar":
            var v18 int64 = -1
            _ = v18
        }
    }
    var v19 string = ("foo bar" + ("" + "a"))
    v20 := GenS1{F0: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v5
    _ = v6
}
