package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    v1 := []int{9223372036854775807, 2147483647, -544, 1}
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    v4 := make(map[string]bool)
    var v5 uint16 = 812
    v6 := []int{32767, -928, int(v5), 1}
    fmt.Println(v5)
    switch ("\n\t" + "foo bar") {
    case "a":
        var v7 int = 9223372036854775807
        _ = v7
    case "hello":
        switch 0 {
        case 4294967295:
            v8 := GenS1{F0: 2}
            _ = v8
        case 128:
            var v9 float64 = 0.0
            _ = v9
        }
    case "foo bar":
        if false {
            v5 = v5
            var v10 string = ("\n\t" + "\n\t")
            for fi11 := 0; fi11 < 8; fi11++ {
                if false {
                    if (!(0 > -208)) {
                        if (42 <= 42) {
                            v5 += v5
                            fmt.Println(v5)
                            v12 := GenS1{F0: 229}
                            _ = v12
                        }
                        var v13 uint = 0
                        fmt.Println(v13)
                        _ = v13
                    }
                    v5 = 820
                }
                if (v10 == v10) {
                    var v14 uint64 = 2
                    var v15 string = "hello"
                    var v16 float64 = ((float64((-2.25 / 3.14159)) + ((1.5 * 0.0) + 0.1)) - 0.0)
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    var v17 float64 = (float64(float32(1.5)) / 0.1)
                    _ = v17
                }
                _ = fi11
            }
            _ = v10
        }
    default:
        var v18 uint = 391
        _ = v18
    }
    v19 := GenS2{F0: 9223372036854775807, F1: 664}
    v5 = 158
    var v20 int8 = 42
    var v21 int64 = 9223372036854775807
    v20 = v20
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
