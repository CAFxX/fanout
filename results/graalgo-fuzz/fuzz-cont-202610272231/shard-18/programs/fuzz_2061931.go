package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint32; F2 int8 }
func genFn0(p0 int32, p1 int8) uint16 {
    var v5 uint16 = 2
    var v6 uint16 = 255
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (v5 + v5)
}
func genFn1() int {
    switch "a" {
    case "hello":
        var v7 uint = 0
        _ = v7
    default:
        var v8 int8 = -128
        _ = v8
    }
    var v9 uint8 = 78
    switch -1 {
    case 2147483647:
        var v10 int32 = 1
        _ = v10
    case -814:
        if false {
            v11 := GenS1{F0: v9}
            var v12 int16 = 1000
            _ = v11
            _ = v12
        } else {
            v13 := []int{128}
            var v14 string = ("" + "hello")
            var v15 bool = (0 != 2)
            _ = v13
            _ = v14
            _ = v15
        }
    case 4294967295:
        v16 := GenS2{F0: 127}
        _ = v16
    default:
        v9 += (v9 * (^(v9)))
    }
    v17 := GenS2{F0: 127}
    if false {
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := []int{fi18}
            switch "\n\t" {
            case "hello":
                _ = len(v19)
            default:
                v9 = (v9 - (v9 - v9))
            }
            v20 := GenS1{F0: v9}
            _ = fi18
            _ = v19
            _ = v20
        }
    } else {
        if (-9223372036854775808 < -866) {
            for fi21 := 0; fi21 < 18; fi21++ {
                v22 := GenS1{F0: v9}
                _ = fi21
                _ = v22
            }
            v23 := []int{32768, 2147483647}
            v9 = (v9 >> ((896) & 63))
            _ = v23
        } else {
            v24 := GenS3{F0: (!true)}
            _ = v24
        }
        fmt.Println(v9)
    }
    _ = v17
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]bool)
    var v2 int = -809
    var v3 uint32 = 42
    var v4 uint32 = 0
    v2 -= v2
    var v25 uint = 2
    v26 := []int{}
    v3 = (v3 / ((v4) | 1))
    var v27 int32 = 1
    v25 = 878
    var v28 string = "foo bar"
    v4 = v4
    v4 -= 4294967295
    switch 0 {
    case 32767:
        if true {
            for fi29 := 0; fi29 < 12; fi29++ {
                v25 += uint(v3)
                _ = fi29
            }
            v30 := GenS4{F0: v28, F1: 2147483648, F2: int8(v25)}
            _ = v30
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                v27 -= -670
                _ = fi31
            }
            v27 += v27
        }
    case -2147483648:
        v32 := GenS3{F0: (false || (-2.25 == 1e10))}
        _ = v32
    case -9223372036854775808:
        var v33 int32 = v27
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
