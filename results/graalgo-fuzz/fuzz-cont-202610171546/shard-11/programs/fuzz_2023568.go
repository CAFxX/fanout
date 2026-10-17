package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() int {
    switch "" {
    case "":
        v4 := GenS1{F0: -27}
        _ = v4
    }
    for fi5 := 0; fi5 < 4; fi5++ {
        switch "" {
        case "hello":
            var v6 uint64 = 4294967295
            _ = v6
        case "":
            var v7 int16 = 127
            _ = v7
        }
        var v8 uint8 = 2
        _ = fi5
        _ = v8
    }
    var v9 int = 42
    var v10 float32 = -2.25
    _ = v10
    _ = v9
    return v9
}

func main() {
    var v1 uint = 298
    var v2 int16 = 1000
    var v3 uint32 = 555
    v2 = 766
    fmt.Println(v3)
    v3 = (v3 / (((v3 / (((v3 % ((v3) | 1))) | 1))) | 1))
    switch -2147483648 {
    case 13:
        var v11 int64 = 9223372036854775807
        _ = v11
    default:
        var v12 int = -2147483648
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        fmt.Println(v1)
        _ = fi13
    }
    fmt.Println(v3)
    var v14 string = "hello"
    v15 := []int{9223372036854775807, genFn0()}
    if len(v15) > 0 {
        v15[0] = -2147483648
    }
    v15 = append(v15, -239)
    if len(v15) > 0 {
        v15[0] = 32768
    }
    fmt.Println(v14, v3, v1)
    for fi16 := 0; fi16 < 8; fi16++ {
        var v17 int64 = -1
        for fi18 := 0; fi18 < 16; fi18++ {
            v1 = (v1 - 640)
            v17 += v17
            fmt.Println(v1, v1, v3)
            _ = fi18
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v1, v2, v2)
    var v19 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
}
