package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string }

func main() {
    var v1 uint = 766
    v2 := make(map[uint32]int)
    v3 := []uint32{2147483648, 925}
    v3 = append(v3, 1)
    var v4 uint64 = 2
    for fi5 := 0; fi5 < 13; fi5++ {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 bool = (42 > uint8(v1))
            _ = fi6
            _ = v7
        }
        if (-627 >= -952) {
            var v8 int16 = 681
            _ = v8
        } else {
            v4 -= (v4 - 18446744073709551615)
        }
        _ = fi5
    }
    v3 = append(v3, 59)
    v3 = append(v3, 2147483647)
    v9 := GenS1{F0: 127}
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v4)
        for fi11 := 0; fi11 < 10; fi11++ {
            v12 := GenS2{F0: "a", F1: (2147483647 <= -1), F2: 42}
            var v13 float32 = 0.0
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi10
    }
    switch "" {
    case "foo bar":
        var v14 int16 = 1000
        _ = v14
    case "\n\t":
        if (0 != -9223372036854775808) {
            v1 = 4294967295
            _ = len(v3)
            var v15 string = "a"
            _ = v15
        }
    default:
        var v16 int16 = 32767
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = v4
    v17 := GenS3{F0: true}
    switch "a" {
    case "":
        v4 = v4
    case "hello":
        for fi18 := 0; fi18 < 12; fi18++ {
            if (2 > 65535) {
                v4 -= (v4 >> (((v1 << ((v1) & 63))) & 63))
            } else {
                v2[42] = fi18
                v3 = append(v3, 2147483648)
            }
            if (205 >= int32(fi18)) {
                fmt.Println(fi18, v4, v1)
                v19 := GenS4{F0: "foo bar"}
                fmt.Println(v1)
                _ = v19
            } else {
                v1 = 0
                v20 := GenS4{F0: "hello"}
                _ = v20
            }
            _ = fi18
        }
    case "foo bar":
        v21 := v2[2147483648]
        _ = v21
    default:
        var v22 float64 = 0.1
        _ = v22
    }
    if (true && true) {
        fmt.Println(v1, v1)
    }
    if true {
        var v23 float32 = 1.5
        var v24 float64 = 3.14159
        v3 = append(v3, 1)
        _ = v23
        _ = v24
    } else {
        delete(v2, 574)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "hello":
        v25 := GenS4{F0: ("a" + "a")}
        _ = v25
    case "a":
        v4 -= 2
    }
    v4 += v4
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
