package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 uint8, p1 uint32) int16 {
    var v5 uint8 = (p0 ^ 1)
    v5 -= (^(v5))
    v6 := []int{0, 32767}
    v5 = v5
    if (-1 == 0) {
        var v7 string = "foo bar"
        if len(v6) > 0 {
            v6[0] = -51
        }
        _ = v7
    } else {
        var v8 bool = (false || true)
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 int = 127
    var v2 uint64 = 0
    var v3 int8 = -128
    var v4 float64 = -2.25
    if ((2 == 42) || false) {
        v4 = v4
        var v9 float64 = v4
        _ = v9
    } else {
        fmt.Println(v4, v3)
        v3 += -42
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint8 = 128
        v11 = (^(v11))
        _ = fi10
        _ = v11
    }
    var v12 int64 = -880
    fmt.Println(v1, v12, v2)
    fmt.Println(v2, v1, v4)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 int16 = 73
        var v15 int32 = -2147483648
        var v16 int16 = (v14 | (v14 | 72))
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 string = ""
    var v18 int16 = 1
    if false {
        v19 := GenS1{F0: 198}
        fmt.Println(v12, v4)
        var v20 int32 = 0
        _ = v19
        _ = v20
    }
    v12 = (v12 & v12)
    var v21 int16 = v18
    v17 = ""
    var v22 int = v1
    var v23 uint16 = 32767
    fmt.Println(v1)
    var v24 uint8 = 128
    if false {
        v25 := GenS1{F0: (v24 / (((^(v24))) | 1))}
        _ = v25
    }
    switch v17 {
    case "a":
        fmt.Println(v1, v3, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
