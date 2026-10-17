package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int = 4294967295
    var v2 uint = 1
    var v3 float32 = (0.0 - 1e10)
    var v4 float32 = float32(1.5)
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 615, F1: 2}
        var v7 bool = (255 > 255)
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := GenS1{F0: 42, F1: 887}
            v4 -= (v4 - v3)
            var v10 int8 = 42
            var v11 bool = (v7 || ((v2 & v2) < (v2 + v2)))
            if v11 {
                var v12 uint = v2
                v13 := GenS1{F0: 42, F1: uint16(v2)}
                _ = v12
                _ = v13
            } else {
                if (!((fi5 | v1) > (v1 ^ fi8))) {
                    var v14 float32 = v4
                    v15 := []int{}
                    switch (^(v1)) {
                    case 127:
                        fmt.Println(v14, fi5)
                    }
                    _ = v14
                    _ = v15
                } else {
                    switch "\n\t" {
                    case "":
                        v16 := []int{4294967295, (fi8 << ((v2) & 63)), fi5, (-(fi5))}
                        _ = v16
                    case "foo bar":
                        v17 := GenS1{F0: 255, F1: 111}
                        _ = v17
                    }
                    fmt.Println(v3, v1)
                    v2 -= v2
                }
                v7 = v7
                var v18 uint32 = 42
                _ = v18
            }
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        v1 += v1
    } else {
        if ((((v2 | v2) > (v2 ^ v2)) && ((v1 | -9223372036854775808) != v1)) || false) {
            v19 := []int{(v1 + (v1 >> (((v2 % (((v2 + (v2 + v2))) | 1))) & 63))), v1, v1}
            _ = v19
        }
        v1 = (v1 * v1)
    }
    if true {
        var v20 uint8 = 255
        v21 := []int{32768}
        _ = v20
        _ = v21
    }
    switch v1 {
    case 128:
        v3 += 1e10
    case 32767:
        switch (v1 * 2147483647) {
        case 1:
            v22 := []int{-9223372036854775808, (^(v1)), v1, (v1 >> ((v2) & 63))}
            _ = v22
        }
    case 127:
        var v23 float64 = (((1.5 / 1e10) - 1.5) - float64(-2.25))
        _ = v23
    default:
        if ((0.1 + v3) <= (-2.25 - v4)) {
            var v24 float32 = v3
            _ = v24
        }
    }
    var v25 float32 = (1e10 / v4)
    switch v1 {
    case 127:
        v1 += v1
    }
    v26 := []int{int(v2), v1}
    v27 := GenS1{F0: 32767, F1: 907}
    var v28 int32 = -392
    for fi29 := 0; fi29 < 9; fi29++ {
        v1 = (v1 - fi29)
        var v30 uint16 = 2
        for fi31 := 0; fi31 < 10; fi31++ {
            v2 -= (^(v2))
            _ = fi31
        }
        _ = fi29
        _ = v30
    }
    switch (^(v1)) {
    case 26:
        if true {
            v4 += float32(float64((3.14159 - (3.14159 / ((v3 + (v3 - float32(1.5))) * (v4 - v3))))))
        }
    case 464:
        fmt.Println(v3, v1, v25)
    }
    v32 := GenS2{F0: 18446744073709551615, F1: (v28 & v28)}
    var v33 int = v1
    switch (v33 + 0) {
    case 1000:
        var v34 uint32 = 0
        _ = v34
    default:
        v35 := GenS3{F0: 0}
        _ = v35
    }
    v26 = append(v26, (v1 ^ (v33 << ((v2) & 63))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
