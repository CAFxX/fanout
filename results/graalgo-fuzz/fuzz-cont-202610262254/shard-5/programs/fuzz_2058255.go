package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint) int8 {
    v5 := []int{-1, 32767}
    _ = len(v5)
    _ = p0
    _ = v5
    return -1
}

func main() {
    v1 := []uint32{2147483648, 665, 0}
    v2 := []uint32{895, 42, 1}
    var v3 uint32 = 938
    var v4 int64 = 1
    v3 = v3
    switch ((("foo bar" + "") + ("\n\t" + "\n\t")) + "") {
    case "foo bar":
        var v6 int32 = 0
        _ = v6
    case "\n\t":
        v4 += v4
    case "":
        if false {
            var v7 string = "a"
            _ = len(v1)
            for fi8 := 0; fi8 < 3; fi8++ {
                var v9 int32 = -2147483648
                v3 = v3
                var v10 uint64 = 1
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
    default:
        if false {
            var v11 int32 = 124
            _ = v11
        } else {
            if false {
                v1 = append(v1, (v3 + v3))
                fmt.Println(v3, v3)
                var v12 string = ("" + ("" + ("\n\t" + ("foo bar" + "a"))))
                _ = v12
            } else {
                fmt.Println(v3, v3)
                var v13 uint = 65535
                _ = v13
            }
            if true {
                v14 := GenS1{F0: (v3 & v3), F1: 132}
                _ = v14
            }
            var v15 uint32 = 2147483648
            _ = v15
        }
    }
    var v16 bool = (42 > -32768)
    fmt.Println(v4)
    v17 := GenS2{F0: ((42 > 4294967295) || ((42 > 1) && (476 >= 457)))}
    var v18 int16 = 42
    var v19 string = ("foo bar" + "a")
    _ = len(v2)
    var v20 int = -186
    if v16 {
        v21 := GenS3{F0: v18}
        for fi22 := 0; fi22 < 18; fi22++ {
            v16 = ((v4 ^ 42) < -9223372036854775808)
            v3 -= v3
            _ = fi22
        }
        _ = v21
    }
    if v16 {
        if len(v1) > 0 {
            v1[0] = (v3 >> ((1) & 63))
        }
    } else {
        var v23 float64 = (float64(((((float32(-2.25) + float32(0.0)) - 0.0) * -2.25) / 1.0)) / 1.5)
        v24 := GenS2{F0: v16}
        var v25 int16 = v18
        _ = v23
        _ = v24
        _ = v25
    }
    v2 = append(v2, 2)
    v26 := GenS3{F0: v18}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
}
