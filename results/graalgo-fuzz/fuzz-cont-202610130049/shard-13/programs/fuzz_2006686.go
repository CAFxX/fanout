package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint32, p1 uint32) string {
    if (false && false) {
        for fi7 := 0; fi7 < 5; fi7++ {
            var v8 int = (fi7 << ((1) & 63))
            v9 := []int{v8, -9223372036854775808, 2147483647, 127}
            var v10 bool = (2147483647 != 42)
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        p1 -= p0
    } else {
        var v11 int16 = 42
        v11 -= 516
        switch -757 {
        case 4294967295:
            fmt.Println(p0, p1, v11)
        case -913:
            fmt.Println(v11)
        case 42:
            v12 := GenS1{F0: (false || true), F1: 729, F2: 255}
            _ = v12
        }
        _ = v11
    }
    p1 = (^(p0))
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 string = ""
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    return ("hello" + "a")
}

func main() {
    var v1 int64 = 127
    var v2 uint8 = 0
    v3 := make(map[uint32]uint32)
    var v4 uint8 = v2
    var v5 float64 = 0.1
    var v6 float64 = v5
    if true {
        var v15 bool = (18446744073709551615 != uint64(v2))
        var v16 int32 = 0
        _ = v15
        _ = v16
    }
    v6 = ((v6 - v6) * (v5 * -2.25))
    for fi17 := 0; fi17 < 10; fi17++ {
        v18 := GenS1{F0: true, F1: 1, F2: 254}
        _ = fi17
        _ = v18
    }
    fmt.Println(v4, v5, v2)
    if ((int8(v4) == -1) && false) {
        v1 = v1
    }
    var v19 int8 = 0
    v20 := v3[0]
    if (((v20 + uint32(v4)) > 2) && true) {
        var v21 int32 = 42
        _ = v21
    } else {
        var v22 float32 = -2.25
        v23 := GenS2{F0: 32768}
        var v24 int16 = -369
        _ = v22
        _ = v23
        _ = v24
    }
    switch "\n\t" {
    case "foo bar":
        if true {
            fmt.Println(v2, v5, v6)
            fmt.Println(v5, v20, v2)
            var v25 int32 = 0
            _ = v25
        }
    }
    switch 32767 {
    case -1:
        fmt.Println(v4)
    default:
        v2 = (v2 % ((v2) | 1))
    }
    fmt.Println(v2)
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 float64 = 0.0
        fmt.Println(v1, v27, v1)
        _ = len(v3)
        _ = fi26
        _ = v27
    }
    var v28 int32 = 35
    if false {
        v29 := GenS3{F0: ((!true) || (1 >= 42)), F1: 255}
        _ = v29
    }
    if true {
        v30 := GenS2{F0: -425}
        switch 694 {
        case 1:
            fmt.Println(v1, v2)
        }
        _ = v30
    }
    if false {
        fmt.Println(v28, v6)
    }
    v31 := []int{int(v4), -114, -824, 92}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
