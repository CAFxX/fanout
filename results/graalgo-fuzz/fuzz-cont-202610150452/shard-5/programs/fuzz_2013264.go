package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int32, p1 uint8, p2 uint32) int {
    var v6 int16 = -32768
    var v7 int = 378
    var v8 string = ("a" + (("hello" + "hello") + "\n\t"))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1(p0 uint32) int8 {
    var v9 uint64 = 1
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := []int{-994}
        fmt.Println(v9, p0, v9)
        _ = fi10
        _ = v11
    }
    if ((0 <= 42) && true) {
        var v12 uint16 = 65535
        _ = v12
    }
    v13 := GenS1{F0: -12, F1: v9}
    v14 := GenS2{F0: 127}
    _ = p0
    _ = v13
    _ = v14
    _ = v9
    return -1
}

func main() {
    var v1 string = ("hello" + ("" + ("\n\t" + ((("\n\t" + ("foo bar" + "a")) + ("foo bar" + (("foo bar" + "foo bar") + ("foo bar" + ("foo bar" + "\n\t"))))) + ""))))
    var v2 float64 = ((-2.25 - ((1e10 * ((float64(0.1) - 0.1) - 1e10)) * 0.0)) + ((0.0 * float64(((-2.25 / (1.5 / ((3.14159 / 1.0) + (1.5 - (((3.14159 - -2.25) / (1e10 - (-2.25 + (3.14159 + 1e10)))) * 0.1))))) - 1.5))) - 1e10))
    v3 := make(map[string]int)
    var v4 int16 = 42
    var v5 bool = (1 < 0)
    v15 := GenS3{F0: v5}
    var v16 int = 0
    v17 := GenS2{F0: 17}
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(fi18)
        _ = fi18
    }
    var v19 int = genFn0(2147483647, 2, 42)
    for fi20 := 0; fi20 < 1; fi20++ {
        for fi21 := 0; fi21 < 15; fi21++ {
            v5 = v5
            v22 := GenS4{F0: 678}
            var v23 uint32 = 65535
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    v2 = ((1.5 * v2) / (v2 / v2))
    if v5 {
        v24 := v3["hello"]
        v25 := v3[((("a" + "foo bar") + v1) + (v1 + v1))]
        _ = v24
        _ = v25
    } else {
        var v26 uint64 = 42
        v27 := GenS5{F0: (v19 >> ((655) & 63)), F1: (-(v16))}
        _ = v26
        _ = v27
    }
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
