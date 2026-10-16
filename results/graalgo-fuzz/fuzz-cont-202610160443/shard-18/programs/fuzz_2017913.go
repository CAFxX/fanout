package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
func genFn0() int8 {
    var v7 bool = false
    switch -2147483648 {
    case -38:
        v7 = v7
    }
    var v8 int8 = 87
    if false {
        v9 := []int{1, 2147483647, 0, -2147483648}
        _ = v9
    } else {
        v7 = ((0.0 / (-2.25 + 0.0)) <= 0.0)
        if v7 {
            var v10 bool = (606 >= 102)
            var v11 float64 = 1.5
            _ = v10
            _ = v11
        } else {
            v8 += -78
            v7 = (1e10 >= (0.1 + (0.0 - float32(float64((3.14159 - 0.1))))))
        }
    }
    v12 := GenS1{F0: v7, F1: (31 <= 2)}
    _ = v12
    _ = v7
    _ = v8
    return (v8 % (((v8 * 127)) | 1))
}

func main() {
    v1 := []uint32{655, 2147483648, 4294967295, 65535}
    v2 := make(map[string]int)
    var v3 uint64 = 586
    var v4 float64 = -2.25
    v5 := []string{"", ("hello" + (("a" + ("a" + "")) + ("" + "\n\t")))}
    v6 := make(map[string]uint32)
    if false {
        for fi13 := 0; fi13 < 10; fi13++ {
            v3 += (v3 ^ v3)
            _ = fi13
        }
        v3 += (v3 >> ((65535) & 63))
        fmt.Println(v3)
    }
    v4 = ((v4 - 3.14159) - 3.14159)
    fmt.Println(v3, v3)
    var v14 uint8 = 42
    for fi15 := 0; fi15 < 8; fi15++ {
        var v16 int64 = 0
        _ = fi15
        _ = v16
    }
    fmt.Println(v14)
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if len(v1) > 0 {
        v1[0] = 249
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
