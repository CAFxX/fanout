package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]int)
    v3 := []int{257, -125}
    v4 := []string{"foo bar", (("hello" + "\n\t") + "a"), "a"}
    var v5 int32 = -1
    fmt.Println(v1)
    v5 += 2147483647
    var v6 float32 = (1.5 * 1e10)
    if false {
        var v7 int = 2147483647
        v7 = 2147483647
        _ = v7
    } else {
        v5 -= (v5 * v5)
        if (0.1 <= 0.1) {
            var v8 uint = 2
            switch 4294967295 {
            case 674:
                v2[32767] = 9223372036854775807
            default:
                var v9 int32 = v5
                _ = v9
            }
            _ = v8
        } else {
            v10 := GenS1{F0: -473, F1: 0}
            if true {
                v1 = 42
            } else {
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (v5 < (v5 >> ((42) & 63))) {
                        var v12 float32 = ((v6 / v6) + v6)
                        var v13 bool = true
                        _ = v12
                        _ = v13
                    } else {
                        v14 := GenS1{F0: 971, F1: 18446744073709551615}
                        _ = v14
                    }
                    var v15 int32 = (v5 - v5)
                    v16 := GenS1{F0: 827, F1: 0}
                    _ = fi11
                    _ = v15
                    _ = v16
                }
            }
            if (!true) {
                switch "hello" {
                case "foo bar":
                    v2[109] = 42
                case "":
                    v2[9223372036854775807] = 949
                case "a":
                    fmt.Println(v6, v1)
                }
            }
            _ = v10
        }
    }
    if false {
        v6 = v6
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            var v18 uint8 = 127
            _ = fi17
            _ = v18
        }
    }
    var v19 bool = (127 < 1)
    v6 = v6
    var v20 int64 = (v1 & (v1 * 9223372036854775807))
    v21 := GenS1{F0: 42, F1: 18446744073709551615}
    v4 = append(v4, "foo bar")
    v6 = v6
    v1 = -9223372036854775808
    v20 = (v20 * v20)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
