package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int; F1 string; F2 int8 }
type GenS3 struct { F0 int8; F1 bool; F2 uint64 }
func genFn0() uint8 {
    if false {
        if (false && true) {
            var v4 bool = false
            v4 = (!true)
            _ = v4
        } else {
            v5 := GenS1{F0: 127}
            var v6 string = "hello"
            _ = v5
            _ = v6
        }
    }
    if true {
        for fi7 := 0; fi7 < 20; fi7++ {
            var v8 int = (fi7 & fi7)
            for fi9 := 0; fi9 < 2; fi9++ {
                switch "\n\t" {
                case "foo bar":
                    var v10 int32 = 42
                    _ = v10
                case "\n\t":
                    var v11 float64 = (3.14159 / (1e10 + float64((((-2.25 + 1e10) * (0.1 / float32(1e10))) / 1.0))))
                    _ = v11
                case "hello":
                    v12 := GenS2{F0: fi9, F1: "", F2: 42}
                    _ = v12
                default:
                    v13 := []int{}
                    _ = v13
                }
                v8 = (^(v8))
                v8 = (fi9 * (v8 - 9223372036854775807))
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    var v14 int64 = 522
    var v15 int8 = 1
    v16 := GenS3{F0: (v15 & -128), F1: (false || ((42 > 42) && (-793 != 127))), F2: 454}
    _ = v14
    _ = v15
    _ = v16
    return 42
}

func main() {
    v1 := []string{"a"}
    var v2 float32 = float32(1e10)
    var v3 uint16 = 2
    v1 = append(v1, "foo bar")
    v3 -= 1
    var v17 uint64 = 0
    v3 = v3
    var v18 uint32 = 0
    var v19 float32 = (v2 + (0.1 - v2))
    var v20 int64 = 553
    var v21 int8 = 1
    var v22 uint32 = 4294967295
    var v23 uint32 = 443
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
