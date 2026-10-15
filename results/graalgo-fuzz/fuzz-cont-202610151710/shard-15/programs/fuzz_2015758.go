package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int8, p1 string, p2 int16) int32 {
    p2 -= p2
    v5 := GenS1{F0: (p0 << ((828) & 63))}
    var v6 int8 = 127
    var v7 uint8 = 127
    p2 = (p2 / (((p2 / (((-(p2))) | 1))) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -5
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    var v4 int16 = -1
    v8 := GenS2{F0: 127, F1: (v4 + 32767), F2: v4}
    var v9 uint64 = 4294967295
    var v10 uint16 = 1
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int32 = 873
        switch -1 {
        case 127:
            var v13 uint = 2
            _ = v13
        case 0:
            v12 += v12
        case 9223372036854775807:
            if false {
                v10 = v10
                if true {
                    v4 -= 0
                    fmt.Println(v12, fi11)
                }
            }
        default:
            for fi14 := 0; fi14 < 11; fi14++ {
                delete(v1, 4294967295)
                _ = fi14
            }
        }
        _ = fi11
        _ = v12
    }
    var v15 uint16 = (v10 - (^(v10)))
    if (false || false) {
        var v16 float32 = (3.14159 - -2.25)
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 string = "foo bar"
            _ = fi17
            _ = v18
        }
        _ = v16
    }
    v15 -= (^(v15))
    if true {
        v19 := []int{243}
        _ = v19
    } else {
        v10 += v15
        v10 = uint16(v9)
        var v20 uint16 = (v15 - 334)
        _ = v20
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        var v22 string = "foo bar"
        _ = fi21
        _ = v22
    }
    if (!false) {
        v1[793] = 65535
        v23 := GenS3{F0: 865}
        var v24 int64 = 431
        _ = v23
        _ = v24
    } else {
        v25 := GenS1{F0: -128}
        var v26 uint32 = 0
        if true {
            for fi27 := 0; fi27 < 8; fi27++ {
                fmt.Println(fi27, v15)
                var v28 uint16 = 1
                _ = fi27
                _ = v28
            }
            v9 = v9
            var v29 float32 = -2.25
            _ = v29
        }
        _ = v25
        _ = v26
    }
    if ((1.5 / 1e10) != 0.0) {
        var v30 uint = 757
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
