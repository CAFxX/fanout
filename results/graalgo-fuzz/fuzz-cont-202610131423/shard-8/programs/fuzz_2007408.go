package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint32 }
type GenS4 struct { F0 int }
func genFn0() uint {
    v6 := []int{-407}
    v7 := GenS1{F0: (!false), F1: 1000}
    var v8 uint16 = 2
    v9 := GenS2{F0: 461, F1: 0}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 255
}
func genFn1(p0 string) int32 {
    var v10 uint64 = 616
    fmt.Println(p0)
    fmt.Println(v10, v10)
    for fi11 := 0; fi11 < 13; fi11++ {
        v12 := []int{(fi11 | (fi11 - -479))}
        var v13 int16 = 1000
        v14 := GenS3{F0: 0, F1: 2147483647, F2: 133}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = v10
    return -617
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 1
    var v3 int64 = -1
    var v4 uint16 = 32767
    v5 := make(map[int]uint32)
    var v15 int32 = 42
    for fi16 := 0; fi16 < 20; fi16++ {
        v17 := GenS1{F0: true, F1: (v2 | (v2 << ((42) & 63)))}
        if false {
            var v18 bool = true
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v2, v3)
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := GenS4{F0: 127}
            fmt.Println(v4, fi19)
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v4 = v4
    if false {
        var v22 int8 = -22
        _ = v22
    }
    for fi23 := 0; fi23 < 4; fi23++ {
        if (int(v4) < fi23) {
            var v24 int32 = (v15 << ((180) & 63))
            _ = v24
        } else {
            v25 := GenS1{F0: true, F1: 32767}
            _ = v25
        }
        v26 := []int{(-(fi23)), (^(fi23)), fi23, (fi23 * (^(fi23)))}
        v2 = -1
        _ = fi23
        _ = v26
    }
    v27 := GenS1{F0: (442 > -2147483648), F1: v2}
    v28 := []int{-1, -1}
    switch 32768 {
    case -601:
        for fi29 := 0; fi29 < 3; fi29++ {
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 uint16 = (v4 >> ((18) & 63))
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    case 32767:
        var v32 float32 = 1e10
        _ = v32
    case -9223372036854775808:
        var v33 float64 = (0.1 / (3.14159 + (1e10 + 1.5)))
        _ = v33
    }
    fmt.Println(v4, v2)
    v34 := GenS4{F0: 832}
    var v35 float32 = (float32(-2.25) - (float32(-2.25) / 3.14159))
    v36 := GenS1{F0: true, F1: 620}
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        var v37 int = -9223372036854775808
        _ = v37
    case "":
        v38 := GenS3{F0: (v15 << ((42) & 63)), F1: v15, F2: 2}
        _ = v38
    default:
        var v39 float64 = 3.14159
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
