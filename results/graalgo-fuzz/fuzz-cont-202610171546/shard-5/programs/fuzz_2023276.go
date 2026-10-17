package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int32 }
func genFn0(p0 int8, p1 uint32, p2 int32) uint {
    v7 := GenS1{F0: (p1 & 806)}
    for fi8 := 0; fi8 < 13; fi8++ {
        switch ("" + "") {
        case "foo bar":
            var v9 uint8 = 1
            _ = v9
        case "a":
            v10 := []int{fi8, fi8, 44}
            _ = v10
        default:
            var v11 uint16 = 255
            _ = v11
        }
        switch "a" {
        case "":
            v12 := GenS2{F0: 18446744073709551615, F1: 13, F2: (p2 << ((4294967295) & 63))}
            _ = v12
        case "a":
            var v13 int16 = int16(p2)
            _ = v13
        case "hello":
            switch -2147483648 {
            case 1:
                fmt.Println(p2, p0, fi8)
            case -667:
                var v14 int64 = 579
                _ = v14
            default:
                var v15 int = (fi8 >> ((0) & 63))
                _ = v15
            }
        }
        _ = fi8
    }
    v16 := []int{128, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v7
    return 0
}
func genFn1(p0 int16) uint8 {
    p0 = 1
    p0 += 0
    var v17 int8 = -1
    fmt.Println(v17)
    if true {
        v18 := []int{}
        _ = v18
    }
    _ = p0
    _ = v17
    return 1
}

func main() {
    var v1 string = "a"
    v2 := make(map[string]uint32)
    var v3 string = (((v1 + "foo bar") + "hello") + v1)
    v4 := make(map[uint32]string)
    var v5 uint32 = 2
    var v6 uint8 = 2
    switch 2147483647 {
    case 9223372036854775807:
        var v19 int64 = 887
        _ = v19
    case -9223372036854775808:
        fmt.Println(v6, v5, v6)
    }
    var v20 string = (((v3 + "") + ("\n\t" + v1)) + v3)
    var v21 int32 = -22
    var v22 int = int(v21)
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 int32 = v21
        _ = fi23
        _ = v24
    }
    if false {
        v25 := []int{v22}
        v26 := GenS1{F0: v5}
        v27 := GenS2{F0: 79, F1: genFn1(0), F2: -943}
        _ = v25
        _ = v26
        _ = v27
    }
    if true {
        var v28 string = v3
        _ = v28
    }
    fmt.Println(v22, v5, v22)
    for fi29 := 0; fi29 < 9; fi29++ {
        var v30 uint64 = 847
        var v31 float64 = (0.0 - 1.5)
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 float32 = (0.0 * (3.14159 / -2.25))
    fmt.Println(v22, v32)
    if false {
        var v33 uint = 255
        var v34 float64 = (0.1 + ((((0.1 + 0.1) + float64(v32)) + float64(v32)) * (0.1 / 0.1)))
        var v35 int64 = 42
        _ = v33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
