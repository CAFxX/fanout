package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int }
func genFn0(p0 int, p1 uint16, p2 string) bool {
    fmt.Println(p2, p0)
    var v7 int64 = -1
    var v8 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 bool = (1000 == 42)
    var v2 float64 = ((3.14159 / ((float64((-2.25 / 1.0)) * 1e10) * -2.25)) * ((((1e10 + 1e10) + -2.25) - 0.0) * 0.0))
    v3 := []int{128}
    var v4 float32 = (0.1 + (-2.25 * 0.0))
    var v5 uint64 = 4294967295
    var v6 int32 = 39
    v5 = v5
    switch "" {
    case "foo bar":
        v6 = int32(v5)
    }
    switch 268 {
    case 127:
        v2 = v2
    case -425:
        v2 += -2.25
    }
    v5 = v5
    var v9 int = 683
    v2 = (((v2 / v2) * 0.1) - (((1e10 / v2) * float64(v4)) / float64(0.1)))
    var v10 uint16 = 78
    v3 = append(v3, v9)
    if (v1 && (("hello" + "\n\t") > "foo bar")) {
        var v11 bool = true
        if v11 {
            v12 := GenS1{F0: -37}
            var v13 uint16 = 672
            _ = v12
            _ = v13
        }
        v14 := GenS2{F0: 0}
        _ = v11
        _ = v14
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: 32767}
            _ = fi15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS3{F0: 1, F1: v9}
        v4 = (v4 + (1e10 / v4))
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int16 = 32767
            fmt.Println(fi17)
            fmt.Println(v2, v4, v6)
            _ = fi19
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v1 = v1
        _ = fi21
    }
    var v22 int8 = 100
    var v23 int16 = -654
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
