package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 bool }
func genFn0(p0 int, p1 string) uint {
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 float64 = 1e10
        v6 := GenS1{F0: true, F1: 2}
        v5 = 1.5
        _ = fi4
        _ = v5
        _ = v6
    }
    p1 = "a"
    p1 = (p1 + ((("hello" + "") + ((p1 + p1) + p1)) + (p1 + p1)))
    switch ("\n\t" + p1) {
    case "hello":
        v7 := GenS2{F0: 1, F1: 255, F2: 42}
        _ = v7
    default:
        var v8 uint8 = 127
        _ = v8
    }
    if true {
        for fi9 := 0; fi9 < 16; fi9++ {
            if true {
                v10 := GenS2{F0: 0, F1: 32767, F2: -1}
                var v11 int16 = 42
                _ = v10
                _ = v11
            }
            var v12 uint64 = 1
            v13 := GenS2{F0: 2, F1: 963, F2: 100}
            _ = fi9
            _ = v12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := []uint32{4294967295}
    var v2 uint16 = 255
    var v3 bool = (-2147483648 == 1)
    var v14 float64 = (0.1 + 0.1)
    v15 := GenS3{F0: 2, F1: 259}
    v16 := GenS4{F0: (v14 >= ((v14 * 1.5) * v14))}
    if len(v1) > 0 {
        v1[0] = 42
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (!false) {
        v2 = v2
        v3 = (v3 && (!false))
    }
    v14 = (((1.5 + v14) * v14) / v14)
    fmt.Println(v14, v2, v2)
    v2 = v2
    v3 = (893 <= 1)
    v3 = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
