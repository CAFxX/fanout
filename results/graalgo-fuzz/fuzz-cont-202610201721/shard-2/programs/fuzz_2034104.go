package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 uint; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint) int64 {
    if true {
        if true {
            p0 = 65535
            p0 = (p0 ^ p0)
        }
    } else {
        v7 := []int{476, -9223372036854775808}
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    fmt.Println(p0)
    p0 = p0
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 float32 = ((3.14159 - 0.0) / 1.0)
        p0 -= (p0 * p0)
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint64 = 42
        p0 -= 957
        v13 -= 638
        _ = fi12
        _ = v13
    }
    _ = p0
    return -9223372036854775808
}
func genFn1(p0 int8) uint32 {
    var v14 int = 57
    switch (("\n\t" + "a") + "") {
    case "hello":
        fmt.Println(v14, p0)
    default:
        var v15 int8 = p0
        _ = v15
    }
    var v16 uint32 = 2147483647
    v17 := []int{(-(v14)), 1, (v14 | (v14 - -2147483648))}
    p0 -= (p0 & p0)
    _ = p0
    _ = v14
    _ = v16
    _ = v17
    return v16
}

func main() {
    var v1 uint16 = 1
    var v2 int16 = 42
    v3 := []int{-9223372036854775808, -9223372036854775808, 4294967295, -1}
    var v4 uint64 = 2
    v5 := make(map[uint32]string)
    var v6 uint8 = 255
    v18 := GenS2{F0: (^(v4)), F1: 42, F2: 0}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v5)
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 10; fi20++ {
            fmt.Println(fi19)
            switch "hello" {
            case "hello":
                switch "foo bar" {
                case "":
                    v4 -= 4294967295
                case "\n\t":
                    v3 = append(v3, 4294967295)
                }
            case "a":
                if (!true) {
                    fmt.Println(fi19, fi20, v6)
                    switch fi20 {
                    case -210:
                        if (true || true) {
                            delete(v5, 2147483648)
                        }
                    case 42:
                        v3 = append(v3, fi19)
                    case -68:
                        v21 := v5[312]
                        _ = v21
                    default:
                        v1 = (v1 ^ (^(v1)))
                    }
                    var v22 int32 = 1
                    _ = v22
                } else {
                    v2 = 299
                    var v23 bool = (!true)
                    fmt.Println(v23, fi20, v1)
                    _ = v23
                }
            default:
                v4 = (v4 | 1)
            }
            _ = fi20
        }
        var v24 uint64 = 630
        v2 += v2
        _ = fi19
        _ = v24
    }
    v6 = (^(v6))
    for fi25 := 0; fi25 < 7; fi25++ {
        v1 = v1
        _ = fi25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        v1 = v1
        v6 = (v6 * v6)
        _ = fi26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS3{F0: 809, F1: 42}
        var v29 float32 = float32((3.14159 * (1.5 + -2.25)))
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := GenS4{F0: 2147483647, F1: int32(v6)}
    for fi31 := 0; fi31 < 19; fi31++ {
        v1 += (^(v1))
        v6 = v6
        _ = fi31
    }
    var v32 uint32 = genFn1(100)
    if false {
        v4 = 564
        fmt.Println(v1)
    }
    if ((true || true) || ((true && false) && (1 == 0))) {
        for fi33 := 0; fi33 < 3; fi33++ {
            switch int(v32) {
            case -387:
                if ((false && true) || false) {
                    v34 := v5[uint32(v4)]
                    _ = v34
                }
            case 4294967295:
                _ = len(v5)
            }
            _ = fi33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
