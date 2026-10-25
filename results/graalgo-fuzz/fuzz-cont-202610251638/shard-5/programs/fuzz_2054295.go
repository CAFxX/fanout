package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int32, p1 uint8) int64 {
    var v5 uint64 = 292
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v5
    return 594
}

func main() {
    var v1 int = 32768
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[int]bool)
    var v6 string = ("" + "\n\t")
    v1 = (v1 + v1)
    fmt.Println(v1, v6)
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := GenS1{F0: 1000}
        v9 := GenS1{F0: 32767}
        var v10 uint8 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    switch 1 {
    case -505:
        v1 -= (v1 - 1)
    case 32768:
        v6 = ("" + (v6 + ""))
    case 13:
        if (((v6 + (v6 + v6)) + v6) == "\n\t") {
            if false {
                v11 := []int{v1}
                _ = len(v11)
                _ = v11
            } else {
                switch -1 {
                case 127:
                    switch (v1 / ((9223372036854775807) | 1)) {
                    case -1:
                        var v12 uint16 = 255
                        _ = v12
                    default:
                        v13 := GenS2{F0: 115}
                        _ = v13
                    }
                default:
                    v14 := GenS3{F0: 793, F1: 186}
                    _ = v14
                }
            }
            var v15 uint64 = 0
            v15 += v15
            _ = v15
        } else {
            v16 := []int{1, v1, v1}
            _ = v16
        }
    default:
        v6 = v6
    }
    switch v6 {
    case "a":
        var v17 int32 = 500
        _ = v17
    }
    v2[(v1 - (v1 * 1))] = 4294967295
    var v18 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
