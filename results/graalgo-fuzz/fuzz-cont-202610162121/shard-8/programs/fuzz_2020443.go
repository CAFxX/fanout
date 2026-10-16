package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 string }
func genFn0() uint8 {
    v5 := []int{652, 0}
    if (42 == 0) {
        var v6 int16 = 42
        v7 := GenS1{F0: -124, F1: "foo bar"}
        _ = v6
        _ = v7
    } else {
        var v8 int = -1
        var v9 float64 = 0.1
        _ = v8
        _ = v9
    }
    v10 := GenS1{F0: 552, F1: (("hello" + ("hello" + ("foo bar" + (("\n\t" + "foo bar") + "a")))) + "\n\t")}
    v11 := GenS1{F0: 1, F1: ((("" + "a") + "a") + "a")}
    _ = v10
    _ = v11
    _ = v5
    return 128
}

func main() {
    var v1 float32 = 1.5
    var v2 float64 = (((1.5 + 1.5) + 0.1) * 0.0)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    v1 = (v1 + float32((((0.0 + v2) * v2) / 3.14159)))
    v1 += -2.25
    v12 := GenS1{F0: 32768, F1: "\n\t"}
    fmt.Println(v4)
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 bool = false
        v14 = false
        if false {
            if false {
                var v15 int64 = 0
                _ = v15
            } else {
                v16 := []int{}
                _ = v16
            }
        }
        _ = fi13
        _ = v14
    }
    switch ((("" + "foo bar") + ("a" + ("foo bar" + ("foo bar" + "")))) + "a") {
    case "foo bar":
        var v17 int16 = -819
        _ = v17
    default:
        if (v1 != v1) {
            var v18 uint32 = 4294967295
            if true {
                for fi19 := 0; fi19 < 14; fi19++ {
                    if (false || true) {
                        for fi20 := 0; fi20 < 16; fi20++ {
                            v21 := GenS2{F0: ((false || false) || (-37 > (fi19 / (((fi19 - -220)) | 1))))}
                            v22 := GenS3{F0: 18446744073709551615}
                            var v23 float32 = (v1 / 3.14159)
                            _ = fi20
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := []int{(fi19 + (-(fi19))), fi19}
                        fmt.Println(v2)
                        _ = v24
                    }
                    if (!true) {
                        var v25 float64 = v2
                        for fi26 := 0; fi26 < 9; fi26++ {
                            fmt.Println(v2)
                            _ = fi26
                        }
                        if (int32(v18) > 2147483647) {
                            v27 := GenS4{F0: 127, F1: ""}
                            _ = v27
                        }
                        _ = v25
                    }
                    _ = fi19
                }
                var v28 int8 = 22
                _ = v28
            }
            _ = v18
        }
    }
    _ = len(v3)
    var v29 uint64 = uint64(v4)
    for fi30 := 0; fi30 < 15; fi30++ {
        var v31 bool = false
        var v32 int16 = 1000
        _ = fi30
        _ = v31
        _ = v32
    }
    var v33 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v4
}
