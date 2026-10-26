package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint }
func genFn0() uint64 {
    var v4 int8 = 75
    var v5 int32 = 0
    _ = v4
    _ = v5
    return 4294967295
}

func main() {
    var v1 int16 = 1
    v2 := []string{"a", "hello", ""}
    v3 := []uint32{4294967295, 237}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 += -32768
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            _ = len(v3)
        } else {
            v1 = v1
            v7 := GenS1{F0: ("foo bar" + ("a" + ("\n\t" + "foo bar")))}
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v1)
    v1 = v1
    var v8 int8 = 127
    var v9 uint16 = 255
    v10 := GenS2{F0: (v1 * v1), F1: 4294967295}
    if len(v3) > 0 {
        v3[0] = 2147483648
    }
    fmt.Println(v1, v1)
    v8 = v8
    for fi11 := 0; fi11 < 7; fi11++ {
        v8 = -128
        var v12 uint8 = 221
        v13 := GenS1{F0: ""}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 uint16 = v9
    for fi15 := 0; fi15 < 9; fi15++ {
        var v16 uint = 255
        if len(v3) > 0 {
            v3[0] = 65535
        }
        v17 := GenS2{F0: (v1 >> (((^(v16))) & 63)), F1: v16}
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 float64 = 0.1
    var v19 uint = 876
    var v20 uint = (v19 << (((v19 - (v19 % ((0) | 1)))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v8
    _ = v9
}
