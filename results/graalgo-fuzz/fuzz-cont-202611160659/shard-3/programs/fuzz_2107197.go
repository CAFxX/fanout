package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 bool, p1 uint8, p2 uint16) uint {
    var v6 float64 = -2.25
    var v7 uint16 = (^(p2))
    v7 -= v7
    fmt.Println(p0, p0)
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return uint(p1)
}
func genFn1(p0 uint8, p1 int8) uint16 {
    var v9 uint32 = 0
    p1 -= (p1 & p1)
    var v10 uint16 = 255
    v11 := GenS1{F0: "", F1: 205, F2: (v10 << ((0) & 63))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return (v10 << ((2) & 63))
}

func main() {
    var v1 bool = (!(1 == 0))
    var v2 uint16 = 32767
    var v3 string = ("" + "")
    var v4 float64 = 1.5
    var v5 uint16 = 1
    var v12 int16 = -32768
    v1 = (-1 <= -124)
    v1 = v1
    var v13 float32 = 0.1
    switch v3 {
    case "hello":
        v14 := GenS2{F0: 1, F1: 232}
        _ = v14
    case "a":
        v13 = 3.14159
    case "foo bar":
        var v15 uint = 255
        _ = v15
    }
    switch 2147483647 {
    case -2147483648:
        v16 := GenS3{F0: (^(v5)), F1: v2, F2: v3}
        _ = v16
    case 4294967295:
        var v17 uint8 = 255
        _ = v17
    case 1:
        var v18 bool = false
        _ = v18
    }
    var v19 uint8 = 42
    v20 := GenS3{F0: (v2 << ((42) & 63)), F1: genFn1((^(v19)), 20), F2: v3}
    switch ((v3 + "foo bar") + ("hello" + ((v3 + v3) + (v3 + ("a" + v3))))) {
    case "hello":
        if false {
            v21 := []int{}
            v21 = append(v21, 1)
            _ = v21
        } else {
            v3 = v3
            v3 = v3
        }
    default:
        v3 = (((v3 + ((("\n\t" + v3) + v3) + v3)) + (("\n\t" + v3) + v3)) + (v3 + "foo bar"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
