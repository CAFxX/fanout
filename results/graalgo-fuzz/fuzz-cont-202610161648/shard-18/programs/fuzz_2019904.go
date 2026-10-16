package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 int8; F2 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]uint32)
    var v3 uint = 660
    var v4 uint64 = 1
    var v5 int64 = 42
    v6 := make(map[string]string)
    if true {
        if true {
            v4 -= v4
            switch "a" {
            case "a":
                for fi7 := 0; fi7 < 11; fi7++ {
                    fmt.Println(v3, v3)
                    var v8 float32 = float32((0.1 + (1.5 + float64(1e10))))
                    v9 := GenS1{F0: -1}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            case "":
                var v10 int = 4294967295
                _ = v10
            }
        }
        if (!true) {
            for fi11 := 0; fi11 < 10; fi11++ {
                if (false && false) {
                    var v12 int32 = -2147483648
                    _ = v12
                }
                var v13 uint = (v3 - v3)
                var v14 uint16 = 0
                _ = fi11
                _ = v13
                _ = v14
            }
            fmt.Println(v5, v5, v3)
            v5 -= 9223372036854775807
        }
        v15 := []int{-262, 42, 32768, 1}
        _ = v15
    }
    v16 := GenS2{F0: 0, F1: 14, F2: 0}
    fmt.Println(v3, v4, v4)
    v5 = 9223372036854775807
    var v17 uint32 = 2
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := v1["foo bar"]
        _ = fi18
        _ = v19
    }
    v17 = v17
    switch ("" + "foo bar") {
    case "hello":
        v3 -= v3
    case "\n\t":
        v4 += (v4 + v4)
    case "foo bar":
        for fi20 := 0; fi20 < 2; fi20++ {
            _ = len(v1)
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 uint = v3
                var v23 float32 = (0.0 + 1e10)
                _ = fi21
                _ = v22
                _ = v23
            }
            v24 := []int{128, fi20, fi20, (-(fi20))}
            _ = fi20
            _ = v24
        }
    }
    v3 -= v3
    var v25 int = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
