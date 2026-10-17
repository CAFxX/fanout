package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 uint32, p1 bool) int32 {
    p1 = ((p0 - (p0 & p0)) > 0)
    var v6 uint8 = 1
    var v7 int32 = 1
    switch 295 {
    case -2147483648:
        v8 := []int{-493, 1, 42, 32767}
        _ = v8
    default:
        if (0.0 != float64(float32(1e10))) {
            v9 := []int{415, 9223372036854775807}
            var v10 uint = 2
            if (!(0.0 < (((0.0 / 1.5) + -2.25) + float32(0.0)))) {
                switch ("" + "foo bar") {
                case "":
                    if true {
                        p1 = true
                    } else {
                        var v11 uint8 = uint8(p0)
                        var v12 uint = 18446744073709551615
                        _ = v11
                        _ = v12
                    }
                case "a":
                    v6 = v6
                default:
                    var v13 bool = p1
                    _ = v13
                }
            }
            _ = v10
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v7
}

func main() {
    var v1 string = "a"
    var v2 bool = (52 >= 127)
    v3 := make(map[string]bool)
    var v4 int64 = -1
    v5 := make(map[uint32]string)
    v14 := GenS1{F0: 65535, F1: v4}
    v2 = (42 != 0)
    if (0 > 1) {
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS1{F0: 65535, F1: (-(v4))}
            for fi17 := 0; fi17 < 9; fi17++ {
                v18 := GenS1{F0: 42, F1: (v4 % ((-798) | 1))}
                v2 = (v1 <= (v1 + v1))
                switch (-(fi17)) {
                case -562:
                    v19 := GenS1{F0: 820, F1: v4}
                    _ = v19
                case 128:
                    var v20 int = 0
                    _ = v20
                default:
                    v1 = (v1 + "\n\t")
                    v2 = ((872 != 2) && v2)
                }
                _ = fi17
                _ = v18
            }
            fmt.Println(v1)
            _ = fi15
            _ = v16
        }
    } else {
        var v21 uint = 65535
        var v22 uint16 = 0
        _ = v21
        _ = v22
    }
    delete(v5, 2)
    var v23 uint32 = 42
    var v24 int8 = -1
    v23 -= (^(v23))
    switch (v1 + ("" + "a")) {
    case "hello":
        v25 := []int{}
        _ = v25
    case "a":
        if v2 {
            if v2 {
                v26 := []int{580}
                v27 := GenS2{F0: 4294967295, F1: 0}
                _ = v26
                _ = v27
            } else {
                fmt.Println(v24, v2, v2)
            }
            if true {
                var v28 int8 = v24
                _ = v28
            }
        }
    default:
        var v29 uint8 = 45
        _ = v29
    }
    switch 42 {
    case 32767:
        var v30 uint = 255
        _ = v30
    }
    var v31 uint32 = 972
    if v2 {
        v23 = v23
        var v32 int64 = v4
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
