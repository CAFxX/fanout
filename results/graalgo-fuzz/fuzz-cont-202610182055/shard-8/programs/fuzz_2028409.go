package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8; F1 int16; F2 int32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int; F1 bool }
type GenS6 struct { F0 int; F1 int8; F2 string }
type GenS7 struct { F0 int16; F1 int8; F2 int16 }
type GenS8 struct { F0 int8; F1 int64 }
type GenS9 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    var v1 int8 = 42
    var v2 uint32 = 2147483647
    var v3 string = ("a" + ("a" + "hello"))
    switch 1 {
    case 2147483647:
        for fi4 := 0; fi4 < 6; fi4++ {
            v5 := []int{}
            v2 = v2
            v1 = int8(fi4)
            _ = fi4
            _ = v5
        }
    case 32767:
        var v6 float64 = (0.0 / 1e10)
        _ = v6
    default:
        v7 := GenS1{F0: 1}
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint = 2
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: 0}
    v11 := GenS3{F0: v1, F1: 0, F2: -2147483648}
    v12 := GenS4{F0: 255, F1: 255}
    v2 = v2
    v13 := GenS3{F0: 0, F1: -901, F2: 691}
    v14 := GenS5{F0: 10, F1: true}
    for fi15 := 0; fi15 < 14; fi15++ {
        if false {
            for fi16 := 0; fi16 < 15; fi16++ {
                if false {
                    v2 -= 87
                    v17 := GenS6{F0: 1, F1: v1, F2: ""}
                    var v18 float64 = float64((0.0 * 1.5))
                    _ = v17
                    _ = v18
                } else {
                    switch fi16 {
                    case -1:
                        v2 = v2
                    case -888:
                        v1 -= v1
                    case 0:
                        v3 = (("foo bar" + "hello") + v3)
                    }
                    switch -2147483648 {
                    case 127:
                        v19 := GenS7{F0: -1, F1: (v1 % ((v1) | 1)), F2: -989}
                        _ = v19
                    case 32768:
                        v20 := GenS3{F0: (-(v1)), F1: -1, F2: 42}
                        var v21 uint16 = 0
                        v3 = v3
                        _ = v20
                        _ = v21
                    default:
                        var v22 string = v3
                        _ = v22
                    }
                    var v23 int32 = -2147483648
                    var v24 int = fi16
                    _ = v23
                    _ = v24
                }
                fmt.Println(v2, v3)
                v2 = v2
                _ = fi16
            }
        } else {
            v25 := GenS8{F0: (v1 / ((v1) | 1)), F1: -863}
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v2, fi26)
                if (651 != 32767) {
                    var v27 uint64 = 42
                    v1 = -1
                    switch 32767 {
                    case 32767:
                        v28 := GenS9{F0: v2, F1: v2, F2: 1}
                        _ = v28
                    case 9223372036854775807:
                        var v29 int64 = 800
                        _ = v29
                    case 9:
                        v30 := []int{}
                        _ = v30
                    }
                    v31 := []int{fi26, fi26, fi15}
                    _ = v27
                    _ = v31
                }
                _ = fi26
            }
            v3 = ""
            _ = v25
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
