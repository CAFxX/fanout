package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 uint; F2 int8 }
func genFn0(p0 string) uint8 {
    v6 := []int{-645}
    v7 := GenS1{F0: 127}
    switch 127 {
    case -2147483648:
        if (42 > 2) {
            if false {
                p0 = p0
            }
            _ = len(v6)
        } else {
            switch -9223372036854775808 {
            case 892:
                v8 := GenS1{F0: 255}
                _ = v8
            case 32768:
                v9 := GenS2{F0: 2}
                _ = v9
            case 4294967295:
                _ = len(v6)
            }
        }
    case 128:
        switch "" {
        case "foo bar":
            fmt.Println(p0)
        case "a":
            var v10 uint32 = 4294967295
            _ = v10
        default:
            v11 := GenS1{F0: 120}
            _ = v11
        }
    default:
        var v12 int32 = -517
        _ = v12
    }
    fmt.Println(p0)
    p0 = p0
    _ = p0
    _ = v6
    _ = v7
    return 206
}

func main() {
    v1 := make(map[string]string)
    var v2 int64 = 0
    v3 := []int{128}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]string)
    fmt.Println(v2)
    for fi13 := 0; fi13 < 5; fi13++ {
        v14 := GenS2{F0: 1}
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v15 := GenS2{F0: 0}
        _ = fi13
        _ = v14
        _ = v15
    }
    if (4294967295 < 121) {
        fmt.Println(v2)
        if false {
            v2 = v2
        } else {
            fmt.Println(v2)
            fmt.Println(v2)
            fmt.Println(v2)
        }
    } else {
        if (0.0 == (0.1 / 1e10)) {
            fmt.Println(v2)
        }
    }
    switch 0 {
    case 128:
        v16 := v4[2147483648]
        _ = v16
    }
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 int8 = 100
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := GenS3{F0: fi19, F1: 255, F2: v18}
            var v21 int16 = 86
            v21 = 1000
            _ = fi19
            _ = v20
            _ = v21
        }
        if true {
            v22 := GenS1{F0: 0}
            for fi23 := 0; fi23 < 18; fi23++ {
                var v24 uint8 = 127
                _ = fi23
                _ = v24
            }
            for fi25 := 0; fi25 < 4; fi25++ {
                v2 = v2
                var v26 float32 = (-2.25 - ((float32(0.0) / 3.14159) + -2.25))
                _ = fi25
                _ = v26
            }
            _ = v22
        } else {
            v27 := GenS2{F0: 65535}
            _ = v27
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v2)
    delete(v5, 2147483647)
    v2 = (v2 % ((v2) | 1))
    v28 := v5[2]
    switch -2147483648 {
    case 40:
        v2 = v2
    case 128:
        switch 127 {
        case -9223372036854775808:
            v29 := GenS1{F0: 239}
            _ = v29
        case 588:
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 uint = 984
                v32 := GenS2{F0: uint16(v31)}
                v2 = (v2 << ((v31) & 63))
                _ = fi30
                _ = v31
                _ = v32
            }
        }
    default:
        if (v2 > -1) {
            if len(v3) > 0 {
                v3[0] = 32768
            }
            v2 += v2
        }
    }
    var v33 string = v28
    var v34 int16 = -495
    var v35 uint8 = 215
    var v36 float64 = 0.0
    fmt.Println(v34)
    var v37 uint64 = 42
    for fi38 := 0; fi38 < 3; fi38++ {
        v39 := GenS2{F0: 65535}
        _ = fi38
        _ = v39
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
