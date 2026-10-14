package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 int8 }
func genFn0() string {
    var v6 int = -412
    if (true && false) {
        if (!(!(1 > 2147483648))) {
            fmt.Println(v6)
        } else {
            v7 := []int{v6, v6, v6}
            v8 := GenS1{F0: 65535, F1: 127, F2: "foo bar"}
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 uint = 644
            v10 = v10
            v11 := GenS2{F0: v6, F1: 0}
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int = 0
        _ = v12
    }
    _ = v6
    return "foo bar"
}

func main() {
    var v1 uint64 = 42
    var v2 uint = 0
    var v3 bool = ((42 != 255) && false)
    v4 := make(map[string]string)
    var v5 int32 = 2147483647
    var v13 uint16 = uint16(v2)
    v5 = (^(v5))
    v14 := GenS2{F0: -734, F1: -60}
    fmt.Println(v2, v13)
    _ = len(v4)
    if (float64(0.0) == 1e10) {
        var v15 int32 = 807
        _ = v15
    }
    var v16 uint64 = v1
    delete(v4, "hello")
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 uint16 = 1
        _ = fi17
        _ = v18
    }
    if (1e10 != -2.25) {
        var v19 int16 = 1000
        var v20 uint32 = 2
        _ = v19
        _ = v20
    }
    var v21 int16 = -32768
    var v22 float64 = 1.5
    if v3 {
        v3 = (int16(v5) != 1000)
    }
    v23 := GenS2{F0: 9223372036854775807, F1: 100}
    var v24 float64 = v22
    v22 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
