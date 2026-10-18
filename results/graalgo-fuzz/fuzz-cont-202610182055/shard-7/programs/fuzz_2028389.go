package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int }
func genFn0(p0 uint16, p1 uint16) uint32 {
    var v5 float64 = (float64(float32(1e10)) / (1e10 * 1e10))
    var v6 int64 = 42
    if true {
        v7 := GenS1{F0: true}
        if true {
            if (3.14159 >= (((3.14159 / (((1e10 * 0.0) - float32(0.0)) - 3.14159)) - float32(float64((((0.1 - 1e10) / 0.1) - (1e10 * ((0.1 + 0.1) + 0.1)))))) / 1e10)) {
                switch "foo bar" {
                case "a":
                    var v8 uint8 = 1
                    _ = v8
                case "\n\t":
                    v9 := []int{}
                    _ = v9
                default:
                    if (false || (0 != 0)) {
                        if true {
                            fmt.Println(p1, p0)
                        } else {
                            switch 42 {
                            case 0:
                                var v10 uint8 = 246
                                _ = v10
                            }
                            v11 := []int{32767}
                            _ = v11
                        }
                        p1 += (p0 - 32767)
                    }
                }
            } else {
                var v12 bool = false
                _ = v12
            }
            v13 := GenS1{F0: false}
            _ = v13
        }
        v14 := []int{4294967295}
        _ = v14
        _ = v7
    } else {
        var v15 bool = (0 >= 0)
        if (!v15) {
            p1 = (^(p0))
        } else {
            p1 -= p0
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            p0 = (p0 * p0)
            _ = fi16
        }
        _ = v15
    }
    if (665 < 1) {
        var v17 float32 = 0.0
        v18 := GenS2{F0: -1}
        v19 := []int{-868, -9223372036854775808}
        _ = v17
        _ = v18
        _ = v19
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 2
}

func main() {
    var v1 string = (("a" + "a") + ("\n\t" + "\n\t"))
    v2 := make(map[uint32]string)
    v3 := []uint32{2147483647, 449, 837}
    var v4 string = v1
    v20 := v2[1]
    var v21 string = ""
    var v22 uint16 = 42
    fmt.Println(v22)
    var v23 int8 = -128
    v24 := GenS2{F0: -944}
    v25 := v2[2147483647]
    v22 += v22
    var v26 float64 = -2.25
    var v27 uint32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
