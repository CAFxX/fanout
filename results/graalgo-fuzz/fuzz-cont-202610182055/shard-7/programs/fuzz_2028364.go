package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 uint8 }
type GenS2 struct { F0 string; F1 uint16; F2 string }
type GenS3 struct { F0 string; F1 int32; F2 int32 }
type GenS4 struct { F0 uint16; F1 int64 }

func main() {
    var v1 float64 = float64(1.5)
    v2 := []int{-659, -221}
    var v3 string = "a"
    v4 := make(map[int]uint32)
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint32 = 42
        _ = fi6
        _ = v7
    }
    v8 := GenS1{F0: (!(1 <= 42)), F1: -1, F2: 121}
    switch "foo bar" {
    case "foo bar":
        var v9 uint = 65535
        _ = v9
    case "hello":
        var v10 bool = ((!false) && ((1 != 1) && false))
        _ = v10
    case "":
        if false {
            v3 = ""
        } else {
            v11 := GenS1{F0: false, F1: 0, F2: 127}
            var v12 float32 = 1.5
            _ = v11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                switch fi13 {
                case -9223372036854775808:
                    v4[fi13] = 0
                default:
                    v14 := GenS2{F0: ("" + (v3 + v3)), F1: 0, F2: "\n\t"}
                    _ = v14
                }
                fmt.Println(v1)
            } else {
                if (1 != -613) {
                    v1 = v1
                    v4[fi13] = 837
                    v15 := GenS2{F0: v3, F1: 32767, F2: v3}
                    _ = v15
                } else {
                    var v16 uint64 = 17
                    _ = v16
                }
                var v17 float32 = 0.0
                var v18 uint64 = 0
                _ = v17
                _ = v18
            }
            v2 = append(v2, 4294967295)
            fmt.Println(v3)
            _ = fi13
        }
    }
    var v19 float64 = 1.5
    var v20 int16 = -1
    switch -1 {
    case 2147483647:
        v2 = append(v2, 4294967295)
    case 128:
        for fi21 := 0; fi21 < 9; fi21++ {
            switch (fi21 >> ((784) & 63)) {
            case 128:
                v2 = append(v2, fi21)
            case -2147483648:
                var v22 float32 = float32((v1 * (((v19 * float64(0.0)) * v19) + v1)))
                _ = v22
            default:
                if false {
                    v23 := GenS3{F0: v3, F1: -2147483648, F2: 1}
                    _ = v23
                }
            }
            fmt.Println(fi21, v3, v3)
            v1 = ((v1 + v19) / (v19 + 0.0))
            _ = fi21
        }
    default:
        v24 := GenS4{F0: 32767, F1: 0}
        _ = v24
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    fmt.Println(v3, v3, v1)
    var v25 int8 = -96
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := GenS2{F0: v3, F1: 2, F2: (v3 + "\n\t")}
        if true {
            v3 = v3
        }
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        var v29 int16 = (v20 - 32767)
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
