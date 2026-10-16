package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 int32; F1 int; F2 string }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 int; F2 uint32 }
func genFn0() int16 {
    var v5 uint32 = 14
    v6 := GenS1{F0: 255, F1: 4294967295}
    var v7 uint = 280
    v8 := GenS1{F0: 40, F1: 4294967295}
    switch "hello" {
    case "\n\t":
        switch "hello" {
        case "a":
            var v9 uint16 = 42
            _ = v9
        case "\n\t":
            for fi10 := 0; fi10 < 3; fi10++ {
                v5 -= (^(v5))
                _ = fi10
            }
        }
    case "":
        v11 := []int{}
        _ = v11
    case "foo bar":
        v12 := []int{0, 32768, 707, -284}
        _ = v12
    }
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint8 = 255
    v2 := []int{2147483647, 2147483647, 4294967295}
    var v3 uint = 42
    var v4 string = ((("" + "\n\t") + "foo bar") + ("" + "foo bar"))
    switch -450 {
    case 106:
        fmt.Println(v1)
    case 42:
        v3 += v3
    case -572:
        if true {
            if true {
                for fi13 := 0; fi13 < 6; fi13++ {
                    var v14 float64 = 1.5
                    if (!false) {
                        v15 := GenS2{F0: 42, F1: (fi13 ^ fi13), F2: v4}
                        v16 := GenS2{F0: 969, F1: -779, F2: v4}
                        fmt.Println(fi13)
                        _ = v15
                        _ = v16
                    }
                    fmt.Println(v1, v1)
                    _ = fi13
                    _ = v14
                }
                var v17 string = v4
                v18 := GenS2{F0: -952, F1: 0, F2: ""}
                _ = v17
                _ = v18
            } else {
                v3 -= (v3 / (((v3 + 1)) | 1))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        } else {
            if true {
                var v19 float32 = 1e10
                _ = v19
            } else {
                for fi20 := 0; fi20 < 12; fi20++ {
                    fmt.Println(fi20, v3, v3)
                    var v21 float64 = float64((float32(1e10) * 1e10))
                    v22 := GenS3{F0: v1, F1: fi20}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint64 = 0
                _ = v23
            }
        }
    default:
        for fi24 := 0; fi24 < 8; fi24++ {
            if (float32(float64(1.5)) > float32(3.14159)) {
                var v25 int = fi24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 uint = v3
                    var v28 int64 = 0
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                v29 := GenS1{F0: 65535, F1: 2}
                _ = v29
            }
            var v30 uint32 = 4294967295
            _ = fi24
            _ = v30
        }
    }
    v31 := GenS4{F0: 42, F1: uint8(v3)}
    v1 = v1
    switch v4 {
    case "\n\t":
        v32 := GenS5{F0: 2147483647}
        _ = v32
    default:
        v1 = (v1 ^ v1)
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v1)
        var v34 uint64 = 1
        _ = fi33
        _ = v34
    }
    v35 := GenS6{F0: 9223372036854775807, F1: 587, F2: 0}
    var v36 uint64 = 0
    var v37 uint32 = 0
    if len(v2) > 0 {
        v2[0] = int(v37)
    }
    _ = len(v2)
    v36 += (v36 % ((v36) | 1))
    v3 = (v3 * (v3 + uint(v1)))
    var v38 bool = (9223372036854775807 == -1)
    var v39 uint8 = v1
    for fi40 := 0; fi40 < 14; fi40++ {
        var v41 int = fi40
        var v42 uint = 758
        _ = fi40
        _ = v41
        _ = v42
    }
    var v43 bool = v38
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v43
}
