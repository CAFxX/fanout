package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int64 {
    var v5 int32 = 687
    v6 := GenS1{F0: 0}
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float32 = (0.1 - 1e10)
        v9 := []int{fi7, fi7}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v5)
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 float64 = ((0.0 + 0.1) - (-2.25 / float64(3.14159)))
    var v3 int8 = 1
    var v4 uint16 = 42
    if true {
        fmt.Println(v1, v1)
        if false {
            fmt.Println(v3, v2, v1)
            v10 := []int{}
            v10 = append(v10, -44)
            _ = v10
        }
    }
    var v11 string = ("foo bar" + "foo bar")
    for fi12 := 0; fi12 < 17; fi12++ {
        for fi13 := 0; fi13 < 11; fi13++ {
            v14 := GenS1{F0: -1}
            var v15 int = fi12
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(fi12)
        v16 := []int{(fi12 ^ (-(fi12))), (fi12 << ((262) & 63))}
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := []int{}
            if len(v19) > 0 {
                v19[0] = 32767
            }
            _ = fi18
            _ = v19
        }
        _ = fi17
    }
    v20 := []int{-553}
    for fi21 := 0; fi21 < 2; fi21++ {
        v4 -= v4
        var v22 int32 = -2147483648
        _ = fi21
        _ = v22
    }
    if (v2 <= 0.0) {
        fmt.Println(v11, v1, v11)
        var v23 int8 = -68
        v3 = v3
        _ = v23
    }
    var v24 uint16 = 2
    v11 = v11
    var v25 uint32 = 42
    if (128 > 41) {
        switch "a" {
        case "\n\t":
            var v26 uint16 = (v24 - 874)
            _ = v26
        default:
            v20 = append(v20, 128)
        }
    }
    var v27 int8 = -91
    fmt.Println(v3, v11, v11)
    var v28 string = "hello"
    v11 = v11
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v20
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
