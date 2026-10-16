package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint8; F1 uint32; F2 string }

func main() {
    var v1 int8 = 1
    var v2 int64 = 196
    var v3 bool = (float64((float32((0.1 + 0.1)) + (-2.25 / -2.25))) < -2.25)
    v4 := make(map[int]string)
    var v5 uint64 = 4294967295
    if v3 {
        v5 -= (v5 | 2)
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 int32 = 0
            if v3 {
                for fi8 := 0; fi8 < 15; fi8++ {
                    v7 = int32(v1)
                    _ = fi8
                }
                var v9 float32 = (1e10 / (0.1 - 1e10))
                if (int(v7) == (fi6 & fi6)) {
                    var v10 int8 = v1
                    fmt.Println(v9, v5)
                    _ = v10
                } else {
                    var v11 float64 = float64(-2.25)
                    _ = v11
                }
                _ = v9
            } else {
                var v12 uint16 = 255
                var v13 int32 = 47
                _ = v12
                _ = v13
            }
            v5 += v5
            _ = fi6
            _ = v7
        }
    }
    v1 += v1
    v14 := v4[32768]
    var v15 int = 965
    v4[v15] = (v14 + (v14 + v14))
    fmt.Println(v1, v3, v15)
    v16 := GenS1{F0: (v14 + "a")}
    v14 = v14
    v17 := GenS2{F0: 127, F1: 4294967295, F2: "foo bar"}
    var v18 uint = 114
    if (v2 != 42) {
        v19 := GenS1{F0: ""}
        v20 := GenS2{F0: 42, F1: 65535, F2: ((v14 + v14) + "hello")}
        var v21 uint = v18
        _ = v19
        _ = v20
        _ = v21
    } else {
        if (uint16(v18) < 255) {
            v18 -= v18
            var v22 uint16 = 1
            var v23 uint = v18
            _ = v22
            _ = v23
        } else {
            var v24 int8 = (v1 / (((-(v1))) | 1))
            _ = v24
        }
        for fi25 := 0; fi25 < 13; fi25++ {
            var v26 float32 = 0.1
            fmt.Println(v18, v2, v5)
            _ = fi25
            _ = v26
        }
        if ((((v14 + v14) + v14) <= (v14 + v14)) && (2 <= v18)) {
            for fi27 := 0; fi27 < 9; fi27++ {
                fmt.Println(v3)
                _ = fi27
            }
            v2 += (v2 / (((v2 % (((v2 >> ((uint(v5)) & 63))) | 1))) | 1))
            var v28 int = (v15 - v15)
            _ = v28
        }
    }
    v4[(^(v15))] = v14
    fmt.Println(v18, v2, v2)
    v4[(v15 | v15)] = "\n\t"
    if true {
        v14 = (v14 + "foo bar")
        switch ("" + (("a" + v14) + "foo bar")) {
        case "hello":
            var v29 uint64 = (v5 + (v5 % (((v5 ^ 18446744073709551615)) | 1)))
            _ = v29
        }
        var v30 int8 = (v1 ^ int8(v2))
        _ = v30
    } else {
        var v31 string = v14
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
