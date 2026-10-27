package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint; F1 int16; F2 uint16 }
func genFn0(p0 string, p1 string) int8 {
    var v5 bool = true
    var v6 int = -1
    for fi7 := 0; fi7 < 2; fi7++ {
        if ((159 <= 42) && v5) {
            v8 := []int{(v6 - (v6 & v6)), v6, (v6 << ((633) & 63)), (fi7 * fi7)}
            _ = v8
        }
        _ = fi7
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 100
}
func genFn1(p0 uint, p1 int16) bool {
    var v9 uint64 = 0
    v10 := GenS1{F0: uint(v9), F1: 42, F2: 127}
    v9 = v9
    for fi11 := 0; fi11 < 6; fi11++ {
        v12 := []int{-473, fi11, fi11}
        var v13 int16 = 1
        v14 := GenS2{F0: (^(v9))}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v9 = 527
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (0 > 127)
}

func main() {
    var v1 float64 = -2.25
    v2 := make(map[uint32]bool)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v15 string = ("\n\t" + "foo bar")
    v16 := GenS3{F0: 4294967295, F1: 308}
    v1 -= 0.0
    switch ((v15 + v15) + v15) {
    case "":
        if (!true) {
            var v17 int = 4294967295
            _ = v17
        } else {
            delete(v4, 836)
            switch (v15 + "a") {
            case "hello":
                if true {
                    var v18 int64 = 1
                    switch 952 {
                    case -318:
                        var v19 uint8 = 113
                        _ = v19
                    case -2147483648:
                        v15 = "a"
                    case 128:
                        v20 := GenS4{F0: v15}
                        _ = v20
                    }
                    _ = v18
                }
            case "foo bar":
                v21 := []int{354, 128, 42}
                _ = v21
            case "a":
                _ = len(v4)
            }
            for fi22 := 0; fi22 < 15; fi22++ {
                v4[985] = v15
                _ = fi22
            }
        }
    case "a":
        fmt.Println(v1, v15)
    default:
        v23 := []int{4294967295, 9223372036854775807, 4294967295}
        _ = v23
    }
    var v24 float64 = v1
    for fi25 := 0; fi25 < 3; fi25++ {
        v24 -= float64(3.14159)
        for fi26 := 0; fi26 < 3; fi26++ {
            v1 -= v1
            v24 += v1
            v24 = v24
            var v27 uint = 49
            _ = fi26
            _ = v27
        }
        v28 := GenS5{F0: -1}
        _ = fi25
        _ = v28
    }
    var v29 int64 = -854
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int16 = 42
        _ = fi30
        _ = v31
    }
    if false {
        v1 -= v1
    }
    v32 := GenS6{F0: 65535, F1: -617, F2: 2}
    for fi33 := 0; fi33 < 14; fi33++ {
        fmt.Println(v15, v15)
        v1 -= (-2.25 * float64((-2.25 - (((0.1 / (3.14159 / 1.0)) / -2.25) - 0.0))))
        _ = fi33
    }
    if true {
        switch "\n\t" {
        case "":
            v34 := GenS4{F0: v15}
            _ = v34
        default:
            if (677 == 4294967295) {
                v15 = v15
            } else {
                v35 := []int{}
                delete(v2, 0)
                _ = v35
            }
        }
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v32
    _ = v4
}
