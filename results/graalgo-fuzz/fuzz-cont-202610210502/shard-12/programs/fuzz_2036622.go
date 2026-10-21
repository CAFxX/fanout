package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 int64 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]int)
    var v3 uint = 1
    v4 := make(map[string]int)
    for fi5 := 0; fi5 < 8; fi5++ {
        v3 -= (^(v3))
        _ = fi5
    }
    switch ("\n\t" + "a") {
    case "":
        if (((true || true) && true) || true) {
            v3 += (v3 / ((4294967295) | 1))
            for fi6 := 0; fi6 < 9; fi6++ {
                v3 = v3
                v7 := []int{}
                var v8 int8 = -1
                _ = fi6
                _ = v7
                _ = v8
            }
        }
    case "hello":
        v3 += (v3 >> ((v3) & 63))
    default:
        for fi9 := 0; fi9 < 19; fi9++ {
            var v10 int16 = -422
            v11 := v1[(fi9 ^ (-(fi9)))]
            var v12 float32 = 0.0
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
    }
    if (!(32768 < 4294967295)) {
        var v13 int16 = -32768
        _ = v13
    }
    v3 -= v3
    if false {
        v14 := GenS1{F0: 0, F1: 42}
        v3 = (v3 & 18446744073709551615)
        _ = v14
    } else {
        v3 += 315
        for fi15 := 0; fi15 < 13; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v17 := GenS2{F0: 100}
                v18 := []int{fi16}
                var v19 uint = (v3 + v3)
                if true {
                    fmt.Println(v19, v19)
                }
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
    }
    var v20 float64 = -2.25
    v21 := GenS3{F0: 2147483647, F1: 0}
    var v22 uint64 = 40
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
