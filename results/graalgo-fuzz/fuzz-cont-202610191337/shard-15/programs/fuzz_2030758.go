package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 int, p1 uint) uint64 {
    for fi6 := 0; fi6 < 11; fi6++ {
        switch -1 {
        case 32767:
            p1 = p1
        case 128:
            v7 := GenS1{F0: 65535, F1: -355}
            _ = v7
        case 0:
            var v8 float64 = 3.14159
            _ = v8
        default:
            if (0 < 2147483647) {
                if (0.1 > (1.5 + (1e10 / 0.1))) {
                    var v9 float32 = (0.1 - 0.1)
                    var v10 int64 = -681
                    var v11 int64 = (v10 + v10)
                    _ = v10
                    _ = v11
                    _ = v9
                }
                fmt.Println(fi6, fi6, p1)
                fmt.Println(p0, p1, p1)
            }
        }
        fmt.Println(p1, p0)
        _ = fi6
    }
    var v12 float64 = float64(3.14159)
    var v13 int16 = 42
    fmt.Println(v13, p0)
    for fi14 := 0; fi14 < 17; fi14++ {
        v12 = -2.25
        fmt.Println(p0, v12)
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 538
}

func main() {
    var v1 bool = true
    var v2 int = 32768
    var v3 string = ((("foo bar" + (("a" + ((("hello" + "foo bar") + "foo bar") + ("\n\t" + "\n\t"))) + ("foo bar" + (("\n\t" + ("" + "\n\t")) + "")))) + "hello") + "\n\t")
    var v4 uint8 = 255
    var v5 uint64 = 2
    v15 := GenS2{F0: v4, F1: 127}
    v1 = true
    if true {
        v16 := GenS1{F0: 65535, F1: (^(v2))}
        _ = v16
    } else {
        v17 := GenS2{F0: uint8(v5), F1: int8(v4)}
        _ = v17
    }
    fmt.Println(v3, v5)
    v18 := []int{2147483647, 32767}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v4 = (^(v4))
    _ = len(v18)
    v5 += genFn0(-1, 42)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
