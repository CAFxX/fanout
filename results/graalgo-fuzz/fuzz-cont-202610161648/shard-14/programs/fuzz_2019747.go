package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool; F1 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 int) uint16 {
    p0 -= 128
    v6 := GenS1{F0: -1, F1: true}
    v7 := GenS1{F0: 0, F1: ((!true) || (!(0.1 >= (((3.14159 * 0.0) + 3.14159) - float64((-2.25 * 1.5))))))}
    p0 += p0
    var v8 int = (p0 / ((p0) | 1))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 string = ((("hello" + "foo bar") + ("hello" + "hello")) + "")
    v2 := make(map[string]uint32)
    var v3 float64 = (0.1 + (0.1 / 1.5))
    var v4 uint32 = 42
    var v5 float64 = v3
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS2{F0: true, F1: 128}
        var v11 float64 = v3
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS3{F0: -128}
    v13 := []int{-2147483648}
    v5 = float64(((1e10 * 0.0) + 1e10))
    var v14 string = "a"
    var v15 uint = 870
    v16 := GenS1{F0: -39, F1: (v5 >= 1e10)}
    v17 := v2[(v14 + ((v1 + v1) + (("a" + v14) + "")))]
    v1 = v1
    v18 := GenS2{F0: false, F1: 245}
    if true {
        v19 := GenS3{F0: 100}
        for fi20 := 0; fi20 < 6; fi20++ {
            v5 -= v3
            v15 = uint(v4)
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi20
        }
        v21 := GenS1{F0: -120, F1: ((false || true) || true)}
        _ = v19
        _ = v21
    }
    var v22 string = v14
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
