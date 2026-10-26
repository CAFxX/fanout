package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 int32 }
func genFn0() string {
    v6 := []int{}
    if (false && (32768 >= 848)) {
        v7 := GenS1{F0: -128, F1: 0, F2: 1}
        v8 := GenS2{F0: 0, F1: 408, F2: 2147483647}
        _ = v7
        _ = v8
    } else {
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v9 int = 924
    _ = v6
    _ = v9
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 int = -1
    var v2 uint8 = 48
    v3 := []int{v1, 42, v1}
    var v4 int64 = -413
    var v5 int16 = 32767
    v10 := GenS2{F0: 2, F1: v4, F2: 253}
    var v11 int64 = (^(v4))
    var v12 string = "a"
    v13 := GenS1{F0: 47, F1: -1, F2: 85}
    var v14 int8 = 127
    if (-451 > v1) {
        v15 := GenS2{F0: 2147483648, F1: 55, F2: 0}
        _ = v15
    } else {
        if len(v3) > 0 {
            v3[0] = v1
        }
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint64 = 42
        _ = fi16
        _ = v17
    }
    var v18 uint16 = 1
    var v19 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
