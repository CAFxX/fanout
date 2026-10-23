package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int, p1 uint16) bool {
    var v4 string = "\n\t"
    v5 := GenS1{F0: (p0 ^ (p0 & p0)), F1: (p1 | (p1 % ((255) | 1))), F2: -809}
    fmt.Println(v4, v4, p0)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return false
}
func genFn1(p0 int8) uint16 {
    v6 := GenS2{F0: 0}
    p0 = 0
    _ = p0
    _ = v6
    return 255
}

func main() {
    var v1 float64 = (3.14159 / float64(1.5))
    var v2 uint16 = 0
    v3 := make(map[uint32]string)
    if false {
        v7 := []int{}
        switch ("hello" + "") {
        case "":
            v8 := GenS1{F0: 2147483647, F1: 2, F2: -624}
            _ = v8
        }
        _ = v7
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            fmt.Println(v2)
            _ = fi10
        }
        _ = fi9
    }
    var v11 int64 = 9223372036854775807
    var v12 uint16 = (v2 >> ((673) & 63))
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := v3[0]
        v11 = v11
        _ = fi13
        _ = v14
    }
    v15 := []int{9223372036854775807, -631}
    v2 -= v12
    v1 -= 0.1
    for fi16 := 0; fi16 < 20; fi16++ {
        if false {
            for fi17 := 0; fi17 < 10; fi17++ {
                v12 = v2
                v18 := GenS3{F0: (v12 - (^(v12)))}
                if true {
                    switch "\n\t" {
                    case "":
                        var v19 float64 = 3.14159
                        _ = v19
                    case "foo bar":
                        v1 += 1.5
                    }
                } else {
                    var v20 int32 = -791
                    if false {
                        var v21 bool = (!(fi16 < fi17))
                        _ = v21
                    }
                    var v22 uint32 = 2
                    _ = v20
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
        }
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
}
