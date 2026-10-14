package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 string }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 uint = 42
    var v2 int8 = -1
    var v3 string = "\n\t"
    fmt.Println(v3)
    for fi4 := 0; fi4 < 20; fi4++ {
        var v5 float64 = (0.0 / ((1.5 + (1e10 + (-2.25 - (-2.25 * 1e10)))) - ((-2.25 + (1.5 / 1.5)) + (1.5 - 3.14159))))
        if (152 >= 255) {
            if (1 == 42) {
                if (!false) {
                    v6 := []int{int(v2)}
                    v7 := GenS1{F0: 888, F1: 65535}
                    v2 += v2
                    _ = v6
                    _ = v7
                }
            } else {
                v8 := GenS2{F0: 1000, F1: int16(v1), F2: v3}
                v9 := []int{42}
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 uint8 = 128
                    if (v1 >= v1) {
                        _ = len(v9)
                    } else {
                        fmt.Println(v11, v1)
                        v12 := GenS2{F0: 42, F1: 42, F2: ("\n\t" + ((v3 + (v3 + v3)) + v3))}
                        v1 = v1
                        _ = v12
                    }
                    _ = fi10
                    _ = v11
                }
                _ = v8
                _ = v9
            }
        } else {
            v1 = 65535
            switch 32768 {
            case 0:
                if ((1.5 < 1e10) && true) {
                    v2 -= 1
                } else {
                    var v13 int = 407
                    _ = v13
                }
            case 4294967295:
                var v14 int32 = 0
                _ = v14
            default:
                var v15 int64 = 358
                _ = v15
            }
            v16 := GenS3{F0: (v1 & v1), F1: -1}
            _ = v16
        }
        _ = fi4
        _ = v5
    }
    if true {
        v3 = (("foo bar" + v3) + v3)
        fmt.Println(v3)
    } else {
        v2 += (v2 ^ (v2 & v2))
        v1 -= (v1 | 4294967295)
    }
    v3 = v3
    var v17 uint8 = 142
    v18 := []int{}
    var v19 uint32 = 65535
    v2 += v2
    var v20 float32 = (0.0 * 0.1)
    var v21 uint32 = (v19 + 118)
    if false {
        v1 += (v1 % (((v1 ^ v1)) | 1))
    }
    fmt.Println(v21, v3, v19)
    v2 = v2
    v1 += v1
    v20 = 0.0
    for fi22 := 0; fi22 < 18; fi22++ {
        if true {
            v17 += (v17 >> ((521) & 63))
        }
        switch fi22 {
        case 42:
            v23 := GenS2{F0: 42, F1: 42, F2: v3}
            _ = v23
        }
        _ = fi22
    }
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
