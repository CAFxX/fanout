package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint64 }
func genFn0() string {
    if true {
        v5 := []int{-675, -338}
        _ = v5
    } else {
        var v6 int = 2147483647
        _ = v6
    }
    var v7 int = -9223372036854775808
    var v8 uint64 = 4294967295
    var v9 float32 = -2.25
    fmt.Println(v9, v7, v7)
    _ = v7
    _ = v8
    _ = v9
    return ("a" + (("" + ("\n\t" + "\n\t")) + "hello"))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{896, 0}
    var v3 uint32 = 2147483647
    v4 := make(map[int]string)
    var v10 string = "\n\t"
    fmt.Println(v10)
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float64 = 0.1
        if true {
            if true {
                v13 := GenS1{F0: 34, F1: -1, F2: 4294967295}
                _ = v13
            }
        } else {
            var v14 bool = true
            v15 := v4[fi11]
            switch fi11 {
            case 1:
                v10 = v10
            case 32768:
                fmt.Println(v12, v3, v15)
            case 127:
                switch (fi11 / ((fi11) | 1)) {
                case 2147483647:
                    v1[2147483647] = fi11
                case 599:
                    var v16 int32 = 0
                    _ = v16
                }
            default:
                for fi17 := 0; fi17 < 2; fi17++ {
                    var v18 bool = false
                    v19 := GenS1{F0: 13, F1: 1, F2: 0}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = v14
            _ = v15
        }
        _ = fi11
        _ = v12
    }
    v10 = "hello"
    if true {
        var v20 float64 = (3.14159 * -2.25)
        _ = v20
    }
    var v21 string = (v10 + v10)
    v22 := v4[1]
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    for fi23 := 0; fi23 < 2; fi23++ {
        var v24 uint16 = 42
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
