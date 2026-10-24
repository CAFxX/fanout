package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }

func main() {
    var v1 uint = 4294967295
    var v2 int16 = 1000
    var v3 string = "hello"
    var v4 float32 = 0.1
    var v5 float32 = 1.5
    if (1 > 283) {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint16 = 462
            _ = fi6
            _ = v7
        }
        v4 = 0.0
        if true {
            switch v3 {
            case "hello":
                v8 := []int{}
                _ = v8
            case "a":
                switch "a" {
                case "":
                    v9 := GenS1{F0: 100, F1: 606}
                    _ = v9
                default:
                    for fi10 := 0; fi10 < 4; fi10++ {
                        v11 := GenS1{F0: 127, F1: 4294967295}
                        var v12 float32 = ((v4 / float32(float64(v5))) + float32(-2.25))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                }
            case "":
                v13 := []int{-13}
                _ = v13
            default:
                var v14 uint16 = 65535
                _ = v14
            }
            var v15 float32 = v5
            _ = v15
        }
    }
    var v16 uint16 = 255
    var v17 float64 = 1e10
    v18 := []int{}
    fmt.Println(v1, v1)
    v16 += (v16 | v16)
    v5 = v5
    v19 := GenS1{F0: 42, F1: 0}
    v16 = (v16 & 2)
    v17 = -2.25
    var v20 float32 = float32(v17)
    var v21 uint32 = 4294967295
    switch 2147483647 {
    case 42:
        var v22 int16 = (v2 & (v2 + (^(v2))))
        _ = v22
    case -71:
        v3 = "\n\t"
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
