package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int8, p1 uint16) uint64 {
    v5 := GenS1{F0: ("hello" + "hello"), F1: 607, F2: 42}
    if (!false) {
        fmt.Println(p1, p1)
    }
    switch -498 {
    case 32768:
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{fi6, (-(fi6)), (^(fi6))}
            var v8 uint16 = 577
            _ = fi6
            _ = v7
            _ = v8
        }
    case 696:
        v9 := []int{32768}
        _ = v9
    }
    var v10 uint = 4294967295
    p0 = 46
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return 0
}
func genFn1(p0 int64, p1 bool) uint64 {
    var v11 float64 = -2.25
    p1 = (737 > 2)
    _ = p0
    _ = p1
    _ = v11
    return 4294967295
}

func main() {
    var v1 int8 = 100
    var v2 uint8 = 213
    var v3 int16 = 0
    var v4 float64 = 0.1
    var v12 int8 = -36
    v13 := GenS2{F0: ("foo bar" + ""), F1: v1}
    var v14 uint = 4294967295
    if true {
        v1 += (v1 / (((-(v12))) | 1))
        v1 += (v1 / ((v12) | 1))
        if false {
            var v15 float32 = -2.25
            for fi16 := 0; fi16 < 20; fi16++ {
                v12 += (v12 + (v1 % ((v1) | 1)))
                v3 += (^(v3))
                var v17 uint16 = 32767
                _ = fi16
                _ = v17
            }
            _ = v15
        }
    } else {
        for fi18 := 0; fi18 < 5; fi18++ {
            v12 = (v1 - -113)
            v19 := GenS3{F0: v1, F1: 255}
            _ = fi18
            _ = v19
        }
    }
    v2 = v2
    var v20 string = "a"
    var v21 int32 = 394
    v22 := []int{}
    v23 := GenS2{F0: "", F1: (-(v12))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
