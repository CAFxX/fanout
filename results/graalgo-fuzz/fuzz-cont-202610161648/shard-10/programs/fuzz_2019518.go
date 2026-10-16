package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64) int64 {
    fmt.Println(p0)
    p0 -= 614
    v5 := []int{}
    _ = p0
    _ = v5
    return 311
}

func main() {
    var v1 uint32 = 1
    var v2 bool = (!false)
    var v3 uint = 18446744073709551615
    var v4 uint64 = 576
    v6 := GenS1{F0: 2147483647, F1: 42, F2: 100}
    if (!(!v2)) {
        if v2 {
            v7 := []int{42, -2147483648, 32767, 32767}
            _ = v7
        }
        var v8 float32 = -2.25
        _ = v8
    }
    if (!false) {
        var v9 int64 = 9223372036854775807
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int32 = 1
            fmt.Println(v3, v2, v1)
            _ = fi10
            _ = v11
        }
        var v12 uint32 = 784
        _ = v12
        _ = v9
    }
    fmt.Println(v1, v4, v2)
    var v13 float64 = 1.5
    switch ("hello" + ("foo bar" + ("a" + "a"))) {
    case "":
        v14 := []int{}
        _ = v14
    case "foo bar":
        if false {
            v15 := GenS2{F0: 3, F1: v1}
            for fi16 := 0; fi16 < 7; fi16++ {
                for fi17 := 0; fi17 < 11; fi17++ {
                    var v18 float32 = 1e10
                    v19 := []int{fi16, (^(fi16))}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v3 = 18446744073709551615
                _ = fi16
            }
            _ = v15
        }
    case "a":
        var v20 int = -2147483648
        _ = v20
    default:
        var v21 float64 = (1.5 / (v13 / v13))
        _ = v21
    }
    v22 := GenS1{F0: 32767, F1: 1, F2: -89}
    var v23 uint8 = 2
    var v24 string = "a"
    v25 := GenS3{F0: 1}
    for fi26 := 0; fi26 < 12; fi26++ {
        v27 := []int{(fi26 % (((^(fi26))) | 1))}
        _ = fi26
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v6
}
