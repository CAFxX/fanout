package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0() uint {
    for fi4 := 0; fi4 < 13; fi4++ {
        var v5 uint = 4294967295
        var v6 float32 = 3.14159
        switch ("hello" + "a") {
        case "a":
            v7 := []int{9223372036854775807, 804}
            _ = v7
        case "foo bar":
            var v8 string = "hello"
            _ = v8
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint32 = 2147483647
        var v11 uint32 = v10
        var v12 uint64 = 4294967295
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        if false {
            var v14 int32 = 42
            _ = v14
        }
        var v15 int8 = 103
        _ = fi13
        _ = v15
    }
    if (279 >= -300) {
        var v16 uint16 = 2
        v17 := GenS1{F0: ((("a" < "") && false) && (true || false)), F1: 224}
        _ = v16
        _ = v17
    }
    var v18 uint16 = 255
    _ = v18
    return 414
}

func main() {
    v1 := []uint32{0, 502, 65535}
    v2 := make(map[int]int)
    var v3 uint32 = 71
    fmt.Println(v3)
    var v19 uint64 = 1
    switch "foo bar" {
    case "hello":
        switch ((("foo bar" + ("\n\t" + "a")) + "\n\t") + "hello") {
        case "hello":
            v3 = v3
        default:
            var v20 float64 = ((float64(0.0) * 1e10) + 0.0)
            _ = v20
        }
    default:
        var v21 float32 = 0.1
        _ = v21
    }
    var v22 bool = (18446744073709551615 <= 0)
    v23 := GenS1{F0: true, F1: 187}
    var v24 int64 = -415
    fmt.Println(v22, v3, v3)
    v3 = 1
    for fi25 := 0; fi25 < 14; fi25++ {
        if (((float64((0.0 - 1.5)) - 1.5) * float64(0.0)) > 0.1) {
            v19 = 4294967295
            var v26 float64 = 0.1
            _ = v26
        }
        _ = fi25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
