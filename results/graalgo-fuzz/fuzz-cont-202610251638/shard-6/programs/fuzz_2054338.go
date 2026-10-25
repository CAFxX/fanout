package main

import (
    "fmt"
)

func genFn0(p0 int, p1 uint64) uint32 {
    v7 := []int{}
    var v8 int = p0
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 int64 = 9223372036854775807
        var v11 uint = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 float64 = (-2.25 - -2.25)
    var v2 uint16 = 496
    v3 := []int{-519, 4294967295, 32768}
    var v4 int8 = 25
    var v5 int = 262
    var v6 uint32 = 65535
    fmt.Println(v1, v4, v2)
    v5 = (v5 % ((428) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        var v12 int32 = -2147483648
        fmt.Println(v2, v5)
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int = v5
        var v15 uint8 = 127
        fmt.Println(fi13, v15)
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 bool = (150 == 127)
    fmt.Println(v2, v6)
    for fi17 := 0; fi17 < 1; fi17++ {
        v5 -= 2147483647
        var v18 bool = true
        _ = fi17
        _ = v18
    }
    fmt.Println(v4, v1, v16)
    for fi19 := 0; fi19 < 13; fi19++ {
        if false {
            v1 = v1
            if len(v3) > 0 {
                v3[0] = fi19
            }
        } else {
            fmt.Println(v1, fi19)
        }
        var v20 int8 = 0
        _ = fi19
        _ = v20
    }
    if (0 >= 42) {
        var v21 uint8 = 2
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
