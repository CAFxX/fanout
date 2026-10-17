package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 int32; F2 bool }
func genFn0(p0 int8) uint16 {
    v7 := []int{128, 0, -102}
    p0 += p0
    _ = p0
    _ = v7
    return 807
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (-128 >= -1)
    var v3 string = "a"
    var v4 uint8 = 127
    var v5 string = v3
    var v6 int32 = int32(v4)
    v8 := GenS1{F0: 42, F1: -1}
    var v9 int16 = 1000
    if (false || ((1.5 * 0.0) < 1e10)) {
        delete(v1, 1)
        v2 = (!(!(((v6 & 42) >= (v6 % (((v6 * 101)) | 1))) || true)))
    }
    v3 = v3
    v10 := []int{128}
    switch 128 {
    case 32768:
        v4 = v4
    case 128:
        v11 := GenS2{F0: 42, F1: v6, F2: v2}
        _ = v11
    case 4294967295:
        v12 := GenS1{F0: (v4 >> ((1) & 63)), F1: 9223372036854775807}
        _ = v12
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 int32 = 912
        v2 = v2
        _ = fi13
        _ = v14
    }
    var v15 bool = (42 != 28)
    if (970 > 2) {
        v5 = "foo bar"
    }
    var v16 float32 = (3.14159 / 1e10)
    v10 = append(v10, 0)
    var v17 int = 32768
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 int = fi18
        var v20 int32 = v6
        if false {
            v15 = true
        } else {
            var v21 uint32 = 65535
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    var v22 int = -774
    switch (v5 + v5) {
    case "a":
        if ((v9 % ((42) | 1)) == -1) {
            delete(v1, -2147483648)
        }
    case "":
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 int64 = 42
            _ = fi23
            _ = v24
        }
    }
    fmt.Println(v16, v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
