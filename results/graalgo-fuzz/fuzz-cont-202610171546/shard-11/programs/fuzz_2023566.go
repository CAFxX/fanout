package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = -790
    v3 := make(map[int]int)
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    fmt.Println(v2)
    var v4 bool = false
    v5 := GenS1{F0: 18446744073709551615, F1: -125}
    for fi6 := 0; fi6 < 14; fi6++ {
        if (65535 > 2) {
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 uint64 = 1
                _ = len(v1)
                _ = fi7
                _ = v8
            }
            var v9 int8 = 27
            _ = v9
        } else {
            if (!((!(21 == 119)) || ((!(!true)) || (0 == 154)))) {
                fmt.Println(v2, v4, v4)
            }
            fmt.Println(fi6, v4, v2)
        }
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(fi6, fi6, v4)
            if (42 <= 892) {
                switch fi6 {
                case 2147483647:
                    v11 := GenS1{F0: 1, F1: 100}
                    var v12 int = -14
                    _ = v11
                    _ = v12
                case 42:
                    var v13 int8 = 127
                    _ = v13
                case -2147483648:
                    var v14 bool = v4
                    v4 = (("hello" + ("foo bar" + "\n\t")) >= "foo bar")
                    var v15 int = (fi6 % ((fi10) | 1))
                    _ = v14
                    _ = v15
                default:
                    fmt.Println(v4, v2)
                }
                v16 := GenS2{F0: 65535, F1: 32768}
                v17 := GenS3{F0: 255, F1: true}
                _ = v16
                _ = v17
            }
            _ = fi10
        }
        _ = fi6
    }
    v4 = (1e10 >= (3.14159 + (0.0 * 0.0)))
    v18 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
