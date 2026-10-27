package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 bool; F1 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 string; F2 int64 }
type GenS5 struct { F0 uint64; F1 uint; F2 int32 }
func genFn0() int32 {
    if (32767 == 255) {
        var v4 string = "foo bar"
        v5 := []int{32768}
        _ = v4
        _ = v5
    }
    switch 32768 {
    case -329:
        var v6 uint32 = 2147483648
        _ = v6
    case 42:
        if false {
            v7 := GenS1{F0: 128, F1: -1}
            var v8 uint64 = 4294967295
            _ = v7
            _ = v8
        } else {
            var v9 float64 = 0.1
            switch 42 {
            case 32768:
                var v10 int16 = -368
                _ = v10
            case 337:
                v11 := GenS1{F0: 252, F1: -87}
                _ = v11
            case 696:
                v9 = (1e10 / v9)
            }
            _ = v9
        }
    case 0:
        var v12 uint16 = 255
        _ = v12
    }
    var v13 int16 = -535
    for fi14 := 0; fi14 < 17; fi14++ {
        var v15 int32 = 183
        v16 := GenS1{F0: 253, F1: -29}
        _ = fi14
        _ = v15
        _ = v16
    }
    for fi17 := 0; fi17 < 3; fi17++ {
        v18 := GenS1{F0: 42, F1: 0}
        _ = fi17
        _ = v18
    }
    _ = v13
    return 0
}

func main() {
    var v1 uint16 = 2
    var v2 int = 32768
    var v3 int64 = -467
    if false {
        switch (v2 ^ v2) {
        case 1:
            for fi19 := 0; fi19 < 9; fi19++ {
                v3 += 761
                fmt.Println(v2, fi19, v1)
                v2 = (fi19 * v2)
                _ = fi19
            }
        }
        v20 := []int{(v2 >> ((770) & 63)), v2}
        if true {
            v21 := GenS2{F0: false, F1: 65535}
            _ = v21
        }
        _ = v20
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int32 = 984
            if true {
                v2 += int(v3)
                fmt.Println(fi23, v1, v1)
                v25 := GenS3{F0: 127}
                switch v2 {
                case 128:
                    v3 -= v3
                    var v26 int8 = 42
                    _ = v26
                case -1:
                    v2 = fi23
                case 127:
                    v1 = v1
                default:
                    v27 := []int{v2, (fi23 - (fi23 ^ 2147483647)), (v2 | fi22)}
                    v28 := GenS1{F0: 0, F1: 0}
                    if true {
                        var v29 int = fi22
                        fmt.Println(v29, v29)
                        _ = v29
                    }
                    _ = v27
                    _ = v28
                }
                _ = v25
            }
            _ = fi23
            _ = v24
        }
        var v30 uint64 = 0
        _ = fi22
        _ = v30
    }
    v31 := GenS4{F0: 1, F1: "foo bar", F2: 1}
    if false {
        v1 -= v1
        var v32 int64 = 42
        v32 += (v3 ^ v32)
        _ = v32
    }
    var v33 int = (^(v2))
    v34 := GenS5{F0: 356, F1: 0, F2: int32(v3)}
    var v35 int16 = 42
    v36 := []int{v2}
    v35 = (v35 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
}
