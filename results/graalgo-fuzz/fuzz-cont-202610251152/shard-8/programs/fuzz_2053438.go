package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8) uint {
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            v8 := []int{}
            p0 += (^(p0))
            _ = fi7
            _ = v8
        }
        v9 := []int{4294967295}
        v10 := GenS1{F0: 4294967295, F1: 1}
        _ = fi6
        _ = v10
        _ = v9
    }
    if false {
        switch (("\n\t" + ("a" + "\n\t")) + "\n\t") {
        case "\n\t":
            v11 := []int{-9223372036854775808}
            _ = v11
        }
        v12 := []int{-948, -432, 0, 9223372036854775807}
        if len(v12) > 0 {
            v12[0] = 42
        }
        _ = v12
    }
    _ = p0
    return 147
}

func main() {
    v1 := []uint32{40, 245, 65535, 0}
    var v2 float32 = 1.5
    v3 := make(map[uint32]uint32)
    var v4 int8 = -128
    var v5 int16 = 1000
    var v13 int32 = -2147483648
    v14 := GenS2{F0: v5}
    fmt.Println(v4)
    for fi15 := 0; fi15 < 11; fi15++ {
        fmt.Println(v5)
        v16 := GenS3{F0: 1}
        _ = fi15
        _ = v16
    }
    fmt.Println(v13, v5)
    var v17 string = ("a" + ("foo bar" + (("a" + "foo bar") + (("\n\t" + (("\n\t" + "hello") + "hello")) + "foo bar"))))
    var v18 uint32 = 65535
    v13 = (^(v13))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
