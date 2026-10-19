package main

import (
    "fmt"
)

func genFn0(p0 uint16) uint {
    switch "hello" {
    case "foo bar":
        p0 -= p0
    default:
        p0 = 65535
    }
    var v7 int = -80
    v7 += v7
    _ = p0
    _ = v7
    return 42
}
func genFn1(p0 int) int32 {
    var v8 string = ("foo bar" + "")
    switch (v8 + "") {
    case "a":
        var v9 int64 = -411
        _ = v9
    case "":
        var v10 uint16 = 42
        _ = v10
    case "foo bar":
        p0 = (^(p0))
    }
    var v11 uint = 4294967295
    _ = p0
    _ = v11
    _ = v8
    return 894
}

func main() {
    var v1 uint16 = 2
    var v2 int32 = 42
    var v3 uint64 = 1
    v4 := []int{128, 32767}
    v5 := make(map[uint32]bool)
    v6 := make(map[uint32]string)
    var v12 float64 = ((3.14159 / 1e10) - (float64(0.0) * 0.1))
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            var v15 int64 = 233
            fmt.Println(v15, fi14)
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    for fi16 := 0; fi16 < 16; fi16++ {
        var v17 uint32 = 2147483648
        var v18 uint16 = v1
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            if len(v4) > 0 {
                v4[0] = fi20
            }
            _ = fi20
        }
        _ = fi19
    }
    _ = len(v4)
    for fi21 := 0; fi21 < 2; fi21++ {
        v3 += (v3 % ((v3) | 1))
        _ = fi21
    }
    fmt.Println(v12)
    switch -2147483648 {
    case 128:
        var v22 float64 = (v12 - v12)
        _ = v22
    }
    var v23 uint64 = v3
    delete(v6, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
