package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint16 }
func genFn0() string {
    var v4 float32 = 0.0
    v5 := GenS1{F0: -515, F1: 647}
    if false {
        v6 := GenS1{F0: 42, F1: 1}
        switch 32767 {
        case 2147483647:
            v4 += v4
        case 886:
            if false {
                fmt.Println(v4)
                v7 := []int{}
                _ = v7
            }
        case 42:
            if true {
                fmt.Println(v4)
                var v8 uint32 = 714
                fmt.Println(v8)
                _ = v8
            }
        default:
            v4 += v4
        }
        v4 = (v4 / (((v4 / (v4 * float32(1.5))) * 0.1) / v4))
        _ = v6
    }
    _ = v4
    _ = v5
    return ("hello" + "foo bar")
}
func genFn1() int64 {
    v9 := []int{}
    var v10 uint16 = 1
    v11 := GenS1{F0: 0, F1: 388}
    _ = v10
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := []uint32{2, 2147483647, 2147483648}
    var v2 uint8 = 128
    var v3 uint8 = v2
    fmt.Println(v2)
    _ = len(v1)
    fmt.Println(v2, v2)
    var v12 int64 = 271
    var v13 int16 = -171
    v12 = genFn1()
    var v14 uint64 = 616
    var v15 int8 = 1
    v16 := GenS2{F0: 764, F1: v2, F2: 32767}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
