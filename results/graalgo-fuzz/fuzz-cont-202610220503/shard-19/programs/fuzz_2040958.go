package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 bool; F1 int8; F2 int }
type GenS4 struct { F0 string; F1 string }
func genFn0(p0 bool, p1 uint32) uint8 {
    v7 := []int{32768, 133, 42, 128}
    p0 = true
    p0 = (-1 >= 32767)
    p1 = 2
    _ = p0
    _ = p1
    _ = v7
    return 117
}

func main() {
    var v1 float32 = ((1.5 + 1.5) * float32((0.0 - 1e10)))
    v2 := make(map[uint32]bool)
    var v3 uint8 = 1
    var v4 string = "foo bar"
    var v5 uint = 1
    var v6 uint32 = 626
    if (!false) {
        for fi8 := 0; fi8 < 8; fi8++ {
            if false {
                v5 = v5
            }
            v6 = (v6 % (((v6 + (v6 * 4294967295))) | 1))
            _ = fi8
        }
        v6 -= v6
    } else {
        switch "foo bar" {
        case "\n\t":
            if true {
                v9 := v2[(v6 / ((2) | 1))]
                switch 127 {
                case 42:
                    var v10 uint64 = 0
                    _ = v10
                }
                v4 = ("foo bar" + v4)
                _ = v9
            }
        }
        v4 = v4
    }
    if true {
        v11 := GenS1{F0: 964}
        fmt.Println(v4)
        v12 := v2[v6]
        _ = v11
        _ = v12
    }
    v6 -= (^(v6))
    v6 = 1
    var v13 float64 = float64((float32((0.0 + (1.5 * 1.5))) + (v1 * v1)))
    v1 -= v1
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := GenS2{F0: int32(v5), F1: 42}
        _ = fi14
        _ = v15
    }
    var v16 int16 = -32768
    var v17 uint8 = v3
    var v18 bool = ((float64(-2.25) >= v13) || (!true))
    var v19 string = v4
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(fi20, v3, v17)
        _ = fi20
    }
    if (float32((v13 + v13)) != v1) {
        if false {
            v5 = v5
            if (-1 <= -55) {
                v13 = v13
            } else {
                v3 -= genFn0(true, v6)
                var v21 float64 = ((v13 - 0.1) / v13)
                _ = v21
            }
            var v22 uint8 = (v17 & (v17 >> ((255) & 63)))
            _ = v22
        } else {
            var v23 uint64 = 1
            delete(v2, v6)
            switch 127 {
            case 32767:
                v24 := GenS3{F0: v18, F1: 1, F2: 1}
                _ = v24
            }
            _ = v23
        }
        v25 := GenS4{F0: "\n\t", F1: ""}
        _ = v25
    } else {
        var v26 bool = true
        var v27 int8 = 100
        v3 = v17
        _ = v26
        _ = v27
    }
    v28 := GenS1{F0: 0}
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
