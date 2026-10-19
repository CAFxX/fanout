package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 int; F2 int8 }
type GenS3 struct { F0 int8 }
func genFn0() int16 {
    var v5 uint = 42
    v5 -= v5
    _ = v5
    return -49
}
func genFn1(p0 bool) uint8 {
    p0 = true
    for fi6 := 0; fi6 < 20; fi6++ {
        if (!p0) {
            var v7 int16 = -32768
            _ = v7
        }
        v8 := []int{(fi6 & (fi6 | (fi6 >> ((42) & 63)))), (^(fi6))}
        if (false || ((-(fi6)) == 875)) {
            for fi9 := 0; fi9 < 3; fi9++ {
                if (!p0) {
                    p0 = (65535 > 42)
                }
                var v10 string = "\n\t"
                v8 = append(v8, fi6)
                _ = fi9
                _ = v10
            }
        }
        _ = fi6
        _ = v8
    }
    var v11 int32 = 217
    var v12 float64 = ((3.14159 + (0.0 * -2.25)) * 0.1)
    for fi13 := 0; fi13 < 13; fi13++ {
        var v14 uint = 0
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = v11
    _ = v12
    return 127
}

func main() {
    var v1 uint16 = 0
    v2 := make(map[string]bool)
    var v3 uint16 = v1
    v4 := make(map[uint32]bool)
    v15 := GenS1{F0: -572, F1: 398}
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint16 = (^(v3))
        var v18 uint16 = (v17 / (((^(v17))) | 1))
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 string = "foo bar"
    fmt.Println(v19)
    _ = len(v4)
    for fi20 := 0; fi20 < 8; fi20++ {
        var v21 int = -9223372036854775808
        v22 := GenS2{F0: 0, F1: (v21 >> ((728) & 63)), F2: -128}
        var v23 uint8 = 68
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: v3, F1: 128, F2: 37}
    v19 = v19
    v19 = v19
    if false {
        v25 := []int{1, 4294967295, 127}
        var v26 int32 = -2147483648
        fmt.Println(v1)
        _ = v25
        _ = v26
    } else {
        var v27 uint64 = 4294967295
        for fi28 := 0; fi28 < 13; fi28++ {
            switch 4294967295 {
            case -1:
                if (!true) {
                    v29 := GenS1{F0: fi28, F1: 1}
                    v30 := GenS3{F0: 0}
                    v19 = "\n\t"
                    _ = v29
                    _ = v30
                } else {
                    if false {
                        v31 := GenS1{F0: 1, F1: 59}
                        _ = v31
                    }
                }
            }
            v32 := []int{1, fi28}
            var v33 int64 = 1
            _ = fi28
            _ = v32
            _ = v33
        }
        if (2 <= 861) {
            var v34 uint32 = 283
            if (true || true) {
                v35 := GenS3{F0: 100}
                v36 := GenS2{F0: 2, F1: -9223372036854775808, F2: 100}
                _ = v35
                _ = v36
            }
            _ = v34
        }
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
