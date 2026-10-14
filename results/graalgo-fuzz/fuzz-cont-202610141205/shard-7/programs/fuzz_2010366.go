package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 uint16; F1 uint32 }
type GenS6 struct { F0 int8 }
func genFn0(p0 uint8) int64 {
    v5 := []int{42}
    var v6 float32 = 0.1
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 uint64, p1 uint16, p2 int) uint16 {
    v7 := []int{}
    switch "\n\t" {
    case "foo bar":
        if (83 <= 0) {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            var v8 int64 = 1
            _ = v8
        } else {
            if (("foo bar" + "hello") < ("hello" + (("a" + ("" + "a")) + "hello"))) {
                for fi9 := 0; fi9 < 12; fi9++ {
                    var v10 int8 = -128
                    _ = fi9
                    _ = v10
                }
            }
        }
    default:
        for fi11 := 0; fi11 < 13; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: 160}
                var v14 uint8 = 48
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    v15 := GenS2{F0: 1, F1: 0}
    if len(v7) > 0 {
        v7[0] = (p2 % (((p2 << ((42) & 63))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v7
    return (p1 + p1)
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]uint32)
    var v3 uint64 = 767
    var v4 uint32 = 4294967295
    v3 = (v3 - (v3 << ((42) & 63)))
    var v16 uint32 = (^(v4))
    _ = len(v2)
    v17 := GenS3{F0: 994, F1: (v16 - 2147483648)}
    v4 = (v16 / ((v16) | 1))
    v3 = (^(v3))
    v4 = (v16 >> ((255) & 63))
    for fi18 := 0; fi18 < 15; fi18++ {
        _ = len(v2)
        if (!true) {
            fmt.Println(v3)
        }
        v16 -= 1
        _ = fi18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        v20 := GenS4{F0: 1, F1: 2147483647}
        for fi21 := 0; fi21 < 15; fi21++ {
            if false {
                v3 = v3
                v4 = (v4 * (^(v4)))
            }
            fmt.Println(v16, fi21, v4)
            v22 := GenS5{F0: 872, F1: (v16 % ((v4) | 1))}
            _ = fi21
            _ = v22
        }
        _ = fi19
        _ = v20
    }
    v23 := GenS5{F0: 48, F1: (v4 % ((0) | 1))}
    v24 := GenS2{F0: 360, F1: 0}
    v16 -= 0
    v3 = (v3 & v3)
    if (1 > uint8(v16)) {
        var v25 bool = true
        switch "foo bar" {
        case "foo bar":
            for fi26 := 0; fi26 < 19; fi26++ {
                if v25 {
                    var v27 uint = 334
                    v4 = v16
                    _ = v27
                } else {
                    if (!(230 < 18446744073709551615)) {
                        v25 = (-171 < -805)
                    } else {
                        for fi28 := 0; fi28 < 5; fi28++ {
                            v3 = v3
                            var v29 string = (("a" + "\n\t") + "")
                            _ = fi28
                            _ = v29
                        }
                        var v30 uint64 = (v3 / ((4294967295) | 1))
                        v2[fi26] = 2147483647
                        _ = v30
                    }
                    v2[fi26] = (^(v4))
                }
                _ = fi26
            }
        case "hello":
            switch ("foo bar" + "\n\t") {
            case "hello":
                v31 := GenS6{F0: 127}
                _ = v31
            default:
                var v32 int32 = 2147483647
                _ = v32
            }
        case "":
            v25 = (2147483648 == v16)
        default:
            v1 = 100
        }
        v33 := GenS5{F0: 744, F1: 0}
        _ = v25
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
