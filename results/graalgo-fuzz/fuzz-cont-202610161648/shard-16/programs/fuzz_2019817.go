package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
func genFn0() uint32 {
    if (true || (!(86 <= 100))) {
        var v7 uint8 = 1
        _ = v7
    } else {
        if false {
            var v8 int16 = 381
            var v9 uint16 = 0
            _ = v8
            _ = v9
        } else {
            var v10 bool = ((float64((1e10 / -2.25)) + 3.14159) < 0.0)
            fmt.Println(v10)
            var v11 string = "hello"
            _ = v10
            _ = v11
        }
    }
    var v12 uint8 = 191
    v13 := []int{0}
    _ = v12
    _ = v13
    return 4294967295
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    var v5 uint32 = 631
    var v6 int64 = 589
    v3[341] = (!(1 >= 411))
    var v14 uint16 = 1
    var v15 bool = true
    var v16 bool = v15
    for fi17 := 0; fi17 < 13; fi17++ {
        v18 := []int{int(v2)}
        if len(v18) >= 2 {
            v18 = v18[1:len(v18)]
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v15, v16)
    if false {
        var v19 uint8 = 128
        _ = v19
    } else {
        if v15 {
            var v20 int8 = 1
            _ = v20
        } else {
            v21 := []int{}
            fmt.Println(v16)
            v16 = ((v14 != 987) && false)
            _ = v21
        }
        if false {
            _ = len(v4)
        } else {
            var v22 uint32 = v5
            for fi23 := 0; fi23 < 11; fi23++ {
                if (v2 != v2) {
                    v24 := []int{}
                    _ = len(v24)
                    v25 := GenS1{F0: 120, F1: genFn0()}
                    _ = v24
                    _ = v25
                }
                for fi26 := 0; fi26 < 18; fi26++ {
                    var v27 int64 = -9223372036854775808
                    _ = fi26
                    _ = v27
                }
                _ = fi23
            }
            _ = v22
        }
    }
    v5 = (v5 ^ genFn0())
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
