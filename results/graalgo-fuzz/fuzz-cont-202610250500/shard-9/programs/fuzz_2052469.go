package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 int }
func genFn0() int {
    v7 := []int{}
    var v8 bool = (1 >= 0)
    v9 := GenS1{F0: "foo bar", F1: 18446744073709551615, F2: -267}
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int32) uint16 {
    var v10 string = ""
    v11 := []int{127}
    switch -2147483648 {
    case 108:
        v12 := GenS1{F0: (v10 + v10), F1: 255, F2: 2147483647}
        _ = v12
    case 4294967295:
        var v13 uint = 42
        _ = v13
    default:
        if false {
            v10 = "foo bar"
        }
    }
    _ = len(v11)
    _ = p0
    _ = v10
    _ = v11
    return 1
}

func main() {
    var v1 float32 = -2.25
    var v2 int = 4294967295
    v3 := make(map[uint32]uint32)
    var v4 int32 = 2147483647
    v5 := make(map[string]uint32)
    v6 := []uint32{962, 42}
    delete(v3, 2147483648)
    var v14 uint32 = 600
    v4 = v4
    var v15 uint16 = 42
    var v16 int16 = 42
    for fi17 := 0; fi17 < 4; fi17++ {
        v18 := GenS1{F0: "\n\t", F1: 18446744073709551615, F2: v2}
        if false {
            v14 += (^(v14))
            var v19 uint8 = 1
            var v20 int32 = (v4 % ((42) | 1))
            _ = v19
            _ = v20
        } else {
            for fi21 := 0; fi21 < 16; fi21++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi21
            }
        }
        _ = fi17
        _ = v18
    }
    v2 = v2
    for fi22 := 0; fi22 < 2; fi22++ {
        var v23 uint32 = (^(v14))
        _ = fi22
        _ = v23
    }
    fmt.Println(v2, v15, v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
