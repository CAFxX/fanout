package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 string; F2 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint; F2 uint8 }
func genFn0(p0 uint32, p1 string) int16 {
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 1000}
    _ = p0
    _ = p1
    _ = v6
    return -250
}
func genFn1(p0 uint) uint {
    fmt.Println(p0)
    v7 := []int{-392, 74, -9223372036854775808, 127}
    if (-887 != 1) {
        switch "a" {
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        case "":
            v7 = append(v7, 127)
        case "hello":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            for fi8 := 0; fi8 < 10; fi8++ {
                v9 := GenS2{F0: 1, F1: (("foo bar" + "foo bar") + ("foo bar" + "hello")), F2: (p0 + (p0 >> ((p0) & 63)))}
                if true {
                    var v10 int8 = 0
                    v10 += (v10 & 100)
                    _ = v10
                } else {
                    var v11 int = fi8
                    v11 = v11
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        }
    }
    var v12 int32 = -1
    _ = p0
    _ = v12
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    var v3 uint64 = 0
    v4 := []string{"hello"}
    v5 := make(map[uint32]uint32)
    var v13 uint8 = 2
    v4 = append(v4, "foo bar")
    fmt.Println(v13)
    switch "a" {
    case "\n\t":
        v13 = (v13 ^ (v13 & v13))
    default:
        var v14 int16 = genFn0(885, "a")
        _ = v14
    }
    v15 := GenS3{F0: 0, F1: (v13 % ((1) | 1)), F2: 889}
    var v16 int32 = 2147483647
    v17 := GenS4{F0: (^(v3)), F1: v13, F2: 2}
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS4{F0: 801, F1: v13, F2: 745}
            var v21 int16 = -32768
            v13 = v13
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi18
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        switch ("a" + ("hello" + "\n\t")) {
        case "":
            fmt.Println(v13, v3)
        case "\n\t":
            v23 := GenS1{F0: int16(v13)}
            _ = v23
        case "foo bar":
            _ = len(v2)
        }
        _ = fi22
    }
    var v24 uint16 = 87
    var v25 uint = genFn1(genFn1(uint(v3)))
    v26 := GenS2{F0: (v13 * (v13 - (v13 * (v13 ^ v13)))), F1: "", F2: 65535}
    var v27 int8 = -1
    for fi28 := 0; fi28 < 16; fi28++ {
        v25 = v25
        _ = len(v4)
        _ = len(v4)
        _ = fi28
    }
    v29 := GenS5{F0: v24, F1: genFn1(651), F2: v13}
    var v30 int32 = (v16 >> ((uint(v24)) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
