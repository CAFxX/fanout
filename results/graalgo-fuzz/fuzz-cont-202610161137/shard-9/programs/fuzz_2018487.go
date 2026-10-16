package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 string }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]uint32)
    var v3 int64 = -9223372036854775808
    var v4 bool = false
    v5 := []int{-578, 2147483647, 4294967295, 42}
    if v4 {
        var v6 uint16 = 42
        _ = v6
    }
    var v7 int8 = int8(v1)
    v8 := GenS1{F0: "hello", F1: 1000, F2: 1}
    var v9 int64 = (v3 | (^(v3)))
    if v4 {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        v10 := GenS1{F0: "a", F1: int16(v3), F2: 0}
        _ = v10
    } else {
        if ((3.14159 <= (0.1 + ((1e10 / (-2.25 - 3.14159)) + float64(1.5)))) && true) {
            if len(v5) > 0 {
                v5[0] = -1
            }
            _ = len(v5)
            v11 := GenS1{F0: ((((("" + "") + "\n\t") + "a") + "\n\t") + "\n\t"), F1: 1000, F2: 218}
            _ = v11
        }
    }
    if len(v5) > 0 {
        v5[0] = 128
    }
    var v12 uint = 65535
    fmt.Println(v3)
    if (true && (false || (32767 <= 830))) {
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 string = ""
            var v15 uint = v12
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        switch "\n\t" {
        case "":
            fmt.Println(v3, v3)
        case "a":
            var v17 int32 = int32(v3)
            _ = v17
        case "hello":
            v18 := GenS1{F0: "foo bar", F1: 1, F2: 220}
            _ = v18
        }
        if (!(!false)) {
            var v19 int = (fi16 & -2147483648)
            _ = v19
        } else {
            v9 -= (v9 ^ int64(fi16))
        }
        for fi20 := 0; fi20 < 6; fi20++ {
            _ = len(v2)
            if v4 {
                var v21 uint = (v12 ^ (v12 << ((v12) & 63)))
                if (226 == 1) {
                    delete(v2, "hello")
                    switch "" {
                    case "":
                        v22 := GenS2{F0: 1, F1: v7, F2: ""}
                        _ = v22
                    default:
                        var v23 int = fi20
                        _ = v23
                    }
                } else {
                    v2[""] = 758
                    var v24 uint16 = 2
                    _ = v24
                }
                var v25 int = (fi20 % (((fi16 >> (((v21 + (^(v12)))) & 63))) | 1))
                _ = v21
                _ = v25
            }
            var v26 int = (^(fi20))
            _ = fi20
            _ = v26
        }
        _ = fi16
    }
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint32 = 2147483647
        delete(v2, "")
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
