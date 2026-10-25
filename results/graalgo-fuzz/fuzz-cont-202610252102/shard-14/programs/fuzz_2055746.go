package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 bool }
type GenS2 struct { F0 bool; F1 uint }
func genFn0(p0 int64, p1 uint16) int {
    if false {
        v7 := GenS1{F0: -536, F1: -232, F2: true}
        _ = v7
    } else {
        if ((p1 > p1) || (-1 >= 42)) {
            var v8 int8 = 100
            _ = v8
        } else {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := []int{325, fi9, (fi9 - 292)}
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v11 float64 = 0.1
                _ = fi9
                _ = v10
                _ = v11
            }
        }
    }
    var v12 float64 = 1e10
    _ = p0
    _ = p1
    _ = v12
    return -1
}

func main() {
    var v1 uint = 762
    v2 := make(map[uint32]string)
    v3 := make(map[uint32]int)
    var v4 bool = true
    var v5 int8 = 127
    var v6 int64 = -228
    switch genFn0(v6, 849) {
    case -2147483648:
        switch -2147483648 {
        case 42:
            v6 += v6
        case -2147483648:
            var v13 uint = 18446744073709551615
            _ = v13
        case 4294967295:
            v14 := GenS1{F0: 1000, F1: 2147483647, F2: (0 <= 0)}
            _ = v14
        default:
            var v15 float64 = float64(0.1)
            _ = v15
        }
    case 9223372036854775807:
        v16 := GenS2{F0: true, F1: (v1 + v1)}
        _ = v16
    }
    var v17 bool = v4
    var v18 uint16 = 0
    var v19 uint32 = 2147483647
    if (91 >= int32(v6)) {
        fmt.Println(v6)
        var v20 string = "hello"
        var v21 int8 = 127
        _ = v20
        _ = v21
    } else {
        var v22 int16 = -32768
        _ = v22
    }
    var v23 uint64 = 439
    var v24 string = (("\n\t" + "foo bar") + (("hello" + "\n\t") + "foo bar"))
    switch v24 {
    case "a":
        v4 = (!v17)
    }
    if true {
        var v25 int16 = 42
        _ = v25
    } else {
        var v26 float32 = 1e10
        _ = v26
    }
    if (32767 <= 1) {
        fmt.Println(v23, v4)
        var v27 string = (v24 + v24)
        for fi28 := 0; fi28 < 19; fi28++ {
            var v29 int8 = 1
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
