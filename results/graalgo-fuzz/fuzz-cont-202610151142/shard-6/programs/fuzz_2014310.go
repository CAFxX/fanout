package main

import (
    "fmt"
)

func genFn0(p0 string) uint {
    fmt.Println(p0)
    p0 = p0
    v7 := []int{128, 4294967295, -2147483648, 32768}
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    v1 := []string{}
    v2 := make(map[int]uint32)
    var v3 float64 = (((1.5 - -2.25) + -2.25) + 1e10)
    var v4 uint16 = 310
    var v5 uint32 = 0
    var v6 float64 = v3
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v3, v5, v5)
        v4 = v4
        v3 = (((3.14159 / v6) * v6) / float64(0.1))
        _ = fi8
    }
    var v9 uint = 65535
    for fi10 := 0; fi10 < 8; fi10++ {
        var v11 bool = true
        var v12 int = 86
        _ = fi10
        _ = v11
        _ = v12
    }
    v13 := v2[128]
    _ = len(v2)
    var v14 int16 = 1
    var v15 int8 = 0
    v13 = v13
    for fi16 := 0; fi16 < 17; fi16++ {
        if false {
            if len(v1) > 0 {
                v1[0] = (("foo bar" + "foo bar") + "a")
            }
            v14 = (v14 ^ 362)
            fmt.Println(v6)
        }
        _ = fi16
    }
    var v17 int32 = -2147483648
    v18 := v2[-2147483648]
    _ = len(v2)
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v19 uint16 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
