package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int16; F1 int64 }
func genFn0() int {
    if true {
        v6 := GenS1{F0: -933}
        _ = v6
    } else {
        for fi7 := 0; fi7 < 17; fi7++ {
            for fi8 := 0; fi8 < 16; fi8++ {
                switch "a" {
                case "foo bar":
                    var v9 float32 = float32((float64(1e10) * 0.0))
                    _ = v9
                case "a":
                    var v10 string = ""
                    _ = v10
                case "\n\t":
                    if (42 == 1) {
                        fmt.Println(fi8, fi7)
                        v11 := []int{}
                        fmt.Println(fi8, fi8)
                        var v12 float32 = 0.1
                        _ = v11
                        _ = v12
                    } else {
                        fmt.Println(fi8, fi7)
                    }
                }
                var v13 string = "a"
                if true {
                    v14 := GenS2{F0: (fi8 % (((fi8 | 32768)) | 1)), F1: 40}
                    _ = v14
                }
                fmt.Println(fi8)
                v15 := []int{}
                _ = fi8
                _ = v13
                _ = v15
            }
            _ = fi7
        }
        var v16 string = "foo bar"
        v17 := []int{1, 9223372036854775807}
        _ = v16
        _ = v17
    }
    var v18 float32 = 3.14159
    _ = v18
    return 2147483647
}

func main() {
    var v1 float64 = (((1e10 - 0.1) + (1.5 / 1.0)) - (-2.25 * (0.0 * float64(-2.25))))
    var v2 uint16 = 54
    v3 := make(map[int]string)
    var v4 int8 = 1
    var v5 int = 127
    v19 := []int{}
    fmt.Println(v2, v4, v5)
    if false {
        v20 := v3[v5]
        v1 = 3.14159
        fmt.Println(v5, v4)
        _ = v20
    }
    var v21 uint32 = 55
    v22 := GenS3{F0: 1, F1: 605}
    var v23 int64 = -642
    var v24 int64 = v23
    if len(v19) > 0 {
        v19[0] = -1
    }
    var v25 uint64 = 863
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
