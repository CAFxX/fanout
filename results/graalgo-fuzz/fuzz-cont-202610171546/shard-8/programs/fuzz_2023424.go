package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 int }
func genFn0(p0 int8, p1 uint32) int16 {
    var v4 uint32 = (p1 ^ (p1 + 65535))
    p0 = (p0 ^ (p0 - 0))
    for fi5 := 0; fi5 < 12; fi5++ {
        switch (-(fi5)) {
        case 42:
            v6 := []int{4294967295, fi5, (^(fi5))}
            _ = v6
        case 89:
            if true {
                p1 = p1
                fmt.Println(v4, fi5)
            } else {
                fmt.Println(p1, p0, p1)
                var v7 int32 = 1
                switch fi5 {
                case 980:
                    p0 = 1
                case 9223372036854775807:
                    var v8 uint16 = 255
                    _ = v8
                default:
                    if (!true) {
                        var v9 int16 = 895
                        _ = v9
                    } else {
                        v4 += (^(v4))
                        v10 := GenS1{F0: 53, F1: -99, F2: (-(fi5))}
                        if (v7 > v7) {
                            var v11 float32 = (1.5 - 1e10)
                            if (!(!false)) {
                                var v12 float32 = 0.1
                                var v13 float64 = 0.1
                                _ = v12
                                _ = v13
                            }
                            _ = v11
                        }
                        _ = v10
                    }
                }
                _ = v7
            }
        }
        _ = fi5
    }
    _ = p0
    _ = p1
    _ = v4
    return -716
}
func genFn1(p0 uint8, p1 int16, p2 uint64) uint8 {
    for fi14 := 0; fi14 < 7; fi14++ {
        p0 += p0
        v15 := []int{-29, (fi14 / (((fi14 | (-(fi14)))) | 1)), (fi14 >> ((902) & 63))}
        _ = fi14
        _ = v15
    }
    v16 := []int{32767, 0, 457}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return uint8(p2)
}

func main() {
    var v1 string = "\n\t"
    var v2 string = v1
    v3 := []uint32{}
    switch "hello" {
    case "\n\t":
        v1 = ""
    case "a":
        for fi17 := 0; fi17 < 3; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                v3 = append(v3, 2147483648)
                _ = fi18
            }
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint64 = 0
                _ = fi19
                _ = v20
            }
            _ = fi17
        }
    case "hello":
        for fi21 := 0; fi21 < 20; fi21++ {
            fmt.Println(fi21, fi21)
            _ = fi21
        }
    }
    v1 = "foo bar"
    for fi22 := 0; fi22 < 19; fi22++ {
        switch fi22 {
        case 127:
            var v23 uint32 = 142
            _ = v23
        }
        _ = fi22
    }
    var v24 int = 32767
    v2 = (v1 + v2)
    v2 = v2
    for fi25 := 0; fi25 < 5; fi25++ {
        fmt.Println(v2, fi25)
        var v26 bool = true
        _ = fi25
        _ = v26
    }
    if false {
        var v27 uint64 = 2
        _ = v27
    } else {
        v24 += (v24 ^ (^(v24)))
        fmt.Println(v24)
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    var v28 uint = 622
    switch (v24 % (((-(v24))) | 1)) {
    case 127:
        var v29 float64 = float64(float32(3.14159))
        _ = v29
    default:
        var v30 float64 = float64((0.0 + (0.0 / 1.0)))
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
}
