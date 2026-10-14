package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int; F1 int64; F2 bool }
func genFn0() uint16 {
    var v6 int32 = -1
    if (0 > 2147483648) {
        fmt.Println(v6)
        v7 := []int{754, 599, 42}
        v8 := GenS1{F0: 109, F1: (127 != 42), F2: 2}
        _ = v7
        _ = v8
    }
    var v9 uint32 = 1
    v6 = v6
    var v10 float64 = 0.0
    _ = v10
    _ = v6
    _ = v9
    return 255
}
func genFn1(p0 uint) uint64 {
    var v11 uint32 = 499
    var v12 int64 = -1
    v13 := GenS2{F0: 1, F1: 1}
    fmt.Println(p0)
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return uint64(v11)
}

func main() {
    var v1 int32 = 621
    var v2 int32 = 42
    v3 := make(map[uint32]bool)
    v4 := make(map[string]string)
    var v5 float64 = 1e10
    v14 := GenS1{F0: 127, F1: false, F2: 504}
    v15 := []int{0, 79}
    for fi16 := 0; fi16 < 15; fi16++ {
        if (!false) {
            v2 += 2147483647
        } else {
            var v17 string = "hello"
            v17 = v17
            if true {
                var v18 uint64 = 816
                var v19 int = int(v18)
                _ = v18
                _ = v19
            }
            _ = v17
        }
        var v20 bool = (true || false)
        _ = fi16
        _ = v20
    }
    if (2 < 1) {
        switch ("" + "\n\t") {
        case "hello":
            v21 := v4["a"]
            _ = v21
        case "":
            fmt.Println(v1, v5)
        default:
            v5 += -2.25
        }
        if (true && false) {
            v5 += (1.5 + (-2.25 - ((v5 + (0.1 + 0.0)) / (v5 * float64((0.1 - (0.0 * 0.1)))))))
            fmt.Println(v2, v5, v5)
        } else {
            fmt.Println(v1, v5, v5)
            v2 = v1
            v22 := GenS2{F0: 255, F1: 42}
            _ = v22
        }
    }
    delete(v4, "foo bar")
    for fi23 := 0; fi23 < 12; fi23++ {
        if len(v15) > 0 {
            v15[0] = fi23
        }
        v1 = v2
        fmt.Println(fi23, v1, v5)
        _ = fi23
    }
    v1 = v1
    delete(v4, "a")
    var v24 uint8 = 57
    if true {
        fmt.Println(v2)
        v25 := GenS3{F0: 0, F1: -9223372036854775808, F2: true}
        _ = v25
    }
    v2 += v2
    v2 -= -2147483648
    switch int(v1) {
    case 42:
        v1 -= (^(v2))
    case 600:
        v24 += 14
    case -954:
        var v26 int = -2147483648
        _ = v26
    }
    v27 := GenS2{F0: 0, F1: 4294967295}
    fmt.Println(v1, v5, v24)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
