package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 string; F2 int32 }

func main() {
    var v1 int = 4294967295
    var v2 uint32 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    var v5 uint8 = 0
    var v6 int = v1
    if true {
        for fi7 := 0; fi7 < 18; fi7++ {
            v8 := GenS1{F0: (((fi7 % ((966) | 1)) > 2147483647) || (false || (!true))), F1: 295, F2: 127}
            var v9 string = "\n\t"
            v5 = 93
            _ = fi7
            _ = v8
            _ = v9
        }
        v10 := GenS2{F0: -684, F1: ((1 <= 1) || (int(v5) < v1))}
        _ = v10
    }
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        var v11 float32 = (0.1 / 1.0)
        _ = v11
    default:
        v1 = 622
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 = fi12
        v5 += (^(v5))
        _ = fi12
    }
    for fi13 := 0; fi13 < 18; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint64 = 0
            var v16 int = (fi13 * (v1 | (fi14 * fi13)))
            switch v6 {
            case 128:
                v16 = (v16 >> ((255) & 63))
            case 0:
                var v17 int64 = -1
                _ = v17
            case -9223372036854775808:
                v18 := []int{(v16 * (fi13 - (fi14 >> ((0) & 63)))), v1, 2147483647, fi14}
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi13
    }
    v2 = (v2 % ((0) | 1))
    var v19 bool = false
    var v20 int8 = 100
    v21 := GenS3{F0: true, F1: 0, F2: 1}
    var v22 uint64 = 289
    v23 := []int{v1, v6, (v1 * v1), (^(v1))}
    v24 := GenS4{F0: int64(v22)}
    v25 := v4[980]
    for fi26 := 0; fi26 < 15; fi26++ {
        v25 = v25
        v19 = (0 >= -831)
        v2 -= (^(v2))
        _ = fi26
    }
    v27 := GenS5{F0: 0, F1: "\n\t", F2: -2147483648}
    if v19 {
        for fi28 := 0; fi28 < 2; fi28++ {
            var v29 uint32 = v2
            _ = fi28
            _ = v29
        }
        for fi30 := 0; fi30 < 16; fi30++ {
            var v31 float32 = 1.5
            var v32 bool = (3.14159 < 0.1)
            var v33 float64 = 1e10
            _ = fi30
            _ = v31
            _ = v32
            _ = v33
        }
        var v34 int32 = -2147483648
        _ = v34
    }
    if false {
        delete(v3, (v6 / ((int(v22)) | 1)))
        var v35 int8 = v20
        _ = len(v3)
        _ = v35
    } else {
        var v36 uint16 = 255
        var v37 int16 = 1
        for fi38 := 0; fi38 < 17; fi38++ {
            delete(v4, (v2 - v2))
            _ = fi38
        }
        _ = v36
        _ = v37
    }
    var v39 int8 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
