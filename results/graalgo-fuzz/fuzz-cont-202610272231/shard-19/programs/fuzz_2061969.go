package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint }
func genFn0() uint {
    v6 := GenS1{F0: 139, F1: -128}
    v7 := []int{-460, -2147483648, -2147483648, 0}
    switch 128 {
    case -554:
        switch "" {
        case "hello":
            var v8 uint32 = 564
            _ = v8
        default:
            var v9 uint64 = 18446744073709551615
            _ = v9
        }
    default:
        v10 := GenS1{F0: 123, F1: 85}
        _ = v10
    }
    var v11 int16 = -1
    var v12 string = "hello"
    _ = v11
    _ = v12
    _ = v6
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 751
    v4 := []int{738, -883}
    var v5 int16 = 0
    fmt.Println(v5)
    var v13 float32 = float32(1e10)
    fmt.Println(v13, v13, v2)
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v1)
        v15 := GenS2{F0: genFn0(), F1: v3}
        v5 -= 631
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        if (100 == 100) {
            v17 := GenS2{F0: 567, F1: v3}
            for fi18 := 0; fi18 < 16; fi18++ {
                fmt.Println(fi18, fi18, v2)
                _ = fi18
            }
            _ = v17
        } else {
            v13 += (3.14159 - v13)
            v1[(v3 + (v3 >> ((235) & 63)))] = "hello"
        }
        v2 = uint64(v3)
        v19 := GenS3{F0: -1}
        _ = fi16
        _ = v19
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 -= -32768
    var v20 uint8 = uint8(v3)
    var v21 uint16 = 0
    for fi22 := 0; fi22 < 3; fi22++ {
        switch (-(fi22)) {
        case 128:
            var v23 bool = true
            _ = v23
        case 2147483647:
            v20 += v20
        default:
            var v24 uint16 = 32767
            _ = v24
        }
        for fi25 := 0; fi25 < 9; fi25++ {
            v13 -= ((v13 + (v13 / float32(float64(v13)))) / v13)
            _ = fi25
        }
        var v26 uint16 = v21
        _ = fi22
        _ = v26
    }
    var v27 int64 = 883
    fmt.Println(v20, v3, v5)
    if true {
        v4 = append(v4, 2147483647)
        if false {
            var v28 uint64 = v2
            if (true && true) {
                v21 = v21
                fmt.Println(v5)
                if (260 == 0) {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        v3 = (v3 + v3)
                        v13 = float32(1.5)
                        _ = fi29
                    }
                }
            }
            for fi30 := 0; fi30 < 3; fi30++ {
                switch "foo bar" {
                case "\n\t":
                    var v31 float32 = (v13 - float32(-2.25))
                    _ = v31
                case "hello":
                    delete(v1, v3)
                case "a":
                    v5 = 1
                default:
                    var v32 float64 = 0.0
                    _ = v32
                }
                _ = fi30
            }
            _ = v28
        }
    }
    v33 := GenS4{F0: 65535}
    var v34 int32 = 86
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
