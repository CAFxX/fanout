package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 int16; F2 uint64 }
type GenS4 struct { F0 uint8; F1 uint32 }
type GenS5 struct { F0 int8; F1 int; F2 int8 }

func main() {
    var v1 uint64 = 609
    v2 := make(map[uint32]int)
    var v3 bool = (1.5 <= (3.14159 / ((float32(0.1) + 0.0) - 3.14159)))
    var v4 int32 = -112
    var v5 int8 = 0
    if v3 {
        v6 := GenS1{F0: (^(v1))}
        switch -986 {
        case 42:
            v4 = v4
        }
        fmt.Println(v1)
        _ = v6
    }
    var v7 uint = uint(v1)
    switch "hello" {
    case "hello":
        if (!false) {
            v1 += v1
        }
    case "\n\t":
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := []int{fi8, fi8}
            v9 = append(v9, 42)
            var v10 uint8 = 96
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    _ = len(v2)
    switch "" {
    case "a":
        if (-1 > -1) {
            for fi11 := 0; fi11 < 10; fi11++ {
                if ((!v3) || (42 <= 4294967295)) {
                    if (v4 == v4) {
                        var v12 uint64 = v1
                        switch (fi11 >> (((^(v7))) & 63)) {
                        case 4294967295:
                            v13 := GenS2{F0: -763}
                            _ = v13
                        case 9223372036854775807:
                            v14 := GenS3{F0: 108, F1: -32768, F2: (v1 << (((^(v7))) & 63))}
                            _ = v14
                        case 127:
                            switch (-(fi11)) {
                            case 750:
                                for fi15 := 0; fi15 < 1; fi15++ {
                                    if (int64(v7) != 0) {
                                        v16 := v2[2]
                                        fmt.Println(v12)
                                        _ = v16
                                    }
                                    v17 := []int{fi11, fi11, (^(fi15))}
                                    if (v3 && (v7 != (^(v7)))) {
                                        if len(v17) >= 2 {
                                            v17 = v17[1:len(v17)]
                                        }
                                    }
                                    fmt.Println(v4, v3, v3)
                                    _ = fi15
                                    _ = v17
                                }
                            default:
                                var v18 uint = (v7 >> (((v7 * (v7 / (((v7 * v7)) | 1)))) & 63))
                                _ = v18
                            }
                        default:
                            v19 := GenS1{F0: 1}
                            _ = v19
                        }
                        var v20 uint64 = v1
                        _ = v12
                        _ = v20
                    }
                    switch fi11 {
                    case 42:
                        v21 := []int{42, 398}
                        _ = v21
                    case 32767:
                        if v3 {
                            if false {
                                v22 := GenS1{F0: 2}
                                if (!false) {
                                    fmt.Println(fi11, fi11, fi11)
                                }
                                _ = v22
                            }
                        } else {
                            for fi23 := 0; fi23 < 4; fi23++ {
                                fmt.Println(fi23, v3)
                                _ = fi23
                            }
                            v2[4294967295] = (fi11 >> (((^(v7))) & 63))
                            switch "a" {
                            case "\n\t":
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    v25 := GenS4{F0: 116, F1: 249}
                                    fmt.Println(fi24, v5)
                                    _ = fi24
                                    _ = v25
                                }
                            }
                        }
                    case -427:
                        var v26 uint16 = 741
                        _ = v26
                    default:
                        if true {
                            v27 := GenS3{F0: 42, F1: 42, F2: 2}
                            v7 = v7
                            v28 := []int{(-(fi11))}
                            _ = v27
                            _ = v28
                        }
                    }
                }
                var v29 uint64 = 4294967295
                _ = fi11
                _ = v29
            }
            v30 := []int{}
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v30
        }
    default:
        v31 := []int{46, int(v7), 32768, 958}
        _ = v31
    }
    v1 = 1
    v32 := GenS5{F0: (v5 - v5), F1: -322, F2: v5}
    fmt.Println(v3)
    if (true || (v3 || v3)) {
        v33 := v2[375]
        _ = v33
    }
    var v34 float64 = (0.0 * -2.25)
    var v35 int16 = int16(v5)
    for fi36 := 0; fi36 < 13; fi36++ {
        var v37 float32 = 1e10
        if true {
            var v38 float64 = 1e10
            _ = v38
        }
        v1 += (v1 + v1)
        _ = fi36
        _ = v37
    }
    var v39 string = "\n\t"
    v39 = ("foo bar" + "a")
    v2[2] = int(v1)
    var v40 float64 = ((0.0 * 1e10) / -2.25)
    v41 := GenS3{F0: 1, F1: (^(v35)), F2: (v1 ^ 2)}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v34
    _ = v35
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v7
}
