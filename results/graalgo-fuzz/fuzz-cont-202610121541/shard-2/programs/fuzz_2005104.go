package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }
func genFn0() uint64 {
    if (0 == 127) {
        var v4 bool = (0 == 18446744073709551615)
        var v5 float32 = 0.0
        _ = v4
        _ = v5
    } else {
        var v6 uint64 = 584
        if (1e10 == 0.0) {
            var v7 int64 = 0
            _ = v7
        }
        _ = v6
    }
    if (1e10 < 1.5) {
        if (380 < 42) {
            switch -1 {
            case 42:
                var v8 int32 = 1
                _ = v8
            case -155:
                var v9 float64 = 1e10
                _ = v9
            }
        } else {
            var v10 uint8 = 127
            v11 := []int{42, -1}
            _ = v10
            _ = v11
        }
        v12 := []int{128, 128, -9223372036854775808}
        var v13 int8 = 127
        _ = v12
        _ = v13
    }
    var v14 int32 = -1
    _ = v14
    return 2
}

func main() {
    var v1 int16 = 1
    var v2 string = ""
    var v3 int32 = 414
    v3 = 1
    v2 = "a"
    var v15 uint = 0
    var v16 uint = (v15 >> ((255) & 63))
    if false {
        switch 127 {
        case 32768:
            v17 := []int{-980, -9223372036854775808, -9223372036854775808}
            _ = v17
        default:
            switch 42 {
            case 996:
                v18 := []int{-779, 957, 127, -9223372036854775808}
                _ = v18
            case 32768:
                v1 += 531
            }
        }
        if false {
            v19 := []int{-142, -441}
            v15 = v16
            v20 := GenS1{F0: (v3 | (v3 | (v3 >> (((^(v16))) & 63)))), F1: (v2 + v2), F2: (4294967295 != 165)}
            _ = v19
            _ = v20
        } else {
            v21 := GenS1{F0: v3, F1: (v2 + "hello"), F2: true}
            var v22 int = -884
            _ = v21
            _ = v22
        }
        if true {
            if (4294967295 >= -2147483648) {
                var v23 uint = (^(v16))
                _ = v23
            }
        }
    } else {
        v24 := []int{}
        switch ("foo bar" + "a") {
        case "":
            v25 := GenS1{F0: v3, F1: v2, F2: false}
            _ = v25
        case "hello":
            v3 -= 42
        case "foo bar":
            var v26 uint = (v15 + 0)
            _ = v26
        }
        fmt.Println(v2, v2, v2)
        _ = v24
    }
    v27 := []int{-1, -271}
    var v28 int16 = 42
    if (true && (!false)) {
        if len(v27) > 0 {
            v27[0] = -2147483648
        }
        if true {
            if len(v27) >= 2 {
                v27 = v27[1:len(v27)]
            }
            var v29 int = -231
            var v30 bool = false
            _ = v29
            _ = v30
        }
        _ = len(v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
