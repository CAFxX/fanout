package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int }

func main() {
    var v1 uint16 = 280
    var v2 bool = (!(42 != 362))
    var v3 float64 = (1.5 * 3.14159)
    v4 := make(map[string]bool)
    v5 := make(map[int]string)
    var v6 uint = 255
    var v7 uint64 = 461
    v6 = uint(v1)
    var v8 uint16 = (v1 * (^(v1)))
    v6 -= (v6 % (((^(v6))) | 1))
    var v9 int16 = 32767
    v8 -= (v8 / ((v8) | 1))
    v2 = true
    v2 = v2
    if v2 {
        var v10 uint = v6
        var v11 uint8 = 49
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 335, F1: 127}
        _ = v12
    }
    var v13 uint8 = 1
    v4[("\n\t" + (("\n\t" + "\n\t") + ""))] = v2
    if false {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 bool = (!false)
            v16 := []int{}
            _ = fi14
            _ = v15
            _ = v16
        }
        fmt.Println(v2, v7, v8)
        fmt.Println(v9)
    }
    var v17 int32 = -61
    v4["\n\t"] = v2
    var v18 uint16 = v1
    switch 553 {
    case 127:
        fmt.Println(v8, v7)
    }
    var v19 bool = (!v2)
    if v2 {
        for fi20 := 0; fi20 < 7; fi20++ {
            switch "\n\t" {
            case "":
                v3 -= (((v3 - v3) / (v3 - v3)) - v3)
            case "hello":
                v3 -= v3
            case "a":
                var v21 uint64 = (v7 << ((153) & 63))
                _ = v21
            default:
                var v22 int8 = -1
                _ = v22
            }
            _ = fi20
        }
        switch 4294967295 {
        case -9223372036854775808:
            v23 := GenS2{F0: ((v2 || false) && true), F1: v7, F2: v9}
            _ = v23
        case 128:
            for fi24 := 0; fi24 < 20; fi24++ {
                for fi25 := 0; fi25 < 15; fi25++ {
                    _ = len(v5)
                    v6 -= 0
                    var v26 float64 = (v3 / (-2.25 + -2.25))
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        case -1:
            v6 = v6
        }
        v27 := GenS2{F0: v2, F1: v7, F2: (^(v9))}
        _ = v27
    } else {
        v2 = ((v17 << (((v6 ^ 255)) & 63)) < (^(v17)))
        for fi28 := 0; fi28 < 7; fi28++ {
            fmt.Println(v13)
            var v29 int32 = v17
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            v31 := GenS3{F0: 913, F1: (v17 & v17), F2: 32767}
            if false {
                for fi32 := 0; fi32 < 13; fi32++ {
                    v18 = 0
                    v33 := []int{(fi30 ^ (fi32 & fi32)), (-(fi32)), (^(fi32)), 4294967295}
                    var v34 int = (-(fi32))
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                var v35 uint8 = v13
                _ = v35
            } else {
                var v36 uint64 = 4294967295
                v37 := []int{(fi30 % ((fi30) | 1)), -1, (fi30 << ((uint(v13)) & 63))}
                _ = v36
                _ = v37
            }
            fmt.Println(v8, fi30)
            _ = fi30
            _ = v31
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
