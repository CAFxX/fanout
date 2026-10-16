package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int64 = -1
    var v2 float32 = ((3.14159 / (0.1 * float32(3.14159))) * (1e10 + (0.1 * float32(1e10))))
    var v3 int8 = -43
    var v4 uint16 = 403
    fmt.Println(v2)
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 int16 = 1000
        _ = fi5
        _ = v6
    }
    if (!true) {
        switch "hello" {
        case "a":
            switch -1 {
            case 32768:
                var v7 int64 = 1
                _ = v7
            default:
                var v8 uint8 = 253
                _ = v8
            }
        }
        if true {
            v9 := []int{}
            var v10 float32 = float32(float64(float32((-2.25 + 1.5))))
            v3 = -65
            _ = v10
            _ = v9
        } else {
            switch "a" {
            case "hello":
                var v11 uint32 = 343
                _ = v11
            case "a":
                v2 = ((1.5 - -2.25) + ((v2 + v2) - (1.5 + (v2 - v2))))
            case "":
                var v12 uint16 = v4
                _ = v12
            }
            if true {
                if (42 < -80) {
                    for fi13 := 0; fi13 < 7; fi13++ {
                        var v14 uint32 = 2147483647
                        _ = fi13
                        _ = v14
                    }
                } else {
                    v15 := GenS1{F0: v4, F1: 4294967295}
                    v16 := []int{-2147483648}
                    _ = v15
                    _ = v16
                }
                switch 4294967295 {
                case 9223372036854775807:
                    var v17 uint32 = 2147483648
                    _ = v17
                }
                v3 = 36
            } else {
                v1 = v1
                v18 := []int{-904}
                _ = v18
            }
        }
    }
    v19 := GenS2{F0: v4, F1: v3}
    v2 = float32(3.14159)
    v20 := GenS3{F0: 2, F1: (^(v1))}
    v21 := []int{42}
    var v22 bool = (!true)
    v23 := GenS1{F0: v4, F1: 874}
    var v24 int32 = 318
    var v25 bool = true
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v26 := GenS1{F0: 32767, F1: 487}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
