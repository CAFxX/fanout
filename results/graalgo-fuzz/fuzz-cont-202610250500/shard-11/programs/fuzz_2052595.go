package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16) uint32 {
    var v7 string = ("hello" + ("foo bar" + ("foo bar" + "a")))
    var v8 int = -777
    fmt.Println(v8, v7, p0)
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{-9223372036854775808, 128, 2147483647, 209}
    v2 := make(map[int]string)
    var v3 int8 = -128
    v4 := make(map[int]bool)
    v5 := make(map[uint32]int)
    var v6 float32 = ((0.1 * 0.0) + 1e10)
    switch 9223372036854775807 {
    case -678:
        if true {
            var v9 int32 = 575
            _ = v9
        }
    case 127:
        var v10 int16 = 1000
        _ = v10
    case 4294967295:
        var v11 int = -9223372036854775808
        _ = v11
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    switch 42 {
    case 42:
        if (42 > 0) {
            var v12 uint32 = 2147483648
            _ = v12
        }
    case 32767:
        if false {
            v13 := GenS1{F0: 255}
            _ = v13
        } else {
            v14 := GenS1{F0: 42}
            var v15 int8 = (v3 + (v3 ^ v3))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v14
            _ = v15
        }
    case -2147483648:
        for fi16 := 0; fi16 < 1; fi16++ {
            v3 = 42
            if true {
                v3 = 100
            }
            v6 = 3.14159
            _ = fi16
        }
    }
    if true {
        fmt.Println(v3, v6)
        var v17 float64 = (float64((v6 - (v6 + (3.14159 * ((v6 / -2.25) + (1.5 / v6)))))) / (1.5 * (((0.0 + -2.25) - 1e10) + (float64(float32((float64(v6) - -2.25))) + float64(((v6 / v6) + v6))))))
        fmt.Println(v6)
        _ = v17
    }
    var v18 bool = ((!(3.14159 != 0.1)) && false)
    _ = len(v1)
    v18 = (("hello" >= "\n\t") && v18)
    fmt.Println(v6)
    v19 := GenS1{F0: 728}
    var v20 uint32 = 479
    v6 = -2.25
    if true {
        var v21 string = (("a" + ("a" + "foo bar")) + "a")
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
