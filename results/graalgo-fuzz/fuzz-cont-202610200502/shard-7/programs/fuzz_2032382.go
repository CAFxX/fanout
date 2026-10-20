package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint; F1 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint64, p1 uint8) uint64 {
    var v5 uint = 523
    v5 = v5
    for fi6 := 0; fi6 < 17; fi6++ {
        v5 -= (^(v5))
        var v7 int32 = 2147483647
        var v8 uint16 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    return (^(p0))
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -9223372036854775808
    var v3 uint64 = 42
    v4 := make(map[int]uint32)
    var v9 string = ""
    v10 := GenS1{F0: -128, F1: 128, F2: 9}
    for fi11 := 0; fi11 < 12; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v2 += 1
            var v13 float32 = 1e10
            v9 = (v9 + ("a" + ""))
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        for fi15 := 0; fi15 < 8; fi15++ {
            var v16 int64 = v2
            v17 := GenS1{F0: 100, F1: 42, F2: 814}
            var v18 string = v9
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            if false {
                v9 = v9
                v20 := GenS1{F0: 103, F1: 167, F2: 1}
                var v21 int64 = 9223372036854775807
                _ = v20
                _ = v21
            } else {
                switch v9 {
                case "a":
                    v9 = "a"
                }
                var v22 uint = 4294967295
                var v23 uint = (v22 & (v22 - 4294967295))
                _ = v22
                _ = v23
            }
            v24 := GenS1{F0: 0, F1: 1, F2: 1}
            _ = fi19
            _ = v24
        }
        _ = fi14
    }
    v2 += (v2 / ((94) | 1))
    for fi25 := 0; fi25 < 16; fi25++ {
        var v26 float32 = float32(((0.0 / 1e10) + 0.0))
        _ = fi25
        _ = v26
    }
    var v27 uint64 = (v3 ^ (^(v3)))
    v3 = genFn0((v27 % (((v27 / (((v27 & v27)) | 1))) | 1)), 205)
    v9 = (v9 + "")
    for fi28 := 0; fi28 < 15; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            v30 := GenS2{F0: 97, F1: 32767}
            _ = fi29
            _ = v30
        }
        v31 := GenS3{F0: 2147483648}
        var v32 uint8 = 255
        _ = fi28
        _ = v31
        _ = v32
    }
    var v33 int16 = 0
    if (!true) {
        var v34 int16 = (v33 + 1000)
        v9 = v9
        _ = v34
    }
    v35 := GenS2{F0: 346, F1: (v33 & v33)}
    v36 := GenS2{F0: 255, F1: v33}
    if false {
        var v37 float32 = 3.14159
        fmt.Println(v3, v2)
        _ = v37
    } else {
        var v38 float64 = 3.14159
        _ = v38
    }
    var v39 uint32 = 2147483648
    v27 += v27
    var v40 string = ("" + ((v9 + v9) + (v9 + v9)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v9
}
