package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
func genFn0(p0 string) int8 {
    v6 := []int{}
    p0 = (p0 + p0)
    _ = p0
    _ = v6
    return -128
}
func genFn1(p0 int16, p1 uint32, p2 uint32) string {
    p0 += (-(p0))
    v7 := GenS1{F0: 1, F1: 9223372036854775807}
    if true {
        var v8 int16 = p0
        switch "\n\t" {
        case "foo bar":
            switch 1 {
            case -2147483648:
                var v9 int = 519
                _ = v9
            case -979:
                var v10 string = ""
                _ = v10
            case 2147483647:
                v11 := []int{-768, -667, -217}
                _ = v11
            default:
                switch ("foo bar" + "hello") {
                case "\n\t":
                    fmt.Println(p0)
                }
            }
        }
        switch 128 {
        case 63:
            p0 += 32767
        }
        _ = v8
    }
    var v12 string = "a"
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    return v12
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int = 341
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 float32 = 1e10
    var v13 float64 = -2.25
    if (42 == 1000) {
        v13 = (float64(v5) * float64(v5))
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        if false {
            if (!false) {
                var v15 string = "hello"
                switch fi14 {
                case 42:
                    v15 = ("a" + "foo bar")
                case -700:
                    if (!(128 != 127)) {
                        fmt.Println(fi14, v2, v5)
                        v5 += ((v5 * (v5 / (float32(0.0) + v5))) - 1e10)
                        v15 = v15
                    }
                case -386:
                    v2 += -130
                default:
                    if (!(2 <= 0)) {
                        var v16 int = (v2 | (fi14 >> ((42) & 63)))
                        v17 := GenS1{F0: -1, F1: -197}
                        var v18 int16 = -32768
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                _ = v15
            }
            v5 -= 1.5
        }
        if (!true) {
            var v19 int64 = -9223372036854775808
            v13 += (((v13 * (-2.25 - 0.1)) * 0.1) * 3.14159)
            var v20 uint32 = 944
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    var v21 int = v2
    var v22 string = "foo bar"
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 9; fi24++ {
            v25 := []int{fi24}
            _ = fi24
            _ = v25
        }
        delete(v1, 65535)
        _ = fi23
    }
    fmt.Println(v21)
    v2 += v2
    v22 = v22
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
