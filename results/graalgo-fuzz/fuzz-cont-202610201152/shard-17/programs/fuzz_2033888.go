package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 int }
type GenS3 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 uint16, p1 string) int16 {
    p1 = "foo bar"
    var v4 int = 384
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 int32) string {
    if true {
        v5 := []int{118}
        var v6 float64 = 0.1
        var v7 float64 = v6
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := GenS1{F0: -128, F1: 0, F2: 255}
    for fi9 := 0; fi9 < 15; fi9++ {
        switch (fi9 * (^(fi9))) {
        case 553:
            var v10 uint16 = 671
            _ = v10
        default:
            var v11 uint64 = 0
            _ = v11
        }
        fmt.Println(fi9, fi9)
        var v12 uint64 = 0
        _ = fi9
        _ = v12
    }
    _ = p0
    _ = v8
    return "a"
}

func main() {
    var v1 uint8 = 170
    v2 := make(map[string]string)
    var v3 string = ""
    v13 := v2[v3]
    var v14 uint64 = 4294967295
    v15 := GenS1{F0: 1, F1: -800, F2: (v1 * (v1 << ((65535) & 63)))}
    if true {
        var v16 int32 = -150
        _ = v16
    } else {
        for fi17 := 0; fi17 < 15; fi17++ {
            fmt.Println(v3, v3, v1)
            _ = fi17
        }
        var v18 uint = 795
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int32 = 0
            for fi21 := 0; fi21 < 7; fi21++ {
                delete(v2, v3)
                v22 := []int{(fi19 * (fi19 | 4294967295)), (fi19 * fi19), int(v20)}
                v23 := GenS1{F0: 71, F1: 850, F2: 99}
                _ = fi21
                _ = v22
                _ = v23
            }
            fmt.Println(v13, v20, v14)
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    var v24 int32 = 0
    var v25 float64 = -2.25
    var v26 uint16 = 449
    switch 42 {
    case -2147483648:
        for fi27 := 0; fi27 < 7; fi27++ {
            _ = len(v2)
            if true {
                v25 = (0.0 + v25)
                if (42 < v24) {
                    v28 := GenS2{F0: ((v3 + v13) + v13), F1: 124, F2: -803}
                    _ = v28
                }
                var v29 int16 = -32768
                _ = v29
            }
            v30 := GenS2{F0: v3, F1: 42, F2: int(v26)}
            _ = fi27
            _ = v30
        }
    case -741:
        if (42 > -500) {
            fmt.Println(v14, v26, v25)
            var v31 uint32 = 4294967295
            _ = v31
        } else {
            switch v13 {
            case "hello":
                var v32 int64 = int64(v1)
                _ = v32
            }
            v33 := GenS3{F0: 32767, F1: -149, F2: -164}
            v34 := []int{-2147483648, int(v26), 2147483647, 836}
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
