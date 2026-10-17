package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
func genFn0(p0 int32, p1 string, p2 int32) string {
    fmt.Println(p0, p1)
    var v6 float64 = (3.14159 - ((1e10 / 0.1) / (float64(-2.25) * 0.1)))
    v7 := GenS1{F0: 255, F1: 2147483647, F2: -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p1
}
func genFn1() uint {
    if (("hello" + "hello") >= ("" + ("a" + ("foo bar" + "a")))) {
        var v8 int16 = -273
        var v9 int8 = -128
        _ = v8
        _ = v9
    } else {
        var v10 uint = 35
        _ = v10
    }
    var v11 uint64 = 116
    v11 = (v11 - (v11 * 230))
    var v12 int = -981
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 uint = 738
    var v2 bool = true
    var v3 uint16 = 65535
    var v4 bool = false
    var v5 uint64 = 4294967295
    if false {
        var v13 string = "foo bar"
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v5 = (v5 - 42)
        if true {
            var v15 int = fi14
            _ = v15
        }
        v16 := []int{fi14, fi14, int(v5)}
        _ = fi14
        _ = v16
    }
    if v4 {
        for fi17 := 0; fi17 < 10; fi17++ {
            v18 := []int{}
            if false {
                switch "" {
                case "foo bar":
                    var v19 int8 = int8(v3)
                    _ = v19
                default:
                    if len(v18) > 0 {
                        v18[0] = (fi17 * (^(fi17)))
                    }
                }
            }
            if true {
                fmt.Println(v3, v2, v4)
                v20 := GenS2{F0: -9223372036854775808, F1: 3}
                v21 := GenS1{F0: v1, F1: fi17, F2: -407}
                _ = v20
                _ = v21
            } else {
                if len(v18) > 0 {
                    v18[0] = fi17
                }
            }
            _ = fi17
            _ = v18
        }
    } else {
        fmt.Println(v2)
        v5 = v5
        v1 = 459
    }
    v22 := GenS1{F0: (v1 / (((v1 / ((v1) | 1))) | 1)), F1: 128, F2: 1}
    switch -964 {
    case 32768:
        fmt.Println(v5, v3, v2)
    }
    v23 := []int{-407, 4294967295, 4294967295, int(v3)}
    v2 = (genFn1() == 1)
    v5 = (v5 | 42)
    v2 = false
    v2 = (!false)
    var v24 int16 = 984
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
