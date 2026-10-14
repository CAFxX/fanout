package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64; F2 int8 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int }
type GenS4 struct { F0 int16; F1 bool; F2 string }
func genFn0(p0 int, p1 bool, p2 string) int64 {
    if (!(42 <= 127)) {
        switch p2 {
        case "\n\t":
            var v4 int16 = -1
            _ = v4
        case "hello":
            v5 := []int{(^(p0)), -372}
            _ = v5
        case "":
            v6 := []int{}
            _ = v6
        }
    }
    fmt.Println(p1)
    p1 = (!p1)
    if true {
        v7 := []int{9223372036854775807}
        var v8 uint8 = 42
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 223
}
func genFn1(p0 int32, p1 int8, p2 uint) int64 {
    var v9 float32 = 0.1
    v10 := GenS1{F0: 69, F1: 0, F2: (p1 ^ (p1 | 80))}
    v9 += (((float32(1e10) * (v9 / (v9 / v9))) / 1.0) + 1e10)
    fmt.Println(p1, p1)
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(p2, fi11, p1)
        if true {
            v9 = v9
            var v12 uint = p2
            _ = v12
        } else {
            var v13 int8 = p1
            v9 = v9
            _ = v13
        }
        v14 := []int{}
        _ = fi11
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return -222
}

func main() {
    v1 := make(map[int]bool)
    v2 := []uint32{2147483647, 1, 42}
    var v3 int32 = -75
    if true {
        if false {
            if false {
                fmt.Println(v3)
                fmt.Println(v3)
                fmt.Println(v3)
            } else {
                for fi15 := 0; fi15 < 11; fi15++ {
                    var v16 uint32 = 1
                    _ = fi15
                    _ = v16
                }
            }
        } else {
            v3 -= v3
            v3 = -727
            var v17 float64 = 1e10
            _ = v17
        }
        fmt.Println(v3)
    } else {
        switch "" {
        case "":
            var v18 int8 = 127
            _ = v18
        case "\n\t":
            fmt.Println(v3)
        case "a":
            v19 := v1[228]
            _ = v19
        }
        fmt.Println(v3)
        var v20 uint16 = 65535
        _ = v20
    }
    var v21 uint16 = 42
    var v22 uint32 = 4294967295
    var v23 int = -498
    v24 := GenS2{F0: "hello", F1: v3, F2: 42}
    switch "" {
    case "a":
        if false {
            for fi25 := 0; fi25 < 17; fi25++ {
                v26 := GenS1{F0: 42, F1: 1, F2: 39}
                _ = fi25
                _ = v26
            }
            v27 := GenS1{F0: 42, F1: -167, F2: 127}
            _ = len(v1)
            _ = v27
        } else {
            fmt.Println(v3, v21)
            var v28 uint64 = 4294967295
            _ = v28
        }
    default:
        for fi29 := 0; fi29 < 3; fi29++ {
            fmt.Println(v21, v3)
            for fi30 := 0; fi30 < 11; fi30++ {
                if true {
                    delete(v1, fi30)
                }
                _ = fi30
            }
            if true {
                switch ("\n\t" + "") {
                case "\n\t":
                    v31 := GenS1{F0: 255, F1: 42, F2: -15}
                    _ = v31
                case "a":
                    if len(v2) > 0 {
                        v2[0] = v22
                    }
                default:
                    v22 = 2147483648
                }
                if ((float32(0.1) + 0.1) > 1.5) {
                    v32 := GenS3{F0: true, F1: 0, F2: (v23 + (-(v23)))}
                    var v33 uint = 818
                    switch fi29 {
                    case 128:
                        v23 = v23
                    default:
                        v34 := GenS2{F0: "\n\t", F1: v3, F2: 127}
                        _ = v34
                    }
                    _ = v32
                    _ = v33
                }
            }
            _ = fi29
        }
    }
    v35 := GenS4{F0: 1000, F1: ((v23 & v23) != (v23 & v23)), F2: ""}
    v23 -= (^(v23))
    v36 := GenS2{F0: "\n\t", F1: (v3 * int32(v21)), F2: int8(v23)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v35
    _ = v36
}
