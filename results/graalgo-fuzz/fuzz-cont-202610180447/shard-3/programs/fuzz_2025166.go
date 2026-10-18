package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint; F1 string; F2 bool }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint) int64 {
    switch ("" + "a") {
    case "a":
        fmt.Println(p0)
    }
    var v5 float32 = -2.25
    v6 := GenS1{F0: 42, F1: 0}
    p0 = p0
    _ = p0
    _ = v5
    _ = v6
    return -1
}

func main() {
    var v1 int32 = 1
    var v2 uint32 = 234
    var v3 string = "foo bar"
    var v4 int16 = 1
    v4 = (v4 + v4)
    if (false && true) {
        v7 := GenS1{F0: (v2 % ((2147483647) | 1)), F1: 2}
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                var v10 int32 = 10
                _ = fi9
                _ = v10
            }
            if false {
                for fi11 := 0; fi11 < 7; fi11++ {
                    v3 = ("" + "hello")
                    fmt.Println(v1)
                    _ = fi11
                }
                v12 := GenS2{F0: 1, F1: 102, F2: v4}
                for fi13 := 0; fi13 < 2; fi13++ {
                    v14 := GenS3{F0: 222, F1: ((((v3 + v3) + (v3 + v3)) + (v3 + ("foo bar" + v3))) + "a"), F2: ((0.1 / float32(float64(3.14159))) == 0.1)}
                    _ = fi13
                    _ = v14
                }
                _ = v12
            } else {
                v3 = (v3 + v3)
                var v15 int64 = -9223372036854775808
                _ = v15
            }
            _ = fi8
        }
        _ = v7
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := []int{fi16, fi16}
        v4 = 1000
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = fi16
        _ = v17
    }
    var v18 uint32 = v2
    v19 := []int{}
    fmt.Println(v1)
    v20 := GenS4{F0: v1}
    v2 -= (v2 >> ((255) & 63))
    if (2 == 42) {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 int64 = 0
            var v23 uint16 = 924
            _ = fi21
            _ = v22
            _ = v23
        }
    } else {
        v24 := GenS1{F0: (v18 << ((0) & 63)), F1: 163}
        for fi25 := 0; fi25 < 13; fi25++ {
            v18 -= 42
            var v26 uint = 414
            for fi27 := 0; fi27 < 12; fi27++ {
                fmt.Println(v2, v3)
                v19 = append(v19, (-(fi25)))
                if (1.5 <= 1e10) {
                    v3 = v3
                    v4 += 867
                    if false {
                        fmt.Println(fi27, v26)
                    } else {
                        v28 := GenS4{F0: v1}
                        v1 -= (^(v1))
                        v18 += v2
                        _ = v28
                    }
                }
                var v29 int = fi27
                _ = fi27
                _ = v29
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
