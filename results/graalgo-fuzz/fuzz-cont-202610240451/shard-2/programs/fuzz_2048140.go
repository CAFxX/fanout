package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 uint16 }
func genFn0(p0 uint32, p1 int16, p2 string) int32 {
    if true {
        if false {
            v6 := []int{42, 4294967295}
            _ = v6
        } else {
            var v7 uint = uint(p0)
            _ = v7
        }
        p2 = (p2 + (p2 + "foo bar"))
    }
    v8 := []int{42, 344, -2147483648, -272}
    if (!true) {
        if false {
            v8 = append(v8, 32768)
            var v9 bool = false
            v9 = v9
            _ = v9
        } else {
            fmt.Println(p2, p2, p1)
            var v10 uint16 = 2
            var v11 float64 = 0.1
            _ = v10
            _ = v11
        }
        var v12 int = 9223372036854775807
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]bool)
    var v3 uint = 2
    v4 := []int{-841, -703, 32768, -753}
    v5 := make(map[int]uint32)
    if len(v4) > 0 {
        v4[0] = 1
    }
    v3 = (v3 >> (((^(v3))) & 63))
    _ = len(v4)
    v13 := GenS1{F0: "hello"}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 bool = (0.1 >= -2.25)
    v14 = ((!false) || (1 > 4294967295))
    v14 = (true && false)
    fmt.Println(v3)
    v3 = v3
    var v15 uint64 = 552
    for fi16 := 0; fi16 < 11; fi16++ {
        v14 = v14
        if (-1 > 1) {
            _ = len(v2)
            if v14 {
                var v17 int = 32767
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(fi16, fi16)
                    v17 -= v17
                    var v19 float32 = 3.14159
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            for fi20 := 0; fi20 < 9; fi20++ {
                switch 32768 {
                case -435:
                    delete(v5, (fi16 + (fi16 % ((562) | 1))))
                case 127:
                    if (0.0 <= -2.25) {
                        v21 := GenS2{F0: -1, F1: v3, F2: 31}
                        _ = len(v4)
                        _ = v21
                    }
                }
                var v22 uint16 = 317
                v15 += v15
                _ = fi20
                _ = v22
            }
        }
        v3 = (^(v3))
        _ = fi16
    }
    var v23 int32 = 0
    v14 = v14
    var v24 int64 = -800
    v25 := GenS3{F0: 1, F1: 1, F2: 42}
    v23 = (v23 ^ (-(v23)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
