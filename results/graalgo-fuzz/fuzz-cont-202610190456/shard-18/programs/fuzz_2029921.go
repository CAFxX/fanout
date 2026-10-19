package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    v1 := []int{42, 128, 32768}
    var v2 int8 = -17
    v3 := make(map[string]string)
    var v4 float64 = float64((0.1 / -2.25))
    var v5 int16 = -669
    v6 := make(map[uint32]uint32)
    v7 := GenS1{F0: 65535}
    fmt.Println(v5, v2, v5)
    if false {
        fmt.Println(v2, v5)
        if true {
            v4 = (float64(1.5) * v4)
            var v8 int64 = -539
            _ = v8
        } else {
            v9 := v3["\n\t"]
            if false {
                for fi10 := 0; fi10 < 3; fi10++ {
                    var v11 int64 = 1
                    v12 := GenS2{F0: 42, F1: v11}
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = len(v1)
            }
            _ = v9
        }
    }
    if false {
        if false {
            var v13 int32 = 455
            delete(v3, ((("\n\t" + "") + "\n\t") + "hello"))
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS2{F0: 65535, F1: -1}
                v5 += v5
                var v16 int = fi14
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = v13
        } else {
            var v17 float32 = (0.1 - (0.1 + 1.5))
            v17 += v17
            if (v17 >= 1e10) {
                for fi18 := 0; fi18 < 11; fi18++ {
                    var v19 int32 = int32(v2)
                    if true {
                        var v20 uint32 = 42
                        v21 := GenS2{F0: 2, F1: 9223372036854775807}
                        _ = v20
                        _ = v21
                    } else {
                        var v22 string = ((("foo bar" + "\n\t") + "\n\t") + "hello")
                        _ = v22
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = v17
        }
        switch 127 {
        case 9223372036854775807:
            for fi23 := 0; fi23 < 15; fi23++ {
                delete(v6, 359)
                v5 = 0
                _ = fi23
            }
        case 32768:
            v24 := GenS2{F0: 42, F1: -762}
            _ = v24
        }
        var v25 float32 = 1.5
        _ = v25
    }
    _ = len(v1)
    v5 += (v5 - (v5 ^ (v5 / ((v5) | 1))))
    _ = len(v1)
    v26 := GenS3{F0: 215, F1: 1}
    var v27 bool = true
    v27 = ((!v27) || (!(0.1 > v4)))
    if len(v1) > 0 {
        v1[0] = -9223372036854775808
    }
    v28 := GenS1{F0: 255}
    fmt.Println(v27)
    if true {
        v27 = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
