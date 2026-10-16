package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 string }
func genFn0() uint16 {
    var v6 uint8 = 255
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 float64 = 3.14159
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{fi7, fi7, (fi7 & (-(fi7)))}
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 65535
}

func main() {
    var v1 int8 = 42
    var v2 int64 = 42
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 uint = 65535
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint64 = 4294967295
        var v13 string = "hello"
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 string = "hello"
    v2 += -1
    v15 := GenS1{F0: 65535}
    v14 = ("a" + "hello")
    v5 -= v5
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 += -9223372036854775808
        if (3.14159 >= -2.25) {
            v1 = 100
            var v17 int32 = -2147483648
            _ = v17
        }
        _ = fi16
    }
    v14 = ((((v14 + "foo bar") + v14) + (v14 + "hello")) + ((v14 + v14) + (v14 + v14)))
    fmt.Println(v14, v14)
    var v18 uint8 = 208
    var v19 uint16 = 527
    v19 = 825
    var v20 string = ("\n\t" + "hello")
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint16 = (v19 | 42)
        _ = len(v4)
        _ = fi21
        _ = v22
    }
    v23 := GenS2{F0: 0, F1: v20}
    var v24 uint64 = 1
    v4[-26] = v20
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
