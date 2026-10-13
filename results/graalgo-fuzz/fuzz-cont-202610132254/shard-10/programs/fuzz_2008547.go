package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int) uint {
    p0 -= (p0 << ((42) & 63))
    var v5 float32 = (1.5 - (0.1 + float32(0.0)))
    var v6 float64 = (1.5 * (1e10 * 0.1))
    p0 += 9223372036854775807
    _ = p0
    _ = v5
    _ = v6
    return 39
}
func genFn1() uint {
    v7 := GenS1{F0: "a"}
    var v8 string = (("" + ("\n\t" + "foo bar")) + "foo bar")
    for fi9 := 0; fi9 < 15; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            switch fi9 {
            case 127:
                v8 = ""
            }
            v8 = ("foo bar" + v8)
            _ = fi10
        }
        v11 := []int{fi9, 2147483647}
        _ = fi9
        _ = v11
    }
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint8 = 0
    var v2 uint64 = 4294967295
    var v3 float32 = 0.0
    var v4 uint8 = 2
    v12 := []int{}
    for fi13 := 0; fi13 < 1; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            v1 += (v1 % ((uint8(v2)) | 1))
            _ = fi14
        }
        fmt.Println(v4)
        var v15 int = (fi13 << ((425) & 63))
        _ = fi13
        _ = v15
    }
    var v16 uint8 = 128
    v1 += v1
    v12 = append(v12, 9223372036854775807)
    var v17 float64 = (1.5 / 1e10)
    fmt.Println(v1, v4)
    v2 = v2
    v18 := GenS1{F0: "hello"}
    var v19 int16 = -424
    if len(v12) > 0 {
        v12[0] = -853
    }
    fmt.Println(v4, v19, v2)
    if ((v19 | (v19 / ((0) | 1))) <= (v19 >> ((0) & 63))) {
        _ = len(v12)
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
