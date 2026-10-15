package main

import (
    "fmt"
)

func genFn0() int64 {
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float32 = 3.14159
        v8 -= (v8 + 3.14159)
        if false {
            var v9 float64 = -2.25
            var v10 float32 = v8
            _ = v10
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    switch -336 {
    case 32767:
        v11 := []int{-1, 127, -997}
        _ = v11
    }
    return 1
}
func genFn1() uint32 {
    v12 := []int{4294967295, 128, -9223372036854775808, 252}
    var v13 int64 = -199
    var v14 int8 = 1
    fmt.Println(v13)
    v13 -= (^(v13))
    _ = v12
    _ = v13
    _ = v14
    return 1
}

func main() {
    v1 := []uint32{1, 4294967295, 42, 42}
    v2 := make(map[int]int)
    var v3 float32 = -2.25
    var v4 float32 = float32((0.1 + 1.5))
    var v5 uint32 = 1
    var v6 int16 = 1
    var v15 int = -338
    var v16 uint16 = 1
    var v17 float64 = float64(v3)
    var v18 bool = (1 >= -99)
    fmt.Println(v16, v3)
    var v19 uint64 = 2
    fmt.Println(v15, v16)
    v16 -= 2
    var v20 uint = 0
    if len(v1) > 0 {
        v1[0] = (^(v5))
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v16 += v16
        var v22 int8 = 42
        _ = fi21
        _ = v22
    }
    var v23 int8 = 0
    var v24 string = ("hello" + ("" + "foo bar"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
