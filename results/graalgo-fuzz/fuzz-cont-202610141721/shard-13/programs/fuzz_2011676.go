package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 int16, p1 uint64, p2 uint16) uint32 {
    var v6 uint8 = 0
    if false {
        var v7 uint64 = 586
        _ = v7
    }
    v8 := GenS1{F0: (v6 + uint8(p2))}
    v9 := GenS2{F0: 100, F1: 42, F2: 26}
    if (!(!((v6 | v6) >= (^(v6))))) {
        switch ("\n\t" + "") {
        case "hello":
            var v10 uint32 = 65535
            _ = v10
        default:
            var v11 string = "a"
            _ = v11
        }
    } else {
        v12 := []int{490, 164, 127}
        switch "" {
        case "hello":
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        default:
            fmt.Println(p1)
        }
        fmt.Println(p0, v6, p2)
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    _ = v9
    return 363
}
func genFn1(p0 int8, p1 uint) int64 {
    fmt.Println(p1, p1)
    if true {
        v13 := GenS3{F0: 2, F1: 0, F2: 0}
        _ = v13
    }
    var v14 uint64 = 42
    _ = p0
    _ = p1
    _ = v14
    return 1
}

func main() {
    var v1 uint8 = 141
    v2 := make(map[uint32]string)
    var v3 uint16 = 0
    var v4 bool = (0 >= 9223372036854775807)
    var v5 int16 = -32768
    v15 := []int{}
    if v4 {
        fmt.Println(v3, v3)
        var v16 int8 = 1
        _ = v16
    }
    var v17 string = (("hello" + ("a" + ("foo bar" + ("a" + ("\n\t" + ("foo bar" + (("\n\t" + "\n\t") + "hello"))))))) + "hello")
    v4 = false
    var v18 uint64 = 4294967295
    switch 127 {
    case -857:
        v17 = v17
    case 42:
        var v19 uint = 683
        _ = v19
    }
    v20 := GenS4{F0: 65535}
    var v21 uint16 = v3
    fmt.Println(v18, v4, v21)
    v18 -= v18
    var v22 int16 = (v5 * v5)
    v23 := GenS1{F0: v1}
    v18 = (v18 - v18)
    fmt.Println(v17, v21)
    v24 := GenS4{F0: 1}
    v17 = v17
    v25 := GenS5{F0: -686, F1: v5, F2: 4294967295}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
