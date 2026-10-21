package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8 }
func genFn0(p0 int32, p1 uint) uint {
    var v6 uint = p1
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            switch "hello" {
            case "":
                if false {
                    var v9 int32 = (p0 + (p0 * (p0 ^ p0)))
                    v10 := []int{-1, int(v9)}
                    _ = v10
                    _ = v9
                }
            case "hello":
                if false {
                    var v11 uint64 = 42
                    _ = v11
                }
            }
            v6 -= 255
            _ = fi8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v6
    return v6
}

func main() {
    var v1 float32 = -2.25
    var v2 int16 = -1
    v3 := make(map[int]uint32)
    var v4 float64 = 0.1
    var v5 float32 = v1
    if false {
        var v12 float64 = v4
        _ = v12
    } else {
        if (2 > 1) {
            v4 += v4
            if true {
                v2 = (v2 ^ -594)
                v13 := GenS1{F0: 65535, F1: 127}
                _ = v13
            } else {
                var v14 float64 = 1e10
                var v15 float64 = v14
                v15 = (v4 + ((v14 + 3.14159) / float64(v1)))
                _ = v14
                _ = v15
            }
            var v16 uint = 42
            _ = v16
        }
    }
    v17 := GenS1{F0: 2147483648, F1: 42}
    fmt.Println(v1, v1, v4)
    var v18 int8 = 65
    v3[-855] = 2
    var v19 uint16 = 32767
    v5 -= (v5 / v5)
    var v20 int64 = -1
    v3[32767] = 2147483647
    v2 -= v2
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 int = fi21
        var v23 string = (("\n\t" + ("\n\t" + "")) + "foo bar")
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
