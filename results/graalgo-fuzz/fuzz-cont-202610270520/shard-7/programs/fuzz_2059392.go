package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 bool; F2 uint16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 uint }

func main() {
    var v1 int = 128
    var v2 string = ("hello" + "")
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 7; fi4++ {
        v5 := GenS1{F0: false}
        if true {
            for fi6 := 0; fi6 < 2; fi6++ {
                fmt.Println(fi4, fi4)
                _ = fi6
            }
            v2 = ("" + v2)
        } else {
            v2 = ("" + v2)
            v7 := []int{(-(v1))}
            _ = v7
        }
        var v8 uint8 = 165
        _ = fi4
        _ = v5
        _ = v8
    }
    switch v2 {
    case "hello":
        v1 = -1
    case "foo bar":
        v1 -= (v1 << ((4294967295) & 63))
    case "a":
        if false {
            for fi9 := 0; fi9 < 20; fi9++ {
                var v10 uint8 = 158
                _ = fi9
                _ = v10
            }
            var v11 uint8 = 127
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := GenS1{F0: false}
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            v14 := []int{716, v1, -9223372036854775808}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            for fi15 := 0; fi15 < 18; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    v3[-9223372036854775808] = (v2 + v2)
                    _ = fi16
                }
                _ = fi15
            }
            _ = v14
        }
    default:
        v17 := []int{v1, v1}
        _ = v17
    }
    var v18 uint64 = 0
    v19 := []int{-2147483648, (v1 % (((v1 << ((827) & 63))) | 1)), (v1 - (v1 - (v1 | (v1 + 1))))}
    var v20 int8 = 42
    v21 := GenS2{F0: 498, F1: ((true || false) && (3.14159 >= ((0.1 - (0.0 / 1.0)) * (0.0 + (0.0 - (float64((-2.25 - (0.1 / 3.14159))) - 1e10)))))), F2: 255}
    var v22 int64 = 611
    for fi23 := 0; fi23 < 18; fi23++ {
        if (1 != -2147483648) {
            if (1 <= 42) {
                v24 := GenS3{F0: 2, F1: fi23}
                _ = len(v3)
                _ = v24
            }
            if true {
                v25 := GenS1{F0: false}
                var v26 int64 = v22
                v27 := GenS3{F0: 608, F1: (v1 << ((0) & 63))}
                _ = v25
                _ = v26
                _ = v27
            }
        } else {
            v2 = v2
        }
        v28 := GenS4{F0: 1}
        _ = len(v19)
        _ = fi23
        _ = v28
    }
    var v29 uint = 65535
    var v30 int32 = 1
    fmt.Println(v18)
    var v31 float64 = 1.5
    v32 := GenS1{F0: (1.5 != -2.25)}
    if true {
        var v33 uint8 = 0
        var v34 int = v1
        v3[(v34 + int(v22))] = (((v2 + "\n\t") + v2) + (v2 + v2))
        _ = v33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
