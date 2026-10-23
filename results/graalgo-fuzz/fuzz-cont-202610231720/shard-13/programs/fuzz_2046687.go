package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 int64 }

func main() {
    var v1 uint32 = 688
    var v2 string = ""
    v3 := []int{int(v1), 42, 128}
    var v4 float32 = (((3.14159 + 3.14159) + 0.0) * -2.25)
    if (1e10 <= -2.25) {
        v5 := GenS1{F0: 2, F1: ((1.5 < v4) && false), F2: 42}
        _ = v5
    } else {
        fmt.Println(v4)
    }
    v3 = append(v3, 1)
    if true {
        v6 := GenS1{F0: 0, F1: true, F2: 1}
        _ = v6
    } else {
        v4 += v4
        v4 += (1e10 + (v4 / float32(-2.25)))
    }
    v4 += (float32(-2.25) - v4)
    if true {
        v7 := GenS1{F0: 510, F1: (false && (((3.14159 - 1e10) + (-2.25 / float64(v4))) > float64(float32(float64((v4 / 1.5)))))), F2: 0}
        fmt.Println(v2)
        _ = v7
    }
    if false {
        v1 -= (^(v1))
        var v8 uint16 = 755
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 float32 = v4
        fmt.Println(fi9)
        _ = fi9
        _ = v10
    }
    if true {
        v2 = v2
        v1 += 2147483647
        if len(v3) > 0 {
            v3[0] = -9223372036854775808
        }
    } else {
        v11 := GenS1{F0: 4294967295, F1: (true || (!(!((0 < -253) && false)))), F2: -32768}
        var v12 bool = (2147483647 <= 2147483647)
        _ = v11
        _ = v12
    }
    v1 = v1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v13 int32 = -222
    v14 := GenS2{F0: 1, F1: 65535}
    v15 := GenS3{F0: (65535 <= 32767), F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
