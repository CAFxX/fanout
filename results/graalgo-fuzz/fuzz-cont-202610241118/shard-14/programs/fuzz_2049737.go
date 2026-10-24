package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
func genFn0(p0 int32, p1 int32) int64 {
    var v6 uint32 = 626
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int32 = p0
        _ = fi7
        _ = v8
    }
    v9 := []int{0}
    var v10 float64 = -2.25
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v6
    _ = v9
    return -529
}

func main() {
    var v1 int32 = 42
    var v2 uint8 = 255
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 uint = 432
    var v11 uint32 = 590
    var v12 uint8 = 128
    v2 = v12
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 1, F1: -67}
        _ = fi13
        _ = v14
    }
    var v15 float64 = float64(1e10)
    fmt.Println(v11, v2, v5)
    fmt.Println(v2, v2, v5)
    for fi16 := 0; fi16 < 18; fi16++ {
        for fi17 := 0; fi17 < 14; fi17++ {
            if false {
                if true {
                    var v18 int64 = 9223372036854775807
                    _ = v18
                } else {
                    if (42 <= int8(v1)) {
                        v1 = (v1 - v1)
                    } else {
                        fmt.Println(fi17)
                        var v19 uint16 = 255
                        _ = v19
                    }
                    fmt.Println(v15, v11)
                    v1 = 1
                    var v20 int8 = 127
                    _ = v20
                }
                var v21 int8 = -128
                var v22 uint = v5
                v4["\n\t"] = true
                _ = v21
                _ = v22
            } else {
                v5 += v5
            }
            v11 = v11
            var v23 bool = ((false && false) && (false && false))
            _ = fi17
            _ = v23
        }
        var v24 int8 = -123
        for fi25 := 0; fi25 < 13; fi25++ {
            v24 += (v24 * -1)
            _ = fi25
        }
        _ = fi16
        _ = v24
    }
    v26 := []int{4294967295, -691, 127}
    v3[("hello" + "hello")] = true
    var v27 uint16 = 255
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := GenS1{F0: 42, F1: 42}
            var v31 uint16 = 9
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    var v32 string = (("hello" + ("foo bar" + "hello")) + "foo bar")
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
