package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 int8) uint {
    p0 += 1
    p0 = (p0 - (p0 | p0))
    switch 32767 {
    case 128:
        fmt.Println(p1, p1)
    case 947:
        v5 := GenS1{F0: (!("foo bar" > "")), F1: 298, F2: 1}
        _ = v5
    case 4294967295:
        v6 := GenS1{F0: false, F1: 18446744073709551615, F2: 42}
        _ = v6
    }
    _ = p0
    _ = p1
    return 255
}
func genFn1(p0 int16, p1 uint16, p2 string) uint16 {
    v7 := []int{-1, 9223372036854775807, 4294967295, -1}
    v8 := GenS2{F0: 0}
    var v9 uint64 = 42
    switch (p2 + p2) {
    case "\n\t":
        fmt.Println(p0, p0, v9)
    case "hello":
        var v10 int64 = 9223372036854775807
        _ = v10
    case "a":
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 453
}

func main() {
    var v1 int = 271
    v2 := []uint32{0, 42, 4294967295}
    v3 := make(map[int]int)
    var v4 uint8 = 213
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(v4, v1)
        var v12 float32 = 1.5
        _ = fi11
        _ = v12
    }
    var v13 int64 = 1
    v13 = -1
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint32 = 2147483648
        var v16 uint = 2
        _ = fi14
        _ = v15
        _ = v16
    }
    v1 += v1
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 float64 = -2.25
        _ = fi17
        _ = v18
    }
    var v19 string = ("foo bar" + "hello")
    v1 = v1
    fmt.Println(v19, v1)
    for fi20 := 0; fi20 < 3; fi20++ {
        _ = len(v2)
        if false {
            var v21 uint8 = v4
            for fi22 := 0; fi22 < 3; fi22++ {
                switch fi20 {
                case 0:
                    var v23 uint = 910
                    _ = v23
                default:
                    switch v19 {
                    case "a":
                        switch v19 {
                        case "foo bar":
                            v24 := GenS3{F0: 1, F1: fi20, F2: 776}
                            _ = v24
                        }
                    }
                }
                v4 = v4
                _ = fi22
            }
            _ = v21
        }
        _ = fi20
    }
    var v25 uint8 = 2
    for fi26 := 0; fi26 < 14; fi26++ {
        for fi27 := 0; fi27 < 5; fi27++ {
            var v28 int64 = -1
            fmt.Println(v25)
            switch fi27 {
            case -402:
                var v29 bool = false
                _ = v29
            default:
                if (42 != 18446744073709551615) {
                    v13 = v13
                    var v30 uint8 = (v25 ^ (^(v25)))
                    _ = v30
                } else {
                    var v31 uint32 = 42
                    v32 := GenS3{F0: 1, F1: (fi27 << ((genFn0(-2147483648, -119)) & 63)), F2: v31}
                    v33 := GenS3{F0: 127, F1: fi26, F2: 445}
                    _ = v31
                    _ = v32
                    _ = v33
                }
            }
            _ = fi27
            _ = v28
        }
        fmt.Println(v4, v1, fi26)
        _ = fi26
    }
    var v34 uint16 = 255
    for fi35 := 0; fi35 < 10; fi35++ {
        v4 = (v4 % ((v4) | 1))
        _ = fi35
    }
    var v36 float64 = (3.14159 + 1e10)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v19
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v36
    _ = v4
}
