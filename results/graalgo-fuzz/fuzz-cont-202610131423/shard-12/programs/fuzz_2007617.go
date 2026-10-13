package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int) int64 {
    var v7 uint8 = 0
    v8 := []int{(p0 ^ (p0 / ((p0) | 1))), (^(p0)), 1, (p0 >> ((255) & 63))}
    p0 = 128
    _ = p0
    _ = v7
    _ = v8
    return 42
}
func genFn1() uint32 {
    var v9 uint8 = 42
    v9 += (v9 / ((v9) | 1))
    var v10 float32 = 0.1
    fmt.Println(v10)
    _ = v10
    _ = v9
    return 355
}

func main() {
    v1 := []string{"foo bar"}
    var v2 uint64 = 2
    var v3 uint16 = 65535
    var v4 int64 = 42
    var v5 int32 = -1
    var v6 int8 = 83
    v11 := GenS1{F0: "hello"}
    _ = len(v1)
    v5 = v5
    v6 -= v6
    for fi12 := 0; fi12 < 6; fi12++ {
        var v13 bool = (((("" + "\n\t") + "hello") + "a") > ("a" + "a"))
        v14 := GenS1{F0: "foo bar"}
        _ = fi12
        _ = v13
        _ = v14
    }
    v2 += uint64(v3)
    if len(v1) > 0 {
        v1[0] = "a"
    }
    v2 = (^(v2))
    var v15 int64 = 518
    var v16 bool = (224 == 65535)
    v17 := GenS1{F0: "a"}
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 uint32 = 0
        _ = fi18
        _ = v19
    }
    var v20 int = 32767
    v6 += 0
    var v21 int16 = int16(v5)
    var v22 bool = v16
    var v23 float32 = (0.0 + 3.14159)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
