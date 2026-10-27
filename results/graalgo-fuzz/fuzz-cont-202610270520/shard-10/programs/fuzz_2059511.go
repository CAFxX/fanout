package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
func genFn0(p0 uint, p1 int32, p2 int16) int8 {
    p0 = p0
    switch 653 {
    case 0:
        p1 += p1
    case -2147483648:
        p2 = p2
    case 32767:
        fmt.Println(p0, p0)
    default:
        fmt.Println(p0, p2, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 100
}
func genFn1() uint64 {
    var v6 int32 = 23
    v7 := []int{664, -751}
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    v1 := []uint32{454, 2147483648}
    v2 := []uint32{1, 42}
    v3 := make(map[string]string)
    var v4 int8 = -128
    v5 := []string{}
    for fi8 := 0; fi8 < 2; fi8++ {
        var v9 int64 = -1
        var v10 int32 = 2147483647
        v10 -= (v10 ^ 1)
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 int = -816
    if (!true) {
        v3["a"] = "\n\t"
        v12 := GenS1{F0: 112, F1: 830}
        v11 -= (^(v11))
        _ = v12
    } else {
        v11 = v11
    }
    v13 := v3["\n\t"]
    v3[(v13 + "hello")] = (((v13 + v13) + v13) + v13)
    v4 += -7
    v11 = (v11 | (v11 << ((0) & 63)))
    var v14 string = (((v13 + "a") + "") + (v13 + ""))
    if (2147483647 < 2) {
        var v15 int16 = 1000
        _ = v15
    } else {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 uint32 = 2
            v14 = v14
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: 2, F1: 921}
        v13 = ("\n\t" + v13)
        _ = v18
    }
    fmt.Println(v11, v14, v13)
    for fi19 := 0; fi19 < 5; fi19++ {
        v13 = "\n\t"
        var v20 int = (fi19 * v11)
        v13 = ("\n\t" + "a")
        _ = fi19
        _ = v20
    }
    var v21 uint32 = 65535
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
