package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int8, p1 uint8, p2 uint8) uint64 {
    p1 -= 41
    v4 := []int{899, -2147483648, 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 327
}
func genFn1() int64 {
    v5 := GenS1{F0: 42, F1: 4294967295}
    if true {
        var v6 float64 = (1e10 + (-2.25 + ((-2.25 - (1.5 * (-2.25 - float64(float32((0.0 - (-2.25 / 1.5))))))) - (-2.25 / 3.14159))))
        _ = v6
    } else {
        var v7 int32 = 1
        if false {
            var v8 uint = 2
            switch "a" {
            case "a":
                for fi9 := 0; fi9 < 11; fi9++ {
                    v8 = 4294967295
                    _ = fi9
                }
            case "\n\t":
                v8 = 1
            case "":
                var v10 int8 = 0
                _ = v10
            }
            v8 -= (v8 % (((v8 ^ (v8 * v8))) | 1))
            _ = v8
        }
        v11 := GenS1{F0: 1, F1: 0}
        _ = v11
        _ = v7
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v13 := []int{(fi12 % ((fi12) | 1)), (^(fi12))}
        if (false && true) {
            var v14 int16 = 1
            _ = v14
        }
        fmt.Println(fi12)
        _ = fi12
        _ = v13
    }
    var v15 float32 = (3.14159 - (3.14159 - (3.14159 / 1e10)))
    _ = v15
    _ = v5
    return 1
}

func main() {
    v1 := []string{(("hello" + ("hello" + "foo bar")) + "a")}
    v2 := []int{}
    var v3 string = "\n\t"
    var v16 float32 = ((3.14159 + (((1.5 + float32((-2.25 * float64((((0.1 - (((0.0 / 1.0) / (-2.25 / 1.5)) * ((0.0 / 1.0) - 1.5))) / float32(0.1)) + 1.5))))) * (float32(0.1) - 1.5)) / ((3.14159 * 3.14159) + (0.0 - 0.0)))) - float32(0.1))
    for fi17 := 0; fi17 < 20; fi17++ {
        v16 = (0.1 / (float32(1e10) * (-2.25 - 0.1)))
        fmt.Println(fi17, v16, v3)
        var v18 string = v3
        _ = fi17
        _ = v18
    }
    if true {
        var v19 int32 = 2147483647
        fmt.Println(v16, v3, v3)
        _ = v19
    } else {
        if ((!(!false)) && (false && false)) {
            _ = len(v2)
        }
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: 1, F1: 0}
            _ = fi20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(fi22, v3, v16)
        if false {
            _ = len(v1)
        }
        v2 = append(v2, fi22)
        _ = fi22
    }
    v2 = append(v2, 167)
    v16 = v16
    v16 = float32(float64(float32(0.1)))
    for fi23 := 0; fi23 < 4; fi23++ {
        var v24 float64 = (3.14159 + 0.0)
        if (-32768 < 42) {
            var v25 int32 = -347
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    fmt.Println(v3, v3)
    var v26 uint16 = 32767
    var v27 float64 = float64(v16)
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v26
    _ = v27
    _ = v3
}
