package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int8 = -128
    var v2 int32 = 42
    var v3 uint = 4294967295
    var v4 int = -253
    v5 := make(map[string]string)
    v6 := []int{-2147483648}
    switch "a" {
    case "foo bar":
        if (0 == 2) {
            v3 += (v3 % (((^(v3))) | 1))
        } else {
            v7 := GenS1{F0: v1}
            _ = v7
        }
    case "":
        v8 := GenS1{F0: (v1 ^ v1)}
        _ = v8
    case "hello":
        v6 = append(v6, (v4 + -9223372036854775808))
    default:
        if true {
            v9 := GenS1{F0: (v1 - (v1 * -1))}
            _ = v9
        }
    }
    v3 -= v3
    _ = len(v6)
    var v10 int16 = 42
    v11 := GenS2{F0: 714, F1: 127, F2: 2}
    var v12 string = ""
    v13 := GenS1{F0: 0}
    fmt.Println(v10, v10)
    switch ((v12 + "a") + "hello") {
    case "\n\t":
        v10 = 0
    default:
        if (v4 == 1) {
            switch ("\n\t" + v12) {
            case "a":
                v14 := GenS3{F0: 42, F1: (v1 / ((v1) | 1))}
                _ = v14
            case "\n\t":
                fmt.Println(v12, v12, v3)
            case "hello":
                fmt.Println(v4, v2, v1)
            }
            if (v12 > "") {
                var v15 uint8 = 42
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = v15
            }
        } else {
            v16 := GenS1{F0: v1}
            v5[v12] = "a"
            v12 = v12
            _ = v16
        }
    }
    var v17 uint8 = 86
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 += v1
        _ = fi18
    }
    v10 = v10
    var v19 uint16 = 42
    for fi20 := 0; fi20 < 4; fi20++ {
        v21 := GenS4{F0: 230}
        if (((!(v2 < (v2 >> (((v3 << (((v3 ^ (v3 - (v3 << ((18446744073709551615) & 63))))) & 63))) & 63)))) || true) && true) {
            for fi22 := 0; fi22 < 10; fi22++ {
                if true {
                    var v23 string = "foo bar"
                    var v24 int64 = 42
                    _ = v23
                    _ = v24
                }
                if (uint(v19) < v3) {
                    var v25 int = (v4 | fi20)
                    _ = v25
                } else {
                    v2 -= (v2 & -2147483648)
                }
                v6 = append(v6, (fi22 * 0))
                _ = fi22
            }
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
