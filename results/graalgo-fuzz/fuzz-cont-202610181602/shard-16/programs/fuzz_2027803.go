package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 uint64; F1 int8 }
func genFn0(p0 int8, p1 int8, p2 bool) uint32 {
    if (533 < 42) {
        v7 := []int{32768, 2147483647}
        _ = v7
    }
    switch "" {
    case "hello":
        fmt.Println(p0)
    case "\n\t":
        v8 := GenS1{F0: 32767, F1: 42, F2: p2}
        _ = v8
    case "":
        v9 := GenS2{F0: -1}
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 uint32, p1 bool) uint8 {
    v10 := []int{845, -193}
    switch -2147483648 {
    case 1:
        p0 = (^(p0))
    case 4294967295:
        var v11 float64 = (1.5 + -2.25)
        _ = v11
    default:
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := GenS1{F0: 464, F1: -49, F2: (255 != 456)}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            fmt.Println(p0)
            _ = fi12
            _ = v13
        }
    }
    var v14 uint32 = 28
    var v15 uint = 318
    _ = p0
    _ = p1
    _ = v10
    _ = v14
    _ = v15
    return 128
}

func main() {
    var v1 bool = true
    v2 := []int{-2147483648}
    var v3 int8 = -1
    var v4 uint8 = 2
    var v5 int = 32768
    var v6 uint8 = v4
    var v16 string = "hello"
    var v17 uint64 = 2
    v18 := GenS1{F0: 32767, F1: -9223372036854775808, F2: v1}
    var v19 bool = false
    switch v16 {
    case "":
        v20 := GenS2{F0: 0}
        _ = v20
    }
    var v21 uint16 = 42
    if (!(42 == 870)) {
        fmt.Println(v16, v5, v21)
        var v22 int = (v5 >> ((44) & 63))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v22
    } else {
        fmt.Println(v1, v16)
    }
    switch v5 {
    case 42:
        if false {
            v23 := GenS2{F0: 2147483647}
            _ = len(v2)
            var v24 int8 = v3
            _ = v23
            _ = v24
        } else {
            var v25 float64 = ((((3.14159 - 1.5) / float64(float32(1.5))) * 0.0) * (3.14159 - float64(1e10)))
            _ = v25
        }
    case 831:
        v26 := GenS1{F0: (v21 | 557), F1: -9223372036854775808, F2: v1}
        _ = v26
    case 127:
        switch ("" + "\n\t") {
        case "\n\t":
            v17 = 357
        }
    default:
        var v27 int32 = 2147483647
        _ = v27
    }
    var v28 int32 = 2147483647
    v2 = append(v2, (v5 % ((-33) | 1)))
    v29 := GenS2{F0: (-(v28))}
    v30 := GenS1{F0: v21, F1: -9223372036854775808, F2: true}
    if (uint8(v17) < genFn1(genFn0(100, (-(v3)), false), (!(!v1)))) {
        var v31 string = ("hello" + (((v16 + v16) + (((v16 + v16) + v16) + v16)) + ("" + v16)))
        var v32 uint64 = v17
        for fi33 := 0; fi33 < 19; fi33++ {
            v34 := GenS3{F0: 0, F1: v32}
            var v35 int8 = v3
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = v31
        _ = v32
    }
    fmt.Println(v4, v6, v16)
    switch v5 {
    case 32768:
        v36 := GenS4{F0: (^(v17)), F1: v3}
        _ = v36
    case 42:
        var v37 uint32 = 2
        _ = v37
    }
    var v38 int32 = (v28 * (v28 % ((0) | 1)))
    v6 = 1
    v1 = (false && (false && v1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
