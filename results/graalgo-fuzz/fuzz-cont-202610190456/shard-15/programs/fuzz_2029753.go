package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 string; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int; F1 int }
func genFn0() string {
    if false {
        var v5 uint64 = 4294967295
        _ = v5
    }
    var v6 bool = (255 == 128)
    _ = v6
    return ""
}
func genFn1(p0 uint32) uint16 {
    v7 := []int{149}
    var v8 int32 = -1
    var v9 int = 181
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return uint16(p0)
}

func main() {
    var v1 uint = 255
    var v2 int64 = int64(v1)
    var v3 int = -737
    var v4 int64 = (v2 << ((1) & 63))
    fmt.Println(v2, v4, v1)
    var v10 float64 = float64(1e10)
    fmt.Println(v10, v2)
    v10 = (((3.14159 - float64(3.14159)) / float64(float32(((v10 * (float64(((float32(v10) - (float32(v10) * -2.25)) / (1e10 + (0.0 + (0.0 / 1e10))))) * v10)) * (((float64(-2.25) - 0.1) * (1.5 - v10)) * float64((-2.25 / 1.0))))))) / 1.5)
    if true {
        v3 -= 42
    } else {
        v1 = v1
        var v11 uint32 = 42
        _ = v11
    }
    if (!((1 == -250) || false)) {
        v12 := []int{4294967295}
        var v13 uint32 = 2
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: ("hello" + "a")}
    switch -1 {
    case -9223372036854775808:
        if true {
            v15 := GenS2{F0: 42, F1: 18446744073709551615}
            _ = v15
        } else {
            v16 := GenS3{F0: "", F1: 128, F2: 42}
            var v17 float32 = 1e10
            _ = v16
            _ = v17
        }
    case 32768:
        var v18 uint16 = 65535
        _ = v18
    default:
        v1 -= v1
    }
    v19 := GenS4{F0: (v3 + v3), F1: (v3 % (((v3 * (v3 * -1))) | 1))}
    v4 = (^(v4))
    v2 -= 9223372036854775807
    v10 = v10
    v20 := GenS2{F0: 0, F1: 18446744073709551615}
    v21 := []int{v3, 127, v3, (v3 << ((v1) & 63))}
    var v22 uint16 = genFn1(162)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
