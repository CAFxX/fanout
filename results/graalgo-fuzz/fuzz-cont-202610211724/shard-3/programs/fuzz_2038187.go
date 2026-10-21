package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32 }
func genFn0(p0 int32) int64 {
    var v7 int64 = 0
    fmt.Println(v7, v7)
    _ = p0
    _ = v7
    return (v7 & 969)
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (((-2.25 - -2.25) - (1e10 - (3.14159 / 1.5))) + ((float64(((0.1 - (-2.25 - 1e10)) - 0.0)) / (1.5 * float64(1.5))) - -2.25))
    v3 := make(map[string]int)
    v4 := make(map[string]string)
    var v5 int64 = -1
    var v6 float32 = -2.25
    if true {
        fmt.Println(v6)
        var v8 int16 = 764
        var v9 uint64 = 1
        _ = v8
        _ = v9
    }
    delete(v4, (("\n\t" + "hello") + "\n\t"))
    if false {
        fmt.Println(v5, v2)
    }
    switch "a" {
    case "a":
        v2 -= float64(0.1)
    case "foo bar":
        if true {
            var v10 int = 2147483647
            v6 += (float32(1.5) + v6)
            _ = v10
        } else {
            v5 = v5
        }
    default:
        v11 := []int{-9223372036854775808, 128, 217}
        _ = v11
    }
    v5 = (v5 % ((genFn0(-434)) | 1))
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int = fi12
        v14 := []int{(v13 + fi12), (v13 | fi12), -898, 9223372036854775807}
        for fi15 := 0; fi15 < 10; fi15++ {
            fmt.Println(v13, fi12)
            if false {
                var v16 uint32 = 442
                v16 += (^(v16))
                _ = v16
            }
            if len(v14) > 0 {
                v14[0] = 2147483647
            }
            _ = fi15
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int64 = genFn0(799)
        _ = fi17
        _ = v18
    }
    var v19 float64 = v2
    v20 := GenS1{F0: 762}
    v21 := GenS2{F0: 127}
    for fi22 := 0; fi22 < 14; fi22++ {
        v2 = (v2 / v19)
        v23 := GenS3{F0: 660}
        _ = fi22
        _ = v23
    }
    var v24 bool = (!true)
    var v25 uint = 65535
    var v26 int = -374
    v5 = (v5 / ((-1) | 1))
    var v27 uint64 = 0
    var v28 float64 = (1.5 + ((float64(-2.25) * float64((1e10 / (float32(float64(float32(float64((float32(v19) / 1.5))))) * v6)))) - v2))
    v24 = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
