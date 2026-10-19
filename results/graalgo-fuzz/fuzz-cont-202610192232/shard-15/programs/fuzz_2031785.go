package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0() uint {
    v6 := []int{}
    var v7 int16 = 1
    v7 += (v7 << ((255) & 63))
    var v8 int = -562
    _ = v6
    _ = v7
    _ = v8
    return 802
}

func main() {
    var v1 int = 4294967295
    var v2 bool = (!(v1 >= (-(v1))))
    var v3 int8 = -128
    v4 := make(map[string]bool)
    var v5 float64 = (1e10 + (3.14159 - 1e10))
    v2 = (42 >= 1)
    switch (v1 % (((v1 << ((42) & 63))) | 1)) {
    case 32767:
        v9 := []int{128, 9223372036854775807, (v1 + 4294967295)}
        _ = v9
    case -948:
        if ((((0.0 - 1.5) / -2.25) / (-2.25 * 1.5)) >= float32((v5 + v5))) {
            var v10 int8 = v3
            _ = v10
        }
    case -689:
        var v11 int16 = 0
        _ = v11
    }
    v12 := []int{(^(v1))}
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int = 1
        _ = fi13
        _ = v14
    }
    var v15 int64 = 790
    v15 = v15
    v1 += v1
    var v16 uint8 = 42
    for fi17 := 0; fi17 < 8; fi17++ {
        v18 := GenS1{F0: v1}
        _ = fi17
        _ = v18
    }
    var v19 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
