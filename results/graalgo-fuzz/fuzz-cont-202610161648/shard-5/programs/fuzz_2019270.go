package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
func genFn0(p0 int16, p1 uint32) string {
    v4 := []int{-2147483648}
    v5 := GenS1{F0: p1, F1: 1, F2: -521}
    p0 += -87
    switch 32767 {
    case 128:
        for fi6 := 0; fi6 < 12; fi6++ {
            v7 := GenS1{F0: (^(p1)), F1: 505, F2: 42}
            switch (fi6 & (fi6 / ((32768) | 1))) {
            case 477:
                var v8 int32 = -1
                _ = v8
            case -1:
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 uint64 = 18446744073709551615
                    var v11 int32 = 42
                    _ = fi9
                    _ = v10
                    _ = v11
                }
            case 42:
                for fi12 := 0; fi12 < 18; fi12++ {
                    p1 = 2147483647
                    _ = fi12
                }
            }
            _ = fi6
            _ = v7
        }
    case -2147483648:
        var v13 string = ("\n\t" + ((("hello" + "hello") + ("hello" + ("" + ("\n\t" + ("\n\t" + ("foo bar" + "a")))))) + "hello"))
        _ = v13
    default:
        var v14 uint64 = 18
        _ = v14
    }
    var v15 bool = (-662 >= 0)
    _ = p0
    _ = p1
    _ = v15
    _ = v4
    _ = v5
    return (("" + (("\n\t" + ("foo bar" + "")) + "hello")) + "a")
}

func main() {
    v1 := []int{}
    var v2 int = -2147483648
    var v3 int16 = 32767
    v1 = append(v1, 283)
    fmt.Println(v2)
    v3 += -672
    if false {
        var v16 uint32 = 749
        v3 -= -341
        var v17 string = "\n\t"
        _ = v16
        _ = v17
    }
    var v18 float32 = -2.25
    var v19 int16 = -32768
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 bool = (true || false)
        _ = fi20
        _ = v21
    }
    var v22 int = (v2 << ((42) & 63))
    for fi23 := 0; fi23 < 3; fi23++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi23
    }
    var v24 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v24
    _ = v3
}
