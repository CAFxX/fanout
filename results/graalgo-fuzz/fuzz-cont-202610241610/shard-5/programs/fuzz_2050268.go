package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int }
type GenS3 struct { F0 uint64; F1 int64; F2 int64 }
type GenS4 struct { F0 int64; F1 int64; F2 int8 }
func genFn0() uint64 {
    var v6 string = "\n\t"
    v6 = "hello"
    v6 = "hello"
    var v7 uint64 = 42
    v8 := []int{1}
    _ = v6
    _ = v7
    _ = v8
    return (^(v7))
}
func genFn1() bool {
    var v9 int32 = 0
    var v10 uint64 = 18446744073709551615
    _ = v10
    _ = v9
    return (true && false)
}

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint16 = 255
    var v5 int16 = -682
    for fi11 := 0; fi11 < 9; fi11++ {
        delete(v3, fi11)
        if false {
            var v12 uint8 = 92
            v13 := []int{(fi11 ^ 32767)}
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        if (-512 < 1) {
            var v15 int16 = v5
            v16 := GenS1{F0: -157}
            v15 = v5
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 12; fi17++ {
                v4 = v4
                v18 := GenS2{F0: 20, F1: 0, F2: fi14}
                v19 := GenS3{F0: v1, F1: -1, F2: -678}
                _ = fi17
                _ = v18
                _ = v19
            }
            var v20 uint32 = 333
            v21 := GenS4{F0: 0, F1: -9223372036854775808, F2: -128}
            _ = v20
            _ = v21
        }
        _ = fi14
    }
    for fi22 := 0; fi22 < 4; fi22++ {
        var v23 float64 = ((1.5 - (1e10 * 1e10)) * 0.1)
        var v24 uint64 = (^(v1))
        v25 := GenS4{F0: 9223372036854775807, F1: 42, F2: 30}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    v26 := GenS3{F0: v1, F1: 740, F2: 257}
    _ = len(v3)
    var v27 int64 = 616
    for fi28 := 0; fi28 < 20; fi28++ {
        if (1 > 2147483647) {
            v1 -= v1
            var v29 uint = 18446744073709551615
            v30 := GenS4{F0: (v27 / ((173) | 1)), F1: (v27 % (((v27 << (((v29 * v29)) & 63))) | 1)), F2: 0}
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 19; fi31++ {
            switch fi28 {
            case -1:
                if (306 != 65535) {
                    v32 := GenS4{F0: v27, F1: (v27 & -1), F2: -85}
                    _ = v32
                }
            case -9223372036854775808:
                var v33 int64 = (v27 + v27)
                _ = v33
            case -455:
                v1 -= 4294967295
            default:
                v34 := GenS4{F0: -993, F1: -9223372036854775808, F2: 127}
                _ = v34
            }
            v35 := GenS4{F0: (v27 << ((687) & 63)), F1: 42, F2: 5}
            _ = fi31
            _ = v35
        }
        var v36 int16 = (v5 / (((v5 & 314)) | 1))
        _ = fi28
        _ = v36
    }
    switch 127 {
    case 127:
        for fi37 := 0; fi37 < 19; fi37++ {
            delete(v2, 42)
            for fi38 := 0; fi38 < 9; fi38++ {
                fmt.Println(fi38, v27)
                _ = fi38
            }
            v39 := []int{(^(fi37)), (fi37 - fi37), 4294967295}
            _ = fi37
            _ = v39
        }
    case 42:
        v40 := []int{-637}
        _ = v40
    default:
        var v41 bool = false
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
