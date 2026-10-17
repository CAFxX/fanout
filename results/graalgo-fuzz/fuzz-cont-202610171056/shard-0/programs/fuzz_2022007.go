package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int, p1 string) uint32 {
    var v4 string = p1
    v5 := []int{p0}
    var v6 string = p1
    _ = len(v5)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 187
}
func genFn1(p0 uint, p1 int, p2 uint) uint32 {
    if true {
        v7 := GenS1{F0: (p0 | p2)}
        _ = v7
    }
    p0 = p0
    p1 += 4294967295
    v8 := []int{p1}
    p2 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 19
}

func main() {
    var v1 int = 575
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v9 int = (v1 + 0)
    for fi10 := 0; fi10 < 7; fi10++ {
        v1 = 0
        _ = fi10
    }
    if (2 > 18446744073709551615) {
        if false {
            switch "foo bar" {
            case "":
                switch v9 {
                case 794:
                    for fi11 := 0; fi11 < 10; fi11++ {
                        v12 := v3["a"]
                        _ = fi11
                        _ = v12
                    }
                default:
                    fmt.Println(v9, v9)
                }
            default:
                var v13 float32 = 3.14159
                _ = v13
            }
            fmt.Println(v1, v1)
            var v14 int32 = -148
            _ = v14
        }
        var v15 bool = ((473 < 532) && false)
        _ = v15
    }
    for fi16 := 0; fi16 < 14; fi16++ {
        v17 := []int{fi16, 531}
        v1 += fi16
        v9 = (v9 * v9)
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: v1}
    fmt.Println(v1, v1)
    for fi19 := 0; fi19 < 11; fi19++ {
        _ = len(v2)
        var v20 int8 = -115
        _ = fi19
        _ = v20
    }
    v21 := GenS3{F0: 2, F1: 1}
    v22 := []int{v1, (v9 >> ((1) & 63)), (v1 - v9), (v1 << ((654) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
