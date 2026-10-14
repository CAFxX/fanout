package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint32 }

func main() {
    var v1 uint = 1
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 32767
    var v4 int64 = 1
    var v5 uint16 = 255
    if false {
        if ((v3 % ((0) | 1)) >= (v3 ^ uint16(v1))) {
            v1 = v1
        }
        fmt.Println(v1)
    }
    v6 := GenS1{F0: 2}
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 uint8 = 134
        switch "foo bar" {
        case "\n\t":
            v4 += 9223372036854775807
        case "foo bar":
            if false {
                v9 := GenS2{F0: 4294967295, F1: -981, F2: -92}
                fmt.Println(fi7, v4, fi7)
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 int16 = 1
                    _ = fi10
                    _ = v11
                }
                _ = v9
            }
        case "hello":
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 string = ("" + "a")
                if false {
                    var v14 string = v13
                    v15 := GenS3{F0: (v8 | 127), F1: 42, F2: 0}
                    fmt.Println(v13, v4, v8)
                    _ = v14
                    _ = v15
                } else {
                    v2[247] = 2147483648
                    if (int32(v3) != 42) {
                        if (!((v1 * 1) > (v1 >> (((v1 >> ((v1) & 63))) & 63)))) {
                            if true {
                                v5 = v5
                                v16 := []int{fi7, -936, fi7}
                                _ = v16
                            }
                            var v17 uint8 = 128
                            _ = len(v2)
                            _ = v17
                        } else {
                            var v18 uint32 = 387
                            var v19 string = ((v13 + v13) + ("" + ""))
                            _ = v18
                            _ = v19
                        }
                    }
                    v2[640] = 2147483647
                }
                _ = fi12
                _ = v13
            }
        }
        if (1 != 0) {
            for fi20 := 0; fi20 < 8; fi20++ {
                v21 := []int{542, fi20}
                switch (fi20 | fi20) {
                case 311:
                    fmt.Println(v5)
                }
                v21 = append(v21, -9223372036854775808)
                _ = fi20
                _ = v21
            }
            v22 := v2[1]
            _ = v22
        }
        _ = fi7
        _ = v8
    }
    if false {
        switch "\n\t" {
        case "foo bar":
            v1 -= v1
        default:
            v1 = (v1 >> ((v1) & 63))
        }
        fmt.Println(v5, v3)
    } else {
        v5 = 853
        v4 -= (v4 ^ v4)
        fmt.Println(v5, v1)
    }
    switch int(v5) {
    case 128:
        var v23 uint = (v1 / ((18446744073709551615) | 1))
        _ = v23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        var v25 int16 = 32767
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS3{F0: 96, F1: 1, F2: uint32(v5)}
            var v28 bool = true
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi24
        _ = v25
    }
    var v29 string = "a"
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
