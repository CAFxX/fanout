package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 int8; F2 bool }
func genFn0(p0 uint32, p1 int16) uint {
    if (false && false) {
        fmt.Println(p1, p0)
    }
    var v7 uint16 = 32767
    p1 -= p1
    _ = p0
    _ = p1
    _ = v7
    return uint(p0)
}

func main() {
    var v1 float64 = 1.5
    var v2 uint16 = 42
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    v6 := make(map[int]int)
    v2 = v2
    switch "foo bar" {
    case "":
        var v8 string = ("foo bar" + "\n\t")
        _ = v8
    case "foo bar":
        if false {
            var v9 uint = 4294967295
            var v10 float32 = 3.14159
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int32 = -2147483648
                var v13 int8 = 0
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
            _ = v9
        }
    case "\n\t":
        for fi14 := 0; fi14 < 13; fi14++ {
            if (true || (400 == 997)) {
                v15 := GenS1{F0: 2147483647, F1: 18446744073709551615}
                _ = v15
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                v17 := GenS2{F0: 0}
                _ = fi16
                _ = v17
            }
            _ = fi14
        }
    default:
        for fi18 := 0; fi18 < 8; fi18++ {
            var v19 string = "\n\t"
            _ = len(v5)
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: 127}
                v1 = float64((3.14159 / 1e10))
                fmt.Println(v2, v19, fi20)
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    fmt.Println(v1, v1)
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := v4[879]
        var v24 string = "hello"
        var v25 uint8 = 42
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println(v2, v1)
    switch "foo bar" {
    case "a":
        if false {
            switch "a" {
            case "\n\t":
                v26 := GenS3{F0: -2147483648, F1: -1, F2: (792 < uint(v2))}
                _ = v26
            }
            switch "hello" {
            case "hello":
                for fi27 := 0; fi27 < 20; fi27++ {
                    var v28 uint16 = v2
                    var v29 int16 = 0
                    _ = fi27
                    _ = v28
                    _ = v29
                }
            default:
                v2 = v2
            }
            v2 = 255
        }
    default:
        switch "" {
        case "\n\t":
            var v30 uint32 = 1
            _ = v30
        case "foo bar":
            var v31 int32 = -16
            _ = v31
        default:
            v32 := GenS3{F0: 32768, F1: -128, F2: false}
            _ = v32
        }
    }
    var v33 float64 = 1.5
    if (!false) {
        var v34 float32 = 0.1
        _ = v34
    } else {
        for fi35 := 0; fi35 < 13; fi35++ {
            v2 -= 42
            _ = fi35
        }
        fmt.Println(v1)
        if (235 != 4294967295) {
            var v36 uint16 = (^(v2))
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
