package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int; F2 int8 }
func genFn0(p0 uint16) int64 {
    var v7 float32 = 0.1
    v8 := GenS1{F0: -1}
    v7 = (v7 + v7)
    _ = p0
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1(p0 int64) bool {
    var v9 bool = true
    switch ("foo bar" + ("\n\t" + "")) {
    case "":
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 float64 = 0.0
            _ = fi10
            _ = v11
        }
    default:
        p0 -= (p0 * 9223372036854775807)
    }
    _ = p0
    _ = v9
    return (42 != 1)
}

func main() {
    var v1 float32 = -2.25
    v2 := make(map[uint32]string)
    var v3 uint16 = 0
    v4 := []int{42, 32767}
    v5 := []uint32{0}
    v6 := []string{"foo bar"}
    v3 = (v3 >> ((893) & 63))
    v12 := GenS1{F0: -2147483648}
    var v13 int64 = 42
    var v14 uint32 = 889
    v14 = (v14 + 42)
    if (v14 == v14) {
        v13 += -1
    }
    delete(v2, 65535)
    switch 768 {
    case -2147483648:
        fmt.Println(v3, v13)
    case -9223372036854775808:
        v15 := GenS2{F0: 42, F1: 2}
        _ = v15
    case 32768:
        v16 := GenS3{F0: 254, F1: 976, F2: 100}
        _ = v16
    default:
        v17 := v2[(v14 & v14)]
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v2[(v14 + v14)] = ("foo bar" + "")
        fmt.Println(v1, v1, fi18)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi18
    }
    switch ("" + "hello") {
    case "a":
        switch 1 {
        case 32768:
            if false {
                var v19 string = ("a" + "hello")
                v5 = append(v5, v14)
                var v20 float64 = 0.0
                _ = v19
                _ = v20
            } else {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v21 int16 = 1000
                var v22 uint8 = 42
                _ = v21
                _ = v22
            }
        case -833:
            v1 = (((v1 / (v1 + (float32(1.5) - v1))) - v1) - v1)
        case 32767:
            var v23 bool = false
            _ = v23
        default:
            var v24 int32 = 42
            _ = v24
        }
    default:
        v3 = v3
    }
    v25 := GenS3{F0: 54, F1: 216, F2: -1}
    v14 = v14
    switch 9223372036854775807 {
    case -1:
        var v26 uint8 = 178
        _ = v26
    }
    fmt.Println(v13, v13)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
