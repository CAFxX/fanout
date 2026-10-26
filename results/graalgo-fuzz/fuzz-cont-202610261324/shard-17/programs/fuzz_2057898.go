package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool; F1 uint64; F2 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 int64; F2 int64 }
func genFn0() int32 {
    var v4 int8 = 25
    if (((true && true) || true) && true) {
        v4 += (v4 - v4)
        if true {
            var v5 int16 = 1
            _ = v5
        }
        fmt.Println(v4)
    }
    switch "a" {
    case "foo bar":
        if false {
            v6 := GenS1{F0: 0, F1: 42}
            v4 -= v4
            v7 := GenS1{F0: 262, F1: (^(v4))}
            _ = v6
            _ = v7
        }
    }
    v8 := GenS2{F0: 0, F1: 0}
    _ = v4
    _ = v8
    return 42
}
func genFn1(p0 int, p1 uint16, p2 uint) int32 {
    switch ("\n\t" + "hello") {
    case "foo bar":
        fmt.Println(p2)
    case "\n\t":
        v9 := []int{1, (p0 / (((p0 % (((p0 / ((p0) | 1))) | 1))) | 1)), (p0 * (p0 * p0))}
        _ = v9
    case "hello":
        p2 -= (^(p2))
    }
    v10 := GenS3{F0: 996}
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 7
}

func main() {
    var v1 bool = (-2147483648 == -2147483648)
    var v2 uint16 = 505
    var v3 uint32 = 65535
    v11 := GenS3{F0: v3}
    v1 = (("\n\t" + "foo bar") == "hello")
    for fi12 := 0; fi12 < 4; fi12++ {
        v2 -= (v2 >> ((2) & 63))
        switch (fi12 + -175) {
        case -615:
            var v13 int64 = -883
            _ = v13
        case 9223372036854775807:
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS4{F0: v1, F1: 2, F2: fi12}
                v2 += (v2 << ((2) & 63))
                _ = fi14
                _ = v15
            }
        case -9223372036854775808:
            var v16 string = ""
            _ = v16
        }
        var v17 uint = 499
        _ = fi12
        _ = v17
    }
    for fi18 := 0; fi18 < 14; fi18++ {
        v3 = (^(v3))
        _ = fi18
    }
    var v19 uint = 41
    if v1 {
        v2 -= (v2 >> ((360) & 63))
    }
    var v20 int8 = 1
    v20 = (v20 + (v20 - v20))
    v21 := []int{32767}
    v21 = append(v21, -413)
    v22 := GenS5{F0: -13}
    if ((^(v3)) < (v3 | (v3 >> (((v19 * (v19 >> ((v19) & 63)))) & 63)))) {
        v23 := GenS6{F0: 18446744073709551615, F1: 42, F2: -274}
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 bool = (((v2 << ((0) & 63)) < 255) || ((1 != 0) || (32767 <= 0)))
            var v26 int = (fi24 ^ fi24)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = v23
    } else {
        var v27 uint8 = 128
        _ = v27
    }
    var v28 bool = false
    if len(v21) > 0 {
        v21[0] = -9223372036854775808
    }
    v21 = append(v21, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
}
