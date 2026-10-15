package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint8) uint16 {
    v5 := GenS1{F0: 470}
    fmt.Println(p0)
    v6 := GenS1{F0: -217}
    var v7 int32 = -1
    v7 = (v7 % ((v7) | 1))
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 690
}
func genFn1(p0 uint16) uint16 {
    p0 += genFn0(42)
    fmt.Println(p0)
    var v8 string = "hello"
    v9 := GenS1{F0: 1000}
    v8 = "\n\t"
    _ = p0
    _ = v8
    _ = v9
    return 12
}

func main() {
    v1 := []uint32{994}
    v2 := []uint32{42, 957}
    v3 := make(map[int]string)
    var v4 uint = 2
    var v10 string = "hello"
    v10 = ""
    v11 := GenS2{F0: 4294967295, F1: 2}
    var v12 int16 = -1
    _ = len(v3)
    if true {
        var v13 uint32 = 4294967295
        v12 -= 1000
        _ = v13
    } else {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 uint32 = 2147483648
            if ((100 < 1) && false) {
                v10 = (("\n\t" + v10) + "foo bar")
                var v16 int16 = v12
                _ = v16
            }
            v17 := GenS2{F0: (v15 >> (((v4 + v4)) & 63)), F1: (v15 & (v15 + 2147483648))}
            _ = fi14
            _ = v15
            _ = v17
        }
    }
    fmt.Println(v12, v12, v10)
    v18 := v3[32768]
    switch -196 {
    case -660:
        var v19 string = v18
        _ = v19
    case 128:
        var v20 float64 = 3.14159
        _ = v20
    default:
        v21 := v3[32768]
        _ = v21
    }
    var v22 float32 = -2.25
    var v23 string = (("hello" + v10) + "")
    v18 = "a"
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS2{F0: 1, F1: 42}
        _ = fi24
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
