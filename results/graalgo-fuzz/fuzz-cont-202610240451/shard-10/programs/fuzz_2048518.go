package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 int16; F1 uint16; F2 int8 }
type GenS7 struct { F0 int8; F1 int32; F2 int64 }
type GenS8 struct { F0 uint; F1 uint8 }
type GenS9 struct { F0 int; F1 int8 }

func main() {
    var v1 uint32 = 0
    var v2 uint16 = 42
    var v3 string = "a"
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]int)
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 42}
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 10; fi9++ {
            if ((!(false || (true || (0 < 9223372036854775807)))) && (113 != 102)) {
                v10 := v5[v3]
                var v11 int16 = 22
                _ = v10
                _ = v11
            }
            v3 = (v3 + v3)
            v1 = v1
            _ = fi9
        }
        _ = fi8
    }
    if true {
        if (!false) {
            for fi12 := 0; fi12 < 13; fi12++ {
                fmt.Println(v1, fi12, fi12)
                for fi13 := 0; fi13 < 19; fi13++ {
                    v14 := GenS2{F0: 101, F1: 0}
                    v15 := GenS3{F0: (v2 & (v2 / ((v2) | 1))), F1: 4294967295}
                    switch "foo bar" {
                    case "a":
                        v3 = ("foo bar" + (v3 + v3))
                    case "hello":
                        var v16 int16 = 42
                        _ = v16
                    default:
                        v17 := GenS4{F0: true}
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                switch fi12 {
                case -241:
                    _ = len(v5)
                case -842:
                    fmt.Println(v1)
                }
                _ = fi12
            }
            v18 := GenS5{F0: 1}
            var v19 float64 = 3.14159
            _ = v18
            _ = v19
        }
    }
    v20 := GenS6{F0: 329, F1: v2, F2: 100}
    var v21 int = 42
    v1 = 2147483647
    var v22 float32 = (((0.0 - (0.1 / (float32((-2.25 - 1.5)) / 3.14159))) * (1.5 - 1.5)) / (0.1 / 0.1))
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 bool = false
        var v25 float32 = v22
        _ = fi23
        _ = v24
        _ = v25
    }
    if true {
        if true {
            var v26 uint16 = uint16(v1)
            v27 := GenS7{F0: 0, F1: 540, F2: 129}
            if false {
                v3 = ((v3 + v3) + v3)
                v28 := GenS8{F0: 1, F1: 0}
                _ = v28
            }
            _ = v26
            _ = v27
        }
        var v29 uint64 = 2
        for fi30 := 0; fi30 < 8; fi30++ {
            if ((v3 + "") >= (v3 + v3)) {
                fmt.Println(v22, v29)
            } else {
                switch v21 {
                case 127:
                    v31 := v5[(v3 + v3)]
                    _ = v31
                case 743:
                    v21 = v21
                default:
                    v32 := GenS9{F0: v21, F1: 127}
                    _ = v32
                }
                fmt.Println(fi30, v21, v21)
            }
            var v33 string = ""
            v34 := []int{(v21 | (v21 + fi30)), (fi30 + (^(fi30)))}
            _ = fi30
            _ = v33
            _ = v34
        }
        _ = v29
    }
    var v35 float64 = (float64(v22) - ((1e10 + (((0.0 / 3.14159) + 1.5) - ((1e10 + 1e10) / 1.5))) * 0.1))
    fmt.Println(v2, v21)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
