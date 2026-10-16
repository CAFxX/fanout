package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 bool }

func main() {
    v1 := []int{335, 0, 4294967295}
    var v2 uint8 = 1
    v3 := []int{-440, 529, 1, 66}
    fmt.Println(v2)
    v2 = (v2 - (^(v2)))
    if false {
        var v4 int = -25
        v4 = -104
        _ = v4
    }
    if true {
        if false {
            if false {
                for fi5 := 0; fi5 < 6; fi5++ {
                    var v6 float64 = (0.0 * 1.5)
                    v6 = v6
                    var v7 float64 = 0.0
                    _ = fi5
                    _ = v6
                    _ = v7
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = -169
                }
            }
        }
        var v8 int64 = 1
        var v9 uint = 2
        _ = v8
        _ = v9
    } else {
        if true {
            if true {
                switch ("" + "\n\t") {
                case "":
                    var v10 int32 = 42
                    _ = v10
                default:
                    if true {
                        v11 := GenS1{F0: 4294967295, F1: 65535, F2: true}
                        v12 := GenS1{F0: 4294967295, F1: 255, F2: true}
                        var v13 uint32 = 2147483647
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
                v14 := GenS1{F0: 65535, F1: 392, F2: (-2147483648 == -505)}
                _ = v14
            }
            v2 -= (v2 & v2)
            var v15 uint8 = v2
            _ = v15
        } else {
            var v16 uint16 = 1
            v17 := GenS1{F0: 4294967295, F1: 42, F2: false}
            _ = v16
            _ = v17
        }
        var v18 int64 = 1
        _ = v18
    }
    v2 += v2
    var v19 float64 = 0.1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v20 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
