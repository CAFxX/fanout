package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32; F1 bool; F2 int32 }
type GenS3 struct { F0 int16; F1 int16; F2 bool }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint16 }
type GenS5 struct { F0 string; F1 uint8 }

func main() {
    var v1 int8 = -1
    v2 := []string{("a" + ("hello" + ("" + ("" + "hello")))), "", "\n\t"}
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    v5 := make(map[string]uint32)
    fmt.Println(v1)
    if (!(-1 != 42)) {
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi6, fi6)
            v7 := GenS1{F0: 519, F1: (4294967295 == 2147483648)}
            var v8 bool = (!false)
            _ = fi6
            _ = v7
            _ = v8
        }
        v1 = v1
        v1 = v1
    } else {
        _ = len(v3)
    }
    v9 := GenS1{F0: -228, F1: (true || false)}
    v1 += (v1 >> ((18446744073709551615) & 63))
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((!false) && (!(!false))) {
            var v11 float64 = float64(-2.25)
            _ = v11
        } else {
            v12 := GenS1{F0: 0, F1: false}
            v2 = append(v2, "\n\t")
            v13 := GenS2{F0: -51, F1: (!(255 == 2)), F2: 806}
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    v14 := GenS1{F0: 2147483647, F1: false}
    v15 := GenS3{F0: -1, F1: -830, F2: true}
    v1 = (^(v1))
    var v16 uint64 = 495
    var v17 int16 = 32767
    var v18 int64 = 956
    v1 += v1
    var v19 int8 = (v1 - v1)
    v17 = (v17 | v17)
    v18 -= (v18 >> ((42) & 63))
    v17 = v17
    v19 = v1
    for fi20 := 0; fi20 < 2; fi20++ {
        switch "\n\t" {
        case "foo bar":
            var v21 string = "a"
            _ = v21
        case "":
            v22 := GenS4{F0: 1, F1: 1, F2: 796}
            _ = v22
        default:
            v23 := GenS5{F0: "\n\t", F1: 196}
            _ = v23
        }
        v16 += (^(v16))
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
