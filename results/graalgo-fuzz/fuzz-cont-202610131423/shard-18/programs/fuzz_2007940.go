package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
func genFn0(p0 int16, p1 int32) int {
    v5 := GenS1{F0: 2, F1: -2147483648}
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 9223372036854775807
}
func genFn1(p0 uint, p1 uint) uint32 {
    var v7 int16 = 1
    v8 := GenS1{F0: 65535, F1: 127}
    fmt.Println(v7, p0, p0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float64 = ((1.5 / (1.5 / 3.14159)) - 3.14159)
    v2 := make(map[uint32]string)
    var v3 int16 = 32767
    var v4 uint16 = 1
    for fi9 := 0; fi9 < 3; fi9++ {
        v4 = v4
        v4 -= 0
        _ = fi9
    }
    if (v4 == v4) {
        v3 -= 52
        v3 = v3
        v1 = 1.5
    } else {
        for fi10 := 0; fi10 < 9; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                fmt.Println(fi11, fi11, v4)
                switch "\n\t" {
                case "foo bar":
                    var v12 uint64 = 335
                    _ = v12
                }
                delete(v2, genFn1(42, 42))
                delete(v2, 0)
                _ = fi11
            }
            var v13 uint16 = (v4 - v4)
            _ = fi10
            _ = v13
        }
        var v14 uint = 2
        _ = v14
    }
    v15 := GenS1{F0: v4, F1: 9223372036854775807}
    var v16 float64 = 3.14159
    v17 := []int{128}
    fmt.Println(v1, v3)
    var v18 float32 = 0.0
    _ = len(v17)
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 string = "\n\t"
        delete(v2, 2147483647)
        v18 = 3.14159
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
